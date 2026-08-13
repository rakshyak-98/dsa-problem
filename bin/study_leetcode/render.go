package main

import (
	"fmt"
	"os"
	"strings"
)

func renderDailyMarkdown(set practiceSet, fetchedAt string) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "# Daily LeetCode — %s\n\n", todayDate())
	fmt.Fprintf(&b, "**%s | %s**\n\n", set.day, set.topic)
	fmt.Fprintf(&b, "Reflex drill: `%s`\n\n", set.reflex)
	fmt.Fprintf(&b, "Warmup: %s\n\n", set.warmup)
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
		fmt.Fprintf(&b, "**%s** · %s%s\n\n", p.Diff, p.Pattern, dailyMark)
		fmt.Fprintf(&b, "%s\n\n", lcURL(p.Slug))
		if p.ReflexFn != "" {
			fmt.Fprintf(&b, "Reflex: `%s`\n\n", p.ReflexFn)
		}
		fmt.Fprintf(&b, "%s\n\n", htmlToText(raw))
		if i < len(set.problems)-1 {
			b.WriteString("---\n\n")
		}
	}
	return b.String(), nil
}

func dailyMarkdownIsCurrent(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	header := "# Daily LeetCode — " + todayDate()
	return strings.HasPrefix(string(data), header)
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
