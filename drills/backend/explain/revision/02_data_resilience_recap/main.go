// REVISION 02 — Data layer + resilience recap (Tuesday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — first step for slow SQL: "explain"
var slowSQLFirstStep = ""

// TODO: EXPLAIN — index helps which clause: "where"
var indexClause = ""

// TODO: EXPLAIN — pool reuses: "connection" or "connections"
var poolReuses = ""

// TODO: EXPLAIN — circuit breaker open behavior: "fail fast"
var breakerOpen = ""

// TODO: EXPLAIN — backoff multiplier base: "2"
var backoffMultiplier = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("explain", contains(slowSQLFirstStep, "explain"), "EXPLAIN ANALYZE first")
	assert("where", contains(indexClause, "where"), "WHERE/JOIN")
	assert("pool", contains(poolReuses, "connection"), "connection reuse")
	assert("fail fast", contains(breakerOpen, "fail"), "open circuit fails fast")
	assert("backoff", norm(backoffMultiplier) == "2", "double each attempt")
	fmt.Println("\nRevision 02 passed.")
}
