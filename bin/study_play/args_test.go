package main

import "testing"

func TestParsePlayArgs(t *testing.T) {
	drillKind, solutionKind, brief, runMath, runMode, parseErr := parsePlayArgs([]string{"--", "--drill", "core", "--brief", "--run", "core"})
	if parseErr || drillKind != "core" || solutionKind != "" || !brief || runMath || runMode != "core" {
		t.Fatal("parsePlayArgs core")
	}
	drillKind, solutionKind, _, _, runMode, parseErr = parsePlayArgs([]string{"--drill", "reflex"})
	if parseErr || drillKind != "reflex" || solutionKind != "" || runMode != "" {
		t.Fatal("parsePlayArgs reflex")
	}
	_, solutionKind, _, _, runMode, parseErr = parsePlayArgs([]string{"--solution", "core"})
	if parseErr || solutionKind != "core" || runMode != "" {
		t.Fatal("parsePlayArgs solution core")
	}
	_, _, _, runMath, runMode, parseErr = parsePlayArgs([]string{"--run-core5"})
	if parseErr || runMode != "core" {
		t.Fatal("run-core5 alias")
	}
	_, _, _, runMath, runMode, parseErr = parsePlayArgs([]string{"--run"})
	if parseErr || runMode != "all" {
		t.Fatal("bare run")
	}
	_, _, _, _, _, parseErr = parsePlayArgs([]string{"--drill"})
	if !parseErr {
		t.Fatal("bare drill should error")
	}
	_, _, _, _, _, parseErr = parsePlayArgs([]string{"--solution"})
	if !parseErr {
		t.Fatal("bare solution should error")
	}
}
