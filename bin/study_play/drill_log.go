package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type fnRecord struct {
	Passes   int    `json:"passes"`
	Fails    int    `json:"fails"`
	LastPass string `json:"lastPass,omitempty"`
	LastFail string `json:"lastFail,omitempty"`
}

type drillLog struct {
	Functions map[string]fnRecord `json:"functions"`
}

var passRE = regexp.MustCompile(`^PASS: (.+)$`)
var failRE = regexp.MustCompile(`FAIL: (.+)`)

func logPath(root string) string {
	return filepath.Join(root, ".drill_log.json")
}

func loadLog(root string) drillLog {
	path := logPath(root)
	data, err := os.ReadFile(path)
	if err != nil {
		return drillLog{Functions: map[string]fnRecord{}}
	}
	var log drillLog
	if err := json.Unmarshal(data, &log); err != nil || log.Functions == nil {
		return drillLog{Functions: map[string]fnRecord{}}
	}
	return log
}

func saveLog(root string, log drillLog) error {
	data, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(logPath(root), data, 0o644)
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func recordResult(log *drillLog, name string, passed bool) {
	rec := log.Functions[name]
	if passed {
		rec.Passes++
		rec.LastPass = today()
	} else {
		rec.Fails++
		rec.LastFail = today()
	}
	log.Functions[name] = rec
}

func parseTestOutput(output string) (passed, failed []string) {
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if m := passRE.FindStringSubmatch(line); len(m) == 2 {
			passed = append(passed, m[1])
		}
		if m := failRE.FindStringSubmatch(line); len(m) == 2 {
			failed = append(failed, m[1])
		}
	}
	return passed, failed
}

func hasTestFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), "_test.go") {
			return true
		}
	}
	return false
}

func runDrillWithLog(drillPath string) (bool, string, error) {
	var cmd *exec.Cmd
	if hasTestFiles(drillPath) {
		cmd = exec.Command("go", "test", "-v", "-count=1", "-vet=off", ".")
	} else {
		cmd = exec.Command("go", "run", ".")
	}
	cmd.Dir = drillPath
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	return err == nil, buf.String(), err
}

func updateLogFromOutput(root string, output string, allFunctions []string) {
	log := loadLog(root)
	passed, failed := parseTestOutput(output)

	seen := map[string]bool{}
	for _, name := range passed {
		recordResult(&log, name, true)
		seen[name] = true
	}
	for _, name := range failed {
		recordResult(&log, name, false)
		seen[name] = true
	}

	// If full pass with no per-test names, mark all drill functions
	if len(passed) == 0 && len(failed) == 0 {
		for _, fn := range allFunctions {
			recordResult(&log, fn, true)
		}
	} else {
		// Functions not reached before failure count as fail only if run failed
		for _, fn := range allFunctions {
			if !seen[fn] && len(failed) > 0 {
				recordResult(&log, fn, false)
			}
		}
	}

	_ = saveLog(root, log)
}

type weakEntry struct {
	name  string
	score float64
	fails int
}

func printWeakFunctions(root string, limit int) {
	log := loadLog(root)
	if len(log.Functions) == 0 {
		fmt.Println("No drill history yet. Run: go run . -- --run")
		return
	}

	var entries []weakEntry
	for name, rec := range log.Functions {
		total := rec.Passes + rec.Fails
		failRate := 0.0
		if total > 0 {
			failRate = float64(rec.Fails) / float64(total)
		}
		score := failRate*10 + float64(rec.Fails)
		if rec.LastPass == "" && rec.Fails > 0 {
			score += 5
		}
		entries = append(entries, weakEntry{name: name, score: score, fails: rec.Fails})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].score == entries[j].score {
			return entries[i].fails > entries[j].fails
		}
		return entries[i].score > entries[j].score
	})

	fmt.Println("\n── WEAK FUNCTIONS (review these first) ────────────────")
	for i, e := range entries {
		if i >= limit {
			break
		}
		rec := log.Functions[e.name]
		fmt.Printf("  %d. %s — fails:%d passes:%d lastPass:%s\n",
			i+1, e.name, rec.Fails, rec.Passes, orDash(rec.LastPass))
	}
	fmt.Println("\n  Tip: go run . -- --reset then blind-write weak functions.")
}

func orDash(s string) string {
	if s == "" {
		return "—"
	}
	return s
}
