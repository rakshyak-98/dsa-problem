// Unified daily practice — reflex read + write in one command
//
// RUN:              go run .          (from repo root)
// RUN with tests:   go run . -- --run reflex -r|-w
// Write Core 5:    go run . -- --drill core
// Reflex read:     go run . -- --track read
// Select track:    go run . -- -t dsa|read|write|backend|cards
// List tracks:     go run . -- --list-tracks
// Cards:           go run . -- --track cards --due
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
		printDrillArgError(opts.track, true, "")
		os.Exit(1)
	}
	if opts.drillUnknown != "" {
		printDrillArgError(opts.track, false, opts.drillUnknown)
		os.Exit(1)
	}
	if opts.solutionMissing {
		printSolutionArgError(opts.track, true, "")
		os.Exit(1)
	}
	if opts.solutionUnknown != "" {
		printSolutionArgError(opts.track, false, opts.solutionUnknown)
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
