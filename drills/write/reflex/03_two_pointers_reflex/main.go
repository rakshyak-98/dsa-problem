// REFLEX DRILL 03 — Two Pointers & Sliding Window
//
// RUN: go run -C drills/write/reflex/03_two_pointers_reflex .
//
// AFTER PASSING: two_pointers/easy/move_zeroes.js
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — remove duplicates from sorted array, return new length
func removeDuplicates(nums []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — move all zeroes to end in-place
func moveZeroes(nums []int) {
	panic("Implement from memory")
}

// TODO: REFLEX — max area container (heights array)
func maxArea(heights []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — is palindrome (alphanumeric only, ignore case)
func isPalindrome(s string) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — max sum of subarray of size k (sliding window fixed)
func maxSumSubarrayK(nums []int, k int) int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	// removeDuplicates — standard, empty, single, all same, no dup, long run
	dup := []int{1, 1, 2, 2, 3}
	assert("removeDuplicates basic len", removeDuplicates(dup) == 3)
	assert("removeDuplicates basic vals", reflect.DeepEqual(dup[:3], []int{1, 2, 3}))
	assert("removeDuplicates empty", removeDuplicates([]int{}) == 0)
	singleDup := []int{7}
	assert("removeDuplicates single", removeDuplicates(singleDup) == 1 && singleDup[0] == 7)
	allSame := []int{2, 2, 2}
	assert("removeDuplicates all same", removeDuplicates(allSame) == 1 && allSame[0] == 2)
	noDup := []int{1, 2, 3, 4}
	assert("removeDuplicates no dup", removeDuplicates(noDup) == 4)
	longRun := []int{1, 1, 1, 2, 2, 3}
	assert("removeDuplicates long run", removeDuplicates(longRun) == 3 && reflect.DeepEqual(longRun[:3], []int{1, 2, 3}))

	// moveZeroes — mixed, single zero, all zeros, no zeros, zeros prefix/suffix
	zeros := []int{0, 1, 0, 3, 12}
	moveZeroes(zeros)
	assert("moveZeroes basic", reflect.DeepEqual(zeros, []int{1, 3, 12, 0, 0}))
	moveZeroesSingle := []int{0}
	moveZeroes(moveZeroesSingle)
	assert("moveZeroes single", reflect.DeepEqual(moveZeroesSingle, []int{0}))
	allZeros := []int{0, 0, 0}
	moveZeroes(allZeros)
	assert("moveZeroes all zeros", reflect.DeepEqual(allZeros, []int{0, 0, 0}))
	noZeros := []int{1, 2, 3}
	moveZeroes(noZeros)
	assert("moveZeroes no zeros", reflect.DeepEqual(noZeros, []int{1, 2, 3}))
	prefix := []int{0, 0, 1, 2}
	moveZeroes(prefix)
	assert("moveZeroes prefix", reflect.DeepEqual(prefix, []int{1, 2, 0, 0}))
	suffix := []int{1, 2, 0, 0}
	moveZeroes(suffix)
	assert("moveZeroes suffix", reflect.DeepEqual(suffix, []int{1, 2, 0, 0}))

	// maxArea — classic, two bars, wide shallow, tall narrow, plateau
	assert("maxArea classic", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	assert("maxArea two", maxArea([]int{1, 1}) == 1)
	assert("maxArea wide shallow", maxArea([]int{1, 2, 1}) == 2)
	assert("maxArea tall narrow", maxArea([]int{4, 3, 2, 1, 4}) == 16)
	assert("maxArea plateau", maxArea([]int{5, 5, 5, 5}) == 15)

	// isPalindrome — classic, false, empty, single, digits, symbols, mixed false
	assert("isPalindrome classic", isPalindrome("A man, a plan, a canal: Panama") == true)
	assert("isPalindrome false", isPalindrome("race a car") == false)
	assert("isPalindrome empty", isPalindrome("") == true)
	assert("isPalindrome single", isPalindrome("a") == true)
	assert("isPalindrome digits", isPalindrome("121") == true)
	assert("isPalindrome symbols only", isPalindrome("., ") == true)
	assert("isPalindrome mixed false", isPalindrome("0P") == false)

	// maxSumSubarrayK — standard, k=1, k=len, negatives, all positive windows
	assert("maxSumSubarrayK basic", maxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3) == 9)
	assert("maxSumSubarrayK k=1", maxSumSubarrayK([]int{4, 2, 9}, 1) == 9)
	assert("maxSumSubarrayK k=len", maxSumSubarrayK([]int{1, 2, 3}, 3) == 6)
	assert("maxSumSubarrayK negatives", maxSumSubarrayK([]int{-1, -2, -3}, 2) == -3)
	assert("maxSumSubarrayK window slide", maxSumSubarrayK([]int{1, 4, 2, 10, 2, 1}, 2) == 14)

	fmt.Println("\nAll two-pointer reflex drills passed.")
	fmt.Println("Primary: two_pointers/easy/move_zeroes.js")
}
