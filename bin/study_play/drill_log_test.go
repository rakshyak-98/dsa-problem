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
