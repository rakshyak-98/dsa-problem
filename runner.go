package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var commandRunner = runIn

func printTrackList() {
	fmt.Println("Available practice tracks:")
	fmt.Println()
	for _, t := range availableTracks {
		fmt.Printf("  %-8s  %s\n", t.name, t.description)
	}
	fmt.Println()
	fmt.Println("Usage: go run . -- --track <name>")
}

func printHelp() {
	fmt.Println("Unified daily practice — select a track with --track")
	fmt.Println()
	printTrackList()
	fmt.Println("Common flags (passed to the selected track):")
	fmt.Println("  --run       run tests / check answers")
	fmt.Println("  --micro     minimum tier (core drills only)")
	fmt.Println("  --catalog   list all drills in the track")
	fmt.Println()
	fmt.Println("Track-specific flags:")
	fmt.Println("  backend:  --cram, --setup")
	fmt.Println("  write:    --reset, --weak, --run-core5, --setup")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run .                         # DSA: read + write")
	fmt.Println("  go run . -- --track backend      # backend interview prep")
	fmt.Println("  go run . -- --track write --run  # test today's write drill")
	fmt.Println("  go run . -- --track backend --cram")
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
		fmt.Fprintf(os.Stderr, "Unknown track %q\n\n", opts.track)
		printTrackList()
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
