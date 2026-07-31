# Core 5 — Solutions Reference

Peek only after an honest blind attempt.

## twoSum
- **Trigger:** pair sums to target
- **Invariant:** map holds value→index for seen elements
- **Bug:** insert before checking complement
- **Complexity:** O(n) time, O(n) space

```go
func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    for i, v := range nums {
        if j, ok := seen[target-v]; ok {
            return []int{j, i}
        }
        seen[v] = i
    }
    return nil
}
```

## binarySearch
- **Trigger:** sorted + exact find
- **Invariant:** answer in [lo, hi] if exists
- **Bug:** infinite loop from wrong mid update
- **Complexity:** O(log n)

## removeDuplicates
- **Trigger:** in-place filter on sorted array
- **Invariant:** prefix [0:write) is unique sorted
- **Bug:** compare to write-1 not read-1

## maxSumSubarrayK
- **Trigger:** fixed window size k
- **Invariant:** window sum always size k after init
- **Bug:** forget to subtract left element when sliding

## frequencyMap
- **Trigger:** count occurrences
- **Invariant:** m[x] is count of x seen so far
- **Bug:** initialize wrong type for empty input
