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
	fmt.Println("  go run ./study_play/variants")
	fmt.Println()
	fmt.Println("━━━ TRACK & VISUALIZE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Tracker:    study_play/study_tracker.html")
	fmt.Println("  Visualizer: visualizer/index.html")
	fmt.Println()
}

func runUnified(root string, passArgs []string, run bool) int {
	printUnifiedHeader()

	fmt.Println("━━━ 1/3  CODE READING (study_code) ━━━━━━━━━━━━━━━━━━━━━")
	if err := commandRunner(filepath.Join(root, "study_code"), passArgs...); err != nil && run {
		return 1
	}

	fmt.Println()
	fmt.Println("━━━ 2/3  REFLEX WRITING (study_play) ━━━━━━━━━━━━━━━━━━━")
	if err := commandRunner(filepath.Join(root, "study_play"), passArgs...); err != nil && run {
		return 1
	}

	printUnifiedFooter()
	return 0
}
