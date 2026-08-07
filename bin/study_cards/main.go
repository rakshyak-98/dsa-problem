// Spaced repetition for interview-prep notes (cards/decks).
//
// RUN:       go run ./bin/study_cards
// DUE:       go run ./bin/study_cards -- --due
// STATS:     go run ./bin/study_cards -- --stats
// CATALOG:   go run ./bin/study_cards -- --catalog
// DECK:      go run ./bin/study_cards -- --deck=jargon --limit=15
package main

import (
	"fmt"
	"os"
)

func main() {
	opts := parseArgs(os.Args[1:])
	if opts.parseErr {
		os.Exit(2)
	}
	if opts.help {
		printHelp()
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	root := cardsRoot(wd)

	if opts.reset {
		if err := runReset(root); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	cards, err := loadCards(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	s, err := loadStore(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch {
	case opts.catalog:
		printDecks(cards)
	case opts.list:
		printList(cards, opts)
	case opts.due:
		printDue(cards, s, opts)
	case opts.stats:
		printStats(cards, s, opts)
	default:
		if err := runReview(root, cards, s, opts); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
