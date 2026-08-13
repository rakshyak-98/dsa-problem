package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var commandRunner = runIn
var core5Runner = runCore5In

var weekdayNames = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func todayName() string {
	return weekdayNames[time.Now().Weekday()]
}

func hasArg(args []string, flag string) bool {
	for _, a := range args {
		if a == flag {
			return true
		}
	}
	return false
}

func withBrief(args []string) []string {
	if hasArg(args, "--brief") {
		return args
	}
	return append(args, "--brief")
}

func printTrackList() {
	for _, t := range availableTracks {
		fmt.Printf("  %-8s%s\n", t.name, t.description)
	}
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Run daily DSA and interview-prep drills from the repo root.

Options:
  -h, --help               display this help message and exit
      --list-tracks        list practice tracks and exit
  -t, --track=NAME         practice track (default: "dsa")
                             dsa: reflex writing (Core 5 + weekday specialty)
                             read: reflex code-reading drills
                             write: same as dsa writing drills
                             backend: interview prep
      --core5              run the Core 5 write drill
      --drill KIND         show today's drill plan (KIND: core or reflex)
      --solution KIND      show solution file path (KIND: core or reflex)
      --run [KIND]         run drill tests (KIND: core or reflex; default: all)
      --catalog            list drills in the active track

LeetCode track (--track=leetcode):
      --set                show today's 10 LeetCode problems (default)
      --catalog            list all weekday practice sets

Read track (--track=read):
      --drill reflex       show today's reflex read plan
      --solution reflex    show read answer key section
      --run [reflex]       run today's reflex read tests

Backend track (--track=backend):
      --drill revision     show today's revision drill path
      --run revision       validate today's revision drill
      --cram               show interview cram schedule

`)
}

func printUnknownTrack(track drillTrack) {
	fmt.Fprintf(os.Stderr, "unknown track %q\n", track)
	fmt.Fprint(os.Stderr, "Valid tracks: dsa, read, write, leetcode, backend\n")
	fmt.Fprint(os.Stderr, "Try 'go run . -- --help' for more information.\n")
}

func printDrillArgError(track drillTrack, missing bool, unknown string) {
	if missing {
		fmt.Fprintln(os.Stderr, "option '--drill' requires an argument")
	} else {
		fmt.Fprintf(os.Stderr, "unknown drill kind %q\n", unknown)
	}
	fmt.Fprintf(os.Stderr, "Valid arguments: %s\n", formatDrillKinds(track))
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func printReadUseTrack() {
	fmt.Fprintln(os.Stderr, "reflex reading drills are on the read track; use --track read")
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func printRunSideConflict() {
	fmt.Fprintln(os.Stderr, "cannot use both -r/--read and -w/--write")
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func printSolutionArgError(track drillTrack, missing bool, unknown string) {
	if missing {
		fmt.Fprintln(os.Stderr, "option '--solution' requires an argument")
	} else {
		fmt.Fprintf(os.Stderr, "unknown solution kind %q\n", unknown)
	}
	fmt.Fprintf(os.Stderr, "Valid arguments: %s\n", formatDrillKinds(track))
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func printUnifiedHeader(track drillTrack) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              DAILY PRACTICE — UNIFIED                    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	switch track {
	case trackDSA:
		fmt.Println("Track: DSA — Core 5 + reflex writing specialty")
	case trackRead:
		fmt.Println("Track: DSA reflex reading — today's specialty")
	case trackWrite:
		fmt.Println("Track: DSA writing — Core 5 + today's reflex specialty")
	case trackLeetcode:
		fmt.Println("Track: LeetCode — 10 full problems matching today's reflex topic")
	case trackBackend:
		fmt.Println("Track: Backend interview — Core 5 explain/write + resume block")
	}
	fmt.Println()
}

func printDSAExtras() {
	fmt.Println()
	fmt.Println("━━━ MATH (daily add-on) ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  Write: go run ./bin/study_play -- --run-math")
	fmt.Println("  Guide: doc/write/MATH_CONCEPTS.md")
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

func runCore5In(root string) error {
	core5Dir := filepath.Join(root, "drills", "write", "core5")
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = core5Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCore5(root string) int {
	if err := core5Runner(root); err != nil {
		return 1
	}
	return 0
}

func runUnified(root string, opts dailyOptions) int {
	if !isKnownTrack(opts.track) {
		printUnknownTrack(opts.track)
		return 1
	}

	drillKind := opts.drillKind
	solutionKind := opts.solutionKind
	briefArgs := withBrief(opts.passArgs)

	switch opts.track {
	case trackDSA:
		if opts.run {
			if opts.runSide == "conflict" {
				printRunSideConflict()
				return 1
			}
			if opts.runSide == "read" {
				printReadUseTrack()
				return 1
			}
			runArgs := opts.passArgs
			if code := runModule(root, "study_play", runArgs, true); code != 0 {
				return code
			}
			return 0
		}
		fmt.Printf("DAILY %s", todayName())
		if drillKind != "" {
			fmt.Printf(" | %s", drillKind)
		} else if solutionKind != "" {
			fmt.Printf(" | solution %s", solutionKind)
		}
		fmt.Println()
		if code := runModule(root, "study_play", briefArgs, false); code != 0 {
			return code
		}
		if drillKind == "" && solutionKind == "" {
			fmt.Println("drill:  go run . -- --drill core")
			fmt.Println("        go run . -- --drill reflex")
			fmt.Println("run:    go run . -- --run core")
			fmt.Println("        go run . -- --run reflex")
			fmt.Println("read:   go run . -- --track read")
			fmt.Println("math:   go run ./bin/study_play -- --run-math")
		}
	case trackRead:
		if code := runModule(root, "study_code", filterReadPassArgs(opts.passArgs), opts.run); code != 0 {
			return code
		}
	case trackWrite:
		if code := runModule(root, "study_play", opts.passArgs, opts.run); code != 0 {
			return code
		}
	case trackLeetcode:
		if opts.run {
			fmt.Fprintln(os.Stderr, "leetcode practice sets are solved on leetcode.com; no local --run")
			fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
			return 1
		}
		if code := runModule(root, "study_leetcode", withBrief(opts.passArgs), false); code != 0 {
			return code
		}
	case trackBackend:
		if code := runModule(root, "study_backend", opts.passArgs, opts.run); code != 0 {
			return code
		}
	}
	return 0
}

func filterReadPassArgs(passArgs []string) []string {
	out := make([]string, 0, len(passArgs))
	for i := 0; i < len(passArgs); i++ {
		a := passArgs[i]
		switch a {
		case "--drill", "--solution":
			if i+1 < len(passArgs) && passArgs[i+1] == "core" {
				i++
				continue
			}
			out = append(out, a)
			if i+1 < len(passArgs) {
				i++
				out = append(out, passArgs[i])
			}
		case "--run":
			out = append(out, a)
			if i+1 < len(passArgs) && passArgs[i+1] == "core" {
				i++
				continue
			}
			if i+1 < len(passArgs) && isDrillKind(passArgs[i+1]) {
				i++
				out = append(out, passArgs[i])
			}
		default:
			out = append(out, a)
		}
	}
	return out
}

func containsAll(s string, parts ...string) bool {
	for _, p := range parts {
		if !strings.Contains(s, p) {
			return false
		}
	}
	return true
}
