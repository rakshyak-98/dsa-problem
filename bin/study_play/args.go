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
	if missing {
		fmt.Fprintln(os.Stderr, "option '--drill' requires an argument")
	} else {
		fmt.Fprintf(os.Stderr, "unknown drill kind %q\n", unknown)
	}
	fmt.Fprintln(os.Stderr, "Valid arguments: core, reflex")
	fmt.Fprintln(os.Stderr, "Try 'go run . -- --help' for more information.")
}

func parsePlayArgs(args []string) (drillKind string, brief, runMath bool, runMode string, drillErr bool) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--drill":
			if i+1 >= len(args) {
				printDrillArgError(true, "")
				return "", brief, runMath, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return "", brief, runMath, runMode, true
			}
			i++
			drillKind = kind
		case "--brief":
			brief = true
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
	return drillKind, brief, runMath, runMode, false
}
