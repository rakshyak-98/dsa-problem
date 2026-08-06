package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestSpecialtyOnly(t *testing.T) {
	if !specialtyOnly(dailyOptions{track: trackDSA}) {
		t.Fatal("default should be specialty only")
	}
	if specialtyOnly(dailyOptions{passArgs: []string{"--verbose"}}) {
		t.Fatal("verbose disables specialty only")
	}
	if specialtyOnly(dailyOptions{passArgs: []string{"--catalog"}}) {
		t.Fatal("catalog disables specialty only")
	}
}

func TestModuleArgs(t *testing.T) {
	args := moduleArgs(dailyOptions{track: trackDSA})
	if len(args) != 1 || args[0] != "--specialty" {
		t.Fatalf("expected --specialty, got %v", args)
	}
	args = moduleArgs(dailyOptions{passArgs: []string{"--verbose"}})
	if len(args) != 1 || args[0] != "--verbose" {
		t.Fatalf("verbose passthrough: %v", args)
	}
}

func TestRunUnifiedDSAPassesSpecialty(t *testing.T) {
	var lastArgs []string
	commandRunner = func(dir string, args ...string) error {
		lastArgs = args
		return nil
	}
	defer func() { commandRunner = runIn }()

	if code := runUnified("/tmp/repo", dailyOptions{track: trackDSA}); code != 0 {
		t.Fatal("expected success")
	}
	if len(lastArgs) == 0 || lastArgs[0] != "--specialty" {
		t.Fatalf("expected --specialty, got %v", lastArgs)
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
	if len(calls) != 1 || !strings.Contains(calls[0], "study_code") {
		t.Fatalf("expected study_code only, got %v", calls)
	}
}

func TestRunUnifiedBackendCramSkipsSpecialty(t *testing.T) {
	var lastArgs []string
	commandRunner = func(dir string, args ...string) error {
		lastArgs = args
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackBackend, passArgs: []string{"--cram"}})
	if code != 0 {
		t.Fatal("expected success")
	}
	for _, a := range lastArgs {
		if a == "--specialty" {
			t.Fatal("cram should not use specialty mode")
		}
	}
}

func TestPrintHelp(t *testing.T) {
	out := captureStdout(printHelp)
	for _, want := range []string{"Usage:", "Options:", "-h, --help", "--track=NAME", "--verbose"} {
		if !strings.Contains(out, want) {
			t.Fatalf("help missing %q:\n%s", want, out)
		}
	}
}

func TestRunUnifiedUnknownTrack(t *testing.T) {
	if code := runUnified("/tmp/repo", dailyOptions{track: "nope"}); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

func TestRunUnifiedFailOnRun(t *testing.T) {
	commandRunner = func(dir string, args ...string) error {
		return errTest
	}
	defer func() { commandRunner = runIn }()

	opts := dailyOptions{track: trackDSA, passArgs: []string{"--run"}, run: true}
	if code := runUnified("/tmp", opts); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

var errTest = &testError{}

type testError struct{}

func (e *testError) Error() string { return "test error" }

func captureStdout(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf strings.Builder
	_, _ = io.Copy(&buf, r)
	return buf.String()
}
