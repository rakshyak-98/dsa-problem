package main

import "testing"

func TestParseDailyArgs(t *testing.T) {
	args, run := parseDailyArgs([]string{"--", "--micro", "--run"})
	if !run || len(args) != 2 {
		t.Fatalf("parseDailyArgs: args=%v run=%v", args, run)
	}
	args, run = parseDailyArgs([]string{"--catalog"})
	if run || len(args) != 0 {
		t.Fatal("unexpected flags")
	}
}

func TestRunInInvalidDir(t *testing.T) {
	if err := runIn(t.TempDir(), "--catalog"); err == nil {
		t.Fatal("expected error for invalid go module dir")
	}
}
