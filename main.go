// Unified daily practice — read + write + asks in one command
//
// RUN:              go run .          (from repo root)
// RUN with tests:   go run . -- --run
// Core only:        go run . -- --micro
package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func repoRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	if _, err := os.Stat(filepath.Join(wd, "bin", "study_play")); err == nil {
		return wd
	}
	if _, err := os.Stat(filepath.Join(wd, "..", "bin", "study_play")); err == nil {
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
	passArgs, run := parseDailyArgs(os.Args[1:])
	if code := runUnified(root, passArgs, run); code != 0 {
		os.Exit(code)
	}
}
