#!/usr/bin/env bash
set -euo pipefail

: "${OPENAI_API_KEY:?OPENAI_API_KEY is required}"

# Defaults (override in workflow env if needed)
: "${CODEX_INSTALL:=npm install -g @openai/codex}"
: "${CODEX_BIN:=codex}"                             # will try this first
: "${CODEX_NPX:=npx @openai/codex}"                 # fallback if global not found
: "${CODEX_CMD:=run --prompt prompts/gke-latest.md}"# codex subcommand & args

echo "[codex] installing…"
# install quietly but keep errors visible
$CODEX_INSTALL

# Show version (best effort)
if command -v "$CODEX_BIN" >/dev/null 2>&1; then
  "$CODEX_BIN" --version || true
else
  $CODEX_NPX --version || true
fi

# Ensure clean tree before running (optional)
git reset --hard
git clean -fd

echo "[codex] running…"
set +e
if command -v "$CODEX_BIN" >/dev/null 2>&1; then
  "$CODEX_BIN" $CODEX_CMD
  CODE=$?
else
  $CODEX_NPX $CODEX_CMD
  CODE=$?
fi
set -e

# Optional: gofmt if present
if command -v gofmt >/dev/null 2>&1; then
  gofmt -s -w .
fi

# Detect diff to decide if we should open a PR
if git diff --quiet; then
  echo "changed=0" >> "$GITHUB_OUTPUT"
  exit $CODE
else
  echo "changed=1" >> "$GITHUB_OUTPUT"
  exit 0
fi
