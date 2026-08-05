//go:build ignore

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
	assert("reverseInPlace", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3}), []int{3, 2, 1}))
	assert("reverseInPlace empty", reflect.DeepEqual(reverseInPlace([]int{}), []int{}))
	assert("reverseInPlace single", reflect.DeepEqual(reverseInPlace([]int{7}), []int{7}))

	assert("indexOfMax", indexOfMax([]int{3, 1, 4, 4}) == 2)
	assert("indexOfMax single", indexOfMax([]int{5}) == 0)

	assert("arraySum", arraySum([]int{1, 2, 3, 4}) == 10)
	assert("arraySum empty", arraySum([]int{}) == 0)

	assert("rotateRight", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}))
	assert("rotateRight k>len", reflect.DeepEqual(rotateRight([]int{1, 2}, 5), []int{2, 1}))
	assert("rotateRight k=0", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 0), []int{1, 2, 3}))

	assert("runningSum", reflect.DeepEqual(runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}))
	assert("runningSum single", reflect.DeepEqual(runningSum([]int{5}), []int{5}))

	fmt.Println("\nAll array reflex drills passed.")
	fmt.Println("Primary: arrays/easy/plus_one.js")
}
