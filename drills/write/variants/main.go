// VARIANT DRILL — same pattern, different ask
//
// GOAL: Recognize when a familiar template applies to a variant problem.
// RUN: go run -C drills/write/variants .
//
// These are the stretch problems: each one reuses a move from the weekday
// rotation but hides it behind a different statement. Kadane, the variable
// window, and product-except-self graduated into drills 01 and 03 — what is
// left here is the next rung up.
package main

import (
	"fmt"
	"reflect"
	"sort"
)

// TODO: REFLEX — two sum on SORTED array (return indices, one solution guaranteed)
func twoSumSorted(nums []int, target int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — merge overlapping intervals; input is unsorted
func mergeIntervals(intervals [][]int) [][]int {
	panic("Implement from memory")
}

// TODO: REFLEX — matrix values in spiral order, clockwise from the top-left
func spiralOrder(matrix [][]int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — longest run of consecutive integers (O(n), not by sorting)
func longestConsecutive(nums []int) int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	_ = sort.Ints // mergeIntervals wants a sort; keep the import honest

	assert("twoSumSorted", reflect.DeepEqual(twoSumSorted([]int{2, 7, 11, 15}, 9), []int{0, 1}))
	assert("twoSumSorted negatives", reflect.DeepEqual(twoSumSorted([]int{-1, 0}, -1), []int{0, 1}))
	assert("twoSumSorted far ends", reflect.DeepEqual(twoSumSorted([]int{1, 3, 4, 9}, 10), []int{0, 3}))

	assert("mergeIntervals basic", reflect.DeepEqual(
		mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
		[][]int{{1, 6}, {8, 10}, {15, 18}}))
	assert("mergeIntervals touching", reflect.DeepEqual(
		mergeIntervals([][]int{{1, 4}, {4, 5}}), [][]int{{1, 5}}))
	assert("mergeIntervals unsorted", reflect.DeepEqual(
		mergeIntervals([][]int{{5, 6}, {1, 3}}), [][]int{{1, 3}, {5, 6}}))
	assert("mergeIntervals nested", reflect.DeepEqual(
		mergeIntervals([][]int{{1, 10}, {2, 3}}), [][]int{{1, 10}}))
	assert("mergeIntervals single", reflect.DeepEqual(
		mergeIntervals([][]int{{1, 2}}), [][]int{{1, 2}}))

	assert("spiralOrder square", reflect.DeepEqual(
		spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}))
	assert("spiralOrder wide", reflect.DeepEqual(
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}),
		[]int{1, 2, 3, 4, 8, 7, 6, 5}))
	assert("spiralOrder single row", reflect.DeepEqual(
		spiralOrder([][]int{{1, 2, 3}}), []int{1, 2, 3}))
	assert("spiralOrder single column", reflect.DeepEqual(
		spiralOrder([][]int{{1}, {2}, {3}}), []int{1, 2, 3}))

	assert("longestConsecutive basic", longestConsecutive([]int{100, 4, 200, 1, 3, 2}) == 4)
	assert("longestConsecutive empty", longestConsecutive([]int{}) == 0)
	assert("longestConsecutive duplicates", longestConsecutive([]int{1, 2, 0, 1}) == 3)
	assert("longestConsecutive single", longestConsecutive([]int{7}) == 1)

	fmt.Println("\nAll variant drills passed.")
}
