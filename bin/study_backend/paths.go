package main

import (
	"os"
	"path/filepath"
)

// User-facing layout:
//   drills/backend/explain/core5/     — daily Core 5 verbal drills
//   drills/backend/explain/blocks/  — resume-themed concept blocks
//   drills/backend/write/core5/     — Go reflex implementations
//   drills/backend/explain/revision/ — weekly cross-topic revision drills

func findRepoRoot(from string) string {
	dir := from
	for {
		if isBackendWorkspace(filepath.Join(dir, "drills", "backend")) {
			return dir
		}
		if _, err := os.Stat(filepath.Join(dir, "bin", "study_backend")); err == nil {
			if _, err2 := os.Stat(filepath.Join(dir, "drills")); err2 == nil {
				return dir
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return from
		}
		dir = parent
	}
}

func isBackendWorkspace(backendDir string) bool {
	_, err := os.Stat(filepath.Join(backendDir, "explain", "core5"))
	return err == nil
}

func explainCore5Dir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "backend", "explain", "core5")
}

func explainBlockDir(repoRoot, block string) string {
	return filepath.Join(repoRoot, "drills", "backend", "explain", "blocks", block)
}

func writeCore5Dir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "backend", "write", "core5")
}

func explainRevisionDir(repoRoot, file string) string {
	return filepath.Join(repoRoot, "drills", "backend", "explain", "revision", file)
}

func scenarioDir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "backend", "scenario")
}
