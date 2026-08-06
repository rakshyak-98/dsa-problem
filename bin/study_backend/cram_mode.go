package main

import (
	"fmt"
	"time"
)

func printCramPlan() {
	interview := time.Date(2026, 8, 7, 16, 0, 0, 0, time.UTC)
	now := time.Now()
	hoursLeft := interview.Sub(now).Hours()

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║     BACKEND INTERVIEW CRAM — Rakshyak Satpathy resume    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Printf("\nInterview: %s (4:00 PM – 5:00 PM)\n", interview.Format("Monday, January 2, 2006 15:04 MST"))
	if hoursLeft > 0 {
		fmt.Printf("Time until interview: ~%.0f hours\n", hoursLeft)
	} else {
		fmt.Println("Interview window is today — run --drill core for warm-up only.")
	}
	fmt.Println()

	if now.Day() == 6 && now.Month() == time.August {
		printAug6Plan()
	} else if now.Day() == 7 && now.Hour() < 16 {
		printAug7Plan(now.Hour())
	} else {
		printGenericCram()
	}
}

func printAug6Plan() {
	fmt.Println("TODAY (Aug 6) — full resume coverage")
	fmt.Println("------------------------------------")
	fmt.Println("Session 1 (2h)  09:00–11:00")
	fmt.Println("  • Core 5 explain + write (--run)")
	fmt.Println("  • Block 01_rest_api_jwt + Block 02_databases_sql")
	fmt.Println("  • Scenario: MST page publish API out loud")
	fmt.Println()
	fmt.Println("Session 2 (2h)  14:00–16:00")
	fmt.Println("  • Block 03_distributed_resilience + Block 04_realtime_webrtc")
	fmt.Println("  • STAR: iBind latency 40% + WebRTC NAT story")
	fmt.Println()
	fmt.Println("Session 3 (2h)  19:00–21:00")
	fmt.Println("  • Block 05_workflows_messaging + Block 06_devops_aws")
	fmt.Println("  • Block 07_compliance_security (ZATCA)")
	fmt.Println()
	fmt.Println("Session 4 (1.5h) 21:30–23:00")
	fmt.Println("  • Block 08_go_systems (BitTorrent + HTTP server)")
	fmt.Println("  • Read STAR_STORIES.md — rehearse each 90 sec")
	fmt.Println("  • Sleep by midnight")
}

func printAug7Plan(hour int) {
	fmt.Println("TOMORROW (Aug 7) — interview day")
	fmt.Println("------------------------------")
	if hour < 10 {
		fmt.Println("Morning (45 min)")
		fmt.Println("  • go run . -- --drill core --run  (Core 5 only)")
		fmt.Println("  • Skim DRILL_CONCEPTS.md headers")
	}
	if hour < 13 {
		fmt.Println("Late morning (45 min)")
		fmt.Println("  • Block 03 + 04 explain drills")
		fmt.Println("  • Rehearse WebRTC NAT debug scenario")
	}
	if hour < 15 {
		fmt.Println("Pre-lunch (30 min)")
		fmt.Println("  • Block 05 + 07 — Airflow idempotency + ZATCA signing")
	}
	fmt.Println("15:30–15:55  Warm-up")
	fmt.Println("  • Core 5 explain OUT LOUD (no peeking)")
	fmt.Println("  • One STAR story: reporting 35% or invoice pipeline")
	fmt.Println("  • Water, quiet room, test mic if remote")
	fmt.Println()
	fmt.Println("16:00–17:00  INTERVIEW")
	fmt.Println("  • Listen fully before answering")
	fmt.Println("  • Structure: context → your action → metric/result")
	fmt.Println("  • Draw boxes for system design; label data flows")
}

func printGenericCram() {
	fmt.Println("Use INTERVIEW_CRAM_PLAN.md for the full schedule.")
	fmt.Println("Run all blocks: go run . -- --catalog")
	fmt.Println("Test yourself:  go run . -- --run")
}
