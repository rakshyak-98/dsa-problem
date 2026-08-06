#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

echo "==> Checking prerequisites"
command -v go >/dev/null || { echo "Go is required (1.22+). Install from https://go.dev/dl/"; exit 1; }
go version

echo
echo "==> Initializing backend interview drills"
(go run -C bin/study_backend . -- --setup)

echo
echo "==> Initializing write drills (7 daily + 2 bonus)"
(go run -C bin/study_play . -- --setup)

echo
echo "==> Verifying helpers"
(go run -C bin/study_play . -- --catalog >/dev/null)
(go run -C bin/study_code . -- --catalog >/dev/null)
(go build -o /dev/null .)
(go build -C drills/write/core5 -o /dev/null .)
(go build -C drills/write/variants -o /dev/null .)

echo
echo "==> Running coverage gate (80%)"
"$ROOT/bin/scripts/test-coverage.sh"

echo
echo "Setup complete."
echo
echo "Next steps:"
echo "  START HERE:                 open doc/drills.md"
echo "  unified daily:              go run ."
echo "  backend interview prep:     go run -C bin/study_backend . -- --cram"
echo "  write drills CLI:           go run -C bin/study_play ."
echo "  Core 5:                     go run -C drills/write/core5 ."
echo "  read drills CLI:            go run -C bin/study_code ."
echo "  study tracker (browser):    open drills/tracker/study_tracker.html"
echo "  visualizer (browser):       open reference/visualizer/index.html"
echo "  problem index:              reference/problems/CATEGORIES.md"
