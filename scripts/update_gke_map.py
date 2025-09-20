#!/usr/bin/env python3
import os, re, sys, json, subprocess, requests
from bs4 import BeautifulSoup

RELEASE_NOTES_URL = "https://cloud.google.com/kubernetes-engine/docs/release-notes"
GKE_GO_PATH = "codecamp/pkg/project/gke.go"
PROMPT_PATH = "prompts/gke-latest.md"
OPENAI_MODEL = os.getenv("OPENAI_MODEL", "gpt-4.1-mini")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")

def top_existing():
    rx = re.compile(r'^\s*Version:\s*"(\d{4}-R\d{2})"', re.MULTILINE)
    with open(GKE_GO_PATH, "r", encoding="utf-8") as f:
        m = rx.search(f.read())
    return m.group(1) if m else None

def fetch_release_notes():
    r = requests.get(RELEASE_NOTES_URL, timeout=60)
    r.raise_for_status()
    return r.text

def sections_above_top(html, top):
    soup = BeautifulSoup(html, "html.parser")
    headers = soup.select('h4[id$="_version_updates"][data-text*="Version updates"]')
    rid_rx = re.compile(r"\((\d{4}-R\d{2})\)\s*Version updates")
    ordered = []
    for h in headers:
        dt = h.get("data-text", "") or h.get_text(strip=True)
        m = rid_rx.search(dt)
        if not m: 
            continue
        ordered.append((m.group(1), h))

    # Find index of TOP_EXISTING on page
    idx = None
    if top:
        for i, (rid, _) in enumerate(ordered):
            if rid == top:
                idx = i
                break

    # New = everything above TOP on the page; if TOP not found, consider all
    new_hdrs = ordered[:idx] if idx is not None else ordered
    out = []
    for rid, h in new_hdrs:
        # find tabset near this header
        tabset = None
        for sib in h.find_all_next():
            if getattr(sib, "name", "") == "div" and "devsite-tabs" in (sib.get("class") or []):
                tabset = sib
                break
        if not tabset:
            continue
        # locate Stable tab -> aria-controls
        stable_tab = None
        for a in tabset.select('a[role="tab"], a[role="button"]'):
            if a.get_text(strip=True).lower() == "stable":
                stable_tab = a
                break
        panel = None
        if stable_tab:
            pid = stable_tab.get("aria-controls") or ""
            if pid:
                panel = tabset.select_one(f'section[role="tabpanel"]#{pid}')
        if not panel:
            panel = tabset.select_one('section#tabpanel-stable-channel[role="tabpanel"]')
        if not panel:
            continue

        # anchor for provenance (closest previous h2/h3 id)
        anchor_el = h.find_previous(["h2", "h3"])
        anchor = f'#{anchor_el.get("id")}' if anchor_el and anchor_el.get("id") else ""
        out.append({"rid": rid, "stable_panel_html": str(panel), "source": RELEASE_NOTES_URL + anchor})
    return out

def call_openai(prompt_text):
    import requests
    url = "https://api.openai.com/v1/responses"
    headers = {"Authorization": f"Bearer {OPENAI_API_KEY}", "Content-Type": "application/json"}
    body = {"model": OPENAI_MODEL, "input": prompt_text, "temperature": 0}
    r = requests.post(url, headers=headers, json=body, timeout=120)
    r.raise_for_status()
    data = r.json()
    # responses API: flatten text
    text = data.get("output", {}).get("text") or data.get("output_text") or ""
    return text.strip()

def merge_entries(entries):
    with open(GKE_GO_PATH, "r", encoding="utf-8") as f:
        content = f.read()
    anchor = "GKEProjectReleases"
    idx = content.find(anchor)
    if idx < 0:
        print("Could not find GKEProjectReleases in file", file=sys.stderr)
        sys.exit(2)
    array_start = content.find("{", idx)
    if array_start < 0:
        print("Could not find array opening brace", file=sys.stderr)
        sys.exit(2)

    def semkey(v):
        m = re.search(r"(\d+)\.(\d+)\.(\d+)", v)
        return tuple(int(x) for x in m.groups()) if m else (0,0,0)

    blocks = []
    for e in entries:
        uniq = sorted(sorted(set(e["RelatedProjectReleases"])), key=semkey)
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

    with open(PROMPT_PATH, "r", encoding="utf-8") as f:
        base_prompt = f.read()

    # Build final prompt sent to the model
    user_payload = {
        "instruction": "Extract ALL Stable-channel Kubernetes versions per release (defaults/newly available/removed). "
                       "Normalize to kube@X.Y.Z (strip -gke*, -autopilot*, +cos*). De-duplicate and sort by full SemVer ascending. "
                       "Return ONLY JSON in the schema below.",
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
    try:
        data = json.loads(raw)
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
