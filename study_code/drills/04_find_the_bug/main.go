// READ DRILL 04 — Find the bug
//
// GOAL: Each function has exactly one logic bug. Name the bug kind + broken line idea.
// RUN:  go run ./drills/04_find_the_bug
package main

import (
	"fmt"
	"strings"
)

// BUG 1 — intended: exact binary search returning index or -1
func brokenSearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi { // BUG: should be lo <= hi
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// TODO: READ — bug kind: "off-by-one" | "wrong-move" | "missing-visit" | "bad-compare"
var bug1Kind = ""

// TODO: READ — fix description must contain "<="
var bug1Fix = ""

// BUG 2 — intended: container with most water (move the shorter line)
func brokenArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0
	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := (right - left) * h
		if area > best {
			best = area
		}
		if height[left] < height[right] {
			right-- // BUG: should advance left (the shorter side)
		} else {
			left++
		}
	}
	return best
}

// TODO: READ — bug kind
var bug2Kind = ""

// TODO: READ — which pointer should move when left is shorter? "left" or "right"
var bug2Move = ""

// BUG 3 — intended: flood fill / mark island; visit neighbors of land
func brokenDFS(grid [][]byte, r, c int) {
	rows, cols := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= rows || c >= cols {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	brokenDFS(grid, r+1, c)
	brokenDFS(grid, r-1, c)
	brokenDFS(grid, r, c+1)
	// BUG: missing left neighbor r, c-1
}

// TODO: READ — bug kind
var bug3Kind = ""

// TODO: READ — missing direction as "left" | "right" | "up" | "down"
var bug3Missing = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

func has(s string, part string) bool {
	return strings.Contains(norm(s), strings.ToLower(part))
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("bug1Kind", norm(bug1Kind) == "off-by-one",
		`loop uses lo < hi instead of lo <= hi`)
	assert("bug1Fix", has(bug1Fix, "<="),
		`describe fix using <=`)

	assert("bug2Kind", norm(bug2Kind) == "wrong-move",
		`moves the wrong pointer`)
	assert("bug2Move", norm(bug2Move) == "left",
		`when left is shorter, move left++`)

	assert("bug3Kind", norm(bug3Kind) == "missing-visit",
		`never recurses to c-1`)
	assert("bug3Missing", norm(bug3Missing) == "left",
		`missing left neighbor`)

	_ = brokenSearch
	_ = brokenArea
	_ = brokenDFS
	fmt.Println("\nFind-the-bug drill passed.")
}
