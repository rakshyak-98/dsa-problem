// SOLUTION — Reflex 10 Math (peek after honest attempt)
package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	a, b = abs(a), abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a) / gcd(a, b) * abs(b)
}

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

func nCr(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 1; i <= k; i++ {
		result = result * (n - k + i) / i
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func powOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
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
	assert("gcd negatives", gcd(-48, 18) == 6)

	assert("lcm basic", lcm(4, 6) == 12)
	assert("lcm coprime", lcm(7, 11) == 77)
	assert("lcm zero", lcm(0, 5) == 0)

	assert("modPow small", modPow(2, 10, 1000) == 24)
	assert("modPow base", modPow(3, 4, 100) == 81)
	assert("modPow mod1", modPow(5, 100, 1) == 0)
	assert("modPow exp0", modPow(2, 0, 100) == 1)

	assert("nCr basic", nCr(5, 2) == 10)
	assert("nCr symmetry", nCr(10, 8) == 45)
	assert("nCr invalid", nCr(5, 6) == 0)
	assert("nCr k0", nCr(5, 0) == 1)
	assert("nCr kn", nCr(5, 5) == 1)

	assert("isPrime true", isPrime(17))
	assert("isPrime false", !isPrime(1))
	assert("isPrime composite", !isPrime(15))
	assert("isPrime two", isPrime(2))

	assert("powOfTwo true", powOfTwo(64))
	assert("powOfTwo false", !powOfTwo(6))
	assert("powOfTwo zero", !powOfTwo(0))
	assert("powOfTwo one", powOfTwo(1))
	assert("powOfTwo two", powOfTwo(2))

	fmt.Println("\nAll math reflex drills passed.")
}
