package main

import (
	"sort"
	"strconv"
)

var reflexBySlug = map[string]string{
	"two-sum":                                        "twoSum",
	"contains-duplicate":                             "containsDuplicate",
	"group-anagrams":                                 "groupAnagrams",
	"first-unique-character-in-a-string":             "firstUniqueChar",
	"valid-palindrome":                               "isPalindrome",
	"container-with-most-water":                      "maxArea",
	"remove-duplicates-from-sorted-array":            "removeDuplicates",
	"move-zeroes":                                    "moveZeroes",
	"binary-search":                                  "binarySearch",
	"search-insert-position":                         "searchInsert",
	"find-minimum-in-rotated-sorted-array":           "findMinRotated",
	"valid-parentheses":                              "isValidParentheses",
	"binary-tree-inorder-traversal":                  "inorderTraversal",
	"binary-tree-preorder-traversal":                 "preorderTraversal",
	"binary-tree-postorder-traversal":                "postorderTraversal",
	"binary-tree-level-order-traversal":              "levelOrderTraversal",
	"maximum-depth-of-binary-tree":                   "maxDepth",
	"daily-temperatures":                             "dailyTemperatures",
	"climbing-stairs":                                "climbStairs",
	"house-robber":                                   "rob",
	"min-cost-climbing-stairs":                       "minCostClimbingStairs",
	"number-of-islands":                              "numIslands",
	"flood-fill":                                     "floodFill",
	"find-if-path-exists-in-graph":                   "shortestPathGrid",
	"subarray-sum-equals-k":                          "subarraySumK",
	"product-of-array-except-self":                   "productExceptSelf",
	"maximum-subarray":                               "maxSubarraySum",
	"longest-substring-without-repeating-characters": "longestUniqueSubstring",
	"koko-eating-bananas":                            "minEatingSpeed",
	"validate-binary-search-tree":                    "isValidBST",
	"coin-change":                                    "coinChange",
	"course-schedule":                                "canFinish",
	"reverse-linked-list":                            "reverseList",
	"linked-list-cycle":                              "hasCycle",
	"middle-of-the-linked-list":                      "middleNode",
	"merge-two-sorted-lists":                         "mergeTwoLists",
	"remove-nth-node-from-end-of-list":               "removeNthFromEnd",
	"rotate-array":                                   "rotateRight",
	// Kept in step with bin/study_play/primaries.go, so a fetched problem maps
	// back to the reflex function that owns its move.
	"reverse-string":                  "reverseInPlace",
	"find-pivot-index":                "runningSum",
	"majority-element":                "frequencyMap",
	"maximum-average-subarray-i":      "maxSumSubarrayK",
	"kth-largest-element-in-an-array": "kthLargest",
	"last-stone-weight":               "lastStoneWeight",
	"merge-k-sorted-lists":            "mergeKSorted",
	"subsets":                         "subsets",
	"permutations":                    "permute",
	"combinations":                    "combine",
	"shortest-path-in-binary-matrix":  "shortestPathGrid",
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
