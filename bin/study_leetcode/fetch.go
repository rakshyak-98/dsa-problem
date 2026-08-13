package main

import (
	"fmt"
	"hash/fnv"
)

func seededShuffle[T any](items []T, date, day, salt string) []T {
	cp := append([]T(nil), items...)
	h := fnv.New64a()
	_, _ = h.Write([]byte(date + day + salt))
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

func shuffledSeeds(slugs []string, date, day string) []string {
	out := seededShuffle(slugs, date, day, "seeds")
	s := make([]string, len(out))
	for i, slug := range out {
		s[i] = slug
	}
	return s
}

func topicFetchSkip(date, day, tag string, total, limit int) int {
	maxSkip := total - limit
	if maxSkip <= 0 {
		return 0
	}
	h := fnv.New64a()
	_, _ = h.Write([]byte(date + day + tag + "skip"))
	seed := int64(h.Sum64())
	return int(seed % int64(maxSkip+1))
}

func fetchTopicProblemPool(meta practiceSet) ([]gqlQuestion, error) {
	date := todayDate()
	bySlug := map[string]gqlQuestion{}

	for _, tag := range meta.topicTags {
		total, _, err := fetchProblemsByTag(tag, 1, 0)
		if err != nil {
			return nil, fmt.Errorf("topic %q: %w", tag, err)
		}
		if total == 0 {
			continue
		}
		limit := topicPoolPageSize
		if total < limit {
			limit = total
		}
		skip := topicFetchSkip(date, meta.day, tag, total, limit)
		_, questions, err := fetchProblemsByTag(tag, limit, skip)
		if err != nil {
			return nil, fmt.Errorf("topic %q page: %w", tag, err)
		}
		for _, q := range questions {
			if q.TitleSlug == "" || q.IsPaidOnly {
				continue
			}
			if !topicMatchQuestion(q, meta.topicTags) {
				continue
			}
			bySlug[q.TitleSlug] = q
		}
	}

	pool := make([]gqlQuestion, 0, len(bySlug))
	for _, q := range bySlug {
		pool = append(pool, q)
	}
	if len(pool) == 0 {
		return nil, fmt.Errorf("no free problems for tags %v", meta.topicTags)
	}
	return seededShuffle(pool, date, meta.day, "pool"), nil
}

func appendUniqueProblems(out []lcProblem, seen map[string]bool, candidates []gqlQuestion, limit int) []lcProblem {
	for _, q := range candidates {
		if len(out) >= limit {
			break
		}
		if seen[q.TitleSlug] {
			continue
		}
		p := questionToProblem(q, "")
		out = append(out, p)
		seen[p.Slug] = true
	}
	return out
}

func fetchTodayProblems(meta practiceSet) ([]lcProblem, error) {
	const want = 10
	seen := map[string]bool{}
	out := make([]lcProblem, 0, want)

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

	pool, poolErr := fetchTopicProblemPool(meta)
	if poolErr == nil {
		out = appendUniqueProblems(out, seen, pool, want)
	}

	if len(out) < want {
		slugs := shuffledSeeds(meta.seedSlugs, todayDate(), meta.day)
		for _, slug := range slugs {
			if len(out) >= want {
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
	}

	if len(out) < want {
		if poolErr != nil {
			return nil, fmt.Errorf("only fetched %d/%d problems for %s: %v", len(out), want, meta.topic, poolErr)
		}
		return nil, fmt.Errorf("only fetched %d/%d problems for %s", len(out), want, meta.topic)
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
