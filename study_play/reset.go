package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed _support/blanks/*
var blankFS embed.FS

func blankContent(file string) ([]byte, error) {
	content, err := blankFS.ReadFile("_support/blanks/" + file + ".go")
	if err != nil {
		return nil, fmt.Errorf("no blank template for %s: %w", file, err)
	}

	// blanks/ use //go:build ignore so they are not compiled; strip before writing
	content = bytes.TrimPrefix(content, []byte("//go:build ignore\n\n"))
	content = bytes.TrimPrefix(content, []byte("//go:build ignore\n"))
	return content, nil
}

func writeDrillFromBlank(file, drillDir string) error {
	content, err := blankContent(file)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(drillDir, 0o755); err != nil {
		return fmt.Errorf("create %s: %w", drillDir, err)
	}

	target := filepath.Join(drillDir, "main.go")
	if err := os.WriteFile(target, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", target, err)
	}
	return nil
}

func setupAllDrills(repoRoot string) error {
	for _, d := range drills {
		drillDir := writeReflexDir(repoRoot, d.file)
		if err := writeDrillFromBlank(d.file, drillDir); err != nil {
			return err
		}
		fmt.Printf("  ✓ drills/write/reflex/%s/main.go\n", d.file)
	}
	for _, file := range bonusDrills {
		drillDir := writeReflexDir(repoRoot, file)
		if err := writeDrillFromBlank(file, drillDir); err != nil {
			return err
		}
		fmt.Printf("  ✓ drills/write/reflex/%s/main.go (bonus)\n", file)
	}
	return nil
}

func resetTodayDrill(today drill, drillDir string) error {
	if err := writeDrillFromBlank(today.file, drillDir); err != nil {
		return err
	}

	fmt.Printf("Reset today's drill → %s\n", filepath.Join(drillDir, "main.go"))
	fmt.Println("All TODO: REFLEX functions restored to panic(\"Implement from memory\").")
	fmt.Printf("Open: drills/write/reflex/%s\n", today.file)
	return nil
}
