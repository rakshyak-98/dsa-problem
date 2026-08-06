// SOLUTION — Reflex 04 Binary Search (peek after honest attempt)
package main

import "fmt"

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func findMinRotated(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

func isTargetPresent(nums []int, target int) bool {
	return binarySearch(nums, target) != -1
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
	assert("binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert("binarySearch single hit", binarySearch([]int{5}, 5) == 0)
	assert("binarySearch single miss", binarySearch([]int{5}, 2) == -1)

	assert("searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert("searchInsert new", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert("searchInsert end", searchInsert([]int{1, 3, 5, 6}, 7) == 4)

	assert("findMinRotated", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert("findMinRotated two", findMinRotated([]int{3, 1}) == 1)

	assert("isTargetPresent true", isTargetPresent([]int{1, 2, 3, 4, 5}, 3))
	assert("isTargetPresent false", !isTargetPresent([]int{1, 2, 3, 4, 5}, 6))

	fmt.Println("\nAll binary search reflex drills passed.")
}
