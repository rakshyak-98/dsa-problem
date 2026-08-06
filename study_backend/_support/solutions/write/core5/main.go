// CORE 5 WRITE — Go reflex for backend patterns (SOLUTION — peek after attempt)
package main

import (
	"fmt"
	"strings"
	"time"
)

func exponentialBackoff(attempt int, base, maxDelay time.Duration) time.Duration {
	delay := base * time.Duration(1<<attempt)
	if delay > maxDelay {
		return maxDelay
	}
	return delay
}

func circuitState(currentState string, failures, threshold int, reset bool) string {
	switch strings.ToLower(currentState) {
	case "open":
		if reset {
			return "half-open"
		}
		return "open"
	case "half-open":
		if failures >= threshold {
			return "open"
		}
		return "half-open"
	default:
		if failures >= threshold {
			return "open"
		}
		return "closed"
	}
}

func httpStatusClass(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "success"
	case code >= 300 && code < 400:
		return "redirect"
	case code >= 400 && code < 500:
		return "client_error"
	case code >= 500 && code < 600:
		return "server_error"
	default:
		return "unknown"
	}
}

func isIdempotentMethod(method string) bool {
	switch strings.ToUpper(method) {
	case "GET", "PUT", "DELETE", "HEAD", "OPTIONS":
		return true
	default:
		return false
	}
}

func parseBearerToken(authHeader string) string {
	parts := strings.Fields(strings.TrimSpace(authHeader))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	d0 := exponentialBackoff(0, 100*time.Millisecond, 2*time.Second)
	d3 := exponentialBackoff(3, 100*time.Millisecond, 2*time.Second)
	assert("backoff attempt0", d0 == 100*time.Millisecond)
	assert("backoff capped", d3 == 800*time.Millisecond)

	assert("circuit closed", circuitState("closed", 2, 5, false) == "closed")
	assert("circuit open", circuitState("closed", 5, 5, false) == "open")
	assert("circuit half", circuitState("open", 0, 5, true) == "half-open")

	assert("status 200", httpStatusClass(200) == "success")
	assert("status 404", httpStatusClass(404) == "client_error")
	assert("status 503", httpStatusClass(503) == "server_error")

	assert("idempotent get", isIdempotentMethod("GET"))
	assert("idempotent put", isIdempotentMethod("PUT"))
	assert("not idempotent post", !isIdempotentMethod("POST"))

	assert("bearer ok", parseBearerToken("Bearer abc123") == "abc123")
	assert("bearer bad", parseBearerToken("Basic xyz") == "")
	assert("bearer trim", parseBearerToken("Bearer  tok ") == "tok")

	fmt.Println("\nCore 5 WRITE passed.")
}
