package main

import "strings"

type drillTrack string

const (
	trackDSA      drillTrack = "dsa"
	trackRead     drillTrack = "read"
	trackWrite    drillTrack = "write"
	trackLeetcode drillTrack = "leetcode"
)

type trackInfo struct {
	name        drillTrack
	title       string
	description string
}

var availableTracks = []trackInfo{
	{trackDSA, "dsa", "reflex writing drills (Core 5 + weekday specialty)"},
	{trackRead, "read", "reflex reading drills only"},
	{trackWrite, "write", "writing drills only"},
	{trackLeetcode, "leetcode", "daily 10-question LeetCode practice set (weekday topic)"},
}

func isKnownTrack(track drillTrack) bool {
	for _, t := range availableTracks {
		if t.name == track {
			return true
		}
	}
	return false
}

func isDrillKind(s string) bool {
	return s == "core" || s == "reflex"
}

func validDrillKinds(track drillTrack) []string {
	switch track {
	case trackRead:
		return []string{"reflex"}
	default:
		return []string{"core", "reflex"}
	}
}

func isDrillKindForTrack(track drillTrack, kind string) bool {
	for _, k := range validDrillKinds(track) {
		if k == kind {
			return true
		}
	}
	return false
}

func formatDrillKinds(track drillTrack) string {
	kinds := validDrillKinds(track)
	out := kinds[0]
	for _, k := range kinds[1:] {
		out += ", " + k
	}
	return out
}

func isRunSideFlag(s string) bool {
	return s == "-r" || s == "--read" || s == "-w" || s == "--write" || s == "-l" || s == "--leetcode"
}

func parseRunSide(s string) string {
	switch s {
	case "-r", "--read":
		return "read"
	case "-l", "--leetcode":
		return "leetcode"
	default:
		return "write"
	}
}

func isRunTarget(s string) bool {
	return s == "leetcode" || isDrillKind(s)
}

func hasRunKind(passArgs []string) bool {
	for i, a := range passArgs {
		if a == "--run" && i+1 < len(passArgs) && isDrillKind(passArgs[i+1]) {
			return true
		}
	}
	return false
}

// dailyProg is the program name shown in usage messages and --version.
const dailyProg = "dsa-drills"

// dailyLongOpts is every long option the root runner accepts. Options it does
// not act on itself (they are forwarded to the active track's module) are still
// listed so a typo is reported here instead of silently ignored downstream.
var dailyLongOpts = []string{
	"help", "version", "list-tracks", "track", "core5",
	"drill", "solution", "run", "catalog",
	"refresh", "show", "levels", "problems", "triggers", "weak", "brief",
	"read", "write", "leetcode", // deprecated run-side selectors; prefer --track
}

var dailyAliases = map[string][]string{}

type dailyOptions struct {
	track           drillTrack
	passArgs        []string
	run             bool
	runSide         string // "read", "write", or ""
	help            bool
	version         bool
	usageErr        bool
	listTracks      bool
	catalog         bool
	core5           bool
	drillKind       string
	drillMissing    bool
	drillUnknown    string
	solutionKind    string
	solutionMissing bool
	solutionUnknown string
}

func parseDailyArgs(args []string) dailyOptions {
	opts := dailyOptions{track: trackDSA}
	norm, ctl := gnuPre(args, gnuSpec{prog: dailyProg, canonical: dailyLongOpts, aliases: dailyAliases})
	switch ctl {
	case gnuHelp:
		opts.help = true
		return opts
	case gnuVersion:
		opts.version = true
		return opts
	case gnuErr:
		opts.usageErr = true
		return opts
	}
	args = norm
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "--list-tracks":
			opts.listTracks = true
		case a == "--catalog":
			opts.catalog = true
			opts.passArgs = append(opts.passArgs, a)
		case a == "--track" || a == "-t":
			if i+1 >= len(args) {
				printNeedsArg("--track")
				opts.usageErr = true
				return opts
			}
			i++
			opts.track = drillTrack(strings.ToLower(args[i]))
		case a == "--core5":
			opts.core5 = true
		case a == "--drill":
			if i+1 >= len(args) {
				opts.drillMissing = true
				continue
			}
			kind := args[i+1]
			if !isDrillKindForTrack(opts.track, kind) {
				opts.drillUnknown = kind
				i++
				continue
			}
			i++
			opts.drillKind = kind
			opts.passArgs = append(opts.passArgs, "--drill", kind)
		case a == "--solution":
			if i+1 >= len(args) {
				opts.solutionMissing = true
				continue
			}
			kind := args[i+1]
			if !isDrillKindForTrack(opts.track, kind) {
				opts.solutionUnknown = kind
				i++
				continue
			}
			i++
			opts.solutionKind = kind
			opts.passArgs = append(opts.passArgs, "--solution", kind)
		case a == "--run":
			opts.run = true
			opts.passArgs = append(opts.passArgs, a)
			if i+1 < len(args) && args[i+1] == "leetcode" {
				i++
				if opts.runSide != "" && opts.runSide != "leetcode" {
					opts.runSide = "conflict"
				} else {
					opts.runSide = "leetcode"
				}
			} else if i+1 < len(args) && isDrillKind(args[i+1]) {
				i++
				opts.passArgs = append(opts.passArgs, args[i])
			}
			for i+1 < len(args) && isRunSideFlag(args[i+1]) {
				i++
				side := parseRunSide(args[i])
				if opts.runSide != "" && opts.runSide != side {
					opts.runSide = "conflict"
				} else {
					opts.runSide = side
				}
				opts.passArgs = append(opts.passArgs, args[i])
			}
		default:
			opts.passArgs = append(opts.passArgs, a)
		}
	}
	return opts
}
