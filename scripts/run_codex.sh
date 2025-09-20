#!/usr/bin/env bash
set -euo pipefail

# Inputs (override via env if your Codex install/run differ)
: "${CODEX_INSTALL:=npm i -g @codex-ai/cli}"
: "${CODEX_CMD:=codex run --prompt prompts/gke-latest.md}"
: "${OPENAI_API_KEY:?OPENAI_API_KEY is required}"

# Optional: Node is present via setup-node in the workflow
shopt -s expand_aliases

# Install Codex CLI
echo "[codex] installing…"
$CODEX_INSTALL >/dev/null

# Ensure clean tree
git reset --hard
git clean -fd

# Run Codex with your prompt
echo "[codex] running prompt…"
eval "$CODEX_CMD"

# Format (optional)
if command -v gofmt >/dev/null 2>&1; then
  gofmt -s -w .
fi

# Did anything change?
if git diff --quiet; then
  echo "changed=0" >> "$GITHUB_OUTPUT"
  exit 0
else
  echo "changed=1" >> "$GITHUB_OUTPUT"
  # leave changes staged/unstaged for the PR action to commit
  exit 0
fi
