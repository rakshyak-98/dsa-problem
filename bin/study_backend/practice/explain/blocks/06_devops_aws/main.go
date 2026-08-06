// BLOCK 06 — Docker, Jenkins, AWS, CI/CD (iBind Jenkins)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — Docker image built from: "layers"
var imageLayers = ""

// TODO: EXPLAIN — CI stage before deploy: "test"
var ciBeforeDeploy = ""

// TODO: EXPLAIN — Lambda best for: "short" or "event" workloads
var lambdaBestFor = ""

// TODO: EXPLAIN — CDN caches: "static" assets
var cdnCaches = ""

// TODO: EXPLAIN — container health check probes: "liveness" or "readiness"
var healthProbe = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("layers", contains(imageLayers, "layer"), "layered filesystem")
	assert("test", contains(ciBeforeDeploy, "test"), "tests in pipeline")
	assert("lambda", contains(lambdaBestFor, "short") || contains(lambdaBestFor, "event"), "short-lived/event-driven")
	assert("cdn", contains(cdnCaches, "static"), "static content edge cache")
	assert("health", contains(healthProbe, "liveness") || contains(healthProbe, "readiness"), "k8s/docker health")
	fmt.Println("\nBlock 06 passed.")
}
