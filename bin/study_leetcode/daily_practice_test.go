package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestFetchTodayProblemsMock(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req gqlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if req.OperationName == "questionOfToday" {
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"activeDailyCodingChallengeQuestion": map[string]any{
						"question": map[string]any{
							"questionFrontendId": "704",
							"title":              "Binary Search",
							"titleSlug":          "binary-search",
							"difficulty":         "Easy",
							"isPaidOnly":         false,
							"topicTags":          []map[string]string{{"name": "Binary Search", "slug": "binary-search"}},
						},
					},
				},
			})
			return
		}
		if req.OperationName == "problemsetQuestionList" {
			filters, _ := req.Variables["filters"].(map[string]any)
			tags, _ := filters["tags"].([]any)
			tag, _ := tags[0].(string)
			limit, _ := req.Variables["limit"].(float64)
			skip, _ := req.Variables["skip"].(float64)
			if limit <= 1 {
				json.NewEncoder(w).Encode(map[string]any{
					"data": map[string]any{
						"problemsetQuestionList": map[string]any{
							"totalNum": 12,
							"data":     []any{},
						},
					},
				})
				return
			}
			questions := make([]map[string]any, 0, int(limit))
			for i := 0; i < int(limit); i++ {
				idx := int(skip) + i + 1
				slug := fmt.Sprintf("topic-%s-%d", tag, idx)
				questions = append(questions, map[string]any{
					"questionFrontendId": fmt.Sprintf("%d", 200+idx),
					"title":              slug,
					"titleSlug":          slug,
					"difficulty":         "Easy",
					"isPaidOnly":         false,
					"topicTags":          []map[string]string{{"name": "Binary Search", "slug": "binary-search"}},
				})
			}
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"problemsetQuestionList": map[string]any{
						"totalNum": 12,
						"data":     questions,
					},
				},
			})
			return
		}
		slug, _ := req.Variables["titleSlug"].(string)
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"question": map[string]any{
					"questionFrontendId": "100",
					"title":              slug,
					"titleSlug":          slug,
					"difficulty":         "Easy",
					"isPaidOnly":         false,
					"topicTags":          []map[string]string{{"name": "Binary Search", "slug": "binary-search"}},
				},
			},
		})
	}))
	defer srv.Close()
	oldURL := leetcodeGraphQL
	leetcodeGraphQL = srv.URL
	defer func() { leetcodeGraphQL = oldURL }()

	meta := practiceSets[3]
	problems, err := fetchTodayProblems(meta)
	if err != nil {
		t.Fatalf("fetchTodayProblems: %v", err)
	}
	if len(problems) != 10 {
		t.Fatalf("expected 10 problems, got %d", len(problems))
	}
	dailyCount := 0
	for _, p := range problems {
		if p.Daily {
			dailyCount++
		}
	}
	if dailyCount != 1 {
		t.Fatalf("expected 1 daily challenge, got %d", dailyCount)
	}
}

func TestFetchTodayProblemsUsesTopicPoolNotSeeds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req gqlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		switch req.OperationName {
		case "questionOfToday":
			json.NewEncoder(w).Encode(map[string]any{"data": map[string]any{
				"activeDailyCodingChallengeQuestion": map[string]any{"question": map[string]any{}},
			}})
		case "problemsetQuestionList":
			limit, _ := req.Variables["limit"].(float64)
			if limit <= 1 {
				json.NewEncoder(w).Encode(map[string]any{"data": map[string]any{
					"problemsetQuestionList": map[string]any{"totalNum": 15, "data": []any{}},
				}})
				return
			}
			questions := make([]map[string]any, 15)
			for i := range questions {
				slug := fmt.Sprintf("pool-problem-%02d", i+1)
				questions[i] = map[string]any{
					"questionFrontendId": fmt.Sprintf("%d", 300+i),
					"title":              slug,
					"titleSlug":          slug,
					"difficulty":         "Easy",
					"isPaidOnly":         false,
					"topicTags":          []map[string]string{{"name": "Array", "slug": "array"}},
				}
			}
			json.NewEncoder(w).Encode(map[string]any{"data": map[string]any{
				"problemsetQuestionList": map[string]any{"totalNum": 15, "data": questions},
			}})
		default:
			t.Fatalf("unexpected operation %q", req.OperationName)
		}
	}))
	defer srv.Close()
	oldURL := leetcodeGraphQL
	leetcodeGraphQL = srv.URL
	defer func() { leetcodeGraphQL = oldURL }()

	meta := practiceSets[0]
	problems, err := fetchTodayProblems(meta)
	if err != nil {
		t.Fatalf("fetchTodayProblems: %v", err)
	}
	if len(problems) != 10 {
		t.Fatalf("expected 10 problems, got %d", len(problems))
	}
	for _, p := range problems {
		if len(p.Slug) < 5 || p.Slug[:5] != "pool-" {
			t.Fatalf("expected topic pool slug, got %q", p.Slug)
		}
	}
}

func TestSeededShuffleDeterministic(t *testing.T) {
	slugs := []string{"a", "b", "c", "d", "e"}
	a := shuffledSeeds(slugs, "2026-08-13", "Thursday")
	b := shuffledSeeds(slugs, "2026-08-13", "Thursday")
	if len(a) != len(slugs) {
		t.Fatalf("expected %d slugs, got %d", len(slugs), len(a))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("shuffle not deterministic: %v vs %v", a, b)
		}
	}
	c := shuffledSeeds(slugs, "2026-08-14", "Thursday")
	same := true
	for i := range a {
		if a[i] != c[i] {
			same = false
			break
		}
	}
	if same {
		t.Fatal("expected different order for different dates")
	}
}

func TestSaveAndLoadDailyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "daily.json")
	f := dailyFile{
		Date:  "2026-08-13",
		Day:   "Thursday",
		Topic: "Binary Search",
		Problems: []lcProblem{
			{Num: 704, Title: "Binary Search", Slug: "binary-search", Diff: "Easy", Pattern: "binary search"},
		},
	}
	if err := saveDailyFile(path, f); err != nil {
		t.Fatal(err)
	}
	loaded, ok := loadDailyFile(path)
	if !ok || loaded.Date != f.Date || len(loaded.Problems) != 1 {
		t.Fatalf("load failed: ok=%v loaded=%+v", ok, loaded)
	}
}

func TestEnsureTodaySetUsesCache(t *testing.T) {
	dir := t.TempDir()
	path := dailyJSONPath(dir)
	cached := dailyFile{
		Date: todayDate(),
		Day:  "Thursday", Topic: "Binary Search", Reflex: "04_binary_search_reflex",
		Problems: make([]lcProblem, 10),
	}
	for i := range cached.Problems {
		cached.Problems[i] = lcProblem{Num: i + 1, Title: "T", Slug: "binary-search", Diff: "Easy", Pattern: "bs"}
	}
	if err := saveDailyFile(path, cached); err != nil {
		t.Fatal(err)
	}

	// hijack path by writing where ensureTodaySet expects - use temp repo layout
	repo := filepath.Join(dir, "repo")
	if err := os.MkdirAll(filepath.Join(repo, "drills", "leetcode"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := saveDailyFile(filepath.Join(repo, "drills", "leetcode", "daily.json"), cached); err != nil {
		t.Fatal(err)
	}

	set, err := ensureTodaySet(repo, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(set.problems) != 10 {
		t.Fatalf("expected cached 10, got %d", len(set.problems))
	}
}

func TestTodayMeta(t *testing.T) {
	meta := todayMeta()
	if meta.day == "" || meta.topic == "" || len(meta.seedSlugs) < 10 {
		t.Fatalf("incomplete meta: %+v", meta)
	}
}

func TestPracticeSetsCatalog(t *testing.T) {
	if len(practiceSets) != 7 {
		t.Fatalf("expected 7 sets, got %d", len(practiceSets))
	}
	for _, s := range practiceSets {
		if len(s.seedSlugs) < 10 {
			t.Fatalf("%s needs seed slugs", s.day)
		}
		if s.reflex == "" || s.warmup == "" || len(s.suggested) < 1 || len(s.topicTags) < 1 {
			t.Fatalf("%s missing metadata", s.day)
		}
	}
}

func TestParseLeetcode(t *testing.T) {
	opts, ctl, perr := parseLeetcode([]string{"--", "--catalog", "--brief"})
	if perr || ctl != gnuOK || !opts.catalog || !opts.brief || opts.showSet || opts.refresh {
		t.Fatalf("catalog brief: %+v ctl=%v perr=%v", opts, ctl, perr)
	}

	opts, _, _ = parseLeetcode([]string{"--refresh"})
	if !opts.refresh || !opts.showSet {
		t.Fatalf("refresh: %+v", opts)
	}

	opts, _, _ = parseLeetcode([]string{"--run"})
	if !opts.showSet || opts.brief {
		t.Fatalf("--run shows today's set without brief: %+v", opts)
	}

	// --set is a retired spelling of --show.
	opts, _, _ = parseLeetcode([]string{"--catalog", "--set"})
	if !opts.showSet {
		t.Fatalf("--set alias -> --show: %+v", opts)
	}

	// Unambiguous abbreviation.
	opts, _, _ = parseLeetcode([]string{"--ref"})
	if !opts.refresh {
		t.Fatalf("--ref -> --refresh: %+v", opts)
	}

	if _, ctl, _ := parseLeetcode([]string{"--help"}); ctl != gnuHelp {
		t.Fatal("--help -> gnuHelp")
	}
	if _, ctl, _ := parseLeetcode([]string{"-V"}); ctl != gnuVersion {
		t.Fatal("-V -> gnuVersion")
	}
	if _, ctl, _ := parseLeetcode([]string{"--bogus"}); ctl != gnuErr {
		t.Fatal("--bogus -> gnuErr")
	}
}

func TestPrintFunctions(t *testing.T) {
	s := practiceSet{
		day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		warmup: "warmup", suggested: []string{"tip"},
		problems: []lcProblem{{Num: 704, Title: "Binary Search", Slug: "binary-search", Diff: "Easy", Pattern: "bs", Daily: true}},
	}
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printTodaySet(s, false)
	printTodaySet(s, true)
	printCatalog()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 50 {
		t.Fatal("print output too short")
	}
}

func TestPrintHelp(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printHelp()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if !bytes.Contains(buf.Bytes(), []byte("daily.json")) {
		t.Fatal("help missing daily.json note")
	}
}

func TestTopicMatchQuestion(t *testing.T) {
	q := gqlQuestion{TopicTags: []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}{{Slug: "binary-search"}}}
	if !topicMatchQuestion(q, []string{"binary-search"}) {
		t.Fatal("should match binary-search")
	}
	if topicMatchQuestion(q, []string{"graph"}) {
		t.Fatal("should not match graph")
	}
}
