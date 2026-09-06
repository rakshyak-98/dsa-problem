// SOLUTION — Reflex 01 Arrays (peek after honest attempt)
package main

import (
	"fmt"
	"reflect"
)

func reverseInPlace(arr []int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func rotateRight(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	k %= n
	out := make([]int, n)
	copy(out, arr[n-k:])
	copy(out[k:], arr[:n-k])
	return out
}

func runningSum(arr []int) []int {
	out := make([]int, len(arr))
	sum := 0
	for i, x := range arr {
		sum += x
		out[i] = sum
	}
	return out
}

func subarraySumK(nums []int, k int) int {
	// Prefix sums seen so far, counted. A subarray ending at i sums to k when
	// some earlier prefix equals running-k, so the map answers in O(1).
	seen := map[int]int{0: 1}
	running, count := 0, 0
	for _, x := range nums {
		running += x
		count += seen[running-k]
		seen[running]++
	}
	return count
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	prefix := 1
	for i := 0; i < n; i++ {
		out[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= suffix
		suffix *= nums[i]
	}
	return out
}

func maxSubarraySum(nums []int) int {
	// Kadane: at each step either extend the run or start fresh at nums[i].
	if len(nums) == 0 {
		return 0
	}
	cur, best := nums[0], nums[0]
	for _, x := range nums[1:] {
		if x > cur+x {
			cur = x
		} else {
			cur += x
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// mergeSort — stable divide and conquer: O(n log n) time, O(n) space.
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		out := make([]int, len(nums))
		copy(out, nums)
		return out
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] { // <= keeps equal keys stable
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}

// quickSort — Lomuto partition around the last element, recurse on each side.
func quickSort(nums []int) []int {
	out := make([]int, len(nums))
	copy(out, nums)
	qsort(out, 0, len(out)-1)
	return out
}

func qsort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	qsort(a, lo, i-1)
	qsort(a, i+1, hi)
}

// sieve — sieve of Eratosthenes: cross out multiples of each prime from p*p.
func sieve(n int) []int {
	primes := []int{}
	if n < 2 {
		return primes
	}
	composite := make([]bool, n+1)
	for p := 2; p*p <= n; p++ {
		if composite[p] {
			continue
		}
		for m := p * p; m <= n; m += p {
			composite[m] = true
		}
	}
	for v := 2; v <= n; v++ {
		if !composite[v] {
			primes = append(primes, v)
		}
	}
	return primes
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	// reverseInPlace — empty, single, two, odd/even, negatives, palindrome
	assert("reverseInPlace basic", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3}), []int{3, 2, 1}))
	assert("reverseInPlace empty", reflect.DeepEqual(reverseInPlace([]int{}), []int{}))
	assert("reverseInPlace single", reflect.DeepEqual(reverseInPlace([]int{7}), []int{7}))
	assert("reverseInPlace two", reflect.DeepEqual(reverseInPlace([]int{1, 2}), []int{2, 1}))
	assert("reverseInPlace even", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3, 4}), []int{4, 3, 2, 1}))
	assert("reverseInPlace negatives", reflect.DeepEqual(reverseInPlace([]int{-1, 0, 1}), []int{1, 0, -1}))
	assert("reverseInPlace palindrome", reflect.DeepEqual(reverseInPlace([]int{1, 2, 1}), []int{1, 2, 1}))

	// rotateRight — k=0/1/len-1/len/>len/double-cycle, empty, single
	assert("rotateRight k=1", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4}, 1), []int{4, 1, 2, 3}))
	assert("rotateRight k=2", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}))
	assert("rotateRight k=0", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 0), []int{1, 2, 3}))
	assert("rotateRight k=len-1", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 2), []int{2, 3, 1}))
	assert("rotateRight k=len", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 3), []int{1, 2, 3}))
	assert("rotateRight k>len", reflect.DeepEqual(rotateRight([]int{1, 2}, 5), []int{2, 1}))
	assert("rotateRight k double cycle", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 6), []int{1, 2, 3}))
	assert("rotateRight empty", reflect.DeepEqual(rotateRight([]int{}, 3), []int{}))
	assert("rotateRight single", reflect.DeepEqual(rotateRight([]int{7}, 5), []int{7}))

	// runningSum — empty, single, negatives, zeros, constant
	assert("runningSum basic", reflect.DeepEqual(runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}))
	assert("runningSum single", reflect.DeepEqual(runningSum([]int{5}), []int{5}))
	assert("runningSum empty", reflect.DeepEqual(runningSum([]int{}), []int{}))
	assert("runningSum negatives", reflect.DeepEqual(runningSum([]int{1, -1, 2}), []int{1, 0, 2}))
	assert("runningSum zeros", reflect.DeepEqual(runningSum([]int{0, 0, 0}), []int{0, 0, 0}))
	assert("runningSum constant", reflect.DeepEqual(runningSum([]int{2, 2, 2}), []int{2, 4, 6}))

	// subarraySumK — empty, no match, whole array, negatives, zeros, overlap
	assert("subarraySumK basic", subarraySumK([]int{1, 1, 1}, 2) == 2)
	assert("subarraySumK whole", subarraySumK([]int{1, 2, 3}, 6) == 1)
	assert("subarraySumK none", subarraySumK([]int{1, 2, 3}, 7) == 0)
	assert("subarraySumK empty", subarraySumK([]int{}, 0) == 0)
	assert("subarraySumK negatives", subarraySumK([]int{1, -1, 0}, 0) == 3)
	assert("subarraySumK zeros", subarraySumK([]int{0, 0, 0}, 0) == 6)
	assert("subarraySumK single hit", subarraySumK([]int{3}, 3) == 1)

	// productExceptSelf — basic, with zero, two zeros, negatives, two elements
	assert("productExceptSelf basic", reflect.DeepEqual(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}))
	assert("productExceptSelf one zero", reflect.DeepEqual(productExceptSelf([]int{1, 0, 3}), []int{0, 3, 0}))
	assert("productExceptSelf two zeros", reflect.DeepEqual(productExceptSelf([]int{0, 0, 3}), []int{0, 0, 0}))
	assert("productExceptSelf negatives", reflect.DeepEqual(productExceptSelf([]int{-1, 2, -3}), []int{-6, 3, -2}))
	assert("productExceptSelf two", reflect.DeepEqual(productExceptSelf([]int{2, 5}), []int{5, 2}))
	assert("productExceptSelf single", reflect.DeepEqual(productExceptSelf([]int{9}), []int{1}))

	// maxSubarraySum — mixed, all negative, all positive, single, empty
	assert("maxSubarraySum basic", maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	assert("maxSubarraySum all negative", maxSubarraySum([]int{-3, -1, -2}) == -1)
	assert("maxSubarraySum all positive", maxSubarraySum([]int{1, 2, 3}) == 6)
	assert("maxSubarraySum single", maxSubarraySum([]int{5}) == 5)
	assert("maxSubarraySum empty", maxSubarraySum([]int{}) == 0)
	assert("maxSubarraySum restart", maxSubarraySum([]int{-5, 8, -1, 3}) == 10)

	// mergeSort / quickSort — empty, single, duplicates, reverse, negatives
	assert("mergeSort basic", reflect.DeepEqual(mergeSort([]int{5, 2, 3, 1}), []int{1, 2, 3, 5}))
	assert("mergeSort duplicates", reflect.DeepEqual(mergeSort([]int{5, 1, 1, 2, 0, 0}), []int{0, 0, 1, 1, 2, 5}))
	assert("mergeSort negatives", reflect.DeepEqual(mergeSort([]int{-3, 4, -1, 0, -2}), []int{-3, -2, -1, 0, 4}))
	assert("quickSort basic", reflect.DeepEqual(quickSort([]int{5, 2, 3, 1}), []int{1, 2, 3, 5}))
	assert("quickSort reverse", reflect.DeepEqual(quickSort([]int{9, 7, 5, 3, 1}), []int{1, 3, 5, 7, 9}))

	// sieve — primes up to n, n below 2
	assert("sieve to 10", reflect.DeepEqual(sieve(10), []int{2, 3, 5, 7}))
	assert("sieve to 1", len(sieve(1)) == 0)
	assert("sieve to 20", reflect.DeepEqual(sieve(20), []int{2, 3, 5, 7, 11, 13, 17, 19}))

	fmt.Println("\nAll array reflex drills passed.")
	fmt.Println("Primary: https://leetcode.com/problems/subarray-sum-equals-k/")
}
