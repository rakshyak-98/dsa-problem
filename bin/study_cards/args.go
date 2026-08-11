package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type options struct {
	help     bool
	review   bool
	due      bool
	stats    bool
	list     bool
	catalog  bool
	reset    bool
	deck     string
	tag      string
	limit    int
	newN     int
	shuffle  bool
	parseErr bool
}

func knownDecks() []string {
	root := cardsRoot(".")
	deckDir := filepath.Join(root, "decks")
	entries, err := os.ReadDir(deckDir)
	if err != nil {
		return nil
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		names = append(names, strings.TrimSuffix(e.Name(), ".json"))
	}
	sort.Strings(names)
	return names
}

func isKnownDeck(name string) bool {
	for _, d := range knownDecks() {
		if strings.EqualFold(d, name) {
			return true
		}
	}
	// Allow through if deck dir not yet populated (tests / fresh clone).
	return len(knownDecks()) == 0
}

func printOptionRequiresArg(flag string) {
	fmt.Fprintf(os.Stderr, "option %q requires an argument\n", flag)
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_cards -- --help' for more information.")
}

func printUnknownOption(arg string) {
	fmt.Fprintf(os.Stderr, "unknown option %q\n", arg)
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_cards -- --help' for more information.")
}

func printUnknownDeck(name string) {
	fmt.Fprintf(os.Stderr, "unknown deck %q\n", name)
	fmt.Fprintf(os.Stderr, "Valid decks: %s\n", strings.Join(knownDecks(), ", "))
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_cards -- --help' for more information.")
}

func printInvalidNumber(flag, value string) {
	fmt.Fprintf(os.Stderr, "invalid argument %q for option %q\n", value, flag)
	fmt.Fprintln(os.Stderr, "Try 'go run ./bin/study_cards -- --help' for more information.")
}

// splitLongOpt handles GNU --name=value forms. Returns name, value, ok.
func splitLongOpt(arg string) (name, value string, hasValue bool) {
	if !strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "--=") {
		return arg, "", false
	}
	body := strings.TrimPrefix(arg, "--")
	if i := strings.IndexByte(body, '='); i >= 0 {
		return "--" + body[:i], body[i+1:], true
	}
	return arg, "", false
}

func takeArg(args []string, i int, flag string) (value string, next int, ok bool) {
	if i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
		printOptionRequiresArg(flag)
		return "", i, false
	}
	return args[i+1], i + 1, true
}

func parsePositiveInt(flag, raw string) (int, bool) {
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		printInvalidNumber(flag, raw)
		return 0, false
	}
	return n, true
}

func parseNonNegInt(flag, raw string) (int, bool) {
	n, err := strconv.Atoi(raw)
	if err != nil || n < 0 {
		printInvalidNumber(flag, raw)
		return 0, false
	}
	return n, true
}

func parseArgs(args []string) options {
	opts := options{limit: 20, newN: 5, review: true, shuffle: true}
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}

	for i := 0; i < len(args); i++ {
		a := args[i]

		// GNU: lone "--" ends option parsing (nothing positional afterward).
		if a == "--" {
			if i+1 < len(args) {
				printUnknownOption(args[i+1])
				opts.parseErr = true
				return opts
			}
			break
		}

		name, inline, hasInline := splitLongOpt(a)
		if hasInline {
			a = name
		}

		switch a {
		case "-h", "--help":
			opts.help = true
			opts.review = false

		case "--review":
			opts.review = true
			opts.due = false
			opts.stats = false
			opts.list = false
			opts.catalog = false
			opts.reset = false

		case "--due":
			opts.due = true
			opts.review = false

		case "--stats":
			opts.stats = true
			opts.review = false

		case "--list":
			opts.list = true
			opts.review = false

		case "--catalog", "--decks":
			// --decks kept as alias; --catalog matches study_play / study_code
			opts.catalog = true
			opts.review = false

		case "--reset":
			opts.reset = true
			opts.review = false

		case "-d", "--deck":
			var val string
			var ok bool
			if hasInline {
				val = inline
				ok = true
				if val == "" {
					printOptionRequiresArg("--deck")
					opts.parseErr = true
					return opts
				}
			} else {
				val, i, ok = takeArg(args, i, a)
				if !ok {
					opts.parseErr = true
					return opts
				}
			}
			if !isKnownDeck(val) {
				printUnknownDeck(val)
				opts.parseErr = true
				return opts
			}
			opts.deck = strings.ToLower(val)

		case "-t", "--tag":
			var val string
			var ok bool
			if hasInline {
				val = inline
				ok = true
				if val == "" {
					printOptionRequiresArg("--tag")
					opts.parseErr = true
					return opts
				}
			} else {
				val, i, ok = takeArg(args, i, a)
				if !ok {
					opts.parseErr = true
					return opts
				}
			}
			opts.tag = val

		case "-n", "--limit":
			var raw string
			var ok bool
			if hasInline {
				raw = inline
				if raw == "" {
					printOptionRequiresArg("--limit")
					opts.parseErr = true
					return opts
				}
			} else {
				raw, i, ok = takeArg(args, i, a)
				if !ok {
					opts.parseErr = true
					return opts
				}
			}
			n, ok := parsePositiveInt("--limit", raw)
			if !ok {
				opts.parseErr = true
				return opts
			}
			opts.limit = n

		case "--new":
			var raw string
			var ok bool
			if hasInline {
				raw = inline
				if raw == "" {
					printOptionRequiresArg("--new")
					opts.parseErr = true
					return opts
				}
			} else {
				raw, i, ok = takeArg(args, i, a)
				if !ok {
					opts.parseErr = true
					return opts
				}
			}
			n, ok := parseNonNegInt("--new", raw)
			if !ok {
				opts.parseErr = true
				return opts
			}
			opts.newN = n

		case "--no-shuffle":
			opts.shuffle = false

		default:
			printUnknownOption(args[i])
			opts.parseErr = true
			return opts
		}
	}
	return opts
}

func printHelp() {
	fmt.Print(`Usage: go run ./bin/study_cards -- [OPTION]...
   or: go run . -- --track cards [OPTION]...

Review spaced-repetition flashcards extracted from doc/ notes.

Options:
  -h, --help               display this help message and exit
      --review             review due (+ new) cards (default)
      --due                show due / new / waiting counts and exit
      --stats              show progress by deck and exit
      --list               list card fronts (respects --deck/--tag/--limit)
      --catalog            list decks and card counts and exit
      --reset              clear local SRS progress (asks for confirmation)
  -d, --deck=NAME          only this deck (see --catalog for names)
  -t, --tag=TAG            only cards with this tag (e.g. back2basics, triage)
  -n, --limit=N            max cards this session (default: 20)
      --new=N              max new cards to introduce (default: 5)
      --no-shuffle         keep deck order instead of shuffling

DSA decks come from doc/backend/ (backend + star interview Q&A).

During review, rate each card:
  1  again   (forgot — due tomorrow)
  2  hard    (remembered with effort)
  3  good    (normal — +1 / +3 / +7 / +21 day ladder)
  4  easy    (too easy — longer interval)
  q          quit and save

Progress is saved to cards/.srs_progress.json (gitignored).
`)
}
