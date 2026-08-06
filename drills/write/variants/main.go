// VARIANT DRILL — same pattern, different ask
//
// GOAL: Recognize when a familiar template applies to a variant problem.
// RUN: go run ./variants
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — two sum on SORTED array (return indices, one solution guaranteed)
func twoSumSorted(nums []int, target int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — max sum of ANY contiguous subarray (variable window / Kadane)
func maxSubarraySum(nums []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	panic("Implement from memory")
}

// TODO: REFLEX — product of array except self (no division)
func productExceptSelf(nums []int) []int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("twoSumSorted", reflect.DeepEqual(twoSumSorted([]int{2, 7, 11, 15}, 9), []int{0, 1}))
	assert("twoSumSorted negatives", reflect.DeepEqual(twoSumSorted([]int{-1, 0}, -1), []int{0, 1}))

	assert("maxSubarraySum", maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	assert("maxSubarraySum all neg", maxSubarraySum([]int{-5, -2, -1}) == -1)
	assert("maxSubarraySum single", maxSubarraySum([]int{3}) == 3)

	assert("lengthOfLongestSubstring", lengthOfLongestSubstring("abcabcbb") == 3)
	assert("lengthOfLongestSubstring empty", lengthOfLongestSubstring("") == 0)

	assert("productExceptSelf", reflect.DeepEqual(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}))
	assert("productExceptSelf zeros", reflect.DeepEqual(productExceptSelf([]int{0, 1, 2}), []int{2, 0, 0}))

	fmt.Println("\nAll variant drills passed.")
	fmt.Println("Primary problems:")
	fmt.Println("  twoSumSorted → hashing/easy/two_sum.js (sorted variant)")
	fmt.Println("  maxSubarraySum → arrays/medium/max_product_subarray.js")
	fmt.Println("  lengthOfLongestSubstring → sliding_window/medium/longest_substring_without_repeating.js")
	fmt.Println("  productExceptSelf → arrays/medium/product_of_array_except_self.js")
}
