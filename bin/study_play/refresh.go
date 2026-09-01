package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// The daily refresh
//
// The old plan printed a list of function names. Names cue recall, and recall
// is the skill that is already strongest — so the session rehearsed what was
// easy and never rehearsed the thing interviews actually test: reading an
// unfamiliar statement and knowing, in seconds, which move it wants.
//
// The refresh inverts that. It opens with cues that never name the function,
// then spends the session on whatever the log says is weakest, newest, or
// stale — instead of on whatever weekday it happens to be.

// drillPriority orders untouched drills by interview value, so "new ground"
// opens the pattern families that show up most, not the next file in the list.
var drillPriority = map[string]int{
	"02_hashing_reflex":       0,
	"04_binary_search_reflex": 1,
	"05_trees_stacks_reflex":  2,
	"03_two_pointers_reflex":  3,
	"01_arrays_reflex":        4,
	"06_dp_reflex":            5,
	"07_graphs_reflex":        6,
	"08_heap_reflex":          8,
	"09_backtrack_reflex":     9,
	"10_linked_list_reflex":   7,
}

const (
	recogniseCount = 4
	rebuildCount   = 3
	newGroundCount = 2
	stretchCount   = 2
)

type refreshSession struct {
	day       string
	specialty drill
	recognise []mastery
	rebuild   []mastery
	newGround []mastery
	stretch   []mastery
}

// weaknessScore ranks attempted functions: frequent failures first, then
// failures that have never been redeemed by a pass.
func weaknessScore(m mastery) float64 {
	total := m.passes + m.fails
	if total == 0 {
		return 0
	}
	score := float64(m.fails)/float64(total)*10 + float64(m.fails)
	if m.passes == 0 {
		score += 5
	}
	if m.daysIdle > m.earned.reviewGap() {
		score += 2
	}
	return score
}

// buildRefresh picks the session for a given day. The date seeds the shuffle so
// the plan is stable all day and rotates tomorrow.
func buildRefresh(root string, today drill, now time.Time) refreshSession {
	report := masteryReport(root)
	seed := int64(now.Year()*10000 + int(now.Month())*100 + now.Day())
	rng := rand.New(rand.NewSource(seed))

	s := refreshSession{day: today.day, specialty: today}
	taken := map[string]bool{}

	// Recognise: cues from functions with some history, one per drill family so
	// the round forces a cross-topic decision rather than a themed one.
	var pool []mastery
	for _, m := range report {
		if m.attempted() {
			pool = append(pool, m)
		}
	}
	if len(pool) < recogniseCount {
		pool = report
	}
	rng.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
	usedDrill := map[string]bool{}
	for _, m := range pool {
		if len(s.recognise) == recogniseCount {
			break
		}
		if usedDrill[m.cue.drill] {
			continue
		}
		usedDrill[m.cue.drill] = true
		taken[m.cue.fn] = true
		s.recognise = append(s.recognise, m)
	}
	for _, m := range pool {
		if len(s.recognise) == recogniseCount {
			break
		}
		if taken[m.cue.fn] {
			continue
		}
		taken[m.cue.fn] = true
		s.recognise = append(s.recognise, m)
	}

	// Rebuild: the weakest attempted functions, whatever weekday owns them.
	var weak []mastery
	for _, m := range report {
		if m.attempted() && m.fails > 0 {
			weak = append(weak, m)
		}
	}
	sort.Slice(weak, func(i, j int) bool {
		si, sj := weaknessScore(weak[i]), weaknessScore(weak[j])
		if si == sj {
			return weak[i].cue.fn < weak[j].cue.fn
		}
		return si > sj
	})
	for _, m := range weak {
		if len(s.rebuild) == rebuildCount {
			break
		}
		s.rebuild = append(s.rebuild, m)
	}

	// New ground: never attempted, highest-value family first.
	var fresh []mastery
	for _, m := range report {
		if !m.attempted() {
			fresh = append(fresh, m)
		}
	}
	sort.Slice(fresh, func(i, j int) bool {
		pi, pj := drillPriority[fresh[i].cue.drill], drillPriority[fresh[j].cue.drill]
		if pi != pj {
			return pi < pj
		}
		if fresh[i].cue.tier != fresh[j].cue.tier {
			return fresh[i].cue.tier < fresh[j].cue.tier
		}
		return fresh[i].cue.fn < fresh[j].cue.fn
	})
	for _, m := range fresh {
		if len(s.newGround) == newGroundCount {
			break
		}
		s.newGround = append(s.newGround, m)
	}

	// Stretch: functions that reached L3 and are due — drill the twist, not the
	// function you already own.
	for _, m := range report {
		if len(s.stretch) == stretchCount {
			break
		}
		if m.earned == levelTransfer && m.due() {
			s.stretch = append(s.stretch, m)
		}
	}
	return s
}

func statusOf(m mastery) string {
	switch {
	case !m.attempted():
		return "never attempted"
	case m.passes == 0:
		return fmt.Sprintf("%d fails, no pass yet", m.fails)
	case m.daysIdle < 0:
		return fmt.Sprintf("%d/%d pass", m.passes, m.passes+m.fails)
	default:
		return fmt.Sprintf("%d/%d pass, %dd idle", m.passes, m.passes+m.fails, m.daysIdle)
	}
}

func printRefresh(s refreshSession, showAnswers bool) {
	fmt.Printf("WRITE refresh | %s | %s\n", s.day, s.specialty.file)
	fmt.Println("Pattern first, code second. Nothing below is picked by weekday except the specialty.")

	fmt.Println("\n1. RECOGNISE — 2 min, no code. Read the cue, say the move out loud.")
	for i, m := range s.recognise {
		fmt.Printf("   %d) you see: %s\n", i+1, m.cue.cue)
		if showAnswers {
			fmt.Printf("      you write: %s   (%s)\n", m.cue.move, m.cue.fn)
		} else {
			fmt.Println("      you write: ______________________________")
		}
	}
	if !showAnswers {
		fmt.Println("   check: go run . -- --refresh --show")
	}

	if len(s.rebuild) > 0 {
		fmt.Println("\n2. REBUILD — weakest in the log. Blind write, then run.")
		for _, m := range s.rebuild {
			printMasteryLine(m)
		}
	}

	if len(s.newGround) > 0 {
		fmt.Println("\n3. NEW GROUND — never attempted, highest-value family first.")
		for _, m := range s.newGround {
			printMasteryLine(m)
		}
	}

	fmt.Printf("\n4. SPECIALTY — %s\n", s.specialty.file)
	fmt.Printf("   %s\n", strings.Join(s.specialty.functions, ", "))
	fmt.Printf("   path: drills/write/reflex/%s/\n", s.specialty.file)
	for _, t := range s.specialty.triggers {
		fmt.Printf("   trigger: %s\n", t)
	}

	if len(s.stretch) > 0 {
		fmt.Println("\n5. STRETCH — you own the move; apply it somewhere new.")
		for _, m := range s.stretch {
			fmt.Printf("   • %s (%s) → %s\n", m.cue.fn, m.earned.label(), m.cue.twist)
		}
	}

	fmt.Println("\nrun:      go run . -- --run reflex")
	fmt.Println("levels:   go run . -- --levels")
	fmt.Println("problems: go run . -- --problems")
}

func printMasteryLine(m mastery) {
	fmt.Printf("   • %-22s %-11s %s\n", m.cue.fn, m.earned.label(), statusOf(m))
	switch m.earned {
	case levelRecall:
		fmt.Printf("     ask:  %s\n", m.cue.ask)
		fmt.Printf("     move: %s\n", m.cue.move)
	case levelPattern:
		fmt.Printf("     cue:  %s\n", m.cue.cue)
		fmt.Printf("     ask:  %s\n", m.cue.ask)
	case levelTransfer:
		fmt.Printf("     twist: %s\n", m.cue.twist)
	}
	printPrimary(m.cue.fn, m.earned)
}

// printLevels is the honest scoreboard: what each function has actually earned,
// grouped by level rather than by weekday.
func printLevels(root string) {
	report := masteryReport(root)
	buckets := map[level][]mastery{}
	var unseen []mastery
	for _, m := range report {
		if !m.attempted() {
			unseen = append(unseen, m)
			continue
		}
		buckets[m.earned] = append(buckets[m.earned], m)
	}

	fmt.Println("UNDERSTANDING LEVELS")
	for _, l := range []level{levelTransfer, levelPattern, levelRecall} {
		group := buckets[l]
		fmt.Printf("\n%s (%d)\n", l.label(), len(group))
		if len(group) == 0 {
			fmt.Println("  —")
			continue
		}
		sort.Slice(group, func(i, j int) bool { return group[i].cue.drill < group[j].cue.drill })
		for _, m := range group {
			due := "  "
			if m.due() {
				due = "→ "
			}
			fmt.Printf("  %s%-22s %-24s %s\n", due, m.cue.fn, m.cue.drill, statusOf(m))
		}
	}

	fmt.Printf("\nnot started (%d)\n", len(unseen))
	byDrill := map[string][]string{}
	for _, m := range unseen {
		byDrill[m.cue.drill] = append(byDrill[m.cue.drill], m.cue.fn)
	}
	var files []string
	for f := range byDrill {
		files = append(files, f)
	}
	sort.Slice(files, func(i, j int) bool { return drillPriority[files[i]] < drillPriority[files[j]] })
	for _, f := range files {
		fmt.Printf("  %-24s %s\n", f, strings.Join(byDrill[f], ", "))
	}
	fmt.Println("\n  → marks a function that is due for review.")
	fmt.Println("  Session for today: go run . -- --refresh")
}
