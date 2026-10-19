package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNormalizeSlug(t *testing.T) {
	for in, want := range map[string]string{
		"two-sum":                                "two-sum",
		" Two-Sum ":                              "two-sum",
		"https://leetcode.com/problems/two-sum/": "two-sum",
		"leetcode.com/problems/3sum/description/?envType=x": "3sum",
	} {
		if got := normalizeSlug(in); got != want {
			t.Errorf("normalizeSlug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestGradeReview(t *testing.T) {
	cases := []struct {
		recognized      bool
		solved          string
		passed, partial bool
	}{
		{true, "clean", true, false},
		{true, "hints", false, true},
		{false, "clean", false, true},
		{true, "solution", false, false},
		{false, "hints", false, false},
	}
	for _, c := range cases {
		if p, pa := gradeReview(c.recognized, c.solved); p != c.passed || pa != c.partial {
			t.Errorf("gradeReview(%v, %s) = %v,%v", c.recognized, c.solved, p, pa)
		}
	}
}

func TestLogCardNewThenReview(t *testing.T) {
	root := t.TempDir()
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	var out bytes.Buffer

	// Guessed hash-lookup, it was heap; solved with hints.
	in := strings.NewReader("hash\nheap\nk most frequent\ncount, then a size-k min-heap\nh\n")
	if err := logCard(root, "https://leetcode.com/problems/top-k-frequent-elements/", in, &out, now); err != nil {
		t.Fatal(err)
	}
	c := loadCards(root).Cards["top-k-frequent-elements"]
	if c.Pattern != "heap" || c.Guess != "hash-lookup" || c.Recognized || c.Misses != 1 || c.Box != 0 || c.Due != "2026-10-20" || c.Logged != "2026-10-19" {
		t.Fatalf("new card = %+v", c)
	}
	if !strings.Contains(out.String(), "No reflex drill covers heap") {
		t.Errorf("expected a gap note:\n%s", out.String())
	}

	// Review tomorrow: recognised, clean; Enter keeps the stored answers.
	out.Reset()
	in = strings.NewReader("heap\n\n\n\nc\n")
	if err := logCard(root, "top-k-frequent-elements", in, &out, now.AddDate(0, 0, 1)); err != nil {
		t.Fatal(err)
	}
	c = loadCards(root).Cards["top-k-frequent-elements"]
	if !c.Recognized || c.Box != 1 || c.Due != "2026-10-23" || c.Reviews != 1 || c.Signal != "k most frequent" || c.Misses != 1 {
		t.Fatalf("reviewed card = %+v", c)
	}
}

func TestLogCardStopsOnEOF(t *testing.T) {
	root := t.TempDir()
	err := logCard(root, "two-sum", strings.NewReader("hash\n"), io.Discard, time.Now())
	if err == nil || !strings.Contains(err.Error(), "nothing saved") {
		t.Fatalf("want an input-ended error, got %v", err)
	}
	if len(loadCards(root).Cards) != 0 {
		t.Fatal("a half-answered card must not be saved")
	}
}

func TestPrintGapsAndReview(t *testing.T) {
	f := cardFile{Cards: map[string]card{
		"a": {Slug: "a", Title: "A", Pattern: "heap", Due: "2026-10-18"},
		"b": {Slug: "b", Title: "B", Pattern: "heap", Due: "2026-10-30"},
		"c": {Slug: "c", Title: "C", Pattern: "dp", Due: "2026-10-19"},
	}}
	out := captureStdout(t, func() { printGaps(f) })
	if !strings.Contains(out, "Gaps — write a reflex drill for: heap") || strings.Contains(out, "for: heap, dp") {
		t.Fatalf("gaps:\n%s", out)
	}
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	out = captureStdout(t, func() { printCardReview(f, now, false) })
	if !strings.Contains(out, "1. A") || !strings.Contains(out, "2. C") || strings.Contains(out, "B ") {
		t.Fatalf("review:\n%s", out)
	}
}

func TestParseLeetcodeCardFlags(t *testing.T) {
	opts, ctl, perr := parseLeetcode([]string{"--log=two-sum", "--answers"})
	if ctl != gnuOK || perr || opts.log != "two-sum" || !opts.answers {
		t.Fatalf("parse: %+v", opts)
	}
	if opts, _, _ := parseLeetcode([]string{"--gaps"}); !opts.gaps {
		t.Fatal("--gaps")
	}
	if _, _, perr := parseLeetcode([]string{"--log"}); !perr {
		t.Fatal("bare --log should be a usage error")
	}
}

func TestMarkdownHidesTagsUntilAnswers(t *testing.T) {
	set := practiceSet{day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		problems: []lcProblem{{Num: 704, Title: "Binary Search", Slug: "binary-search", Diff: "Easy", Pattern: "array, binary search", ReflexFn: "binarySearch"}}}
	ans := renderAnswersMarkdown(set)
	if !strings.Contains(ans, "array, binary search") || !strings.Contains(ans, "binarySearch") {
		t.Fatalf("answers:\n%s", ans)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}
