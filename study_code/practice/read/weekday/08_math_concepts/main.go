// READ DRILL 08 — Math concepts (formulas & traces)
//
// GOAL: Recall DSA math formulas and trace small numeric examples.
// RUN:  go run ./drills/08_math_concepts
//
// REFERENCE: study_play/docs/MATH_CONCEPTS.md
package main

import (
	"fmt"
	"strings"
)

func egcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// TODO: READ — trace egcd(48, 18)
var m1GCD = 0

func modPow(base, exp, mod int) int {
	if mod == 1 {
		return 0
	}
	base %= mod
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

// TODO: READ — modPow(3, 4, 100)
var m2ModPow = 0

// TODO: READ — C(8, 3) combinations
var m3nCr = 0

// TODO: READ — sum of integers 1 through 20: n(n+1)/2
var m4Sum = 0

// TODO: READ — naive fib(7) where fib(0)=0, fib(1)=1
var m5Fib = 0

// TODO: READ — T(n)=2T(n/2)+O(n) time complexity: "O(n log n)" or "O(n^2)"
var m6Master = ""

// TODO: READ — T(n)=T(n/2)+O(1) time complexity
var m7Recurrence = ""

// TODO: READ — popcount (set bits) in 13 (binary 1101)
var m8Bits = 0

// TODO: READ — is 1 a prime? "yes" or "no"
var m9Prime = ""

// TODO: READ — height of complete binary tree with 15 nodes: floor(log2 n)
var m10Height = 0

func norm(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func assert(name string, cond bool, hint string) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s — %s", name, hint))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("m1GCD", m1GCD == egcd(48, 18), `Euclidean: 48,18 → 18,12 → 12,6 → 6,0 → 6`)
	assert("m2ModPow", m2ModPow == modPow(3, 4, 100), `3^4 = 81`)
	assert("m3nCr", m3nCr == 56, `8!/(3!5!) = 56`)
	assert("m4Sum", m4Sum == 210, `20*21/2 = 210`)

	fib7 := 13
	assert("m5Fib", m5Fib == fib7, `0,1,1,2,3,5,8,13`)

	okMaster := norm(m6Master) == "o(nlogn)" || norm(m6Master) == "o(n*logn)"
	assert("m6Master", okMaster, `merge-sort shape → Θ(n log n)`)

	okRec := norm(m7Recurrence) == "o(logn)" || norm(m7Recurrence) == "o(log2n)"
	assert("m7Recurrence", okRec, `halve each step → O(log n)`)

	assert("m8Bits", m8Bits == 3, `1101 has three 1-bits`)
	assert("m9Prime", norm(m9Prime) == "no", `1 is not prime by definition`)
	assert("m10Height", m10Height == 3, `floor(log2 15) = 3`)

	_ = egcd
	_ = modPow
	fmt.Println("\nMath concepts drill passed.")
	fmt.Println("Reference: study_play/docs/MATH_CONCEPTS.md")
}
