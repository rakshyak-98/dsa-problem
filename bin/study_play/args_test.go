package main

import "testing"

func TestParsePlay(t *testing.T) {
	opts, ctl, perr := parsePlay([]string{"--", "--drill", "reflex", "--brief", "--run", "reflex"})
	if ctl != gnuOK || perr || opts.drillKind != "reflex" || opts.solutionKind != "" || !opts.brief || opts.runMode != "reflex" {
		t.Fatalf("parsePlay reflex run: %+v ctl=%v perr=%v", opts, ctl, perr)
	}

	opts, ctl, perr = parsePlay([]string{"--drill", "reflex"})
	if ctl != gnuOK || perr || opts.drillKind != "reflex" || opts.runMode != "" {
		t.Fatalf("parsePlay reflex: %+v", opts)
	}

	// --option=value form.
	opts, _, _ = parsePlay([]string{"--drill=reflex", "--run=reflex"})
	if opts.drillKind != "reflex" || opts.runMode != "reflex" {
		t.Fatalf("parsePlay --opt=value: %+v", opts)
	}

	opts, _, _ = parsePlay([]string{"--solution", "reflex"})
	if opts.solutionKind != "reflex" || opts.runMode != "" {
		t.Fatalf("parsePlay solution reflex: %+v", opts)
	}

	opts, _, _ = parsePlay([]string{"--run"})
	if opts.runMode != "all" {
		t.Fatalf("parsePlay bare run: %+v", opts)
	}

	// Session-view flags land on the struct.
	opts, _, _ = parsePlay([]string{"--show"})
	if !opts.show {
		t.Fatalf("parsePlay show: %+v", opts)
	}
	if _, ctl, _ := parsePlay([]string{"--levels"}); ctl != gnuErr {
		t.Fatal("--levels was removed and should be a usage error")
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
