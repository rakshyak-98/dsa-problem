//go:build ignore

// REFLEX DRILL 08 — Heaps (bonus)
//
// RUN: go run -C drills/write/reflex/08_heap_reflex .
//
// AFTER PASSING: heaps/medium/kth_largest_element_in_an_array.js
package main

import (
	"fmt"
)

// TODO: REFLEX — kth largest element (0-indexed: 1st largest = max)
func kthLargest(nums []int, k int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — last stone weight after smashing two largest each round
func lastStoneWeight(stones []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — merge k sorted slices into one sorted slice
func mergeKSorted(lists [][]int) []int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("kthLargest", kthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	assert("kthLargest dup", kthLargest([]int{3, 3, 3, 3}, 2) == 3)
	assert("kthLargest k=1", kthLargest([]int{1, 2, 3}, 1) == 3)
	assert("kthLargest single", kthLargest([]int{42}, 1) == 42)

	assert("lastStoneWeight", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
	assert("lastStoneWeight single", lastStoneWeight([]int{5}) == 5)
	assert("lastStoneWeight cancel", lastStoneWeight([]int{5, 5}) == 0)

	got := mergeKSorted([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	assert("mergeKSorted", len(got) == 8 && got[0] == 1 && got[len(got)-1] == 6)
	assert("mergeKSorted empty", len(mergeKSorted([][]int{})) == 0)
	assert("mergeKSorted one list", reflectDeepEqual(mergeKSorted([][]int{{1, 2}}), []int{1, 2}))
	assert("mergeKSorted with empty list", reflectDeepEqual(mergeKSorted([][]int{{}, {1, 3}, {2}}), []int{1, 2, 3}))
	assert("mergeKSorted full", reflectDeepEqual(mergeKSorted([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4}))

	fmt.Println("\nAll heap reflex drills passed.")
	fmt.Println("Primary: heaps/medium/kth_largest_element_in_an_array.js")
}

func reflectDeepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
