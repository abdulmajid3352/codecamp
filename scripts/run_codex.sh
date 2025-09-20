#!/usr/bin/env bash
set -euo pipefail

: "${OPENAI_API_KEY:?OPENAI_API_KEY is required}"

CODEX_PROMPT="${CODEX_PROMPT:-prompts/gke-latest.md}"

if [ ! -f "$CODEX_PROMPT" ]; then
  echo "ERROR: Prompt file not found: $CODEX_PROMPT"
  exit 2
fi

echo "[codex] installing @openai/codex…"
npm install -g @openai/codex >/dev/null 2>&1 || npm install -g @openai/codex

if command -v codex >/dev/null 2>&1; then
  codex --version || true
else
  npx @openai/codex --version || true
fi

git reset --hard
git clean -fd

echo "[codex] running with prompt: $CODEX_PROMPT"
set +e
if command -v codex >/dev/null 2>&1; then
  codex "$CODEX_PROMPT"
  RUN_CODE=$?
else
  npx @openai/codex "$CODEX_PROMPT"
  RUN_CODE=$?
fi
set -e

if command -v gofmt >/dev/null 2>&1; then
  gofmt -s -w .
fi

if git diff --quiet; then
  echo "changed=0" >> "$GITHUB_OUTPUT"
  echo "[codex] no changes"
  exit "$RUN_CODE"   # fail the job if Codex errored
else
  echo "changed=1" >> "$GITHUB_OUTPUT"
  echo "[codex] changes detected"
  exit 0
fi
