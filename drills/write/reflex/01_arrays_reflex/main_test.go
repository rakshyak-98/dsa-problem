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

func TestReverseInPlace(t *testing.T) {
	assert(t, "reverseInPlace basic", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3}), []int{3, 2, 1}))
	assert(t, "reverseInPlace empty", reflect.DeepEqual(reverseInPlace([]int{}), []int{}))
	assert(t, "reverseInPlace single", reflect.DeepEqual(reverseInPlace([]int{7}), []int{7}))
	assert(t, "reverseInPlace two", reflect.DeepEqual(reverseInPlace([]int{1, 2}), []int{2, 1}))
	assert(t, "reverseInPlace even", reflect.DeepEqual(reverseInPlace([]int{1, 2, 3, 4}), []int{4, 3, 2, 1}))
	assert(t, "reverseInPlace negatives", reflect.DeepEqual(reverseInPlace([]int{-1, 0, 1}), []int{1, 0, -1}))
	assert(t, "reverseInPlace palindrome", reflect.DeepEqual(reverseInPlace([]int{1, 2, 1}), []int{1, 2, 1}))
}

func TestIndexOfMax(t *testing.T) {
	assert(t, "indexOfMax basic", indexOfMax([]int{3, 1, 4, 4}) == 2)
	assert(t, "indexOfMax single", indexOfMax([]int{5}) == 0)
	assert(t, "indexOfMax empty", indexOfMax([]int{}) == 0)
	assert(t, "indexOfMax ties", indexOfMax([]int{5, 5, 5}) == 0)
	assert(t, "indexOfMax at start", indexOfMax([]int{9, 1, 2}) == 0)
	assert(t, "indexOfMax at end", indexOfMax([]int{1, 2, 9}) == 2)
	assert(t, "indexOfMax negatives", indexOfMax([]int{-10, -3, -7}) == 1)
}

func TestArraySum(t *testing.T) {
	assert(t, "arraySum basic", arraySum([]int{1, 2, 3, 4}) == 10)
	assert(t, "arraySum empty", arraySum([]int{}) == 0)
	assert(t, "arraySum single", arraySum([]int{5}) == 5)
	assert(t, "arraySum negatives", arraySum([]int{-1, 2, -3}) == -2)
	assert(t, "arraySum all negative", arraySum([]int{-2, -3}) == -5)
	assert(t, "arraySum zeros", arraySum([]int{0, 0, 0}) == 0)
}

func TestRotateRight(t *testing.T) {
	assert(t, "rotateRight k=1", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4}, 1), []int{4, 1, 2, 3}))
	assert(t, "rotateRight k=2", reflect.DeepEqual(rotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}))
	assert(t, "rotateRight k=0", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 0), []int{1, 2, 3}))
	assert(t, "rotateRight k=len-1", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 2), []int{2, 3, 1}))
	assert(t, "rotateRight k=len", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 3), []int{1, 2, 3}))
	assert(t, "rotateRight k>len", reflect.DeepEqual(rotateRight([]int{1, 2}, 5), []int{2, 1}))
	assert(t, "rotateRight k double cycle", reflect.DeepEqual(rotateRight([]int{1, 2, 3}, 6), []int{1, 2, 3}))
	assert(t, "rotateRight empty", reflect.DeepEqual(rotateRight([]int{}, 3), []int{}))
	assert(t, "rotateRight single", reflect.DeepEqual(rotateRight([]int{7}, 5), []int{7}))
}

func TestRunningSum(t *testing.T) {
	assert(t, "runningSum basic", reflect.DeepEqual(runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}))
	assert(t, "runningSum single", reflect.DeepEqual(runningSum([]int{5}), []int{5}))
	assert(t, "runningSum empty", reflect.DeepEqual(runningSum([]int{}), []int{}))
	assert(t, "runningSum negatives", reflect.DeepEqual(runningSum([]int{1, -1, 2}), []int{1, 0, 2}))
	assert(t, "runningSum zeros", reflect.DeepEqual(runningSum([]int{0, 0, 0}), []int{0, 0, 0}))
	assert(t, "runningSum constant", reflect.DeepEqual(runningSum([]int{2, 2, 2}), []int{2, 4, 6}))
}

func TestAll(t *testing.T) {
	t.Run("reverseInPlace", TestReverseInPlace)
	t.Run("indexOfMax", TestIndexOfMax)
	t.Run("arraySum", TestArraySum)
	t.Run("rotateRight", TestRotateRight)
	t.Run("runningSum", TestRunningSum)
	fmt.Println("\nAll array reflex drills passed.")
}
