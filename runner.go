package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var commandRunner = runIn

func printTrackList() {
	for _, t := range availableTracks {
		fmt.Printf("  %-8s%s\n", t.name, t.description)
	}
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Run daily practice drills.

Options:
  -h, --help               display this help message and exit
      --list-tracks        list practice tracks and exit
  -t, --track=NAME         practice track: dsa, read, write, or backend
                             (default: "dsa")
      --run                run tests (forwarded to track)
      --micro              core drills only (forwarded to track)
      --catalog            list drills in track (forwarded to track)

`)
}

func printUnknownTrack(track drillTrack) {
	fmt.Fprintf(os.Stderr, "unknown track %q\n", track)
	fmt.Fprint(os.Stderr, "Valid tracks: dsa, read, write, backend\n")
	fmt.Fprint(os.Stderr, "Try 'go run . -- --help' for more information.\n")
}

func printUnifiedHeader(track drillTrack) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              DAILY PRACTICE — UNIFIED                    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	switch track {
	case trackDSA:
		fmt.Println("Track: DSA — Question literacy → Read → Core 5 → Write specialty")
	case trackRead:
		fmt.Println("Track: DSA reading — Core Read 3 + today's specialty")
	case trackWrite:
		fmt.Println("Track: DSA writing — Core 5 + today's reflex specialty")
	case trackBackend:
		fmt.Println("Track: Backend interview — Core 5 explain/write + resume block")
	}
	fmt.Println()
}

func printDSAExtras() {
	fmt.Println()
	fmt.Println("━━━ VARIANTS (optional stretch) ━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  go run -C drills/write/variants .")
	fmt.Println()
	fmt.Println("━━━ TRACK & VISUALIZE ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Tracker:    drills/tracker/study_tracker.html")
	fmt.Println("  Visualizer: reference/visualizer/index.html")
	fmt.Println()
}

func runModule(root, module string, passArgs []string, run bool) int {
	if err := commandRunner(filepath.Join(root, "bin", module), passArgs...); err != nil && run {
		return 1
	}
	return 0
}

func runUnified(root string, opts dailyOptions) int {
	if !isKnownTrack(opts.track) {
		printUnknownTrack(opts.track)
		return 1
	}

	printUnifiedHeader(opts.track)

	switch opts.track {
	case trackDSA:
		fmt.Println("━━━ 1/2  CODE READING ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		if code := runModule(root, "study_code", opts.passArgs, opts.run); code != 0 {
			return code
		}
		fmt.Println()
		fmt.Println("━━━ 2/2  REFLEX WRITING ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		if code := runModule(root, "study_play", opts.passArgs, opts.run); code != 0 {
			return code
		}
		printDSAExtras()
	case trackRead:
		if code := runModule(root, "study_code", opts.passArgs, opts.run); code != 0 {
			return code
		}
	case trackWrite:
		if code := runModule(root, "study_play", opts.passArgs, opts.run); code != 0 {
			return code
		}
		printDSAExtras()
	case trackBackend:
		if code := runModule(root, "study_backend", opts.passArgs, opts.run); code != 0 {
			return code
		}
	}
	return 0
}

func containsAll(s string, parts ...string) bool {
	for _, p := range parts {
		if !strings.Contains(s, p) {
			return false
		}
	}
	return true
}
