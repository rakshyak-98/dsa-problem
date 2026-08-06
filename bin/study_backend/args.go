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

func parseBackendArgs(args []string) (drillKind string, catalog, cram, setup bool, runMode string, drillErr bool) {
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--drill":
			if i+1 >= len(args) {
				printDrillArgError(true, "")
				return "", catalog, cram, setup, runMode, true
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				printDrillArgError(false, kind)
				return "", catalog, cram, setup, runMode, true
			}
			i++
			drillKind = kind
		case "--catalog":
			catalog = true
		case "--cram":
			cram = true
		case "--setup":
			setup = true
		case "--run":
			if i+1 < len(args) && isRunKind(args[i+1]) {
				i++
				runMode = args[i]
			} else {
				runMode = "all"
			}
		}
	}
	return drillKind, catalog, cram, setup, runMode, false
}
