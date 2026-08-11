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
	printDSAExtras()
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

	code := runUnified("/tmp/repo", dailyOptions{track: trackDSA, drillKind: "core", passArgs: []string{"--drill", "core"}})
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 2 {
		t.Fatalf("expected 2 calls, got %d", len(calls))
	}
}

func TestRunUnifiedReadOnly(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackRead})
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_code") {
		t.Fatalf("expected study_code only, got %v", calls)
	}
}

func TestRunUnifiedBackend(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackBackend, passArgs: []string{"--cram"}})
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_backend") {
		t.Fatalf("expected study_backend only, got %v", calls)
	}
}

func TestRunUnifiedCards(t *testing.T) {
	var gotDir string
	var gotArgs []string
	commandRunner = func(dir string, args ...string) error {
		gotDir = dir
		gotArgs = append([]string{}, args...)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{
		track:    trackCards,
		run:      true,
		passArgs: []string{"--run", "core", "-r", "--due", "--deck=jargon"},
	})
	if code != 0 {
		t.Fatal("expected success")
	}
	if !containsAll(gotDir, "study_cards") {
		t.Fatalf("expected study_cards, got %s", gotDir)
	}
	if len(gotArgs) != 2 || gotArgs[0] != "--due" || gotArgs[1] != "--deck=jargon" {
		t.Fatalf("expected filtered card args, got %v", gotArgs)
	}
}

func TestFilterCardsPassArgs(t *testing.T) {
	got := filterCardsPassArgs([]string{"--run", "reflex", "-w", "--stats", "--tag=jwt"})
	if len(got) != 2 || got[0] != "--stats" || got[1] != "--tag=jwt" {
		t.Fatalf("got %v", got)
	}
}

func TestFilterReadPassArgs(t *testing.T) {
	got := filterReadPassArgs([]string{"--brief", "--drill", "core", "--solution", "reflex"})
	if len(got) != 3 || got[0] != "--brief" || got[1] != "--solution" || got[2] != "reflex" {
		t.Fatalf("got %v", got)
	}
	got = filterReadPassArgs([]string{"--run", "core", "-w"})
	if len(got) != 2 || got[0] != "--run" || got[1] != "-w" {
		t.Fatalf("got %v", got)
	}
}

func TestRunUnifiedCoreReadRemoved(t *testing.T) {
	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		runSide:  "read",
		passArgs: []string{"--run", "core", "-r"},
	}
	if code := runUnified("/tmp/repo", opts); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
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
	for _, want := range []string{"Usage:", "Options:", "-h, --help", "--track=NAME", "--core5", "--drill reflex", "--run reflex -r", "(default:", "cards"} {
		if !strings.Contains(out, want) {
			t.Fatalf("help missing %q:\n%s", want, out)
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

func TestRunUnifiedRunReadOnly(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		runSide:  "read",
		passArgs: []string{"--run", "reflex", "-r"},
	}
	if code := runUnified("/tmp/repo", opts); code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 1 || !containsAll(calls[0], "study_code") {
		t.Fatalf("expected study_code only, got %v", calls)
	}
}

func TestRunUnifiedRunRequiresSide(t *testing.T) {
	opts := dailyOptions{
		track:    trackDSA,
		run:      true,
		passArgs: []string{"--run", "reflex"},
	}
	if code := runUnified("/tmp/repo", opts); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
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
		runSide:  "read",
		passArgs: []string{"--run", "core", "-r"},
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
	for _, want := range []string{"requires an argument", "Valid arguments: core, reflex"} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing %q:\n%s", want, out)
		}
	}
}

func TestPrintDrillArgErrorReadTrack(t *testing.T) {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	printDrillArgError(trackRead, false, "core")
	w.Close()
	os.Stderr = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "Valid arguments: reflex") {
		t.Fatalf("read track help: %s", buf.String())
	}
}

func TestRunCore5(t *testing.T) {
	called := false
	core5Runner = func(root string) error {
		called = true
		if !strings.Contains(root, "repo") {
			t.Fatalf("unexpected root: %s", root)
		}
		return nil
	}
	defer func() { core5Runner = runCore5In }()

	if code := runCore5("/tmp/repo"); code != 0 {
		t.Fatalf("expected success, got %d", code)
	}
	if !called {
		t.Fatal("expected core5Runner to be called")
	}
}

func TestRunCore5Fail(t *testing.T) {
	core5Runner = func(root string) error { return errTest }
	defer func() { core5Runner = runCore5In }()

	if code := runCore5("/tmp/repo"); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}
