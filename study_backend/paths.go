package main

import (
	"os"
	"path/filepath"
)

// User-facing layout:
//   drills/backend/explain/core5/     — daily Core 5 verbal drills
//   drills/backend/explain/blocks/  — resume-themed concept blocks
//   drills/backend/write/core5/     — Go reflex implementations
//   drills/backend/scenario/        — STAR / mock interview prompts

func findRepoRoot(from string) string {
	dir := from
	for {
		if isBackendWorkspace(filepath.Join(dir, "drills", "backend")) {
			return dir
		}
		if _, err := os.Stat(filepath.Join(dir, "study_backend")); err == nil {
			if _, err2 := os.Stat(filepath.Join(dir, "study_play")); err2 == nil {
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
	p := filepath.Join(repoRoot, "drills", "backend", "explain", "core5")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_backend", "practice", "explain", "core5")
	}
	return p
}

func explainBlockDir(repoRoot, block string) string {
	p := filepath.Join(repoRoot, "drills", "backend", "explain", "blocks", block)
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_backend", "practice", "explain", "blocks", block)
	}
	return p
}

func writeCore5Dir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "backend", "write", "core5")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_backend", "practice", "write", "core5")
	}
	return p
}

func scenarioDir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "backend", "scenario")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_backend", "practice", "scenario")
	}
	return p
}
