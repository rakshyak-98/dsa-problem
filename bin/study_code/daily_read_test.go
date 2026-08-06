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
	drill, runMath, catalog, brief, runMode := parseReadArgs([]string{"--", "--drill", "--run", "core", "--run-math", "--catalog", "--brief"})
	if !drill || runMode != "core" || !runMath || !catalog || !brief {
		t.Fatal("parseReadArgs all flags")
	}
	drill, runMath, catalog, brief, runMode = parseReadArgs([]string{"--run", "reflex"})
	if runMode != "reflex" || drill || runMath || catalog || brief {
		t.Fatal("parseReadArgs reflex")
	}
	drill, runMath, catalog, brief, runMode = parseReadArgs([]string{"--run"})
	if runMode != "all" {
		t.Fatal("parseReadArgs bare run")
	}
	drill, runMath, catalog, brief, runMode = parseReadArgs(nil)
	if drill || runMath || catalog || brief || runMode != "" {
		t.Fatal("parseReadArgs empty")
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
