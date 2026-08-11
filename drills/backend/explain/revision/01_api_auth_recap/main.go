// REVISION 01 — API & auth recap (Monday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — HTTP status for successful GET: "200"
var statusOK = ""

// TODO: EXPLAIN — idempotent update verb: "put"
var idempotentVerb = ""

// TODO: EXPLAIN — JWT sent in HTTP: "header"
var jwtLocation = ""

// TODO: EXPLAIN — safe POST retry header: "idempotency-key"
var postRetryHeader = ""

// TODO: EXPLAIN — API version in URL prefix: "v1" or "/v1"
var apiVersion = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("200", norm(statusOK) == "200", "200 OK")
	assert("put", contains(idempotentVerb, "put"), "PUT is idempotent")
	assert("header", norm(jwtLocation) == "header", "Authorization header")
	assert("idempotency", contains(postRetryHeader, "idempotency"), "Idempotency-Key")
	assert("version", contains(apiVersion, "v1"), "URL versioning")
	fmt.Println("\nRevision 01 passed.")
}
