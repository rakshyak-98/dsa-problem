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

	opts = parseDailyArgs([]string{"--", "--run", "reflex"})
	if !opts.run || len(opts.passArgs) != 2 || opts.passArgs[1] != "reflex" {
		t.Fatalf("parseDailyArgs reflex: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--track", "backend", "--cram"})
	if opts.track != trackBackend || len(opts.passArgs) != 1 || opts.passArgs[0] != "--cram" {
		t.Fatalf("backend track: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--track=cards", "--due", "--deck=jargon"})
	if opts.track != trackCards || len(opts.passArgs) != 2 || opts.passArgs[0] != "--due" || opts.passArgs[1] != "--deck=jargon" {
		t.Fatalf("cards track: %+v", opts)
	}

	opts = parseDailyArgs([]string{"-t", "write", "--run"})
	if opts.track != trackWrite || !opts.run {
		t.Fatalf("write track: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--catalog"})
	if opts.run || len(opts.passArgs) != 1 || opts.passArgs[0] != "--catalog" {
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
}

func TestIsKnownTrack(t *testing.T) {
	if !isKnownTrack(trackDSA) || !isKnownTrack(trackBackend) || !isKnownTrack(trackCards) {
		t.Fatal("known tracks")
	}
	if isKnownTrack("nope") {
		t.Fatal("unknown track should fail")
	}
}

func TestRunInInvalidDir(t *testing.T) {
	if err := runIn(t.TempDir(), "--catalog"); err == nil {
		t.Fatal("expected error for invalid go module dir")
	}
}
