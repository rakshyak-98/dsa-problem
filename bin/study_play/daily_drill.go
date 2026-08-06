// Daily reflex practice helper — Core 5 + specialty drill
//
// RUN:              go run .
// RUN with tests:   go run . -- --run
// Core 5 only:      go run . -- --micro
// Full catalog:     go run . -- --catalog
// Reset today:      go run . -- --reset
// First-time setup: go run . -- --setup
package main

import (
	"fmt"
	"os"
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
}

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
}

func printBanner() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════╗
║         DAILY REFLEX PRACTICE — ESSENTIAL PACK           ║
╚══════════════════════════════════════════════════════════╝`)
}

func printCore5() {
	fmt.Print(`── CORE 5 (every day — target < 8 min) ───────────────
  Say the ask, then write blind.

`)
	for i, fn := range core5 {
		fmt.Printf("  %d. %s\n", i+1, fn.name)
		fmt.Printf("     Ask:     %s\n", fn.ask)
		fmt.Printf("     Pattern: %s\n", fn.pattern)
		fmt.Printf("     Target:  %ds\n\n", fn.sec)
	}
}

func printCatalog() {
	printBanner()
	fmt.Println("Essential catalog — own every function blind.")
	fmt.Println()
	for _, entry := range essentialCatalog {
		fmt.Printf("  %s\n", entry.group)
		for _, fn := range entry.fns {
			fmt.Printf("    [ ] %s\n", fn)
		}
		fmt.Println()
	}
	fmt.Println("Guide: doc/write/DAILY_30MIN_DRILL.md")
}

func printMicro(today drill) {
	printBanner()
	fmt.Printf("Minimum / micro day (%s) — Core 5 only.\n\n", today.day)
	printCore5()
	fmt.Printf(`When ready for specialty:
  Open: drills/write/reflex/%s
  Run:  go run . -- --run
`, today.file)
}

func printWriteSpecialty(today drill) {
	fmt.Println("════════════════════════════════════════")
	fmt.Printf(" SPECIALTY — %s (%s)\n", today.day, today.file)
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("  Patterns: %s\n", today.patterns)
	fmt.Printf("  Warm-up: %s\n", today.understandWarmup)
	fmt.Println("  Functions:")
	for _, fn := range today.functions {
		fmt.Printf("    • %s\n", fn)
	}
	fmt.Printf("\n  Open:  drills/write/reflex/%s/main.go\n", today.file)
	fmt.Printf("  Run:   go run -C drills/write/reflex/%s .\n", today.file)
	fmt.Println("  Solutions: drills/solutions/reflex/<today's drill>/ (after honest attempt)")
	fmt.Println()
}

func printToday(today drill) {
	printBanner()
	fmt.Printf("Today: %s\nSpecialty file: drills/write/reflex/%s\nSpecialty set:  %s\n\n",
		today.day, today.file, today.patterns)

	if today.day == "Sunday" {
		fmt.Println("Note: Sunday is rest from new problems. Graphs specialty is optional; still do Core 5.")
		fmt.Println()
	}

	fmt.Print(`── TIERS ──────────────────────────────────────────────
  Minimum   ~20-30 min   Core 5 + log
  Reflex    ~30-40 min   Core 5 + today's specialty drill
  Standard  45-60 min    Reflex + ONE primary from STUDY_PLAN.md
`)
	printCore5()

	fmt.Println("── SPECIALTY FUNCTIONS (blind write after Core 5) ──")
	for _, fn := range today.functions {
		fmt.Printf("  • %s\n", fn)
	}

	fmt.Print(`
── REFLEX CLOCK ────────────────────────────────────────
  0-2 min    Full trigger scan (out loud)
  2-10 min   Core 5 blind write
  10-12 min  Understand warm-up (specialty)
  12-32 min  Specialty TODO: REFLEX functions
  32-37 min  Run tests & fix once (no solutions)
  37-40 min  Log fails + revisit (+3 days)

── TODAY'S TRIGGERS ────────────────────────────────────`)
	for _, t := range today.triggers {
		fmt.Printf("  • %s\n", t)
	}

	fmt.Println("\n── ALWAYS-ON TRIGGERS (scan daily) ─────────────────────")
	for _, t := range allTriggers {
		fmt.Printf("  • %s\n", t)
	}

	fmt.Printf(`
── UNDERSTAND WARM-UP (say aloud) ──────────────────────
  %s

── COMMANDS ────────────────────────────────────────────
  Core 5 drill:    go run -C drills/write/core5 .
  Open specialty:  drills/write/reflex/%s
  Run specialty:   go run . -- --run
  Run Core 5:      go run . -- --run-core5
  Reset specialty: go run . -- --reset
  Weak functions:  go run . -- --weak
  Variants:        go run -C drills/write/variants .
  Solutions:       drills/solutions/reflex/<today's drill>/ (after honest attempt)
  Core 5 only:     go run . -- --micro
  Full catalog:    go run . -- --catalog
  Full guide:      doc/write/DAILY_30MIN_DRILL.md
`, today.understandWarmup, today.file)

	printAskWarmup(today.day)
	printProblemMap(today.file)
	printCore5Problems()
	printVisualizerLink(today.file)
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
	if hasFlag("--specialty") {
		if hasFlag("--micro") {
			return
		}
		printWriteSpecialty(today)
		if hasFlag("--run") {
			fmt.Println("Running specialty tests...")
			fmt.Println()
			ok, output, _ := runDrillWithLog(drillPath)
			fmt.Print(output)
			if !ok {
				updateLogFromOutput(repoRoot, output, today.functions)
				fmt.Println()
				fmt.Println("Tests failed — good data. Fix blind, then re-run.")
				os.Exit(1)
			}
			updateLogFromOutput(repoRoot, output, today.functions)
		}
		return
	}
	if hasFlag("--micro") {
		printMicro(today)
		return
	}
	if hasFlag("--reset") {
		if err := resetTodayDrill(today, drillPath); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	printToday(today)

	if hasFlag("--run-core5") {
		core5Path := writeCore5Dir(repoRoot)
		fmt.Println("Running Core 5 tests...")
		fmt.Println()
		ok, output, _ := runDrillWithLog(core5Path)
		if !ok {
			fmt.Print(output)
			fmt.Println()
			fmt.Println("Core 5 failed — fix blind, then re-run.")
			os.Exit(1)
		}
		fmt.Print(output)
		updateLogFromOutput(repoRoot, output, []string{"twoSum", "binarySearch", "removeDuplicates", "maxSumSubarrayK", "frequencyMap"})
		fmt.Println("Core 5 logged.")
		return
	}

	if hasFlag("--run") {
		fmt.Println("Running specialty tests...")
		fmt.Println()
		ok, output, _ := runDrillWithLog(drillPath)
		fmt.Print(output)
		if !ok {
			updateLogFromOutput(repoRoot, output, today.functions)
			fmt.Println()
			fmt.Println("Tests failed — good data. Fix blind, then re-run.")
			os.Exit(1)
		}
		updateLogFromOutput(repoRoot, output, today.functions)
		fmt.Println("Specialty logged. Next: solve today's primary problem (see map above).")
	} else {
		fmt.Println("Tip: after Core 5 + specialty, run  go run . -- --run")
		fmt.Println()
	}
}
