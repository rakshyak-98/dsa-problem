package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintUnifiedHeaderFooter(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printUnifiedHeader(trackDSA)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	if len(out) < 20 {
		t.Fatal("header/footer too short")
	}
}

func TestRunUnifiedDSA(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackDSA, drillKind: "reflex", passArgs: []string{"--drill", "reflex"}})
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_play") {
		t.Fatalf("expected study_play only, got %v", calls)
	}
}

func TestRunUnifiedCatalog(t *testing.T) {
	cases := []struct {
		track  drillTrack
		module string
	}{
		{trackDSA, "study_play"},
		{trackWrite, "study_play"},
		{trackLeetcode, "study_leetcode"},
	}
	for _, tc := range cases {
		var calls [][]string
		commandRunner = func(dir string, args ...string) error {
			calls = append(calls, append([]string{dir}, args...))
			return nil
		}
		code := runUnified("/tmp/repo", dailyOptions{
			track:    tc.track,
			catalog:  true,
			passArgs: []string{"--catalog"},
		})
		commandRunner = runIn
		if code != 0 {
			t.Fatalf("%s: expected success, got %d", tc.track, code)
		}
		if len(calls) != 1 || !containsAll(calls[0][0], tc.module) {
			t.Fatalf("%s: expected %s, got %v", tc.track, tc.module, calls)
		}
		if len(calls[0]) != 2 || calls[0][1] != "--catalog" {
			t.Fatalf("%s: expected sole --catalog arg, got %v", tc.track, calls[0])
		}
	}
}

func TestRunUnifiedCatalogFail(t *testing.T) {
	commandRunner = func(dir string, args ...string) error { return errTest }
	defer func() { commandRunner = runIn }()

	if code := runUnified("/tmp/repo", dailyOptions{track: trackDSA, catalog: true}); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

func TestRunUnifiedLeetcode(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackLeetcode})
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_leetcode") {
		t.Fatalf("expected study_leetcode only, got %v", calls)
	}
}

func TestRunUnifiedLeetcodeRun(t *testing.T) {
	calls := [][]string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, append([]string{dir}, args...))
		return nil
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{track: trackLeetcode, run: true, passArgs: []string{"--run"}}
	if code := runUnified("/tmp/repo", opts); code != 0 {
		t.Fatalf("expected success, got %d", code)
	}
	if len(calls) != 1 || !containsAll(calls[0][0], "study_leetcode") || !containsAll(strings.Join(calls[0], " "), "--run") {
		t.Fatalf("expected study_leetcode --run, got %v", calls)
	}
}

func TestRunUnifiedDSALeetcodeSide(t *testing.T) {
	calls := [][]string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, append([]string{dir}, args...))
		return nil
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		runSide:  "leetcode",
		passArgs: []string{"--run", "-l"},
	}
	if code := runUnified("/tmp/repo", opts); code != 0 {
		t.Fatalf("expected success, got %d", code)
	}
	if len(calls) != 1 || !containsAll(calls[0][0], "study_leetcode") {
		t.Fatalf("expected study_leetcode, got %v", calls)
	}
}

func TestFilterLeetcodePassArgs(t *testing.T) {
	got := filterLeetcodePassArgs([]string{"--run", "leetcode", "-l", "--refresh"})
	if len(got) != 2 || got[0] != "--run" || got[1] != "--refresh" {
		t.Fatalf("got %v", got)
	}
	got = filterLeetcodePassArgs([]string{"--run", "reflex", "-w"})
	if len(got) != 1 || got[0] != "--run" {
		t.Fatalf("got %v", got)
	}
}

func TestPrintHelp(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printHelp()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	for _, want := range []string{"Usage:", "-h, --help", "-V, --version", "--track=NAME", "--drill[=KIND]", "--refresh", "--problems", "Exit status:"} {
		if !strings.Contains(out, want) {
			t.Fatalf("help missing %q:\n%s", want, out)
		}
	}
	for _, gone := range []string{"cards", "backend", "revision", "cram"} {
		if strings.Contains(out, gone) {
			t.Fatalf("help should not mention %q track", gone)
		}
	}
	if strings.Contains(out, "Examples:") {
		t.Fatal("help should not include examples section")
	}
}

func TestRunUnifiedUnknownTrack(t *testing.T) {
	if code := runUnified("/tmp/repo", dailyOptions{track: "nope"}); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

func TestRunUnifiedRunWriteDefault(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		passArgs: []string{"--run", "reflex"},
	}
	if code := runUnified("/tmp/repo", opts); code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_play") {
		t.Fatalf("expected study_play only, got %v", calls)
	}
}

func TestRunUnifiedFailOnRun(t *testing.T) {
	commandRunner = func(dir string, args ...string) error {
		return errTest
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		passArgs: []string{"--run", "reflex"},
	}
	if code := runUnified("/tmp", opts); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

var errTest = &testError{}

type testError struct{}

func (e *testError) Error() string { return "test error" }

func TestPrintDrillArgError(t *testing.T) {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	printDrillArgError(trackDSA, true, "")
	w.Close()
	os.Stderr = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	for _, want := range []string{"requires an argument", "Valid arguments: reflex"} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing %q:\n%s", want, out)
		}
	}
}

func TestRunUnifiedReset(t *testing.T) {
	var gotDir string
	var gotArgs []string
	commandRunner = func(dir string, args ...string) error {
		gotDir, gotArgs = dir, args
		return nil
	}
	defer func() { commandRunner = runIn }()

	opts := parseDailyArgs([]string{"--reset"})
	if !opts.reset {
		t.Fatalf("--reset not parsed: %+v", opts)
	}
	if code := runUnified("/tmp/repo", opts); code != 0 {
		t.Fatal("expected success")
	}
	if !containsAll(gotDir, "study_play") || len(gotArgs) != 1 || gotArgs[0] != "--reset" {
		t.Fatalf("expected study_play --reset, got %s %v", gotDir, gotArgs)
	}
}
