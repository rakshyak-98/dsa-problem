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
	"reverseInPlace": {"Reverse String", "reverse-string", "Easy", "Reverse Words in a String II"},
	"indexOfMax":     {"Largest Number At Least Twice of Others", "largest-number-at-least-twice-of-others", "Easy", "Third Maximum Number"},
	"arraySum":       {"Running Sum of 1d Array", "running-sum-of-1d-array", "Easy", "Find Pivot Index"},
	"rotateRight":    {"Rotate Array", "rotate-array", "Medium", "Rotate List"},
	"runningSum":     {"Find Pivot Index", "find-pivot-index", "Easy", "Subarray Sum Equals K"},

	// 02 — hashing
	"twoSum":            {"Two Sum", "two-sum", "Easy", "3Sum"},
	"containsDuplicate": {"Contains Duplicate", "contains-duplicate", "Easy", "Contains Duplicate II"},
	"frequencyMap":      {"Majority Element", "majority-element", "Easy", "Top K Frequent Elements"},
	"firstUniqueChar":   {"First Unique Character in a String", "first-unique-character-in-a-string", "Easy", "Longest Substring Without Repeating Characters"},
	"groupAnagrams":     {"Group Anagrams", "group-anagrams", "Medium", "Longest Consecutive Sequence"},

	// 03 — two pointers & window
	"removeDuplicates": {"Remove Duplicates from Sorted Array", "remove-duplicates-from-sorted-array", "Easy", "Remove Duplicates from Sorted Array II"},
	"moveZeroes":       {"Move Zeroes", "move-zeroes", "Easy", "Sort Colors"},
	"maxArea":          {"Container With Most Water", "container-with-most-water", "Medium", "Trapping Rain Water"},
	"isPalindrome":     {"Valid Palindrome", "valid-palindrome", "Easy", "Valid Palindrome II"},
	"maxSumSubarrayK":  {"Maximum Average Subarray I", "maximum-average-subarray-i", "Easy", "Longest Substring Without Repeating Characters"},

	// 04 — binary search
	"binarySearch":    {"Binary Search", "binary-search", "Easy", "Search in Rotated Sorted Array"},
	"searchInsert":    {"Search Insert Position", "search-insert-position", "Easy", "Find First and Last Position of Element in Sorted Array"},
	"findMinRotated":  {"Find Minimum in Rotated Sorted Array", "find-minimum-in-rotated-sorted-array", "Medium", "Median of Two Sorted Arrays"},
	"isTargetPresent": {"Find Smallest Letter Greater Than Target", "find-smallest-letter-greater-than-target", "Easy", "Search a 2D Matrix"},

	// 05 — trees & stacks
	"inorderTraversal":    {"Binary Tree Inorder Traversal", "binary-tree-inorder-traversal", "Easy", "Kth Smallest Element in a BST"},
	"preorderTraversal":   {"Binary Tree Preorder Traversal", "binary-tree-preorder-traversal", "Easy", "Serialize and Deserialize Binary Tree"},
	"postorderTraversal":  {"Binary Tree Postorder Traversal", "binary-tree-postorder-traversal", "Easy", "Diameter of Binary Tree"},
	"levelOrderTraversal": {"Binary Tree Level Order Traversal", "binary-tree-level-order-traversal", "Medium", "Binary Tree Right Side View"},
	"maxDepth":            {"Maximum Depth of Binary Tree", "maximum-depth-of-binary-tree", "Easy", "Balanced Binary Tree"},
	"isValidParentheses":  {"Valid Parentheses", "valid-parentheses", "Easy", "Minimum Remove to Make Valid Parentheses"},
	"dailyTemperatures":   {"Daily Temperatures", "daily-temperatures", "Medium", "Largest Rectangle in Histogram"},

	// 06 — dp
	"fib":                   {"Fibonacci Number", "fibonacci-number", "Easy", "N-th Tribonacci Number"},
	"climbStairs":           {"Climbing Stairs", "climbing-stairs", "Easy", "Decode Ways"},
	"minCostClimbingStairs": {"Min Cost Climbing Stairs", "min-cost-climbing-stairs", "Easy", "Triangle"},
	"rob":                   {"House Robber", "house-robber", "Medium", "House Robber II"},

	// 07 — graphs
	"numIslands":       {"Number of Islands", "number-of-islands", "Medium", "Max Area of Island"},
	"floodFill":        {"Flood Fill", "flood-fill", "Easy", "Surrounded Regions"},
	"shortestPathGrid": {"Shortest Path in Binary Matrix", "shortest-path-in-binary-matrix", "Medium", "Rotting Oranges"},

	// 08 — heaps
	"kthLargest":      {"Kth Largest Element in an Array", "kth-largest-element-in-an-array", "Medium", "Kth Largest Element in a Stream"},
	"lastStoneWeight": {"Last Stone Weight", "last-stone-weight", "Easy", "Minimum Cost to Connect Sticks"},
	"mergeKSorted":    {"Merge k Sorted Lists", "merge-k-sorted-lists", "Hard", "Smallest Range Covering Elements from K Lists"},

	// 09 — backtracking
	"subsets": {"Subsets", "subsets", "Medium", "Subsets II"},
	"permute": {"Permutations", "permutations", "Medium", "Permutations II"},
	"combine": {"Combinations", "combinations", "Medium", "Combination Sum"},

	// 10 — math
	"gcd":      {"Greatest Common Divisor of Strings", "greatest-common-divisor-of-strings", "Easy", "Find Greatest Common Divisor of Array"},
	"lcm":      {"Range Addition II", "range-addition-ii", "Easy", "Nth Magical Number"},
	"modPow":   {"Pow(x, n)", "powx-n", "Medium", "Super Pow"},
	"nCr":      {"Pascal's Triangle", "pascals-triangle", "Easy", "Unique Paths"},
	"isPrime":  {"Count Primes", "count-primes", "Medium", "Prime Arrangements"},
	"powOfTwo": {"Power of Two", "power-of-two", "Easy", "Number of 1 Bits"},
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
