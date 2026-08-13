#!/usr/bin/env python3
"""Patch reflex drill main() blocks with comprehensive test variations."""

from __future__ import annotations

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]

DRILLS = [
    "01_arrays_reflex",
    "02_hashing_reflex",
    "03_two_pointers_reflex",
    "04_binary_search_reflex",
    "05_trees_stacks_reflex",
    "06_dp_reflex",
    "07_graphs_reflex",
    "08_heap_reflex",
    "09_backtrack_reflex",
    "10_math_reflex",
]

WRITE_MAINS: dict[str, str] = {}
SOLUTION_MAINS: dict[str, str] = {}

# fmt: off
WRITE_MAINS["01_arrays_reflex"] = r'''
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
'''

WRITE_MAINS["02_hashing_reflex"] = r'''
func main() {
	// twoSum — basic, negatives, duplicates, zero, both negative, distant indices
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert("twoSum basic", len(got) == 2 && got[0] == 0 && got[1] == 1)
	assert("twoSum negatives", reflectDeepEqual(twoSum([]int{-1, -2, -3, -4, -5}, -8), []int{2, 4}))
	assert("twoSum duplicates", reflectDeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}))
	assert("twoSum zero", reflectDeepEqual(twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3}))
	assert("twoSum both negative", reflectDeepEqual(twoSum([]int{-3, 4, 3, 90}, 0), []int{0, 2}))
	assert("twoSum distant", reflectDeepEqual(twoSum([]int{1, 2, 3, 4, 5}, 9), []int{3, 4}))

	// containsDuplicate — found, not found, single, empty, pair, triple run
	assert("containsDuplicate true", containsDuplicate([]int{1, 2, 3, 1}) == true)
	assert("containsDuplicate false", containsDuplicate([]int{1, 2, 3, 4}) == false)
	assert("containsDuplicate single", containsDuplicate([]int{1}) == false)
	assert("containsDuplicate empty", containsDuplicate([]int{}) == false)
	assert("containsDuplicate pair", containsDuplicate([]int{2, 2}) == true)
	assert("containsDuplicate triple", containsDuplicate([]int{1, 1, 1}) == true)

	// frequencyMap — multi, empty, single, repeated, distinct
	freq := frequencyMap([]string{"a", "b", "a", "c"})
	assert("frequencyMap basic", freq["a"] == 2 && freq["b"] == 1 && freq["c"] == 1)
	assert("frequencyMap empty", len(frequencyMap([]string{})) == 0)
	assert("frequencyMap single", frequencyMap([]string{"x"})["x"] == 1)
	assert("frequencyMap all same", frequencyMap([]string{"z", "z", "z"})["z"] == 3)
	assert("frequencyMap distinct", len(frequencyMap([]string{"a", "b", "c"})) == 3)

	// firstUniqueChar — middle, none, single, empty, last unique
	assert("firstUniqueChar basic", firstUniqueChar("leetcode") == "l")
	assert("firstUniqueChar none", firstUniqueChar("aabb") == "")
	assert("firstUniqueChar single", firstUniqueChar("z") == "z")
	assert("firstUniqueChar empty", firstUniqueChar("") == "")
	assert("firstUniqueChar last", firstUniqueChar("aabbcd") == "c")

	// groupAnagrams — multi-group, empty, single, no shared, all anagrams
	groups := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	ok := len(groups) == 3
	if ok {
		found := false
		for _, g := range groups {
			if len(g) == 3 {
				for _, w := range g {
					if w == "eat" {
						found = true
					}
				}
			}
		}
		ok = found
	}
	assert("groupAnagrams basic", ok)
	assert("groupAnagrams empty", len(groupAnagrams([]string{})) == 0)
	assert("groupAnagrams single", len(groupAnagrams([]string{"abc"})) == 1)
	assert("groupAnagrams no shared", len(groupAnagrams([]string{"ab", "cd", "ef"})) == 3)
	allA := groupAnagrams([]string{"abc", "bca", "cab"})
	assert("groupAnagrams all anagrams", len(allA) == 1 && len(allA[0]) == 3)

	fmt.Println("\nAll hashing reflex drills passed.")
'''

WRITE_MAINS["03_two_pointers_reflex"] = r'''
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
'''

WRITE_MAINS["04_binary_search_reflex"] = r'''
func main() {
	// binarySearch — hit/miss, empty, single, first/last/mid, two-element
	assert("binarySearch found mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert("binarySearch missing mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert("binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert("binarySearch single hit", binarySearch([]int{5}, 5) == 0)
	assert("binarySearch single miss", binarySearch([]int{5}, 2) == -1)
	assert("binarySearch first", binarySearch([]int{1, 2, 3, 4, 5}, 1) == 0)
	assert("binarySearch last", binarySearch([]int{1, 2, 3, 4, 5}, 5) == 4)
	assert("binarySearch two hit", binarySearch([]int{1, 3}, 3) == 1)
	assert("binarySearch two miss", binarySearch([]int{1, 3}, 2) == -1)

	// searchInsert — exist, middle, end, start, empty, single, equal front
	assert("searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert("searchInsert new mid", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert("searchInsert end", searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	assert("searchInsert empty", searchInsert([]int{}, 5) == 0)
	assert("searchInsert start", searchInsert([]int{2, 4, 6}, 1) == 0)
	assert("searchInsert single", searchInsert([]int{5}, 5) == 0)
	assert("searchInsert after single", searchInsert([]int{5}, 7) == 1)

	// findMinRotated — pivot, two elem, sorted, single, pivot at end
	assert("findMinRotated pivot", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert("findMinRotated two", findMinRotated([]int{3, 1}) == 1)
	assert("findMinRotated sorted", findMinRotated([]int{1, 2, 3, 4}) == 1)
	assert("findMinRotated single", findMinRotated([]int{2}) == 2)
	assert("findMinRotated pivot end", findMinRotated([]int{2, 3, 4, 5, 1}) == 1)

	// isTargetPresent — mirrors search outcomes
	assert("isTargetPresent true mid", isTargetPresent([]int{1, 2, 3, 4, 5}, 3) == true)
	assert("isTargetPresent false", isTargetPresent([]int{1, 2, 3, 4, 5}, 6) == false)
	assert("isTargetPresent empty", isTargetPresent([]int{}, 1) == false)
	assert("isTargetPresent first", isTargetPresent([]int{1, 2, 3}, 1) == true)
	assert("isTargetPresent last", isTargetPresent([]int{1, 2, 3}, 3) == true)

	fmt.Println("\nAll binary search reflex drills passed.")
'''

WRITE_MAINS["05_trees_stacks_reflex"] = r'''
func main() {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}},
	}
	single := &TreeNode{Val: 5}
	leftSkew := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	// inorderTraversal — full, nil, single, left skew
	assert("inorderTraversal basic", reflect.DeepEqual(inorderTraversal(tree), []int{2, 1, 4, 3}))
	assert("inorderTraversal nil", reflect.DeepEqual(inorderTraversal(nil), []int{}))
	assert("inorderTraversal single", reflect.DeepEqual(inorderTraversal(single), []int{5}))
	assert("inorderTraversal left skew", reflect.DeepEqual(inorderTraversal(leftSkew), []int{3, 2, 1}))

	// maxDepth — full, nil, single, left skew
	assert("maxDepth basic", maxDepth(tree) == 3)
	assert("maxDepth nil", maxDepth(nil) == 0)
	assert("maxDepth single", maxDepth(single) == 1)
	assert("maxDepth left skew", maxDepth(leftSkew) == 3)

	// isValidParentheses — valid types, invalid, empty, unmatched open/close, nested, interleaved
	assert("isValidParentheses basic", isValidParentheses("()[]{}") == true)
	assert("isValidParentheses invalid", isValidParentheses("(]") == false)
	assert("isValidParentheses empty", isValidParentheses("") == true)
	assert("isValidParentheses open only", isValidParentheses("(") == false)
	assert("isValidParentheses close only", isValidParentheses(")") == false)
	assert("isValidParentheses nested", isValidParentheses("((()))") == true)
	assert("isValidParentheses interleaved false", isValidParentheses("([)]") == false)
	assert("isValidParentheses mixed valid", isValidParentheses("{[()()]}") == true)

	// dailyTemperatures — classic, single, decreasing, equal, increasing, pair
	assert("dailyTemperatures basic", reflect.DeepEqual(
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	))
	assert("dailyTemperatures single", reflect.DeepEqual(dailyTemperatures([]int{50}), []int{0}))
	assert("dailyTemperatures decreasing", reflect.DeepEqual(dailyTemperatures([]int{5, 4, 3}), []int{0, 0, 0}))
	assert("dailyTemperatures equal", reflect.DeepEqual(dailyTemperatures([]int{70, 70, 70}), []int{0, 0, 0}))
	assert("dailyTemperatures increasing", reflect.DeepEqual(dailyTemperatures([]int{60, 61, 62}), []int{1, 1, 0}))
	assert("dailyTemperatures pair", reflect.DeepEqual(dailyTemperatures([]int{55, 56}), []int{1, 0}))

	fmt.Println("\nAll trees/stacks reflex drills passed.")
'''

WRITE_MAINS["06_dp_reflex"] = r'''
func main() {
	// fib — base cases and several n
	assert("fib ten", fib(10) == 55)
	assert("fib zero", fib(0) == 0)
	assert("fib one", fib(1) == 1)
	assert("fib two", fib(2) == 1)
	assert("fib three", fib(3) == 2)
	assert("fib five", fib(5) == 5)
	assert("fib twenty", fib(20) == 6765)

	// minCostClimbingStairs — 3-step, 2-step, single, equal, cheap first
	assert("minCostClimbingStairs basic", minCostClimbingStairs([]int{10, 15, 20}) == 15)
	assert("minCostClimbingStairs two", minCostClimbingStairs([]int{1, 100}) == 1)
	assert("minCostClimbingStairs single", minCostClimbingStairs([]int{5}) == 0)
	assert("minCostClimbingStairs equal", minCostClimbingStairs([]int{5, 5, 5}) == 5)
	assert("minCostClimbingStairs cheap start", minCostClimbingStairs([]int{0, 1, 1, 1}) == 1)

	// rob — classic, single/two, empty, alternating, all same, skip middle
	assert("rob basic", rob([]int{2, 7, 9, 3, 1}) == 12)
	assert("rob single", rob([]int{5}) == 5)
	assert("rob two pick max", rob([]int{2, 1}) == 2)
	assert("rob empty", rob([]int{}) == 0)
	assert("rob alternating", rob([]int{5, 1, 5, 1}) == 10)
	assert("rob all same", rob([]int{3, 3, 3}) == 6)
	assert("rob endpoints", rob([]int{2, 1, 2}) == 4)

	// climbStairs — 0..6
	assert("climbStairs five", climbStairs(5) == 8)
	assert("climbStairs one", climbStairs(1) == 1)
	assert("climbStairs two", climbStairs(2) == 2)
	assert("climbStairs three", climbStairs(3) == 3)
	assert("climbStairs zero", climbStairs(0) == 0)
	assert("climbStairs four", climbStairs(4) == 5)
	assert("climbStairs six", climbStairs(6) == 13)

	fmt.Println("\nAll DP reflex drills passed.")
'''

WRITE_MAINS["07_graphs_reflex"] = r'''
func main() {
	grid1 := [][]byte{
		{'1', '1', '0'},
		{'0', '1', '0'},
		{'1', '0', '1'},
	}

	// numIslands — multi, all water/land, empty, single, diagonal, row
	assert("numIslands basic", numIslands(grid1) == 3)
	assert("numIslands all water", numIslands([][]byte{{'0'}}) == 0)
	assert("numIslands all land", numIslands([][]byte{{'1', '1'}, {'1', '1'}}) == 1)
	assert("numIslands empty", numIslands([][]byte{}) == 0)
	assert("numIslands single land", numIslands([][]byte{{'1'}}) == 1)
	assert("numIslands diagonal", numIslands([][]byte{{'1', '0'}, {'0', '1'}}) == 2)
	assert("numIslands row", numIslands([][]byte{{'1', '1', '1', '0', '1'}}) == 2)

	img := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	want := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	// floodFill — region, same color, single pixel, corner, isolated pixel
	assert("floodFill basic", reflect.DeepEqual(floodFill(cloneGrid(img), 1, 1, 2), want))
	sameColor := [][]int{{3}}
	assert("floodFill same color", reflect.DeepEqual(floodFill(cloneGrid(sameColor), 0, 0, 3), sameColor))
	singlePixel := [][]int{{0}}
	assert("floodFill single", reflect.DeepEqual(floodFill(cloneGrid(singlePixel), 0, 0, 9), [][]int{{9}}))
	corner := [][]int{{1, 0}, {0, 0}}
	assert("floodFill corner", reflect.DeepEqual(floodFill(cloneGrid(corner), 0, 0, 7), [][]int{{7, 0}, {0, 0}}))
	isolated := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	assert("floodFill isolated", reflect.DeepEqual(floodFill(cloneGrid(isolated), 1, 1, 9), [][]int{{1, 0, 1}, {0, 9, 0}, {1, 0, 1}}))

	pathGrid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}

	// shortestPathGrid — detour, blocked, single, start/end blocked, straight, end wall
	assert("shortestPathGrid detour", shortestPathGrid(pathGrid) == 4)
	assert("shortestPathGrid blocked", shortestPathGrid([][]int{{0, 1}, {1, 0}}) == -1)
	assert("shortestPathGrid single", shortestPathGrid([][]int{{0}}) == 1)
	assert("shortestPathGrid start blocked", shortestPathGrid([][]int{{1}}) == -1)
	openPath := [][]int{{0, 0, 0, 0}}
	assert("shortestPathGrid straight", shortestPathGrid(openPath) == 3)
	endBlocked := [][]int{{0, 0}, {0, 1}}
	assert("shortestPathGrid end blocked", shortestPathGrid(endBlocked) == -1)
	open2x2 := [][]int{{0, 0}, {0, 0}}
	assert("shortestPathGrid open 2x2", shortestPathGrid(open2x2) == 2)

	fmt.Println("\nAll graph reflex drills passed.")
'''

WRITE_MAINS["08_heap_reflex"] = r'''
func main() {
	// kthLargest — general, duplicates, k=1, single, k=len, negatives
	assert("kthLargest basic", kthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	assert("kthLargest dup", kthLargest([]int{3, 3, 3, 3}, 2) == 3)
	assert("kthLargest k=1", kthLargest([]int{1, 2, 3}, 1) == 3)
	assert("kthLargest single", kthLargest([]int{42}, 1) == 42)
	assert("kthLargest k=len", kthLargest([]int{4, 2, 9}, 3) == 2)
	assert("kthLargest negatives", kthLargest([]int{-1, -2, -3}, 2) == -2)

	// lastStoneWeight — classic, single, equal cancel, chain cancel, three equal
	assert("lastStoneWeight basic", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
	assert("lastStoneWeight single", lastStoneWeight([]int{5}) == 5)
	assert("lastStoneWeight cancel", lastStoneWeight([]int{5, 5}) == 0)
	assert("lastStoneWeight chain", lastStoneWeight([]int{6, 3, 3}) == 0)
	assert("lastStoneWeight three equal", lastStoneWeight([]int{4, 4, 4}) == 4)

	// mergeKSorted — multi, empty outer, one list, empty inner, full merge, duplicates
	got := mergeKSorted([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	assert("mergeKSorted basic", len(got) == 8 && got[0] == 1 && got[len(got)-1] == 6)
	assert("mergeKSorted empty", len(mergeKSorted([][]int{})) == 0)
	assert("mergeKSorted one list", reflectDeepEqual(mergeKSorted([][]int{{1, 2}}), []int{1, 2}))
	assert("mergeKSorted with empty list", reflectDeepEqual(mergeKSorted([][]int{{}, {1, 3}, {2}}), []int{1, 2, 3}))
	assert("mergeKSorted full", reflectDeepEqual(mergeKSorted([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4}))
	assert("mergeKSorted duplicates", reflectDeepEqual(mergeKSorted([][]int{{1, 1}, {1}}), []int{1, 1, 1}))

	fmt.Println("\nAll heap reflex drills passed.")
'''

WRITE_MAINS["09_backtrack_reflex"] = r'''
func main() {
	// subsets — counts and explicit contents for n=0,1,2,3
	assert("subsets n=3 count", len(subsets([]int{1, 2, 3})) == 8)
	assert("subsets empty input", reflect.DeepEqual(subsets([]int{}), [][]int{{}}))
	assert("subsets single", reflect.DeepEqual(subsets([]int{1}), [][]int{{}, {1}}))
	assert("subsets n=2 contents", reflect.DeepEqual(
		sortedSlices(subsets([]int{1, 2})),
		sortedSlices([][]int{{}, {1}, {2}, {1, 2}}),
	))

	// permute — counts and explicit for n=1,2,3
	assert("permute n=3 count", len(permute([]int{1, 2, 3})) == 6)
	assert("permute single", reflect.DeepEqual(permute([]int{7}), [][]int{{7}}))
	assert("permute n=2 contents", reflect.DeepEqual(
		sortedSlices(permute([]int{1, 2})),
		sortedSlices([][]int{{1, 2}, {2, 1}}),
	))

	// combine — varied n,k including k=0, k=1, k=n, k>n invalid
	combs := combine(4, 2)
	assert("combine n=4 k=2 count", len(combs) == 6)
	assert("combine n=4 k=2 sample", reflect.DeepEqual(sortedSlices(combs)[0], []int{1, 2}))
	assert("combine k=1", len(combine(3, 1)) == 3)
	assert("combine k=n", len(combine(3, 3)) == 1)
	assert("combine n=1", reflect.DeepEqual(combine(1, 1), [][]int{{1}}))
	assert("combine k=0", reflect.DeepEqual(combine(4, 0), [][]int{{}}))
	assert("combine k=2 n=2", reflect.DeepEqual(sortedSlices(combine(2, 2)), sortedSlices([][]int{{1, 2}})))

	fmt.Println("\nAll backtracking reflex drills passed.")
'''

WRITE_MAINS["10_math_reflex"] = r'''
func main() {
	// gcd — basic, coprime, zero sides, both zero, negatives
	assert("gcd basic", gcd(48, 18) == 6)
	assert("gcd coprime", gcd(17, 13) == 1)
	assert("gcd zero b", gcd(0, 7) == 7)
	assert("gcd zero a", gcd(7, 0) == 7)
	assert("gcd both zero", gcd(0, 0) == 0)
	assert("gcd negatives", gcd(-48, 18) == 6)
	assert("gcd both negative", gcd(-48, -18) == 6)

	// lcm — basic, coprime, zero, same number
	assert("lcm basic", lcm(4, 6) == 12)
	assert("lcm coprime", lcm(7, 11) == 77)
	assert("lcm zero", lcm(0, 5) == 0)
	assert("lcm same", lcm(8, 8) == 8)

	// modPow — small, base case, mod 1, exp 0/1, even exp
	assert("modPow small", modPow(2, 10, 1000) == 24)
	assert("modPow base", modPow(3, 4, 100) == 81)
	assert("modPow mod1", modPow(5, 100, 1) == 0)
	assert("modPow exp0", modPow(2, 0, 100) == 1)
	assert("modPow exp1", modPow(7, 1, 10) == 7)
	assert("modPow even exp", modPow(2, 8, 100) == 56)

	// nCr — basic, symmetry, invalid, k=0, k=n, n=0
	assert("nCr basic", nCr(5, 2) == 10)
	assert("nCr symmetry", nCr(10, 8) == 45)
	assert("nCr invalid high", nCr(5, 6) == 0)
	assert("nCr invalid negative", nCr(5, -1) == 0)
	assert("nCr k0", nCr(5, 0) == 1)
	assert("nCr kn", nCr(5, 5) == 1)
	assert("nCr n0k0", nCr(0, 0) == 1)

	// isPrime — 2, odd prime, composite, 1, 0, square, even composite
	assert("isPrime seventeen", isPrime(17) == true)
	assert("isPrime one", isPrime(1) == false)
	assert("isPrime composite", isPrime(15) == false)
	assert("isPrime two", isPrime(2) == true)
	assert("isPrime zero", isPrime(0) == false)
	assert("isPrime square", isPrime(9) == false)
	assert("isPrime even composite", isPrime(4) == false)

	// powOfTwo — powers, non-powers, zero/negative
	assert("powOfTwo sixtyfour", powOfTwo(64) == true)
	assert("powOfTwo six", powOfTwo(6) == false)
	assert("powOfTwo zero", powOfTwo(0) == false)
	assert("powOfTwo one", powOfTwo(1) == true)
	assert("powOfTwo two", powOfTwo(2) == true)
	assert("powOfTwo eight", powOfTwo(8) == true)
	assert("powOfTwo negative", powOfTwo(-4) == false)

	fmt.Println("\nAll math reflex drills passed.")
'''
# fmt: on

FOOTERS = {
    "01_arrays_reflex": '\tfmt.Println("Primary: arrays/easy/plus_one.js")\n',
    "02_hashing_reflex": '\tfmt.Println("Primary: hashing/easy/two_sum.js")\n',
    "03_two_pointers_reflex": '\tfmt.Println("Primary: two_pointers/easy/move_zeroes.js")\n',
    "04_binary_search_reflex": '\tfmt.Println("Primary: binary_search/easy/search_insertion_position.js")\n',
    "05_trees_stacks_reflex": '\tfmt.Println("Primary: stacks/easy/valid_parentheses.js")\n',
    "06_dp_reflex": '\tfmt.Println("Primary: dynamic_programming/easy/fibonacci_number.js")\n',
    "07_graphs_reflex": '\tfmt.Println("Primary: graphs/medium/number_of_islands.js")\n',
    "08_heap_reflex": '\tfmt.Println("Primary: heaps/medium/kth_largest_element_in_an_array.js")\n',
    "09_backtrack_reflex": '\tfmt.Println("Primary: backtracking/medium/subsets.js")\n',
    "10_math_reflex": '\tfmt.Println("Reference: doc/write/MATH_CONCEPTS.md")\n',
}


def to_solution_main(write_main: str) -> str:
    main = write_main
    # bool comparisons -> direct bool expressions for solution style
    main = re.sub(r"containsDuplicate\(([^)]+)\) == true", r"containsDuplicate(\1)", main)
    main = re.sub(r"containsDuplicate\(([^)]+)\) == false", r"!containsDuplicate(\1)", main)
    main = re.sub(r"isPalindrome\(([^)]+)\) == true", r"isPalindrome(\1)", main)
    main = re.sub(r"isPalindrome\(([^)]+)\) == false", r"!isPalindrome(\1)", main)
    main = re.sub(r"isTargetPresent\(([^)]+)\) == true", r"isTargetPresent(\1)", main)
    main = re.sub(r"isTargetPresent\(([^)]+)\) == false", r"!isTargetPresent(\1)", main)
    main = re.sub(r"isValidParentheses\(([^)]+)\) == true", r"isValidParentheses(\1)", main)
    main = re.sub(r"isValidParentheses\(([^)]+)\) == false", r"!isValidParentheses(\1)", main)
    main = re.sub(r"isPrime\(([^)]+)\) == true", r"isPrime(\1)", main)
    main = re.sub(r"isPrime\(([^)]+)\) == false", r"!isPrime(\1)", main)
    main = re.sub(r"powOfTwo\(([^)]+)\) == true", r"powOfTwo(\1)", main)
    main = re.sub(r"powOfTwo\(([^)]+)\) == false", r"!powOfTwo(\1)", main)
    return main


for name in WRITE_MAINS:
    SOLUTION_MAINS[name] = to_solution_main(WRITE_MAINS[name])


def patch_main(path: Path, new_main: str) -> None:
    text = path.read_text()
    pattern = re.compile(r"func main\(\) \{.*?\n\}", re.DOTALL)
    if not pattern.search(text):
        raise SystemExit(f"no main() found in {path}")
    updated = pattern.sub(lambda _m: new_main.strip(), text, count=1)
    path.write_text(updated)


def main() -> None:
    for drill in DRILLS:
        write_main = WRITE_MAINS[drill].strip()
        if drill in FOOTERS:
            write_main += "\n" + FOOTERS[drill].rstrip()
        write_main += "\n}"

        solution_main = SOLUTION_MAINS[drill].strip()
        if drill in FOOTERS:
            solution_main += "\n" + FOOTERS[drill].rstrip()
        solution_main += "\n}"

        write_path = ROOT / "drills/write/reflex" / drill / "main.go"
        blank_path = ROOT / "bin/study_play/_support/blanks" / f"{drill}.go"
        solution_path = ROOT / "drills/solutions/reflex" / drill / "main.go"

        for path, body in [
            (write_path, write_main),
            (blank_path, write_main),
            (solution_path, solution_main),
        ]:
            patch_main(path, body)
            print(f"patched {path.relative_to(ROOT)}")


if __name__ == "__main__":
    main()
