# Drill 04 — Binary Search

## binarySearch
- **Template:** lo, hi := 0, n-1; for lo <= hi; mid; shrink

## searchInsert / lower bound
- **Return lo** when loop ends (first index ≥ target)
- **Bug:** using hi < lo condition wrong

## findMinRotated
- **Pattern:** compare mid with hi; shrink toward unsorted half
- **Bug:** using lo < hi when duplicates exist (this drill assumes distinct)

## minEatingSpeed
- **Trigger:** "minimum rate / capacity that still works" — the input is not sorted
- **Pattern:** binary search the *answer space* (1..max pile), not an index
- **Requires** `feasible(x)` monotonic: if x works, everything above x works
- **Bug:** integer division when counting hours — use `(p + speed - 1) / speed`
