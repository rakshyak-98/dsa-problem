package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type lcProblem struct {
	Num      int    `json:"num"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Diff     string `json:"diff"`
	Pattern  string `json:"pattern"`
	ReflexFn string `json:"reflexFn,omitempty"`
	Daily    bool   `json:"daily,omitempty"`
}

type practiceSet struct {
	day       string
	topic     string
	reflex    string
	topicTags []string
	seedSlugs []string
	warmup    string
	suggested []string
	problems  []lcProblem
}

type dailyFile struct {
	Date      string      `json:"date"`
	Day       string      `json:"day"`
	Topic     string      `json:"topic"`
	Reflex    string      `json:"reflex"`
	Warmup    string      `json:"warmup"`
	Suggested []string    `json:"suggested"`
	FetchedAt string      `json:"fetchedAt"`
	Problems  []lcProblem `json:"problems"`
}

func todayDate() string {
	return time.Now().Format("2006-01-02")
}

func loadDailyFile(path string) (dailyFile, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return dailyFile{}, false
	}
	var f dailyFile
	if err := json.Unmarshal(data, &f); err != nil {
		return dailyFile{}, false
	}
	return f, true
}

func saveDailyFile(path string, f dailyFile) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}

func fileToSet(f dailyFile) practiceSet {
	return practiceSet{
		day:       f.Day,
		topic:     f.Topic,
		reflex:    f.Reflex,
		warmup:    f.Warmup,
		suggested: f.Suggested,
		problems:  f.Problems,
	}
}

func setToFile(s practiceSet) dailyFile {
	return dailyFile{
		Date:      todayDate(),
		Day:       s.day,
		Topic:     s.topic,
		Reflex:    s.reflex,
		Warmup:    s.warmup,
		Suggested: s.suggested,
		FetchedAt: time.Now().UTC().Format(time.RFC3339),
		Problems:  s.problems,
	}
}

func ensureTodaySet(repoRoot string, refresh bool) (practiceSet, error) {
	meta := todayMeta()
	path := dailyJSONPath(repoRoot)

	if !refresh {
		if cached, ok := loadDailyFile(path); ok && cached.Date == todayDate() && len(cached.Problems) == 10 {
			return fileToSet(cached), nil
		}
	}

	problems, err := fetchTodayProblems(meta)
	if err != nil {
		if cached, ok := loadDailyFile(path); ok && len(cached.Problems) > 0 {
			fmt.Fprintf(os.Stderr, "warning: leetcode fetch failed (%v); using cached set from %s\n", err, cached.Date)
			return fileToSet(cached), nil
		}
		return practiceSet{}, err
	}

	set := practiceSet{
		day:       meta.day,
		topic:     meta.topic,
		reflex:    meta.reflex,
		warmup:    meta.warmup,
		suggested: meta.suggested,
		problems:  problems,
	}
	if err := saveDailyFile(path, setToFile(set)); err != nil {
		return set, fmt.Errorf("save daily.json: %w", err)
	}
	return set, nil
}
