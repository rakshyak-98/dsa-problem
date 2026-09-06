package main

import "testing"

func TestParseDailyArgs(t *testing.T) {
	opts := parseDailyArgs([]string{"--", "--drill", "core", "--run", "core"})
	if !opts.run || opts.drillKind != "core" || len(opts.passArgs) != 4 || opts.passArgs[1] != "core" || opts.track != trackDSA {
		t.Fatalf("parseDailyArgs: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--", "--drill", "reflex"})
	if opts.drillKind != "reflex" || len(opts.passArgs) != 2 || opts.passArgs[1] != "reflex" {
		t.Fatalf("parseDailyArgs reflex: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--drill"})
	if !opts.drillMissing || opts.drillKind != "" {
		t.Fatalf("drill missing: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--drill", "nope"})
	if opts.drillUnknown != "nope" || opts.drillKind != "" {
		t.Fatalf("drill unknown: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--", "--run", "reflex", "-r"})
	if !opts.run || opts.runSide != "read" || len(opts.passArgs) != 3 {
		t.Fatalf("parseDailyArgs run reflex -r: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--run", "reflex", "--write"})
	if !opts.run || opts.runSide != "write" {
		t.Fatalf("parseDailyArgs run reflex --write: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--run", "reflex"})
	if !opts.run || opts.runSide != "" {
		t.Fatalf("parseDailyArgs run reflex no side: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--run", "reflex", "-r", "-w"})
	if opts.runSide != "conflict" {
		t.Fatalf("expected run side conflict: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--", "--run", "leetcode"})
	if !opts.run || opts.runSide != "leetcode" || len(opts.passArgs) != 1 || opts.passArgs[0] != "--run" {
		t.Fatalf("parseDailyArgs run leetcode: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--run", "-l"})
	if !opts.run || opts.runSide != "leetcode" {
		t.Fatalf("parseDailyArgs run -l: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--run", "reflex", "-l"})
	if !opts.run || opts.runSide != "leetcode" {
		t.Fatalf("parseDailyArgs run reflex -l prefers leetcode: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--", "--run", "reflex"})
	if !opts.run || len(opts.passArgs) != 2 || opts.passArgs[1] != "reflex" {
		t.Fatalf("parseDailyArgs reflex: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--track", "leetcode", "--catalog"})
	if opts.track != trackLeetcode || !opts.catalog || len(opts.passArgs) != 1 || opts.passArgs[0] != "--catalog" {
		t.Fatalf("leetcode track: %+v", opts)
	}

	opts = parseDailyArgs([]string{"-t", "write", "--run"})
	if opts.track != trackWrite || !opts.run {
		t.Fatalf("write track: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--catalog"})
	if opts.run || !opts.catalog || opts.track != trackDSA || len(opts.passArgs) != 1 || opts.passArgs[0] != "--catalog" {
		t.Fatalf("catalog passthrough: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--list-tracks"})
	if !opts.listTracks {
		t.Fatal("expected listTracks")
	}

	opts = parseDailyArgs([]string{"--help"})
	if !opts.help {
		t.Fatal("expected help")
	}

	opts = parseDailyArgs([]string{"--", "--solution", "reflex"})
	if opts.solutionKind != "reflex" || len(opts.passArgs) != 2 || opts.passArgs[1] != "reflex" {
		t.Fatalf("parseDailyArgs solution reflex: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--solution"})
	if !opts.solutionMissing || opts.solutionKind != "" {
		t.Fatalf("solution missing: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--core5"})
	if !opts.core5 || opts.track != trackDSA {
		t.Fatalf("core5: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--track", "read", "--drill", "core"})
	if opts.drillUnknown != "core" || opts.drillKind != "" {
		t.Fatalf("read track should reject core drill: %+v", opts)
	}
}

func TestParseDailyArgsGNU(t *testing.T) {
	// --version is handled by the shared front-end.
	if opts := parseDailyArgs([]string{"--version"}); !opts.version || opts.help {
		t.Fatalf("expected version: %+v", opts)
	}
	if opts := parseDailyArgs([]string{"-V"}); !opts.version {
		t.Fatalf("expected version from -V: %+v", opts)
	}

	// --option=value is split for the switch that follows.
	opts := parseDailyArgs([]string{"--drill=core", "--run=reflex"})
	if opts.drillKind != "core" || !opts.run {
		t.Fatalf("--opt=value: %+v", opts)
	}
	if opts := parseDailyArgs([]string{"--track=leetcode"}); opts.track != trackLeetcode {
		t.Fatalf("--track=value: %+v", opts)
	}

	// Unambiguous abbreviation resolves to the full option.
	if opts := parseDailyArgs([]string{"--ref"}); !containsStrDaily(opts.passArgs, "--refresh") {
		t.Fatalf("--ref should expand to --refresh: %+v", opts)
	}

	// An unknown or ambiguous option is a usage error, not a silent passthrough.
	if opts := parseDailyArgs([]string{"--bogus"}); !opts.usageErr {
		t.Fatalf("--bogus should be a usage error: %+v", opts)
	}
	if opts := parseDailyArgs([]string{"--r"}); !opts.usageErr {
		t.Fatalf("ambiguous --r should be a usage error: %+v", opts)
	}
}

func containsStrDaily(s []string, want string) bool {
	for _, v := range s {
		if v == want {
			return true
		}
	}
	return false
}

func TestIsKnownTrack(t *testing.T) {
	if !isKnownTrack(trackDSA) || !isKnownTrack(trackRead) || !isKnownTrack(trackWrite) || !isKnownTrack(trackLeetcode) {
		t.Fatal("known tracks")
	}
	if isKnownTrack("nope") || isKnownTrack("backend") {
		t.Fatal("unknown track should fail")
	}
}

func TestRunInInvalidDir(t *testing.T) {
	if err := runIn(t.TempDir(), "--catalog"); err == nil {
		t.Fatal("expected error for invalid go module dir")
	}
}
