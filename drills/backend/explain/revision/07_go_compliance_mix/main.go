// REVISION 07 — Go systems + compliance mix (Sunday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — ZATCA step before submit: "sign" or "signing"
var zatcaBeforeSubmit = ""

// TODO: EXPLAIN — schema validation format: "xsd"
var schemaValidation = ""

// TODO: EXPLAIN — Go cancellation primitive: "context"
var goCancel = ""

// TODO: EXPLAIN — HTTP graceful stop method: "shutdown"
var httpGraceful = ""

// TODO: EXPLAIN — BitTorrent metadata encoding: "bencode"
var torrentEncoding = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("sign", contains(zatcaBeforeSubmit, "sign"), "cryptographic signing")
	assert("xsd", contains(schemaValidation, "xsd"), "XSD validation")
	assert("context", contains(goCancel, "context"), "context.Context")
	assert("shutdown", contains(httpGraceful, "shutdown"), "http.Server.Shutdown")
	assert("bencode", contains(torrentEncoding, "bencode"), "bencode metadata")
	fmt.Println("\nRevision 07 passed.")
}
