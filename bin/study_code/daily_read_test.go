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
	drillKind, runMath, catalog, brief, runMode, drillErr := parseReadArgs([]string{"--", "--drill", "core", "--run", "core", "--run-math", "--catalog", "--brief"})
	if drillErr || drillKind != "core" || runMode != "core" || !runMath || !catalog || !brief {
		t.Fatal("parseReadArgs all flags")
	}
	drillKind, runMath, catalog, brief, runMode, drillErr = parseReadArgs([]string{"--drill", "reflex"})
	if drillErr || drillKind != "reflex" || runMath || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs reflex drill")
	}
	drillKind, runMath, catalog, brief, runMode, drillErr = parseReadArgs([]string{"--run", "reflex"})
	if drillErr || runMode != "reflex" || drillKind != "" || runMath || catalog || brief {
		t.Fatal("parseReadArgs reflex run")
	}
	drillKind, runMath, catalog, brief, runMode, drillErr = parseReadArgs([]string{"--run"})
	if drillErr || runMode != "all" {
		t.Fatal("parseReadArgs bare run")
	}
	drillKind, runMath, catalog, brief, runMode, drillErr = parseReadArgs(nil)
	if drillErr || drillKind != "" || runMath || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs empty")
	}
	_, _, _, _, _, drillErr = parseReadArgs([]string{"--drill"})
	if !drillErr {
		t.Fatal("bare drill should error")
	}
}

func TestRunDrillInvalid(t *testing.T) {
	if err := runDrill("nonexistent_drill"); err == nil {
		t.Fatal("expected runDrill error")
	}
}

func TestCore3Items(t *testing.T) {
	if len(core3) != 3 {
		t.Fatal("core3 length")
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
	printDrill(false)
	printReflexDrill(drills[0], false)
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
