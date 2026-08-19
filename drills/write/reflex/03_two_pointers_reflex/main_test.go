package main

import (
	"fmt"
	"reflect"
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

func TestRemoveDuplicates(t *testing.T) {
	dup := []int{1, 1, 2, 2, 3}
	assert(t, "removeDuplicates basic len", removeDuplicates(dup) == 3)
	assert(t, "removeDuplicates basic vals", reflect.DeepEqual(dup[:3], []int{1, 2, 3}))
	assert(t, "removeDuplicates empty", removeDuplicates([]int{}) == 0)
	singleDup := []int{7}
	assert(t, "removeDuplicates single", removeDuplicates(singleDup) == 1 && singleDup[0] == 7)
	allSame := []int{2, 2, 2}
	assert(t, "removeDuplicates all same", removeDuplicates(allSame) == 1 && allSame[0] == 2)
	noDup := []int{1, 2, 3, 4}
	assert(t, "removeDuplicates no dup", removeDuplicates(noDup) == 4)
	longRun := []int{1, 1, 1, 2, 2, 3}
	assert(t, "removeDuplicates long run", removeDuplicates(longRun) == 3 && reflect.DeepEqual(longRun[:3], []int{1, 2, 3}))
}

func TestMoveZeroes(t *testing.T) {
	zeros := []int{0, 1, 0, 3, 12}
	moveZeroes(zeros)
	assert(t, "moveZeroes basic", reflect.DeepEqual(zeros, []int{1, 3, 12, 0, 0}))
	moveZeroesSingle := []int{0}
	moveZeroes(moveZeroesSingle)
	assert(t, "moveZeroes single", reflect.DeepEqual(moveZeroesSingle, []int{0}))
	allZeros := []int{0, 0, 0}
	moveZeroes(allZeros)
	assert(t, "moveZeroes all zeros", reflect.DeepEqual(allZeros, []int{0, 0, 0}))
	noZeros := []int{1, 2, 3}
	moveZeroes(noZeros)
	assert(t, "moveZeroes no zeros", reflect.DeepEqual(noZeros, []int{1, 2, 3}))
	prefix := []int{0, 0, 1, 2}
	moveZeroes(prefix)
	assert(t, "moveZeroes prefix", reflect.DeepEqual(prefix, []int{1, 2, 0, 0}))
	suffix := []int{1, 2, 0, 0}
	moveZeroes(suffix)
	assert(t, "moveZeroes suffix", reflect.DeepEqual(suffix, []int{1, 2, 0, 0}))
}

func TestMaxArea(t *testing.T) {
	assert(t, "maxArea classic", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	assert(t, "maxArea two", maxArea([]int{1, 1}) == 1)
	assert(t, "maxArea wide shallow", maxArea([]int{1, 2, 1}) == 2)
	assert(t, "maxArea tall narrow", maxArea([]int{4, 3, 2, 1, 4}) == 16)
	assert(t, "maxArea plateau", maxArea([]int{5, 5, 5, 5}) == 15)
}

func TestIsPalindrome(t *testing.T) {
	assert(t, "isPalindrome classic", isPalindrome("A man, a plan, a canal: Panama"))
	assert(t, "isPalindrome false", !isPalindrome("race a car"))
	assert(t, "isPalindrome empty", isPalindrome(""))
	assert(t, "isPalindrome single", isPalindrome("a"))
	assert(t, "isPalindrome digits", isPalindrome("121"))
	assert(t, "isPalindrome symbols only", isPalindrome("., "))
	assert(t, "isPalindrome mixed false", !isPalindrome("0P"))
}

func TestMaxSumSubarrayK(t *testing.T) {
	assert(t, "maxSumSubarrayK basic", maxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3) == 9)
	assert(t, "maxSumSubarrayK k=1", maxSumSubarrayK([]int{4, 2, 9}, 1) == 9)
	assert(t, "maxSumSubarrayK k=len", maxSumSubarrayK([]int{1, 2, 3}, 3) == 6)
	assert(t, "maxSumSubarrayK negatives", maxSumSubarrayK([]int{-1, -2, -3}, 2) == -3)
	assert(t, "maxSumSubarrayK window slide", maxSumSubarrayK([]int{1, 4, 2, 10, 2, 1}, 2) == 14)
}

func TestAll(t *testing.T) {
	t.Run("removeDuplicates", TestRemoveDuplicates)
	t.Run("moveZeroes", TestMoveZeroes)
	t.Run("maxArea", TestMaxArea)
	t.Run("isPalindrome", TestIsPalindrome)
	t.Run("maxSumSubarrayK", TestMaxSumSubarrayK)
	fmt.Println("\nAll two-pointer reflex drills passed.")
}
