//go:build ignore

// REFLEX DRILL 09 — Backtracking (bonus)
//
// RUN: go run -C drills/write/reflex/09_backtrack_reflex .
//
// AFTER PASSING: backtracking/medium/subsets.js
package main

import (
	"fmt"
	"reflect"
	"sort"
)

// TODO: REFLEX — return all subsets of nums (power set)
func subsets(nums []int) [][]int {
	panic("Implement from memory")
}

// TODO: REFLEX — return all permutations of nums
func permute(nums []int) [][]int {
	panic("Implement from memory")
}

// TODO: REFLEX — combinations of k numbers chosen from 1..n
func combine(n, k int) [][]int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
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

func main() {
	assert("subsets count", len(subsets([]int{1, 2, 3})) == 8)
	assert("subsets empty", reflect.DeepEqual(subsets([]int{}), [][]int{{}}))
	assert("subsets single", reflect.DeepEqual(subsets([]int{1}), [][]int{{}, {1}}))

	perms := permute([]int{1, 2, 3})
	assert("permute count", len(perms) == 6)
	assert("permute single", reflect.DeepEqual(permute([]int{7}), [][]int{{7}}))

	combs := combine(4, 2)
	assert("combine count", len(combs) == 6)
	assert("combine sample", reflect.DeepEqual(sortedSlices(combs)[0], []int{1, 2}))
	assert("combine k=1", len(combine(3, 1)) == 3)
	assert("combine k=n", len(combine(3, 3)) == 1)
	assert("combine n=1", reflect.DeepEqual(combine(1, 1), [][]int{{1}}))

	fmt.Println("\nAll backtracking reflex drills passed.")
	fmt.Println("Primary: backtracking/medium/subsets.js")
}
