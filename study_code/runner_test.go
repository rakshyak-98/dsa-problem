package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRunStudyCodeCatalog(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(false, false, true)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || len(buf.String()) < 30 {
		t.Fatal("catalog mode")
	}
}

func TestRunStudyCodeMicro(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(true, false, false)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || !strings.Contains(buf.String(), "micro") {
		t.Fatal("micro mode")
	}
}

func TestRunStudyCodeDefault(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(false, false, false)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || len(buf.String()) < 50 {
		t.Fatal("default mode")
	}
}
