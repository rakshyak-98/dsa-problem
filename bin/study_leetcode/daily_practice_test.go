package main

import (
	"bytes"
	"encoding/json"
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
	meta.seedSlugs = []string{
		"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "s11",
	}
	problems, err := fetchTodayProblems(meta)
	if err != nil {
		t.Fatalf("fetchTodayProblems: %v", err)
	}
	if len(problems) != 10 {
		t.Fatalf("expected 10 problems, got %d", len(problems))
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

func TestParseLeetcodeArgs(t *testing.T) {
	help, catalog, brief, showSet, refresh, parseErr := parseLeetcodeArgs([]string{"--", "--catalog", "--brief"})
	if parseErr || !catalog || !brief || showSet || help || refresh {
		t.Fatalf("catalog brief: help=%v catalog=%v brief=%v showSet=%v refresh=%v err=%v", help, catalog, brief, showSet, refresh, parseErr)
	}
	_, _, _, _, refresh, parseErr = parseLeetcodeArgs([]string{"--refresh"})
	if parseErr || !refresh {
		t.Fatal("refresh flag")
	}
	_, _, _, showSet, _, parseErr = parseLeetcodeArgs([]string{"--run"})
	if parseErr || !showSet {
		t.Fatal("--run should show today's set")
	}
	_, _, brief, _, _, parseErr = parseLeetcodeArgs([]string{"--run"})
	if parseErr || brief {
		t.Fatal("--run should not use brief mode")
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
