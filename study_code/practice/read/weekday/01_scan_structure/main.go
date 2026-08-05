// READ DRILL 01 — Scan structure
//
// GOAL: Map signature, control flow, and state BEFORE tracing values.
// RUN: go run -C drills/read/weekday/01_scan_structure .
package main

import (
	"fmt"
	"strings"
)

// SNIPPET 1 — do not rename; answer questions below
func snip1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	best := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[best] {
			best = i
		}
	}
	return best
}

// TODO: READ — return type meaning: "index" or "value"
var s1Returns = ""

// TODO: READ — empty input returns? (as string, e.g. "-1")
var s1Empty = ""

// TODO: READ — number of loops (integer as string, e.g. "1")
var s1Loops = ""

// SNIPPET 2
func snip2(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	if n == 0 {
		return out
	}
	out[0] = nums[0]
	for i := 1; i < n; i++ {
		out[i] = out[i-1] + nums[i]
	}
	return out
}

// TODO: READ — mutates nums? "yes" or "no"
var s2Mutates = ""

// TODO: READ — builds which structure? "prefix" or "suffix" or "freq"
var s2Builds = ""

// TODO: READ — early exit / special case for? "empty" or "sorted" or "negative"
var s2Special = ""

// SNIPPET 3
func snip3(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

// TODO: READ — mutates grid? "yes" or "no"
var s3Mutates = ""

// TODO: READ — outer structure: "nested loops" or "single loop" or "queue"
var s3Outer = ""

// TODO: READ — helper style: "dfs" or "bfs" or "union"
var s3Helper = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("s1Returns", norm(s1Returns) == "index", `snip1 returns an index, not the max value`)
	assert("s1Empty", norm(s1Empty) == "-1", `empty → -1`)
	assert("s1Loops", norm(s1Loops) == "1", `one for-loop`)

	assert("s2Mutates", norm(s2Mutates) == "no", `writes to out, not nums`)
	assert("s2Builds", strings.Contains(norm(s2Builds), "prefix"), `running/prefix sums`)
	assert("s2Special", strings.Contains(norm(s2Special), "empty"), `n==0 early path`)

	assert("s3Mutates", norm(s3Mutates) == "yes", `marks visited by writing '0'`)
	assert("s3Outer", strings.Contains(norm(s3Outer), "nested"), `double for over cells`)
	assert("s3Helper", norm(s3Helper) == "dfs", `recursive 4-direction flood`)

	_ = snip1
	_ = snip2
	_ = snip3
	fmt.Println("\nScan structure drill passed.")
}
