// SOLUTION — Reflex 02 Hashing (peek after honest attempt)
package main

import (
	"slices"
	"fmt"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = struct{}{}
	}
	return false
}

func frequencyMap(arr []string) map[string]int {
	m := make(map[string]int)
	for _, x := range arr {
		m[x]++
	}
	return m
}

func firstUniqueChar(s string) string {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	for _, ch := range s {
		if freq[ch] == 1 {
			return string(ch)
		}
	}
	return ""
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, w := range strs {
		b := []byte(w)
		slices.Sort(b)
		key := string(b)
		m[key] = append(m[key], w)
	}
	out := make([][]string, 0, len(m))
	for _, g := range m {
		out = append(out, g)
	}
	return out
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func reflectDeepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// twoSum — basic, negatives, duplicates, zero, both negative, distant indices
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert("twoSum basic", len(got) == 2 && got[0] == 0 && got[1] == 1)
	assert("twoSum negatives", reflectDeepEqual(twoSum([]int{-1, -2, -3, -4, -5}, -8), []int{2, 4}))
	assert("twoSum duplicates", reflectDeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}))
	assert("twoSum zero", reflectDeepEqual(twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3}))
	assert("twoSum both negative", reflectDeepEqual(twoSum([]int{-3, 4, 3, 90}, 0), []int{0, 2}))
	assert("twoSum distant", reflectDeepEqual(twoSum([]int{1, 2, 3, 4, 5}, 9), []int{3, 4}))

	// containsDuplicate — found, not found, single, empty, pair, triple run
	assert("containsDuplicate true", containsDuplicate([]int{1, 2, 3, 1}))
	assert("containsDuplicate false", !containsDuplicate([]int{1, 2, 3, 4}))
	assert("containsDuplicate single", !containsDuplicate([]int{1}))
	assert("containsDuplicate empty", !containsDuplicate([]int{}))
	assert("containsDuplicate pair", containsDuplicate([]int{2, 2}))
	assert("containsDuplicate triple", containsDuplicate([]int{1, 1, 1}))

	// frequencyMap — multi, empty, single, repeated, distinct
	freq := frequencyMap([]string{"a", "b", "a", "c"})
	assert("frequencyMap basic", freq["a"] == 2 && freq["b"] == 1 && freq["c"] == 1)
	assert("frequencyMap empty", len(frequencyMap([]string{})) == 0)
	assert("frequencyMap single", frequencyMap([]string{"x"})["x"] == 1)
	assert("frequencyMap all same", frequencyMap([]string{"z", "z", "z"})["z"] == 3)
	assert("frequencyMap distinct", len(frequencyMap([]string{"a", "b", "c"})) == 3)

	// firstUniqueChar — middle, none, single, empty, last unique
	assert("firstUniqueChar basic", firstUniqueChar("leetcode") == "l")
	assert("firstUniqueChar none", firstUniqueChar("aabb") == "")
	assert("firstUniqueChar single", firstUniqueChar("z") == "z")
	assert("firstUniqueChar empty", firstUniqueChar("") == "")
	assert("firstUniqueChar last", firstUniqueChar("aabbcd") == "c")

	// groupAnagrams — multi-group, empty, single, no shared, all anagrams
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
	assert("groupAnagrams basic", ok)
	assert("groupAnagrams empty", len(groupAnagrams([]string{})) == 0)
	assert("groupAnagrams single", len(groupAnagrams([]string{"abc"})) == 1)
	assert("groupAnagrams no shared", len(groupAnagrams([]string{"ab", "cd", "ef"})) == 3)
	allA := groupAnagrams([]string{"abc", "bca", "cab"})
	assert("groupAnagrams all anagrams", len(allA) == 1 && len(allA[0]) == 3)

	fmt.Println("\nAll hashing reflex drills passed.")
	fmt.Println("Primary: hashing/easy/two_sum.js")
}
