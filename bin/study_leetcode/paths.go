package main

import (
	"os"
	"path/filepath"
)

func findRepoRoot(from string) string {
	dir := from
	for {
		if _, err := os.Stat(filepath.Join(dir, "drills", "leetcode")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return from
		}
		dir = parent
	}
}

func dailyJSONPath(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "leetcode", "daily.json")
}
