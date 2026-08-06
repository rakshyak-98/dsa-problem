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
	printUnifiedHeader()
	printUnifiedFooter()
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

func TestRunUnifiedMocked(t *testing.T) {
	calls := []string{}
	commandRunner = func(dir string, args ...string) error {
		calls = append(calls, dir)
		return nil
	}
	defer func() { commandRunner = runIn }()

	code := runUnified("/tmp/repo", []string{"--micro"}, false)
	if code != 0 {
		t.Fatal("expected success")
	}
	if len(calls) != 2 {
		t.Fatalf("expected 2 calls, got %d", len(calls))
	}
}

func TestRunUnifiedFailOnRun(t *testing.T) {
	commandRunner = func(dir string, args ...string) error {
		return errTest
	}
	defer func() { commandRunner = runIn }()

	if code := runUnified("/tmp", []string{"--run"}, true); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

var errTest = &testError{}

type testError struct{}

func (e *testError) Error() string { return "test error" }

func containsAll(s string, parts ...string) bool {
	for _, p := range parts {
		if !bytes.Contains([]byte(s), []byte(p)) {
			return false
		}
	}
	return true
}
