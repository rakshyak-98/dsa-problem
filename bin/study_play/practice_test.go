package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestEveryCatalogFunctionHasPatternAndHint(t *testing.T) {
	for _, d := range drills {
		for _, fn := range d.functions {
			id, ok := fnPattern[fn]
			if !ok {
				t.Errorf("%s has no fnPattern entry", fn)
				continue
			}
			if _, ok := patternByID(id); !ok {
				t.Errorf("%s: unknown pattern %q", fn, id)
			}
			if fnHints[fn] == "" {
				t.Errorf("%s has no hint", fn)
			}
		}
	}
	seen := map[string]bool{}
	for _, p := range patterns {
		if seen[p.id] {
			t.Errorf("duplicate pattern id %q", p.id)
		}
		seen[p.id] = true
	}
}

func TestRecognitionBankIsSound(t *testing.T) {
	ids := map[string]bool{}
	for _, q := range recogBank {
		if ids[q.id] {
			t.Errorf("duplicate question id %q", q.id)
		}
		ids[q.id] = true
		for _, id := range append([]string{q.pattern, q.lookalike}, q.alts...) {
			if _, ok := patternByID(id); !ok {
				t.Errorf("%s: unknown pattern %q", q.id, id)
			}
		}
		if q.lookalike == q.pattern {
			t.Errorf("%s: look-alike equals the answer", q.id)
		}
		if q.why == "" || q.notIt == "" {
			t.Errorf("%s: missing signal or look-alike reason", q.id)
		}
	}
}

func TestResolvePattern(t *testing.T) {
	cases := map[string]string{"variable-window": "variable-window", "var": "variable-window", " DIJ ": "dijkstra", "mono": "mono-stack"}
	for in, want := range cases {
		if got, ok := resolvePattern(in); !ok || got != want {
			t.Errorf("resolvePattern(%q) = %q,%v want %q", in, got, ok, want)
		}
	}
	for _, in := range []string{"", "nope", "t"} { // "t" matches several ids
		if _, ok := resolvePattern(in); ok {
			t.Errorf("resolvePattern(%q) should fail", in)
		}
	}
}

func TestNextReview(t *testing.T) {
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	if box, due := nextReview(0, true, false, now); box != 1 || due != "2026-10-22" {
		t.Errorf("pass from 0: box %d due %s", box, due)
	}
	if box, due := nextReview(3, false, false, now); box != 0 || due != "2026-10-20" {
		t.Errorf("miss: box %d due %s", box, due)
	}
	if box, _ := nextReview(2, false, true, now); box != 2 {
		t.Errorf("partial keeps the box, got %d", box)
	}
	if box, _ := nextReview(len(reviewIntervals)-1, true, false, now); box != len(reviewIntervals)-1 {
		t.Errorf("box should cap, got %d", box)
	}
}

// A drill with one passing, one blank, one wrong and one never-ending function.
const gradeFixtureMain = `package main

func good() int  { return 1 }
func blank() int { panic("Implement from memory") }
func wrong() int { return 2 }
func spin() []int {
	var s []int
	for {
		s = append(s, 1)
		if len(s) < 0 {
			return s
		}
	}
}

func main() {}
`

const gradeFixtureTest = `package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func TestGood(t *testing.T)  { assert(t, "good basic", good() == 1) }
func TestBlank(t *testing.T) { assert(t, "blank basic", blank() == 1) }
func TestWrong(t *testing.T) { assert(t, "wrong first", true); assert(t, "wrong value", wrong() == 1) }
func TestSpin(t *testing.T)  { assert(t, "spin basic", len(spin()) == 0) }
`

func TestGradeFunctionsIsolatesEachTest(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"go.mod":       "module fixture\n\ngo 1.22\n",
		"main.go":      gradeFixtureMain,
		"main_test.go": gradeFixtureTest,
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	results, err := gradeFunctions(dir, "fixture", []string{"good", "blank", "wrong", "spin", "missing"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]fnStatus{"good": statusPass, "blank": statusTodo, "wrong": statusFail, "spin": statusTimeout, "missing": statusNoTest}
	for _, r := range results {
		if r.status != want[r.fn] {
			t.Errorf("%s: status %v, want %v (detail %q)", r.fn, r.status, want[r.fn], r.detail)
		}
	}
	if results[2].detail != "wrong value" {
		t.Errorf("wrong: detail %q, want the failing assert", results[2].detail)
	}
}

func TestGradeFunctionsReportsBuildErrors(t *testing.T) {
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module fixture\n\ngo 1.22\n"), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() { x }\n"), 0o644)
	if _, err := gradeFunctions(dir, "fixture", []string{"x"}); err == nil || !strings.Contains(err.Error(), "does not compile") {
		t.Fatalf("want a compile error, got %v", err)
	}
}

func TestRecordGradesScoresOnlyTheFirstRunOfADay(t *testing.T) {
	log := drillLog{Functions: map[string]fnRecord{}}
	morning := time.Date(2026, 10, 19, 9, 0, 0, 0, time.Local)
	log.Functions["twoSum"] = fnRecord{Started: morning.Add(-5 * time.Minute).Format(time.RFC3339)}

	recordGrades(&log, []fnResult{{fn: "twoSum", status: statusFail}, {fn: "rob", status: statusTodo}}, morning)
	rec := log.Functions["twoSum"]
	if rec.Fails != 1 || rec.Box != 0 || rec.Due != "2026-10-20" || len(rec.Attempts) != 1 || rec.Attempts[0].FirstTry {
		t.Fatalf("after first-try miss: %+v", rec)
	}
	if _, ok := log.Functions["rob"]; ok {
		t.Fatal("a todo function should not be recorded")
	}

	later := morning.Add(10 * time.Minute)
	recordGrades(&log, []fnResult{{fn: "twoSum", status: statusPass}}, later)
	rec = log.Functions["twoSum"]
	a := rec.Attempts[0]
	if rec.Passes != 0 || rec.Due != "2026-10-20" || a.Runs != 2 || !a.Passed || a.SecsToPass != 15*60 {
		t.Fatalf("a same-day fix must not reschedule or count as a pass: %+v %+v", rec, a)
	}

	tomorrow := morning.AddDate(0, 0, 1)
	recordGrades(&log, []fnResult{{fn: "twoSum", status: statusPass}}, tomorrow)
	rec = log.Functions["twoSum"]
	if rec.Passes != 1 || rec.Box != 1 || rec.Due != "2026-10-23" || !rec.Attempts[1].FirstTry {
		t.Fatalf("next-day first-try pass: %+v", rec)
	}
}

func TestDueFunctions(t *testing.T) {
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	log := drillLog{Functions: map[string]fnRecord{}}
	for _, d := range drills {
		for _, fn := range d.functions {
			log.Functions[fn] = fnRecord{Passes: 1, Box: 3, Due: "2026-12-01"} // not due
		}
	}
	log.Functions["floodFill"] = fnRecord{Fails: 3, Due: "2026-10-10"}
	log.Functions["numIslands"] = fnRecord{Fails: 1, Due: "2026-10-18"}
	log.Functions["dfs"] = fnRecord{Fails: 1, Due: "2026-10-19"}
	log.Functions["twoSum"] = fnRecord{Passes: 2, Box: 1, Due: "2026-10-19"}
	log.Functions["rob"] = fnRecord{Fails: 1, Due: "2026-10-01"} // in the skipped file

	got := dueFunctions(log, now, 4, "06_dp_reflex")
	var names []string
	for _, it := range got {
		names = append(names, it.fn)
	}
	// One per file first (most overdue first), then a second from a file;
	// at most two from the graphs file, nothing from the skipped file,
	// nothing not yet due.
	if strings.Join(names, ",") != "floodFill,twoSum,numIslands" {
		t.Fatalf("due = %v", names)
	}
	if got[0].overdue != 9 {
		t.Errorf("floodFill overdue = %d, want 9", got[0].overdue)
	}
}

func TestBlankInFileKeepsTheOtherFunctions(t *testing.T) {
	repo := t.TempDir()
	if err := setupAllDrills(repo); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(writeReflexDir(repo, "06_dp_reflex"), "main.go")
	src, _ := os.ReadFile(path)
	written := strings.Replace(string(src), "func climbStairs(n int) int {\n\tpanic(\"Implement from memory\")",
		"func climbStairs(n int) int {\n\treturn 42 // mine", 1)
	written = strings.Replace(written, "func rob(nums []int) int {\n\tpanic(\"Implement from memory\")",
		"func rob(nums []int) int {\n\treturn 7 // mine", 1)
	if written == string(src) {
		t.Fatal("fixture replace did not apply; blank stub format changed?")
	}
	_ = os.WriteFile(path, []byte(written), 0o644)

	if err := blankFunctions(repo, []string{"rob"}, time.Now()); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(path)
	if !strings.Contains(string(out), "return 42 // mine") {
		t.Error("climbStairs was not left alone")
	}
	if strings.Contains(string(out), "return 7 // mine") {
		t.Error("rob was not blanked")
	}
	if loadLog(repo).Functions["rob"].Started == "" {
		t.Error("blanking should stamp the start time")
	}
}

func TestParseFnList(t *testing.T) {
	fns, err := parseFnList("twoSum, rob,twoSum")
	if err != nil || strings.Join(fns, ",") != "twoSum,rob" {
		t.Fatalf("parseFnList = %v, %v", fns, err)
	}
	if _, err := parseFnList("twosum"); err == nil || !strings.Contains(err.Error(), "twoSum") {
		t.Fatalf("want a suggestion for twosum, got %v", err)
	}
}

func TestRunRecognition(t *testing.T) {
	qs := []recogQ{recogBank[0], recogBank[1], recogBank[2]} // hash-lookup, two-pointers, prefix-sum
	in := strings.NewReader("hash\nzzz\ntwo-p\n?\n")
	var out bytes.Buffer
	right, wrong := runRecognition(qs, in, &out)
	if strings.Join(right, ",") != "two-sum,two-sum-sorted" || strings.Join(wrong, ",") != "subarray-sum-k" {
		t.Fatalf("right=%v wrong=%v\n%s", right, wrong, out.String())
	}
	if !strings.Contains(out.String(), "not a pattern id") || !strings.Contains(out.String(), "not variable-window") {
		t.Errorf("missing retry prompt or look-alike:\n%s", out.String())
	}

	right, wrong = runRecognition(qs, strings.NewReader("hash\n"), &out)
	if len(right) != 1 || len(wrong) != 0 {
		t.Fatalf("EOF should stop grading: right=%v wrong=%v", right, wrong)
	}
}

func TestPickRecognition(t *testing.T) {
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	misses := map[string]skillRecord{
		"koko":     {Misses: 2, LastMiss: "2026-10-18"},
		"rpn":      {Misses: 1, LastMiss: "2026-10-17"},
		"subsets":  {Misses: 3, LastMiss: "2026-09-01"}, // too old
		"vanished": {Misses: 5, LastMiss: "2026-10-18"}, // not in the bank
	}
	qs := pickRecognition(recogBank, misses, now, recognizeCount)
	if len(qs) != recognizeCount || qs[0].id != "koko" || qs[1].id != "rpn" {
		t.Fatalf("picked %v", qs)
	}
	seen := map[string]bool{}
	for _, q := range qs {
		if seen[q.id] {
			t.Fatalf("duplicate pick %s", q.id)
		}
		seen[q.id] = true
	}
	if again := pickRecognition(recogBank, nil, now, recognizeCount); again[0].id != pickRecognition(recogBank, nil, now, recognizeCount)[0].id {
		t.Fatal("a day's round should be stable")
	}
	if got := pickRecognition(nil, nil, now, 3); len(got) != 0 {
		t.Fatal("empty bank")
	}
}

func TestRecordRecognition(t *testing.T) {
	log := drillLog{Functions: map[string]fnRecord{}}
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	recordRecognition(&log, nil, []string{"koko", "koko"}, now)
	recordRecognition(&log, []string{"koko"}, nil, now)
	if log.Recognition["koko"].Misses != 1 {
		t.Fatalf("koko = %+v", log.Recognition["koko"])
	}
	recordRecognition(&log, []string{"koko"}, nil, now)
	if _, ok := log.Recognition["koko"]; ok {
		t.Fatal("a recovered question should drop out")
	}
}

func TestPruneRetiredBacksUpFirst(t *testing.T) {
	root := t.TempDir()
	log := drillLog{Functions: map[string]fnRecord{
		"twoSum": {Passes: 1}, "twoSum basic": {Passes: 1}, "arraySum": {Fails: 9}, "arraySum basic": {Passes: 3},
	}}
	if err := saveLog(root, log); err != nil {
		t.Fatal(err)
	}
	pruneRetired(root, &log)
	for _, keep := range []string{"twoSum", "twoSum basic"} {
		if _, ok := log.Functions[keep]; !ok {
			t.Errorf("pruned %s", keep)
		}
	}
	for _, drop := range []string{"arraySum", "arraySum basic"} {
		if _, ok := log.Functions[drop]; ok {
			t.Errorf("kept retired %s", drop)
		}
	}
	if data, err := os.ReadFile(logPath(root) + ".bak"); err != nil || !strings.Contains(string(data), "arraySum") {
		t.Fatalf("backup missing or incomplete: %v", err)
	}
}

func TestParsePlayPracticeFlags(t *testing.T) {
	opts, ctl, perr := parsePlay([]string{"--run", "--fn=twoSum,rob"})
	if ctl != gnuOK || perr || opts.runMode != "all" || opts.fns != "twoSum,rob" {
		t.Fatalf("--run --fn: %+v", opts)
	}
	opts, _, _ = parsePlay([]string{"--hint", "rob", "--recog"})
	if opts.hint != "rob" || !opts.recognize {
		t.Fatalf("--hint/--recognize: %+v", opts)
	}
	if _, _, perr := parsePlay([]string{"--fn"}); !perr {
		t.Fatal("bare --fn should be a usage error")
	}
}

func captureOut(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// A fake repo whose 06_dp_reflex holds a written rob and a blank climbStairs.
func TestGradeAndLogEndToEnd(t *testing.T) {
	repo := t.TempDir()
	dir := writeReflexDir(repo, "06_dp_reflex")
	_ = os.MkdirAll(dir, 0o755)
	_ = os.WriteFile(filepath.Join(repo, "go.mod"), []byte("module fake\n\ngo 1.22\n"), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte(`package main

func rob(nums []int) int { return 4 }
func climbStairs(n int) int { panic("Implement from memory") }
func main() {}
`), 0o644)
	assertHelper := gradeFixtureTest[:strings.Index(gradeFixtureTest, "func TestGood")]
	_ = os.WriteFile(filepath.Join(dir, "main_test.go"), []byte(assertHelper+`
func TestRob(t *testing.T)         { assert(t, "rob basic", rob([]int{1, 2, 3}) == 4) }
func TestClimbStairs(t *testing.T) { assert(t, "climbStairs basic", climbStairs(3) == 3) }
`), 0o644)

	var ok bool
	out := captureOut(t, func() { ok = gradeAndLog(repo, []string{"rob", "climbStairs"}) })
	if ok {
		t.Fatal("a todo function means not everything passed")
	}
	for _, want := range []string{"PASS: rob basic", "✓ rob", "first try ✓", "· climbStairs", "not written yet"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
	log := loadLog(repo)
	if log.Functions["rob"].Box != 1 {
		t.Errorf("rob = %+v", log.Functions["rob"])
	}

	// Second run the same day: a pass is noted, the schedule is not moved.
	out = captureOut(t, func() { gradeAndLog(repo, []string{"rob"}) })
	if !strings.Contains(out, "passing after 2 runs") {
		t.Errorf("second run:\n%s", out)
	}

	// A build error is reported and grades nothing.
	_ = os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() { x }\n"), 0o644)
	if captureOut(t, func() { ok = gradeAndLog(repo, []string{"rob"}) }); ok {
		t.Fatal("a compile error is not a pass")
	}
}

func TestPrintGradesFirstTryMissShowsHint(t *testing.T) {
	now := time.Now()
	log := drillLog{Functions: map[string]fnRecord{}}
	results := []fnResult{
		{fn: "rob", status: statusFail, detail: "rob basic", output: "FAIL: rob basic\n"},
		{fn: "dfs", status: statusNoTest},
	}
	recordGrades(&log, results, now)
	out := captureOut(t, func() { printGrades(log, results, now) })
	for _, want := range []string{"✗ rob", "first try ✗", "hint: " + fnHints["rob"], "no TestDfs"} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q:\n%s", want, out)
		}
	}
	if statusTimeout.String() != "timeout" || formatSecs(75) != "1m15s" || formatSecs(9) != "9s" {
		t.Error("String/formatSecs")
	}
	if firstPanicLine("x\npanic: boom\n") != "panic: boom" || !strings.Contains(firstPanicLine(""), "no assert") {
		t.Error("firstPanicLine")
	}
}

func TestPrintersForReviewHintsAndGap(t *testing.T) {
	now := time.Date(2026, 10, 19, 9, 0, 0, 0, time.UTC)
	log := drillLog{Functions: map[string]fnRecord{
		"floodFill": {Fails: 3, LastFail: "2026-10-04"},
		"twoSum":    {Passes: 1, Box: 2, Due: "2026-10-15", LastPass: "2026-10-08"},
		"rob":       {Passes: 1, Box: 1, Due: "2026-10-19"},
		"dfs":       {Passes: 2},
	}}
	out := captureOut(t, func() {
		printReview(log, now, "", false)
		printReview(log, now, "", true)
		printSessionGap(log, now)
	})
	for _, want := range []string{"REVIEW", "overdue 4d · box 2", "due today · box 1", "3 fails so far", "--fn=", "review: ", "last session: 11 days ago"} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q:\n%s", want, out)
		}
	}
	if describeDue(dueItem{fresh: true}) != "never drilled" || describeDue(dueItem{fresh: true, rec: fnRecord{Passes: 1}}) != "not scheduled yet" {
		t.Error("describeDue")
	}
	same := captureOut(t, func() {
		printSessionGap(drillLog{Functions: map[string]fnRecord{"x": {LastPass: "2026-10-19"}}}, now)
		printSessionGap(drillLog{Functions: map[string]fnRecord{"x": {LastFail: "2026-10-18"}}}, now)
		printSessionGap(drillLog{Functions: map[string]fnRecord{}}, now)
	})
	if !strings.Contains(same, "today") || !strings.Contains(same, "yesterday") {
		t.Errorf("session gap:\n%s", same)
	}

	out = captureOut(t, func() {
		if err := printHints("rob, dailyTemperatures"); err != nil {
			t.Error(err)
		}
	})
	if !strings.Contains(out, "pattern:   dp") || !strings.Contains(out, "mono-stack") || !strings.Contains(out, "practice:") {
		t.Errorf("hints:\n%s", out)
	}
	if err := printHints("nope"); err == nil {
		t.Error("unknown function should error")
	}
}

func TestRecognizeIncludesCardsAndSaves(t *testing.T) {
	repo := t.TempDir()
	f := cardFile{Cards: map[string]card{
		"top-k": {Slug: "top-k", Title: "Top K", Pattern: "heap", Signal: "k most frequent", Insight: "size-k heap", Logged: "2026-10-01", Due: "2026-10-18"},
		"bad":   {Slug: "bad", Pattern: "not-a-pattern"},
	}}
	if err := saveCards(repo, f); err != nil {
		t.Fatal(err)
	}
	qs := cardQuestions(loadCards(repo))
	if len(qs) != 1 || qs[0].id != "lc:top-k" || !qs[0].accepts("heap") || !strings.Contains(qs[0].notIt, "size-k heap") {
		t.Fatalf("cardQuestions = %+v", qs)
	}
	if due := dueCards(f, time.Date(2026, 10, 19, 0, 0, 0, 0, time.UTC)); len(due) != 1 || due[0].Slug != "top-k" {
		t.Fatalf("dueCards = %+v", due)
	}

	var out bytes.Buffer
	if err := recognize(repo, strings.NewReader("?\n?\n?\n?\n?\n"), &out); err != nil {
		t.Fatal(err)
	}
	if got := len(loadLog(repo).Recognition); got != recognizeCount {
		t.Fatalf("recorded %d misses, want %d:\n%s", got, recognizeCount, out.String())
	}
	if err := recognize(repo, strings.NewReader(""), &out); err != nil {
		t.Fatal("an empty round saves nothing and is not an error")
	}
}
