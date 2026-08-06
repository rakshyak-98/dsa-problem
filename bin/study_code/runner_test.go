package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRunStudyCodeSpecialty(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(false, false, false, true)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()
	if code != 0 || !strings.Contains(out, "SPECIALTY") || strings.Contains(out, "CORE READ") {
		t.Fatalf("specialty mode: %q", out)
	}
}

func TestRunStudyCodeCatalog(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(false, false, true, false)
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
	code := runStudyCode(true, false, false, false)
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
	code := runStudyCode(false, false, false, false)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || len(buf.String()) < 50 {
		t.Fatal("default mode")
	}
}
