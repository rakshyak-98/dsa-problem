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
		patterns:  "reverse, rotate, prefix, prefix+map, Kadane, merge/quick sort, sieve",
		functions: []string{"reverseInPlace", "rotateRight", "runningSum", "subarraySumK", "productExceptSelf", "maxSubarraySum", "mergeSort", "quickSort", "sieve"},
		triggers: []string{
			"in-place mutate / reverse → two pointers L/R swap",
			"range sum / running total → prefix array",
			"count subarrays summing to k → prefix sum + map of seen prefixes",
			"best contiguous run → Kadane (extend or restart at nums[i])",
			"rotate by k → k %= n, then reverse sections (or copy with modulo)",
			"implement a sort → merge sort (stable, split+merge) or quick sort (in-place partition)",
			"all primes up to n → sieve of Eratosthenes (cross out multiples from p*p)",
		},
		understandWarmup: "Subarray sum equals k: count contiguous runs, not pairs — a prefix seen before means the gap between them sums to k.",
	},
	{
		day: "Tuesday", file: "02_hashing_reflex",
		patterns:  "two-sum, dup, freq map, unique char, anagrams, XOR single",
		functions: []string{"twoSum", "containsDuplicate", "frequencyMap", "firstUniqueChar", "groupAnagrams", "singleNumber"},
		triggers: []string{
			"pair sums to target → map + complement",
			"seen before? → set / map",
			"same letter multiset → sorted key or count signature",
			"one unpaired value, O(1) space → XOR the whole array (no hash set)",
		},
		understandWarmup: "Two-sum: return indices of two numbers that add to target (usually one valid pair).",
	},
	{
		day: "Wednesday", file: "03_two_pointers_reflex",
		patterns:  "dedupe, zeroes, container, palindrome, fixed + variable window, Dutch flag",
		functions: []string{"removeDuplicates", "moveZeroes", "maxArea", "isPalindrome", "maxSumSubarrayK", "longestUniqueSubstring", "dutchFlag"},
		triggers: []string{
			"sorted + two values → L/R two pointers",
			"in-place filter / dedupe → read/write pointers",
			"subarray length k stated → fixed sliding window",
			"longest/shortest run with a property → variable window, shrink from the left",
			"array of a few distinct values, sort in one pass → Dutch national flag (low/mid/high)",
		},
		understandWarmup: "Fixed vs variable window: k given in the problem means fixed; \"longest such that…\" means the left edge moves on its own.",
	},
	{
		day: "Thursday", file: "04_binary_search_reflex",
		patterns:  "exact BS, lower bound, rotated min, search the answer, quickselect, fast power, gcd",
		functions: []string{"binarySearch", "searchInsert", "findMinRotated", "minEatingSpeed", "quickSelect", "fastPow", "gcd"},
		triggers: []string{
			"sorted + find exact → lo <= hi, mid compare",
			"first position ≥ target → lower bound / searchInsert",
			"rotated sorted min → decide which half is sorted",
			"minimum rate / capacity that still works → binary search the answer space",
			"kth largest / smallest in O(n) average → quickselect (partition, recurse one half)",
			"x^n fast, or huge exponent → binary exponentiation (square base, halve exponent)",
			"gcd / lcm / reduce a fraction → Euclid (a, b = b, a mod b)",
		},
		understandWarmup: "Searching the answer: when the input is not sorted but \"does speed s work?\" is monotonic, binary search over s.",
	},
	{
		day: "Friday", file: "05_trees_stacks_reflex",
		patterns:  "inorder, preorder, postorder, level-order, depth, BST bounds, parens, mono stack",
		functions: []string{"inorderTraversal", "preorderTraversal", "postorderTraversal", "levelOrderTraversal", "maxDepth", "isValidBST", "isValidParentheses", "dailyTemperatures"},
		triggers: []string{
			"tree order without recursion → stack / iterative DFS",
			"row-by-row tree visit → BFS queue (level-order)",
			"BST validity → carry (lo, hi) bounds down, never compare to the parent alone",
			"matching brackets → stack of opens",
			"next greater element → monotonic decreasing stack",
		},
		understandWarmup: "Valid parentheses: every closer must match the latest unmatched opener.",
	},
	{
		day: "Saturday", file: "06_dp_reflex",
		patterns:  "climb stairs, min cost, rob, coin change",
		functions: []string{"climbStairs", "minCostClimbingStairs", "rob", "coinChange"},
		triggers: []string{
			"min cost / max ways on a line → 1D DP",
			"define dp[i] in English before coding",
			"rob houses → cannot take adjacent → max(take, skip)",
			"fewest items making a total, reuse allowed → unbounded knapsack over amounts",
		},
		understandWarmup: "Coin change: dp[a] is the fewest coins making exactly a; greedy fails, so try every coin at every amount.",
	},
	{
		day: "Sunday", file: "07_graphs_reflex",
		patterns:  "islands, flood fill, BFS path, topological order, DFS/BFS, Dijkstra",
		functions: []string{"numIslands", "floodFill", "shortestPathGrid", "canFinish", "dfs", "bfs", "bfsShortestPath", "topoSort", "dijkstra"},
		triggers: []string{
			"grid regions / components → DFS or BFS + visited",
			"shortest path unweighted grid → BFS",
			"flood fill → DFS/BFS from start, recolor connected cells",
			"prerequisites / ordering / \"is there a cycle\" → topological sort on in-degrees",
			"visit every reachable node → DFS (recursion/stack) or BFS (queue)",
			"fewest edges between two nodes → BFS from the source, count levels",
			"cheapest weighted path, non-negative edges → Dijkstra (min-heap frontier)",
		},
		understandWarmup: "Topological sort: repeatedly take a node whose prerequisites are all met; anything left over sits on a cycle.",
	},
}

var bonusDrills = []string{
	"08_heap_reflex",
	"09_backtrack_reflex",
	"10_linked_list_reflex",
}

var essentialCatalog = []struct {
	group string
	fns   []string
}{
	{"Arrays & prefix", []string{"reverseInPlace", "rotateRight", "runningSum", "subarraySumK", "productExceptSelf", "maxSubarraySum", "mergeSort", "quickSort", "sieve"}},
	{"Hashing", []string{"twoSum", "containsDuplicate", "frequencyMap", "firstUniqueChar", "groupAnagrams", "singleNumber"}},
	{"Two pointers & window", []string{"removeDuplicates", "moveZeroes", "maxArea", "isPalindrome", "maxSumSubarrayK", "longestUniqueSubstring", "dutchFlag"}},
	{"Binary search", []string{"binarySearch", "searchInsert", "findMinRotated", "minEatingSpeed", "quickSelect", "fastPow", "gcd"}},
	{"Trees & stacks", []string{"inorderTraversal", "preorderTraversal", "postorderTraversal", "levelOrderTraversal", "maxDepth", "isValidBST", "isValidParentheses", "dailyTemperatures"}},
	{"DP", []string{"climbStairs", "minCostClimbingStairs", "rob", "coinChange"}},
	{"Graphs", []string{"numIslands", "floodFill", "shortestPathGrid", "canFinish", "dfs", "bfs", "bfsShortestPath", "topoSort", "dijkstra"}},
	{"Heaps (bonus)", []string{"kthLargest", "lastStoneWeight", "mergeKSorted"}},
	{"Backtracking (bonus)", []string{"subsets", "permute", "combine"}},
	{"Linked lists (bonus)", []string{"reverseList", "hasCycle", "middleNode", "mergeTwoLists", "removeNthFromEnd"}},
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
	"prerequisites / ordering → topological sort",
	"count subarrays summing to k → prefix sum + map",
	"longest run with a property → variable sliding window",
	"min rate / capacity that works → binary search the answer",
	"reverse / detect a loop in a list → prev-cur rewire, or slow+fast pointers",
	"sort in O(n log n) → merge sort (stable) or quick sort (in place)",
	"kth largest / smallest in O(n) average → quickselect (partition, recurse one side)",
	"array of a few distinct values → three-way partition (Dutch national flag)",
	"x^n fast, or huge exponent → binary exponentiation (square, halve the exponent)",
	"gcd / lcm / reduce a fraction → Euclid (a, b = b, a mod b)",
	"all primes up to n → sieve of Eratosthenes",
	"one unpaired value, O(1) space → XOR the whole array",
	"visit every reachable node → DFS (stack / recursion) or BFS (queue)",
	"fewest edges between two nodes → BFS from the source, count levels",
	"directed graph, valid order / cycle check → Kahn topological sort",
	"cheapest weighted path, non-negative edges → Dijkstra (min-heap frontier)",
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

// formatInColumns lays names out in fixed-width columns, filled top-to-bottom
// then left-to-right like `ls`. colContent is the widest name across the whole
// listing, so callers pass one shared value and every block's columns line up.
// termWidth caps the line length (80 is the GNU default for a non-tty).
func formatInColumns(fns []string, indent string, colContent, termWidth int) string {
	if len(fns) == 0 {
		return ""
	}

	colWidth := colContent + 2
	cols := (termWidth - len(indent)) / colWidth
	if cols < 1 {
		cols = 1
	}
	rows := (len(fns) + cols - 1) / cols

	var b strings.Builder
	for row := 0; row < rows; row++ {
		b.WriteString(indent)
		for col := 0; col < cols; col++ {
			idx := col*rows + row
			if idx >= len(fns) {
				break
			}
			fn := fns[idx]
			b.WriteString(fn)
			if (col+1)*rows+row < len(fns) {
				b.WriteString(strings.Repeat(" ", colWidth-len(fn)))
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func printCatalog() {
	fmt.Println("Write drills, grouped by pattern:")

	colContent := 0
	for _, entry := range essentialCatalog {
		for _, fn := range entry.fns {
			if len(fn) > colContent {
				colContent = len(fn)
			}
		}
	}

	for _, entry := range essentialCatalog {
		fmt.Println()
		fmt.Printf("%s:\n", entry.group)
		fmt.Print(formatInColumns(entry.fns, "  ", colContent, 80))
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

// hasFlag reports whether flag was passed on the command line. Retained for
// callers that inspect os.Args directly; main uses the parsed playOpts.
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

	opts, ctl, parseErr := parsePlay(os.Args[1:])
	switch ctl {
	case gnuHelp:
		printHelp()
		return
	case gnuVersion:
		printVersion(playProg)
		return
	case gnuErr:
		os.Exit(2)
	}
	if parseErr {
		os.Exit(2)
	}
	brief := opts.brief
	runMode := opts.runMode

	if opts.weak {
		printWeakFunctions(repoRoot, 5)
		return
	}
	if opts.setup {
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

	if opts.catalog {
		printCatalog()
		return
	}
	if opts.triggers {
		printAllTriggers()
		return
	}
	if opts.problems {
		printProblemSet(repoRoot)
		return
	}
	if opts.levels {
		printLevels(repoRoot)
		return
	}
	if opts.refresh {
		printRefresh(buildRefresh(repoRoot, today, time.Now()), opts.show)
		return
	}
	if opts.drillKind == "core" {
		printDrill(today, brief)
		return
	}
	if opts.drillKind == "reflex" {
		printReflexDrill(today, brief)
		return
	}
	if opts.solutionKind == "core" {
		printSolutionCore(brief)
		return
	}
	if opts.solutionKind == "reflex" {
		printSolutionReflex(today, brief)
		return
	}
	if opts.reset {
		if err := resetTodayDrill(today, drillPath); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if runMode == "" {
		printToday(today, brief)
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
