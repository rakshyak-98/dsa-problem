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
		{"rotateRight", "arrays/easy/shuffle_the_array.js", "Interleave two halves of an array"},
		{"runningSum", "misc/easy/running_sum_of_1d_array.js", "Running total while scanning"},
		{"subarraySumK", "misc/medium/subarray_sum_divisible_by_k.js", "Count subarrays by prefix remainder"},
		{"productExceptSelf", "misc/medium/product_of_array_except_self.js", "Product of all other elements"},
		{"maxSubarraySum", "arrays/medium/max_product_subarray.js", "Best contiguous run, product variant"},
		{"mergeSort", "sorting/medium/sort_an_array.js", "Stable O(n log n) sort by split and merge"},
		{"quickSort", "sorting/medium/sort_an_array.js", "In-place O(n log n) sort by partitioning"},
		{"sieve", "math/medium/count_primes.js", "All primes up to n via sieve of Eratosthenes"},
	},
	"02_hashing_reflex": {
		{"twoSum", "hashing/easy/two_sum.js", "Return indices of two numbers that add to target"},
		{"containsDuplicate", "hashing/easy/contains_duplicates.js", "Return true if any value appears twice"},
		{"frequencyMap", "hashing/easy/top_k_ferquent_element.js", "Count frequency of each element"},
		{"firstUniqueChar", "hashing/easy/first_unique_character_in_a_string.js", "First non-repeating character"},
		{"groupAnagrams", "hashing/medium/group_anagram.js", "Group strings that are anagrams"},
		{"singleNumber", "bit_manipulation/easy/single_number.js", "The one unpaired value via XOR"},
	},
	"03_two_pointers_reflex": {
		{"removeDuplicates", "two_pointers/easy/remove_duplicates_from_sorted_array.js", "In-place dedupe sorted array, return new length"},
		{"moveZeroes", "two_pointers/easy/move_zeroes.js", "Move all zeroes to end in-place"},
		{"maxArea", "two_pointers/medium/container_with_most_water.js", "Max water between two vertical lines"},
		{"isPalindrome", "two_pointers/easy/valid_palindrome.js", "Is string a palindrome ignoring non-alphanumeric"},
		{"maxSumSubarrayK", "misc/easy/maximum_average_subarray_1.js", "Max sum of subarray of size k"},
		{"longestUniqueSubstring", "strings/medium/longest_substring_without_repeating.js", "Longest run with no repeat"},
		{"dutchFlag", "sorting/medium/sort_colors.js", "One-pass three-way partition of 0/1/2"},
	},
	"04_binary_search_reflex": {
		{"binarySearch", "binary_search/easy/search_insertion_position.js", "Find target index in sorted array"},
		{"searchInsert", "binary_search/easy/search_insertion_position.js", "Insertion index for target in sorted array"},
		{"findMinRotated", "binary_search/medium/find_minimum_in_rotated_sorted_array.js", "Minimum in rotated sorted array"},
		{"minEatingSpeed", "binary_search/medium/koko_eating_bananas.js", "Slowest rate that still finishes in time"},
		{"quickSelect", "sorting/medium/kth_largest_element_in_an_array.js", "Kth smallest without fully sorting"},
		{"fastPow", "math/medium/powx_n.js", "x^n in O(log n) by binary exponentiation"},
		{"gcd", "math/easy/find_greatest_common_divisor_of_array.js", "Greatest common divisor via Euclid"},
	},
	"05_trees_stacks_reflex": {
		{"inorderTraversal", "trees/easy/binary_tree_inorder_traversal.js", "Return inorder traversal of binary tree"},
		{"preorderTraversal", "trees/easy/binary_tree_preorder_traversal.js", "Return preorder traversal of binary tree"},
		{"postorderTraversal", "trees/easy/binary_tree_postorder_traversal.js", "Return postorder traversal of binary tree"},
		{"levelOrderTraversal", "trees/medium/binary_tree_level_order_traversal.js", "Return level-order traversal of binary tree"},
		{"maxDepth", "trees/easy/maximum_depth_of_binary_tree.js", "Maximum depth of binary tree"},
		{"isValidBST", "trees/medium/validate_binary_search_tree.js", "Is every node within its ancestor bounds"},
		{"isValidParentheses", "stacks/easy/valid_parentheses.js", "Are brackets properly matched"},
		{"dailyTemperatures", "stacks/medium/daily_temperatures.js", "Days until warmer temperature"},
	},
	"06_dp_reflex": {
		{"minCostClimbingStairs", "dynamic_programming/easy/min_cost_climbing_staris.js", "Min cost to reach top of stairs"},
		{"rob", "dynamic_programming/medium/house_robber.js", "Max money robbing non-adjacent houses"},
		{"climbStairs", "dynamic_programming/easy/climbing_stairs.js", "Count ways to climb n stairs"},
		{"coinChange", "dynamic_programming/medium/coin_change.js", "Fewest coins making an exact total"},
	},
	"07_graphs_reflex": {
		{"numIslands", "graphs/medium/number_of_islands.js", "Count connected land regions in grid"},
		{"floodFill", "graphs/easy/flood_fill.js", "Recolor connected pixels from start"},
		{"shortestPathGrid", "graphs/medium/shortest_path_in_binary_matrix.js", "Shortest path in unweighted grid"},
		{"canFinish", "graphs/medium/course_schedule.js", "Ordering with prerequisites / cycle check"},
		{"dfs", "graphs/medium/number_of_provinces.js", "Reachable nodes / components via depth-first search"},
		{"bfs", "graphs/medium/snakes_and_ladders.js", "Reachable nodes level by level via breadth-first search"},
		{"bfsShortestPath", "graphs/medium/minimum_genetic_mutation.js", "Fewest edges between two nodes via BFS"},
		{"topoSort", "graphs/medium/course_schedule_ii.js", "Linear order of a DAG via Kahn's algorithm"},
		{"dijkstra", "graphs/medium/network_delay_time.js", "Weighted shortest paths from a source"},
	},
	"08_heap_reflex": {
		{"kthLargest", "heaps/medium/kth_largest_element_in_an_array.js", "Kth largest element in array"},
		{"lastStoneWeight", "heaps/easy/last_stone_weight.js", "Simulate stone smashing with max heap"},
		{"mergeKSorted", "heaps/hard/merge_k_sorted_lists.js", "Merge k sorted linked lists"},
	},
	"09_backtrack_reflex": {
		{"subsets", "backtracking/medium/subsets.js", "Return all subsets of nums"},
		{"permute", "backtracking/medium/permutations.js", "Return all permutations of nums"},
		{"combine", "backtracking/medium/combinations.js", "All combinations of k numbers from 1..n"},
	},
	"10_linked_list_reflex": {
		{"reverseList", "linked_list/easy/reverse_linked_list.js", "Reverse the list in place"},
		{"hasCycle", "linked_list/easy/linked_list_cycle.js", "Detect a loop with slow + fast"},
		{"middleNode", "linked_list/easy/middle_of_the_linked_list.js", "Middle node in one pass"},
		{"mergeTwoLists", "linked_list/easy/merge_two_sorted_lists.js", "Merge two sorted lists"},
		{"removeNthFromEnd", "linked_list/medium/remove_nth_node_from_end.js", "Drop the nth node from the end"},
	},
}

var core5Problems = []problemLink{
	{"twoSum", "hashing/easy/two_sum.js", "Return indices of two numbers that add to target"},
	{"binarySearch", "binary_search/easy/search_insertion_position.js", "Find target index in sorted array"},
	{"removeDuplicates", "two_pointers/easy/remove_duplicates_from_sorted_array.js", "In-place dedupe sorted array"},
	{"maxSumSubarrayK", "sliding_window/easy/maximum_average_subarray_1.js", "Max sum of subarray of size k"},
	{"frequencyMap", "hashing/easy/top_k_ferquent_element.js", "Count frequency of each element"},
}

// printProblemMap shows what to solve once the drill's tests pass. The
// headline is the curated LeetCode primary from primaries.go — the local
// reference/problems/ mirror is only a partial index, so it is shown as a
// secondary hint and never as the target.
func printProblemMap(drillFile string) {
	links, ok := problemMap[drillFile]
	if !ok {
		return
	}
	fmt.Println("\n── AFTER TESTS PASS: SOLVE THESE ──────────────────────")
	for _, l := range links {
		p, hasPrimary := primaries[l.function]
		if !hasPrimary {
			fmt.Printf("  • %s → %s\n", l.function, l.problem)
			continue
		}
		fmt.Printf("  • %-20s %-6s %s\n", l.function, p.diff, p.title)
		if c, ok := cueByFn[l.function]; ok {
			fmt.Printf("    ask: %s\n", c.ask)
		}
		fmt.Printf("    %s\n", problemURL(p.slug))
	}
	fmt.Println("\n  Solve without reopening the drill. Stuck twice → that is an L1 function again.")
}

func printCore5Problems() {
	fmt.Println("\n── CORE 5 → PRIMARY PROBLEMS ──────────────────────────")
	for _, l := range core5Problems {
		p, ok := primaries[l.function]
		if !ok {
			fmt.Printf("  • %s → %s\n", l.function, l.problem)
			continue
		}
		fmt.Printf("  • %-20s %-6s %-40s %s\n", l.function, p.diff, p.title, problemURL(p.slug))
	}
}
