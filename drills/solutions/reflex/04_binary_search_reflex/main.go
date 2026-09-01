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

func minEatingSpeed(piles []int, h int) int {
	// The search space is the answer itself (1..max pile), not an index. feasible
	// is monotonic: if speed s finishes in time, so does every speed above it.
	hours := func(speed int) int {
		total := 0
		for _, p := range piles {
			total += (p + speed - 1) / speed
		}
		return total
	}
	lo, hi := 1, 0
	for _, p := range piles {
		if p > hi {
			hi = p
		}
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hours(mid) <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
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

	// minEatingSpeed — exact fit, one pile, h == len(piles), huge h
	assert("minEatingSpeed basic", minEatingSpeed([]int{3, 6, 7, 11}, 8) == 4)
	assert("minEatingSpeed tight", minEatingSpeed([]int{30, 11, 23, 4, 20}, 5) == 30)
	assert("minEatingSpeed loose", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6) == 23)
	assert("minEatingSpeed single pile", minEatingSpeed([]int{12}, 3) == 4)
	assert("minEatingSpeed huge h", minEatingSpeed([]int{1, 1, 1}, 100) == 1)
	assert("minEatingSpeed h equals piles", minEatingSpeed([]int{5, 5, 5}, 3) == 5)

	fmt.Println("\nAll binary search reflex drills passed.")
	fmt.Println("Primary: binary_search/easy/search_insertion_position.js")
}
