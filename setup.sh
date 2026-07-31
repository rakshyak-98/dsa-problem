#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

echo "==> Checking prerequisites"
command -v go >/dev/null || { echo "Go is required (1.22+). Install from https://go.dev/dl/"; exit 1; }
go version

echo
echo "==> Initializing study_play reflex drills"
(cd study_play && go run . -- --setup)

echo
echo "==> Verifying study_play helper"
(cd study_play && go run . -- --catalog >/dev/null)

echo
echo "==> Verifying study_code helper"
(cd study_code && go run . -- --catalog >/dev/null)

echo
echo "Setup complete."
echo
echo "Next steps:"
echo "  study_play (write drills):  cd study_play && go run ."
echo "  study_code (read drills):   cd study_code && go run ."
echo "  study tracker (browser):    open study_play/study_tracker.html"
echo "  visualizer (browser):       open visualizer/index.html"
echo "  problem index:              problems/CATEGORIES.md"
