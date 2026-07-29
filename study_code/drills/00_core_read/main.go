// CORE READ 3 — every day
//
// GOAL: Practice the first five passes of READING_PATTERNS.md on short snippets.
// HOW:  Fill every TODO: READ. Do not peek at answers/.
// RUN:  go run ./drills/00_core_read
package main

import (
	"fmt"
	"strings"
)

// =============================================================================
// SNIPPET A — fill answers below after Pass 1–5
// =============================================================================

func snippetA(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > best {
			best = sum
		}
	}
	return best
}

// TODO: READ — pattern label (lowercase). Example: "fixed sliding window"
var aPattern = ""

// TODO: READ — result of snippetA([]int{2, 1, 5, 1, 3, 2}, 3)
var aTrace = 0

// TODO: READ — time complexity as "O(n)" / "O(n^2)" / "O(log n)" etc.
var aTime = ""

// =============================================================================
// SNIPPET B
// =============================================================================

func snippetB(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		need := target - n
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

// TODO: READ — pattern label
var bPattern = ""

// TODO: READ — snippetB([]int{2, 7, 11, 15}, 9) as "i,j" e.g. "0,1"
var bTrace = ""

// TODO: READ — does it mutate nums? "yes" or "no"
var bMutates = ""

// =============================================================================
// SNIPPET C
// =============================================================================

func snippetC(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// TODO: READ — pattern label
var cPattern = ""

// TODO: READ — snippetC("abba") → "true" or "false"
var cTrace = ""

// TODO: READ — one-sentence ask (must contain the word "palindrome")
var cAsk = ""

func norm(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func hasAll(s string, parts ...string) bool {
	s = norm(s)
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
	assert("A pattern", hasAll(aPattern, "window") || hasAll(aPattern, "sliding"),
		`set aPattern e.g. "fixed sliding window"`)
	assert("A trace", aTrace == 9, "snippetA([2,1,5,1,3,2], 3) max window sum is 9")
	assert("A time", norm(aTime) == "o(n)", `set aTime to "O(n)"`)

	assert("B pattern", hasAll(bPattern, "map") || hasAll(bPattern, "hash") || hasAll(bPattern, "complement"),
		`set bPattern e.g. "hash complement" or "map two sum"`)
	assert("B trace", norm(bTrace) == "0,1" || norm(bTrace) == "0 1",
		`set bTrace to "0,1"`)
	assert("B mutates", norm(bMutates) == "no", `set bMutates to "no"`)

	assert("C pattern", hasAll(cPattern, "two") && hasAll(cPattern, "pointer") || hasAll(cPattern, "opposite"),
		`set cPattern e.g. "two pointers opposite ends"`)
	assert("C trace", norm(cTrace) == "true", `set cTrace to "true"`)
	assert("C ask", hasAll(cAsk, "palindrome"),
		`set cAsk to a sentence containing "palindrome"`)

	_ = snippetA
	_ = snippetB
	_ = snippetC
	fmt.Println("\nCore Read 3 passed.")
}
