// REFLEX DRILL 01 — Arrays
//
// GOAL: Write every function from memory in < 3 minutes each.
// RUN: go run -C drills/write/reflex/01_arrays_reflex .
//
// AFTER PASSING: arrays/easy/plus_one.js
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — reverse array in-place, return the same slice
// INVARIANT: L/R move inward; swap until they meet
func reverseInPlace(arr []int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — return index of max element (first max if ties)
func indexOfMax(arr []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — return sum of all elements
func arraySum(arr []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — rotate right by k steps (use modulo on k)
func rotateRight(arr []int, k int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — return new slice: running sum (prefix as output)
func runningSum(arr []int) []int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	// reverseInPlace — empty, single, two, odd/even, negatives, palindrome
	assert("reverseInPlace basic", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3}), []int{3, 2, 1}))
	assert("reverseInPlace empty", reflect.DeepEqual(reverseInPlace([]int{}), []int{}))
	assert("reverseInPlace single", reflect.DeepEqual(reverseInPlace([]int{7}), []int{7}))
	assert("reverseInPlace two", reflect.DeepEqual(reverseInPlace([]int{1, 2}), []int{2, 1}))
	assert("reverseInPlace even", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3, 4}), []int{4, 3, 2, 1}))
	assert("reverseInPlace negatives", reflect.DeepEqual(reverseInPlace([]int{-1, 0, 1}), []int{1, 0, -1}))
	assert("reverseInPlace palindrome", reflect.DeepEqual(reverseInPlace([]int{1, 2, 1}), []int{1, 2, 1}))

	// indexOfMax — empty, single, ties, start/end, negatives
	assert("indexOfMax basic", indexOfMax([]int{3, 1, 4, 4}) == 2)
	assert("indexOfMax single", indexOfMax([]int{5}) == 0)
	assert("indexOfMax empty", indexOfMax([]int{}) == 0)
	assert("indexOfMax ties", indexOfMax([]int{5, 5, 5}) == 0)
	assert("indexOfMax at start", indexOfMax([]int{9, 1, 2}) == 0)
	assert("indexOfMax at end", indexOfMax([]int{1, 2, 9}) == 2)
	assert("indexOfMax negatives", indexOfMax([]int{-10, -3, -7}) == 1)

	// arraySum — empty, single, mixed/all negative, zeros
	assert("arraySum basic", arraySum([]int{1, 2, 3, 4}) == 10)
	assert("arraySum empty", arraySum([]int{}) == 0)
	assert("arraySum single", arraySum([]int{5}) == 5)
	assert("arraySum negatives", arraySum([]int{-1, 2, -3}) == -2)
	assert("arraySum all negative", arraySum([]int{-2, -3}) == -5)
	assert("arraySum zeros", arraySum([]int{0, 0, 0}) == 0)

	// rotateRight — k=0/1/len-1/len/>len/double-cycle, empty, single
	assert("rotateRight k=1", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4}, 1), []int{4, 1, 2, 3}))
	assert("rotateRight k=2", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}))
	assert("rotateRight k=0", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 0), []int{1, 2, 3}))
	assert("rotateRight k=len-1", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 2), []int{2, 3, 1}))
	assert("rotateRight k=len", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 3), []int{1, 2, 3}))
	assert("rotateRight k>len", reflect.DeepEqual(rotateRight([]int{1, 2}, 5), []int{2, 1}))
	assert("rotateRight k double cycle", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 6), []int{1, 2, 3}))
	assert("rotateRight empty", reflect.DeepEqual(rotateRight([]int{}, 3), []int{}))
	assert("rotateRight single", reflect.DeepEqual(rotateRight([]int{7}, 5), []int{7}))

	// runningSum — empty, single, negatives, zeros, constant
	assert("runningSum basic", reflect.DeepEqual(runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}))
	assert("runningSum single", reflect.DeepEqual(runningSum([]int{5}), []int{5}))
	assert("runningSum empty", reflect.DeepEqual(runningSum([]int{}), []int{}))
	assert("runningSum negatives", reflect.DeepEqual(runningSum([]int{1, -1, 2}), []int{1, 0, 2}))
	assert("runningSum zeros", reflect.DeepEqual(runningSum([]int{0, 0, 0}), []int{0, 0, 0}))
	assert("runningSum constant", reflect.DeepEqual(runningSum([]int{2, 2, 2}), []int{2, 4, 6}))

	fmt.Println("\nAll array reflex drills passed.")
	fmt.Println("Primary: arrays/easy/plus_one.js")
}
