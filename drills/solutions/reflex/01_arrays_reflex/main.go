// SOLUTION — Reflex 01 Arrays (peek after honest attempt)
package main

import (
	"fmt"
	"reflect"
)

func reverseInPlace(arr []int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func indexOfMax(arr []int) int {
	best := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[best] {
			best = i
		}
	}
	return best
}

func arraySum(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func rotateRight(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	k %= n
	out := make([]int, n)
	copy(out, arr[n-k:])
	copy(out[k:], arr[:n-k])
	return out
}

func runningSum(arr []int) []int {
	out := make([]int, len(arr))
	sum := 0
	for i, x := range arr {
		sum += x
		out[i] = sum
	}
	return out
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
}
