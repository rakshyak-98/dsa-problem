// Mock interview scenarios — practice OUT LOUD (no auto-check)
//
// RUN: go run .  (prints prompts)
package main

import "fmt"

var scenarios = []struct {
	title  string
	prompt string
	hints  []string
}{
	{
		title: "MST — Page publish API",
		prompt: "Design REST endpoints for creating, versioning, and publishing pages with role-based permissions.",
		hints: []string{
			"Resources: /pages, /pages/{id}/versions, /pages/{id}/publish",
			"Draft vs published snapshot; audit who published",
			"403 vs 404 for unauthorized vs missing",
		},
	},
	{
		title: "MST — Slow reporting query",
		prompt: "Analytics query on 10M rows takes 30s. How did you get 35% improvement?",
		hints: []string{
			"EXPLAIN → full scan on filter column",
			"Composite index on (tenant_id, created_at)",
			"Rewrite subquery; consider summary table",
		},
	},
	{
		title: "iBind — NAT video failure",
		prompt: "Users on corporate VPN cannot connect video. Debug approach?",
		hints: []string{
			"Signaling OK? ICE candidates include relay?",
			"Deploy TURN; test UDP blocked",
			"Connection state timeline in client logs",
		},
	},
	{
		title: "iBind — Downstream outage",
		prompt: "Payment service flaky. How do you protect your API?",
		hints: []string{
			"Timeout every outbound call",
			"Retry idempotent reads with backoff",
			"Circuit breaker; degrade gracefully",
		},
	},
	{
		title: "Opscale — Duplicate invoices",
		prompt: "Airflow task retry created duplicate submissions. Fix?",
		hints: []string{
			"Idempotent task keyed by invoice ID",
			"Check submission log before POST",
			"DLQ for permanent failures",
		},
	},
	{
		title: "Opscale — ZATCA rejection",
		prompt: "Authority rejected invoice. Trace root cause.",
		hints: []string{
			"XSD validation errors in logs",
			"Signature certificate expiry",
			"Compare payload hash to audit store",
		},
	},
	{
		title: "Go — HTTP server design",
		prompt: "Walk through your HTTP server middleware chain.",
		hints: []string{
			"Logging → recovery → auth → handler",
			"Context carries request ID + user",
			"Graceful shutdown on SIGTERM",
		},
	},
	{
		title: "System design — Rate limiter",
		prompt: "Design a rate limiter for 10k RPS API.",
		hints: []string{
			"Token bucket or sliding window per client",
			"Redis for distributed counters",
			"429 + Retry-After header",
		},
	},
}

func main() {
	fmt.Println("MOCK SCENARIOS — 3 minutes each, out loud")
	fmt.Println("============================================")
	for i, s := range scenarios {
		fmt.Printf("\n%d. %s\n", i+1, s.title)
		fmt.Printf("   Q: %s\n", s.prompt)
		fmt.Println("   (After answering, reveal hints:)")
		for _, h := range s.hints {
			fmt.Printf("     • %s\n", h)
		}
	}
	fmt.Println("\nFull STAR scripts: study_backend/docs/STAR_STORIES.md")
}
