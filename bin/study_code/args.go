package main

import (
	"fmt"
	"os"
)

const readProg = "study_code"

var readLongOpts = []string{
	"help", "version",
	"drill", "solution", "run", "catalog", "brief",
	"read", "write", // consumed by the root runner when picking a side
}

var readAliases = map[string][]string{}

func readSpec() gnuSpec {
	return gnuSpec{prog: readProg, canonical: readLongOpts, aliases: readAliases}
}

func isRunKind(s string) bool {
	return s == "reflex"
}

func isDrillKind(s string) bool {
	return isRunKind(s)
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Reflex code-reading drills for today's weekday specialty. With no option it
prints today's plan.

Long options may be abbreviated while unambiguous, and every value-taking
option also accepts the --option=value form. The only KIND is "reflex", so
the value may be omitted.

  -h, --help            display this help and exit
  -V, --version         display version information and exit
      --drill[=reflex]  show today's reflex reading plan
      --solution[=reflex]
                        show the reading answer-key section for today
      --run[=reflex]    run today's reflex reading tests
      --catalog         list the weekday reading drills
      --brief           one-line output for the unified daily runner

Exit status: 0 success, 1 a drill failed, 2 a command-line usage error.
`)
}

func printCoreReadRemoved() {
	fmt.Fprintf(os.Stderr, "%s: core reading drills were removed; use --run (reflex is implied)\n", readProg)
	fmt.Fprintln(os.Stderr, gnuTryHelp(readProg))
}

// readOpts is the fully parsed command line for the read-drill CLI.
type readOpts struct {
	drillKind    string // "reflex" | ""
	solutionKind string // "reflex" | ""
	runMode      string // "reflex" | ""
	catalog      bool
	brief        bool
}

// parseRead applies the GNU front-end, then reads the normalised tokens.
func parseRead(args []string) (opts readOpts, ctl gnuCtl, parseErr bool) {
	norm, ctl := gnuPre(args, readSpec())
	if ctl != gnuOK {
		return opts, ctl, false
	}
	// value for --drill / --solution / --run: an attached KIND token, else the
	// implied default "reflex". "core" is a removed spelling and is rejected.
	kindValue := func(i *int) (string, bool) {
		if *i+1 < len(norm) {
			next := norm[*i+1]
			if next == "core" {
				printCoreReadRemoved()
				return "", false
			}
			if isRunKind(next) {
				*i++
				return next, true
			}
		}
		return "reflex", true
	}
	for i := 0; i < len(norm); i++ {
		switch norm[i] {
		case "--catalog":
			opts.catalog = true
		case "--brief":
			opts.brief = true
		case "--read", "--write", "-r", "-w":
			// selected by the root runner; nothing to do here
		case "--drill":
			kind, ok := kindValue(&i)
			if !ok {
				return opts, gnuOK, true
			}
			opts.drillKind = kind
		case "--solution":
			kind, ok := kindValue(&i)
			if !ok {
				return opts, gnuOK, true
			}
			opts.solutionKind = kind
		case "--run":
			kind, ok := kindValue(&i)
			if !ok {
				return opts, gnuOK, true
			}
			opts.runMode = kind
		}
	}
	return opts, gnuOK, false
}
