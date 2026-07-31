# Drill 04 — Binary Search

## binarySearch
- **Template:** lo, hi := 0, n-1; for lo <= hi; mid; shrink

## searchInsert / lower bound
- **Return lo** when loop ends (first index ≥ target)
- **Bug:** using hi < lo condition wrong

## findMinRotated
- **Pattern:** compare mid with hi; shrink toward unsorted half
- **Bug:** using lo < hi when duplicates exist (this drill assumes distinct)
