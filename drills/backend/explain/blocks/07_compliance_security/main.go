// BLOCK 07 — XML validation, signing, ZATCA compliance (Opscale)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — XSD validates: "structure" or "schema"
var xsdValidates = ""

// TODO: EXPLAIN — digital signature proves: "integrity" and "authenticity"
var signatureProves = ""

// TODO: EXPLAIN — signing hash algorithm example: "sha256"
var hashAlgo = ""

// TODO: EXPLAIN — replay attack prevented by: "nonce" or "timestamp" or "id"
var replayPrevention = ""

// TODO: EXPLAIN — audit trail stores: "who" and "when" changes
var auditStores = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("xsd", contains(xsdValidates, "structure") || contains(xsdValidates, "schema"), "XML schema")
	assert("signature", contains(signatureProves, "integrity") || contains(signatureProves, "authentic"), "integrity/authenticity")
	assert("sha", contains(hashAlgo, "sha"), "SHA-256 common")
	assert("replay", contains(replayPrevention, "nonce") || contains(replayPrevention, "timestamp") || contains(replayPrevention, "id"), "unique request id")
	assert("audit", contains(auditStores, "who") || contains(auditStores, "when"), "actor + timestamp")
	fmt.Println("\nBlock 07 passed.")
}
