package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Card is one flashcard loaded from cards/decks/*.json.
type Card struct {
	ID      string   `json:"id"`
	Deck    string   `json:"deck"`
	Tags    []string `json:"tags"`
	Front   string   `json:"front"`
	Back    string   `json:"back"`
	Source  string   `json:"source"`
	Section string   `json:"section"`
}

// Progress tracks SM-2-lite state for one card.
// Intervals follow the study plan: +1, +3, +7, +21 days (then grow).
type Progress struct {
	Due        string  `json:"due"`         // YYYY-MM-DD
	Interval   int     `json:"interval"`    // days until next review
	Ease       float64 `json:"ease"`        // SM-2 ease factor
	Reps       int     `json:"reps"`        // successful reviews in a row
	Lapses     int     `json:"lapses"`      // times forgotten
	Last       string  `json:"last"`        // YYYY-MM-DD of last review
	LastRating int     `json:"last_rating"` // 1=again .. 4=easy
	Reviews    int     `json:"reviews"`     // total reviews
}

type store struct {
	Version int                  `json:"version"`
	Updated string               `json:"updated"`
	Cards   map[string]*Progress `json:"cards"`
}

func cardsRoot(start string) string {
	candidates := []string{
		filepath.Join(start, "cards"),
		filepath.Join(start, "..", "cards"),
		filepath.Join(start, "..", "..", "cards"),
	}
	for _, c := range candidates {
		if st, err := os.Stat(c); err == nil && st.IsDir() {
			abs, _ := filepath.Abs(c)
			return abs
		}
	}
	return filepath.Join(start, "cards")
}

func loadCards(root string) ([]Card, error) {
	deckDir := filepath.Join(root, "decks")
	entries, err := os.ReadDir(deckDir)
	if err != nil {
		return nil, fmt.Errorf("read decks: %w (expected %s)", err, deckDir)
	}
	var all []Card
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(deckDir, e.Name()))
		if err != nil {
			return nil, err
		}
		var deck []Card
		if err := json.Unmarshal(raw, &deck); err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		all = append(all, deck...)
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].Deck != all[j].Deck {
			return all[i].Deck < all[j].Deck
		}
		return all[i].ID < all[j].ID
	})
	return all, nil
}

func progressPath(root string) string {
	return filepath.Join(root, ".srs_progress.json")
}

func loadStore(root string) (*store, error) {
	path := progressPath(root)
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &store{Version: 1, Cards: map[string]*Progress{}}, nil
		}
		return nil, err
	}
	var s store
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, err
	}
	if s.Cards == nil {
		s.Cards = map[string]*Progress{}
	}
	return &s, nil
}

func saveStore(root string, s *store) error {
	s.Updated = time.Now().Format(time.RFC3339)
	s.Version = 1
	raw, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(progressPath(root), append(raw, '\n'), 0o644)
}

func today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func parseDay(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

func formatDay(t time.Time) string {
	return t.Format("2006-01-02")
}

func (p *Progress) isNew() bool {
	return p == nil || (p.Reviews == 0 && p.Due == "" && p.Reps == 0)
}

func (p *Progress) isDue(on time.Time) bool {
	if p == nil || p.isNew() {
		return false
	}
	due := parseDay(p.Due)
	if due.IsZero() {
		return true
	}
	return !due.After(on)
}

// scheduleIntervals mirrors doc/write/STUDY_PLAN.md spaced repetition.
var scheduleIntervals = []int{1, 3, 7, 21}

func applyRating(p *Progress, rating int, on time.Time) {
	if p.Ease == 0 {
		p.Ease = 2.5
	}
	p.Last = formatDay(on)
	p.LastRating = rating
	p.Reviews++

	switch rating {
	case 1: // again — reset
		p.Lapses++
		p.Reps = 0
		p.Interval = 1
		p.Ease = maxFloat(1.3, p.Ease-0.2)
	case 2: // hard — short step
		if p.Reps == 0 {
			p.Interval = 1
		} else {
			p.Interval = maxInt(1, int(float64(p.Interval)*1.2))
		}
		p.Ease = maxFloat(1.3, p.Ease-0.15)
		p.Reps++
	case 3: // good — study-plan ladder then ease growth
		if p.Reps < len(scheduleIntervals) {
			p.Interval = scheduleIntervals[p.Reps]
		} else {
			p.Interval = maxInt(1, int(float64(p.Interval)*p.Ease))
		}
		p.Reps++
	case 4: // easy — skip ahead one rung
		p.Ease += 0.15
		if p.Reps+1 < len(scheduleIntervals) {
			p.Interval = scheduleIntervals[p.Reps+1]
			p.Reps += 2
		} else if p.Reps < len(scheduleIntervals) {
			p.Interval = scheduleIntervals[p.Reps]
			p.Reps++
			p.Interval = maxInt(p.Interval, int(float64(p.Interval)*1.3))
		} else {
			p.Interval = maxInt(1, int(float64(p.Interval)*p.Ease*1.3))
			p.Reps++
		}
	default:
		p.Interval = 1
	}
	p.Due = formatDay(on.AddDate(0, 0, p.Interval))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getProgress(s *store, id string) *Progress {
	if p, ok := s.Cards[id]; ok {
		return p
	}
	return nil
}

func ensureProgress(s *store, id string) *Progress {
	if p, ok := s.Cards[id]; ok {
		return p
	}
	p := &Progress{Ease: 2.5}
	s.Cards[id] = p
	return p
}

func cardMatches(c Card, deck, tag string) bool {
	if deck != "" && !strings.EqualFold(c.Deck, deck) {
		return false
	}
	if tag == "" {
		return true
	}
	tag = strings.ToLower(tag)
	for _, t := range c.Tags {
		if strings.ToLower(t) == tag {
			return true
		}
	}
	return strings.Contains(strings.ToLower(c.Deck), tag) ||
		strings.Contains(strings.ToLower(c.Section), tag)
}

func wrap(text string, width int) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}
	var lines []string
	var line strings.Builder
	for _, w := range words {
		if line.Len() == 0 {
			line.WriteString(w)
			continue
		}
		if line.Len()+1+len(w) > width {
			lines = append(lines, line.String())
			line.Reset()
			line.WriteString(w)
			continue
		}
		line.WriteByte(' ')
		line.WriteString(w)
	}
	if line.Len() > 0 {
		lines = append(lines, line.String())
	}
	return strings.Join(lines, "\n")
}
