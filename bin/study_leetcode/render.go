package main

import (
	"fmt"
	"os"
	"strings"
)

func renderDailyMarkdown(set practiceSet, fetchedAt string) (string, error) {
	var b strings.Builder
	// The topic, tags and reflex links name the pattern, which is the thing
	// to work out yourself, so they sit in an Answers section at the end.
	fmt.Fprintf(&b, "# Daily LeetCode — %s\n\n", todayDate())
	b.WriteString("Before each problem, write down the pattern and the words that point to it.\n")
	b.WriteString("Then check the Answers section at the end, solve, and record it with\n")
	b.WriteString("`go run . -- --log=SLUG`.\n\n")
	if fetchedAt != "" {
		fmt.Fprintf(&b, "Fetched: %s\n\n", fetchedAt)
	}
	fmt.Fprintf(&b, "---\n\n")

	for i, p := range set.problems {
		raw, err := fetchQuestionContent(p.Slug)
		if err != nil {
			return "", fmt.Errorf("problem #%d %s: %w", p.Num, p.Slug, err)
		}
		dailyMark := ""
		if p.Daily {
			dailyMark = " · daily challenge"
		}
		fmt.Fprintf(&b, "## %d. %s (#%d)\n\n", i+1, p.Title, p.Num)
		fmt.Fprintf(&b, "**%s**%s\n\n", p.Diff, dailyMark)
		fmt.Fprintf(&b, "%s\n\n", lcURL(p.Slug))
		fmt.Fprintf(&b, "%s\n\n", htmlToText(raw))
		b.WriteString("---\n\n")
	}
	b.WriteString(renderAnswersMarkdown(set))
	return b.String(), nil
}

// renderAnswersMarkdown lists what each problem tests, for checking your guess.
func renderAnswersMarkdown(set practiceSet) string {
	var b strings.Builder
	b.WriteString("## Answers (check after naming each pattern)\n\n")
	if set.topic != "" {
		fmt.Fprintf(&b, "Topic: %s · reflex drill `%s`\n\n", set.topic, set.reflex)
	}
	if set.warmup != "" {
		fmt.Fprintf(&b, "Warmup: %s\n\n", set.warmup)
	}
	for i, p := range set.problems {
		fmt.Fprintf(&b, "%d. %s (`%s`) — %s", i+1, p.Title, p.Slug, orNone(p.Pattern))
		if p.ReflexFn != "" {
			fmt.Fprintf(&b, " · reflex `%s`", p.ReflexFn)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func orNone(s string) string {
	if s == "" {
		return "(no tags)"
	}
	return s
}

func dailyMarkdownIsCurrent(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	header := "# Daily LeetCode — " + todayDate()
	return strings.HasPrefix(string(data), header)
}

func writeDailyOutputs(repoRoot string, set practiceSet, fetchedAt string, refresh bool) error {
	var msgs []string
	if err := writeDailyMarkdown(repoRoot, set, fetchedAt, refresh); err != nil {
		msgs = append(msgs, "daily.md: "+err.Error())
	}
	if err := writeDailyGo(repoRoot, set, fetchedAt, refresh); err != nil {
		msgs = append(msgs, "daily.go: "+err.Error())
	}
	if len(msgs) > 0 {
		return fmt.Errorf("%s", strings.Join(msgs, "; "))
	}
	return nil
}

func writeDailyMarkdown(repoRoot string, set practiceSet, fetchedAt string, refresh bool) error {
	path := dailyMarkdownPath(repoRoot)
	if !refresh && dailyMarkdownIsCurrent(path) {
		return nil
	}
	body, err := renderDailyMarkdown(set, fetchedAt)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(body), 0o644)
}
