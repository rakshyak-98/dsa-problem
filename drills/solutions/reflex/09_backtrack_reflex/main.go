// SOLUTION — Reflex 09 Backtracking (peek after honest attempt)
package main

import (
	"fmt"
	"reflect"
	"sort"
)

func subsets(nums []int) [][]int {
	var out [][]int
	var path []int
	var dfs func(int)
	dfs = func(start int) {
		cp := append([]int{}, path...)
		out = append(out, cp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return out
}

func permute(nums []int) [][]int {
	var out [][]int
	path := make([]int, len(nums))
	used := make([]bool, len(nums))
	var dfs func(int)
	dfs = func(depth int) {
		if depth == len(nums) {
			cp := append([]int{}, path...)
			out = append(out, cp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path[depth] = nums[i]
			dfs(depth + 1)
			used[i] = false
		}
	}
	dfs(0)
	return out
}

func combine(n, k int) [][]int {
	var out [][]int
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			cp := append([]int{}, path...)
			out = append(out, cp)
			return
		}
		need := k - len(path)
		for i := start; i <= n; i++ {
			if n-i+1 < need {
				break
			}
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return out
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
	// subsets — counts and explicit contents for n=0,1,2,3
	assert("subsets n=3 count", len(subsets([]int{1, 2, 3})) == 8)
	assert("subsets empty input", reflect.DeepEqual(subsets([]int{}), [][]int{{}}))
	assert("subsets single", reflect.DeepEqual(subsets([]int{1}), [][]int{{}, {1}}))
	assert("subsets n=2 contents", reflect.DeepEqual(
		sortedSlices(subsets([]int{1, 2})),
		sortedSlices([][]int{{}, {1}, {2}, {1, 2}}),
	))

	// permute — counts and explicit for n=1,2,3
	assert("permute n=3 count", len(permute([]int{1, 2, 3})) == 6)
	assert("permute single", reflect.DeepEqual(permute([]int{7}), [][]int{{7}}))
	assert("permute n=2 contents", reflect.DeepEqual(
		sortedSlices(permute([]int{1, 2})),
		sortedSlices([][]int{{1, 2}, {2, 1}}),
	))

	// combine — varied n,k including k=0, k=1, k=n, k>n invalid
	combs := combine(4, 2)
	assert("combine n=4 k=2 count", len(combs) == 6)
	assert("combine n=4 k=2 sample", reflect.DeepEqual(sortedSlices(combs)[0], []int{1, 2}))
	assert("combine k=1", len(combine(3, 1)) == 3)
	assert("combine k=n", len(combine(3, 3)) == 1)
	assert("combine n=1", reflect.DeepEqual(combine(1, 1), [][]int{{1}}))
	assert("combine k=0", reflect.DeepEqual(combine(4, 0), [][]int{{}}))
	assert("combine k=2 n=2", reflect.DeepEqual(sortedSlices(combine(2, 2)), sortedSlices([][]int{{1, 2}})))

	fmt.Println("\nAll backtracking reflex drills passed.")
	fmt.Println("Primary: backtracking/medium/subsets.js")
}
