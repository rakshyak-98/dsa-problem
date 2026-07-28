// REFLEX DRILL 01 — Arrays
//
// GOAL: Write every function from memory in < 3 minutes each.
// RUN: go run ./drills/01_arrays_reflex
//
// After filling TODOs, all tests should pass.
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — reverse array in-place, return the same slice
func reverseInPlace(arr []int) []int {
	j := len(arr) - 1
	i := 0
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		j--
		i++
	}
	return arr
}

// TODO: REFLEX — return index of max element (first max if ties)
func indexOfMax(arr []int) int {
	maxIndex := 0
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	return maxIndex
}

// TODO: REFLEX — return sum of all elements
func arraySum(arr []int) int {
	sum := 0
	i := 0
	for i < len(arr) {
		sum += arr[i]
		i++
	}
	return sum
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
	assert("indexOfMax", indexOfMax([]int{3, 1, 4, 4}) == 2)
	assert("arraySum", arraySum([]int{1, 2, 3, 4}) == 10)
	assert("rotateRight", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}))
	assert("rotateRight k>len", reflect.DeepEqual(rotateRight([]int{1, 2}, 5), []int{2, 1}))
	assert("runningSum", reflect.DeepEqual(runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}))
	fmt.Println("\nAll array reflex drills passed.")
}
