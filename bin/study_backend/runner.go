package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runBackend(micro, run, catalog, cram, setup bool) int {
	if setup {
		if err := runSetup(); err != nil {
			fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
			return 1
		}
		return 0
	}
	if catalog {
		printCatalog()
		return 0
	}
	if cram {
		printCramPlan()
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("study_backend — interview prep from your resume")
	fmt.Println("Fill TODO: EXPLAIN / TODO: REFLEX, then go run . -- --run")
	fmt.Println()

	printCoreExplain()
	printCoreWrite()

	if micro {
		fmt.Println("(--micro) Blocks skipped. Still counts as Minimum tier.")
		if run {
			if err := runCoreDrills(); err != nil {
				return 1
			}
		}
		return 0
	}

	b := todayBlock()
	printBlock(b)

	if run {
		if err := runCoreDrills(); err != nil {
			return 1
		}
		fmt.Printf("── Running block %s ──\n", b.file)
		if err := runExplainBlock(b.file); err != nil {
			return 1
		}
	}
	return 0
}

func runCoreDrills() error {
	root := findRepoRoot(mustCwd())
	fmt.Println("── Core 5 EXPLAIN ──")
	if err := goRun(explainCore5Dir(root)); err != nil {
		fmt.Fprintf(os.Stderr, "explain core5 failed: %v\n", err)
		return err
	}
	fmt.Println()
	fmt.Println("── Core 5 WRITE (Go) ──")
	if err := goRun(writeCore5Dir(root)); err != nil {
		fmt.Fprintf(os.Stderr, "write core5 failed: %v\n", err)
		return err
	}
	return nil
}

func runExplainBlock(block string) error {
	return goRun(explainBlockDir(findRepoRoot(mustCwd()), block))
}

func goRun(dir string) error {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func mustCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return cwd
}
