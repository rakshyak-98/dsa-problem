// BLOCK 04 — WebSocket, WebRTC, STUN/TURN (iBind video platform)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — WebRTC media path is: "peer" or "p2p"
var mediaPath = ""

// TODO: EXPLAIN — signaling carries: "sdp" or "session description"
var signalingCarries = ""

// TODO: EXPLAIN — ICE gathers: "candidates"
var iceGathers = ""

// TODO: EXPLAIN — STUN discovers public: "address" or "ip"
var stunDiscovers = ""

// TODO: EXPLAIN — TURN used when direct P2P: "fails" or "blocked"
var turnWhen = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("p2p", contains(mediaPath, "peer") || contains(mediaPath, "p2p"), "peer-to-peer media")
	assert("sdp", contains(signalingCarries, "sdp") || contains(signalingCarries, "session"), "SDP offer/answer")
	assert("candidates", contains(iceGathers, "candidate"), "ICE candidates")
	assert("stun", contains(stunDiscovers, "address") || contains(stunDiscovers, "ip"), "public IP")
	assert("turn", contains(turnWhen, "fail") || contains(turnWhen, "block"), "NAT blocks direct path")
	fmt.Println("\nBlock 04 passed.")
}
