package main

import (
	"fmt"
	"os"
)

func runStudyCode(drillKind, solutionKind string, catalog, brief bool, runMode string) int {
	if catalog {
		printCatalog()
		return 0
	}
	if drillKind == "reflex" {
		printReflexDrill(todayDrill(), brief)
		return 0
	}
	if solutionKind == "reflex" {
		printSolutionReflex(todayDrill(), brief)
		return 0
	}

	d := todayDrill()

	if runMode == "" {
		printToday(d, brief)
		return 0
	}

	if err := runDrill(d.file); err != nil {
		fmt.Fprintf(os.Stderr, "reflex read failed: %v\n", err)
		return 1
	}
	return 0
}
