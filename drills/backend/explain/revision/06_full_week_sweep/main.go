// REVISION 06 — Full-week trigger sweep (Saturday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// Trigger 1 — Slow SQL: lead with "explain"
var triggerSlowSQL = ""

// Trigger 2 — API design: lead with "resource" or "rest"
var triggerAPI = ""

// Trigger 3 — Auth: JWT claim for expiry: "exp"
var triggerAuth = ""

// Trigger 4 — Outage: after timeout use: "retry" or "backoff"
var triggerOutage = ""

// Trigger 5 — Real-time: signaling uses: "websocket"
var triggerRealtime = ""

// Trigger 6 — Batch job: tasks must be: "idempotent"
var triggerBatch = ""

// Trigger 7 — Deploy: container format: "docker"
var triggerDeploy = ""

// Trigger 8 — Go concurrency: lightweight thread: "goroutine"
var triggerGo = ""

// Trigger 9 — Invoice rejected: validate with: "xsd"
var triggerInvoice = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("sql", contains(triggerSlowSQL, "explain"), "EXPLAIN plan")
	assert("api", contains(triggerAPI, "resource") || contains(triggerAPI, "rest"), "REST resources")
	assert("auth", norm(triggerAuth) == "exp", "exp claim")
	assert("outage", contains(triggerOutage, "retry") || contains(triggerOutage, "backoff"), "retry/backoff")
	assert("realtime", contains(triggerRealtime, "websocket") || contains(triggerRealtime, "ws"), "WebSocket")
	assert("batch", contains(triggerBatch, "idempotent"), "idempotent tasks")
	assert("deploy", contains(triggerDeploy, "docker"), "Docker image")
	assert("go", contains(triggerGo, "goroutine"), "goroutines")
	assert("invoice", contains(triggerInvoice, "xsd"), "XSD validation")
	fmt.Println("\nRevision 06 passed — full week sweep.")
}
