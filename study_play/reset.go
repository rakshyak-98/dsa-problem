package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed blanks/*
var blankFS embed.FS

func resetTodayDrill(today drill, drillDir string) error {
	blankName := today.file + ".go"
	content, err := blankFS.ReadFile("blanks/" + blankName)
	if err != nil {
		return fmt.Errorf("no blank template for %s: %w", today.file, err)
	}

	// blanks/ use //go:build ignore so they are not compiled; strip before writing
	content = bytes.TrimPrefix(content, []byte("//go:build ignore\n\n"))
	content = bytes.TrimPrefix(content, []byte("//go:build ignore\n"))

	target := filepath.Join(drillDir, "main.go")
	if err := os.WriteFile(target, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", target, err)
	}

	fmt.Printf("Reset today's drill → %s\n", target)
	fmt.Println("All TODO: REFLEX functions restored to panic(\"Implement from memory\").")
	fmt.Printf("Open: study_play/drills/%s\n", today.file)
	return nil
}
