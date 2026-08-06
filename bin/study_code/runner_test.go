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
	code := runStudyCode(false, false, true, false, "")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || len(buf.String()) < 30 {
		t.Fatal("catalog mode")
	}
}

func TestRunStudyCodeDrill(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(true, false, false, false, "")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || !strings.Contains(buf.String(), "core 3") {
		t.Fatal("drill mode")
	}
}

func TestRunStudyCodeDefault(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	code := runStudyCode(false, false, false, false, "")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if code != 0 || len(buf.String()) < 20 {
		t.Fatal("default mode")
	}
}
