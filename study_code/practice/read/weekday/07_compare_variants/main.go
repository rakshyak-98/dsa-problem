// READ DRILL 07 — Compare variants
//
// GOAL: Same ask, two shapes. Name the shared ask and the tradeoff.
// RUN:  go run ./drills/07_compare_variants
package main

import (
	"fmt"
	"strings"
)

// Pair 1 — two sum style
func v1a(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func v1b(nums []int, target int) []int {
	seen := map[int]int{}
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

// TODO: READ — shared ask must include "sum" and "target"
var pair1Ask = ""

// TODO: READ — v1a time "O(n^2)"; v1b time "O(n)" — fill "n^2 vs n" exactly that phrase
var pair1Trade = ""

// Pair 2 — climb stairs / fib ways
func v2a(n int) int {
	if n <= 2 {
		return n
	}
	return v2a(n-1) + v2a(n-2)
}

func v2b(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// TODO: READ — shared ask must include "ways" or "climb" or "stair"
var pair2Ask = ""

// TODO: READ — which is better for large n? "v2a" or "v2b"
var pair2Pick = ""

// Pair 3 — reverse string/array
func v3a(nums []int) []int {
	out := make([]int, len(nums))
	for i := range nums {
		out[len(nums)-1-i] = nums[i]
	}
	return out
}

func v3b(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// TODO: READ — does v3b mutate input? "yes" or "no"
var pair3Mutates = ""

// TODO: READ — extra array space: which uses O(n) extra? "v3a" or "v3b"
var pair3Space = ""

func has(s string, parts ...string) bool {
	s = strings.ToLower(s)
	for _, p := range parts {
		if !strings.Contains(s, strings.ToLower(p)) {
			return false
		}
	}
	return true
}

func anyHas(s string, parts ...string) bool {
	s = strings.ToLower(s)
	for _, p := range parts {
		if strings.Contains(s, strings.ToLower(p)) {
			return true
		}
	}
	return false
}

func norm(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("pair1Ask", has(pair1Ask, "sum") && has(pair1Ask, "target"),
		`indices of two numbers that sum to target`)
	assert("pair1Trade", strings.Contains(norm(pair1Trade), "n^2") && strings.Contains(norm(pair1Trade), "n"),
		`use phrase like "n^2 vs n"`)

	assert("pair2Ask", anyHas(pair2Ask, "ways", "climb", "stair"),
		`number of ways to climb n stairs`)
	assert("pair2Pick", norm(pair2Pick) == "v2b",
		`iterative DP beats naive recursion`)

	assert("pair3Mutates", norm(pair3Mutates) == "yes", `v3b swaps in place`)
	assert("pair3Space", norm(pair3Space) == "v3a", `v3a allocates out`)

	_ = v1a
	_ = v1b
	_ = v2a
	_ = v2b
	_ = v3a
	_ = v3b
	fmt.Println("\nCompare-variants drill passed.")
}
