// CORE 5 — daily reflex essentials (run every session)
//
// GOAL: All five under 8 minutes from memory.
// RUN:  go run ./core5
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — return indices of two values that sum to target
func twoSum(nums []int, target int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — return index of target in sorted array, or -1
func binarySearch(nums []int, target int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — in-place unique prefix length on sorted array
func removeDuplicates(nums []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — max sum of any contiguous window of size k
func maxSumSubarrayK(nums []int, k int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — map each value to its count
func frequencyMap(arr []string) map[string]int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert("twoSum", len(got) == 2 && ((got[0] == 0 && got[1] == 1) || (got[0] == 1 && got[1] == 0)))
	assert("twoSum negatives", reflect.DeepEqual(twoSum([]int{-1, -2, -3, -4, -5}, -8), []int{2, 4}))

	assert("binarySearch found", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert("binarySearch missing", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert("binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert("binarySearch single", binarySearch([]int{5}, 5) == 0)

	dup := []int{1, 1, 2, 2, 3}
	assert("removeDuplicates len", removeDuplicates(dup) == 3)
	assert("removeDuplicates vals", reflect.DeepEqual(dup[:3], []int{1, 2, 3}))
	assert("removeDuplicates empty", removeDuplicates([]int{}) == 0)
	assert("removeDuplicates single", removeDuplicates([]int{7}) == 1)

	assert("maxSumSubarrayK", maxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3) == 9)
	assert("maxSumSubarrayK k=1", maxSumSubarrayK([]int{4, 2, 9}, 1) == 9)

	freq := frequencyMap([]string{"a", "b", "a", "c"})
	assert("frequencyMap", freq["a"] == 2 && freq["b"] == 1 && freq["c"] == 1)
	assert("frequencyMap empty", len(frequencyMap([]string{})) == 0)

	fmt.Println("\nCore 5 passed. Primary: hashing/easy/two_sum.js")
}
