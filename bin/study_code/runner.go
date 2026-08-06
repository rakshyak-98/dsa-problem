package main

import (
	"fmt"
	"os"
)

func runStudyCode(drill, run, runMath, catalog, brief bool) int {
	if catalog {
		printCatalog()
		return 0
	}
	if drill {
		printDrill(brief)
		return 0
	}

	d := todayDrill()
	printToday(d, brief)

	if runMath {
		if err := runDrill(mathReadFile); err != nil {
			fmt.Fprintf(os.Stderr, "math read failed: %v\n", err)
			return 1
		}
		return 0
	}

	if run {
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
