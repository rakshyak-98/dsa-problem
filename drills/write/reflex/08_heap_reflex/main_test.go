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

func TestKthLargest(t *testing.T) {
	assert(t, "kthLargest basic", kthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	assert(t, "kthLargest dup", kthLargest([]int{3, 3, 3, 3}, 2) == 3)
	assert(t, "kthLargest k=1", kthLargest([]int{1, 2, 3}, 1) == 3)
	assert(t, "kthLargest single", kthLargest([]int{42}, 1) == 42)
	assert(t, "kthLargest k=len", kthLargest([]int{4, 2, 9}, 3) == 2)
	assert(t, "kthLargest negatives", kthLargest([]int{-1, -2, -3}, 2) == -2)
}

func TestLastStoneWeight(t *testing.T) {
	assert(t, "lastStoneWeight basic", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
	assert(t, "lastStoneWeight single", lastStoneWeight([]int{5}) == 5)
	assert(t, "lastStoneWeight cancel", lastStoneWeight([]int{5, 5}) == 0)
	assert(t, "lastStoneWeight chain", lastStoneWeight([]int{6, 3, 3}) == 0)
	assert(t, "lastStoneWeight three equal", lastStoneWeight([]int{4, 4, 4}) == 4)
}

func TestMergeKSorted(t *testing.T) {
	got := mergeKSorted([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	assert(t, "mergeKSorted basic", len(got) == 8 && got[0] == 1 && got[len(got)-1] == 6)
	assert(t, "mergeKSorted empty", len(mergeKSorted([][]int{})) == 0)
	assert(t, "mergeKSorted one list", reflectDeepEqual(mergeKSorted([][]int{{1, 2}}), []int{1, 2}))
	assert(t, "mergeKSorted with empty list", reflectDeepEqual(mergeKSorted([][]int{{}, {1, 3}, {2}}), []int{1, 2, 3}))
	assert(t, "mergeKSorted full", reflectDeepEqual(mergeKSorted([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4}))
	assert(t, "mergeKSorted duplicates", reflectDeepEqual(mergeKSorted([][]int{{1, 1}, {1}}), []int{1, 1, 1}))
}

func TestAll(t *testing.T) {
	t.Run("kthLargest", TestKthLargest)
	t.Run("lastStoneWeight", TestLastStoneWeight)
	t.Run("mergeKSorted", TestMergeKSorted)
	fmt.Println("\nAll heap reflex drills passed.")
}
