package main

import (
	"os"
	"path/filepath"
)

func findRepoRoot(from string) string {
	dir := from
	for {
		if _, err := os.Stat(filepath.Join(dir, "drills", "read", "core")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return from
		}
		dir = parent
	}
}

func readDrillDir(repoRoot, file string) string {
	if file == "00_core_read" {
		return filepath.Join(repoRoot, "drills", "read", "core", file)
	}
	return filepath.Join(repoRoot, "drills", "read", "weekday", file)
}

func readAnswersPath(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "read", "answers", "ANSWER_KEY.md")
}
