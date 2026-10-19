package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Grading a drill
//
// Every test runs in its own process. A blank stub panics with
// "Implement from memory", and in one shared `go test` run that panic killed
// every test after it, so half-written drills were logged as all-fail. Here a
// blank function is reported as todo and the others are still graded.
//
// Each run is also capped in time and memory: an infinite loop that keeps
// appending (a traversal that forgets to step right) can otherwise eat all RAM
// before any test timeout fires.

type fnStatus int

const (
	statusPass fnStatus = iota
	statusFail
	statusTodo    // still the blank stub
	statusTimeout // ran past drillTestTimeout or drillMemLimit
	statusNoTest  // no Test<Function> in the drill's tests
)

func (s fnStatus) String() string {
	return [...]string{"pass", "fail", "todo", "timeout", "no test"}[s]
}

type fnResult struct {
	fn     string
	file   string
	status fnStatus
	detail string // first failing assert, timeout reason, or panic line
	output string // PASS:/FAIL: lines from the run
}

const (
	drillTestTimeout = 2 * time.Second
	drillMemLimit    = 1 << 30 // bytes of resident memory; Linux only (read from /proc)
	blankPanic       = "Implement from memory"
)

// gradeFunctions builds file's tests once and runs the test for each of fns on
// its own. A build error (the drill does not compile) is returned as err with
// the compiler output, and nothing is graded.
func gradeFunctions(drillPath, file string, fns []string) ([]fnResult, error) {
	tmp, err := os.MkdirTemp("", "drill-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmp)

	bin := filepath.Join(tmp, "drill.test")
	build := exec.Command("go", "test", "-c", "-vet=off", "-o", bin, ".")
	build.Dir = drillPath
	if out, err := build.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("%s does not compile:\n%s", file, strings.TrimSpace(string(out)))
	}

	list, err := exec.Command(bin, "-test.list", ".").Output()
	if err != nil {
		return nil, fmt.Errorf("list tests in %s: %w", file, err)
	}
	tests := strings.Fields(string(list))

	results := make([]fnResult, 0, len(fns))
	for _, fn := range fns {
		test := testFor(fn, tests)
		if test == "" {
			results = append(results, fnResult{fn: fn, file: file, status: statusNoTest})
			continue
		}
		r := runOneTest(bin, drillPath, test)
		r.fn, r.file = fn, file
		results = append(results, r)
	}
	return results, nil
}

// testFor finds Test<Function> for fn, ignoring case (TestGCD tests gcd).
func testFor(fn string, tests []string) string {
	for _, t := range tests {
		if strings.EqualFold(t, "Test"+fn) {
			return t
		}
	}
	return ""
}

func runOneTest(bin, dir, test string) fnResult {
	ctx, cancel := context.WithTimeout(context.Background(), drillTestTimeout+3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin,
		"-test.run", "^"+test+"$", "-test.v", "-test.timeout", drillTestTimeout.String())
	cmd.Dir = dir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Start(); err != nil {
		return fnResult{status: statusFail, detail: err.Error()}
	}

	overMem := make(chan bool, 1)
	done := make(chan struct{})
	go watchMemory(cmd.Process.Pid, drillMemLimit, done, overMem, cancel)
	err := cmd.Wait()
	close(done)

	out := buf.String()
	passed, failed := parseTestOutput(out)
	r := fnResult{output: assertLines(passed, failed)}
	switch {
	case len(overMem) > 0:
		r.status, r.detail = statusTimeout, fmt.Sprintf("used over %d MB — a loop that never ends?", drillMemLimit>>20)
	case strings.Contains(out, "test timed out") || ctx.Err() != nil:
		r.status, r.detail = statusTimeout, fmt.Sprintf("still running after %s — a loop that never ends?", drillTestTimeout)
	case strings.Contains(out, blankPanic):
		r.status = statusTodo
	case err == nil:
		r.status = statusPass
	default:
		r.status = statusFail
		if len(failed) > 0 {
			r.detail = failed[0]
		} else {
			r.detail = firstPanicLine(out)
		}
	}
	return r
}

func assertLines(passed, failed []string) string {
	var b strings.Builder
	for _, p := range passed {
		b.WriteString("PASS: " + p + "\n")
	}
	for _, f := range failed {
		b.WriteString("FAIL: " + f + "\n")
	}
	return b.String()
}

func firstPanicLine(out string) string {
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "panic: ") {
			return strings.TrimSpace(line)
		}
	}
	return "test failed (no assert output)"
}

// watchMemory kills the run when its resident memory passes limit. It reads
// /proc/<pid>/statm, so it only acts on Linux; elsewhere the time limit is the
// only guard.
func watchMemory(pid int, limit int64, done <-chan struct{}, over chan<- bool, kill func()) {
	path := fmt.Sprintf("/proc/%d/statm", pid)
	page := int64(os.Getpagesize())
	tick := time.NewTicker(25 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-done:
			return
		case <-tick.C:
			data, err := os.ReadFile(path)
			if err != nil {
				return
			}
			fields := strings.Fields(string(data))
			if len(fields) < 2 {
				return
			}
			pages, err := strconv.ParseInt(fields[1], 10, 64)
			if err == nil && pages*page > limit {
				over <- true
				kill()
				return
			}
		}
	}
}

// recordGrades writes one day's results into the log and reschedules each
// function. Only the first graded run of a day moves the schedule; later runs
// the same day just record that (and when) you got it passing. Todo and
// no-test results are not recorded.
func recordGrades(log *drillLog, results []fnResult, now time.Time) {
	day := dateOf(now)
	for _, r := range results {
		if r.status == statusTodo || r.status == statusNoTest {
			continue
		}
		passed := r.status == statusPass
		rec := log.Functions[r.fn]
		n := len(rec.Attempts)
		if n > 0 && rec.Attempts[n-1].Date == day {
			a := &rec.Attempts[n-1]
			a.Runs++
			if passed && !a.Passed {
				a.Passed = true
				a.SecsToPass = secsSince(a.Start, now)
			}
		} else {
			start := now.Format(time.RFC3339)
			if t, err := time.Parse(time.RFC3339, rec.Started); err == nil && dateOf(t) == day {
				start = rec.Started
			}
			a := attempt{Date: day, FirstTry: passed, Passed: passed, Runs: 1, Start: start}
			if passed {
				a.SecsToPass = secsSince(start, now)
			}
			rec.Attempts = append(rec.Attempts, a)
			if len(rec.Attempts) > maxAttempts {
				rec.Attempts = rec.Attempts[len(rec.Attempts)-maxAttempts:]
			}
			if passed {
				rec.Passes++
				rec.LastPass = day
			} else {
				rec.Fails++
				rec.LastFail = day
			}
			rec.Box, rec.Due = nextReview(rec.Box, passed, false, now)
		}
		log.Functions[r.fn] = rec
	}
}

func secsSince(start string, now time.Time) int {
	t, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return 0
	}
	if s := int(now.Sub(t).Seconds()); s > 0 {
		return s
	}
	return 0
}

// printGrades prints the per-assert lines and then one line per function: its
// status, whether this run moved the schedule, and a hint for a first-try miss.
func printGrades(log drillLog, results []fnResult, now time.Time) {
	for _, r := range results {
		fmt.Print(r.output)
	}
	day := dateOf(now)
	fmt.Println("\n── RESULT ─────────────────────────────────────────────")
	for _, r := range results {
		mark := map[fnStatus]string{statusPass: "✓", statusFail: "✗", statusTodo: "·", statusTimeout: "⏱", statusNoTest: "?"}[r.status]
		line := fmt.Sprintf("  %s %-24s %-8s", mark, r.fn, r.status)
		rec := log.Functions[r.fn]
		var a attempt
		if n := len(rec.Attempts); n > 0 && rec.Attempts[n-1].Date == day {
			a = rec.Attempts[n-1]
		}
		switch {
		case r.status == statusTodo:
			line += " not written yet"
		case r.status == statusNoTest:
			line += " no Test" + strings.ToUpper(r.fn[:1]) + r.fn[1:] + " in main_test.go"
		case a.Runs == 1:
			verdict := "first try ✓"
			if !a.FirstTry {
				verdict = "first try ✗"
			}
			line += fmt.Sprintf(" %s · next review %s", verdict, rec.Due)
			if a.Passed && a.SecsToPass > 0 {
				line += " · " + formatSecs(a.SecsToPass)
			}
		case r.status == statusPass:
			line += fmt.Sprintf(" passing after %d runs", a.Runs)
			if a.SecsToPass > 0 {
				line += " · " + formatSecs(a.SecsToPass)
			}
			line += " · schedule already set today"
		}
		if r.detail != "" {
			line += "\n      " + r.detail
		}
		fmt.Println(line)
		if a.Runs == 1 && !a.FirstTry {
			if h, ok := fnHints[r.fn]; ok {
				fmt.Printf("      hint: %s\n", h)
			}
		}
	}
}

func formatSecs(s int) string {
	if s < 60 {
		return fmt.Sprintf("%ds", s)
	}
	return fmt.Sprintf("%dm%02ds", s/60, s%60)
}

// gradeAndLog grades fns (grouped by drill file), records the results, and
// returns whether every one passed.
func gradeAndLog(repoRoot string, fns []string) bool {
	now := time.Now()
	log := loadLog(repoRoot)
	pruneLogNoise(&log)
	pruneRetired(repoRoot, &log)

	allPass := true
	var results []fnResult
	for _, g := range groupByFile(fns) {
		rs, err := gradeFunctions(writeReflexDir(repoRoot, g.file), g.file, g.fns)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			allPass = false
			continue
		}
		results = append(results, rs...)
	}
	recordGrades(&log, results, now)
	if err := saveLog(repoRoot, log); err != nil {
		fmt.Fprintf(os.Stderr, "%s: save drill log: %v\n", playProg, err)
	}
	printGrades(log, results, now)
	for _, r := range results {
		if r.status != statusPass {
			allPass = false
		}
	}
	return allPass
}

type fileGroup struct {
	file string
	fns  []string
}

// groupByFile splits fns by the drill file that holds them, in catalog order.
func groupByFile(fns []string) []fileGroup {
	want := map[string]bool{}
	for _, fn := range fns {
		want[fn] = true
	}
	var out []fileGroup
	for _, d := range drills {
		g := fileGroup{file: d.file}
		for _, fn := range d.functions {
			if want[fn] {
				g.fns = append(g.fns, fn)
			}
		}
		if len(g.fns) > 0 {
			out = append(out, g)
		}
	}
	return out
}

// drillFileOf returns the drill file holding fn.
func drillFileOf(fn string) (string, bool) {
	for _, d := range drills {
		for _, f := range d.functions {
			if f == fn {
				return d.file, true
			}
		}
	}
	return "", false
}
