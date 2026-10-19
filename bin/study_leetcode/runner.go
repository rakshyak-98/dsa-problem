package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// revealAnswers is set by --answers: show the topic, tags and reflex links
// that name each problem's pattern. They are hidden by default because
// naming the pattern is the first thing to practise.
var revealAnswers bool

func printProblem(p lcProblem, n int) {
	dailyMark := ""
	if p.Daily {
		dailyMark = " [daily challenge]"
	}
	tags := ""
	if revealAnswers {
		tags = p.Pattern
	}
	fmt.Printf("  %2d. #%d  %-42s  %-6s  %s%s\n", n, p.Num, p.Title, p.Diff, tags, dailyMark)
	fmt.Printf("      %s\n", lcURL(p.Slug))
	if p.ReflexFn != "" && revealAnswers {
		fmt.Printf("      reflex: %s\n", p.ReflexFn)
	}
}

func printTodaySet(s practiceSet, brief bool) {
	if brief {
		fmt.Printf("leetcode:  10 problems — name each pattern before solving\n")
		return
	}
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║         DAILY LEETCODE PRACTICE — 10 QUESTIONS           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	if revealAnswers {
		fmt.Printf("TODAY %s | Topic: %s\n", s.day, s.topic)
		fmt.Printf("Reflex drill (separate): drills/write/reflex/%s\n", s.reflex)
		fmt.Printf("Warmup: %s\n", s.warmup)
	} else {
		fmt.Printf("TODAY %s — topic and tags hidden: name each pattern first (--answers shows them)\n", s.day)
	}
	fmt.Printf("Saved:  drills/leetcode/daily.json\n")
	fmt.Printf("        drills/leetcode/daily.go    (full problem statements)\n")
	fmt.Printf("        drills/leetcode/daily.md    (markdown copy)\n")
	fmt.Println()
	fmt.Println("For each: write the pattern and the words that point to it, then solve on")
	fmt.Println("LeetCode, then record it: go run . -- --log=SLUG")
	fmt.Println()
	fmt.Println("━━━ TODAY'S 10 PROBLEMS (from LeetCode API) ━━━━━━━━━━━━")
	for i, p := range s.problems {
		printProblem(p, i+1)
	}
	fmt.Println()
	if revealAnswers {
		fmt.Println("━━━ SUGGESTED ORDER ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		for i, tip := range s.suggested {
			fmt.Printf("  %d. %s\n", i+1, tip)
		}
		fmt.Println()
	}
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
		printCardReview(loadCards(repoRoot), time.Now(), brief)
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
	revealAnswers = opts.answers
	if opts.gaps {
		printGaps(loadCards(repoRoot))
		return
	}
	if opts.log != "" {
		if err := logCard(repoRoot, opts.log, os.Stdin, os.Stdout, time.Now()); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", leetcodeProg, err)
			os.Exit(1)
		}
		return
	}
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
