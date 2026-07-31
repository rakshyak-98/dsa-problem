//go:build ignore

// REFLEX DRILL 08 — Heaps (bonus)
//
// RUN: go run ./drills/08_heap_reflex
//
// AFTER PASSING: heaps/medium/kth_largest_element_in_an_array.js
package main

import (
	"container/heap"
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

	assert("lastStoneWeight", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
	assert("lastStoneWeight single", lastStoneWeight([]int{5}) == 5)

	got := mergeKSorted([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	assert("mergeKSorted", len(got) == 8 && got[0] == 1 && got[len(got)-1] == 6)

	_ = heap.IntHeap{} // ensure container/heap import is valid for solutions
	fmt.Println("\nAll heap reflex drills passed.")
	fmt.Println("Primary: heaps/medium/kth_largest_element_in_an_array.js")
}
