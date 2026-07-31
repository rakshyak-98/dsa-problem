// REFLEX DRILL 02 — Hashing
//
// GOAL: Hash map patterns = automatic.
// RUN: go run ./drills/02_hashing_reflex
package main

import (
	"fmt"
)

// TODO: REFLEX — classic two sum (return indices)
func twoSum(nums []int, target int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — return true if any value appears twice
func containsDuplicate(nums []int) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — return map of value -> frequency
func frequencyMap(arr []string) map[string]int {
	panic("Implement from memory")
}

// TODO: REFLEX — first character that appears only once; return "" if none
func firstUniqueChar(s string) string {
	panic("Implement from memory")
}

// TODO: REFLEX — group anagrams (return slice of groups)
func groupAnagrams(strs []string) [][]string {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert("twoSum", len(got) == 2 && got[0] == 0 && got[1] == 1)
	assert("containsDuplicate true", containsDuplicate([]int{1, 2, 3, 1}) == true)
	assert("containsDuplicate false", containsDuplicate([]int{1, 2, 3, 4}) == false)

	freq := frequencyMap([]string{"a", "b", "a", "c"})
	assert("frequencyMap", freq["a"] == 2 && freq["b"] == 1)

	assert("firstUniqueChar", firstUniqueChar("leetcode") == "l")
	assert("firstUniqueChar none", firstUniqueChar("aabb") == "")

	groups := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	ok := len(groups) == 3
	if ok {
		found := false
		for _, g := range groups {
			if len(g) == 3 {
				for _, w := range g {
					if w == "eat" {
						found = true
					}
				}
			}
		}
		ok = found
	}
	assert("groupAnagrams count", ok)

	fmt.Println("\nAll hashing reflex drills passed.")
}
