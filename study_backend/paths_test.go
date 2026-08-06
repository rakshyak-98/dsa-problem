package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindRepoRoot(t *testing.T) {
	root := findRepoRoot(mustCwd())
	if _, err := os.Stat(filepath.Join(root, "study_backend")); err != nil {
		t.Fatalf("expected study_backend under %s", root)
	}
}

func TestExplainCore5Dir(t *testing.T) {
	root := findRepoRoot(mustCwd())
	p := explainCore5Dir(root)
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("explain core5 dir missing: %s", p)
	}
}

func TestWriteCore5Dir(t *testing.T) {
	root := findRepoRoot(mustCwd())
	p := writeCore5Dir(root)
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("write core5 dir missing: %s", p)
	}
}

func TestTodayBlockCount(t *testing.T) {
	if len(blocks) != 8 {
		t.Fatalf("expected 8 resume blocks, got %d", len(blocks))
	}
	b := todayBlock()
	if b.file == "" {
		t.Fatal("todayBlock returned empty file")
	}
}
