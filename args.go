package main

import "strings"

type drillTrack string

const (
	trackDSA       drillTrack = "dsa"
	trackBackend   drillTrack = "backend"
	trackRead      drillTrack = "read"
	trackWrite     drillTrack = "write"
	trackLeetcode  drillTrack = "leetcode"
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
	return s == "core" || s == "reflex" || s == "revision"
}

func validDrillKinds(track drillTrack) []string {
	switch track {
	case trackRead:
		return []string{"reflex"}
	case trackBackend:
		return []string{"core", "reflex", "revision"}
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
		switch {
		case a == "--help" || a == "-h":
			opts.help = true
		case a == "--list-tracks":
			opts.listTracks = true
		case a == "--track" || a == "-t":
			if i+1 >= len(args) {
				opts.help = true
				continue
			}
			i++
			opts.track = drillTrack(strings.ToLower(args[i]))
		case strings.HasPrefix(a, "--track="):
			opts.track = drillTrack(strings.ToLower(strings.TrimPrefix(a, "--track=")))
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
