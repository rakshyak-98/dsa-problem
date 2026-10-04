package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestProblemMapCoverage(t *testing.T) {
	for _, d := range drills {
		links, ok := problemMap[d.file]
		if !ok {
			t.Fatalf("missing problem map for %s", d.file)
		}
		if len(links) == 0 {
			t.Fatalf("empty links for %s", d.file)
		}
		for _, l := range links {
			if l.function == "" || l.problem == "" || l.ask == "" {
				t.Fatalf("incomplete link in %s", d.file)
			}
		}
	}
}

func TestPrintProblemMapNoPanic(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printProblemMap("01_arrays_reflex")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 20 {
		t.Fatal("printProblemMap empty")
	}
}
