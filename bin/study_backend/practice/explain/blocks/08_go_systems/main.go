// BLOCK 08 — Go systems programming (BitTorrent + HTTP server projects)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — goroutine scheduling model: "m:n" or "multiplexed"
var goroutineModel = ""

// TODO: EXPLAIN — context.Context used for: "cancellation" or "deadline"
var contextFor = ""

// TODO: EXPLAIN — BitTorrent metadata encoding: "bencode"
var torrentEncoding = ""

// TODO: EXPLAIN — HTTP keep-alive reuses: "connection" or "tcp"
var keepAliveReuses = ""

// TODO: EXPLAIN — graceful shutdown closes: "listener" then drains requests
var gracefulShutdown = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("m:n", contains(goroutineModel, "m") || contains(goroutineModel, "multiplex"), "M goroutines on N threads")
	assert("context", contains(contextFor, "cancel") || contains(contextFor, "deadline"), "cancel/deadline")
	assert("bencode", contains(torrentEncoding, "bencode"), "bencode format")
	assert("keepalive", contains(keepAliveReuses, "connection") || contains(keepAliveReuses, "tcp"), "TCP reuse")
	assert("shutdown", contains(gracefulShutdown, "listener") || contains(gracefulShutdown, "drain"), "stop accepting, drain")
	fmt.Println("\nBlock 08 passed.")
}
