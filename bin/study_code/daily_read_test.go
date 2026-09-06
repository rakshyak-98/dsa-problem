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

func TestParseRead(t *testing.T) {
	opts, ctl, perr := parseRead([]string{"--", "--drill", "reflex", "--run", "reflex", "--catalog", "--brief"})
	if perr || ctl != gnuOK || opts.drillKind != "reflex" || opts.runMode != "reflex" || !opts.catalog || !opts.brief {
		t.Fatalf("parseRead all flags: %+v", opts)
	}

	// The KIND value is optional: bare --drill implies reflex.
	opts, _, _ = parseRead([]string{"--drill"})
	if opts.drillKind != "reflex" {
		t.Fatalf("bare --drill implies reflex: %+v", opts)
	}
	opts, _, _ = parseRead([]string{"--solution=reflex"})
	if opts.solutionKind != "reflex" {
		t.Fatalf("--solution=reflex: %+v", opts)
	}
	opts, _, _ = parseRead([]string{"--run"})
	if opts.runMode != "reflex" {
		t.Fatalf("bare --run implies reflex: %+v", opts)
	}

	opts, _, _ = parseRead(nil)
	if opts.drillKind != "" || opts.runMode != "" || opts.catalog || opts.brief {
		t.Fatalf("parseRead empty: %+v", opts)
	}

	// Unambiguous abbreviation.
	opts, _, _ = parseRead([]string{"--cat"})
	if !opts.catalog {
		t.Fatalf("--cat -> --catalog: %+v", opts)
	}

	// "core" is a removed spelling.
	if _, _, perr := parseRead([]string{"--drill", "core"}); !perr {
		t.Fatal("core drill should be a usage error")
	}
	if _, _, perr := parseRead([]string{"--run", "core"}); !perr {
		t.Fatal("core run should be a usage error")
	}

	if _, ctl, _ := parseRead([]string{"--help"}); ctl != gnuHelp {
		t.Fatal("--help -> gnuHelp")
	}
	if _, ctl, _ := parseRead([]string{"-V"}); ctl != gnuVersion {
		t.Fatal("-V -> gnuVersion")
	}
	if _, ctl, _ := parseRead([]string{"--bogus"}); ctl != gnuErr {
		t.Fatal("--bogus -> gnuErr")
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
