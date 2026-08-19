// REFLEX DRILL 03 — Two Pointers & Sliding Window
//
// RUN: go run -C drills/write/reflex/03_two_pointers_reflex .
//
// AFTER PASSING: two_pointers/easy/move_zeroes.js
package main

// TODO: REFLEX — remove duplicates from sorted array, return new length
func removeDuplicates(nums []int) int {
	write, read := 0, 0
	for read < len(nums) {
		if nums[read] != nums[write] {
			nums[write] = nums[read]
			write++
		}
		read++
	}
	return write
}

// TODO: REFLEX — move all zeroes to end in-place
func moveZeroes(nums []int) {
	panic("Implement from memory")
}

// TODO: REFLEX — max area container (heights array)
func maxArea(heights []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — is palindrome (alphanumeric only, ignore case)
func isPalindrome(s string) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — max sum of subarray of size k (sliding window fixed)
func maxSumSubarrayK(nums []int, k int) int {
	panic("Implement from memory")
}

func main() {}
