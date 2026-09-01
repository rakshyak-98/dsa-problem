package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func TestBinarySearch(t *testing.T) {
	assert(t, "binarySearch found mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert(t, "binarySearch missing mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert(t, "binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert(t, "binarySearch single hit", binarySearch([]int{5}, 5) == 0)
	assert(t, "binarySearch single miss", binarySearch([]int{5}, 2) == -1)
	assert(t, "binarySearch first", binarySearch([]int{1, 2, 3, 4, 5}, 1) == 0)
	assert(t, "binarySearch last", binarySearch([]int{1, 2, 3, 4, 5}, 5) == 4)
	assert(t, "binarySearch two hit", binarySearch([]int{1, 3}, 3) == 1)
	assert(t, "binarySearch two miss", binarySearch([]int{1, 3}, 2) == -1)
}

func TestSearchInsert(t *testing.T) {
	assert(t, "searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert(t, "searchInsert new mid", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert(t, "searchInsert end", searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	assert(t, "searchInsert empty", searchInsert([]int{}, 5) == 0)
	assert(t, "searchInsert start", searchInsert([]int{2, 4, 6}, 1) == 0)
	assert(t, "searchInsert single", searchInsert([]int{5}, 5) == 0)
	assert(t, "searchInsert after single", searchInsert([]int{5}, 7) == 1)
}

func TestFindMinRotated(t *testing.T) {
	assert(t, "findMinRotated pivot", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert(t, "findMinRotated two", findMinRotated([]int{3, 1}) == 1)
	assert(t, "findMinRotated sorted", findMinRotated([]int{1, 2, 3, 4}) == 1)
	assert(t, "findMinRotated single", findMinRotated([]int{2}) == 2)
	assert(t, "findMinRotated pivot end", findMinRotated([]int{2, 3, 4, 5, 1}) == 1)
}

func TestMinEatingSpeed(t *testing.T) {
	assert(t, "minEatingSpeed basic", minEatingSpeed([]int{3, 6, 7, 11}, 8) == 4)
	assert(t, "minEatingSpeed tight", minEatingSpeed([]int{30, 11, 23, 4, 20}, 5) == 30)
	assert(t, "minEatingSpeed loose", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6) == 23)
	assert(t, "minEatingSpeed single pile", minEatingSpeed([]int{12}, 3) == 4)
	assert(t, "minEatingSpeed huge h", minEatingSpeed([]int{1, 1, 1}, 100) == 1)
	assert(t, "minEatingSpeed h equals piles", minEatingSpeed([]int{5, 5, 5}, 3) == 5)
}

func TestAll(t *testing.T) {
	t.Run("binarySearch", TestBinarySearch)
	t.Run("searchInsert", TestSearchInsert)
	t.Run("findMinRotated", TestFindMinRotated)
	t.Run("minEatingSpeed", TestMinEatingSpeed)
	fmt.Println("\nAll binary search reflex drills passed.")
}
