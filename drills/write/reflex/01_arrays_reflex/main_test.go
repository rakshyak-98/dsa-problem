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

func TestSubarraySumK(t *testing.T) {
	assert(t, "subarraySumK basic", subarraySumK([]int{1, 1, 1}, 2) == 2)
	assert(t, "subarraySumK whole", subarraySumK([]int{1, 2, 3}, 6) == 1)
	assert(t, "subarraySumK none", subarraySumK([]int{1, 2, 3}, 7) == 0)
	assert(t, "subarraySumK empty", subarraySumK([]int{}, 0) == 0)
	assert(t, "subarraySumK negatives", subarraySumK([]int{1, -1, 0}, 0) == 3)
	assert(t, "subarraySumK zeros", subarraySumK([]int{0, 0, 0}, 0) == 6)
	assert(t, "subarraySumK single hit", subarraySumK([]int{3}, 3) == 1)
}

func TestProductExceptSelf(t *testing.T) {
	assert(t, "productExceptSelf basic", reflect.DeepEqual(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}))
	assert(t, "productExceptSelf one zero", reflect.DeepEqual(productExceptSelf([]int{1, 0, 3}), []int{0, 3, 0}))
	assert(t, "productExceptSelf two zeros", reflect.DeepEqual(productExceptSelf([]int{0, 0, 3}), []int{0, 0, 0}))
	assert(t, "productExceptSelf negatives", reflect.DeepEqual(productExceptSelf([]int{-1, 2, -3}), []int{-6, 3, -2}))
	assert(t, "productExceptSelf two", reflect.DeepEqual(productExceptSelf([]int{2, 5}), []int{5, 2}))
	assert(t, "productExceptSelf single", reflect.DeepEqual(productExceptSelf([]int{9}), []int{1}))
}

func TestMaxSubarraySum(t *testing.T) {
	assert(t, "maxSubarraySum basic", maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	assert(t, "maxSubarraySum all negative", maxSubarraySum([]int{-3, -1, -2}) == -1)
	assert(t, "maxSubarraySum all positive", maxSubarraySum([]int{1, 2, 3}) == 6)
	assert(t, "maxSubarraySum single", maxSubarraySum([]int{5}) == 5)
	assert(t, "maxSubarraySum empty", maxSubarraySum([]int{}) == 0)
	assert(t, "maxSubarraySum restart", maxSubarraySum([]int{-5, 8, -1, 3}) == 10)
}

func TestMergeSort(t *testing.T) {
	assert(t, "mergeSort basic", reflect.DeepEqual(mergeSort([]int{5, 2, 3, 1}), []int{1, 2, 3, 5}))
	assert(t, "mergeSort empty", reflect.DeepEqual(mergeSort([]int{}), []int{}))
	assert(t, "mergeSort single", reflect.DeepEqual(mergeSort([]int{7}), []int{7}))
	assert(t, "mergeSort duplicates", reflect.DeepEqual(mergeSort([]int{5, 1, 1, 2, 0, 0}), []int{0, 0, 1, 1, 2, 5}))
	assert(t, "mergeSort already sorted", reflect.DeepEqual(mergeSort([]int{1, 2, 3, 4}), []int{1, 2, 3, 4}))
	assert(t, "mergeSort negatives", reflect.DeepEqual(mergeSort([]int{-3, 4, -1, 0, -2}), []int{-3, -2, -1, 0, 4}))
}

func TestQuickSort(t *testing.T) {
	assert(t, "quickSort basic", reflect.DeepEqual(quickSort([]int{5, 2, 3, 1}), []int{1, 2, 3, 5}))
	assert(t, "quickSort empty", reflect.DeepEqual(quickSort([]int{}), []int{}))
	assert(t, "quickSort single", reflect.DeepEqual(quickSort([]int{7}), []int{7}))
	assert(t, "quickSort duplicates", reflect.DeepEqual(quickSort([]int{3, 3, 3, 1, 2}), []int{1, 2, 3, 3, 3}))
	assert(t, "quickSort reverse", reflect.DeepEqual(quickSort([]int{9, 7, 5, 3, 1}), []int{1, 3, 5, 7, 9}))
	assert(t, "quickSort negatives", reflect.DeepEqual(quickSort([]int{-3, 4, -1, 0, -2}), []int{-3, -2, -1, 0, 4}))
}

func TestSieve(t *testing.T) {
	assert(t, "sieve to 10", reflect.DeepEqual(sieve(10), []int{2, 3, 5, 7}))
	assert(t, "sieve to 2", reflect.DeepEqual(sieve(2), []int{2}))
	assert(t, "sieve to 1", len(sieve(1)) == 0)
	assert(t, "sieve to 0", len(sieve(0)) == 0)
	assert(t, "sieve to 20", reflect.DeepEqual(sieve(20), []int{2, 3, 5, 7, 11, 13, 17, 19}))
	assert(t, "sieve prime bound", reflect.DeepEqual(sieve(13), []int{2, 3, 5, 7, 11, 13}))
}

func TestAll(t *testing.T) {
	t.Run("reverseInPlace", TestReverseInPlace)
	t.Run("rotateRight", TestRotateRight)
	t.Run("runningSum", TestRunningSum)
	t.Run("subarraySumK", TestSubarraySumK)
	t.Run("productExceptSelf", TestProductExceptSelf)
	t.Run("maxSubarraySum", TestMaxSubarraySum)
	t.Run("mergeSort", TestMergeSort)
	t.Run("quickSort", TestQuickSort)
	t.Run("sieve", TestSieve)
	fmt.Println("\nAll array reflex drills passed.")
}
