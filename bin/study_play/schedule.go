package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Review queue
//
// The weekday drill is one topic. The review queue mixes in a few functions
// from other topics that are due under spaced repetition, so every pattern
// keeps coming back and you practise switching between them.

const (
	reviewCount   = 4 // functions in the daily review block
	reviewPerFile = 2 // at most this many from one drill file
)

type dueItem struct {
	fn      string
	file    string
	overdue int  // days past due; 0 = due today
	fresh   bool // not on the schedule yet
	rec     fnRecord
	rank    float64
}

// dueFunctions returns up to n functions due for review, skipping skipFile
// (today's main drill). Functions already on the schedule rank by how overdue
// they are; ones not on it yet (history from before the schedule existed)
// rank by how weak they are.
func dueFunctions(log drillLog, now time.Time, n int, skipFile string) []dueItem {
	today := dateOf(now)
	var items []dueItem
	for _, d := range drills {
		if d.file == skipFile {
			continue
		}
		for _, fn := range d.functions {
			rec := log.Functions[fn]
			it := dueItem{fn: fn, file: d.file, rec: rec}
			if rec.Due != "" {
				if rec.Due > today {
					continue
				}
				due, _ := time.Parse(dateLayout, rec.Due)
				it.overdue = daysBetween(due, now)
				it.rank = 10 + float64(it.overdue) + weakScore(rec, now)
				if rec.Box == 0 {
					it.rank += 5
				}
			} else {
				it.fresh = true
				it.rank = weakScore(rec, now)
			}
			items = append(items, it)
		}
	}
	sort.SliceStable(items, func(i, j int) bool { return items[i].rank > items[j].rank })

	// One per file first, so the block spans as many topics as possible; only
	// then a second (up to reviewPerFile) from the same file.
	perFile := map[string]int{}
	taken := map[string]bool{}
	var out []dueItem
	for limit := 1; limit <= reviewPerFile; limit++ {
		for _, it := range items {
			if len(out) >= n {
				return out
			}
			if taken[it.fn] || perFile[it.file] >= limit {
				continue
			}
			perFile[it.file]++
			taken[it.fn] = true
			out = append(out, it)
		}
	}
	return out
}

// lastSession is the most recent day any function was graded, or "".
func lastSession(log drillLog) string {
	last := ""
	for _, rec := range log.Functions {
		for _, d := range []string{rec.LastPass, rec.LastFail} {
			if d > last {
				last = d
			}
		}
	}
	return last
}

func printSessionGap(log drillLog, now time.Time) {
	last := lastSession(log)
	if last == "" {
		return
	}
	t, err := time.Parse(dateLayout, last)
	if err != nil {
		return
	}
	switch gap := daysBetween(t, now); {
	case gap == 0:
		fmt.Println("last session: today")
	case gap == 1:
		fmt.Println("last session: yesterday")
	default:
		fmt.Printf("last session: %d days ago (%s) — reviews pile up while you are away\n", gap, last)
	}
}

func describeDue(it dueItem) string {
	switch {
	case it.fresh && it.rec.Fails > 0:
		return fmt.Sprintf("not scheduled yet · %d fails so far", it.rec.Fails)
	case it.fresh && it.rec.Passes > 0:
		return "not scheduled yet"
	case it.fresh:
		return "never drilled"
	case it.overdue == 0:
		return fmt.Sprintf("due today · box %d", it.rec.Box)
	default:
		return fmt.Sprintf("overdue %dd · box %d", it.overdue, it.rec.Box)
	}
}

// printReview prints today's review block: mixed-topic functions to write cold.
func printReview(log drillLog, now time.Time, skipFile string, brief bool) {
	items := dueFunctions(log, now, reviewCount, skipFile)
	if len(items) == 0 {
		return
	}
	names := make([]string, len(items))
	for i, it := range items {
		names[i] = it.fn
	}
	list := strings.Join(names, ",")
	if brief {
		fmt.Printf("       review: %s\n", strings.Join(names, ", "))
		fmt.Printf("       blank:  go run . -- --fn=%s\n", list)
		return
	}
	fmt.Println("\n── REVIEW — mixed topics, due today ───────────────────")
	fmt.Println("  Name the pattern and write the skeleton first, then the code. Run once.")
	for i, it := range items {
		fmt.Printf("  %d. %-24s %-24s %s\n", i+1, it.fn, it.file, describeDue(it))
	}
	fmt.Printf("\n  blank: go run . -- --fn=%s\n", list)
	fmt.Printf("  check: go run . -- --run --fn=%s\n", list)
}
