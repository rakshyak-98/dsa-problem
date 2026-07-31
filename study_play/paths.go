package main

import (
	"os"
	"path/filepath"
	"time"
)

// User-facing drill layout (repo root):
//   drills/write/reflex/   — weekday + bonus reflex drills
//   drills/write/core5/    — daily Core 5
//   drills/write/variants/ — medium variants
//   drills/read/core/      — core reading drill
//   drills/read/weekday/   — reading specialty drills
//   drills/solutions/     — write drill solutions (peek after attempt)
//   drills/tracker/        — browser study tracker

func findRepoRoot(from string) string {
	dir := from
	for {
		if isDrillWorkspace(filepath.Join(dir, "drills")) {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return from
		}
		dir = parent
	}
}

func isDrillWorkspace(drillsDir string) bool {
	_, err := os.Stat(filepath.Join(drillsDir, "write", "reflex"))
	if err == nil {
		return true
	}
	// module-internal path when symlinks unavailable
	_, err = os.Stat(filepath.Join("study_play", "practice", "write", "reflex"))
	return err == nil
}

func writeReflexDir(repoRoot, drillFile string) string {
	p := filepath.Join(repoRoot, "drills", "write", "reflex", drillFile)
	if _, err := os.Stat(filepath.Dir(p)); err != nil {
		return filepath.Join(repoRoot, "study_play", "practice", "write", "reflex", drillFile)
	}
	return p
}

func writeCore5Dir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "write", "core5")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_play", "practice", "write", "core5")
	}
	return p
}

func writeVariantsDir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "write", "variants")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_play", "practice", "write", "variants")
	}
	return p
}

func trackerDir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "tracker")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_play", "practice", "tracker")
	}
	return p
}

func solutionsDir(repoRoot string) string {
	p := filepath.Join(repoRoot, "drills", "solutions")
	if _, err := os.Stat(p); err != nil {
		return filepath.Join(repoRoot, "study_play", "_support", "solutions_reference")
	}
	return p
}

func todayDrillFromWeekday(weekday time.Weekday) drill {
	dayIndex := int(weekday)
	drillIndex := dayIndex - 1
	if dayIndex == 0 {
		drillIndex = 6
	}
	return drills[drillIndex]
}

func resolvePlayPaths(cwd, drillFile string) (repoRoot, drillPath string) {
	repoRoot = findRepoRoot(cwd)
	return repoRoot, writeReflexDir(repoRoot, drillFile)
}
