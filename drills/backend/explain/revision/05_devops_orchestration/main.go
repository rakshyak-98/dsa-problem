// REVISION 05 — DevOps + orchestration recap (Friday)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — Docker health: traffic gate uses: "readiness"
var trafficHealth = ""

// TODO: EXPLAIN — Docker health: restart probe: "liveness"
var restartHealth = ""

// TODO: EXPLAIN — short event handler on AWS: "lambda"
var eventCompute = ""

// TODO: EXPLAIN — static asset edge cache: "cdn" or "cloudfront"
var staticEdge = ""

// TODO: EXPLAIN — pipeline runs tests before deploy: "ci" or "continuous integration"
var preDeployGate = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("readiness", contains(trafficHealth, "readiness"), "readiness gates traffic")
	assert("liveness", contains(restartHealth, "liveness"), "liveness restarts pod")
	assert("lambda", contains(eventCompute, "lambda"), "Lambda for events")
	assert("cdn", contains(staticEdge, "cdn") || contains(staticEdge, "cloudfront"), "CDN at edge")
	assert("ci", contains(preDeployGate, "ci") || contains(preDeployGate, "continuous"), "CI before CD")
	fmt.Println("\nRevision 05 passed.")
}
