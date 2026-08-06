package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDailyAsksComplete(t *testing.T) {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, day := range days {
		ask, ok := dailyAsks[day]
		if !ok {
			t.Fatalf("missing ask for %s", day)
		}
		if ask.statement == "" || len(ask.hints) < 4 {
			t.Fatalf("incomplete ask for %s", day)
		}
	}
}

func TestPrintAskWarmup(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printAskWarmup("Monday")
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	if len(buf.String()) < 30 {
		t.Fatal("printAskWarmup empty")
	}
}
