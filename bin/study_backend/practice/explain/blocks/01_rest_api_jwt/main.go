// BLOCK 01 — REST API design & JWT (MST + iBind)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — HTTP status for resource created: "201"
var statusCreated = ""

// TODO: EXPLAIN — safe retry for update without side effects: "put" (idempotent)
var idempotentUpdate = ""

// TODO: EXPLAIN — JWT segments count: "3"
var jwtSegmentCount = ""

// TODO: EXPLAIN — validate JWT signature with what key type for RS256? "public"
var jwtVerifyKey = ""

// TODO: EXPLAIN — API versioning strategy you used: "url" or "header" (either ok)
var apiVersioning = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("201", norm(statusCreated) == "201", "201 Created")
	assert("put", contains(idempotentUpdate, "put"), "PUT is idempotent")
	assert("3 parts", norm(jwtSegmentCount) == "3", "header.payload.signature")
	assert("public key", contains(jwtVerifyKey, "public"), "RS256 uses public key")
	assert("versioning", contains(apiVersioning, "url") || contains(apiVersioning, "header"), "path or header version")
	fmt.Println("\nBlock 01 passed.")
}
