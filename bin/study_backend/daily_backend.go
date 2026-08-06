package main

import (
	"fmt"
	"os"
	"time"
)

type coreItem struct {
	topic  string
	prompt string
	sec    int
}

type block struct {
	day     string
	file    string
	topic   string
	resume  string
	focus   []string
	warmup  string
	scenario string
}

var coreExplain = []coreItem{
	{"JWT auth flow", "Access token vs refresh token; where validated; what claims matter", 120},
	{"Circuit breaker", "Closed/Open/Half-open states; when to trip; what fails fast", 90},
	{"SQL slow query", "Three levers: indexes, query rewrite, pooling/caching", 90},
	{"REST vs WebSocket", "When polling/REST wins vs persistent duplex channel", 60},
	{"Retry + backoff", "Why backoff; jitter; idempotency requirement", 90},
}

var coreWrite = []coreItem{
	{"exponentialBackoff", "Compute delay for attempt n with cap", 90},
	{"circuitState", "Next state given failures and threshold", 90},
	{"httpStatusClass", "Map status code to class name", 60},
	{"isIdempotentMethod", "PUT/GET/DELETE vs POST", 60},
	{"parseBearerToken", "Extract token from Authorization header", 90},
}

var blocks = []block{
	{
		day: "Monday", file: "01_rest_api_jwt",
		topic: "REST API design & JWT",
		resume: "MST page-builder APIs; iBind REST + JWT",
		warmup: "Before status codes: what is the resource contract (input, output, errors)?",
		focus: []string{
			"Resource naming, versioning, pagination",
			"Structured errors {code, message, details}",
			"JWT: header.payload.signature, validate exp/iss/aud",
		},
		scenario: "Design publish API for a page with versioning and role checks.",
	},
	{
		day: "Tuesday", file: "02_databases_sql",
		topic: "Databases & query optimization",
		resume: "MySQL schema (MST); 35% reporting gain; connection pooling (iBind)",
		warmup: "EXPLAIN the query before adding an index.",
		focus: []string{
			"Normalization vs denormalization for analytics",
			"Index types: B-tree, composite, covering",
			"Connection pool sizing, N+1, read replicas",
		},
		scenario: "Reporting query is slow on 10M rows — walk through your MST fix.",
	},
	{
		day: "Wednesday", file: "03_distributed_resilience",
		topic: "Distributed systems & resilience",
		resume: "Circuit breaker + exponential backoff (iBind); incident response (Opscale)",
		warmup: "What failure is transient vs requires human intervention?",
		focus: []string{
			"Timeouts, retries, circuit breaker, bulkhead",
			"Idempotency keys for safe retries",
			"CAP, eventual consistency, saga vs 2PC",
		},
		scenario: "Downstream invoice service flakes — how do you isolate blast radius?",
	},
	{
		day: "Thursday", file: "04_realtime_webrtc",
		topic: "WebSocket, WebRTC, STUN/TURN",
		resume: "P2P video conferencing platform (iBind)",
		warmup: "Signaling is not the media path — separate concerns.",
		focus: []string{
			"WebSocket upgrade, rooms, heartbeats",
			"SDP offer/answer, ICE candidates",
			"STUN vs TURN; NAT hairpin failures",
		},
		scenario: "Users on corporate NAT cannot connect — debug checklist.",
	},
	{
		day: "Friday", file: "05_workflows_messaging",
		topic: "Airflow, RabbitMQ, orchestration",
		resume: "ZATCA invoice DAGs (Opscale); RabbitMQ in skills",
		warmup: "Every task should be safe to rerun.",
		focus: []string{
			"DAG structure, sensors, SLAs",
			"Idempotent tasks, retry policies",
			"Queue patterns: work queue, pub/sub, dead letter",
		},
		scenario: "Invoice pipeline task fails mid-batch — recovery without duplicates.",
	},
	{
		day: "Saturday", file: "06_devops_aws",
		topic: "Docker, Jenkins, AWS, CI/CD",
		resume: "Jenkins pipelines (iBind); AWS EC2, Lambda, CDN",
		warmup: "CI verifies; CD deploys — list your pipeline stages.",
		focus: []string{
			"Multi-stage Dockerfile, health checks",
			"Jenkins: build → test → scan → deploy",
			"Lambda vs EC2; CDN for static assets",
		},
		scenario: "Walk through Jenkins setup you built at iBind.",
	},
	{
		day: "Sunday", file: "07_compliance_security",
		topic: "XML, signing, compliance",
		resume: "Saudi ZATCA e-invoicing (Opscale)",
		warmup: "Compliance flows need audit trail + non-repudiation.",
		focus: []string{
			"XML schema validation (XSD)",
			"Cryptographic signing workflow",
			"Secure submission, replay protection",
		},
		scenario: "Invoice rejected by ZATCA — how do you trace root cause?",
	},
	{
		day: "Interview eve", file: "08_go_systems",
		topic: "Go systems programming",
		resume: "BitTorrent client; HTTP server from scratch",
		warmup: "Goroutine per connection vs worker pool — tradeoffs.",
		focus: []string{
			"Goroutines, channels, context cancellation",
			"TCP lifecycle, HTTP/1.1 keep-alive",
			"bencode, peer wire protocol, piece scheduling",
		},
		scenario: "Explain your HTTP server middleware chain design.",
	},
}

func todayBlock() block {
	// Interview cram: Aug 6 eve → block 08; Aug 7 morning → rotate through high-priority
	now := time.Now()
	if now.Year() == 2026 && now.Month() == time.August && now.Day() == 6 {
		return blocks[7] // Go systems + projects
	}
	if now.Year() == 2026 && now.Month() == time.August && now.Day() == 7 && now.Hour() < 16 {
		return blocks[2] // resilience — common backend interview topic
	}
	wd := int(now.Weekday())
	idx := wd % len(blocks)
	return blocks[idx]
}

func printCoreExplain() {
	fmt.Println("════════════════════════════════════════")
	fmt.Println(" CORE 5 EXPLAIN  (every session)")
	fmt.Println("════════════════════════════════════════")
	for i, c := range coreExplain {
		fmt.Printf("  %d. [%ds] %s\n     → %s\n", i+1, c.sec, c.topic, c.prompt)
	}
	fmt.Println()
	fmt.Println("  File: drills/backend/explain/core5/")
	fmt.Println("  Say answers out loud, then fill TODO: EXPLAIN vars.")
	fmt.Println()
}

func printCoreWrite() {
	fmt.Println("════════════════════════════════════════")
	fmt.Println(" CORE 5 WRITE  (Go reflex)")
	fmt.Println("════════════════════════════════════════")
	for i, c := range coreWrite {
		fmt.Printf("  %d. [%ds] %s\n", i+1, c.sec, c.topic)
	}
	fmt.Println()
	fmt.Println("  File: drills/backend/write/core5/")
	fmt.Println()
}

func printBlock(b block) {
	fmt.Println("════════════════════════════════════════")
	fmt.Printf(" RESUME BLOCK — %s (%s)\n", b.day, b.file)
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("  Topic:   %s\n", b.topic)
	fmt.Printf("  Resume:  %s\n", b.resume)
	fmt.Printf("  Warm-up: %s\n", b.warmup)
	fmt.Println("  Focus:")
	for _, f := range b.focus {
		fmt.Printf("    • %s\n", f)
	}
	fmt.Printf("\n  Scenario (practice 3 min out loud): %s\n", b.scenario)
	fmt.Printf("\n  Open:  drills/backend/explain/blocks/%s/\n", b.file)
	fmt.Printf("  STAR:  doc/backend/STAR_STORIES.md\n")
	fmt.Println()
}

func printCatalog() {
	fmt.Println("BACKEND INTERVIEW BLOCK CATALOG (from resume)")
	fmt.Println("----------------------------------------------")
	for _, b := range blocks {
		fmt.Printf("%-12s  %-22s  %s\n", b.day, b.file, b.topic)
	}
	fmt.Println()
	fmt.Println("Core explain: drills/backend/explain/core5/")
	fmt.Println("Core write:   drills/backend/write/core5/")
	fmt.Println("Cram plan:    doc/backend/INTERVIEW_CRAM_PLAN.md")
}

func main() {
	drillKind, catalog, cram, setup, runMode, drillErr := parseBackendArgs(os.Args[1:])
	if drillErr {
		os.Exit(1)
	}
	if code := runBackend(drillKind, catalog, cram, setup, runMode); code != 0 {
		os.Exit(code)
	}
}
