package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestVisualizerLinks(t *testing.T) {
	for _, d := range drills {
		if _, ok := visualizerLinks[d.file]; !ok {
			t.Fatalf("missing visualizer link for %s", d.file)
		}
	}
	if _, ok := visualizerLinks["core5"]; !ok {
		t.Fatal("missing core5 visualizer")
	}
}

func TestPrintVisualizerLink(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printVisualizerLink("02_hashing_reflex")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 10 {
		t.Fatal("printVisualizerLink empty")
	}
}
