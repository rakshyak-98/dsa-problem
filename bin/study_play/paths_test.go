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
	repo := t.TempDir()
	reflex := filepath.Join(repo, "drills", "write", "reflex", "01_arrays_reflex")
	if err := os.MkdirAll(reflex, 0o755); err != nil {
		t.Fatal(err)
	}
	_, drillPath := resolvePlayPaths(repo, "01_arrays_reflex")
	if drillPath != reflex {
		t.Fatalf("resolve local: %s", drillPath)
	}
}

func TestFindRepoRoot(t *testing.T) {
	repo := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, "drills", "write", "reflex"), 0o755); err != nil {
		t.Fatal(err)
	}
	nested := filepath.Join(repo, "bin", "study_play")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	if findRepoRoot(nested) != repo {
		t.Fatal("expected repo root from nested tool dir")
	}
}

func TestWriteDrillPathHelpers(t *testing.T) {
	repo := t.TempDir()
	core5 := filepath.Join(repo, "drills", "write", "core5")
	variants := filepath.Join(repo, "drills", "write", "variants")
	tracker := filepath.Join(repo, "drills", "tracker")
	for _, p := range []string{core5, variants, tracker} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if writeCore5Dir(repo) != core5 {
		t.Fatalf("core5: %s", writeCore5Dir(repo))
	}
	if writeVariantsDir(repo) != variants {
		t.Fatalf("variants: %s", writeVariantsDir(repo))
	}
	if trackerDir(repo) != tracker {
		t.Fatalf("tracker: %s", trackerDir(repo))
	}

}

func TestSolutionsDir(t *testing.T) {
	repo := t.TempDir()
	solutions := filepath.Join(repo, "drills", "solutions")
	if err := os.MkdirAll(solutions, 0o755); err != nil {
		t.Fatal(err)
	}
	if solutionsDir(repo) != solutions {
		t.Fatalf("solutions: %s", solutionsDir(repo))
	}
}

func TestWriteReflexDir(t *testing.T) {
	repo := t.TempDir()
	want := filepath.Join(repo, "drills", "write", "reflex", "01_arrays_reflex")
	if writeReflexDir(repo, "01_arrays_reflex") != want {
		t.Fatalf("got %s", writeReflexDir(repo, "01_arrays_reflex"))
	}
}

func TestPrintTodaySunday(t *testing.T) {
	sun := todayDrillFromWeekday(time.Sunday)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printToday(sun, false)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !strings.Contains(buf.String(), "Sunday") {
		t.Fatal("expected sunday note")
	}
}
