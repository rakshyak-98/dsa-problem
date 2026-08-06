// Unified daily practice — read + write + asks in one command
//
// RUN:              go run .          (from repo root)
// RUN with tests:   go run . -- --run
// Core only:        go run . -- --drill
// Select track:     go run . -- -t dsa|read|write|backend
// List tracks:      go run . -- --list-tracks
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
	opts := parseDailyArgs(os.Args[1:])
	if opts.help {
		printHelp()
		return
	}
	if opts.listTracks {
		printTrackList()
		return
	}
	if code := runUnified(root, opts); code != 0 {
		os.Exit(code)
	}
}
