package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestApplyRatingLadder(t *testing.T) {
	on := time.Date(2026, 8, 7, 0, 0, 0, 0, time.Local)
	p := &Progress{Ease: 2.5}

	applyRating(p, 3, on) // good → +1
	if p.Interval != 1 || p.Due != "2026-08-08" {
		t.Fatalf("first good: got interval=%d due=%s", p.Interval, p.Due)
	}
	applyRating(p, 3, on) // → +3
	if p.Interval != 3 {
		t.Fatalf("second good: got interval=%d", p.Interval)
	}
	applyRating(p, 3, on) // → +7
	if p.Interval != 7 {
		t.Fatalf("third good: got interval=%d", p.Interval)
	}
	applyRating(p, 3, on) // → +21
	if p.Interval != 21 {
		t.Fatalf("fourth good: got interval=%d", p.Interval)
	}
}

func TestAgainResets(t *testing.T) {
	on := today()
	p := &Progress{Ease: 2.5, Reps: 3, Interval: 7}
	applyRating(p, 1, on)
	if p.Reps != 0 || p.Interval != 1 || p.Lapses != 1 {
		t.Fatalf("again: reps=%d interval=%d lapses=%d", p.Reps, p.Interval, p.Lapses)
	}
}

func TestIsDue(t *testing.T) {
	on := time.Date(2026, 8, 7, 0, 0, 0, 0, time.Local)
	p := &Progress{Due: "2026-08-07", Reviews: 1, Reps: 1}
	if !p.isDue(on) {
		t.Fatal("expected due today")
	}
	p.Due = "2026-08-08"
	if p.isDue(on) {
		t.Fatal("expected not due")
	}
	if (&Progress{}).isDue(on) {
		t.Fatal("new card should not count as due")
	}
}

func TestCardMatches(t *testing.T) {
	c := Card{Deck: "jargon", Tags: []string{"dsa", "jargon"}, Section: "Graphs"}
	if !cardMatches(c, "jargon", "") {
		t.Fatal("deck filter")
	}
	if cardMatches(c, "math", "") {
		t.Fatal("wrong deck")
	}
	if !cardMatches(c, "", "dsa") {
		t.Fatal("tag filter")
	}
	if cardMatches(c, "", "backend") {
		t.Fatal("wrong tag")
	}
}

func TestParseArgsGNU(t *testing.T) {
	opts := parseArgs([]string{"--", "--stats", "--deck=math", "--limit=5", "--new=2"})
	if opts.parseErr {
		t.Fatal("unexpected parse error")
	}
	if !opts.stats || opts.deck != "math" || opts.limit != 5 || opts.newN != 2 {
		t.Fatalf("%+v", opts)
	}

	opts = parseArgs([]string{"--due", "-d", "jargon", "-t", "dsa", "-n", "10"})
	if opts.parseErr || !opts.due || opts.deck != "jargon" || opts.tag != "dsa" || opts.limit != 10 {
		t.Fatalf("%+v", opts)
	}

	opts = parseArgs([]string{"--catalog"})
	if opts.parseErr || !opts.catalog || opts.review {
		t.Fatalf("%+v", opts)
	}

	opts = parseArgs([]string{"--deck=nope"})
	if !opts.parseErr {
		t.Fatal("expected unknown deck error")
	}

	opts = parseArgs([]string{"--limit"})
	if !opts.parseErr {
		t.Fatal("expected missing arg error")
	}

	opts = parseArgs([]string{"--bogus"})
	if !opts.parseErr {
		t.Fatal("expected unknown option")
	}

	// bare subcommands are rejected (GNU: options only)
	opts = parseArgs([]string{"stats"})
	if !opts.parseErr {
		t.Fatal("expected bare subcommand rejected")
	}
}

func TestPrintHelpGNU(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printHelp()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	for _, want := range []string{
		"Usage:",
		"Options:",
		"-h, --help",
		"--deck=NAME",
		"--tag=TAG",
		"--limit=N",
		"--new=N",
		"--catalog",
		"--due",
		"--stats",
		"(default:",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("help missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "Examples:") {
		t.Fatal("help should not include examples section")
	}
}
