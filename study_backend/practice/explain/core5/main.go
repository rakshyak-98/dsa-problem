// CORE 5 EXPLAIN — backend interview essentials (every session)
//
// GOAL: Answer out loud first, then fill every TODO: EXPLAIN.
// RUN:  go run .
// Peek: study_backend/_support/answers/explain_core5.md
package main

import (
	"fmt"
	"strings"
)

// Q1 — JWT authentication
// TODO: EXPLAIN — three parts of a JWT, separated by what character? (e.g. "dot")
var jwtPartsSeparator = ""

// TODO: EXPLAIN — where is JWT usually sent? one word: "header"
var jwtLocation = ""

// TODO: EXPLAIN — claim that enforces expiry? "exp"
var jwtExpiryClaim = ""

// Q2 — Circuit breaker
// TODO: EXPLAIN — three states: comma-separated "closed,open,half-open"
var circuitStates = ""

// TODO: EXPLAIN — open circuit behavior: "fail fast"
var circuitOpenBehavior = ""

// Q3 — SQL optimization
// TODO: EXPLAIN — index helps which clause? keyword "where"
var indexHelps = ""

// TODO: EXPLAIN — pooling reduces what? keyword "connection"
var poolingReduces = ""

// Q4 — REST vs WebSocket
// TODO: EXPLAIN — WebSocket starts as HTTP "upgrade"
var wsStartsAs = ""

// TODO: EXPLAIN — server-push chat transport: "websocket"
var chatTransport = ""

// Q5 — Retry with backoff
// TODO: EXPLAIN — jitter prevents "thundering herd"
var jitterPrevents = ""

// TODO: EXPLAIN — retries need "idempotent" operations
var retryRequires = ""

func norm(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func containsAny(s string, opts ...string) bool {
	n := norm(s)
	for _, o := range opts {
		if strings.Contains(n, norm(o)) {
			return true
		}
	}
	return false
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("jwt separator", containsAny(jwtPartsSeparator, "dot", "."), "dots separate parts")
	assert("jwt location", norm(jwtLocation) == "header", "Authorization header")
	assert("jwt exp", norm(jwtExpiryClaim) == "exp", "exp claim")
	assert("circuit states", containsAny(circuitStates, "closed,open,half-open", "closed, open, half-open"),
		"closed → open → half-open")
	assert("circuit open", containsAny(circuitOpenBehavior, "fail fast", "failfast"), "fail fast")
	assert("index helps", containsAny(indexHelps, "where"), "WHERE / JOIN")
	assert("pooling", containsAny(poolingReduces, "connection"), "connection reuse")
	assert("ws upgrade", norm(wsStartsAs) == "upgrade", "HTTP Upgrade")
	assert("chat transport", norm(chatTransport) == "websocket", "duplex push")
	assert("jitter", containsAny(jitterPrevents, "thundering", "herd"), "thundering herd")
	assert("idempotent", containsAny(retryRequires, "idempotent"), "idempotent ops")

	fmt.Println("\nCore 5 EXPLAIN passed. Next: drills/backend/write/core5/")
}
