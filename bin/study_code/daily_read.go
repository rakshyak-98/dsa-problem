// Daily code-reading helper — Core Read 3 + specialty drill
//
// RUN:              go run .
// RUN with checks:  go run . -- --run
// Math add-on:      go run . -- --run-math
// Core only:        go run . -- --drill
// Full catalog:     go run . -- --catalog
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type coreItem struct {
	skill  string
	prompt string
	sec    int
}

type drill struct {
	day    string
	file   string
	skill  string
	focus  []string
	warmup string
}

var core3 = []coreItem{
	{"Signature + state", "Name in/out and what each binder means across the loop", 90},
	{"Trace one sample", "Build a step table until you can predict the return", 120},
	{"Pattern + bound", "Name the template and give O(time) / O(space)", 60},
}

var drills = []drill{
	{
		day: "Monday", file: "01_scan_structure",
		skill: "Structure scan",
		focus: []string{
			"params / return type / mutation",
			"loop nest vs single pass",
			"early returns and base cases",
		},
		warmup: "Before logic: mark every loop and every return.",
	},
	{
		day: "Tuesday", file: "02_trace_execution",
		skill: "Hand trace",
		focus: []string{
			"pick the given sample",
			"update binders each step",
			"predict return before checking",
		},
		warmup: "If the table disagrees with your gut, trust the table.",
	},
	{
		day: "Wednesday", file: "03_name_the_pattern",
		skill: "Pattern from shape",
		focus: []string{
			"ignore misleading names",
			"match skeleton to template",
			"one breath label",
		},
		warmup: "Shape beats function name — names may be wrong on purpose.",
	},
	{
		day: "Thursday", file: "04_find_the_bug",
		skill: "Bug hunt",
		focus: []string{
			"off-by-one on bounds",
			"wrong pointer move",
			"missing visit / wrong compare",
		},
		warmup: "Assume one bug. Find the line that breaks the invariant.",
	},
	{
		day: "Friday", file: "05_complexity_glance",
		skill: "Complexity from shape",
		focus: []string{
			"nested vs amortized two-pointer",
			"map/set space",
			"recursion depth",
		},
		warmup: "Count how many times the inner work runs across the whole input.",
	},
	{
		day: "Saturday", file: "06_reconstruct_ask",
		skill: "Ask from code",
		focus: []string{
			"one sentence, no jargon",
			"index vs value vs count",
			"contiguous vs not",
		},
		warmup: "If a junior couldn't understand your sentence, rewrite it.",
	},
	{
		day: "Sunday", file: "07_compare_variants",
		skill: "Compare two solutions",
		focus: []string{
			"same ask, different structure",
			"name the tradeoff (time/space/clarity)",
			"when you'd pick each",
		},
		warmup: "Diff the skeletons first; ignore shared boilerplate.",
	},
}

const mathReadFile = "08_math_concepts"

func todayDrill() drill {
	wd := int(time.Now().Weekday())
	idx := (wd + 6) % 7
	return drills[idx]
}

func drillOpenPath(file string) string {
	if file == "00_core_read" {
		return "drills/read/core/00_core_read/main.go"
	}
	return fmt.Sprintf("drills/read/weekday/%s/main.go", file)
}

func printDrill(brief bool) {
	if brief {
		fmt.Println("read:  drills/read/core/00_core_read/")
		return
	}
	fmt.Println("READ | core 3")
	fmt.Println("path: drills/read/core/00_core_read/")
}

func printToday(d drill, brief bool) {
	if brief {
		fmt.Printf("read:  %s\n", d.file)
		return
	}
	fmt.Printf("READ %s | %s — %s\n", d.day, d.file, d.skill)
	fmt.Printf("path: %s\n", drillOpenPath(d.file))
	fmt.Println("run:    go run . -- --run core")
	fmt.Println("        go run . -- --run reflex")
}

func printCatalog() {
	fmt.Println("READ catalog")
	for _, d := range drills {
		fmt.Printf("%-9s  %-22s  %s\n", d.day, d.file, d.skill)
	}
	fmt.Printf("%-9s  %-22s  %s\n", "Daily", mathReadFile, "Math formulas & traces")
}

func runDrill(file string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := readDrillDir(findRepoRoot(cwd), file)
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	drill, runMath, catalog, brief, runMode := parseReadArgs(os.Args[1:])
	if code := runStudyCode(drill, runMath, catalog, brief, runMode); code != 0 {
		os.Exit(code)
	}
}
