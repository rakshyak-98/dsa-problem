package main

import (
	"fmt"
	"os"
)

func isRunKind(s string) bool {
	return s == "core" || s == "reflex"
}

func isDrillKind(s string) bool {
	return isRunKind(s)
}

func printHelp() {
	fmt.Print(`Usage: go run ./bin/study_play -- [OPTION]...

Reflex writing drills (Core 5 + today's weekday specialty).

Options:
  -h, --help               display this help message and exit
      --drill core|reflex  show Core 5 or today's reflex write plan
      --solution core|reflex
                           show solution file path
      --run [core|reflex]  run write drill tests (default: core + reflex)
      --run-core5          run Core 5 tests only
      --run-math           run math reflex add-on drill
      --catalog            list weekday write drills
      --brief              one-line output for unified daily runner
      --weak               show weakest functions from tracker log
      --setup              generate drills from blank templates
      --reset              reset today's drill to blank template

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
		fmt.Fprintf(os.Stderr, "option %q requires an argument\n", flag)
	} else {
		fmt.Fprintf(os.Stderr, "unknown %s %q\n", label, unknown)
	}
	fmt.Fprintln(os.Stderr, "Valid arguments: core, reflex")
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_play -- --help' for more information.")
}

func parsePlayArgs(args []string) (drillKind, solutionKind string, help, brief, runMath bool, runMode string, parseErr bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			help = true
		case "--drill":
			if i+1 >= len(args) {
				printDrillArgError(true, "")
				return "", "", help, brief, runMath, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return "", "", help, brief, runMath, runMode, true
			}
			i++
			drillKind = kind
		case "--solution":
			if i+1 >= len(args) {
				printSolutionArgError(true, "")
				return "", "", help, brief, runMath, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printSolutionArgError(false, kind)
				return "", "", help, brief, runMath, runMode, true
			}
			i++
			solutionKind = kind
		case "--brief":
			brief = true
		case "-r", "--read", "-w", "--write":
			// consumed by root CLI when selecting read/write side
		case "--run-math":
			runMath = true
		case "--run-core5":
			runMode = "core"
		case "--run":
			if i+1 < len(args) && isRunKind(args[i+1]) {
				i++
				runMode = args[i]
			} else {
				runMode = "all"
			}
		}
	}
	return drillKind, solutionKind, help, brief, runMath, runMode, false
}
