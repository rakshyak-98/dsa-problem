package main

import (
	"fmt"
	"os"
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
	fmt.Println()
	fmt.Println("Solve on LeetCode — log progress in drills/tracker/study_tracker.html")
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
	fmt.Println("refresh: go run . -- --track leetcode -- --refresh")
	fmt.Println("catalog: go run . -- --track leetcode -- --catalog")
	fmt.Println("reflex:  go run . -- --drill reflex   (in-repo micro-functions)")
}

func printCatalog() {
	fmt.Println("LEETCODE PRACTICE catalog — 10 API-fetched problems per weekday")
	fmt.Println()
	for _, s := range practiceSets {
		fmt.Printf("%-9s  %-18s  %-22s  tags: %v\n", s.day, s.topic, s.reflex, s.topicTags)
	}
	fmt.Println()
	fmt.Println("Run today's set: go run . -- --track leetcode")
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
	help, catalog, brief, showSet, refresh, parseErr := parseLeetcodeArgs(os.Args[1:])
	if parseErr {
		os.Exit(1)
	}
	if help {
		printHelp()
		return
	}
	repoRoot := findRepoRoot(mustGetwd())
	if code := runStudyLeetcode(repoRoot, catalog, brief, showSet, refresh); code != 0 {
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
