package main

import (
	"sort"
	"strconv"
)

var reflexBySlug = map[string]string{
	"two-sum":                             "twoSum",
	"contains-duplicate":                  "containsDuplicate",
	"group-anagrams":                      "groupAnagrams",
	"first-unique-character-in-a-string":  "firstUniqueChar",
	"valid-palindrome":                    "isPalindrome",
	"container-with-most-water":           "maxArea",
	"remove-duplicates-from-sorted-array": "removeDuplicates",
	"move-zeroes":                         "moveZeroes",
	"binary-search":                       "binarySearch",
	"search-insert-position":              "searchInsert",
	"find-minimum-in-rotated-sorted-array": "findMinRotated",
	"valid-parentheses":                   "isValidParentheses",
	"binary-tree-inorder-traversal":       "inorderTraversal",
	"binary-tree-preorder-traversal":      "preorderTraversal",
	"binary-tree-postorder-traversal":     "postorderTraversal",
	"binary-tree-level-order-traversal":   "levelOrderTraversal",
	"maximum-depth-of-binary-tree":        "maxDepth",
	"daily-temperatures":                  "dailyTemperatures",
	"climbing-stairs":                     "climbStairs",
	"house-robber":                        "rob",
	"min-cost-climbing-stairs":            "minCostClimbingStairs",
	"number-of-islands":                   "numIslands",
	"flood-fill":                          "floodFill",
	"find-if-path-exists-in-graph":        "shortestPathGrid",
	"rotate-array":                        "rotateRight",
}

func sortProblemsByDifficulty(ps []lcProblem) {
	order := map[string]int{"Easy": 0, "Medium": 1, "Hard": 2}
	sort.SliceStable(ps, func(i, j int) bool {
		di, dj := order[ps[i].Diff], order[ps[j].Diff]
		if di != dj {
			return di < dj
		}
		return ps[i].Num < ps[j].Num
	})
}

func parseQuestionNum(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
