//go:build ignore

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
	// gcd — basic, coprime, zero sides, both zero, negatives
	assert("gcd basic", gcd(48, 18) == 6)
	assert("gcd coprime", gcd(17, 13) == 1)
	assert("gcd zero b", gcd(0, 7) == 7)
	assert("gcd zero a", gcd(7, 0) == 7)
	assert("gcd both zero", gcd(0, 0) == 0)
	assert("gcd negatives", gcd(-48, 18) == 6)
	assert("gcd both negative", gcd(-48, -18) == 6)

	// lcm — basic, coprime, zero, same number
	assert("lcm basic", lcm(4, 6) == 12)
	assert("lcm coprime", lcm(7, 11) == 77)
	assert("lcm zero", lcm(0, 5) == 0)
	assert("lcm same", lcm(8, 8) == 8)

	// modPow — small, base case, mod 1, exp 0/1, even exp
	assert("modPow small", modPow(2, 10, 1000) == 24)
	assert("modPow base", modPow(3, 4, 100) == 81)
	assert("modPow mod1", modPow(5, 100, 1) == 0)
	assert("modPow exp0", modPow(2, 0, 100) == 1)
	assert("modPow exp1", modPow(7, 1, 10) == 7)
	assert("modPow even exp", modPow(2, 8, 100) == 56)

	// nCr — basic, symmetry, invalid, k=0, k=n, n=0
	assert("nCr basic", nCr(5, 2) == 10)
	assert("nCr symmetry", nCr(10, 8) == 45)
	assert("nCr invalid high", nCr(5, 6) == 0)
	assert("nCr invalid negative", nCr(5, -1) == 0)
	assert("nCr k0", nCr(5, 0) == 1)
	assert("nCr kn", nCr(5, 5) == 1)
	assert("nCr n0k0", nCr(0, 0) == 1)

	// isPrime — 2, odd prime, composite, 1, 0, square, even composite
	assert("isPrime seventeen", isPrime(17) == true)
	assert("isPrime one", isPrime(1) == false)
	assert("isPrime composite", isPrime(15) == false)
	assert("isPrime two", isPrime(2) == true)
	assert("isPrime zero", isPrime(0) == false)
	assert("isPrime square", isPrime(9) == false)
	assert("isPrime even composite", isPrime(4) == false)

	// powOfTwo — powers, non-powers, zero/negative
	assert("powOfTwo sixtyfour", powOfTwo(64) == true)
	assert("powOfTwo six", powOfTwo(6) == false)
	assert("powOfTwo zero", powOfTwo(0) == false)
	assert("powOfTwo one", powOfTwo(1) == true)
	assert("powOfTwo two", powOfTwo(2) == true)
	assert("powOfTwo eight", powOfTwo(8) == true)
	assert("powOfTwo negative", powOfTwo(-4) == false)

	fmt.Println("\nAll math reflex drills passed.")
	fmt.Println("Reference: doc/write/MATH_CONCEPTS.md")
}
