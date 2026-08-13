package main

import (
	"fmt"
	"os"
)

func printProblem(p lcProblem, n int) {
	fmt.Printf("  %2d. #%d  %-42s  %-6s  %s\n", n, p.num, p.title, p.diff, p.pattern)
	fmt.Printf("      %s\n", lcURL(p.num, slugFor(p.num)))
	if p.reflexFn != "" {
		fmt.Printf("      reflex: %s\n", p.reflexFn)
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
	fmt.Println()
	fmt.Println("Solve on LeetCode — log progress in drills/tracker/study_tracker.html")
	fmt.Println()
	fmt.Println("━━━ TODAY'S 10 PROBLEMS ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, p := range s.problems {
		printProblem(p, i+1)
	}
	fmt.Println()
	fmt.Println("━━━ SUGGESTED ORDER ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, tip := range s.suggested {
		fmt.Printf("  %d. %s\n", i+1, tip)
	}
	fmt.Println()
	fmt.Println("catalog: go run . -- --track leetcode -- --catalog")
	fmt.Println("reflex:  go run . -- --drill reflex   (in-repo micro-functions)")
}

func printCatalog() {
	fmt.Println("LEETCODE PRACTICE catalog — 10 problems per weekday")
	fmt.Println()
	for _, s := range practiceSets {
		fmt.Printf("%-9s  %-18s  %-22s  %d problems\n", s.day, s.topic, s.reflex, len(s.problems))
	}
	fmt.Println()
	fmt.Println("Run today's set: go run . -- --track leetcode")
}

func runStudyLeetcode(catalog, brief, showSet bool) int {
	if catalog {
		printCatalog()
		return 0
	}
	if showSet {
		printTodaySet(todaySet(), brief)
		return 0
	}
	return 0
}

func main() {
	help, catalog, brief, showSet, parseErr := parseLeetcodeArgs(os.Args[1:])
	if parseErr {
		os.Exit(1)
	}
	if help {
		printHelp()
		return
	}
	if code := runStudyLeetcode(catalog, brief, showSet); code != 0 {
		os.Exit(code)
	}
}
