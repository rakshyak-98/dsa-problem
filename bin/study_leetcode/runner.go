package main

import (
	"fmt"
	"os"
	"strings"
)

func printProblem(p lcProblem, n int) {
	dailyMark := ""
	if p.Daily {
		dailyMark = " [daily challenge]"
	}
	fmt.Printf("  %2d. #%d  %-42s  %-6s  %s%s\n", n, p.Num, p.Title, p.Diff, p.Pattern, dailyMark)
	fmt.Printf("      %s\n", lcURL(p.Slug))
	if p.ReflexFn != "" {
		fmt.Printf("      reflex: %s\n", p.ReflexFn)
	}
}

func printTodaySet(s practiceSet, brief bool) {
	if brief {
		fmt.Printf("leetcode:  %s — 10 problems (%s)\n", s.topic, s.reflex)
		return
	}
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║         DAILY LEETCODE PRACTICE — 10 QUESTIONS           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("TODAY %s | Topic: %s\n", s.day, s.topic)
	fmt.Printf("Reflex drill (separate): drills/write/reflex/%s\n", s.reflex)
	fmt.Printf("Warmup: %s\n", s.warmup)
	fmt.Printf("Saved:  drills/leetcode/daily.json\n")
	fmt.Printf("        drills/leetcode/daily.go    (full problem statements)\n")
	fmt.Printf("        drills/leetcode/daily.md    (markdown copy)\n")
	fmt.Println()
	fmt.Println("Solve on LeetCode — then run: go run . -- --levels")
	fmt.Println()
	fmt.Println("━━━ TODAY'S 10 PROBLEMS (from LeetCode API) ━━━━━━━━━━━━")
	for i, p := range s.problems {
		printProblem(p, i+1)
	}
	fmt.Println()
	fmt.Println("━━━ SUGGESTED ORDER ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, tip := range s.suggested {
		fmt.Printf("  %d. %s\n", i+1, tip)
	}
	fmt.Println()
	fmt.Println("run:     go run . -- --run leetcode")
	fmt.Println("         go run . -- --run -l")
	fmt.Println("refresh: go run . -- --run leetcode -- --refresh")
	fmt.Println("catalog: go run . -- --track leetcode -- --catalog")
	fmt.Println("reflex:  go run . -- --drill reflex   (in-repo micro-functions)")
}

func printCatalog() {
	fmt.Println("LeetCode practice sets, by weekday (10 problems each):")
	fmt.Println()
	for _, s := range practiceSets {
		fmt.Printf("  %-9s  %-20s  %-24s  %s\n",
			s.day, s.topic, s.reflex, strings.Join(s.topicTags, ", "))
	}
	fmt.Println()
	fmt.Println("Run today's set with 'go run . -- --run leetcode'.")
}

func runStudyLeetcode(repoRoot string, catalog, brief, showSet, refresh bool) int {
	if catalog {
		printCatalog()
		return 0
	}
	if showSet {
		set, err := ensureTodaySet(repoRoot, refresh)
		if err != nil {
			fmt.Fprintf(os.Stderr, "leetcode fetch failed: %v\n", err)
			return 1
		}
		printTodaySet(set, brief)
		return 0
	}
	return 0
}

func main() {
	opts, ctl, parseErr := parseLeetcode(os.Args[1:])
	switch ctl {
	case gnuHelp:
		printHelp()
		return
	case gnuVersion:
		printVersion(leetcodeProg)
		return
	case gnuErr:
		os.Exit(2)
	}
	if parseErr {
		os.Exit(2)
	}
	repoRoot := findRepoRoot(mustGetwd())
	if code := runStudyLeetcode(repoRoot, opts.catalog, opts.brief, opts.showSet, opts.refresh); code != 0 {
		os.Exit(code)
	}
}

func mustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return wd
}
