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
//   drills/read/weekday/   — reflex reading drills
//   drills/solutions/      — write drill solutions (peek after attempt)
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
	return err == nil
}

func writeReflexDir(repoRoot, drillFile string) string {
	return filepath.Join(repoRoot, "drills", "write", "reflex", drillFile)
}

func writeCore5Dir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "write", "core5")
}

func writeVariantsDir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "write", "variants")
}

func trackerDir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "tracker")
}

func solutionsDir(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "solutions")
}

func solutionCorePath(repoRoot string) string {
	return filepath.Join(solutionsDir(repoRoot), "core5.md")
}

func solutionReflexMain(repoRoot, drillFile string) string {
	return filepath.Join(solutionsDir(repoRoot), "reflex", drillFile, "main.go")
}

func solutionReflexNotes(repoRoot, drillFile string) string {
	return filepath.Join(solutionsDir(repoRoot), drillFile+".md")
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
