package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Print(`Usage: go run ./bin/study_leetcode -- [OPTION]...

Daily 10-question LeetCode practice set aligned with today's reflex topic.
Fetches live problem data from the LeetCode GraphQL API and saves to
drills/leetcode/daily.json and drills/leetcode/daily.md (full statements).

Options:
  -h, --help               display this help message and exit
      --catalog            list all weekday practice sets
      --set                show today's 10 problems (default)
      --run                fetch + show today's 10 problems (unified runner)
      --refresh            re-fetch from LeetCode even if daily.json is current
      --brief              one-line output for unified daily runner

`)
}

func parseLeetcodeArgs(args []string) (help, catalog, brief, showSet, refresh bool, parseErr bool) {
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
		case "--run", "-l", "--leetcode":
			showSet = true
			brief = false
		case "--refresh":
			refresh = true
		case "--brief":
			brief = true
		case "-r", "--read", "-w", "--write":
			// consumed by root CLI when selecting run side
		default:
			fmt.Fprintf(os.Stderr, "unknown option %q\n", args[i])
			fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_leetcode -- --help' for more information.")
			return help, catalog, brief, showSet, refresh, true
		}
	}
	return help, catalog, brief, showSet, refresh, false
}
