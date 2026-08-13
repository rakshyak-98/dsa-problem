// REFLEX DRILL 04 — Binary Search
//
// RUN: go run -C drills/write/reflex/04_binary_search_reflex .
//
// AFTER PASSING: binary_search/easy/search_insertion_position.js
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
	// binarySearch — hit/miss, empty, single, first/last/mid, two-element
	assert("binarySearch found mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert("binarySearch missing mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert("binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert("binarySearch single hit", binarySearch([]int{5}, 5) == 0)
	assert("binarySearch single miss", binarySearch([]int{5}, 2) == -1)
	assert("binarySearch first", binarySearch([]int{1, 2, 3, 4, 5}, 1) == 0)
	assert("binarySearch last", binarySearch([]int{1, 2, 3, 4, 5}, 5) == 4)
	assert("binarySearch two hit", binarySearch([]int{1, 3}, 3) == 1)
	assert("binarySearch two miss", binarySearch([]int{1, 3}, 2) == -1)

	// searchInsert — exist, middle, end, start, empty, single, equal front
	assert("searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert("searchInsert new mid", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert("searchInsert end", searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	assert("searchInsert empty", searchInsert([]int{}, 5) == 0)
	assert("searchInsert start", searchInsert([]int{2, 4, 6}, 1) == 0)
	assert("searchInsert single", searchInsert([]int{5}, 5) == 0)
	assert("searchInsert after single", searchInsert([]int{5}, 7) == 1)

	// findMinRotated — pivot, two elem, sorted, single, pivot at end
	assert("findMinRotated pivot", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert("findMinRotated two", findMinRotated([]int{3, 1}) == 1)
	assert("findMinRotated sorted", findMinRotated([]int{1, 2, 3, 4}) == 1)
	assert("findMinRotated single", findMinRotated([]int{2}) == 2)
	assert("findMinRotated pivot end", findMinRotated([]int{2, 3, 4, 5, 1}) == 1)

	// isTargetPresent — mirrors search outcomes
	assert("isTargetPresent true mid", isTargetPresent([]int{1, 2, 3, 4, 5}, 3) == true)
	assert("isTargetPresent false", isTargetPresent([]int{1, 2, 3, 4, 5}, 6) == false)
	assert("isTargetPresent empty", isTargetPresent([]int{}, 1) == false)
	assert("isTargetPresent first", isTargetPresent([]int{1, 2, 3}, 1) == true)
	assert("isTargetPresent last", isTargetPresent([]int{1, 2, 3}, 3) == true)

	fmt.Println("\nAll binary search reflex drills passed.")
	fmt.Println("Primary: binary_search/easy/search_insertion_position.js")
}
