// READ DRILL 02 — Trace execution
//
// GOAL: Hand-simulate. Predict returns from a step table, not vibes.
// RUN:  go run ./drills/02_trace_execution
package main

import (
	"fmt"
	"strings"
)

func snip1(nums []int) int {
	write := 0
	for read := 0; read < len(nums); read++ {
		if read == 0 || nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

// Trace snip1 on nums = [1, 1, 2, 2, 3] (sorted).
// TODO: READ — final write (new length)
var t1Len = 0

// TODO: READ — nums[:write] as "a,b,c" e.g. "1,2,3"
var t1Prefix = ""

func snip2(height []int) int {
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
			left++
		} else {
			right--
		}
	}
	return best
}

// Trace snip2 on height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
// TODO: READ — final best area
var t2Best = 0

func snip3(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
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

// Trace snip3 on nums = [1, 3, 5, 7, 9], target = 7
// Record mid values in order as "a,b,c" (the mid indices compared)
// TODO: READ
var t3Mids = ""

// TODO: READ — return value
var t3Ret = 0

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("t1Len", t1Len == 3, "unique prefix length is 3")
	assert("t1Prefix", norm(t1Prefix) == "1,2,3", `prefix becomes 1,2,3`)
	assert("t2Best", t2Best == 49, "classic container — best area 49")
	assert("t3Mids", norm(t3Mids) == "2,3", "mids: 2 (val 5), then 3 (val 7)")
	assert("t3Ret", t3Ret == 3, "7 is at index 3")

	// sanity that snippets match expected (reader may run mentally only)
	a := []int{1, 1, 2, 2, 3}
	assert("snip1 live", snip1(a) == 3, "internal")
	assert("snip2 live", snip2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49, "internal")
	assert("snip3 live", snip3([]int{1, 3, 5, 7, 9}, 7) == 3, "internal")

	fmt.Println("\nTrace execution drill passed.")
}
