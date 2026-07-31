package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestTodayDrillFromWeekday(t *testing.T) {
	mon := todayDrillFromWeekday(time.Monday)
	if mon.file != "01_arrays_reflex" {
		t.Fatalf("monday: %s", mon.file)
	}
	sun := todayDrillFromWeekday(time.Sunday)
	if sun.file != "07_graphs_reflex" {
		t.Fatalf("sunday: %s", sun.file)
	}
}

func TestResolvePlayPaths(t *testing.T) {
	dir := t.TempDir()
	drillDir := filepath.Join(dir, "drills", "01_arrays_reflex")
	if err := os.MkdirAll(drillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	playRoot, drillPath := resolvePlayPaths(dir, "01_arrays_reflex")
	if playRoot != dir || drillPath != filepath.Join(dir, "drills", "01_arrays_reflex") {
		t.Fatalf("resolve local: root=%s path=%s", playRoot, drillPath)
	}

	nested := filepath.Join(dir, "study_play", "drills", "02_hashing_reflex")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	playRoot, drillPath = resolvePlayPaths(dir, "02_hashing_reflex")
	if playRoot != filepath.Join(dir, "study_play") {
		t.Fatalf("resolve nested root: %s", playRoot)
	}
}

func TestPrintTodaySunday(t *testing.T) {
	sun := todayDrillFromWeekday(time.Sunday)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printToday(sun)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "Sunday") {
		t.Fatal("expected sunday note")
	}
}
