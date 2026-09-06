package main

import (
	"fmt"
	"os"
)

const playProg = "study_play"

// playLongOpts is every long option this drill CLI accepts. Anything else that
// starts with "--" is reported as an error instead of being silently ignored.
var playLongOpts = []string{
	"help", "version",
	"drill", "solution", "run",
	"refresh", "show", "levels", "problems", "triggers",
	"catalog", "brief", "weak", "setup", "reset",
	"read", "write", // consumed by the root runner when picking a side
}

// playAliases rewrites retired spellings to the current option.
var playAliases = map[string][]string{
	"run-core5": {"--run", "core"}, // superseded by: --run=core
}

func playSpec() gnuSpec {
	return gnuSpec{prog: playProg, canonical: playLongOpts, aliases: playAliases}
}

func isRunKind(s string) bool {
	return s == "core" || s == "reflex"
}

func isDrillKind(s string) bool {
	return isRunKind(s)
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Reflex writing drills: Core 5 plus today's weekday specialty. With no option
it prints today's plan.

Long options may be abbreviated while unambiguous, and every value-taking
option also accepts the --option=value form.

  -h, --help            display this help and exit
  -V, --version         display version information and exit

      --drill=KIND      show a drill plan (KIND: core or reflex)
      --solution=KIND   show the solution file path (KIND: core or reflex)
      --run[=KIND]      run drill tests and log the result
                          core    the Core 5 tests
                          reflex  today's specialty tests
                          (default: both)
      --catalog         list the weekday write drills
      --brief           one-line output for the unified daily runner

Session views (writing track):
      --refresh         today's level-matched session (recognise then rebuild)
      --show            with --refresh, reveal the recognition answers
      --levels          what each function has earned: L1 / L2 / L3
      --problems        the curated primary problem per function, by level
      --triggers        the full cross-topic pattern-trigger table
      --weak            the weakest functions in the drill log

Maintenance:
      --setup           scaffold every drill from the blank templates
      --reset           reset today's drill to its blank template

Deprecated (still accepted): --run-core5 is now --run=core.

Exit status: 0 success, 1 a drill failed, 2 a command-line usage error.
`)
}

func printDrillArgError(missing bool, unknown string) {
	printKindArgError("--drill", "drill kind", missing, unknown)
}

func printSolutionArgError(missing bool, unknown string) {
	printKindArgError("--solution", "solution kind", missing, unknown)
}

func printKindArgError(flag, label string, missing bool, unknown string) {
	if missing {
		fmt.Fprintf(os.Stderr, "%s: option '%s' requires an argument\n", playProg, flag)
	} else {
		fmt.Fprintf(os.Stderr, "%s: unknown %s %q\n", playProg, label, unknown)
	}
	fmt.Fprintln(os.Stderr, "valid arguments: core, reflex")
	fmt.Fprintln(os.Stderr, gnuTryHelp(playProg))
}

// playOpts is the fully parsed command line for the write-drill CLI.
type playOpts struct {
	brief        bool
	drillKind    string // "core" | "reflex" | ""
	solutionKind string // "core" | "reflex" | ""
	runMode      string // "core" | "reflex" | "all" | ""
	refresh      bool
	show         bool
	levels       bool
	problems     bool
	triggers     bool
	catalog      bool
	weak         bool
	setup        bool
	reset        bool
}

// parsePlay applies the GNU front-end, then reads the normalised tokens. The
// returned gnuCtl tells main whether to print help/version or exit on a usage
// error; parseErr covers a bad KIND value caught here.
func parsePlay(args []string) (opts playOpts, ctl gnuCtl, parseErr bool) {
	norm, ctl := gnuPre(args, playSpec())
	if ctl != gnuOK {
		return opts, ctl, false
	}
	for i := 0; i < len(norm); i++ {
		switch norm[i] {
		case "--brief":
			opts.brief = true
		case "--read", "--write", "-r", "-w":
			// selected by the root runner; nothing to do here
		case "--refresh":
			opts.refresh = true
		case "--show":
			opts.show = true
		case "--levels":
			opts.levels = true
		case "--problems":
			opts.problems = true
		case "--triggers":
			opts.triggers = true
		case "--catalog":
			opts.catalog = true
		case "--weak":
			opts.weak = true
		case "--setup":
			opts.setup = true
		case "--reset":
			opts.reset = true
		case "--drill":
			kind, ok := takeKind(norm, &i)
			if !ok {
				printDrillArgError(true, "")
				return opts, gnuOK, true
			}
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return opts, gnuOK, true
			}
			opts.drillKind = kind
		case "--solution":
			kind, ok := takeKind(norm, &i)
			if !ok {
				printSolutionArgError(true, "")
				return opts, gnuOK, true
			}
			if !isDrillKind(kind) {
				printSolutionArgError(false, kind)
				return opts, gnuOK, true
			}
			opts.solutionKind = kind
		case "--run":
			if i+1 < len(norm) && isRunKind(norm[i+1]) {
				i++
				opts.runMode = norm[i]
			} else {
				opts.runMode = "all"
			}
		}
	}
	return opts, gnuOK, false
}

// takeKind consumes the token after a value-taking option, if present.
func takeKind(args []string, i *int) (string, bool) {
	if *i+1 >= len(args) {
		return "", false
	}
	*i++
	return args[*i], true
}
