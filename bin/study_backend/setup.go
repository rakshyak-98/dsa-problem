package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func runSetup() error {
	root := findRepoRoot(mustCwd())
	dst := filepath.Join(root, "drills", "backend")
	if _, err := os.Stat(filepath.Join(dst, "explain", "core5")); err != nil {
		return fmt.Errorf("drills/backend/ missing; expected explain/core5 under %s", dst)
	}
	fmt.Println("Backend drills ready at drills/backend/ (canonical; no seed copy).")
	return nil
}
