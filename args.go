package main

import "strings"

type drillTrack string

const (
	trackDSA     drillTrack = "dsa"
	trackBackend drillTrack = "backend"
	trackRead    drillTrack = "read"
	trackWrite   drillTrack = "write"
)

type trackInfo struct {
	name        drillTrack
	title       string
	description string
}

var availableTracks = []trackInfo{
	{trackDSA, "dsa", "reading and reflex writing"},
	{trackRead, "read", "reading drills only"},
	{trackWrite, "write", "writing drills only"},
	{trackBackend, "backend", "interview prep"},
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

func isRunSideFlag(s string) bool {
	return s == "-r" || s == "--read" || s == "-w" || s == "--write"
}

func parseRunSide(s string) string {
	if s == "-r" || s == "--read" {
		return "read"
	}
	return "write"
}

func hasRunKind(passArgs []string) bool {
	for i, a := range passArgs {
		if a == "--run" && i+1 < len(passArgs) && isDrillKind(passArgs[i+1]) {
			return true
		}
	}
	return false
}

type dailyOptions struct {
	track           drillTrack
	passArgs        []string
	run             bool
	runSide         string // "read", "write", or ""
	help            bool
	listTracks      bool
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
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch a {
		case "--help", "-h":
			opts.help = true
		case "--list-tracks":
			opts.listTracks = true
		case "--track", "-t":
			if i+1 >= len(args) {
				opts.help = true
				continue
			}
			i++
			opts.track = drillTrack(strings.ToLower(args[i]))
		case "--core5":
			opts.core5 = true
		case "--drill":
			if i+1 >= len(args) {
				opts.drillMissing = true
				continue
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				opts.drillUnknown = kind
				i++
				continue
			}
			i++
			opts.drillKind = kind
			opts.passArgs = append(opts.passArgs, "--drill", kind)
		case "--solution":
			if i+1 >= len(args) {
				opts.solutionMissing = true
				continue
			}
			kind := args[i+1]
			if !isDrillKind(kind) {
				opts.solutionUnknown = kind
				i++
				continue
			}
			i++
			opts.solutionKind = kind
			opts.passArgs = append(opts.passArgs, "--solution", kind)
		case "--run":
			opts.run = true
			opts.passArgs = append(opts.passArgs, a)
			if i+1 < len(args) && isDrillKind(args[i+1]) {
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
