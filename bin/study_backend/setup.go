package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func copyTree(srcRoot, dstRoot string) error {
	return filepath.Walk(srcRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		return copyFile(path, filepath.Join(dstRoot, rel))
	})
}

func runSetup() error {
	root := findRepoRoot(mustCwd())
	srcPractice := filepath.Join(root, "bin", "study_backend", "practice")
	dstDrills := filepath.Join(root, "drills", "backend")

	tasks := []struct{ src, dst string }{
		{"explain/core5", "explain/core5"},
		{"explain/blocks", "explain/blocks"},
		{"write/core5", "write/core5"},
		{"scenario/mock_scenarios", "scenario/mock_scenarios"},
	}
	for _, t := range tasks {
		if err := copyTree(
			filepath.Join(srcPractice, t.src),
			filepath.Join(dstDrills, t.dst),
		); err != nil {
			return fmt.Errorf("copy %s: %w", t.src, err)
		}
	}
	fmt.Println("Copied practice drills → drills/backend/")
	return nil
}
