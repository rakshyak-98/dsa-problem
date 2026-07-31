package main

import (
	"os"
	"path/filepath"
	"time"
)

func todayDrillFromWeekday(weekday time.Weekday) drill {
	dayIndex := int(weekday)
	drillIndex := dayIndex - 1
	if dayIndex == 0 {
		drillIndex = 6
	}
	return drills[drillIndex]
}

func resolvePlayPaths(root, drillFile string) (playRoot, drillPath string) {
	playRoot = root
	drillPath = filepath.Join(root, "drills", drillFile)
	if _, err := os.Stat(drillPath); err != nil {
		alt := filepath.Join(root, "study_play", "drills", drillFile)
		if _, err2 := os.Stat(alt); err2 == nil {
			drillPath = alt
			playRoot = filepath.Join(root, "study_play")
		}
	}
	return playRoot, drillPath
}
