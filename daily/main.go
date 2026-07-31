// Unified daily practice — read + write + asks in one command
//
// RUN:              go run .
// RUN with tests:   go run . -- --run
// Core only:        go run . -- --micro
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func repoRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	if _, err := os.Stat(filepath.Join(wd, "study_play")); err == nil {
		return wd
	}
	if _, err := os.Stat(filepath.Join(wd, "..", "study_play")); err == nil {
		p, _ := filepath.Abs(filepath.Join(wd, ".."))
		return p
	}
	return wd
}

func runIn(dir string, args ...string) error {
	cmd := exec.Command("go", append([]string{"run", "."}, args...)...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	root := repoRoot()
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}

	micro := false
	run := false
	passArgs := []string{}
	for _, a := range args {
		switch a {
		case "--micro":
			micro = true
			passArgs = append(passArgs, a)
		case "--run":
			run = true
			passArgs = append(passArgs, a)
		}
	}
	_ = micro

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              DAILY DSA PRACTICE — UNIFIED                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Order: Question literacy → Read → Core 5 → Write specialty → Primary problem")
	fmt.Println()

	fmt.Println("━━━ 1/3  CODE READING (study_code) ━━━━━━━━━━━━━━━━━━━━━")
	if err := runIn(filepath.Join(root, "study_code"), passArgs...); err != nil && run {
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("━━━ 2/3  REFLEX WRITING (study_play) ━━━━━━━━━━━━━━━━━━━")
	if err := runIn(filepath.Join(root, "study_play"), passArgs...); err != nil && run {
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("━━━ 3/3  VARIANTS (optional stretch) ━━━━━━━━━━━━━━━━━")
	fmt.Println("  go run ./study_play/variants")
	fmt.Println()
	fmt.Println("━━━ TRACK & VISUALIZE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Tracker:    study_play/study_tracker.html")
	fmt.Println("  Visualizer: visualizer/index.html")
	fmt.Println()
}
