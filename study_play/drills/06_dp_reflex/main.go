// REFLEX DRILL 06 — Dynamic Programming (1D)
//
// RUN: go run ./drills/06_dp_reflex
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
	assert("minCostClimbingStairs", minCostClimbingStairs([]int{10, 15, 20}) == 15)
	assert("rob", rob([]int{2, 7, 9, 3, 1}) == 12)
	assert("climbStairs", climbStairs(5) == 8)

	fmt.Println("\nAll DP reflex drills passed.")
}
