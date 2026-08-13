package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultLeetcodeGraphQL = "https://leetcode.com/graphql"

var leetcodeGraphQL = defaultLeetcodeGraphQL

var httpClient = &http.Client{Timeout: 20 * time.Second}

type gqlRequest struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables,omitempty"`
	OperationName string         `json:"operationName,omitempty"`
}

type gqlQuestion struct {
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	Difficulty         string `json:"difficulty"`
	IsPaidOnly         bool   `json:"isPaidOnly"`
	TopicTags          []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"topicTags"`
}

func postGraphQL(req gqlRequest, out any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest(http.MethodPost, leetcodeGraphQL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Referer", "https://leetcode.com/")
	httpReq.Header.Set("Origin", "https://leetcode.com")
	httpReq.Header.Set("User-Agent", "dsa-problem-study-leetcode/1.0")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("leetcode api status %d: %s", resp.StatusCode, truncate(string(raw), 120))
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func fetchDailyChallenge() (lcProblem, error) {
	var resp struct {
		Data struct {
			Active struct {
				Date string `json:"date"`
				Link string `json:"link"`
				Q    gqlQuestion `json:"question"`
			} `json:"activeDailyCodingChallengeQuestion"`
		} `json:"data"`
		Errors []any `json:"errors"`
	}
	err := postGraphQL(gqlRequest{
		Query: `query questionOfToday {
			activeDailyCodingChallengeQuestion {
				date
				link
				question {
					questionFrontendId
					title
					titleSlug
					difficulty
					isPaidOnly
					topicTags { name slug }
				}
			}
		}`,
		OperationName: "questionOfToday",
	}, &resp)
	if err != nil {
		return lcProblem{}, err
	}
	q := resp.Data.Active.Q
	if q.TitleSlug == "" {
		return lcProblem{}, fmt.Errorf("empty daily challenge response")
	}
	p := questionToProblem(q, "")
	p.Daily = true
	return p, nil
}

func fetchQuestionBySlug(slug string) (lcProblem, error) {
	q, err := fetchQuestionRaw(slug)
	if err != nil {
		return lcProblem{}, err
	}
	return questionToProblem(q, ""), nil
}

func questionToProblem(q gqlQuestion, reflexFn string) lcProblem {
	num := parseQuestionNum(q.QuestionFrontendID)
	if reflexFn == "" {
		reflexFn = reflexBySlug[q.TitleSlug]
	}
	return lcProblem{
		Num:      num,
		Title:    q.Title,
		Slug:     q.TitleSlug,
		Diff:     q.Difficulty,
		Pattern:  patternFromTags(q.TopicTags),
		ReflexFn: reflexFn,
	}
}

func patternFromTags(tags []struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}) string {
	if len(tags) == 0 {
		return "leetcode"
	}
	names := make([]string, 0, len(tags))
	for _, t := range tags {
		names = append(names, strings.ToLower(t.Name))
	}
	return strings.Join(names, ", ")
}

const topicPoolPageSize = 50

func fetchProblemsByTag(tag string, limit, skip int) (int, []gqlQuestion, error) {
	if limit <= 0 {
		limit = topicPoolPageSize
	}
	var resp struct {
		Data struct {
			ProblemsetQuestionList struct {
				Total     int           `json:"totalNum"`
				Questions []gqlQuestion `json:"data"`
			} `json:"problemsetQuestionList"`
		} `json:"data"`
		Errors []any `json:"errors"`
	}
	err := postGraphQL(gqlRequest{
		Query: `query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
			problemsetQuestionList: questionList(
				categorySlug: $categorySlug
				limit: $limit
				skip: $skip
				filters: $filters
			) {
				totalNum
				data {
					questionFrontendId
					title
					titleSlug
					difficulty
					isPaidOnly
					topicTags { name slug }
				}
			}
		}`,
		Variables: map[string]any{
			"categorySlug": "",
			"skip":         skip,
			"limit":        limit,
			"filters":      map[string]any{"tags": []string{tag}},
		},
		OperationName: "problemsetQuestionList",
	}, &resp)
	if err != nil {
		return 0, nil, err
	}
	list := resp.Data.ProblemsetQuestionList
	return list.Total, list.Questions, nil
}

func matchesTopicTags(q gqlQuestion, required []string) bool {
	if len(required) == 0 {
		return true
	}
	have := map[string]bool{}
	for _, t := range q.TopicTags {
		have[t.Slug] = true
	}
	for _, want := range required {
		if have[want] {
			return true
		}
	}
	return false
}
