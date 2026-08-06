// REFLEX DRILL 10 — Math (number theory & combinatorics)
//
// GOAL: Core math helpers for DSA = automatic.
// RUN: go run ./drills/10_math_reflex
//
// REFERENCE: doc/write/MATH_CONCEPTS.md
package main

import (
	"fmt"
)

// TODO: REFLEX — greatest common divisor (Euclidean algorithm)
func gcd(a, b int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — least common multiple using gcd
func lcm(a, b int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — (base^exp) % mod using fast exponentiation
func modPow(base, exp, mod int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — combinations C(n,k); return 0 if k < 0 or k > n
func nCr(n, k int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — return true if n is prime (n >= 2)
func isPrime(n int) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — return true if n is a positive power of 2
func powOfTwo(n int) bool {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("gcd basic", gcd(48, 18) == 6)
	assert("gcd coprime", gcd(17, 13) == 1)
	assert("gcd zero", gcd(0, 7) == 7)

	assert("lcm basic", lcm(4, 6) == 12)
	assert("lcm coprime", lcm(7, 11) == 77)

	assert("modPow small", modPow(2, 10, 1000) == 24)
	assert("modPow base", modPow(3, 4, 100) == 81)
	assert("modPow mod1", modPow(5, 100, 1) == 0)

	assert("nCr basic", nCr(5, 2) == 10)
	assert("nCr symmetry", nCr(10, 8) == 45)
	assert("nCr invalid", nCr(5, 6) == 0)

	assert("isPrime true", isPrime(17) == true)
	assert("isPrime false", isPrime(1) == false)
	assert("isPrime composite", isPrime(15) == false)

	assert("powOfTwo true", powOfTwo(64) == true)
	assert("powOfTwo false", powOfTwo(6) == false)
	assert("powOfTwo zero", powOfTwo(0) == false)

	fmt.Println("\nAll math reflex drills passed.")
	fmt.Println("Reference: doc/write/MATH_CONCEPTS.md")
}
