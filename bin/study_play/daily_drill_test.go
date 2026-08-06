package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestHasFlag(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()
	os.Args = []string{"daily_drill", "--", "--run", "--weak"}
	if !hasFlag("--run") || !hasFlag("--weak") || hasFlag("--nope") {
		t.Fatal("hasFlag")
	}
}

func TestDrillRotation(t *testing.T) {
	if len(drills) != 7 {
		t.Fatalf("expected 7 drills, got %d", len(drills))
	}
	for i, d := range drills {
		if d.file == "" || d.day == "" {
			t.Fatalf("drill %d incomplete", i)
		}
	}
}

func TestPrintFunctionsNoPanic(t *testing.T) {
	capture := func(fn func()) string {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		fn()
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		return buf.String()
	}
	out := capture(func() { printCatalog() })
	if len(out) < 50 {
		t.Fatal("printCatalog empty")
	}
	out = capture(func() { printToday(drills[0], false) })
	if len(out) < 30 {
		t.Fatal("printToday empty")
	}
	out = capture(func() { printDrill(drills[0], false) })
	if len(out) < 20 {
		t.Fatal("printDrill empty")
	}
}

func TestEssentialCatalog(t *testing.T) {
	if len(essentialCatalog) < 7 {
		t.Fatal("catalog too small")
	}
}

func TestCore5Metadata(t *testing.T) {
	if len(core5) != 5 {
		t.Fatal("core5 length")
	}
	if len(allTriggers) < 10 {
		t.Fatal("triggers list")
	}
}
func TestBonusDrills(t *testing.T) {
	if len(bonusDrills) != 3 {
		t.Fatal("bonus drills")
	}
}
