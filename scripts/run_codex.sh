#!/usr/bin/env bash
set -euo pipefail

# --- config ---
# Set via GitHub Secrets
: "${OPENAI_API_KEY:?OPENAI_API_KEY is required}"

# Path to your prompt file (override in workflow with CODEX_PROMPT if needed)
CODEX_PROMPT="${CODEX_PROMPT:-prompts/gke-latest.md}"

# --- sanity checks ---
if [ ! -f "$CODEX_PROMPT" ]; then
  echo "ERROR: Prompt file not found: $CODEX_PROMPT"
  echo "Tip: put your prompt at prompts/gke-latest.md or set CODEX_PROMPT env."
  exit 2
fi

# --- install codex CLI ---
echo "[codex] installing @openai/codex…"
npm install -g @openai/codex

# show version (best effort)
if command -v codex >/dev/null 2>&1; then
  codex --version || true
else
  npx @openai/codex --version || true
fi

# ensure clean tree (optional, keeps the working dir predictable)
git reset --hard
git clean -fd

# --- run codex ---
echo "[codex] running with prompt: $CODEX_PROMPT"
set +e
if command -v codex >/dev/null 2>&1; then
  codex run --prompt "$CODEX_PROMPT"
  RUN_CODE=$?
else
  npx @openai/codex run --prompt "$CODEX_PROMPT"
  RUN_CODE=$?
fi
set -e

# optional: format Go if available
if command -v gofmt >/dev/null 2>&1; then
  gofmt -s -w .
fi

# --- diff detection for PR ---
if git diff --quiet; then
  echo "changed=0" >> "$GITHUB_OUTPUT"
  echo "[codex] no changes"
  exit "$RUN_CODE"
else
  echo "changed=1" >> "$GITHUB_OUTPUT"
  echo "[codex] changes detected"
  exit 0
fi
