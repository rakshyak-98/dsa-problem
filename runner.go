package main

import (
	"fmt"
	"path/filepath"
)

var commandRunner = runIn

func printUnifiedHeader() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              DAILY DSA PRACTICE — UNIFIED                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Order: Question literacy → Read → Core 5 → Write specialty → Primary problem")
	fmt.Println()
}

func printUnifiedFooter() {
	fmt.Println()
	fmt.Println("━━━ 3/3  VARIANTS (optional stretch) ━━━━━━━━━━━━━━━━━")
	fmt.Println("  go run -C drills/write/variants .")
	fmt.Println()
	fmt.Println("━━━ TRACK & VISUALIZE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Tracker:    drills/tracker/study_tracker.html")
	fmt.Println("  Visualizer: reference/visualizer/index.html")
	fmt.Println()
}

func runUnified(root string, passArgs []string, run bool) int {
	printUnifiedHeader()

	fmt.Println("━━━ 1/3  CODE READING (read drills) ━━━━━━━━━━━━━━━━━━━━━")
	if err := commandRunner(filepath.Join(root, "bin", "study_code"), passArgs...); err != nil && run {
		return 1
	}

	fmt.Println()
	fmt.Println("━━━ 2/3  REFLEX WRITING (write drills) ━━━━━━━━━━━━━━━━━━")
	if err := commandRunner(filepath.Join(root, "bin", "study_play"), passArgs...); err != nil && run {
		return 1
	}

	printUnifiedFooter()
	return 0
}
