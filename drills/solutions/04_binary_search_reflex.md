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

## quickSelect
- **Trigger:** "kth largest / smallest" with an O(n) average requirement
- **Pattern:** partition; `p == k-1` is the answer, else recurse into the one side holding rank k
- **Same discard-a-half idea as binary search**, on an unsorted array

## fastPow
- **Trigger:** large exponent, or "compute x^n without the library call"
- **Pattern:** fold n bit by bit — square x each step, multiply in on a set bit
- **Bug:** forgetting `n < 0 → x = 1/x, n = -n`

## gcd
- **Trigger:** reduce a fraction, tile evenly, `lcm(a,b) = a/gcd*b`
- **Pattern:** Euclid — `for b != 0 { a, b = b, a%b }`; return `a`
