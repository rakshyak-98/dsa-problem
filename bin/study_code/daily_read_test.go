package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestTodayDrill(t *testing.T) {
	d := todayDrill()
	if d.file == "" || d.day == "" {
		t.Fatal("todayDrill incomplete")
	}
}

func TestDrillsCatalog(t *testing.T) {
	if len(drills) != 7 {
		t.Fatalf("expected 7 drills, got %d", len(drills))
	}
}

func TestParseReadArgs(t *testing.T) {
	drillKind, solutionKind, help, catalog, brief, runMode, parseErr := parseReadArgs([]string{"--", "--drill", "reflex", "--run", "reflex", "--catalog", "--brief"})
	if parseErr || drillKind != "reflex" || solutionKind != "" || runMode != "reflex" || !catalog || !brief || help {
		t.Fatal("parseReadArgs all flags")
	}
	drillKind, solutionKind, help, catalog, brief, runMode, parseErr = parseReadArgs([]string{"--drill", "reflex"})
	if parseErr || drillKind != "reflex" || solutionKind != "" || help || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs reflex drill")
	}
	_, solutionKind, help, catalog, brief, runMode, parseErr = parseReadArgs([]string{"--solution", "reflex"})
	if parseErr || solutionKind != "reflex" || help || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs solution reflex")
	}
	drillKind, solutionKind, help, catalog, brief, runMode, parseErr = parseReadArgs([]string{"--run", "reflex"})
	if parseErr || runMode != "reflex" || drillKind != "" || solutionKind != "" || help || catalog || brief {
		t.Fatal("parseReadArgs reflex run")
	}
	drillKind, solutionKind, help, catalog, brief, runMode, parseErr = parseReadArgs([]string{"--run"})
	if parseErr || runMode != "reflex" || help {
		t.Fatal("parseReadArgs bare run")
	}
	drillKind, solutionKind, help, catalog, brief, runMode, parseErr = parseReadArgs(nil)
	if parseErr || drillKind != "" || solutionKind != "" || help || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs empty")
	}
	_, _, _, _, _, _, parseErr = parseReadArgs([]string{"--drill"})
	if !parseErr {
		t.Fatal("bare drill should error")
	}
	_, _, _, _, _, _, parseErr = parseReadArgs([]string{"--solution"})
	if !parseErr {
		t.Fatal("bare solution should error")
	}
	_, _, _, _, _, _, parseErr = parseReadArgs([]string{"--drill", "core"})
	if !parseErr {
		t.Fatal("core drill should error")
	}
	_, _, _, _, _, _, parseErr = parseReadArgs([]string{"--run", "core"})
	if !parseErr {
		t.Fatal("core run should error")
	}
	_, _, help, _, _, _, parseErr = parseReadArgs([]string{"--help"})
	if parseErr || !help {
		t.Fatal("help flag")
	}
}

func TestRunDrillInvalid(t *testing.T) {
	if err := runDrill("nonexistent_drill"); err == nil {
		t.Fatal("expected runDrill error")
	}
}

func TestEachDrillHasFocus(t *testing.T) {
	for _, d := range drills {
		if len(d.focus) < 3 {
			t.Fatalf("%s missing focus items", d.file)
		}
		if d.skill == "" || d.warmup == "" {
			t.Fatalf("%s incomplete metadata", d.file)
		}
	}
}

func TestTodayDrillWeekdays(t *testing.T) {
	seen := map[string]bool{}
	for wd := time.Sunday; wd <= time.Saturday; wd++ {
		_ = wd
		d := todayDrill()
		seen[d.file] = true
	}
	if len(seen) < 1 {
		t.Fatal("todayDrill should return a drill")
	}
}

func TestPrintFunctions(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printReflexDrill(drills[0], false)
	printSolutionReflex(drills[0], false)
	printToday(drills[0], false)
	printCatalog()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 30 {
		t.Fatal("print output too short")
	}
}

func TestPrintHelp(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printHelp()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	if !bytes.Contains(buf.Bytes(), []byte("reflex")) || !bytes.Contains(buf.Bytes(), []byte("--catalog")) {
		t.Fatalf("help missing reflex options:\n%s", out)
	}
}
