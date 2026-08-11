// REVISION 04 — Real-time + messaging recap (Thursday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — relay when P2P fails: "turn"
var p2pRelay = ""

// TODO: EXPLAIN — signaling transport: "websocket"
var signalingTransport = ""

// TODO: EXPLAIN — Airflow graph type: "dag"
var workflowGraph = ""

// TODO: EXPLAIN — poison messages go to: "dlq" or "dead letter"
var poisonQueue = ""

// TODO: EXPLAIN — competing consumers pattern: "work queue"
var queuePattern = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("turn", contains(p2pRelay, "turn"), "TURN relays media")
	assert("ws", contains(signalingTransport, "websocket") || contains(signalingTransport, "ws"), "WebSocket signaling")
	assert("dag", contains(workflowGraph, "dag"), "directed acyclic graph")
	assert("dlq", contains(poisonQueue, "dlq") || contains(poisonQueue, "dead"), "dead letter queue")
	assert("work queue", contains(queuePattern, "work"), "work queue pattern")
	fmt.Println("\nRevision 04 passed.")
}
