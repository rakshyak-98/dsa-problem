// READ DRILL 03 — Name the pattern
//
// GOAL: Ignore names. Match skeleton → pattern label.
// RUN:  go run ./drills/03_name_the_pattern
package main

import (
	"fmt"
	"strings"
)

func alpha(s string) int {
	left := 0
	best := 0
	last := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := last[ch]; ok && idx >= left {
			left = idx + 1
		}
		last[ch] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// TODO: READ — pattern (must include "window" or "sliding")
var pAlpha = "sliding"

func beta(nums []int) []int {
	stack := []int{}
	out := make([]int, len(nums))
	for i := range out {
		out[i] = -1
	}
	for i, n := range nums {
		for len(stack) > 0 && n > nums[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[j] = n
		}
		stack = append(stack, i)
	}
	return out
}

// TODO: READ — pattern (must include "stack")
var pBeta = "stack"

func gamma(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

// TODO: READ — pattern (rotated / binary)
var pGamma = "binary"

func delta(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	a, b := cost[0], cost[1]
	for i := 2; i < n; i++ {
		a, b = b, cost[i]+min(a, b)
	}
	return min(a, b)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// TODO: READ — pattern (must include "dp" or "dynamic")
var pDelta = "dp"

func epsilon(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// TODO: READ — pattern
var pEpsilon = "dynamic"

func has(s string, parts ...string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	for _, p := range parts {
		if !strings.Contains(s, strings.ToLower(p)) {
			return false
		}
	}
	return true
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("alpha", has(pAlpha, "window") || has(pAlpha, "sliding"),
		`longest substring without repeat → variable sliding window`)
	assert("beta", has(pBeta, "stack"),
		`next greater → monotonic stack`)
	assert("gamma", has(pGamma, "binary") || has(pGamma, "rotated"),
		`min in rotated sorted → binary search on rotated array`)
	assert("delta", has(pDelta, "dp") || has(pDelta, "dynamic"),
		`min cost climb → 1D DP`)
	assert("epsilon", has(pEpsilon, "two") && has(pEpsilon, "pointer") || has(pEpsilon, "reverse"),
		`in-place reverse → two pointers`)

	_ = alpha
	_ = beta
	_ = gamma
	_ = delta
	_ = epsilon
	fmt.Println("\nName-the-pattern drill passed.")
}
