package main

import "testing"

func TestParsePlayArgs(t *testing.T) {
	drill, brief, runMath, runMode := parsePlayArgs([]string{"--", "--drill", "--brief", "--run", "core"})
	if !drill || !brief || runMath || runMode != "core" {
		t.Fatal("parsePlayArgs core")
	}
	_, _, runMath, runMode = parsePlayArgs([]string{"--run-core5"})
	if runMode != "core" {
		t.Fatal("run-core5 alias")
	}
	_, _, runMath, runMode = parsePlayArgs([]string{"--run"})
	if runMode != "all" {
		t.Fatal("bare run")
	}
}
