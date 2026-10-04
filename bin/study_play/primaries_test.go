package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestEveryCatalogFunctionHasAPrimaryProblem(t *testing.T) {
	for _, entry := range essentialCatalog {
		for _, fn := range entry.fns {
			p, ok := primaries[fn]
			if !ok {
				t.Errorf("%s (%s) has no primary problem", fn, entry.group)
				continue
			}
			if p.title == "" || p.slug == "" || p.nextUp == "" {
				t.Errorf("%s: incomplete primary", fn)
			}
			switch p.diff {
			case "Easy", "Medium", "Hard":
			default:
				t.Errorf("%s: bad difficulty %q", fn, p.diff)
			}
		}
	}
}

func TestOwningFunction(t *testing.T) {
	cases := []struct {
		in   string
		want string
		ok   bool
	}{
		{"rotateRight k=1", "rotateRight", true},
		{"rotateRight", "rotateRight", true},
		{"maxSumSubarrayK window slide", "maxSumSubarrayK", true},
		{"removeDuplicates empty", "removeDuplicates", true},
		{"TestRotateRight (0.00s)", "", false},
		{"TestAll/maxSumSubarrayK (0.00s)", "", false},
		{"rotateRightward", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		got, ok := owningFunction(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("owningFunction(%q) = (%q, %v), want (%q, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func writeTestLog(t *testing.T, fns map[string]fnRecord) string {
	t.Helper()
	root := t.TempDir()
	if err := saveLog(root, drillLog{Functions: fns}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".drill_log.json")); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestPrintersDoNotPanic(t *testing.T) {
	mustPrint(t, "printProblemSet", printProblemSet)
	mustPrint(t, "printProblemMap", func() { printProblemMap("03_two_pointers_reflex") })
}

func mustPrint(t *testing.T, name string, fn func()) {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Errorf("%s printed nothing", name)
	}
}
