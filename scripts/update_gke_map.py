#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import re
import sys
import json
import time
import subprocess
import requests
from bs4 import BeautifulSoup

RELEASE_NOTES_URL = "https://cloud.google.com/kubernetes-engine/docs/release-notes"
GKE_GO_PATH = "pkg/project/gke.go"
PROMPT_PATH = "prompts/gke-latest.md"
OPENAI_MODEL = os.getenv("OPENAI_MODEL", "gpt-o3")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")

def top_existing():
    rx = re.compile(r'^\s*Version:\s*"(\d{4}-R\d{2})"', re.MULTILINE)
    with open(GKE_GO_PATH, "r", encoding="utf-8") as f:
        m = rx.search(f.read())
    return m.group(1) if m else None

def fetch_release_notes():
    headers = {"User-Agent": "chkk-gke-bot/1.0 (+https://github.com/your-org/your-repo)"}
    last_exc = None
    for attempt in range(3):
        try:
            r = requests.get(RELEASE_NOTES_URL, headers=headers, timeout=60)
            if r.status_code >= 500:
                time.sleep(1 + attempt)
                continue
            r.raise_for_status()
            return r.text
        except Exception as e:
            last_exc = e
            time.sleep(1 + attempt)
    if last_exc:
        raise last_exc
    raise RuntimeError("Unable to fetch release notes")

def sections_above_top(html, top):
    from bs4 import Tag
    soup = BeautifulSoup(html, "html.parser")

    headers = soup.find_all(
        lambda t: t.name == "h4"
        and (t.get("id", "").endswith("-version-updates") or t.get("id", "").endswith("_version_updates"))
        and (t.get("data-text", "") or t.get_text(" ", strip=True)).strip().endswith("Version updates")
    )
    rid_rx = re.compile(r"\((\d{4}-R\d{2})\)\s*Version updates", re.I)

    ordered = []
    for h in headers:
        label = h.get("data-text", "") or h.get_text(" ", strip=True)
        m = rid_rx.search(label)
        if m:
            ordered.append((m.group(1), h))

    idx = None
    if top:
        for i, (rid, _) in enumerate(ordered):
            if rid == top:
                idx = i
                break

    new_hdrs = ordered[:idx] if idx is not None else ordered

    out = []
    for rid, h in new_hdrs:
        tabset = None
        for sib in h.next_siblings:
            if isinstance(sib, Tag) and sib.name == "div" and "devsite-tabs" in (sib.get("class") or []):
                tabset = sib
                break
            if isinstance(sib, Tag) and sib.name in ("h3", "h4"):
                if rid_rx.search(sib.get_text(" ", strip=True) or ""):
                    break

        stable_panel = None
        if tabset:
            stable_tab = None
            for a in tabset.select('a[role="tab"], a[role="button"]'):
                if a.get_text(strip=True).lower() == "stable":
                    stable_tab = a
                    break
            if stable_tab:
                pid = stable_tab.get("aria-controls") or ""
                if pid:
                    stable_panel = tabset.select_one(f'section[role="tabpanel"]#{pid}')
            if not stable_panel:
                stable_panel = tabset.select_one('section#tabpanel-stable-channel[role="tabpanel"]')

        if not stable_panel:
            block_nodes = []
            for sib in h.next_siblings:
                if isinstance(sib, Tag) and sib.name in ("h3", "h4") and rid_rx.search(sib.get_text(" ", strip=True) or ""):
                    break
                block_nodes.append(sib)
            block_html = "".join(str(n) for n in block_nodes)
            block = BeautifulSoup(block_html, "html.parser")
            candidates = [
                el for el in block.find_all(["h5", "h6", "strong", "p", "li"])
                if "stable" in el.get_text(" ", strip=True).lower()
            ]
            stable_panel = candidates[0] if candidates else None

        if not stable_panel:
            continue

        anc = ""
        parent_anchor = h.find_previous(["h2", "h3"])
        if parent_anchor and parent_anchor.get("id"):
            anc = f'#{parent_anchor.get("id")}'

        out.append({
            "rid": rid,
            "stable_panel_html": str(stable_panel),
            "source": RELEASE_NOTES_URL + anc
        })

    print(f"[debug] TOP={top} | headers_on_page={len(ordered)} | idx_of_top={idx} | new_above_top={len(new_hdrs)} | extracted={len(out)} | R_ids={[x['rid'] for x in out]}")
    return out

def _parse_json_lenient(text: str):
    try:
        return json.loads(text)
    except Exception:
        s = text.strip()
        if s.startswith("```"):
            s = re.sub(r"^```(?:json)?\s*", "", s, flags=re.DOTALL)
            s = re.sub(r"\s*```$", "", s, flags=re.DOTALL)
        if "{" in s and "}" in s:
            s2 = s[s.find("{"):s.rfind("}") + 1]
            return json.loads(s2)
        raise

def call_openai(prompt_text: str) -> str:
    headers = {
        "Authorization": f"Bearer {OPENAI_API_KEY}",
        "Content-Type": "application/json",
    }
    system_msg = "You are a precise data extractor. Return ONLY strict JSON with no prose and no code fences."

    chat_url = "https://api.openai.com/v1/chat/completions"
    chat_body_json = {
        "model": OPENAI_MODEL,
        "messages": [
            {"role": "system", "content": system_msg},
            {"role": "user", "content": prompt_text},
        ],
        "response_format": {"type": "json_object"},
    }
    try:
        r = requests.post(chat_url, headers=headers, json=chat_body_json, timeout=180)
        if r.status_code == 200:
            data = r.json()
            text = data["choices"][0]["message"]["content"]
            return (text or "").strip()
        else:
            print(f"[openai-chat/json] {r.status_code} {r.text}", file=sys.stderr)
    except Exception as e:
        print(f"[openai-chat/json] exception: {e}", file=sys.stderr)

    chat_body_plain = {
        "model": OPENAI_MODEL,
        "messages": [
            {"role": "system", "content": system_msg},
            {"role": "user", "content": prompt_text},
        ],
    }
    try:
        r = requests.post(chat_url, headers=headers, json=chat_body_plain, timeout=180)
        if r.status_code == 200:
            data = r.json()
            text = data["choices"][0]["message"]["content"]
            return (text or "").strip()
        else:
            print(f"[openai-chat] {r.status_code} {r.text}", file=sys.stderr)
    except Exception as e:
        print(f"[openai-chat] exception: {e}", file=sys.stderr)

    resp_url = "https://api.openai.com/v1/responses"
    resp_body = {"model": OPENAI_MODEL, "input": prompt_text}
    r2 = requests.post(resp_url, headers=headers, json=resp_body, timeout=180)
    if r2.status_code >= 400:
        print(f"[openai-responses] {r2.status_code} {r2.text}", file=sys.stderr)
    r2.raise_for_status()
    data = r2.json()
    text = data.get("output_text")
    if not text:
        try:
            text = data["output"][0]["content"][0]["text"]
        except Exception:
            try:
                text = data["choices"][0]["message"]["content"]
            except Exception:
                text = ""
    return (text or "").strip()

def merge_entries(entries):
    with open(GKE_GO_PATH, "r", encoding="utf-8") as f:
        content = f.read()

    m = re.search(r"(?:var\s+)?GKEProjectReleases\s*=\s*\[\]\s*ProjectRelease\s*{", content)
    if not m:
        print("Could not find GKEProjectReleases array", file=sys.stderr)
        sys.exit(2)

    array_start = m.end() - 1

    def semkey(v):
        mm = re.search(r"(\d+)\.(\d+)\.(\d+)", v)
        return tuple(int(x) for x in mm.groups()) if mm else (0, 0, 0)

    blocks = []
    for e in entries:
        uniq = sorted(set(e["RelatedProjectReleases"]), key=semkey)
        inner = ",\n            ".join(f"\"{v}\"" for v in uniq)
        block = f"""
    {{
        Project: GKE.ID,
        Version: "{e["Version"]}",
        RelatedProjectReleases: []string{{
            {inner}
        }},
    }},
    // source: {e.get("source","")}
"""
        blocks.append(block)

    new_content = content[:array_start+1] + "".join(blocks) + content[array_start+1:]
    if new_content == content:
        return False

    with open(GKE_GO_PATH, "w", encoding="utf-8") as f:
        f.write(new_content)

    try:
        subprocess.run(["gofmt", "-s", "-w", "."], check=False)
    except Exception:
        pass
    return True

def main():
    if not OPENAI_API_KEY:
        print("OPENAI_API_KEY not set", file=sys.stderr)
        sys.exit(2)

    top = top_existing()
    html = fetch_release_notes()
    sections = sections_above_top(html, top)
    if not sections:
        print("No new releases above TOP; exiting.")
        sys.exit(0)

    try:
        with open(PROMPT_PATH, "r", encoding="utf-8") as f:
            base_prompt = f.read()
    except Exception as e:
        print(f"Failed to read prompt file {PROMPT_PATH}: {e}", file=sys.stderr)
        sys.exit(2)

    user_payload = {
        "instruction": (
            "Extract ALL Stable-channel Kubernetes versions per release "
            "(defaults/newly available/removed). Normalize to kube@X.Y.Z "
            "(strip -gke*, -autopilot*, +cos*). De-duplicate and sort by "
            "full SemVer ascending. Return ONLY JSON in the schema below."
        ),
        "schema": {
            "entries": [{
                "Version": "YYYY-RXX",
                "RelatedProjectReleases": ["kube@1.30.14"],
                "source": "<url#anchor>"
            }]
        },
        "topExisting": top,
        "releases": sections
    }
    final_prompt = base_prompt + "\n\nContext:\n" + json.dumps(user_payload, ensure_ascii=False)

    raw = call_openai(final_prompt)
    if not raw:
        print("Model returned empty output.", file=sys.stderr)
        sys.exit(2)

    try:
        data = _parse_json_lenient(raw)
        entries = data.get("entries", [])
    except Exception:
        print("Model did not return valid JSON.\nRaw:\n", raw, file=sys.stderr)
        sys.exit(2)

    if not entries:
        print("No entries to add; exiting.")
        sys.exit(0)

    changed = merge_entries(entries)
    sys.exit(10 if changed else 0)

if __name__ == "__main__":
    main()
