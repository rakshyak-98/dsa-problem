package main

import (
	"fmt"
	"os"
	"strings"
)

const leetcodeProg = "study_leetcode"

var leetcodeLongOpts = []string{
	"help", "version",
	"catalog", "show", "run", "refresh", "brief",
	"answers", "log", "gaps",
	"write", "leetcode", // consumed by the root runner when picking a side
}

var leetcodeAliases = map[string][]string{
	"set": {"--show"}, // superseded by: --show
}

func leetcodeSpec() gnuSpec {
	return gnuSpec{prog: leetcodeProg, canonical: leetcodeLongOpts, aliases: leetcodeAliases}
}

func printHelp() {
	fmt.Print(`Usage: go run . -- [OPTION]...

Daily 10-question LeetCode practice set aligned with today's reflex topic.
Statements are fetched from the LeetCode GraphQL API and written to
drills/leetcode/daily.json (plus daily.go and daily.md). With no option it
shows today's set.

Long options may be abbreviated while unambiguous, and every value-taking
option also accepts the --option=value form.

  -h, --help       display this help and exit
  -V, --version    display version information and exit
      --show       show today's 10 problems (default)
      --run        fetch if needed, then show today's 10 problems
      --refresh    re-fetch from the LeetCode API even if the cache is current
      --catalog    list every weekday practice set
      --brief      one-line output for the unified daily runner
      --answers    also show the topic, tags and reflex links (hidden by
                   default: name each pattern before you look)

Solved-problem cards (drills/leetcode/cards.json):
      --log=SLUG   record a solved problem, or review one that is due: your
                   pattern guess, the signal, the one-line insight. Schedules
                   the next re-solve (1, 3, 7, 14, 30, 60 days)
      --gaps       patterns on your cards that no reflex drill covers yet

Deprecated (still accepted): --set is now --show.

Exit status: 0 success, 1 a fetch failed, 2 a command-line usage error.
`)
}

// leetcodeOpts is the fully parsed command line for the LeetCode set CLI.
type leetcodeOpts struct {
	catalog bool
	brief   bool
	showSet bool
	refresh bool
	answers bool
	gaps    bool
	log     string // --log slug
}

func parseLeetcode(args []string) (opts leetcodeOpts, ctl gnuCtl, parseErr bool) {
	opts.showSet = true
	norm, ctl := gnuPre(args, leetcodeSpec())
	if ctl != gnuOK {
		return opts, ctl, false
	}
	for i := 0; i < len(norm); i++ {
		switch a := norm[i]; a {
		case "--answers":
			opts.answers = true
		case "--gaps":
			opts.gaps = true
		case "--log":
			if i+1 >= len(norm) || strings.HasPrefix(norm[i+1], "-") {
				fmt.Fprintf(os.Stderr, "%s: option '--log' requires a problem slug\n", leetcodeProg)
				fmt.Fprintln(os.Stderr, gnuTryHelp(leetcodeProg))
				return opts, gnuOK, true
			}
			i++
			opts.log = norm[i]
		case "--catalog":
			opts.catalog = true
			opts.showSet = false
		case "--show":
			opts.showSet = true
		case "--run", "-l", "--leetcode":
			opts.showSet = true
			opts.brief = false
		case "--refresh":
			opts.refresh = true
		case "--brief":
			opts.brief = true
		case "--write", "-w":
			// selected by the root runner; nothing to do here
		}
	}
	return opts, gnuOK, false
}
