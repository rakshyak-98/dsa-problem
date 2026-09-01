// REFLEX DRILL 03 — Two Pointers & Sliding Window
//
// RUN: go run -C drills/write/reflex/03_two_pointers_reflex .
//
// AFTER PASSING: two_pointers/easy/move_zeroes.js
package main

import "unicode"

// TODO: REFLEX — remove duplicates from sorted array, return new length
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	write := 0
	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[write] {
			write++
			nums[write] = nums[read]
		}
	}
	return write + 1
}

// TODO: REFLEX — move all zeroes to end in-place
func moveZeroes(nums []int) {
	write := 0
	for read := range nums {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}
	for i := write; i < len(nums); i++ {
		nums[write] = 0
		write++
	}
}

// TODO: REFLEX — max area container (heights array)
func maxArea(heights []int) int {
	maxArea := 0
	left, right := 0, len(heights)-1
	for left < right {
		h := min(heights[left], heights[right])
		area := h * (right - left)
		if maxArea < area {
			maxArea = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// TODO: REFLEX — is palindrome (alphanumeric only, ignore case)
func isPalindrome(s string) bool {
	filter := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filter = append(filter, unicode.ToLower(ch))
		}
	}
	left, right := 0, len(filter)-1
	for left < right {
		if filter[left] != filter[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// TODO: REFLEX — max sum of subarray of size k (sliding window fixed)
func maxSumSubarrayK(nums []int, k int) int {
	sum := 0
	for i := range k {
		sum += nums[i]
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		if sum > best {
			best = sum
		}
	}
	return best
}

// TODO: REFLEX — length of the longest substring with no repeated character
func longestUniqueSubstring(s string) int {
	panic("Implement from memory")
}

func main() {}
