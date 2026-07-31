package main

import "fmt"

type askPrompt struct {
	statement string
	hints     []string
}

var dailyAsks = map[string]askPrompt{
	"Monday": {
		statement: "Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.",
		hints: []string{
			"Ask: Move last k elements to the front",
			"Input: array, k (k may be larger than n)",
			"Output: rotated array (in-place or new?)",
			"Edge: empty, k=0, k % n",
			"Pattern: reverse sections or modulo indexing",
		},
	},
	"Tuesday": {
		statement: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
		hints: []string{
			"Ask: Find two indices whose values sum to target",
			"Input: unsorted ints, one valid pair guaranteed",
			"Output: two indices (not values)",
			"Edge: negatives, duplicate values",
			"Pattern: hash map + complement",
		},
	},
	"Wednesday": {
		statement: "You are given an integer array height of length n. There are n vertical lines. Find two lines that together with the x-axis form a container that holds the most water.",
		hints: []string{
			"Ask: Maximize area = width × min(height[i], height[j])",
			"Input: heights array",
			"Output: max area (integer)",
			"Edge: two elements, all same height",
			"Pattern: two pointers L/R",
		},
	},
	"Thursday": {
		statement: "Given a sorted array of distinct integers and a target value, return the index if found. If not, return the index where it would be inserted.",
		hints: []string{
			"Ask: Lower bound / insertion position",
			"Input: sorted distinct array, target",
			"Output: index (0..n)",
			"Edge: insert before all, after all",
			"Pattern: binary search, first index ≥ target",
		},
	},
	"Friday": {
		statement: "Given a string s containing just '(', ')', '{', '}', '[' and ']', determine if the input string is valid.",
		hints: []string{
			"Ask: Every closing bracket matches most recent unmatched opener",
			"Input: bracket string",
			"Output: boolean",
			"Edge: empty, single opener, wrong type close",
			"Pattern: stack",
		},
	},
	"Saturday": {
		statement: "You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay cost[i], you can climb one or two steps. Return the minimum cost to reach the top.",
		hints: []string{
			"Ask: Min total cost to reach index n (past last step)",
			"Input: cost per step",
			"Output: min total cost",
			"Edge: n=1, n=2",
			"Pattern: 1D DP — dp[i] = cost[i] + min(dp[i-1], dp[i-2])",
		},
	},
	"Sunday": {
		statement: "Given an m×n 2D binary grid which represents a map of '1's (land) and '0's (water), return the number of islands.",
		hints: []string{
			"Ask: Count connected components of land",
			"Input: grid of '0' and '1'",
			"Output: island count",
			"Edge: all water, all land, single cell",
			"Pattern: DFS/BFS + visited",
		},
	},
}

func printAskWarmup(day string) {
	ask, ok := dailyAsks[day]
	if !ok {
		return
	}
	fmt.Println("\n── QUESTION LITERACY (before coding) ──────────────────")
	fmt.Printf("  Problem: %s\n\n", ask.statement)
	fmt.Println("  Fill in (on paper or aloud):")
	for _, h := range ask.hints {
		fmt.Printf("    • %s\n", h)
	}
	fmt.Println("\n  Full asks pack: study_play/_support/asks/README.md")
}
