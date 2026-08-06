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

type dailyOptions struct {
	track      drillTrack
	passArgs   []string
	run        bool
	help       bool
	listTracks bool
}

func specialtyOnly(opts dailyOptions) bool {
	if opts.help || opts.listTracks {
		return false
	}
	for _, a := range opts.passArgs {
		switch a {
		case "--catalog", "--setup", "--cram", "--weak", "--reset", "--verbose":
			return false
		}
	}
	return true
}

func moduleArgs(opts dailyOptions) []string {
	if !specialtyOnly(opts) {
		return opts.passArgs
	}
	return append([]string{"--specialty"}, opts.passArgs...)
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
		case "--run":
			opts.run = true
			opts.passArgs = append(opts.passArgs, a)
		case "--verbose":
			opts.passArgs = append(opts.passArgs, a)
		default:
			opts.passArgs = append(opts.passArgs, a)
		}
	}
	return opts
}
