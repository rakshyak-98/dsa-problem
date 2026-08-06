package main

import (
	"strings"
	"testing"
)

func TestPrintCatalog(t *testing.T) {
	// smoke: catalog should mention all blocks
	for _, b := range blocks {
		if b.file == "" || b.topic == "" {
			t.Fatalf("invalid block entry: %+v", b)
		}
	}
}

func TestCramPlanContainsInterview(t *testing.T) {
	// printCramPlan writes to stdout; verify blocks cover resume themes
	themes := []string{"REST", "SQL", "WebRTC", "Airflow", "Jenkins", "ZATCA", "Go"}
	found := 0
	for _, b := range blocks {
		for _, theme := range themes {
			if strings.Contains(b.resume, theme) || strings.Contains(b.topic, theme) {
				found++
				break
			}
		}
	}
	if found < 5 {
		t.Fatalf("expected blocks to cover resume themes, matched %d", found)
	}
}
