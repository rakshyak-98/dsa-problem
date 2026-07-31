# Variants — Solutions Reference

## twoSumSorted
- **Pattern:** L/R pointers; if sum < target move left++, else right--
- **Different from hash** because array is sorted

## maxSubarraySum (Kadane)
- **cur = max(nums[i], cur+nums[i]); best = max(best, cur)**

## lengthOfLongestSubstring
- **Sliding window + last seen index map**
- **Shrink left** when duplicate inside window

## productExceptSelf
- **Prefix products left→right, suffix right→left**
- **Or:** output as prefix pass then multiply suffix in place
