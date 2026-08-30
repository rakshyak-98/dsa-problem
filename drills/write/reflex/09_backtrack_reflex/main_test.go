package main

import (
	"fmt"
	"reflect"
	"sort"
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

func sortedSlices(in [][]int) [][]int {
	out := append([][]int(nil), in...)
	for _, s := range out {
		sort.Ints(s)
	}
	sort.Slice(out, func(i, j int) bool {
		if len(out[i]) != len(out[j]) {
			return len(out[i]) < len(out[j])
		}
		for k := range out[i] {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

func TestSubsets(t *testing.T) {
	assert(t, "subsets n=3 count", len(subsets([]int{1, 2, 3})) == 8)
	assert(t, "subsets empty input", reflect.DeepEqual(subsets([]int{}), [][]int{{}}))
	assert(t, "subsets single", reflect.DeepEqual(subsets([]int{1}), [][]int{{}, {1}}))
	assert(t, "subsets n=2 contents", reflect.DeepEqual(
		sortedSlices(subsets([]int{1, 2})),
		sortedSlices([][]int{{}, {1}, {2}, {1, 2}}),
	))
}

func TestPermute(t *testing.T) {
	assert(t, "permute n=3 count", len(permute([]int{1, 2, 3})) == 6)
	assert(t, "permute single", reflect.DeepEqual(permute([]int{7}), [][]int{{7}}))
	assert(t, "permute n=2 contents", reflect.DeepEqual(
		sortedSlices(permute([]int{1, 2})),
		sortedSlices([][]int{{1, 2}, {2, 1}}),
	))
}

func TestCombine(t *testing.T) {
	combs := combine(4, 2)
	assert(t, "combine n=4 k=2 count", len(combs) == 6)
	assert(t, "combine n=4 k=2 sample", reflect.DeepEqual(sortedSlices(combs)[0], []int{1, 2}))
	assert(t, "combine k=1", len(combine(3, 1)) == 3)
	assert(t, "combine k=n", len(combine(3, 3)) == 1)
	assert(t, "combine n=1", reflect.DeepEqual(combine(1, 1), [][]int{{1}}))
	assert(t, "combine k=0", reflect.DeepEqual(combine(4, 0), [][]int{{}}))
	assert(t, "combine k=2 n=2", reflect.DeepEqual(sortedSlices(combine(2, 2)), sortedSlices([][]int{{1, 2}})))
}

func TestAll(t *testing.T) {
	t.Run("subsets", TestSubsets)
	t.Run("permute", TestPermute)
	t.Run("combine", TestCombine)
	fmt.Println("\nAll backtracking reflex drills passed.")
}
