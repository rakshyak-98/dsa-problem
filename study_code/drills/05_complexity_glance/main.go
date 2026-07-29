// READ DRILL 05 — Complexity at a glance
//
// GOAL: Bound time/space from shape. Watch for amortized two-pointer.
// RUN:  go run ./drills/05_complexity_glance
package main

import (
	"fmt"
	"strings"
)

func c1(nums []int) int {
	left := 0
	best := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > 10 && left <= right {
			sum -= nums[left]
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// TODO: READ — time as "O(n)" or "O(n^2)"
var c1Time = ""

// TODO: READ — extra space as "O(1)" or "O(n)"
var c1Space = ""

func c2(nums []int) [][]int {
	n := len(nums)
	out := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					out = append(out, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return out
}

// TODO: READ — time
var c2Time = ""

func c3(nums []int, k int) bool {
	seen := map[int]bool{}
	for _, x := range nums {
		if seen[k-x] {
			return true
		}
		seen[x] = true
	}
	return false
}

// TODO: READ — time
var c3Time = ""

// TODO: READ — space
var c3Space = ""

func c4(n int) int {
	if n <= 1 {
		return n
	}
	return c4(n-1) + c4(n-2)
}

// TODO: READ — time (naive fib): "O(2^n)" or "O(n)" or "O(n^2)"
var c4Time = ""

func norm(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("c1Time", norm(c1Time) == "o(n)",
		`left only moves forward → amortized O(n), not O(n^2)`)
	assert("c1Space", norm(c1Space) == "o(1)", `a few ints`)

	assert("c2Time", norm(c2Time) == "o(n^3)" || norm(c2Time) == "o(n3)",
		`three nested loops`)

	assert("c3Time", norm(c3Time) == "o(n)", `one pass + map`)
	assert("c3Space", norm(c3Space) == "o(n)", `map can hold n keys`)

	assert("c4Time", norm(c4Time) == "o(2^n)",
		`naive double recursion → O(2^n)`)

	_ = c1
	_ = c2
	_ = c3
	_ = c4
	fmt.Println("\nComplexity drill passed.")
}
