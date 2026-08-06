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
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func parseReadArgs(args []string) (drillKind, solutionKind string, runMath, catalog, brief bool, runMode string, parseErr bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--drill":
			if i+1 >= len(args) {
				printDrillArgError(true, "")
				return "", "", runMath, catalog, brief, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return "", "", runMath, catalog, brief, runMode, true
			}
			i++
			drillKind = kind
		case "--solution":
			if i+1 >= len(args) {
				printSolutionArgError(true, "")
				return "", "", runMath, catalog, brief, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printSolutionArgError(false, kind)
				return "", "", runMath, catalog, brief, runMode, true
			}
			i++
			solutionKind = kind
		case "--run-math":
			runMath = true
		case "--catalog":
			catalog = true
		case "--brief":
			brief = true
		case "-r", "--read", "-w", "--write":
			// consumed by root CLI when selecting read/write side
		case "--run":
			if i+1 < len(args) && isRunKind(args[i+1]) {
				i++
				runMode = args[i]
			} else {
				runMode = "all"
			}
		}
	}
	return drillKind, solutionKind, runMath, catalog, brief, runMode, false
}
