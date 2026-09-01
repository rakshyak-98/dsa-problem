package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEveryCatalogFunctionHasACue(t *testing.T) {
	for _, entry := range essentialCatalog {
		for _, fn := range entry.fns {
			if _, ok := cueByFn[fn]; !ok {
				t.Errorf("%s (%s) has no pattern cue", fn, entry.group)
			}
		}
	}
}

func TestEveryCueIsComplete(t *testing.T) {
	for _, c := range cues {
		if c.ask == "" || c.cue == "" || c.move == "" || c.twist == "" {
			t.Errorf("%s: incomplete cue", c.fn)
		}
		if c.tier < levelRecall || c.tier > levelTransfer {
			t.Errorf("%s: tier %d out of range", c.fn, c.tier)
		}
		if _, ok := drillPriority[c.drill]; !ok {
			t.Errorf("%s: drill %q is not in drillPriority", c.fn, c.drill)
		}
	}
}

func TestEveryCueHasAPrimaryProblem(t *testing.T) {
	for _, c := range cues {
		p, ok := primaries[c.fn]
		if !ok {
			t.Errorf("%s: no primary problem", c.fn)
			continue
		}
		if p.title == "" || p.slug == "" || p.nextUp == "" {
			t.Errorf("%s: incomplete primary", c.fn)
		}
		switch p.diff {
		case "Easy", "Medium", "Hard":
		default:
			t.Errorf("%s: bad difficulty %q", c.fn, p.diff)
		}
	}
	for fn := range primaries {
		if _, ok := cueByFn[fn]; !ok {
			t.Errorf("primary %q has no drill function", fn)
		}
	}
}

// A cue only makes sense as a recognition question if it never names the
// function it is asking about.
func TestCuesDoNotLeakTheirAnswer(t *testing.T) {
	for _, c := range cues {
		if containsFold(c.cue, c.fn) {
			t.Errorf("%s: cue names the function, so it cannot be a quiz", c.fn)
		}
	}
}

func containsFold(haystack, needle string) bool {
	lower := func(s string) string {
		b := []byte(s)
		for i := range b {
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] += 'a' - 'A'
			}
		}
		return string(b)
	}
	h, n := lower(haystack), lower(needle)
	for i := 0; i+len(n) <= len(h); i++ {
		if h[i:i+len(n)] == n {
			return true
		}
	}
	return false
}

func TestOwningFunction(t *testing.T) {
	cases := []struct {
		in   string
		want string
		ok   bool
	}{
		{"rotateRight k=1", "rotateRight", true},
		{"rotateRight", "rotateRight", true},
		{"maxSumSubarrayK window slide", "maxSumSubarrayK", true},
		{"removeDuplicates empty", "removeDuplicates", true},
		{"TestRotateRight (0.00s)", "", false},
		{"TestAll/maxSumSubarrayK (0.00s)", "", false},
		{"rotateRightward", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		got, ok := owningFunction(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("owningFunction(%q) = (%q, %v), want (%q, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestEarnedLevelIsCappedByTier(t *testing.T) {
	// arraySum is a tier-1 function: passing it fifty times still does not make
	// it a transfer question.
	easy := cueByFn["arraySum"]
	if got := earnedLevel(easy, 50, 0, true); got != levelRecall {
		t.Errorf("arraySum earned %s, want L1 (its tier)", got.label())
	}
	hard := cueByFn["groupAnagrams"]
	if got := earnedLevel(hard, 10, 0, true); got != levelTransfer {
		t.Errorf("groupAnagrams earned %s, want L3", got.label())
	}
	if got := earnedLevel(hard, 3, 1, true); got != levelPattern {
		t.Errorf("groupAnagrams with a fail earned %s, want L2", got.label())
	}
	if got := earnedLevel(hard, 0, 4, false); got != levelRecall {
		t.Errorf("never passed earned %s, want L1", got.label())
	}
	// A pass from long ago does not hold a transfer grade.
	if got := earnedLevel(hard, 10, 0, false); got == levelTransfer {
		t.Error("a stale pass should not still count as L3")
	}
}

func TestDueRespectsReviewGap(t *testing.T) {
	fresh := mastery{cue: cueByFn["twoSum"], passes: 3, daysIdle: 1, earned: levelPattern}
	if fresh.due() {
		t.Error("an L2 function passed yesterday is not due yet")
	}
	stale := mastery{cue: cueByFn["twoSum"], passes: 3, daysIdle: 9, earned: levelPattern}
	if !stale.due() {
		t.Error("an L2 function idle 9 days is due")
	}
	never := mastery{cue: cueByFn["twoSum"], daysIdle: -1}
	if !never.due() {
		t.Error("a never-attempted function is always due")
	}
}

func TestMasteryReportCoversTheWholeCatalog(t *testing.T) {
	root := t.TempDir()
	report := masteryReport(root)
	if len(report) != len(cues) {
		t.Fatalf("report has %d functions, catalog has %d", len(report), len(cues))
	}
	for _, m := range report {
		if m.attempted() {
			t.Errorf("%s: attempted with an empty log", m.cue.fn)
		}
		if m.earned != levelRecall {
			t.Errorf("%s: earned %s with an empty log", m.cue.fn, m.earned.label())
		}
	}
}

func TestBuildRefreshOnAnEmptyLog(t *testing.T) {
	root := t.TempDir()
	s := buildRefresh(root, drills[0], time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC))
	if len(s.recognise) != recogniseCount {
		t.Errorf("recognise round has %d cues, want %d", len(s.recognise), recogniseCount)
	}
	if len(s.rebuild) != 0 {
		t.Error("nothing can be weak when nothing has been attempted")
	}
	if len(s.newGround) != newGroundCount {
		t.Errorf("new ground has %d, want %d", len(s.newGround), newGroundCount)
	}
	// Hashing is the highest-priority untouched family.
	for _, m := range s.newGround {
		if m.cue.drill != "02_hashing_reflex" {
			t.Errorf("new ground picked %s from %s, want 02_hashing_reflex first", m.cue.fn, m.cue.drill)
		}
	}
}

func TestBuildRefreshIsStableWithinADayAndRotatesAcrossDays(t *testing.T) {
	root := writeTestLog(t, map[string]fnRecord{
		"twoSum":       {Passes: 4, LastPass: "2026-08-30"},
		"binarySearch": {Passes: 4, LastPass: "2026-08-30"},
		"maxArea":      {Passes: 4, LastPass: "2026-08-30"},
		"fib":          {Passes: 4, LastPass: "2026-08-30"},
		"numIslands":   {Passes: 4, LastPass: "2026-08-30"},
		"rob":          {Passes: 4, LastPass: "2026-08-30"},
	})
	day1 := time.Date(2026, 9, 1, 8, 0, 0, 0, time.UTC)
	day1Again := time.Date(2026, 9, 1, 21, 0, 0, 0, time.UTC)
	day2 := time.Date(2026, 9, 2, 8, 0, 0, 0, time.UTC)

	a := recogniseNames(buildRefresh(root, drills[0], day1))
	b := recogniseNames(buildRefresh(root, drills[0], day1Again))
	c := recogniseNames(buildRefresh(root, drills[0], day2))

	if a != b {
		t.Errorf("same day gave different rounds: %q vs %q", a, b)
	}
	if a == c {
		t.Errorf("consecutive days gave the identical round: %q", a)
	}
}

func TestRefreshRebuildsTheWeakestFirst(t *testing.T) {
	root := writeTestLog(t, map[string]fnRecord{
		"maxArea":      {Passes: 20, Fails: 1, LastPass: "2026-08-30"},
		"moveZeroes":   {Passes: 1, Fails: 14, LastPass: "2026-08-10"},
		"isPalindrome": {Passes: 5, Fails: 3, LastPass: "2026-08-20"},
	})
	s := buildRefresh(root, drills[0], time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC))
	if len(s.rebuild) == 0 {
		t.Fatal("nothing scheduled for rebuild")
	}
	if s.rebuild[0].cue.fn != "moveZeroes" {
		t.Errorf("rebuild starts with %s, want moveZeroes (worst fail rate)", s.rebuild[0].cue.fn)
	}
}

func recogniseNames(s refreshSession) string {
	out := ""
	for _, m := range s.recognise {
		out += m.cue.fn + " "
	}
	return out
}

func writeTestLog(t *testing.T, fns map[string]fnRecord) string {
	t.Helper()
	root := t.TempDir()
	if err := saveLog(root, drillLog{Functions: fns}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".drill_log.json")); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestPrintersDoNotPanic(t *testing.T) {
	root := writeTestLog(t, map[string]fnRecord{"twoSum": {Passes: 9, LastPass: today()}})
	mustPrint(t, "printLevels", func() { printLevels(root) })
	mustPrint(t, "printProblemSet", func() { printProblemSet(root) })
	mustPrint(t, "printAllTriggers", printAllTriggers)
	mustPrint(t, "printTriggers", func() { printTriggers(drills[2]) })
	s := buildRefresh(root, drills[2], time.Now())
	mustPrint(t, "printRefresh hidden", func() { printRefresh(s, false) })
	mustPrint(t, "printRefresh shown", func() { printRefresh(s, true) })
	mustPrint(t, "printProblemMap", func() { printProblemMap("03_two_pointers_reflex") })
	mustPrint(t, "printCore5Problems", printCore5Problems)
}

func mustPrint(t *testing.T, name string, fn func()) {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Errorf("%s printed nothing", name)
	}
}
