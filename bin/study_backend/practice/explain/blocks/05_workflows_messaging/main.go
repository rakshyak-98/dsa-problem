// BLOCK 05 — Airflow, RabbitMQ, orchestration (Opscale ZATCA pipelines)
// RUN: go run .
package main

import (
	"fmt"
	"strings"
)

// TODO: EXPLAIN — Airflow unit of work: "task"
var airflowUnit = ""

// TODO: EXPLAIN — DAG means: "directed acyclic graph"
var dagMeaning = ""

// TODO: EXPLAIN — failed task safe rerun needs: "idempotent"
var airflowIdempotent = ""

// TODO: EXPLAIN — RabbitMQ dead letter queue for: "poison" or "failed"
var dlqFor = ""

// TODO: EXPLAIN — work queue pattern: competing: "consumers"
var workQueueConsumers = ""

func norm(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func contains(s, sub string) bool { return strings.Contains(norm(s), norm(sub)) }

func assert(name string, cond bool, hint string) {
	if !cond { panic(fmt.Sprintf("FAIL: %s — %s", name, hint)) }
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("task", contains(airflowUnit, "task"), "operators/tasks")
	assert("dag", contains(dagMeaning, "directed") && contains(dagMeaning, "acyclic"), "DAG definition")
	assert("idempotent", contains(airflowIdempotent, "idempotent"), "safe retries")
	assert("dlq", contains(dlqFor, "poison") || contains(dlqFor, "fail"), "poison messages")
	assert("consumers", contains(workQueueConsumers, "consumer"), "competing consumers")
	fmt.Println("\nBlock 05 passed.")
}
