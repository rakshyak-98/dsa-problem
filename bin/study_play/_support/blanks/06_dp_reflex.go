//go:build ignore

// REFLEX DRILL 06 — Dynamic Programming (1D)
//
// RUN: go run -C drills/write/reflex/06_dp_reflex .
//
// AFTER PASSING: dynamic_programming/easy/fibonacci_number.js
package main

import "fmt"

// TODO: REFLEX — nth Fibonacci (O(n) tabulation)
func fib(n int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — min cost to climb stairs (cost[i] to step i)
func minCostClimbingStairs(cost []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — max money from house robber (no adjacent houses)
func rob(nums []int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — count ways to climb n stairs (1 or 2 steps)
func climbStairs(n int) int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	assert("fib", fib(10) == 55)
	assert("fib base", fib(0) == 0)
	assert("fib one", fib(1) == 1)

	assert("minCostClimbingStairs", minCostClimbingStairs([]int{10, 15, 20}) == 15)
	assert("minCostClimbingStairs two", minCostClimbingStairs([]int{1, 100}) == 1)

	assert("rob", rob([]int{2, 7, 9, 3, 1}) == 12)
	assert("rob single", rob([]int{5}) == 5)
	assert("rob two", rob([]int{2, 1}) == 2)

	assert("climbStairs", climbStairs(5) == 8)
	assert("climbStairs one", climbStairs(1) == 1)
	assert("climbStairs two", climbStairs(2) == 2)

	fmt.Println("\nAll DP reflex drills passed.")
	fmt.Println("Primary: dynamic_programming/easy/fibonacci_number.js")
}
