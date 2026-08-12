package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type coreItem struct {
	topic  string
	prompt string
	sec    int
}

type block struct {
	day      string
	file     string
	topic    string
	resume   string
	focus    []string
	warmup   string
	scenario string
}

type revisionDay struct {
	day           string
	file          string
	label         string
	topics        []string
	revisitBlocks []string
	activity      string
}

type topicGroup struct {
	group  string
	topics []string
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

// weekdayBlocks — Mon→Sun resume blocks (block 08 is bonus / Sunday revision add-on).
var weekdayBlocks = []block{
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
}

var bonusBlock = block{
	day: "Bonus", file: "08_go_systems",
	topic: "Go systems programming",
	resume: "BitTorrent client; HTTP server from scratch",
	warmup: "Goroutine per connection vs worker pool — tradeoffs.",
	focus: []string{
		"Goroutines, channels, context cancellation",
		"TCP lifecycle, HTTP/1.1 keep-alive",
		"bencode, peer wire protocol, piece scheduling",
	},
	scenario: "Explain your HTTP server middleware chain design.",
}

// blocks is the full catalog (weekday + bonus) for cram/catalog views.
var blocks = append(weekdayBlocks, bonusBlock)

// revisionCycle — weekly cross-topic revision (re-run prior blocks from memory).
var revisionCycle = []revisionDay{
	{
		day: "Monday", file: "01_api_auth_recap",
		label: "API & auth recap",
		topics: []string{
			"REST resource naming, status codes, error contract",
			"JWT flow: access/refresh, exp/iss/aud validation",
			"Idempotency-Key for safe POST retries",
		},
		revisitBlocks: []string{},
		activity:      "Say trigger rows (API design, Auth) out loud; run today's block 01 from memory.",
	},
	{
		day: "Tuesday", file: "02_data_resilience_recap",
		label: "Data layer + resilience recap",
		topics: []string{
			"EXPLAIN → index → rewrite → pool/cache",
			"Composite index column order",
			"Circuit breaker states + exponential backoff",
		},
		revisitBlocks: []string{"01_rest_api_jwt"},
		activity:      "Re-run block 01 explain without peeking; then today's block 02.",
	},
	{
		day: "Wednesday", file: "03_distributed_realtime",
		label: "Distributed + real-time recap",
		topics: []string{
			"CAP trade-off during partition",
			"Saga vs 2PC",
			"WebSocket upgrade vs REST polling",
		},
		revisitBlocks: []string{"02_databases_sql"},
		activity:      "Re-run block 02; draw outage isolation diagram for block 03.",
	},
	{
		day: "Thursday", file: "04_realtime_messaging",
		label: "Real-time + messaging recap",
		topics: []string{
			"SDP/ICE/STUN/TURN NAT checklist",
			"Airflow idempotent tasks + DLQ",
			"RabbitMQ work queue vs pub/sub",
		},
		revisitBlocks: []string{"03_distributed_resilience"},
		activity:      "Re-run block 03; rehearse WebRTC NAT debug scenario out loud.",
	},
	{
		day: "Friday", file: "05_devops_orchestration",
		label: "DevOps + orchestration recap",
		topics: []string{
			"Jenkins pipeline stages",
			"Docker liveness vs readiness",
			"Lambda vs EC2 vs CDN quick map",
		},
		revisitBlocks: []string{"04_realtime_webrtc", "05_workflows_messaging"},
		activity:      "Re-run blocks 04–05; walk Jenkins setup from resume.",
	},
	{
		day: "Saturday", file: "06_full_week_sweep",
		label: "Full-week trigger sweep",
		topics: []string{
			"All 9 trigger-table rows (DRILL_CONCEPTS.md)",
			"Cross-block scenario: flake → retry → breaker → DLQ",
			"STAR story pick: latency 40% or ZATCA pipeline",
		},
		revisitBlocks: []string{
			"01_rest_api_jwt", "02_databases_sql", "03_distributed_resilience",
			"04_realtime_webrtc", "05_workflows_messaging", "06_devops_aws",
		},
		activity:      "Run all 6 weekday blocks (--run reflex)",
	},
	{
		day: "Sunday", file: "07_go_compliance_mix",
		label: "Go systems + compliance mix",
		topics: []string{
			"ZATCA: XSD → sign → submit → audit",
			"Go: goroutine pools, context cancel, graceful shutdown",
			"HTTP keep-alive + BitTorrent piece scheduling",
		},
		revisitBlocks: []string{"07_compliance_security", "08_go_systems"},
		activity:      "Re-run blocks 07 + 08; rest or light review.",
	},
}

// backendTopics — collected topic index for catalog and docs.
var backendTopics = []topicGroup{
	{group: "API & auth", topics: []string{
		"REST design", "JWT", "versioning", "pagination", "idempotency", "structured errors",
	}},
	{group: "Data layer", topics: []string{
		"MySQL indexing", "EXPLAIN", "connection pooling", "schema design", "read replicas", "N+1",
	}},
	{group: "Distributed systems", topics: []string{
		"Retry + backoff", "circuit breaker", "CAP", "saga vs 2PC", "idempotency keys", "bulkhead",
	}},
	{group: "Real-time", topics: []string{
		"WebSocket", "WebRTC signaling", "SDP/ICE", "STUN/TURN", "NAT debug",
	}},
	{group: "Workflows & messaging", topics: []string{
		"Airflow DAGs", "idempotent tasks", "RabbitMQ", "DLQ", "pub/sub",
	}},
	{group: "DevOps & cloud", topics: []string{
		"Docker multi-stage", "Jenkins pipelines", "EC2", "Lambda", "CDN", "health checks",
	}},
	{group: "Compliance & security", topics: []string{
		"ZATCA XML/XSD", "cryptographic signing", "audit trail", "replay protection",
	}},
	{group: "Go systems", topics: []string{
		"goroutines/channels", "context", "HTTP server", "BitTorrent", "worker pools",
	}},
}

func weekdayIndex(wd time.Weekday) int {
	dayIndex := int(wd)
	idx := dayIndex - 1
	if dayIndex == 0 {
		idx = 6
	}
	return idx
}

func todayBlock() block {
	// Interview cram: Aug 6 eve → block 08; Aug 7 morning → resilience
	now := time.Now()
	if now.Year() == 2026 && now.Month() == time.August && now.Day() == 6 {
		return bonusBlock
	}
	if now.Year() == 2026 && now.Month() == time.August && now.Day() == 7 && now.Hour() < 16 {
		return weekdayBlocks[2]
	}
	return weekdayBlocks[weekdayIndex(now.Weekday())]
}

func todayRevision() revisionDay {
	return revisionCycle[weekdayIndex(time.Now().Weekday())]
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

func printRevision(r revisionDay) {
	fmt.Println("════════════════════════════════════════")
	fmt.Printf(" WEEKLY REVISION — %s (%s)\n", r.day, r.file)
	fmt.Println("════════════════════════════════════════")
	fmt.Printf("  Label:    %s\n", r.label)
	fmt.Println("  Topics:")
	for _, t := range r.topics {
		fmt.Printf("    • %s\n", t)
	}
	if len(r.revisitBlocks) > 0 {
		fmt.Println("  Revisit blocks (no peeking):")
		for _, b := range r.revisitBlocks {
			fmt.Printf("    → drills/backend/explain/blocks/%s/\n", b)
		}
	}
	fmt.Printf("\n  Activity: %s\n", r.activity)
	fmt.Printf("\n  Open:  drills/backend/explain/revision/%s/\n", r.file)
	fmt.Println()
}

func printCatalog() {
	fmt.Println("BACKEND INTERVIEW BLOCK CATALOG (from resume)")
	fmt.Println("----------------------------------------------")
	for _, b := range weekdayBlocks {
		fmt.Printf("%-9s  %-22s  %s\n", b.day, b.file, b.topic)
	}
	fmt.Printf("%-9s  %-22s  %s\n", bonusBlock.day, bonusBlock.file, bonusBlock.topic)
	fmt.Println()
	fmt.Println("WEEKLY REVISION CYCLE")
	fmt.Println("---------------------")
	for _, r := range revisionCycle {
		fmt.Printf("%-9s  %-24s  %s\n", r.day, r.file, r.label)
	}
	fmt.Println()
	fmt.Println("TOPIC GROUPS")
	fmt.Println("------------")
	for _, g := range backendTopics {
		fmt.Printf("%-28s  %s\n", g.group, strings.Join(g.topics, ", "))
	}
	fmt.Println()
	fmt.Println("Core explain: drills/backend/explain/core5/")
	fmt.Println("Core write:   drills/backend/write/core5/")
	fmt.Println("Revision:     drills/backend/explain/revision/")
	fmt.Println("Weekly plan:  doc/backend/WEEKLY_REVISION.md")
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
