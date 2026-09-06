package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
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

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert(t, "twoSum basic", len(got) == 2 && got[0] == 0 && got[1] == 1)
	assert(t, "twoSum negatives", reflectDeepEqual(twoSum([]int{-1, -2, -3, -4, -5}, -8), []int{2, 4}))
	assert(t, "twoSum duplicates", reflectDeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}))
	assert(t, "twoSum zero", reflectDeepEqual(twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3}))
	assert(t, "twoSum both negative", reflectDeepEqual(twoSum([]int{-3, 4, 3, 90}, 0), []int{0, 2}))
	assert(t, "twoSum distant", reflectDeepEqual(twoSum([]int{1, 2, 3, 4, 5}, 9), []int{3, 4}))
}

func TestContainsDuplicate(t *testing.T) {
	assert(t, "containsDuplicate true", containsDuplicate([]int{1, 2, 3, 1}) == true)
	assert(t, "containsDuplicate false", containsDuplicate([]int{1, 2, 3, 4}) == false)
	assert(t, "containsDuplicate single", containsDuplicate([]int{1}) == false)
	assert(t, "containsDuplicate empty", containsDuplicate([]int{}) == false)
	assert(t, "containsDuplicate pair", containsDuplicate([]int{2, 2}) == true)
	assert(t, "containsDuplicate triple", containsDuplicate([]int{1, 1, 1}) == true)
}

func TestFrequencyMap(t *testing.T) {
	freq := frequencyMap([]string{"a", "b", "a", "c"})
	assert(t, "frequencyMap basic", freq["a"] == 2 && freq["b"] == 1 && freq["c"] == 1)
	assert(t, "frequencyMap empty", len(frequencyMap([]string{})) == 0)
	assert(t, "frequencyMap single", frequencyMap([]string{"x"})["x"] == 1)
	assert(t, "frequencyMap all same", frequencyMap([]string{"z", "z", "z"})["z"] == 3)
	assert(t, "frequencyMap distinct", len(frequencyMap([]string{"a", "b", "c"})) == 3)
}

func TestFirstUniqueChar(t *testing.T) {
	assert(t, "firstUniqueChar basic", firstUniqueChar("leetcode") == "l")
	assert(t, "firstUniqueChar none", firstUniqueChar("aabb") == "")
	assert(t, "firstUniqueChar single", firstUniqueChar("z") == "z")
	assert(t, "firstUniqueChar empty", firstUniqueChar("") == "")
	assert(t, "firstUniqueChar last", firstUniqueChar("aabbcd") == "c")
}

func TestGroupAnagrams(t *testing.T) {
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
	assert(t, "groupAnagrams basic", ok)
	assert(t, "groupAnagrams empty", len(groupAnagrams([]string{})) == 0)
	assert(t, "groupAnagrams single", len(groupAnagrams([]string{"abc"})) == 1)
	assert(t, "groupAnagrams no shared", len(groupAnagrams([]string{"ab", "cd", "ef"})) == 3)
	allA := groupAnagrams([]string{"abc", "bca", "cab"})
	assert(t, "groupAnagrams all anagrams", len(allA) == 1 && len(allA[0]) == 3)
}

func TestSingleNumber(t *testing.T) {
	assert(t, "singleNumber basic", singleNumber([]int{2, 2, 1}) == 1)
	assert(t, "singleNumber middle", singleNumber([]int{4, 1, 2, 1, 2}) == 4)
	assert(t, "singleNumber single", singleNumber([]int{7}) == 7)
	assert(t, "singleNumber negatives", singleNumber([]int{-1, -1, -3}) == -3)
	assert(t, "singleNumber zero", singleNumber([]int{0, 1, 1}) == 0)
	assert(t, "singleNumber last", singleNumber([]int{5, 3, 3}) == 5)
}

func TestAll(t *testing.T) {
	t.Run("twoSum", TestTwoSum)
	t.Run("containsDuplicate", TestContainsDuplicate)
	t.Run("frequencyMap", TestFrequencyMap)
	t.Run("firstUniqueChar", TestFirstUniqueChar)
	t.Run("groupAnagrams", TestGroupAnagrams)
	t.Run("singleNumber", TestSingleNumber)
	fmt.Println("\nAll hashing reflex drills passed.")
}
