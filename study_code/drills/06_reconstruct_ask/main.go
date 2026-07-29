// READ DRILL 06 — Reconstruct the ask
//
// GOAL: Write the problem in one plain sentence from the code alone.
// RUN:  go run ./drills/06_reconstruct_ask
package main

import (
	"fmt"
	"strings"
)

func snipA(nums []int, k int) bool {
	seen := map[int]int{}
	for i, x := range nums {
		if j, ok := seen[x]; ok && i-j <= k {
			return true
		}
		seen[x] = i
	}
	return false
}

// TODO: READ — one sentence; must include "duplicate" and "k"
var askA = ""

func snipB(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if ch == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, ch)
	}
	return string(stack)
}

// TODO: READ — must include "star" or "*" or "delete" or "remove"
var askB = ""

func snipC(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	take, skip := nums[0], 0
	for i := 1; i < len(nums); i++ {
		take, skip = skip+nums[i], max(take, skip)
	}
	return max(take, skip)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO: READ — must include "adjacent" or "rob" or "house"
var askC = ""

func snipD(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var ca, cb [26]int
	for i := 0; i < len(a); i++ {
		ca[a[i]-'a']++
		cb[b[i]-'a']++
	}
	return ca == cb
}

// TODO: READ — must include "anagram"
var askD = ""

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

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("askA", has(askA, "duplicate") && has(askA, "k"),
		`contains nearby duplicate within distance k`)
	assert("askB", anyHas(askB, "star", "*", "delete", "remove", "backspace"),
		`remove previous char when seeing *`)
	assert("askC", anyHas(askC, "adjacent", "rob", "house"),
		`max sum with no two adjacent`)
	assert("askD", has(askD, "anagram"),
		`check if two strings are anagrams`)

	_ = snipA
	_ = snipB
	_ = snipC
	_ = snipD
	fmt.Println("\nReconstruct-ask drill passed.")
}
