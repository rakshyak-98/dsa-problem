package main

import (
	"fmt"
	"os"
)

func runStudyCode(micro, run, catalog, specialty bool) int {
	if catalog {
		printCatalog()
		return 0
	}

	if specialty {
		if micro {
			return 0
		}
		d := todayDrill()
		printSpecialty(d)
		if run {
			fmt.Printf("── Running specialty %s ──\n", d.file)
			if err := runDrill(d.file); err != nil {
				fmt.Fprintf(os.Stderr, "specialty failed: %v\n", err)
				return 1
			}
		}
		return 0
	}

	fmt.Println()
	fmt.Println("study_code — code reading drills")
	fmt.Println("Fill TODO: READ answers, then go run . -- --run")
	fmt.Println()

	printCore()
	if micro {
		fmt.Println("(--micro) Specialty skipped. Still counts as Minimum tier.")
		return 0
	}

	d := todayDrill()
	printSpecialty(d)

	if run {
		fmt.Println("── Running Core Read 3 ──")
		if err := runDrill("00_core_read"); err != nil {
			fmt.Fprintf(os.Stderr, "core failed: %v\n", err)
			return 1
		}
		fmt.Println()
		fmt.Printf("── Running specialty %s ──\n", d.file)
		if err := runDrill(d.file); err != nil {
			fmt.Fprintf(os.Stderr, "specialty failed: %v\n", err)
			return 1
		}
	}
	return 0
}
