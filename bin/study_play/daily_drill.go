// Daily reflex practice helper — Core 5 + specialty drill
//
// RUN:              go run .
// RUN with tests:   go run . -- --run
// Core 5 only:      go run . -- --drill
// Full catalog:     go run . -- --catalog
// Reset today:      go run . -- --reset
// First-time setup: go run . -- --setup
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type coreFn struct {
	name    string
	ask     string
	pattern string
	sec     int
}

type drill struct {
	day              string
	file             string
	patterns         string
	functions        []string
	triggers         []string
	understandWarmup string
}

var core5 = []coreFn{
	{"twoSum(nums, target)", "indices of two values that sum to target", "map + complement", 90},
	{"binarySearch(nums, target)", "index of target in sorted array, or -1", "lo <= hi binary search", 60},
	{"removeDuplicates(nums)", "in-place unique prefix length on sorted array", "read/write two pointers", 90},
	{"maxSumSubarrayK(nums, k)", "max sum of any contiguous window of size k", "fixed sliding window", 90},
	{"frequencyMap(arr)", "map each value to its count", "freq map loop", 60},
}

var drills = []drill{
	{
		day: "Monday", file: "01_arrays_reflex",
		patterns:  "reverse, max index, sum, rotate, prefix",
		functions: []string{"reverseInPlace", "indexOfMax", "arraySum", "rotateRight", "runningSum"},
		triggers: []string{
			"in-place mutate / reverse → two pointers L/R swap",
			"running total → prefix / accumulate",
			"rotate by k → k %= n, then reverse sections (or copy with modulo)",
		},
		understandWarmup: "Rotate right by k: same-length array; last k elements move to the front.",
	},
	{
		day: "Tuesday", file: "02_hashing_reflex",
		patterns:  "two-sum, dup, freq map, unique char, anagrams",
		functions: []string{"twoSum", "containsDuplicate", "frequencyMap", "firstUniqueChar", "groupAnagrams"},
		triggers: []string{
			"pair sums to target → map + complement",
			"seen before? → set / map",
			"same letter multiset → sorted key or count signature",
		},
		understandWarmup: "Two-sum: return indices of two numbers that add to target (usually one valid pair).",
	},
	{
		day: "Wednesday", file: "03_two_pointers_reflex",
		patterns:  "dedupe, zeroes, container, palindrome, window",
		functions: []string{"removeDuplicates", "moveZeroes", "maxArea", "isPalindrome", "maxSumSubarrayK"},
		triggers: []string{
			"sorted + two values → L/R two pointers",
			"in-place filter / dedupe → read/write pointers",
			"subarray length k → fixed sliding window",
		},
		understandWarmup: "Container with most water: max area between two lines = width × min(height).",
	},
	{
		day: "Thursday", file: "04_binary_search_reflex",
		patterns:  "exact BS, lower bound, rotated min, present?",
		functions: []string{"binarySearch", "searchInsert", "findMinRotated", "isTargetPresent"},
		triggers: []string{
			"sorted + find exact → lo <= hi, mid compare",
			"first position ≥ target → lower bound / searchInsert",
			"rotated sorted min → decide which half is sorted",
		},
		understandWarmup: "Lower bound: first index where value is ≥ target (insertion point).",
	},
	{
		day: "Friday", file: "05_trees_stacks_reflex",
		patterns:  "inorder, depth, parens, mono stack",
		functions: []string{"inorderTraversal", "maxDepth", "isValidParentheses", "dailyTemperatures"},
		triggers: []string{
			"tree order without recursion → stack / iterative inorder",
			"matching brackets → stack of opens",
			"next greater element → monotonic decreasing stack",
		},
		understandWarmup: "Valid parentheses: every closer must match the latest unmatched opener.",
	},
	{
		day: "Saturday", file: "06_dp_reflex",
		patterns:  "fib, min cost, rob, climb stairs",
		functions: []string{"fib", "minCostClimbingStairs", "rob", "climbStairs"},
		triggers: []string{
			"min cost / max ways on a line → 1D DP",
			"define dp[i] in English before coding",
			"rob houses → cannot take adjacent → max(take, skip)",
		},
		understandWarmup: "Min cost climbing: from i you pay cost[i], then jump 1 or 2 steps; reach top with min total.",
	},
	{
		day: "Sunday", file: "07_graphs_reflex",
		patterns:  "islands, flood fill, BFS path",
		functions: []string{"numIslands", "floodFill", "shortestPathGrid"},
		triggers: []string{
			"grid regions / components → DFS or BFS + visited",
			"shortest path unweighted grid → BFS",
			"flood fill → DFS/BFS from start, recolor connected cells",
		},
		understandWarmup: "Island count: each unvisited land cell starts one DFS/BFS component.",
	},
}

var bonusDrills = []string{
	"08_heap_reflex",
	"09_backtrack_reflex",
	"10_math_reflex",
}

const mathReflexFile = "10_math_reflex"

var essentialCatalog = []struct {
	group string
	fns   []string
}{
	{"Arrays & prefix", []string{"reverseInPlace", "indexOfMax", "arraySum", "rotateRight", "runningSum"}},
	{"Hashing", []string{"twoSum", "containsDuplicate", "frequencyMap", "firstUniqueChar", "groupAnagrams"}},
	{"Two pointers & window", []string{"removeDuplicates", "moveZeroes", "maxArea", "isPalindrome", "maxSumSubarrayK"}},
	{"Binary search", []string{"binarySearch", "searchInsert", "findMinRotated", "isTargetPresent"}},
	{"Trees & stacks", []string{"inorderTraversal", "maxDepth", "isValidParentheses", "dailyTemperatures"}},
	{"DP", []string{"fib", "climbStairs", "minCostClimbingStairs", "rob"}},
	{"Graphs", []string{"numIslands", "floodFill", "shortestPathGrid"}},
	{"Math", []string{"gcd", "lcm", "modPow", "nCr", "isPrime", "powOfTwo"}},
	{"Heaps (bonus)", []string{"kthLargest", "lastStoneWeight", "mergeKSorted"}},
	{"Backtracking (bonus)", []string{"subsets", "permute", "combine"}},
}

var allTriggers = []string{
	"pair sums to target → map + complement",
	"duplicates / seen before → set or freq map",
	"sorted + two values / area → L/R two pointers",
	"in-place filter / dedupe → read/write pointers",
	"subarray size k → fixed sliding window",
	"sorted + find / insert → binary search or lower bound",
	"matching brackets → stack",
	"next greater → monotonic stack",
	"min cost / ways → 1D DP (define dp[i] first)",
	"grid regions / fill → DFS or BFS + visited",
	"shortest unweighted path → BFS",
	"gcd / lcm / modPow → Euclidean + fast exponentiation",
	"count combinations → nCr with symmetry k = min(k, n-k)",
}

func core5Names() string {
	names := make([]string, len(core5))
	for i, fn := range core5 {
		name := fn.name
		if idx := strings.Index(name, "("); idx > 0 {
			name = name[:idx]
		}
		names[i] = name
	}
	return strings.Join(names, ", ")
}

func printCatalog() {
	fmt.Println("WRITE catalog")
	for _, entry := range essentialCatalog {
		fmt.Printf("%s: %s\n", entry.group, strings.Join(entry.fns, ", "))
	}
}

func printDrill(today drill, brief bool) {
	if brief {
		fmt.Println("write: drills/write/core5/")
		return
	}
	fmt.Printf("WRITE %s | core 5\n", today.day)
	fmt.Printf("core5: %s\n", core5Names())
	fmt.Println("path: drills/write/core5/")
}

func printToday(today drill, brief bool) {
	if brief {
		fmt.Printf("write: %s\n", today.file)
		fmt.Printf("       core5: %s\n", core5Names())
		fmt.Printf("       specialty: %s\n", strings.Join(today.functions, ", "))
		return
	}
	fmt.Printf("WRITE %s | %s\n", today.day, today.file)
	fmt.Printf("core5: %s\n", core5Names())
	fmt.Printf("specialty: %s\n", strings.Join(today.functions, ", "))
	fmt.Printf("path: drills/write/reflex/%s/\n", today.file)
	fmt.Println("run:    go run . -- --run core")
	fmt.Println("        go run . -- --run reflex")
	fmt.Printf("math:   go run . -- --run-math  (%s)\n", mathReflexFile)
}

func hasFlag(flag string) bool {
	for _, a := range os.Args[1:] {
		if a == flag {
			return true
		}
	}
	return false
}

func main() {
	today := todayDrillFromWeekday(time.Now().Weekday())

	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repoRoot := findRepoRoot(root)
	_, drillPath := resolvePlayPaths(root, today.file)

	drill, brief, runMath, runMode := parsePlayArgs(os.Args[1:])

	if hasFlag("--weak") {
		printWeakFunctions(repoRoot, 5)
		return
	}
	if hasFlag("--setup") {
		fmt.Println("Setting up write reflex drills from blanks/ ...")
		fmt.Println()
		if err := setupAllDrills(repoRoot); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("Setup complete. Start with: go run .")
		return
	}

	if hasFlag("--catalog") {
		printCatalog()
		return
	}
	if drill {
		printDrill(today, brief)
		return
	}
	if hasFlag("--reset") {
		if err := resetTodayDrill(today, drillPath); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	printToday(today, brief)

	if runMath {
		mathPath := writeReflexDir(repoRoot, mathReflexFile)
		ok, output, _ := runDrillWithLog(mathPath)
		fmt.Print(output)
		mathFns := []string{"gcd", "lcm", "modPow", "nCr", "isPrime", "powOfTwo"}
		if !ok {
			updateLogFromOutput(repoRoot, output, mathFns)
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, mathFns)
		return
	}

	core5Fns := []string{"twoSum", "binarySearch", "removeDuplicates", "maxSumSubarrayK", "frequencyMap"}
	core5Path := writeCore5Dir(repoRoot)

	switch runMode {
	case "core":
		ok, output, _ := runDrillWithLog(core5Path)
		fmt.Print(output)
		if !ok {
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, core5Fns)
	case "reflex":
		ok, output, _ := runDrillWithLog(drillPath)
		fmt.Print(output)
		if !ok {
			updateLogFromOutput(repoRoot, output, today.functions)
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, today.functions)
	case "all":
		ok, output, _ := runDrillWithLog(core5Path)
		fmt.Print(output)
		if !ok {
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, core5Fns)
		ok, output, _ = runDrillWithLog(drillPath)
		fmt.Print(output)
		if !ok {
			updateLogFromOutput(repoRoot, output, today.functions)
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, today.functions)
	}
}
