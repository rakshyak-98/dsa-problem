// REFLEX DRILL 04 — Binary Search
//
// RUN: go run ./drills/04_binary_search_reflex
package main

import "fmt"

// TODO: REFLEX — return index of target or -1
func binarySearch(nums []int, target int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — insert position (first index where nums[i] >= target)
func searchInsert(nums []int, target int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — find minimum in rotated sorted array (no duplicates)
func findMinRotated(nums []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — is target in nums (classic template with lo <= hi)
func isTargetPresent(nums []int, target int) bool {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("binarySearch found", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert("binarySearch missing", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert("searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert("searchInsert new", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert("findMinRotated", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert("isTargetPresent", isTargetPresent([]int{1, 2, 3, 4, 5}, 3) == true)

	fmt.Println("\nAll binary search reflex drills passed.")
}
