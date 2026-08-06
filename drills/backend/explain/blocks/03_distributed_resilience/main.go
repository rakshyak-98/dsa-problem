// BLOCK 03 — Distributed systems & resilience (iBind circuit breaker)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — exponential backoff multiplier base commonly: "2"
var backoffBase = ""

// TODO: EXPLAIN — circuit breaker prevents: "cascade" or "cascading"
var breakerPrevents = ""

// TODO: EXPLAIN — idempotency key sent in: "header" (common)
var idempotencyKeyLocation = ""

// TODO: EXPLAIN — CAP: pick two of consistency, availability, partition tolerance — you sacrifice one during partition: "consistency" or "availability"
var capSacrifice = ""

// TODO: EXPLAIN — saga pattern coordinates via: "events" or "choreography" or "orchestration"
var sagaCoordination = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("backoff 2", norm(backoffBase) == "2", "double each attempt")
	assert("cascade", contains(breakerPrevents, "cascade"), "cascading failures")
	assert("idempotency header", contains(idempotencyKeyLocation, "header"), "Idempotency-Key header")
	assert("cap", contains(capSacrifice, "consistency") || contains(capSacrifice, "availability"), "pick CP or AP")
	assert("saga", contains(sagaCoordination, "event") || contains(sagaCoordination, "choreography") || contains(sagaCoordination, "orchestration"), "distributed tx pattern")
	fmt.Println("\nBlock 03 passed.")
}
