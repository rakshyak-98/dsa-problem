package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseTestOutput(t *testing.T) {
	out := "PASS: twoSum\nPASS: binarySearch\nFAIL: removeDuplicates\n"
	passed, failed := parseTestOutput(out)
	if len(passed) != 2 || len(failed) != 1 {
		t.Fatalf("parse: passed=%v failed=%v", passed, failed)
	}
	if passed[0] != "twoSum" || failed[0] != "removeDuplicates" {
		t.Fatal("parse names")
	}
}

func TestRecordResultAndLogIO(t *testing.T) {
	dir := t.TempDir()
	log := drillLog{Functions: map[string]fnRecord{}}
	recordResult(&log, "twoSum", true)
	recordResult(&log, "twoSum", false)
	if log.Functions["twoSum"].Passes != 1 || log.Functions["twoSum"].Fails != 1 {
		t.Fatal("recordResult counts")
	}
	if err := saveLog(dir, log); err != nil {
		t.Fatal(err)
	}
	loaded := loadLog(dir)
	if loaded.Functions["twoSum"].Passes != 1 {
		t.Fatal("loadLog")
	}
}

func TestUpdateLogFromOutput(t *testing.T) {
	dir := t.TempDir()
	updateLogFromOutput(dir, "PASS: a\nPASS: b\n", []string{"a", "b", "c"})
	log := loadLog(dir)
	if log.Functions["a"].Passes != 1 || log.Functions["b"].Passes != 1 {
		t.Fatal("update pass")
	}
	updateLogFromOutput(dir, "FAIL: x\n", []string{"x", "y"})
	log = loadLog(dir)
	if log.Functions["x"].Fails < 1 {
		t.Fatal("update fail")
	}
	updateLogFromOutput(dir, "", []string{"all", "funcs"})
	log = loadLog(dir)
	if log.Functions["all"].Passes < 1 {
		t.Fatal("update bulk pass")
	}
}

func TestPrintWeakFunctions(t *testing.T) {
	dir := t.TempDir()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printWeakFunctions(dir, 3)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "No drill history") {
		t.Fatal("expected empty history message")
	}
	updateLogFromOutput(dir, "PASS: a\n", []string{"a"})
	r, w, _ = os.Pipe()
	os.Stdout = w
	printWeakFunctions(dir, 3)
	w.Close()
	os.Stdout = old
	buf.Reset()
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "a") {
		t.Fatal("expected weak function listed")
	}
}

func TestLoadLogInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(logPath(dir), []byte("{bad"), 0o644); err != nil {
		t.Fatal(err)
	}
	log := loadLog(dir)
	if log.Functions == nil {
		t.Fatal("expected empty functions map")
	}
}

func TestPrintWeakFunctionsSorted(t *testing.T) {
	dir := t.TempDir()
	updateLogFromOutput(dir, "FAIL: slow\n", []string{"slow"})
	updateLogFromOutput(dir, "PASS: fast\n", []string{"fast"})
	updateLogFromOutput(dir, "FAIL: slow\n", []string{"slow"})
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printWeakFunctions(dir, 2)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "slow") {
		t.Fatal("expected slow ranked weak")
	}
}

func TestOrDash(t *testing.T) {
	if orDash("") != "—" || orDash("2026-01-01") != "2026-01-01" {
		t.Fatal("orDash")
	}
}

func TestLogPath(t *testing.T) {
	p := logPath("/tmp/foo")
	if filepath.Base(p) != ".drill_log.json" {
		t.Fatal("logPath")
	}
	_ = os.Remove(p)
}

// `go test -v` frames every test with lines like "--- FAIL: TestRotateRight
// (0.00s)". An unanchored FAIL pattern captured those as if they were drill
// functions, and they then dominated the weak list.
func TestParseTestOutputIgnoresGoTestFraming(t *testing.T) {
	output := `=== RUN   TestRotateRight
PASS: rotateRight k=0
FAIL: rotateRight k=1
    main_test.go:13: FAIL: rotateRight k=1
--- FAIL: TestRotateRight (0.00s)
=== RUN   TestArraySum
PASS: reverseInPlace basic
--- PASS: TestArraySum (0.00s)
FAIL
FAIL	github.com/rakshyak-98/dsa-problem/drills/write/reflex/01_arrays_reflex	0.004s
`
	passed, failed := parseTestOutput(output)
	wantPass := []string{"rotateRight k=0", "reverseInPlace basic"}
	wantFail := []string{"rotateRight k=1"}
	if !equalStrings(passed, wantPass) {
		t.Errorf("passed = %q, want %q", passed, wantPass)
	}
	if !equalStrings(failed, wantFail) {
		t.Errorf("failed = %q, want %q", failed, wantFail)
	}
}

func TestUpdateLogRollsAssertsUpToTheirFunction(t *testing.T) {
	root := t.TempDir()
	output := `PASS: reverseInPlace basic
PASS: reverseInPlace empty
PASS: reverseInPlace negatives
PASS: rotateRight k=0
FAIL: rotateRight k=1
--- FAIL: TestRotateRight (0.00s)
`
	updateLogFromOutput(root, output, []string{"reverseInPlace", "rotateRight", "runningSum"})
	log := loadLog(root)

	// Three passing asserts are still one pass for the function.
	if got := log.Functions["reverseInPlace"]; got.Passes != 1 || got.Fails != 0 {
		t.Errorf("reverseInPlace = %+v, want 1 pass 0 fails", got)
	}
	// One failing assert fails the whole function, even alongside a pass.
	if got := log.Functions["rotateRight"]; got.Passes != 0 || got.Fails != 1 {
		t.Errorf("rotateRight = %+v, want 0 passes 1 fail", got)
	}
	// Never reached during a failing run: unfinished, not untouched.
	if got := log.Functions["runningSum"]; got.Fails != 1 {
		t.Errorf("runningSum = %+v, want 1 fail", got)
	}
	// Per-assert detail is kept.
	if got := log.Functions["reverseInPlace basic"]; got.Passes != 1 {
		t.Errorf("assert detail lost: %+v", got)
	}
	// Framing lines are never recorded.
	for name := range log.Functions {
		if name == "TestRotateRight (0.00s)" {
			t.Error("go test framing leaked into the log")
		}
	}
}

func TestPruneLogNoiseKeepsRealEntries(t *testing.T) {
	log := drillLog{Functions: map[string]fnRecord{
		"twoSum":                          {Passes: 3},
		"rotateRight k=1":                 {Fails: 2},
		"TestRotateRight (0.00s)":         {Fails: 9},
		"TestAll/maxSumSubarrayK (0.00s)": {Fails: 7},
		"TestAll (0.00s)":                 {Fails: 8},
	}}
	pruneLogNoise(&log)
	for _, keep := range []string{"twoSum", "rotateRight k=1"} {
		if _, ok := log.Functions[keep]; !ok {
			t.Errorf("pruned a real entry: %s", keep)
		}
	}
	for _, drop := range []string{"TestRotateRight (0.00s)", "TestAll/maxSumSubarrayK (0.00s)", "TestAll (0.00s)"} {
		if _, ok := log.Functions[drop]; ok {
			t.Errorf("kept framing noise: %s", drop)
		}
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
