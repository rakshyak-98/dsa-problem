package main

import "testing"

func TestParsePlayArgs(t *testing.T) {
	drillKind, solutionKind, help, brief, runMath, runMode, parseErr := parsePlayArgs([]string{"--", "--drill", "core", "--brief", "--run", "core"})
	if parseErr || drillKind != "core" || solutionKind != "" || !brief || runMath || runMode != "core" || help {
		t.Fatal("parsePlayArgs core")
	}
	drillKind, solutionKind, help, _, _, runMode, parseErr = parsePlayArgs([]string{"--drill", "reflex"})
	if parseErr || drillKind != "reflex" || solutionKind != "" || help || runMode != "" {
		t.Fatal("parsePlayArgs reflex")
	}
	_, solutionKind, help, _, _, runMode, parseErr = parsePlayArgs([]string{"--solution", "core"})
	if parseErr || solutionKind != "core" || help || runMode != "" {
		t.Fatal("parsePlayArgs solution core")
	}
	_, _, help, _, runMath, runMode, parseErr = parsePlayArgs([]string{"--run-core5"})
	if parseErr || runMath || runMode != "core" || help {
		t.Fatal("parsePlayArgs run-core5")
	}
	_, _, help, _, runMath, runMode, parseErr = parsePlayArgs([]string{"--run"})
	if parseErr || runMath || runMode != "all" || help {
		t.Fatal("parsePlayArgs bare run")
	}
	_, _, help, _, _, _, parseErr = parsePlayArgs([]string{"--drill"})
	if !parseErr || help {
		t.Fatal("bare drill should error")
	}
	_, _, help, _, _, _, parseErr = parsePlayArgs([]string{"--solution"})
	if !parseErr || help {
		t.Fatal("bare solution should error")
	}
	_, _, help, _, _, _, parseErr = parsePlayArgs([]string{"--help"})
	if parseErr || !help {
		t.Fatal("help flag")
	}
}
