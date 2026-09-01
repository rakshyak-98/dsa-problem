# Variants — Solutions Reference

The stretch set: every one reuses a move from the weekday rotation, but the
statement does not hand you the pattern. Kadane, the variable window, and
product-except-self used to live here; they are now drills 01 and 03.

## twoSumSorted
- **Pattern:** L/R pointers; if sum < target move left++, else right--
- **Different from the hash version** because the array is sorted — O(1) space
- **Step up:** 3Sum (fix one index, two-pointer the rest)

## mergeIntervals
- **Pattern:** sort by start, then sweep keeping one open interval
- **Merge when** `next.start <= open.end`; extend `open.end = max(open.end, next.end)`
- **Trap:** touching intervals (`[1,4]`,`[4,5]`) do merge — the test is `<=`, not `<`
- **Step up:** insert interval into an already-sorted list without re-sorting

## spiralOrder
- **Pattern:** four moving bounds — top, bottom, left, right — closing inward
- **After each pass** shrink that bound and check `top <= bottom && left <= right`
  before the next direction, or a single row gets emitted twice
- **Step up:** rotate image 90° in place (transpose, then reverse each row)

## longestConsecutive
- **Pattern:** set membership, not sorting — sorting is O(n log n) and misses the point
- **Only start counting** at `x` when `x-1` is absent, so each run is walked once
- **Step up:** longest arithmetic subsequence (needs DP, the set trick stops working)
