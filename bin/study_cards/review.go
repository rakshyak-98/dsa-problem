package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// rand is auto-seeded in Go 1.20+.

type queueItem struct {
	card  Card
	prog  *Progress
	isNew bool
}

func buildQueue(cards []Card, s *store, opts options, on time.Time) []queueItem {
	var due, neu []queueItem
	for _, c := range cards {
		if !cardMatches(c, opts.deck, opts.tag) {
			continue
		}
		p := getProgress(s, c.ID)
		if p == nil || p.isNew() {
			neu = append(neu, queueItem{card: c, isNew: true})
			continue
		}
		if p.isDue(on) {
			due = append(due, queueItem{card: c, prog: p})
		}
	}
	if opts.shuffle {
		rand.Shuffle(len(due), func(i, j int) { due[i], due[j] = due[j], due[i] })
		rand.Shuffle(len(neu), func(i, j int) { neu[i], neu[j] = neu[j], neu[i] })
	}
	if opts.newN >= 0 && len(neu) > opts.newN {
		neu = neu[:opts.newN]
	}
	q := append(due, neu...)
	if opts.limit > 0 && len(q) > opts.limit {
		q = q[:opts.limit]
	}
	return q
}

func runReview(root string, cards []Card, s *store, opts options) error {
	on := today()
	q := buildQueue(cards, s, opts, on)
	if len(q) == 0 {
		fmt.Println("Nothing due. Nice work.")
		if opts.deck != "" || opts.tag != "" {
			fmt.Println("Try dropping --deck/--tag, or raise --new to learn fresh cards.")
		} else {
			fmt.Println("Use: go run ./bin/study_cards -- --new 10")
		}
		return nil
	}

	dueCount, newCount := 0, 0
	for _, it := range q {
		if it.isNew {
			newCount++
		} else {
			dueCount++
		}
	}
	fmt.Printf("Session: %d cards (%d due, %d new)", len(q), dueCount, newCount)
	if opts.deck != "" {
		fmt.Printf("  deck=%s", opts.deck)
	}
	if opts.tag != "" {
		fmt.Printf("  tag=%s", opts.tag)
	}
	fmt.Println()
	fmt.Println("Rate: 1=again  2=hard  3=good  4=easy  |  s=show  q=quit")
	fmt.Println(strings.Repeat("─", 60))

	in := bufio.NewReader(os.Stdin)
	reviewed := 0
	for i, it := range q {
		c := it.card
		fmt.Printf("\n[%d/%d]  %s", i+1, len(q), c.Deck)
		if c.Section != "" {
			fmt.Printf(" · %s", c.Section)
		}
		if it.isNew {
			fmt.Print("  (new)")
		}
		fmt.Println()
		fmt.Println()
		fmt.Println(wrap(c.Front, 72))
		fmt.Println()
		fmt.Print("Answer ready? [Enter=show, q=quit] ")
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(strings.ToLower(line))
		if line == "q" || line == "quit" {
			break
		}

		fmt.Println()
		fmt.Println("── answer ──")
		fmt.Println(wrap(c.Back, 72))
		if c.Source != "" {
			fmt.Printf("\n(source: %s)\n", c.Source)
		}
		fmt.Println()

		rating := 0
		for rating == 0 {
			fmt.Print("Rate 1–4 (again/hard/good/easy) or q: ")
			line, err = in.ReadString('\n')
			if err != nil {
				rating = -1
				break
			}
			line = strings.TrimSpace(strings.ToLower(line))
			switch line {
			case "q", "quit":
				rating = -1
			case "1", "a", "again":
				rating = 1
			case "2", "h", "hard":
				rating = 2
			case "3", "g", "good":
				rating = 3
			case "4", "e", "easy":
				rating = 4
			default:
				fmt.Println("  enter 1, 2, 3, 4, or q")
			}
		}
		if rating < 0 {
			break
		}
		p := ensureProgress(s, c.ID)
		applyRating(p, rating, on)
		reviewed++
		fmt.Printf("  → next due %s (interval %dd)\n", p.Due, p.Interval)
	}

	if err := saveStore(root, s); err != nil {
		return err
	}
	fmt.Printf("\nSaved. Reviewed %d card(s). Progress → cards/.srs_progress.json\n", reviewed)
	return nil
}

func printDue(cards []Card, s *store, opts options) {
	on := today()
	due, neu, later := 0, 0, 0
	for _, c := range cards {
		if !cardMatches(c, opts.deck, opts.tag) {
			continue
		}
		p := getProgress(s, c.ID)
		if p == nil || p.isNew() {
			neu++
			continue
		}
		if p.isDue(on) {
			due++
		} else {
			later++
		}
	}
	fmt.Printf("Due today:  %d\n", due)
	fmt.Printf("New left:   %d\n", neu)
	fmt.Printf("Not yet:    %d\n", later)
	fmt.Printf("Filter:     deck=%q tag=%q\n", opts.deck, opts.tag)
}

func printStats(cards []Card, s *store, opts options) {
	on := today()
	type agg struct {
		total, neu, due, learning, mature int
		reviews, lapses                   int
	}
	byDeck := map[string]*agg{}
	var all agg
	for _, c := range cards {
		if !cardMatches(c, opts.deck, opts.tag) {
			continue
		}
		a := byDeck[c.Deck]
		if a == nil {
			a = &agg{}
			byDeck[c.Deck] = a
		}
		a.total++
		all.total++
		p := getProgress(s, c.ID)
		if p == nil || p.isNew() {
			a.neu++
			all.neu++
			continue
		}
		a.reviews += p.Reviews
		a.lapses += p.Lapses
		all.reviews += p.Reviews
		all.lapses += p.Lapses
		if p.isDue(on) {
			a.due++
			all.due++
		}
		if p.Interval >= 21 {
			a.mature++
			all.mature++
		} else {
			a.learning++
			all.learning++
		}
	}
	fmt.Println("Spaced repetition stats")
	fmt.Println(strings.Repeat("─", 60))
	fmt.Printf("%-12s %6s %6s %6s %8s %8s\n", "deck", "total", "new", "due", "learning", "mature")
	names := make([]string, 0, len(byDeck))
	for n := range byDeck {
		names = append(names, n)
	}
	sortStrings(names)
	for _, n := range names {
		a := byDeck[n]
		fmt.Printf("%-12s %6d %6d %6d %8d %8d\n", n, a.total, a.neu, a.due, a.learning, a.mature)
	}
	fmt.Println(strings.Repeat("─", 60))
	fmt.Printf("%-12s %6d %6d %6d %8d %8d\n", "ALL", all.total, all.neu, all.due, all.learning, all.mature)
	fmt.Printf("Total reviews logged: %d   lapses: %d\n", all.reviews, all.lapses)
}

func sortStrings(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func printDecks(cards []Card) {
	counts := map[string]int{}
	for _, c := range cards {
		counts[c.Deck]++
	}
	names := make([]string, 0, len(counts))
	for n := range counts {
		names = append(names, n)
	}
	sortStrings(names)
	fmt.Println("Decks (from cards/decks/):")
	total := 0
	for _, n := range names {
		fmt.Printf("  %-12s %4d cards\n", n, counts[n])
		total += counts[n]
	}
	fmt.Printf("  %-12s %4d cards\n", "TOTAL", total)
}

func printList(cards []Card, opts options) {
	n := 0
	for _, c := range cards {
		if !cardMatches(c, opts.deck, opts.tag) {
			continue
		}
		n++
		if opts.limit > 0 && n > opts.limit {
			fmt.Printf("… truncated at --limit %d\n", opts.limit)
			return
		}
		fmt.Printf("%s  [%s] %s\n", c.ID, c.Deck, c.Front)
	}
	fmt.Printf("(%d cards)\n", n)
}

func runReset(root string) error {
	fmt.Print("Reset all SRS progress? Type YES to confirm: ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	if strings.TrimSpace(line) != "YES" {
		fmt.Println("Cancelled.")
		return nil
	}
	s := &store{Version: 1, Cards: map[string]*Progress{}}
	if err := saveStore(root, s); err != nil {
		return err
	}
	fmt.Println("Progress cleared.")
	return nil
}
