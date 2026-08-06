#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

echo "==> Checking prerequisites"
command -v go >/dev/null || { echo "Go is required (1.22+). Install from https://go.dev/dl/"; exit 1; }
go version

echo
echo "==> Initializing backend interview drills"
(cd study_backend && go run . -- --setup)

echo
echo "==> Initializing write drills (7 daily + 2 bonus)"
(cd study_play && go run . -- --setup)

echo
echo "==> Verifying helpers"
(cd study_play && go run . -- --catalog >/dev/null)
(cd study_code && go run . -- --catalog >/dev/null)
(cd daily && go build -o /dev/null .)
(cd study_play/practice/write/core5 && go build -o /dev/null .)
(cd study_play/practice/write/variants && go build -o /dev/null .)

echo
echo "==> Running coverage gate (80%)"
"$ROOT/tools/scripts/test-coverage.sh"

echo
echo "Setup complete."
echo
echo "Next steps:"
echo "  backend interview prep:     cd study_backend && go run . -- --cram"
echo "  START HERE:                 open drills/README.md"
echo "  unified daily:              cd daily && go run ."
echo "  write drills CLI:           cd study_play && go run ."
echo "  Core 5:                     go run -C study_play/practice/write/core5 ."
echo "  read drills CLI:            cd study_code && go run ."
echo "  study tracker (browser):    open drills/tracker/study_tracker.html"
echo "  visualizer (browser):       open visualizer/index.html"
echo "  problem index:              problems/CATEGORIES.md"
