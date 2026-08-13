package main

import (
	"fmt"
	"hash/fnv"
)

func shuffledSeeds(slugs []string, date, day string) []string {
	cp := append([]string(nil), slugs...)
	h := fnv.New64a()
	_, _ = h.Write([]byte(date + day))
	seed := int64(h.Sum64())
	for i := len(cp) - 1; i > 0; i-- {
		seed = seed*1664525 + 1013904223
		j := int(seed % int64(i+1))
		if j < 0 {
			j = -j
		}
		cp[i], cp[j] = cp[j], cp[i]
	}
	return cp
}

func fetchTodayProblems(meta practiceSet) ([]lcProblem, error) {
	seen := map[string]bool{}
	out := make([]lcProblem, 0, 10)

	daily, derr := fetchDailyChallenge()
	if derr == nil {
		qDaily, err := fetchQuestionRaw(daily.Slug)
		if err == nil && !qDaily.IsPaidOnly && topicMatchQuestion(qDaily, meta.topicTags) {
			p := questionToProblem(qDaily, "")
			p.Daily = true
			out = append(out, p)
			seen[p.Slug] = true
		}
	}

	slugs := shuffledSeeds(meta.seedSlugs, todayDate(), meta.day)
	for _, slug := range slugs {
		if len(out) >= 10 {
			break
		}
		if seen[slug] {
			continue
		}
		q, err := fetchQuestionRaw(slug)
		if err != nil {
			continue
		}
		if q.IsPaidOnly {
			continue
		}
		if !topicMatchQuestion(q, meta.topicTags) {
			continue
		}
		p := questionToProblem(q, "")
		out = append(out, p)
		seen[slug] = true
	}

	if len(out) < 10 {
		return nil, fmt.Errorf("only fetched %d/10 problems for %s", len(out), meta.topic)
	}
	sortProblemsByDifficulty(out)
	return out, nil
}

func fetchQuestionRaw(slug string) (gqlQuestion, error) {
	var resp struct {
		Data struct {
			Question gqlQuestion `json:"question"`
		} `json:"data"`
	}
	err := postGraphQL(gqlRequest{
		Query: `query getQuestion($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				questionFrontendId
				title
				titleSlug
				difficulty
				isPaidOnly
				topicTags { name slug }
			}
		}`,
		Variables:     map[string]any{"titleSlug": slug},
		OperationName: "getQuestion",
	}, &resp)
	if err != nil {
		return gqlQuestion{}, err
	}
	if resp.Data.Question.TitleSlug == "" {
		return gqlQuestion{}, fmt.Errorf("question not found: %s", slug)
	}
	return resp.Data.Question, nil
}

func topicMatchQuestion(q gqlQuestion, required []string) bool {
	return matchesTopicTags(q, required)
}
