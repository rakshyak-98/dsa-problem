//go:build ignore

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
	dup := []int{1, 1, 2, 2, 3}
	assert("removeDuplicates len", removeDuplicates(dup) == 3)
	assert("removeDuplicates vals", reflect.DeepEqual(dup[:3], []int{1, 2, 3}))
	assert("removeDuplicates empty", removeDuplicates([]int{}) == 0)
	singleDup := []int{7}
	assert("removeDuplicates single", removeDuplicates(singleDup) == 1 && singleDup[0] == 7)
	allSame := []int{2, 2, 2}
	assert("removeDuplicates all same", removeDuplicates(allSame) == 1 && allSame[0] == 2)

	zeros := []int{0, 1, 0, 3, 12}
	moveZeroes(zeros)
	assert("moveZeroes", reflect.DeepEqual(zeros, []int{1, 3, 12, 0, 0}))
	moveZeroesSingle := []int{0}
	moveZeroes(moveZeroesSingle)
	assert("moveZeroes single", reflect.DeepEqual(moveZeroesSingle, []int{0}))
	allZeros := []int{0, 0, 0}
	moveZeroes(allZeros)
	assert("moveZeroes all zeros", reflect.DeepEqual(allZeros, []int{0, 0, 0}))
	noZeros := []int{1, 2, 3}
	moveZeroes(noZeros)
	assert("moveZeroes no zeros", reflect.DeepEqual(noZeros, []int{1, 2, 3}))

	assert("maxArea", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	assert("maxArea two", maxArea([]int{1, 1}) == 1)
	assert("maxArea wide shallow", maxArea([]int{1, 2, 1}) == 2)

	assert("isPalindrome", isPalindrome("A man, a plan, a canal: Panama") == true)
	assert("isPalindrome false", isPalindrome("race a car") == false)
	assert("isPalindrome empty", isPalindrome("") == true)
	assert("isPalindrome single", isPalindrome("a") == true)
	assert("isPalindrome digits", isPalindrome("121") == true)

	assert("maxSumSubarrayK", maxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3) == 9)
	assert("maxSumSubarrayK k=1", maxSumSubarrayK([]int{4, 2, 9}, 1) == 9)
	assert("maxSumSubarrayK k=len", maxSumSubarrayK([]int{1, 2, 3}, 3) == 6)
	assert("maxSumSubarrayK negatives", maxSumSubarrayK([]int{-1, -2, -3}, 2) == -3)

	fmt.Println("\nAll two-pointer reflex drills passed.")
	fmt.Println("Primary: two_pointers/easy/move_zeroes.js")
}
