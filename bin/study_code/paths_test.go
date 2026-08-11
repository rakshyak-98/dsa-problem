package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDrillPaths(t *testing.T) {
	repo := t.TempDir()
	weekdayDrill := filepath.Join(repo, "drills", "read", "weekday", "01_scan_structure")
	if err := os.MkdirAll(weekdayDrill, 0o755); err != nil {
		t.Fatal(err)
	}
	weekday := readDrillDir(repo, "01_scan_structure")
	if weekday != weekdayDrill {
		t.Fatalf("weekday: got %s want %s", weekday, weekdayDrill)
	}
}

func TestFindRepoRoot(t *testing.T) {
	repo := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, "drills", "read", "weekday"), 0o755); err != nil {
		t.Fatal(err)
	}
	sub := filepath.Join(repo, "bin", "study_code")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	got := findRepoRoot(sub)
	if got != repo {
		t.Fatalf("findRepoRoot: got %s want %s", got, repo)
	}
}
