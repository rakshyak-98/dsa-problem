// Daily code-reading helper — reflex specialty drill
//
// RUN:              go run .
// RUN with checks:  go run . -- --run
// Reflex only:      go run . -- --drill reflex
// Full catalog:     go run . -- --catalog
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type drill struct {
	day    string
	file   string
	skill  string
	focus  []string
	warmup string
}

var drills = []drill{
	{
		day: "Monday", file: "01_scan_structure",
		skill: "Structure scan",
		focus: []string{
			"params / return type / mutation",
			"loop nest vs single pass",
			"early returns and base cases",
		},
		warmup: "Before logic: mark every loop and every return.",
	},
	{
		day: "Tuesday", file: "02_trace_execution",
		skill: "Hand trace",
		focus: []string{
			"pick the given sample",
			"update binders each step",
			"predict return before checking",
		},
		warmup: "If the table disagrees with your gut, trust the table.",
	},
	{
		day: "Wednesday", file: "03_name_the_pattern",
		skill: "Pattern from shape",
		focus: []string{
			"ignore misleading names",
			"match skeleton to template",
			"one breath label",
		},
		warmup: "Shape beats function name — names may be wrong on purpose.",
	},
	{
		day: "Thursday", file: "04_find_the_bug",
		skill: "Bug hunt",
		focus: []string{
			"off-by-one on bounds",
			"wrong pointer move",
			"missing visit / wrong compare",
		},
		warmup: "Assume one bug. Find the line that breaks the invariant.",
	},
	{
		day: "Friday", file: "05_complexity_glance",
		skill: "Complexity from shape",
		focus: []string{
			"nested vs amortized two-pointer",
			"map/set space",
			"recursion depth",
		},
		warmup: "Count how many times the inner work runs across the whole input.",
	},
	{
		day: "Saturday", file: "06_reconstruct_ask",
		skill: "Ask from code",
		focus: []string{
			"one sentence, no jargon",
			"index vs value vs count",
			"contiguous vs not",
		},
		warmup: "If a junior couldn't understand your sentence, rewrite it.",
	},
	{
		day: "Sunday", file: "07_compare_variants",
		skill: "Compare two solutions",
		focus: []string{
			"same ask, different structure",
			"name the tradeoff (time/space/clarity)",
			"when you'd pick each",
		},
		warmup: "Diff the skeletons first; ignore shared boilerplate.",
	},
}

// readQ is one small-skill question for reading code: the mechanics (loop
// bounds, index and window math, invariants, cost) that a reader must get right
// before the shape of the algorithm even matters.
type readQ struct {
	q      string
	answer string
}

// showReadAnswers is set by --show so answers print under each question.
var showReadAnswers bool

var readSkills = map[string][]readQ{
	"Monday": {
		{"A function takes nums []int and returns nothing. What does that tell you before reading the body?",
			"It mutates nums in place (or has another side effect); the result lives in the argument."},
		{"One loop nested inside another versus two loops one after the other: what does each do to the cost, at a glance?",
			"Nested multiplies the work (O(n^2)); sequential loops add (O(n))."},
		{"The condition is i+1 < len(a) and the body reads a[i+1]. For len 5, what is the last i that runs?",
			"i = 3, which reads a[4]; i = 4 would read a[5], out of range."},
	},
	"Tuesday": {
		{"Trace: sum = 0; for i = 1; i <= 4; i++ { sum += i }. What is sum at the end?",
			"10 (1+2+3+4)."},
		{"left=0, right=4 inclusive, mid=(left+right)/2, and nums[mid] < target so left = mid+1. What are mid and left after one step?",
			"mid = 2, left = 3."},
		{"Inclusive window with left=2, right=5. What is its length, and its length after right++?",
			"4 (right-left+1), then 5."},
	},
	"Wednesday": {
		{"A write index w and read index r; copy when nums[r] != nums[w-1]. Name the pattern and say what w means.",
			"Read/write two pointers; w is the length of the kept prefix (the next free slot)."},
		{"A map of value to index, checked for target-x before inserting x. Which classic is this, and why look up first?",
			"Two sum with a complement map; inserting first could pair an element with itself."},
		{"A stack of indices whose values decrease; elements pop when a bigger value arrives. Name it and say what a pop means.",
			"Monotonic stack; the popped element's next greater value is the one that just arrived."},
	},
	"Thursday": {
		{"for lo <= hi { mid := (lo+hi)/2; if nums[mid] < t { lo = mid } else { hi = mid-1 } } — what breaks?",
			"lo = mid can stall forever (lo=3, hi=4); it must be lo = mid+1."},
		{"for i := 0; i <= len(a); i++ { use a[i] } — what is the bug?",
			"At i == len(a) the read is out of range; the bound should be i < len(a)."},
		{"The window is shrunk with an if, but one new element may need several left moves. What is the bug?",
			"It must be a for (while) loop; a single shrink may leave the window still invalid."},
	},
	"Friday": {
		{"for i := 0; i < n; i++ { for j := i; j < n; j++ { ... } } — how many inner iterations and what is the big-O?",
			"n(n+1)/2 iterations; O(n^2)."},
		{"A while-loop inside a for-loop over right moves left forward only. What is the total cost?",
			"O(n): left makes at most n moves across the whole run (amortized)."},
		{"Recursion on a balanced tree versus a skewed tree: how deep does the call stack go?",
			"O(log n) versus O(n)."},
	},
	"Saturday": {
		{"The code returns indices i, j with nums[i]+nums[j]==target using a map. State the ask in one sentence without naming a map.",
			"Find two positions whose values add up to the target."},
		{"A loop keeps cur = max(x, cur+x) and best = max(best, cur). State the ask in plain words.",
			"The largest sum of any contiguous run of the array."},
		{"A search with hi = mid and loop lo < hi returns lo. What does the result mean?",
			"The first index where the condition holds (a lower bound); n if it never does."},
	},
	"Sunday": {
		{"Recursive DFS versus DFS with an explicit stack: what differs?",
			"Same O(depth) extra space and visit order; recursion can overflow the call stack on deep input, an explicit stack lives on the heap."},
		{"Pair sum by sort plus two pointers versus a hash map: compare the tradeoffs.",
			"Sorting is O(n log n) time with O(1) extra space but loses original indices; the map is O(n) time and O(n) space and keeps indices."},
		{"Memoized recursion versus bottom-up DP: when would you pick each?",
			"Memo visits only reachable states and is easy to derive; bottom-up avoids recursion depth and allows space optimisation."},
	},
}

func printReadSkills(day string) {
	qs := readSkills[day]
	if len(qs) == 0 {
		return
	}
	fmt.Println("\n── SMALL SKILLS (say the answer and the reason) ────────")
	for i, q := range qs {
		fmt.Printf("  %d) %s\n", i+1, q.q)
		if showReadAnswers {
			fmt.Printf("     answer: %s\n", q.answer)
		}
	}
	if !showReadAnswers {
		fmt.Println("  check: go run . -- --show")
	}
}

func todayDrill() drill {
	wd := int(time.Now().Weekday())
	idx := (wd + 6) % 7
	return drills[idx]
}

func drillOpenPath(file string) string {
	return fmt.Sprintf("drills/read/weekday/%s/main.go", file)
}

func printReflexDrill(d drill, brief bool) {
	if brief {
		fmt.Printf("read:  %s\n", d.file)
		return
	}
	fmt.Printf("READ %s | %s — %s\n", d.day, d.file, d.skill)
	fmt.Printf("path: %s\n", drillOpenPath(d.file))
	printReadSkills(d.day)
}

func printSolutionReflex(d drill, brief bool) {
	if brief {
		fmt.Println("read:  drills/read/answers/ANSWER_KEY.md")
		return
	}
	fmt.Printf("READ solution %s | %s — %s\n", d.day, d.file, d.skill)
	fmt.Println("path: drills/read/answers/ANSWER_KEY.md")
	fmt.Printf("section: %s\n", d.file)
}

func printToday(d drill, brief bool) {
	if brief {
		fmt.Printf("read:  %s\n", d.file)
		return
	}
	fmt.Printf("READ %s | %s — %s\n", d.day, d.file, d.skill)
	fmt.Printf("path: %s\n", drillOpenPath(d.file))
	printReadSkills(d.day)
	fmt.Println("\nrun:    go run . -- --run reflex")
}

func printCatalog() {
	fmt.Println("Reflex read drills, by weekday:")
	fmt.Println()
	for _, d := range drills {
		fmt.Printf("  %-9s  %-22s  %s\n", d.day, d.file, d.skill)
	}
}

func runDrill(file string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := readDrillDir(findRepoRoot(cwd), file)
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	opts, ctl, parseErr := parseRead(os.Args[1:])
	switch ctl {
	case gnuHelp:
		printHelp()
		return
	case gnuVersion:
		printVersion(readProg)
		return
	case gnuErr:
		os.Exit(2)
	}
	if parseErr {
		os.Exit(2)
	}
	showReadAnswers = opts.show
	if code := runStudyCode(opts.drillKind, opts.solutionKind, opts.catalog, opts.brief, opts.runMode); code != 0 {
		os.Exit(code)
	}
}
