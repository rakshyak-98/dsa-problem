package main

import "testing"

func TestParseDailyArgs(t *testing.T) {
	opts := parseDailyArgs([]string{"--", "--drill", "--run"})
	if !opts.run || len(opts.passArgs) != 2 || opts.track != trackDSA {
		t.Fatalf("parseDailyArgs: %+v", opts)
	}

	opts = parseDailyArgs([]string{"--track", "backend", "--cram"})
	if opts.track != trackBackend || len(opts.passArgs) != 1 || opts.passArgs[0] != "--cram" {
		t.Fatalf("backend track: %+v", opts)
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
}

func TestIsKnownTrack(t *testing.T) {
	if !isKnownTrack(trackDSA) || !isKnownTrack(trackBackend) {
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
