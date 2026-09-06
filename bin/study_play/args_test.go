package main

import "testing"

func TestParsePlay(t *testing.T) {
	opts, ctl, perr := parsePlay([]string{"--", "--drill", "core", "--brief", "--run", "core"})
	if ctl != gnuOK || perr || opts.drillKind != "core" || opts.solutionKind != "" || !opts.brief || opts.runMode != "core" {
		t.Fatalf("parsePlay core: %+v ctl=%v perr=%v", opts, ctl, perr)
	}

	opts, ctl, perr = parsePlay([]string{"--drill", "reflex"})
	if ctl != gnuOK || perr || opts.drillKind != "reflex" || opts.runMode != "" {
		t.Fatalf("parsePlay reflex: %+v", opts)
	}

	// --option=value form.
	opts, _, _ = parsePlay([]string{"--drill=reflex", "--run=core"})
	if opts.drillKind != "reflex" || opts.runMode != "core" {
		t.Fatalf("parsePlay --opt=value: %+v", opts)
	}

	opts, _, _ = parsePlay([]string{"--solution", "core"})
	if opts.solutionKind != "core" || opts.runMode != "" {
		t.Fatalf("parsePlay solution core: %+v", opts)
	}

	// --run-core5 is a retired spelling of --run=core.
	opts, _, perr = parsePlay([]string{"--run-core5"})
	if perr || opts.runMode != "core" {
		t.Fatalf("parsePlay run-core5 alias: %+v perr=%v", opts, perr)
	}

	opts, _, _ = parsePlay([]string{"--run"})
	if opts.runMode != "all" {
		t.Fatalf("parsePlay bare run: %+v", opts)
	}

	// Session-view flags land on the struct.
	opts, _, _ = parsePlay([]string{"--refresh", "--show"})
	if !opts.refresh || !opts.show {
		t.Fatalf("parsePlay refresh/show: %+v", opts)
	}
	opts, _, _ = parsePlay([]string{"--weak"})
	if !opts.weak {
		t.Fatalf("parsePlay weak: %+v", opts)
	}

	// Unambiguous abbreviation.
	opts, _, _ = parsePlay([]string{"--prob"})
	if !opts.problems {
		t.Fatalf("parsePlay --prob -> --problems: %+v", opts)
	}

	_, _, perr = parsePlay([]string{"--drill"})
	if !perr {
		t.Fatal("bare --drill should be a usage error")
	}
	_, _, perr = parsePlay([]string{"--solution"})
	if !perr {
		t.Fatal("bare --solution should be a usage error")
	}
	_, _, perr = parsePlay([]string{"--drill", "bogus"})
	if !perr {
		t.Fatal("unknown drill kind should be a usage error")
	}

	if _, ctl, _ := parsePlay([]string{"--help"}); ctl != gnuHelp {
		t.Fatal("--help -> gnuHelp")
	}
	if _, ctl, _ := parsePlay([]string{"-V"}); ctl != gnuVersion {
		t.Fatal("-V -> gnuVersion")
	}
	if _, ctl, _ := parsePlay([]string{"--nonsense"}); ctl != gnuErr {
		t.Fatal("--nonsense -> gnuErr")
	}
}
