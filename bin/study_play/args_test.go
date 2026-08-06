package main

import "testing"

func TestParsePlayArgs(t *testing.T) {
	drillKind, brief, runMath, runMode, drillErr := parsePlayArgs([]string{"--", "--drill", "core", "--brief", "--run", "core"})
	if drillErr || drillKind != "core" || !brief || runMath || runMode != "core" {
		t.Fatal("parsePlayArgs core")
	}
	drillKind, _, runMath, runMode, drillErr = parsePlayArgs([]string{"--drill", "reflex"})
	if drillErr || drillKind != "reflex" || runMath || runMode != "" {
		t.Fatal("parsePlayArgs reflex")
	}
	_, _, runMath, runMode, drillErr = parsePlayArgs([]string{"--run-core5"})
	if drillErr || runMode != "core" {
		t.Fatal("run-core5 alias")
	}
	_, _, runMath, runMode, drillErr = parsePlayArgs([]string{"--run"})
	if drillErr || runMode != "all" {
		t.Fatal("bare run")
	}
	_, _, _, _, drillErr = parsePlayArgs([]string{"--drill"})
	if !drillErr {
		t.Fatal("bare drill should error")
	}
	_, _, _, _, drillErr = parsePlayArgs([]string{"--drill", "nope"})
	if !drillErr {
		t.Fatal("unknown drill kind should error")
	}
}
