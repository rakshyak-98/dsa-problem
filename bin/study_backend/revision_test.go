package main

import (
	"testing"
	"time"
)

func TestWeekdayBlockMapping(t *testing.T) {
	cases := []struct {
		wd   time.Weekday
		want string
	}{
		{time.Monday, "01_rest_api_jwt"},
		{time.Tuesday, "02_databases_sql"},
		{time.Wednesday, "03_distributed_resilience"},
		{time.Thursday, "04_realtime_webrtc"},
		{time.Friday, "05_workflows_messaging"},
		{time.Saturday, "06_devops_aws"},
		{time.Sunday, "07_compliance_security"},
	}
	for _, tc := range cases {
		got := weekdayBlocks[weekdayIndex(tc.wd)].file
		if got != tc.want {
			t.Errorf("weekday %s: got block %q, want %q", tc.wd, got, tc.want)
		}
	}
}

func TestRevisionCycleMatchesWeekdays(t *testing.T) {
	if len(revisionCycle) != len(weekdayBlocks) {
		t.Fatalf("revision cycle len %d != weekday blocks %d", len(revisionCycle), len(weekdayBlocks))
	}
	days := []time.Weekday{
		time.Monday, time.Tuesday, time.Wednesday, time.Thursday,
		time.Friday, time.Saturday, time.Sunday,
	}
	for _, wd := range days {
		r := revisionCycle[weekdayIndex(wd)]
		if r.file == "" || r.label == "" {
			t.Fatalf("empty revision entry for %s: %+v", wd, r)
		}
	}
}

func TestBackendTopicsNonEmpty(t *testing.T) {
	if len(backendTopics) < 8 {
		t.Fatalf("expected at least 8 topic groups, got %d", len(backendTopics))
	}
	for _, g := range backendTopics {
		if g.group == "" || len(g.topics) == 0 {
			t.Fatalf("invalid topic group: %+v", g)
		}
	}
}
