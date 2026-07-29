// Daily code-reading helper — Core Read 3 + specialty drill
//
// RUN:              go run .
// RUN with checks:  go run . -- --run
// Core only:        go run . -- --micro
// Full catalog:     go run . -- --catalog
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type coreItem struct {
	skill  string
	prompt string
	sec    int
}

type drill struct {
	day     string
	file    string
	skill   string
	focus   []string
	warmup  string
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

func todayDrill() drill {
	// Go weekday: Sunday=0 … Saturday=6 → map to our table Mon=0…Sun=6
	wd := int(time.Now().Weekday())
	idx := (wd + 6) % 7 // Mon→0 … Sun→6
	return drills[idx]
}

func printCore() {
	fmt.Println("════════════════════════════════════════")
	fmt.Println(" CORE READ 3  (every day)")
	fmt.Println("════════════════════════════════════════")
	for i, c := range core3 {
		fmt.Printf("  %d. [%ds] %s\n     → %s\n", i+1, c.sec, c.skill, c.prompt)
	}
	fmt.Println()
	fmt.Println("  File: drills/00_core_read/")
	fmt.Println("  Method: READING_PATTERNS.md (6 passes)")
	fmt.Println()
}

func printSpecialty(d drill) {
	fmt.Println("════════════════════════════════════════")
	fmt.Printf(" SPECIALTY — %s (%s)\n", d.day, d.file)
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("  Skill: %s\n", d.skill)
	fmt.Printf("  Warm-up: %s\n", d.warmup)
	fmt.Println("  Focus:")
	for _, f := range d.focus {
		fmt.Printf("    • %s\n", f)
	}
	fmt.Printf("\n  Open:  drills/%s/main.go\n", d.file)
	fmt.Printf("  Run:   go run ./drills/%s\n", d.file)
	fmt.Println("  Answers only after fails: drills/answers/")
	fmt.Println()
}

func printCatalog() {
	fmt.Println("ESSENTIAL READING CATALOG")
	fmt.Println("-------------------------")
	for _, d := range drills {
		fmt.Printf("%-9s  %-22s  %s\n", d.day, d.file, d.skill)
	}
	fmt.Println()
	fmt.Println("Core: drills/00_core_read/  |  Method: READING_PATTERNS.md")
}

func runDrill(file string) error {
	dir := filepath.Join("drills", file)
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	args := os.Args[1:]
	// allow `go run . -- --run`
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}

	micro := false
	run := false
	catalog := false
	for _, a := range args {
		switch a {
		case "--micro":
			micro = true
		case "--run":
			run = true
		case "--catalog":
			catalog = true
		}
	}

	if catalog {
		printCatalog()
		return
	}

	fmt.Println()
	fmt.Println("study_code — code reading drills")
	fmt.Println("Fill TODO: READ answers, then go run . -- --run")
	fmt.Println()

	printCore()
	if micro {
		fmt.Println("(--micro) Specialty skipped. Still counts as Minimum tier.")
		return
	}

	d := todayDrill()
	printSpecialty(d)

	if run {
		fmt.Println("── Running Core Read 3 ──")
		if err := runDrill("00_core_read"); err != nil {
			fmt.Fprintf(os.Stderr, "core failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
		fmt.Printf("── Running specialty %s ──\n", d.file)
		if err := runDrill(d.file); err != nil {
			fmt.Fprintf(os.Stderr, "specialty failed: %v\n", err)
			os.Exit(1)
		}
	}
}
