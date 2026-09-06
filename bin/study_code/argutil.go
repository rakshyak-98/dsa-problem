package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// cliVersion is reported by --version. Bump it when the option surface changes.
const cliVersion = "0.2.0"

// GNU-style option handling shared by every entry point.
//
// The hand-written switch parsers below only understand exact long options in
// the form "--name value". gnuPre adds the conventions the GNU standards ask
// for and the old parsers lacked:
//
//   - a single leading "--" (the `go run .` separator) is dropped
//   - "--name=value" is accepted and split into two tokens
//   - unambiguous abbreviations are expanded ("--ref" -> "--refresh")
//   - "--help"/"-h" and "--version"/"-V" are handled here
//   - deprecated spellings are rewritten to the current option
//   - an unknown or ambiguous --option is a usage error (exit status 2)
//
// Short options and bare operands pass straight through for the caller's own
// switch to interpret.

type gnuCtl int

const (
	gnuOK      gnuCtl = iota // parsing may continue with the returned args
	gnuHelp                  // caller should print help and exit 0
	gnuVersion               // caller should print the version and exit 0
	gnuErr                   // message already written to stderr; exit 2
)

// gnuSpec is one program's long-option vocabulary.
type gnuSpec struct {
	prog      string              // program name for messages and --version
	canonical []string            // every accepted --long, without the dashes
	aliases   map[string][]string // deprecated --long -> replacement tokens
}

func gnuTryHelp(prog string) string {
	return fmt.Sprintf("Try '%s --help' for more information.", prog)
}

func printVersion(prog string) {
	fmt.Printf("%s %s\n", prog, cliVersion)
}

// resolveLong maps a possibly-abbreviated long name to its canonical spelling.
// ok is false when nothing matches or the abbreviation is ambiguous, in which
// case matches lists every option the prefix hit.
func resolveLong(name string, canonical []string) (canon string, ok bool, matches []string) {
	for _, c := range canonical {
		if c == name {
			return c, true, nil
		}
	}
	var hits []string
	for _, c := range canonical {
		if strings.HasPrefix(c, name) {
			hits = append(hits, c)
		}
	}
	if len(hits) == 1 {
		return hits[0], true, nil
	}
	sort.Strings(hits)
	return name, false, hits
}

func gnuPre(args []string, spec gnuSpec) ([]string, gnuCtl) {
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	out := make([]string, 0, len(args)+2)
	for i := 0; i < len(args); i++ {
		a := args[i]

		switch {
		case a == "--":
			out = append(out, args[i+1:]...) // explicit end of options
			return out, gnuOK
		case a == "-h", a == "--help":
			return nil, gnuHelp
		case a == "-V", a == "--version":
			return nil, gnuVersion
		case !strings.HasPrefix(a, "--"):
			out = append(out, a) // short option or operand
			continue
		}

		name, val, hasVal := a, "", false
		if eq := strings.IndexByte(a, '='); eq >= 0 {
			name, val, hasVal = a[:eq], a[eq+1:], true
		}
		bare := strings.TrimPrefix(name, "--")

		if repl, ok := spec.aliases[bare]; ok {
			out = append(out, repl...)
			if hasVal {
				out = append(out, val)
			}
			continue
		}

		canon, ok, matches := resolveLong(bare, spec.canonical)
		if !ok {
			if len(matches) > 1 {
				fmt.Fprintf(os.Stderr, "%s: option '--%s' is ambiguous; possibilities:", spec.prog, bare)
				for _, m := range matches {
					fmt.Fprintf(os.Stderr, " '--%s'", m)
				}
				fmt.Fprintln(os.Stderr)
			} else {
				fmt.Fprintf(os.Stderr, "%s: unrecognized option '--%s'\n", spec.prog, bare)
			}
			fmt.Fprintln(os.Stderr, gnuTryHelp(spec.prog))
			return nil, gnuErr
		}
		switch canon {
		case "help":
			return nil, gnuHelp
		case "version":
			return nil, gnuVersion
		}

		out = append(out, "--"+canon)
		if hasVal {
			out = append(out, val)
		}
	}
	return out, gnuOK
}

// hasArg reports whether flag appears literally in args.
func hasArg(args []string, flag string) bool {
	for _, a := range args {
		if a == flag {
			return true
		}
	}
	return false
}
