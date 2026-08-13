package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHtmlToText(t *testing.T) {
	raw := `<p>Given <code>nums</code> and <code>target</code>.</p>
<p><strong class="example">Example 1:</strong></p>
<pre>
<strong>Input:</strong> nums = [2,7], target = 9
<strong>Output:</strong> [0,1]
</pre>
<ul><li><code>2 &lt;= nums.length</code></li></ul>`

	got := htmlToText(raw)
	for _, want := range []string{
		"`nums`", "`target`", "**Example 1:**", "```", "Input:", "- `2 <= nums.length`",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("htmlToText missing %q in:\n%s", want, got)
		}
	}
}

func TestRenderDailyMarkdownMock(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req gqlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if req.OperationName == "getQuestionContent" {
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"question": map[string]any{
						"content": "<p>Search in sorted array.</p>",
					},
				},
			})
			return
		}
		http.Error(w, "unexpected op", 400)
	}))
	defer srv.Close()
	oldURL := leetcodeGraphQL
	leetcodeGraphQL = srv.URL
	defer func() { leetcodeGraphQL = oldURL }()

	set := practiceSet{
		day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		problems: []lcProblem{
			{Num: 704, Title: "Binary Search", Slug: "binary-search", Diff: "Easy", Pattern: "binary search"},
		},
	}
	md, err := renderDailyMarkdown(set, "2026-08-13T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"# Daily LeetCode —", "Binary Search", "704", "leetcode.com/problems/binary-search",
		"Search in sorted array",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("render missing %q in:\n%s", want, md)
		}
	}
}

func TestWriteDailyMarkdownSkipsWhenCurrent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "drills", "leetcode", "daily.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	header := "# Daily LeetCode — " + todayDate() + "\nbody"
	if err := os.WriteFile(path, []byte(header), 0o644); err != nil {
		t.Fatal(err)
	}
	set := practiceSet{problems: []lcProblem{{Slug: "would-fail-if-fetched"}}}
	if err := writeDailyMarkdown(dir, set, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != header {
		t.Fatalf("expected cache skip, got %q", data)
	}
}

func TestDailyMarkdownPath(t *testing.T) {
	if p := dailyMarkdownPath("/tmp/repo"); p != "/tmp/repo/drills/leetcode/daily.md" {
		t.Fatalf("unexpected path: %s", p)
	}
}

func TestRenderDailyGoMock(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req gqlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if req.OperationName == "getQuestionContent" {
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"question": map[string]any{
						"content": "<p>Search in sorted array.</p>",
					},
				},
			})
			return
		}
		http.Error(w, "unexpected op", 400)
	}))
	defer srv.Close()
	oldURL := leetcodeGraphQL
	leetcodeGraphQL = srv.URL
	defer func() { leetcodeGraphQL = oldURL }()

	set := practiceSet{
		day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		problems: []lcProblem{
			{Num: 704, Title: "Binary Search", Slug: "binary-search", Diff: "Easy", Pattern: "binary search"},
		},
	}
	src, err := renderDailyGo(set, "2026-08-13T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"package main", "DAILY LEETCODE —", "Binary Search", "Search in sorted array",
		"go run ./drills/leetcode/daily.go",
	} {
		if !strings.Contains(src, want) {
			t.Fatalf("renderDailyGo missing %q", want)
		}
	}
	if err := os.WriteFile(filepath.Join(t.TempDir(), "daily.go"), []byte(src), 0o644); err == nil {
		// compile check via go vet on generated source would need module; syntax check only
		if !strings.HasSuffix(strings.TrimSpace(src), "}") {
			t.Fatal("generated go missing closing brace")
		}
	}
}

func TestWriteDailyGoSkipsWhenCurrent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "drills", "leetcode", "daily.go")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	header := "// DAILY LEETCODE — " + todayDate() + "\npackage main\n"
	if err := os.WriteFile(path, []byte(header), 0o644); err != nil {
		t.Fatal(err)
	}
	set := practiceSet{problems: []lcProblem{{Slug: "would-fail-if-fetched"}}}
	if err := writeDailyGo(dir, set, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != header {
		t.Fatalf("expected cache skip, got %q", data)
	}
}

func TestDailyGoPath(t *testing.T) {
	if p := dailyGoPath("/tmp/repo"); p != "/tmp/repo/drills/leetcode/daily.go" {
		t.Fatalf("unexpected path: %s", p)
	}
}
