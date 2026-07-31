package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBlankContent(t *testing.T) {
	content, err := blankContent("01_arrays_reflex")
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Fatal("empty blank")
	}
	if _, err := blankContent("nonexistent_drill"); err == nil {
		t.Fatal("expected error")
	}
}

func TestWriteDrillFromBlank(t *testing.T) {
	dir := t.TempDir()
	if err := writeDrillFromBlank("01_arrays_reflex", dir); err != nil {
		t.Fatal(err)
	}
	mainPath := filepath.Join(dir, "main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) < 100 {
		t.Fatal("main.go too short")
	}
}

func TestSetupAllDrills(t *testing.T) {
	dir := t.TempDir()
	if err := setupAllDrills(dir); err != nil {
		t.Fatal(err)
	}
	for _, d := range drills {
		if _, err := os.Stat(filepath.Join(dir, "drills", d.file, "main.go")); err != nil {
			t.Fatalf("missing %s: %v", d.file, err)
		}
	}
	for _, file := range bonusDrills {
		if _, err := os.Stat(filepath.Join(dir, "drills", file, "main.go")); err != nil {
			t.Fatalf("missing bonus %s: %v", file, err)
		}
	}
}

func TestResetTodayDrill(t *testing.T) {
	dir := t.TempDir()
	drillDir := filepath.Join(dir, "drills", "01_arrays_reflex")
	today := drills[0]
	if err := resetTodayDrill(today, drillDir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(drillDir, "main.go")); err != nil {
		t.Fatal(err)
	}
}
