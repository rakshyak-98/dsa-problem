#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
THRESHOLD="${COVERAGE_THRESHOLD:-80}"

echo "==> Running tests with coverage (threshold: ${THRESHOLD}%)"

OUT=$(mktemp)
trap 'rm -f "$OUT"' EXIT

(cd "$ROOT/study_play" && go test . ./_support/templates -covermode=atomic) | tee -a "$OUT"
(cd "$ROOT/daily" && go test . -covermode=atomic) | tee -a "$OUT"
(cd "$ROOT/study_code" && go test . -covermode=atomic) | tee -a "$OUT"

echo
echo "==> Coverage by package"
grep 'coverage:' "$OUT" || true

TOTAL=$(grep -oE 'coverage: [0-9.]+%' "$OUT" | awk -F'[: %]+' '{sum+=$2; n++} END { if (n>0) printf "%.1f", sum/n; else print "0" }')

echo
echo "Average coverage: ${TOTAL}%"

awk -v total="$TOTAL" -v threshold="$THRESHOLD" 'BEGIN {
  if (total + 0 < threshold + 0) {
    printf "FAIL: average coverage %.1f%% is below %d%% threshold\n", total, threshold
    exit 1
  }
  printf "PASS: average coverage %.1f%% meets %d%% threshold\n", total, threshold
}'
