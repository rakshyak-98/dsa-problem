// Daily reflex practice helper — Core 5 + specialty drill
//
// RUN:              go run .
// Today's session:  go run . -- --refresh        (level-matched, log-driven)
// Reveal answers:   go run . -- --refresh --show
// Level scoreboard: go run . -- --levels
// Problem set:      go run . -- --problems
// RUN with tests:   go run . -- --run
// Core 5 only:      go run . -- --drill core
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
		patterns:  "inorder, preorder, postorder, level-order, depth, parens, mono stack",
		functions: []string{"inorderTraversal", "preorderTraversal", "postorderTraversal", "levelOrderTraversal", "maxDepth", "isValidParentheses", "dailyTemperatures"},
		triggers: []string{
			"tree order without recursion → stack / iterative DFS",
			"row-by-row tree visit → BFS queue (level-order)",
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
	{"Trees & stacks", []string{"inorderTraversal", "preorderTraversal", "postorderTraversal", "levelOrderTraversal", "maxDepth", "isValidParentheses", "dailyTemperatures"}},
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

func formatInColumns(fns []string, indent string) string {
	if len(fns) == 0 {
		return ""
	}

	termWidth := 120
	indentLen := len(indent)
	availWidth := termWidth - indentLen

	maxWidth := 0
	for _, fn := range fns {
		if len(fn) > maxWidth {
			maxWidth = len(fn)
		}
	}
	colWidth := maxWidth + 2

	cols := availWidth / colWidth
	if cols < 1 {
		cols = 1
	}

	rows := (len(fns) + cols - 1) / cols

	var result strings.Builder
	for row := 0; row < rows; row++ {
		result.WriteString(indent)
		for col := 0; col < cols; col++ {
			idx := col*rows + row
			if idx >= len(fns) {
				break
			}
			fn := fns[idx]
			result.WriteString(fn)
			if col < cols-1 && idx < len(fns)-1 {
				result.WriteString(strings.Repeat(" ", colWidth-len(fn)))
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func printCatalog() {
	fmt.Println("WRITE catalog")
	for _, entry := range essentialCatalog {
		fmt.Printf("%s:\n%s\n", entry.group, formatInColumns(entry.fns, "  "))
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
	fmt.Println("\n── SAY THE ASK, THEN WRITE ────────────────────────────")
	for _, fn := range core5 {
		fmt.Printf("  • %-28s %s\n", fn.name, fn.ask)
		fmt.Printf("    %-28s %s · target %ds\n", "", fn.pattern, fn.sec)
	}
	printCore5Problems()
}

func printReflexDrill(today drill, brief bool) {
	if brief {
		fmt.Printf("write: %s\n", today.file)
		return
	}
	fmt.Printf("WRITE %s | %s\n", today.day, today.file)
	fmt.Printf("specialty: %s\n", strings.Join(today.functions, ", "))
	fmt.Printf("path: drills/write/reflex/%s/\n", today.file)
}

func printSolutionCore(brief bool) {
	if brief {
		fmt.Println("write: drills/solutions/core5.md")
		return
	}
	fmt.Println("WRITE solution | core 5")
	fmt.Println("path: drills/solutions/core5.md")
}

func printSolutionReflex(today drill, brief bool) {
	if brief {
		fmt.Printf("write: drills/solutions/reflex/%s/main.go\n", today.file)
		return
	}
	fmt.Printf("WRITE solution %s | %s\n", today.day, today.file)
	fmt.Printf("path: drills/solutions/reflex/%s/main.go\n", today.file)
	fmt.Printf("notes: drills/solutions/%s.md\n", today.file)
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

	// Recognition before recall: the triggers and the one-sentence ask are read
	// out loud before any code is typed. They exist so the session rehearses
	// "which move does this want", not just "can I still type this function".
	printTriggers(today)
	printAskWarmup(today.day)
	printProblemMap(today.file)

	fmt.Println("\nrun:    go run . -- --run core")
	fmt.Println("        go run . -- --run reflex")
	fmt.Printf("math:   go run . -- --run-math  (%s)\n", mathReflexFile)
	fmt.Println("today:  go run . -- --refresh   (level-matched session)")
}

func printTriggers(today drill) {
	fmt.Println("\n── PATTERN TRIGGERS (say these before you type) ────────")
	for _, t := range today.triggers {
		fmt.Printf("  • %s\n", t)
	}
	if today.understandWarmup != "" {
		fmt.Printf("\n  Warm-up ask: %s\n", today.understandWarmup)
	}
}

// printAllTriggers is the full cross-topic table — the one to scan on a day
// when nothing else gets done.
func printAllTriggers() {
	fmt.Println("PATTERN TRIGGERS — all topics")
	for _, t := range allTriggers {
		fmt.Printf("  • %s\n", t)
	}
	fmt.Println("\n  Quiz yourself instead: go run . -- --refresh")
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

	drillKind, solutionKind, help, brief, runMath, runMode, parseErr := parsePlayArgs(os.Args[1:])
	if parseErr {
		os.Exit(1)
	}
	if help {
		printHelp()
		return
	}

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
	if hasFlag("--triggers") {
		printAllTriggers()
		return
	}
	if hasFlag("--problems") {
		printProblemSet(repoRoot)
		return
	}
	if hasFlag("--levels") {
		printLevels(repoRoot)
		return
	}
	if hasFlag("--refresh") {
		printRefresh(buildRefresh(repoRoot, today, time.Now()), hasFlag("--show"))
		return
	}
	if drillKind == "core" {
		printDrill(today, brief)
		return
	}
	if drillKind == "reflex" {
		printReflexDrill(today, brief)
		return
	}
	if solutionKind == "core" {
		printSolutionCore(brief)
		return
	}
	if solutionKind == "reflex" {
		printSolutionReflex(today, brief)
		return
	}
	if hasFlag("--reset") {
		if err := resetTodayDrill(today, drillPath); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if runMode == "" {
		printToday(today, brief)
	}

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
