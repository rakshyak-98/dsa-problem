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

type dailyOptions struct {
	track        drillTrack
	passArgs     []string
	run          bool
	help         bool
	listTracks   bool
	core5        bool
	drillKind    string
	drillMissing bool
	drillUnknown string
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
		case "--run":
			opts.run = true
			opts.passArgs = append(opts.passArgs, a)
			if i+1 < len(args) && isDrillKind(args[i+1]) {
				i++
				opts.passArgs = append(opts.passArgs, args[i])
			}
		default:
			opts.passArgs = append(opts.passArgs, a)
		}
	}
	return opts
}
