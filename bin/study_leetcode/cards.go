package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// Solved-problem cards
//
// A problem solved once is mostly gone a few weeks later. --log keeps the part
// worth remembering (the pattern, the wording that signals it, the one-line
// insight) and schedules a re-solve. A re-solve starts from the statement:
// name the pattern first, then write the skeleton.

const (
	cardReviewCount = 4 // due cards shown with today's set
	gapMinCards     = 2 // a drill-less pattern on this many cards is a gap
)

// normalizeSlug accepts a slug or a problem URL.
func normalizeSlug(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if i := strings.Index(s, "/problems/"); i >= 0 {
		s = s[i+len("/problems/"):]
	}
	if i := strings.IndexAny(s, "/?#"); i >= 0 {
		s = s[:i]
	}
	return s
}

// titleFor finds a title for slug in today's set, else the existing card.
func titleFor(repoRoot, slug string, existing card) string {
	if f, ok := loadDailyFile(dailyJSONPath(repoRoot)); ok {
		for _, p := range f.Problems {
			if p.Slug == slug {
				return p.Title
			}
		}
	}
	if existing.Title != "" {
		return existing.Title
	}
	return slug
}

type prompter struct {
	sc  *bufio.Scanner
	out io.Writer
}

// ask prints q and returns the trimmed answer; def is used for an empty line.
func (p prompter) ask(q, def string) (string, error) {
	if def != "" {
		fmt.Fprintf(p.out, "  %s [Enter = %s]: ", q, def)
	} else {
		fmt.Fprintf(p.out, "  %s: ", q)
	}
	if !p.sc.Scan() {
		fmt.Fprintln(p.out)
		return "", fmt.Errorf("input ended; nothing saved")
	}
	a := strings.TrimSpace(p.sc.Text())
	if a == "" {
		return def, nil
	}
	return a, nil
}

// askPattern asks until it gets a pattern id. allowNone accepts "?" or an
// empty line as "no idea".
func (p prompter) askPattern(q, def string, allowNone bool) (string, error) {
	for tries := 0; tries < 3; tries++ {
		a, err := p.ask(q, def)
		if err != nil {
			return "", err
		}
		if allowNone && (a == "" || a == "?") {
			return "", nil
		}
		if id, ok := resolvePattern(a); ok {
			return id, nil
		}
		fmt.Fprintf(p.out, "    %q is not a pattern id (or matches several). Ids:\n", a)
		fmt.Fprintf(p.out, "    %s\n", strings.Join(patternIDs(), " "))
	}
	return "", fmt.Errorf("no pattern id given; nothing saved")
}

func (p prompter) askSolved() (string, error) {
	for tries := 0; tries < 3; tries++ {
		a, err := p.ask("Solved? [c]lean / [h]ints / [s]olution", "")
		if err != nil {
			return "", err
		}
		switch strings.ToLower(a) {
		case "c", "clean":
			return "clean", nil
		case "h", "hints", "hint":
			return "hints", nil
		case "s", "solution":
			return "solution", nil
		}
	}
	return "", fmt.Errorf("answer c, h or s; nothing saved")
}

// grade turns one review into a schedule move: recognised and solved clean
// moves the card up; one of the two keeps its box; neither, or reading the
// solution, sends it back to tomorrow.
func gradeReview(recognized bool, solved string) (passed, partial bool) {
	clean := solved == "clean"
	switch {
	case recognized && clean:
		return true, false
	case solved == "solution":
		return false, false
	case recognized || clean:
		return false, true
	}
	return false, false
}

// logCard records a new card or reviews an existing one, reading answers
// from in.
func logCard(repoRoot, rawSlug string, in io.Reader, out io.Writer, now time.Time) error {
	slug := normalizeSlug(rawSlug)
	if slug == "" {
		return fmt.Errorf("--log needs a problem slug, like two-sum")
	}
	f := loadCards(repoRoot)
	c, exists := f.Cards[slug]
	c.Slug = slug
	c.Title = titleFor(repoRoot, slug, c)

	p := prompter{sc: bufio.NewScanner(in), out: out}
	if exists {
		fmt.Fprintf(out, "Reviewing %s (%s) — %d earlier review(s)\n", c.Title, slug, c.Reviews)
	} else {
		fmt.Fprintf(out, "New card: %s (%s)\n", c.Title, slug)
	}

	guess, err := p.askPattern("Pattern you named BEFORE solving (id, or ? if none)", "", true)
	if err != nil {
		return err
	}
	def := c.Pattern
	if def == "" {
		def = guess
	}
	actual, err := p.askPattern("Actual pattern id", def, false)
	if err != nil {
		return err
	}
	signal, err := p.ask("Signal — the words in the statement that point to it", c.Signal)
	if err != nil {
		return err
	}
	insight, err := p.ask("Insight — the one-line trick", c.Insight)
	if err != nil {
		return err
	}
	solved, err := p.askSolved()
	if err != nil {
		return err
	}

	recognized := guess == actual
	passed, partial := gradeReview(recognized, solved)
	c.Pattern, c.Signal, c.Insight = actual, signal, insight
	c.Guess, c.Recognized, c.Solved = guess, recognized, solved
	if exists {
		c.Reviews++
	} else {
		c.Logged = dateOf(now)
	}
	if !recognized {
		c.Misses++
	}
	c.Reviewed = dateOf(now)
	c.Box, c.Due = nextReview(c.Box, passed, partial, now)
	f.Cards[slug] = c
	if err := saveCards(repoRoot, f); err != nil {
		return err
	}

	verdict := "✓ recognised"
	if !recognized {
		verdict = "✗ not recognised (guess: " + orNone(guess) + ")"
	}
	fmt.Fprintf(out, "\nSaved %s: %s · %s · %s, solved %s\n", slug, actual, verdict, describeBox(passed, partial), solved)
	fmt.Fprintf(out, "Next re-solve: %s (box %d)\n", c.Due, c.Box)
	if fns := drillFunctionsFor(actual); len(fns) > 0 {
		fmt.Fprintf(out, "Reflex drill for this pattern: %s\n", strings.Join(fns, ", "))
	} else {
		fmt.Fprintf(out, "No reflex drill covers %s yet — see: go run . -- --gaps\n", actual)
	}
	return nil
}

func describeBox(passed, partial bool) string {
	switch {
	case passed:
		return "moved up"
	case partial:
		return "same box"
	}
	return "back to tomorrow"
}

// drillFunctionsFor lists the reflex drill functions that practise pattern id.
func drillFunctionsFor(id string) []string {
	var fns []string
	for fn, p := range fnPattern {
		if p == id {
			fns = append(fns, fn)
		}
	}
	sort.Strings(fns)
	return fns
}

// printCardReview lists the cards due for a re-solve today.
func printCardReview(f cardFile, now time.Time, brief bool) {
	due := dueCards(f, now)
	if brief {
		if len(due) > 0 {
			var slugs []string
			for i, c := range due {
				if i == cardReviewCount {
					break
				}
				slugs = append(slugs, c.Slug)
			}
			fmt.Printf("           re-solve: %s (go run . -- --log=SLUG after each)\n", strings.Join(slugs, ", "))
		}
		return
	}
	if len(f.Cards) == 0 {
		fmt.Println("cards:   none yet — after solving one: go run . -- --log=SLUG")
		return
	}
	if len(due) == 0 {
		fmt.Printf("cards:   %d logged, none due today\n", len(f.Cards))
		return
	}
	fmt.Println()
	fmt.Println("━━━ RE-SOLVE — cards due today ━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Read the statement only. Name the pattern within 60s, write the skeleton,")
	fmt.Println("  then the code if you missed it. Record it: go run . -- --log=SLUG")
	for i, c := range due {
		if i == cardReviewCount {
			fmt.Printf("  … and %d more due\n", len(due)-cardReviewCount)
			break
		}
		fmt.Printf("  %d. %-40s due %s · box %d\n", i+1, c.Title, c.Due, c.Box)
		fmt.Printf("     %s\n", lcURL(c.Slug))
	}
}

// printGaps shows the patterns on your cards and which have no reflex drill.
// A drill-less pattern on gapMinCards or more cards is the next drill to write.
func printGaps(f cardFile) {
	if len(f.Cards) == 0 {
		fmt.Println("No cards yet. After solving a problem: go run . -- --log=SLUG")
		return
	}
	count := map[string]int{}
	misses := map[string]int{}
	for _, c := range f.Cards {
		count[c.Pattern]++
		misses[c.Pattern] += c.Misses
	}
	ids := make([]string, 0, len(count))
	for id := range count {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		if count[ids[i]] != count[ids[j]] {
			return count[ids[i]] > count[ids[j]]
		}
		return ids[i] < ids[j]
	})

	fmt.Println("Patterns on your cards:")
	var gaps []string
	for _, id := range ids {
		fns := drillFunctionsFor(id)
		cover := fmt.Sprintf("%d reflex function(s)", len(fns))
		if len(fns) == 0 {
			cover = "no reflex drill yet"
			if count[id] >= gapMinCards {
				gaps = append(gaps, id)
				cover += "  ← gap"
			}
		}
		fmt.Printf("  %-16s %2d card(s) · %2d miss(es) · %s\n", id, count[id], misses[id], cover)
	}
	fmt.Println()
	if len(gaps) == 0 {
		fmt.Printf("No gaps: every pattern on %d+ cards has a reflex drill.\n", gapMinCards)
		return
	}
	fmt.Printf("Gaps — write a reflex drill for: %s\n", strings.Join(gaps, ", "))
	fmt.Println("(a blank, a stub + main_test.go, a solution, then primaries.go, reflex_map.go")
	fmt.Println(" and the drills list in bin/study_play; see AGENTS.md)")
}
