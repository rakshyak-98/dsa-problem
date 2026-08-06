package main

import (
	"bytes"
	"io"
	"os"
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
	if len(out) < 100 {
		t.Fatal("header/footer too short")
	}
	if !containsAll(out, "UNIFIED", "variants", "Visualizer") {
		t.Fatal("missing sections")
	}
}

func TestRunUnifiedDSA(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", dailyOptions{track: trackDSA, passArgs: []string{"--micro"}})
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
