// Unified daily practice — read + write + asks in one command
//
// RUN:              go run .          (from repo root)
// RUN with tests:   go run . -- --run
// Core only:        go run . -- --drill core
// Select track:     go run . -- -t dsa|read|write|backend|cards
// List tracks:      go run . -- --list-tracks
// Cards:            go run . -- --track cards --due
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
	cmd.Stdin = os.Stdin
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
	if opts.drillMissing {
		printDrillArgError(true, "")
		os.Exit(1)
	}
	if opts.drillUnknown != "" {
		printDrillArgError(false, opts.drillUnknown)
		os.Exit(1)
	}
	if opts.solutionMissing {
		printSolutionArgError(true, "")
		os.Exit(1)
	}
	if opts.solutionUnknown != "" {
		printSolutionArgError(false, opts.solutionUnknown)
		os.Exit(1)
	}
	if opts.core5 {
		if code := runCore5(root); code != 0 {
			os.Exit(code)
		}
		return
	}
	if code := runUnified(root, opts); code != 0 {
		os.Exit(code)
	}
}
