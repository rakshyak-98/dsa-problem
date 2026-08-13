package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTodaySet(t *testing.T) {
	s := todaySet()
	if s.day == "" || s.topic == "" || len(s.problems) != 10 {
		t.Fatalf("todaySet incomplete: %+v", s)
	}
}

func TestPracticeSetsCatalog(t *testing.T) {
	if len(practiceSets) != 7 {
		t.Fatalf("expected 7 sets, got %d", len(practiceSets))
	}
	for _, s := range practiceSets {
		if len(s.problems) != 10 {
			t.Fatalf("%s should have 10 problems, got %d", s.day, len(s.problems))
		}
		if s.reflex == "" || s.warmup == "" || len(s.suggested) < 1 {
			t.Fatalf("%s missing metadata", s.day)
		}
		for _, p := range s.problems {
			if p.num <= 0 || p.title == "" || p.diff == "" || p.pattern == "" {
				t.Fatalf("%s problem incomplete: %+v", s.day, p)
			}
			if slugFor(p.num) == "" {
				t.Fatalf("missing slug for #%d", p.num)
			}
		}
	}
}

func TestParseLeetcodeArgs(t *testing.T) {
	help, catalog, brief, showSet, parseErr := parseLeetcodeArgs([]string{"--", "--catalog", "--brief"})
	if parseErr || !catalog || !brief || showSet || help {
		t.Fatalf("catalog brief: help=%v catalog=%v brief=%v showSet=%v err=%v", help, catalog, brief, showSet, parseErr)
	}
	_, _, _, _, parseErr = parseLeetcodeArgs([]string{"--run"})
	if !parseErr {
		t.Fatal("run should error")
	}
	_, _, _, showSet, parseErr = parseLeetcodeArgs(nil)
	if parseErr || !showSet {
		t.Fatal("default should show set")
	}
	_, _, _, _, parseErr = parseLeetcodeArgs([]string{"--nope"})
	if !parseErr {
		t.Fatal("unknown flag should error")
	}
}

func TestPrintFunctions(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printTodaySet(practiceSets[0], false)
	printTodaySet(practiceSets[0], true)
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
	if !bytes.Contains(buf.Bytes(), []byte("leetcode.com")) {
		t.Fatal("help missing leetcode note")
	}
}

func TestSlugForKnown(t *testing.T) {
	if slugFor(704) != "binary-search" {
		t.Fatal("expected binary-search slug")
	}
}
