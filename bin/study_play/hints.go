package main

import (
	"fmt"
	"strings"
)

// fnHints is the technique for each drill function. The stub comments in
// drills/write/reflex only say what to compute, so choosing the technique is
// part of the drill; the hint appears with --hint=NAME or after a first-try
// miss.
var fnHints = map[string]string{
	"reverseInPlace":    "two pointers from both ends; swap and step inward until they meet",
	"rotateRight":       "k %= n; reverse all, then reverse the first k and the rest",
	"runningSum":        "out[i] = out[i-1] + arr[i]",
	"subarraySumK":      "prefix sum + map of prefix counts seeded {0:1}; add count[sum-k] before storing sum",
	"productExceptSelf": "left-to-right prefix products into out, then multiply by a right-to-left suffix product",
	"maxSubarraySum":    "Kadane: cur = max(x, cur+x); best = max(best, cur)",
	"mergeSort":         "split in half, sort each side, merge two sorted halves taking <= from the left (stable)",
	"quickSort":         "partition around a pivot (Lomuto/Hoare), recurse on each side",
	"sieve":             "sieve of Eratosthenes: for p*p <= n, cross out multiples starting at p*p",

	"twoSum":            "map value -> index; look up target-x before storing x",
	"containsDuplicate": "set of seen values; a hit means a duplicate",
	"frequencyMap":      "map value -> count, one pass",
	"firstUniqueChar":   "count every char, then a second pass for the first with count 1",
	"groupAnagrams":     "key = sorted string (or 26-count signature); map key -> group",
	"singleNumber":      "XOR everything: pairs cancel (a^a = 0), the single value is left",

	"removeDuplicates":       "write pointer w; copy nums[r] to nums[w] when it differs from nums[w-1]",
	"moveZeroes":             "write pointer for the next non-zero; then fill the tail with zeroes (or swap)",
	"maxArea":                "two pointers at both ends; record area, move the shorter wall inward",
	"isPalindrome":           "two pointers from both ends, skipping non-alphanumerics, compare lowercased",
	"maxSumSubarrayK":        "fixed window: sum the first k, then add nums[i] and drop nums[i-k]",
	"longestUniqueSubstring": "variable window: last-seen map; jump left past a repeat, best = max(right-left+1)",
	"dutchFlag":              "Dutch flag: low/mid/high; 0 swaps to low (both advance), 2 swaps to high (mid stays)",

	"binarySearch":   "lo<=hi; mid = lo+(hi-lo)/2; move past mid on each side",
	"searchInsert":   "lower bound: lo<hi over [0,n); nums[mid] < target -> lo=mid+1 else hi=mid",
	"findMinRotated": "compare nums[mid] with nums[hi]: bigger -> min is right of mid, else hi=mid",
	"minEatingSpeed": "binary search the speed in [1, max pile]; feasible(s) = sum(ceil(p/s)) <= h",
	"quickSelect":    "quickselect: partition, then recurse only into the side holding index k-1",
	"fastPow":        "binary exponentiation: square x and halve n; multiply in when n is odd; n<0 -> 1/x",
	"gcd":            "Euclid: a, b = b, a%b until b == 0",

	"inorderTraversal":    "stack: go left pushing every node; pop and visit; then go to the right child",
	"preorderTraversal":   "stack: pop and visit; push right, then left (so left pops first)",
	"postorderTraversal":  "preorder with left pushed first (root, right, left), then reverse the output",
	"levelOrderTraversal": "queue: pop the front, visit, push left then right",
	"maxDepth":            "nil -> 0; otherwise 1 + max(depth(left), depth(right))",
	"isValidBST":          "recurse with (lo, hi) bounds: left narrows hi to the node, right raises lo",
	"isValidParentheses":  "stack of openers; a closer must match the top; the stack must end empty",
	"dailyTemperatures":   "monotonic stack of indices waiting for a warmer day; a warmer day pops and answers them",

	"climbStairs":           "dp[i] = dp[i-1] + dp[i-2]; dp[0] = dp[1] = 1",
	"minCostClimbingStairs": "dp[i] = cost[i] + min(dp[i-1], dp[i-2]); answer min(dp[n-1], dp[n-2])",
	"rob":                   "dp[i] = max(dp[i-1], dp[i-2] + nums[i]) — skip this house or rob it",
	"coinChange":            "dp[a] = min over coins of dp[a-c] + 1; dp[0] = 0, others start at infinity",

	"numIslands":       "scan every cell; each unvisited '1' starts a DFS/BFS that sinks its island; count++",
	"floodFill":        "DFS/BFS from (sr, sc) over 4-neighbours with the start colour; return early if colour is unchanged",
	"shortestPathGrid": "BFS from the start, marking visited on enqueue; the first time you reach the end is shortest",
	"canFinish":        "Kahn: indegree counts, queue the zero-indegree courses; finished == n means no cycle",
	"dfs":              "recursive: visit, mark, then recurse into each unvisited neighbour in order",
	"bfs":              "queue + visited marked on enqueue; pop, visit, enqueue unvisited neighbours",
	"bfsShortestPath":  "BFS with dist[]; dist[v] = dist[u] + 1 when v is first enqueued",
	"topoSort":         "Kahn with a min-heap or sorted queue for ties; fewer than n out -> cycle -> nil",
	"dijkstra":         "min-heap of (dist, node); skip stale pops; relax dist[v] = dist[u] + w",
}

// printHint shows a function's pattern, technique and practice problem.
func printHint(fn string) error {
	file, ok := drillFileOf(fn)
	if !ok {
		return fmt.Errorf("unknown function %q%s", fn, suggestFn(fn))
	}
	fmt.Printf("%s  (%s)\n", fn, file)
	if p, ok := patternByID(fnPattern[fn]); ok {
		fmt.Printf("  pattern:   %s — %s\n", p.id, p.name)
		fmt.Printf("  signal:    %s\n", p.signal)
	}
	if h, ok := fnHints[fn]; ok {
		fmt.Printf("  technique: %s\n", h)
	}
	if p, ok := primaries[fn]; ok {
		fmt.Printf("  practice:  %s (%s) %s\n", p.title, p.diff, problemURL(p.slug))
	}
	return nil
}

// printHints prints printHint for each name in a comma list.
func printHints(list string) error {
	for i, fn := range strings.Split(list, ",") {
		if i > 0 {
			fmt.Println()
		}
		if err := printHint(strings.TrimSpace(fn)); err != nil {
			return err
		}
	}
	return nil
}
