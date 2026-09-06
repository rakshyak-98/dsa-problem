package main

import (
	"fmt"
	"sort"
)

// Primary problems
//
// Each drilled function gets one canonical problem to solve *after* the drill
// passes, plus one step up. The pair is chosen so the difficulty matches the
// level the function is graded at: an L1 function gets a problem that only
// needs the move typed out, an L3 function gets one where the move has to be
// recognised through a disguise first.
//
// These are LeetCode slugs rather than paths into reference/problems/, which
// indexes only a partial local mirror — bin/study_leetcode resolves slugs
// against the live problem set.
type primaryProblem struct {
	title  string
	slug   string
	diff   string // Easy | Medium | Hard
	nextUp string // the same pattern, one rung harder
}

var primaries = map[string]primaryProblem{
	// 01 — arrays & prefix
	"reverseInPlace":    {"Reverse String", "reverse-string", "Easy", "Reverse Words in a String II"},
	"rotateRight":       {"Rotate Array", "rotate-array", "Medium", "Rotate List"},
	"runningSum":        {"Find Pivot Index", "find-pivot-index", "Easy", "Subarray Sum Equals K"},
	"subarraySumK":      {"Subarray Sum Equals K", "subarray-sum-equals-k", "Medium", "Continuous Subarray Sum"},
	"productExceptSelf": {"Product of Array Except Self", "product-of-array-except-self", "Medium", "Maximum Product Subarray"},
	"maxSubarraySum":    {"Maximum Subarray", "maximum-subarray", "Medium", "Maximum Product Subarray"},

	// 02 — hashing
	"twoSum":            {"Two Sum", "two-sum", "Easy", "3Sum"},
	"containsDuplicate": {"Contains Duplicate", "contains-duplicate", "Easy", "Contains Duplicate II"},
	"frequencyMap":      {"Majority Element", "majority-element", "Easy", "Top K Frequent Elements"},
	"firstUniqueChar":   {"First Unique Character in a String", "first-unique-character-in-a-string", "Easy", "Longest Substring Without Repeating Characters"},
	"groupAnagrams":     {"Group Anagrams", "group-anagrams", "Medium", "Longest Consecutive Sequence"},

	// 03 — two pointers & window
	"removeDuplicates":       {"Remove Duplicates from Sorted Array", "remove-duplicates-from-sorted-array", "Easy", "Remove Duplicates from Sorted Array II"},
	"moveZeroes":             {"Move Zeroes", "move-zeroes", "Easy", "Sort Colors"},
	"maxArea":                {"Container With Most Water", "container-with-most-water", "Medium", "Trapping Rain Water"},
	"isPalindrome":           {"Valid Palindrome", "valid-palindrome", "Easy", "Valid Palindrome II"},
	"maxSumSubarrayK":        {"Maximum Average Subarray I", "maximum-average-subarray-i", "Easy", "Longest Substring Without Repeating Characters"},
	"longestUniqueSubstring": {"Longest Substring Without Repeating Characters", "longest-substring-without-repeating-characters", "Medium", "Minimum Window Substring"},

	// 04 — binary search
	"binarySearch":   {"Binary Search", "binary-search", "Easy", "Search in Rotated Sorted Array"},
	"searchInsert":   {"Search Insert Position", "search-insert-position", "Easy", "Find First and Last Position of Element in Sorted Array"},
	"findMinRotated": {"Find Minimum in Rotated Sorted Array", "find-minimum-in-rotated-sorted-array", "Medium", "Median of Two Sorted Arrays"},
	"minEatingSpeed": {"Koko Eating Bananas", "koko-eating-bananas", "Medium", "Capacity To Ship Packages Within D Days"},

	// 05 — trees & stacks
	"inorderTraversal":    {"Binary Tree Inorder Traversal", "binary-tree-inorder-traversal", "Easy", "Kth Smallest Element in a BST"},
	"preorderTraversal":   {"Binary Tree Preorder Traversal", "binary-tree-preorder-traversal", "Easy", "Serialize and Deserialize Binary Tree"},
	"postorderTraversal":  {"Binary Tree Postorder Traversal", "binary-tree-postorder-traversal", "Easy", "Diameter of Binary Tree"},
	"levelOrderTraversal": {"Binary Tree Level Order Traversal", "binary-tree-level-order-traversal", "Medium", "Binary Tree Right Side View"},
	"maxDepth":            {"Maximum Depth of Binary Tree", "maximum-depth-of-binary-tree", "Easy", "Balanced Binary Tree"},
	"isValidBST":          {"Validate Binary Search Tree", "validate-binary-search-tree", "Medium", "Recover Binary Search Tree"},
	"isValidParentheses":  {"Valid Parentheses", "valid-parentheses", "Easy", "Minimum Remove to Make Valid Parentheses"},
	"dailyTemperatures":   {"Daily Temperatures", "daily-temperatures", "Medium", "Largest Rectangle in Histogram"},

	// 06 — dp
	"climbStairs":           {"Climbing Stairs", "climbing-stairs", "Easy", "Decode Ways"},
	"minCostClimbingStairs": {"Min Cost Climbing Stairs", "min-cost-climbing-stairs", "Easy", "Triangle"},
	"rob":                   {"House Robber", "house-robber", "Medium", "House Robber II"},
	"coinChange":            {"Coin Change", "coin-change", "Medium", "Coin Change II"},

	// 07 — graphs
	"numIslands":       {"Number of Islands", "number-of-islands", "Medium", "Max Area of Island"},
	"floodFill":        {"Flood Fill", "flood-fill", "Easy", "Surrounded Regions"},
	"shortestPathGrid": {"Shortest Path in Binary Matrix", "shortest-path-in-binary-matrix", "Medium", "Rotting Oranges"},
	"canFinish":        {"Course Schedule", "course-schedule", "Medium", "Course Schedule II"},

	// 08 — heaps
	"kthLargest":      {"Kth Largest Element in an Array", "kth-largest-element-in-an-array", "Medium", "Kth Largest Element in a Stream"},
	"lastStoneWeight": {"Last Stone Weight", "last-stone-weight", "Easy", "Minimum Cost to Connect Sticks"},
	"mergeKSorted":    {"Merge k Sorted Lists", "merge-k-sorted-lists", "Hard", "Smallest Range Covering Elements from K Lists"},

	// 10 — linked lists
	"reverseList":      {"Reverse Linked List", "reverse-linked-list", "Easy", "Reverse Linked List II"},
	"hasCycle":         {"Linked List Cycle", "linked-list-cycle", "Easy", "Linked List Cycle II"},
	"middleNode":       {"Middle of the Linked List", "middle-of-the-linked-list", "Easy", "Reorder List"},
	"mergeTwoLists":    {"Merge Two Sorted Lists", "merge-two-sorted-lists", "Easy", "Merge k Sorted Lists"},
	"removeNthFromEnd": {"Remove Nth Node From End of List", "remove-nth-node-from-end-of-list", "Medium", "Remove Duplicates from Sorted List II"},

	// 09 — backtracking
	"subsets": {"Subsets", "subsets", "Medium", "Subsets II"},
	"permute": {"Permutations", "permutations", "Medium", "Permutations II"},
	"combine": {"Combinations", "combinations", "Medium", "Combination Sum"},

	// 01 — sorting & sieve (folded into the arrays drill)
	"mergeSort": {"Sort an Array", "sort-an-array", "Medium", "Count of Smaller Numbers After Self"},
	"quickSort": {"Sort an Array", "sort-an-array", "Medium", "Kth Largest Element in an Array"},
	"sieve":     {"Count Primes", "count-primes", "Medium", "Closest Prime Numbers in Range"},

	// 02 — XOR single number (folded into the hashing drill)
	"singleNumber": {"Single Number", "single-number", "Easy", "Single Number II"},

	// 03 — Dutch national flag (folded into the two-pointer drill)
	"dutchFlag": {"Sort Colors", "sort-colors", "Medium", "Wiggle Sort II"},

	// 04 — divide & conquer / log reduction (folded into the binary search drill)
	"quickSelect": {"Kth Largest Element in an Array", "kth-largest-element-in-an-array", "Medium", "K Closest Points to Origin"},
	"fastPow":     {"Pow(x, n)", "powx-n", "Medium", "Super Pow"},
	"gcd":         {"Find Greatest Common Divisor of Array", "find-greatest-common-divisor-of-array", "Easy", "Water and Jug Problem"},

	// 07 — adjacency-list graph algorithms (folded into the graphs drill)
	"dfs":             {"Number of Provinces", "number-of-provinces", "Medium", "Number of Connected Components in an Undirected Graph"},
	"bfs":             {"Snakes and Ladders", "snakes-and-ladders", "Medium", "Word Ladder"},
	"bfsShortestPath": {"Minimum Genetic Mutation", "minimum-genetic-mutation", "Medium", "Word Ladder"},
	"topoSort":        {"Course Schedule II", "course-schedule-ii", "Medium", "Alien Dictionary"},
	"dijkstra":        {"Network Delay Time", "network-delay-time", "Medium", "Cheapest Flights Within K Stops"},
}

func problemURL(slug string) string {
	return "https://leetcode.com/problems/" + slug + "/"
}

// printPrimary shows the problem to solve after the drill passes. Above L1 it
// also shows the step up, because a function you already own is only worth
// practising on a statement that hides it.
func printPrimary(fn string, earned level) {
	p, ok := primaries[fn]
	if !ok {
		return
	}
	fmt.Printf("     solve: %s (%s) %s\n", p.title, p.diff, problemURL(p.slug))
	if earned > levelRecall && p.nextUp != "" {
		fmt.Printf("     then:  %s\n", p.nextUp)
	}
}

// printProblemSet lists the curated problem for every function, ordered by the
// level it is currently graded at — the answer to "what should I actually
// solve next".
func printProblemSet(root string) {
	report := masteryReport(root)
	sort.Slice(report, func(i, j int) bool {
		if report[i].earned != report[j].earned {
			return report[i].earned < report[j].earned
		}
		pi, pj := drillPriority[report[i].cue.drill], drillPriority[report[j].cue.drill]
		if pi != pj {
			return pi < pj
		}
		return report[i].cue.fn < report[j].cue.fn
	})

	fmt.Println("PROBLEM SET — one primary per function, ordered by the level you are at")
	current := level(0)
	for _, m := range report {
		if m.earned != current {
			current = m.earned
			fmt.Printf("\n%s — solve these as written\n", current.label())
		}
		p, ok := primaries[m.cue.fn]
		if !ok {
			continue
		}
		fmt.Printf("  %-22s %-6s %-44s %s\n", m.cue.fn, p.diff, p.title, problemURL(p.slug))
		if m.earned > levelRecall {
			fmt.Printf("  %-22s %-6s step up: %s\n", "", "", p.nextUp)
		}
	}
	fmt.Println("\n  Level is earned from .drill_log.json: go run . -- --levels")
}
