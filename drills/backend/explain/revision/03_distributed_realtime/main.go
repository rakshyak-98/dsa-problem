// REVISION 03 — Distributed + real-time recap (Wednesday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — during partition sacrifice consistency or availability: either "consistency" or "availability"
var capTradeoff = ""

// TODO: EXPLAIN — microservices distributed tx pattern: "saga"
var distributedTx = ""

// TODO: EXPLAIN — jitter prevents: "thundering herd"
var jitterPrevents = ""

// TODO: EXPLAIN — WebSocket starts as HTTP: "upgrade"
var wsHandshake = ""

// TODO: EXPLAIN — idempotency key location: "header"
var idempotencyKey = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("cap", contains(capTradeoff, "consistency") || contains(capTradeoff, "availability"), "CP or AP")
	assert("saga", contains(distributedTx, "saga"), "saga over 2PC")
	assert("jitter", contains(jitterPrevents, "thundering") || contains(jitterPrevents, "herd"), "thundering herd")
	assert("upgrade", norm(wsHandshake) == "upgrade", "HTTP Upgrade")
	assert("header", contains(idempotencyKey, "header"), "Idempotency-Key header")
	fmt.Println("\nRevision 03 passed.")
}
