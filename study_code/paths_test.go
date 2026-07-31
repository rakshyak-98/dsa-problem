package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDrillDir(t *testing.T) {
	repo := t.TempDir()
	coreDrill := filepath.Join(repo, "drills", "read", "core", "00_core_read")
	if err := os.MkdirAll(coreDrill, 0o755); err != nil {
		t.Fatal(err)
	}
	weekdayDrill := filepath.Join(repo, "drills", "read", "weekday", "01_scan_structure")
	if err := os.MkdirAll(weekdayDrill, 0o755); err != nil {
		t.Fatal(err)
	}
	core := readDrillDir(repo, "00_core_read")
	if core != coreDrill {
		t.Fatalf("core: got %s want %s", core, coreDrill)
	}
	weekday := readDrillDir(repo, "01_scan_structure")
	if weekday != weekdayDrill {
		t.Fatalf("weekday: got %s want %s", weekday, weekdayDrill)
	}
}

func TestReadDrillDirFallback(t *testing.T) {
	repo := t.TempDir()
	core := readDrillDir(repo, "00_core_read")
	if core != filepath.Join(repo, "study_code", "practice", "read", "core", "00_core_read") {
		t.Fatal(core)
	}
	weekday := readDrillDir(repo, "01_scan_structure")
	if weekday != filepath.Join(repo, "study_code", "practice", "read", "weekday", "01_scan_structure") {
		t.Fatal(weekday)
	}
}
func TestFindRepoRootRead(t *testing.T) {
	repo := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, "drills", "read", "core"), 0o755); err != nil {
		t.Fatal(err)
	}
	if findRepoRoot(filepath.Join(repo, "study_code")) != repo {
		t.Fatal("repo root")
	}
}
