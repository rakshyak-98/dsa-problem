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

func TestDailyJSONPath(t *testing.T) {
	p := dailyJSONPath("/tmp/repo")
	if p != "/tmp/repo/drills/leetcode/daily.json" {
		t.Fatalf("unexpected path: %s", p)
	}
}
