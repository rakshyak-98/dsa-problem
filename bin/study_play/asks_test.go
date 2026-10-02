package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDailyAsksComplete(t *testing.T) {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, day := range days {
		ask, ok := dailyAsks[day]
		if !ok {
			t.Fatalf("missing ask for %s", day)
		}
		if ask.statement == "" || len(ask.hints) < 4 {
			t.Fatalf("incomplete ask for %s", day)
		}
	}
}

func TestPrintAskWarmup(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printAskWarmup("Monday")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 30 {
		t.Fatal("printAskWarmup empty")
	}
}

func TestSkillBankIsSound(t *testing.T) {
	known := map[string]bool{}
	for _, id := range skillOrder {
		known[id] = true
		if n := len(skillQuestions(id)); n < 4 {
			t.Errorf("skill %s has %d questions, want at least 4", id, n)
		}
	}
	seen := map[string]bool{}
	for _, q := range skillBank {
		if !known[q.skill] || q.q == "" || q.answer == "" {
			t.Errorf("bad skill question %+v", q)
		}
		if seen[q.q] {
			t.Errorf("duplicate question %q", q.q)
		}
		seen[q.q] = true
	}
	for _, d := range drills {
		if len(skillFocus[d.day]) < skillFocusSkills {
			t.Errorf("%s: too few focus skills", d.day)
		}
		for _, id := range skillFocus[d.day] {
			if !known[id] {
				t.Errorf("%s: unknown focus skill %q", d.day, id)
			}
		}
	}
}

func TestSkillSetShapeAndStability(t *testing.T) {
	now := time.Date(2026, 10, 1, 9, 0, 0, 0, time.UTC)
	f, r, c := skillSet("Thursday", now, nil)
	f2, _, c2 := skillSet("Thursday", now.Add(8*time.Hour), nil)
	if len(f) != skillFocusSkills*skillFocusPerSkill || len(r) != 0 || len(c) != skillCrossCount {
		t.Fatalf("shape: %d focus, %d redo, %d cross", len(f), len(r), len(c))
	}
	for i := range f {
		if f[i] != f2[i] {
			t.Fatal("focus changed within a day")
		}
	}
	for i := range c {
		if c[i] != c2[i] {
			t.Fatal("cross changed within a day")
		}
	}
	inFocus := map[string]bool{}
	for _, q := range f {
		inFocus[q.skill] = true
	}
	crossSeen := map[string]bool{}
	for _, q := range c {
		if inFocus[q.skill] || crossSeen[q.skill] {
			t.Errorf("cross-topic repeats skill %s", q.skill)
		}
		crossSeen[q.skill] = true
	}
}

// The rotation must reach every skill and every question; otherwise some
// topics would silently never be practised.
func TestSkillRotationCoversWholeBank(t *testing.T) {
	skills := map[string]bool{}
	qs := map[string]bool{}
	start := time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)
	for day := 0; day < 700; day++ {
		now := start.AddDate(0, 0, day)
		f, r, c := skillSet(drills[(int(now.Weekday())+6)%7].day, now, nil)
		for _, q := range append(append(f, r...), c...) {
			skills[q.skill] = true
			qs[q.q] = true
		}
	}
	for _, id := range skillOrder {
		if !skills[id] {
			t.Errorf("skill %s never scheduled", id)
		}
	}
	if len(qs) != len(skillBank) {
		t.Errorf("only %d of %d questions ever scheduled", len(qs), len(skillBank))
	}
}

func TestPrintSkillQuestionsRevealsOnlyWithShow(t *testing.T) {
	now := time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC)
	capture := func(show bool) string {
		revealSkillAnswers = show
		defer func() { revealSkillAnswers = false }()
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		printSkillQuestionsAt("Monday", now, nil)
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		return buf.String()
	}
	if hidden := capture(false); strings.Contains(hidden, "answer:") {
		t.Fatal("answers shown without --show")
	}
	if shown := capture(true); strings.Count(shown, "answer:") != skillFocusSkills*skillFocusPerSkill+skillCrossCount {
		t.Fatalf("--show should reveal all answers:\n%s", shown)
	}
}

func TestWeakSkillsComeBackAsRedo(t *testing.T) {
	now := time.Date(2026, 10, 1, 9, 0, 0, 0, time.UTC)
	f, r, c := skillSet("Thursday", now, []string{"heap", "dp", "graph"})
	if len(r) != skillMaxRedo {
		t.Fatalf("want %d redo questions, got %d", skillMaxRedo, len(r))
	}
	if len(f)+len(r)+len(c) != skillFocusSkills*skillFocusPerSkill+skillCrossCount {
		t.Fatalf("total changed: %d/%d/%d", len(f), len(r), len(c))
	}
	if r[0].skill != "heap" || r[1].skill != "dp" {
		t.Fatalf("redo should follow the weak order: %+v", r)
	}
	for _, q := range c {
		if q.skill == "heap" || q.skill == "dp" {
			t.Fatalf("cross repeats a redo skill: %s", q.skill)
		}
	}
}

func TestRecordAndWeakSkills(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	if err := recordSkills(root, []string{"heap", "index"}, true); err != nil {
		t.Fatal(err)
	}
	if err := recordSkills(root, []string{"heap"}, true); err != nil {
		t.Fatal(err)
	}
	got := weakSkills(loadLog(root), now)
	if len(got) != 2 || got[0] != "heap" || got[1] != "index" {
		t.Fatalf("most-missed first, got %v", got)
	}
	if err := recordSkills(root, []string{"index"}, false); err != nil {
		t.Fatal(err)
	}
	if got := weakSkills(loadLog(root), now); len(got) != 1 || got[0] != "heap" {
		t.Fatalf("recovered skill should drop out, got %v", got)
	}
	if err := recordSkills(root, []string{"nope"}, true); err == nil {
		t.Fatal("unknown skill should be rejected")
	}
	// A stale miss ages out of the window.
	if got := weakSkills(loadLog(root), now.AddDate(0, 0, skillMissWindow+2)); len(got) != 0 {
		t.Fatalf("stale misses should expire, got %v", got)
	}
	// Function history in the same log is untouched.
	if loadLog(root).Functions == nil {
		t.Fatal("Functions map lost")
	}
}

func TestParsePlaySkillFlags(t *testing.T) {
	opts, _, perr := parsePlay([]string{"--missed=index,range", "--got", "heap"})
	if perr || opts.missed != "index,range" || opts.got != "heap" {
		t.Fatalf("missed/got: %+v perr=%v", opts, perr)
	}
	if _, _, perr := parsePlay([]string{"--missed"}); !perr {
		t.Fatal("bare --missed should be a usage error")
	}
}
