#!/usr/bin/env bash
set -euo pipefail

: "${OPENAI_API_KEY:?OPENAI_API_KEY is required}"

# Path to your prompt markdown; change via env if needed.
CODEX_PROMPT="${CODEX_PROMPT:-prompts/gke-latest.md}"

if [ ! -f "$CODEX_PROMPT" ]; then
  echo "ERROR: Prompt file not found: $CODEX_PROMPT"
  exit 2
fi

echo "[codex] installing @openai/codex…"
npm install -g @openai/codex >/dev/null 2>&1 || npm install -g @openai/codex
codex --version || true

# Ensure a clean workspace view for diffing
git reset --hard
git clean -fd

# Run Codex in NON-INTERACTIVE mode with web search and auto-approvals.
# Pass the prompt CONTENTS, not the filename, to avoid TUI + /dev/tty.
PROMPT_TEXT="$(cat "$CODEX_PROMPT")"

echo "[codex] exec (non-interactive) with search + full-auto…"
set +e
codex exec --full-auto --search "$PROMPT_TEXT"
RUN_CODE=$?
set -e

# Optional formatting
command -v gofmt >/dev/null 2>&1 && gofmt -s -w .

# Expose whether files changed for the PR step
if git diff --quiet; then
  echo "changed=0" >> "$GITHUB_OUTPUT"
  echo "[codex] no changes"
  # If Codex errored and produced no changes, fail the job
  exit "$RUN_CODE"
else
  echo "changed=1" >> "$GITHUB_OUTPUT"
  echo "[codex] changes detected"
  # Success; let the PR step commit & open the PR
  exit 0
fi
