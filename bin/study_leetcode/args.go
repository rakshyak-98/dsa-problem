package main

import (
	"fmt"
	"os"
)

func isRunKind(s string) bool {
	return false // LeetCode sets are solved externally; no local --run
}

func printHelp() {
	fmt.Print(`Usage: go run ./bin/study_leetcode -- [OPTION]...

Daily 10-question LeetCode practice set aligned with today's reflex topic.
Solve on leetcode.com — separate from in-repo reflex drills.

Options:
  -h, --help               display this help message and exit
      --catalog            list all weekday practice sets
      --set                show today's 10 problems (default)
      --brief              one-line output for unified daily runner

`)
}

func parseLeetcodeArgs(args []string) (help, catalog, brief, showSet bool, parseErr bool) {
	showSet = true
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			help = true
		case "--catalog":
			catalog = true
			showSet = false
		case "--set":
			showSet = true
		case "--brief":
			brief = true
		case "--run":
			fmt.Fprintln(os.Stderr, "leetcode practice sets are solved on leetcode.com; no local --run")
			fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_leetcode -- --help' for more information.")
			return help, catalog, brief, showSet, true
		default:
			fmt.Fprintf(os.Stderr, "unknown option %q\n", args[i])
			fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_leetcode -- --help' for more information.")
			return help, catalog, brief, showSet, true
		}
	}
	return help, catalog, brief, showSet, false
}
