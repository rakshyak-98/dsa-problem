// Daily LeetCode practice sets — 10 full problems per weekday topic.
//
// RUN:              go run .
// Full catalog:     go run . -- --catalog
// Today's set only: go run . -- --set
package main

import (
	"fmt"
	"time"
)

type lcProblem struct {
	num      int
	title    string
	diff     string
	pattern  string
	reflexFn string // links to matching reflex drill function, if any
}

type practiceSet struct {
	day       string
	topic     string
	reflex    string // matching write reflex drill file
	problems  []lcProblem
	warmup    string
	suggested []string // study tips
}

func lcURL(num int, slug string) string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", slug)
}

var practiceSets = []practiceSet{
	{
		day: "Monday", topic: "Arrays", reflex: "01_arrays_reflex",
		warmup: "Before coding: name the array shape (sorted? in-place?) and whether you need a prefix or two indices.",
		suggested: []string{
			"Start with Two Sum and Best Time to Buy — map + single pass patterns",
			"Save in-place problems (Rotate, Move Zeroes) for when you're warmed up",
		},
		problems: []lcProblem{
			{1, "Two Sum", "Easy", "hash map + complement", "twoSum"},
			{121, "Best Time to Buy and Sell Stock", "Easy", "single pass min + max profit", ""},
			{217, "Contains Duplicate", "Easy", "set / sort", "containsDuplicate"},
			{53, "Maximum Subarray", "Easy", "Kadane / running max", ""},
			{238, "Product of Array Except Self", "Medium", "prefix + suffix without division", ""},
			{88, "Merge Sorted Array", "Easy", "write from end two pointers", ""},
			{189, "Rotate Array", "Easy", "reverse sections or modulo copy", "rotateRight"},
			{283, "Move Zeroes", "Easy", "read/write two pointers", "moveZeroes"},
			{66, "Plus One", "Easy", "carry from end", ""},
			{169, "Majority Element", "Easy", "Boyer-Moore or hash count", ""},
		},
	},
	{
		day: "Tuesday", topic: "Hashing", reflex: "02_hashing_reflex",
		warmup: "Ask: do I need counts, seen-before, or a canonical key for grouping?",
		suggested: []string{
			"Group Anagrams and Top K Frequent are interview staples",
			"Longest Consecutive is a great 'sort vs hash' tradeoff question",
		},
		problems: []lcProblem{
			{242, "Valid Anagram", "Easy", "freq map or sort", ""},
			{49, "Group Anagrams", "Medium", "sorted key or count signature", "groupAnagrams"},
			{347, "Top K Frequent Elements", "Medium", "freq map + bucket/heap", ""},
			{128, "Longest Consecutive Sequence", "Medium", "set + expand from each start", ""},
			{1, "Two Sum", "Easy", "map + complement", "twoSum"},
			{217, "Contains Duplicate", "Easy", "set membership", "containsDuplicate"},
			{387, "First Unique Character in a String", "Easy", "freq map + second pass", "firstUniqueChar"},
			{205, "Isomorphic Strings", "Easy", "two maps or normalized key", ""},
			{202, "Happy Number", "Easy", "set detect cycle", ""},
			{350, "Intersection of Two Arrays II", "Easy", "freq map on smaller array", ""},
		},
	},
	{
		day: "Wednesday", topic: "Two Pointers", reflex: "03_two_pointers_reflex",
		warmup: "Sorted input → L/R. In-place filter → read/write. Subarray length → sliding window.",
		suggested: []string{
			"Container and 3Sum reward clean pointer discipline",
			"Longest Substring Without Repeating is the classic variable window",
		},
		problems: []lcProblem{
			{125, "Valid Palindrome", "Easy", "L/R skip non-alnum", "isPalindrome"},
			{167, "Two Sum II", "Easy", "sorted L/R two pointers", ""},
			{15, "3Sum", "Medium", "sort + fix i + L/R", ""},
			{11, "Container With Most Water", "Medium", "L/R move shorter side", "maxArea"},
			{42, "Trapping Rain Water", "Hard", "L/R max height", ""},
			{26, "Remove Duplicates from Sorted Array", "Easy", "read/write dedupe", "removeDuplicates"},
			{283, "Move Zeroes", "Easy", "read/write filter", "moveZeroes"},
			{977, "Squares of a Sorted Array", "Easy", "merge from ends", ""},
			{3, "Longest Substring Without Repeating Characters", "Medium", "variable sliding window", ""},
			{209, "Minimum Size Subarray Sum", "Medium", "variable window on sum", ""},
		},
	},
	{
		day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		warmup: "Exact find → lo <= hi. First >= target → lower bound. Answer space → BS on answer.",
		suggested: []string{
			"Search Insert and First Bad Version are warm-up templates",
			"Koko Eating Bananas introduces binary search on the answer",
		},
		problems: []lcProblem{
			{704, "Binary Search", "Easy", "classic lo <= hi", "binarySearch"},
			{35, "Search Insert Position", "Easy", "lower bound", "searchInsert"},
			{153, "Find Minimum in Rotated Sorted Array", "Medium", "rotated half sorted", "findMinRotated"},
			{33, "Search in Rotated Sorted Array", "Medium", "pick sorted half", ""},
			{278, "First Bad Version", "Easy", "lower bound on predicate", ""},
			{69, "Sqrt(x)", "Easy", "BS on answer range", ""},
			{367, "Valid Perfect Square", "Easy", "BS on sqrt candidate", ""},
			{744, "Find Smallest Letter Greater Than Target", "Easy", "lower bound on letters", ""},
			{852, "Peak Index in Mountain Array", "Medium", "BS on mountain slope", ""},
			{875, "Koko Eating Bananas", "Medium", "BS on eating speed", ""},
		},
	},
	{
		day: "Friday", topic: "Trees & Stacks", reflex: "05_trees_stacks_reflex",
		warmup: "Tree traversal → stack/recursion. Matching brackets → stack. Next greater → monotonic stack.",
		suggested: []string{
			"Valid Parentheses and Inorder Traversal mirror today's reflex drill",
			"Daily Temperatures is the monotonic stack template",
		},
		problems: []lcProblem{
			{20, "Valid Parentheses", "Easy", "stack of openers", "isValidParentheses"},
			{94, "Binary Tree Inorder Traversal", "Easy", "stack iterative inorder", "inorderTraversal"},
			{104, "Maximum Depth of Binary Tree", "Easy", "DFS depth", "maxDepth"},
			{226, "Invert Binary Tree", "Easy", "recursive swap children", ""},
			{100, "Same Tree", "Easy", "parallel DFS compare", ""},
			{101, "Symmetric Tree", "Easy", "mirror DFS", ""},
			{112, "Path Sum", "Easy", "DFS accumulate target", ""},
			{739, "Daily Temperatures", "Medium", "monotonic decreasing stack", "dailyTemperatures"},
			{155, "Min Stack", "Medium", "stack + min tracking", ""},
			{102, "Binary Tree Level Order Traversal", "Medium", "BFS queue", ""},
		},
	},
	{
		day: "Saturday", topic: "Dynamic Programming", reflex: "06_dp_reflex",
		warmup: "Define dp[i] in English first. Linear DP → look back 1–2 states. Grid → 2D table.",
		suggested: []string{
			"Climbing Stairs and House Robber are the 1D DP foundations",
			"Coin Change and Word Break test state definition under pressure",
		},
		problems: []lcProblem{
			{70, "Climbing Stairs", "Easy", "1D ways dp", "climbStairs"},
			{198, "House Robber", "Medium", "take/skip adjacent", "rob"},
			{746, "Min Cost Climbing Stairs", "Easy", "min of last two steps", "minCostClimbingStairs"},
			{322, "Coin Change", "Medium", "unbounded knapsack min coins", ""},
			{300, "Longest Increasing Subsequence", "Medium", "dp or patience sorting", ""},
			{139, "Word Break", "Medium", "dp on prefix reachable", ""},
			{91, "Decode Ways", "Medium", "string dp 1–2 char lookback", ""},
			{62, "Unique Paths", "Medium", "grid dp", ""},
			{152, "Maximum Product Subarray", "Medium", "track min and max product", ""},
			{416, "Partition Equal Subset Sum", "Medium", "0/1 knapsack subset", ""},
		},
	},
	{
		day: "Sunday", topic: "Graphs", reflex: "07_graphs_reflex",
		warmup: "Grid component → DFS/BFS + visited. Shortest unweighted path → BFS. Dependencies → topo sort.",
		suggested: []string{
			"Number of Islands and Flood Fill match today's reflex patterns",
			"Course Schedule is the canonical cycle / topo sort question",
		},
		problems: []lcProblem{
			{200, "Number of Islands", "Medium", "DFS/BFS grid components", "numIslands"},
			{733, "Flood Fill", "Easy", "DFS recolor connected", "floodFill"},
			{133, "Clone Graph", "Medium", "BFS/DFS + hash clone", ""},
			{207, "Course Schedule", "Medium", "cycle detect / topo sort", ""},
			{994, "Rotting Oranges", "Medium", "multi-source BFS", ""},
			{127, "Word Ladder", "Hard", "BFS shortest transform", ""},
			{130, "Surrounded Regions", "Medium", "DFS from border", ""},
			{417, "Pacific Atlantic Water Flow", "Medium", "reverse flow DFS from oceans", ""},
			{743, "Network Delay Time", "Medium", "Dijkstra / BFS weighted", ""},
			{1971, "Find if Path Exists in Graph", "Easy", "DFS/BFS reachability", "shortestPathGrid"},
		},
	},
}

func todaySet() practiceSet {
	wd := int(time.Now().Weekday())
	idx := (wd + 6) % 7
	return practiceSets[idx]
}

func slugFor(num int) string {
	slugs := map[int]string{
		1: "two-sum", 3: "longest-substring-without-repeating-characters",
		11: "container-with-most-water", 15: "3sum", 20: "valid-parentheses",
		26: "remove-duplicates-from-sorted-array", 33: "search-in-rotated-sorted-array",
		35: "search-insert-position", 42: "trapping-rain-water", 49: "group-anagrams",
		53: "maximum-subarray", 62: "unique-paths", 66: "plus-one", 69: "sqrtx",
		70: "climbing-stairs", 88: "merge-sorted-array", 91: "decode-ways",
		94: "binary-tree-inorder-traversal", 100: "same-tree", 101: "symmetric-tree",
		102: "binary-tree-level-order-traversal", 104: "maximum-depth-of-binary-tree",
		112: "path-sum", 121: "best-time-to-buy-and-sell-stock", 125: "valid-palindrome",
		127: "word-ladder", 128: "longest-consecutive-sequence", 130: "surrounded-regions",
		133: "clone-graph", 139: "word-break", 152: "maximum-product-subarray",
		153: "find-minimum-in-rotated-sorted-array", 155: "min-stack",
		167: "two-sum-ii-input-array-is-sorted", 169: "majority-element",
		189: "rotate-array", 197: "rising-temperature", 198: "house-robber",
		200: "number-of-islands", 202: "happy-number", 205: "isomorphic-strings",
		207: "course-schedule", 209: "minimum-size-subarray-sum", 217: "contains-duplicate",
		226: "invert-binary-tree", 238: "product-of-array-except-self", 242: "valid-anagram",
		278: "first-bad-version", 283: "move-zeroes", 300: "longest-increasing-subsequence",
		322: "coin-change", 347: "top-k-frequent-elements", 350: "intersection-of-two-arrays-ii",
		367: "valid-perfect-square", 387: "first-unique-character-in-a-string",
		416: "partition-equal-subset-sum", 417: "pacific-atlantic-water-flow",
		704: "binary-search", 733: "flood-fill", 739: "daily-temperatures",
		743: "network-delay-time", 744: "find-smallest-letter-greater-than-target",
		746: "min-cost-climbing-stairs", 852: "peak-index-in-a-mountain-array",
		875: "koko-eating-bananas", 977: "squares-of-a-sorted-array",
		994: "rotting-oranges", 1971: "find-if-path-exists-in-graph",
	}
	if s, ok := slugs[num]; ok {
		return s
	}
	return fmt.Sprintf("problem-%d", num)
}
