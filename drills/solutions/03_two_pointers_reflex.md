# Drill 03 — Two Pointers & Window

## removeDuplicates
- **Invariant:** nums[0:write) unique sorted
- **Move write** when nums[read] != nums[write-1]

## moveZeroes
- **Pattern:** write pointer for non-zero; swap or fill zeros after

## maxArea
- **Pattern:** L/R; move shorter side inward
- **Bug:** compute area before moving wrong pointer

## maxSumSubarrayK
- **Pattern:** fixed window — build first k, slide with +/- ends

## isPalindrome
- **Pattern:** skip non-alphanumeric; compare lowercased
