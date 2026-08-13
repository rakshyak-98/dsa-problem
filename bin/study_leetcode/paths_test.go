package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindRepoRoot(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	root := findRepoRoot(cwd)
	if _, err := os.Stat(filepath.Join(root, "drills", "leetcode")); err != nil {
		t.Fatalf("expected drills/leetcode under %s: %v", root, err)
	}
}

func TestPracticeSetPath(t *testing.T) {
	p := practiceSetPath("/tmp/repo", "monday.md")
	if p != "/tmp/repo/drills/leetcode/sets/monday.md" {
		t.Fatalf("unexpected path: %s", p)
	}
}
