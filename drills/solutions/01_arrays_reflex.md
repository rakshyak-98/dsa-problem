# Drill 01 — Arrays

## reverseInPlace
- **Trigger:** in-place reverse
- **Pattern:** two pointers L/R swap
- **Bug:** forget to return same slice

## rotateRight
- **Trigger:** rotate by k
- **Pattern:** k %= n; reverse whole, reverse [0:k), reverse [k:n)
- **Bug:** k not reduced modulo n

## runningSum
- **Trigger:** prefix output
- **Pattern:** out[i] = out[i-1] + arr[i]
- **Bug:** mutate input instead of new slice

## subarraySumK
- **Trigger:** "how many subarrays sum to k" — contiguous, not pairs
- **Pattern:** running prefix + map of prefix counts; add `seen[running-k]`
- **Seed `seen[0]=1`** or every subarray starting at index 0 is missed
- **Bug:** counting indices instead of counts — duplicates prefixes need `++`, not `=`

## productExceptSelf
- **Trigger:** every position needs both sides, division banned
- **Pattern:** prefix pass left→right into out, then a suffix multiplier right→left
- **Bug:** allocating a second array — the suffix can multiply in place

## maxSubarraySum
- **Trigger:** best contiguous run
- **Pattern:** Kadane — `cur = max(x, cur+x)`, `best = max(best, cur)`
- **Bug:** initialising `best = 0`, which returns 0 for an all-negative input

## mergeSort
- **Trigger:** "implement a sort", stability matters, or count while sorting
- **Pattern:** split in half, recurse both, merge with two pointers (`<=` = stable)
- **Bug:** returning the input slice on the base case aliases the caller's array — copy it

## quickSort
- **Trigger:** in-place O(n log n) sort, no second array
- **Pattern:** Lomuto — pivot last, `i` tracks the `< pivot` frontier, swap pivot home
- **Bug:** recursing on `lo..i` instead of `lo..i-1` never terminates

## sieve
- **Trigger:** "all primes up to n", many primality checks over a range
- **Pattern:** cross out multiples of each `p` starting at `p*p`; survivors are prime
- **Cost:** O(n log log n) time, O(n) space
