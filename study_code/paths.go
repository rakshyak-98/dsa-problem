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
		if _, err := os.Stat(filepath.Join(dir, "study_code", "practice", "read", "core")); err == nil {
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
		p := filepath.Join(repoRoot, "drills", "read", "core", file)
		if _, err := os.Stat(p); err == nil {
			return p
		}
		return filepath.Join(repoRoot, "study_code", "practice", "read", "core", file)
	}
	p := filepath.Join(repoRoot, "drills", "read", "weekday", file)
	if _, err := os.Stat(p); err == nil {
		return p
	}
	return filepath.Join(repoRoot, "study_code", "practice", "read", "weekday", file)
}
