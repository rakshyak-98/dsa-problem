package main

import (
	"fmt"
	"os"
)

func isRunKind(s string) bool {
	return s == "reflex"
}

func isDrillKind(s string) bool {
	return isRunKind(s)
}

func printHelp() {
	fmt.Print(`Usage: go run ./bin/study_code -- [OPTION]...

Run reflex code-reading drills for today's weekday specialty.

Options:
  -h, --help               display this help message and exit
      --drill reflex       show today's reflex read plan
      --solution reflex    show answer key section for today
      --run [reflex]       run today's read drill tests (default: reflex)
      --catalog            list weekday read drills
      --brief              one-line output for unified daily runner

`)
}

func printDrillArgError(missing bool, unknown string) {
	printKindArgError("--drill", "drill kind", missing, unknown)
}

func printSolutionArgError(missing bool, unknown string) {
	printKindArgError("--solution", "solution kind", missing, unknown)
}

func printRunArgError(unknown string) {
	fmt.Fprintf(os.Stderr, "unknown run kind %q\n", unknown)
	fmt.Fprintln(os.Stderr, "Valid arguments: reflex")
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_code -- --help' for more information.")
}

func printCoreReadRemoved() {
	fmt.Fprintln(os.Stderr, "core reading drills were removed; use --run reflex")
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_code -- --help' for more information.")
}

func printKindArgError(flag, label string, missing bool, unknown string) {
	if missing {
		fmt.Fprintf(os.Stderr, "option %q requires an argument\n", flag)
	} else {
		fmt.Fprintf(os.Stderr, "unknown %s %q\n", label, unknown)
	}
	fmt.Fprintln(os.Stderr, "Valid arguments: reflex")
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_code -- --help' for more information.")
}

func parseReadArgs(args []string) (drillKind, solutionKind string, help, catalog, brief bool, runMode string, parseErr bool) {
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
				return "", "", help, catalog, brief, runMode, true
			}
			kind := args[i+1]
			if kind == "core" {
				printCoreReadRemoved()
				return "", "", help, catalog, brief, runMode, true
			}
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return "", "", help, catalog, brief, runMode, true
			}
			i++
			drillKind = kind
		case "--solution":
			if i+1 >= len(args) {
				printSolutionArgError(true, "")
				return "", "", help, catalog, brief, runMode, true
			}
			kind := args[i+1]
			if kind == "core" {
				printCoreReadRemoved()
				return "", "", help, catalog, brief, runMode, true
			}
			if !isDrillKind(kind) {
				printSolutionArgError(false, kind)
				return "", "", help, catalog, brief, runMode, true
			}
			i++
			solutionKind = kind
		case "--catalog":
			catalog = true
		case "--brief":
			brief = true
		case "-r", "--read", "-w", "--write":
			// consumed by root CLI when selecting read/write side
		case "--run":
			if i+1 < len(args) {
				kind := args[i+1]
				if kind == "core" {
					printCoreReadRemoved()
					return "", "", help, catalog, brief, runMode, true
				}
				if isRunKind(kind) {
					i++
					runMode = kind
					continue
				}
				printRunArgError(kind)
				return "", "", help, catalog, brief, runMode, true
			}
			runMode = "reflex"
		}
	}
	return drillKind, solutionKind, help, catalog, brief, runMode, false
}
