package main

import "fmt"

type problemLink struct {
	function string
	problem  string
	ask      string
}

var problemMap = map[string][]problemLink{
	"01_arrays_reflex": {
		{"reverseInPlace", "arrays/easy/plus_one.js", "Increment array as if it were a number"},
		{"indexOfMax", "arrays/easy/largest_number_at_least_twice_of_others.js", "Find dominant element or return -1"},
		{"arraySum", "arrays/easy/concatenation_of_array.js", "Build concatenated array from two halves"},
		{"rotateRight", "arrays/easy/shuffle_the_array.js", "Interleave two halves of an array"},
		{"runningSum", "arrays/easy/find_closest_number_to_zero.js", "Track running best while scanning"},
	},
	"02_hashing_reflex": {
		{"twoSum", "hashing/easy/two_sum.js", "Return indices of two numbers that add to target"},
		{"containsDuplicate", "hashing/easy/contains_duplicates.js", "Return true if any value appears twice"},
		{"frequencyMap", "hashing/easy/top_k_ferquent_element.js", "Count frequency of each element"},
		{"firstUniqueChar", "hashing/easy/first_unique_character_in_a_string.js", "First non-repeating character"},
		{"groupAnagrams", "hashing/medium/group_anagram.js", "Group strings that are anagrams"},
	},
	"03_two_pointers_reflex": {
		{"removeDuplicates", "two_pointers/easy/remove_duplicates_from_sorted_array.js", "In-place dedupe sorted array, return new length"},
		{"moveZeroes", "two_pointers/easy/move_zeroes.js", "Move all zeroes to end in-place"},
		{"maxArea", "two_pointers/medium/container_with_most_water.js", "Max water between two vertical lines"},
		{"isPalindrome", "two_pointers/easy/valid_palindrome.js", "Is string a palindrome ignoring non-alphanumeric"},
		{"maxSumSubarrayK", "sliding_window/easy/maximum_average_subarray_1.js", "Max sum of subarray of size k"},
	},
	"04_binary_search_reflex": {
		{"binarySearch", "binary_search/easy/search_insertion_position.js", "Find target index in sorted array"},
		{"searchInsert", "binary_search/easy/search_insertion_position.js", "Insertion index for target in sorted array"},
		{"findMinRotated", "binary_search/medium/find_minimum_in_rotated_sorted_array.js", "Minimum in rotated sorted array"},
		{"isTargetPresent", "binary_search/easy/find_smallest_letter_greater_than_target.js", "Is target present in sorted data"},
	},
	"05_trees_stacks_reflex": {
		{"inorderTraversal", "trees/easy/binary_tree_inorder_traversal.js", "Return inorder traversal of binary tree"},
		{"maxDepth", "trees/easy/maximum_depth_of_binary_tree.js", "Maximum depth of binary tree"},
		{"isValidParentheses", "stacks/easy/valid_parentheses.js", "Are brackets properly matched"},
		{"dailyTemperatures", "stacks/medium/daily_temperatures.js", "Days until warmer temperature"},
	},
	"06_dp_reflex": {
		{"fib", "dynamic_programming/easy/fibonacci_number.js", "Nth Fibonacci number"},
		{"minCostClimbingStairs", "dynamic_programming/easy/min_cost_climbing_staris.js", "Min cost to reach top of stairs"},
		{"rob", "dynamic_programming/medium/house_robber.js", "Max money robbing non-adjacent houses"},
		{"climbStairs", "dynamic_programming/easy/climbing_stairs.js", "Count ways to climb n stairs"},
	},
	"07_graphs_reflex": {
		{"numIslands", "graphs/medium/number_of_islands.js", "Count connected land regions in grid"},
		{"floodFill", "graphs/easy/flood_fill.js", "Recolor connected pixels from start"},
		{"shortestPathGrid", "graphs/medium/shortest_path_in_binary_matrix.js", "Shortest path in unweighted grid"},
	},
	"08_heap_reflex": {
		{"kthLargest", "heaps/medium/kth_largest_element_in_an_array.js", "Kth largest element in array"},
		{"lastStoneWeight", "heaps/easy/last_stone_weight.js", "Simulate stone smashing with max heap"},
		{"mergeKSortedLists", "heaps/hard/merge_k_sorted_lists.js", "Merge k sorted linked lists"},
	},
	"09_backtrack_reflex": {
		{"subsets", "backtracking/medium/subsets.js", "Return all subsets of nums"},
		{"permute", "backtracking/medium/permutations.js", "Return all permutations of nums"},
		{"combine", "backtracking/medium/combinations.js", "All combinations of k numbers from 1..n"},
	},
}

var core5Problems = []problemLink{
	{"twoSum", "hashing/easy/two_sum.js", "Return indices of two numbers that add to target"},
	{"binarySearch", "binary_search/easy/search_insertion_position.js", "Find target index in sorted array"},
	{"removeDuplicates", "two_pointers/easy/remove_duplicates_from_sorted_array.js", "In-place dedupe sorted array"},
	{"maxSumSubarrayK", "sliding_window/easy/maximum_average_subarray_1.js", "Max sum of subarray of size k"},
	{"frequencyMap", "hashing/easy/top_k_ferquent_element.js", "Count frequency of each element"},
}

func printProblemMap(drillFile string) {
	links, ok := problemMap[drillFile]
	if !ok {
		return
	}
	fmt.Println("\n── DRILL → PRIMARY PROBLEMS ───────────────────────────")
	for _, l := range links {
		fmt.Printf("  • %s → %s\n", l.function, l.problem)
		fmt.Printf("    Ask: %s\n", l.ask)
	}
	fmt.Println("\n  After tests pass: solve the primary without peeking at this drill.")
}

func printCore5Problems() {
	fmt.Println("\n── CORE 5 → PRIMARY PROBLEMS ──────────────────────────")
	for _, l := range core5Problems {
		fmt.Printf("  • %s → %s\n", l.function, l.problem)
	}
}
