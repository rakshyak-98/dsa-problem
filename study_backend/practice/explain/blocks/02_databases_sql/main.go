// BLOCK 02 — Databases & SQL optimization (MST 35%, iBind pooling)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — default MySQL InnoDB index structure: "b-tree" or "btree"
var defaultIndexType = ""

// TODO: EXPLAIN — composite index column order matters for left-prefix rule: "yes"
var compositeOrderMatters = ""

// TODO: EXPLAIN — N+1 problem fixed by: "eager" or "join" or "batch"
var nPlusOneFix = ""

// TODO: EXPLAIN — connection pool avoids repeated: "handshake" or "tcp"
var poolAvoids = ""

// TODO: EXPLAIN — EXPLAIN shows: "plan" or "execution plan"
var explainShows = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("btree", contains(defaultIndexType, "b"), "B-tree indexes")
	assert("prefix", norm(compositeOrderMatters) == "yes", "leftmost prefix rule")
	assert("n+1", contains(nPlusOneFix, "eager") || contains(nPlusOneFix, "join") || contains(nPlusOneFix, "batch"), "batch/join")
	assert("pool", contains(poolAvoids, "handshake") || contains(poolAvoids, "tcp"), "TCP/auth handshake")
	assert("explain", contains(explainShows, "plan"), "execution plan")
	fmt.Println("\nBlock 02 passed.")
}
