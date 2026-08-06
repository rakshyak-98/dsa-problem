package main

import (
	"fmt"
	"os"
)

func runStudyCode(drillKind string, runMath, catalog, brief bool, runMode string) int {
	if catalog {
		printCatalog()
		return 0
	}
	if drillKind == "core" {
		printDrill(brief)
		return 0
	}
	if drillKind == "reflex" {
		printReflexDrill(todayDrill(), brief)
		return 0
	}

	d := todayDrill()

	if runMode == "" {
		printToday(d, brief)
	}

	if runMath {
		if err := runDrill(mathReadFile); err != nil {
			fmt.Fprintf(os.Stderr, "math read failed: %v\n", err)
			return 1
		}
		return 0
	}

	switch runMode {
	case "core":
		if err := runDrill("00_core_read"); err != nil {
			fmt.Fprintf(os.Stderr, "core failed: %v\n", err)
			return 1
		}
	case "reflex":
		if err := runDrill(d.file); err != nil {
			fmt.Fprintf(os.Stderr, "specialty failed: %v\n", err)
			return 1
		}
	case "all":
		if err := runDrill("00_core_read"); err != nil {
			fmt.Fprintf(os.Stderr, "core failed: %v\n", err)
			return 1
		}
		if err := runDrill(d.file); err != nil {
			fmt.Fprintf(os.Stderr, "specialty failed: %v\n", err)
			return 1
		}
	}
	return 0
}
