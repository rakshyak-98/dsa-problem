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

func withBrief(args []string) []string {
	if hasArg(args, "--brief") {
		return args
	}
	return append(args, "--brief")
}

func printTrackList() {
	for _, t := range availableTracks {
		fmt.Printf("  %-10s%s\n", t.name, t.description)
	}
}

func printNeedsArg(opt string) {
	fmt.Fprintf(os.Stderr, "%s: option '%s' requires an argument\n", dailyProg, opt)
	fmt.Fprintln(os.Stderr, gnuTryHelp(dailyProg))
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Run the daily DSA practice session from the repo root. With no options it
prints today's plan for the default track.

Long options may be abbreviated while unambiguous, and any option that takes
a value also accepts the --option=value form.

  -h, --help               display this help and exit
  -V, --version            display version information and exit
      --list-tracks        list the practice tracks and exit

  -t, --track=NAME         select the practice track (default: dsa)
                             dsa       Core 5 + today's reflex writing specialty
                             read      reflex code-reading drills
                             write     writing drills only
                             leetcode  daily 10-question practice set
      --drill[=KIND]       show a drill plan (KIND: core or reflex; default reflex)
      --solution[=KIND]    show the matching solution file (KIND: core or reflex)
      --run[=KIND]         run drill tests and log the result
                             (KIND: core, reflex, or leetcode)
      --core5              run the standalone Core 5 write drill
      --catalog            list every drill in the active track

Writing track (dsa / write):
      --refresh            today's level-matched session (recognise then rebuild)
      --show               with --refresh, reveal the recognition answers
      --levels             what each function has earned: L1 / L2 / L3
      --problems           the curated primary problem per function, by level
      --triggers           the full cross-topic pattern-trigger table
      --weak               the weakest functions in the drill log

LeetCode track (--track=leetcode):
      --run                fetch and show today's 10 problems
      --refresh            re-fetch from the LeetCode API and rewrite daily.json
      --catalog            list every weekday practice set

Read track (--track=read):
      --drill[=reflex]     show today's reflex reading plan
      --solution[=reflex]  show the reading answer-key section
      --run[=reflex]       run today's reflex reading tests

Deprecated (still accepted): --read, --write, --leetcode as run-side selectors
(-r, -w, -l) are superseded by --track.

Exit status: 0 success, 1 a drill failed, 2 a command-line usage error.
`)
}

func printUnknownTrack(track drillTrack) {
	fmt.Fprintf(os.Stderr, "unknown track %q\n", track)
	fmt.Fprint(os.Stderr, "Valid tracks: dsa, read, write, leetcode\n")
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
	fmt.Fprintln(os.Stderr, "cannot combine -r/--read, -w/--write, and -l/--leetcode on the same --run")
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
	}
	fmt.Println()
}

func printDSAExtras() {
	fmt.Println()
	fmt.Println("━━━ VARIANTS (optional stretch) ━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  go run -C drills/write/variants .")
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

// catalogModule maps a track to the drill module that owns its --catalog listing.
func catalogModule(track drillTrack) string {
	switch track {
	case trackRead:
		return "study_code"
	case trackLeetcode:
		return "study_leetcode"
	default:
		return "study_play"
	}
}

func runUnified(root string, opts dailyOptions) int {
	if !isKnownTrack(opts.track) {
		printUnknownTrack(opts.track)
		return 1
	}

	// --catalog is a terminal listing: delegate straight to the active track's
	// module with no other args, skipping the daily header/footer scaffolding.
	if opts.catalog {
		return runModule(root, catalogModule(opts.track), []string{"--catalog"}, true)
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
			if opts.runSide == "leetcode" {
				if code := runLeetcode(root, opts.passArgs, true); code != 0 {
					return code
				}
				return 0
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
			fmt.Println("        go run . -- --run leetcode")
			fmt.Println("        go run . -- --run -l")
			fmt.Println("read:   go run . -- --track read")
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
			if code := runLeetcode(root, opts.passArgs, true); code != 0 {
				return code
			}
			return 0
		}
		if code := runModule(root, "study_leetcode", withBrief(opts.passArgs), false); code != 0 {
			return code
		}
	}
	return 0
}

func runLeetcode(root string, passArgs []string, run bool) int {
	args := filterLeetcodePassArgs(passArgs)
	if run && !hasArg(args, "--run") {
		args = append(args, "--run")
	}
	return runModule(root, "study_leetcode", args, run)
}

func filterLeetcodePassArgs(passArgs []string) []string {
	out := make([]string, 0, len(passArgs))
	for i := 0; i < len(passArgs); i++ {
		a := passArgs[i]
		switch a {
		case "-r", "--read", "-w", "--write", "-l", "--leetcode":
			continue
		case "--run":
			out = append(out, a)
			if i+1 < len(passArgs) && isRunTarget(passArgs[i+1]) {
				i++
			}
		default:
			out = append(out, a)
		}
	}
	return out
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
