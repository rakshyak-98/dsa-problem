package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func coreExplainTopics() string {
	topics := make([]string, len(coreExplain))
	for i, c := range coreExplain {
		topics[i] = c.topic
	}
	return strings.Join(topics, ", ")
}

func coreWriteTopics() string {
	topics := make([]string, len(coreWrite))
	for i, c := range coreWrite {
		topics[i] = c.topic
	}
	return strings.Join(topics, ", ")
}

func runBackend(drillKind string, catalog, cram, setup bool, runMode string) int {
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

	if drillKind == "core" {
		fmt.Println("BACKEND | core 5")
		fmt.Printf("explain: %s\n", coreExplainTopics())
		fmt.Printf("write: %s\n", coreWriteTopics())
		fmt.Println("path: drills/backend/explain/core5/ + drills/backend/write/core5/")
		if runMode == "core" || runMode == "all" {
			if err := runCoreDrills(); err != nil {
				return 1
			}
		}
		return 0
	}
	if drillKind == "reflex" {
		b := todayBlock()
		fmt.Printf("BACKEND %s | %s — %s\n", b.day, b.file, b.topic)
		fmt.Printf("path: drills/backend/explain/blocks/%s/\n", b.file)
		return 0
	}

	b := todayBlock()
	fmt.Printf("BACKEND %s | %s — %s\n", b.day, b.file, b.topic)
	fmt.Printf("path: drills/backend/explain/blocks/%s/\n", b.file)
	fmt.Println("run:    go run . -- --run core")
	fmt.Println("        go run . -- --run reflex")

	switch runMode {
	case "core":
		if err := runCoreDrills(); err != nil {
			return 1
		}
	case "reflex":
		if err := runExplainBlock(b.file); err != nil {
			return 1
		}
	case "all":
		if err := runCoreDrills(); err != nil {
			return 1
		}
		if err := runExplainBlock(b.file); err != nil {
			return 1
		}
	}
	return 0
}

func runCoreDrills() error {
	root := findRepoRoot(mustCwd())
	if err := goRun(explainCore5Dir(root)); err != nil {
		fmt.Fprintf(os.Stderr, "explain core5 failed: %v\n", err)
		return err
	}
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
