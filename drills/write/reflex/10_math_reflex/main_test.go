package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func TestGcd(t *testing.T) {
	assert(t, "gcd basic", gcd(48, 18) == 6)
	assert(t, "gcd coprime", gcd(17, 13) == 1)
	assert(t, "gcd zero b", gcd(0, 7) == 7)
	assert(t, "gcd zero a", gcd(7, 0) == 7)
	assert(t, "gcd both zero", gcd(0, 0) == 0)
	assert(t, "gcd negatives", gcd(-48, 18) == 6)
	assert(t, "gcd both negative", gcd(-48, -18) == 6)
}

func TestLcm(t *testing.T) {
	assert(t, "lcm basic", lcm(4, 6) == 12)
	assert(t, "lcm coprime", lcm(7, 11) == 77)
	assert(t, "lcm zero", lcm(0, 5) == 0)
	assert(t, "lcm same", lcm(8, 8) == 8)
}

func TestModPow(t *testing.T) {
	assert(t, "modPow small", modPow(2, 10, 1000) == 24)
	assert(t, "modPow base", modPow(3, 4, 100) == 81)
	assert(t, "modPow mod1", modPow(5, 100, 1) == 0)
	assert(t, "modPow exp0", modPow(2, 0, 100) == 1)
	assert(t, "modPow exp1", modPow(7, 1, 10) == 7)
	assert(t, "modPow even exp", modPow(2, 8, 100) == 56)
}

func TestNCr(t *testing.T) {
	assert(t, "nCr basic", nCr(5, 2) == 10)
	assert(t, "nCr symmetry", nCr(10, 8) == 45)
	assert(t, "nCr invalid high", nCr(5, 6) == 0)
	assert(t, "nCr invalid negative", nCr(5, -1) == 0)
	assert(t, "nCr k0", nCr(5, 0) == 1)
	assert(t, "nCr kn", nCr(5, 5) == 1)
	assert(t, "nCr n0k0", nCr(0, 0) == 1)
}

func TestIsPrime(t *testing.T) {
	assert(t, "isPrime seventeen", isPrime(17) == true)
	assert(t, "isPrime one", isPrime(1) == false)
	assert(t, "isPrime composite", isPrime(15) == false)
	assert(t, "isPrime two", isPrime(2) == true)
	assert(t, "isPrime zero", isPrime(0) == false)
	assert(t, "isPrime square", isPrime(9) == false)
	assert(t, "isPrime even composite", isPrime(4) == false)
}

func TestPowOfTwo(t *testing.T) {
	assert(t, "powOfTwo sixtyfour", powOfTwo(64) == true)
	assert(t, "powOfTwo six", powOfTwo(6) == false)
	assert(t, "powOfTwo zero", powOfTwo(0) == false)
	assert(t, "powOfTwo one", powOfTwo(1) == true)
	assert(t, "powOfTwo two", powOfTwo(2) == true)
	assert(t, "powOfTwo eight", powOfTwo(8) == true)
	assert(t, "powOfTwo negative", powOfTwo(-4) == false)
}

func TestAll(t *testing.T) {
	t.Run("gcd", TestGcd)
	t.Run("lcm", TestLcm)
	t.Run("modPow", TestModPow)
	t.Run("nCr", TestNCr)
	t.Run("isPrime", TestIsPrime)
	t.Run("powOfTwo", TestPowOfTwo)
	fmt.Println("\nAll math reflex drills passed.")
}
