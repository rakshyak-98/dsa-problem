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

func TestFib(t *testing.T) {
	assert(t, "fib ten", fib(10) == 55)
	assert(t, "fib zero", fib(0) == 0)
	assert(t, "fib one", fib(1) == 1)
	assert(t, "fib two", fib(2) == 1)
	assert(t, "fib three", fib(3) == 2)
	assert(t, "fib five", fib(5) == 5)
	assert(t, "fib twenty", fib(20) == 6765)
}

func TestMinCostClimbingStairs(t *testing.T) {
	assert(t, "minCostClimbingStairs basic", minCostClimbingStairs([]int{10, 15, 20}) == 15)
	assert(t, "minCostClimbingStairs two", minCostClimbingStairs([]int{1, 100}) == 1)
	assert(t, "minCostClimbingStairs single", minCostClimbingStairs([]int{5}) == 0)
	assert(t, "minCostClimbingStairs equal", minCostClimbingStairs([]int{5, 5, 5}) == 5)
	assert(t, "minCostClimbingStairs cheap start", minCostClimbingStairs([]int{0, 1, 1, 1}) == 1)
}

func TestRob(t *testing.T) {
	assert(t, "rob basic", rob([]int{2, 7, 9, 3, 1}) == 12)
	assert(t, "rob single", rob([]int{5}) == 5)
	assert(t, "rob two pick max", rob([]int{2, 1}) == 2)
	assert(t, "rob empty", rob([]int{}) == 0)
	assert(t, "rob alternating", rob([]int{5, 1, 5, 1}) == 10)
	assert(t, "rob all same", rob([]int{3, 3, 3}) == 6)
	assert(t, "rob endpoints", rob([]int{2, 1, 2}) == 4)
}

func TestClimbStairs(t *testing.T) {
	assert(t, "climbStairs five", climbStairs(5) == 8)
	assert(t, "climbStairs one", climbStairs(1) == 1)
	assert(t, "climbStairs two", climbStairs(2) == 2)
	assert(t, "climbStairs three", climbStairs(3) == 3)
	assert(t, "climbStairs zero", climbStairs(0) == 0)
	assert(t, "climbStairs four", climbStairs(4) == 5)
	assert(t, "climbStairs six", climbStairs(6) == 13)
}

func TestAll(t *testing.T) {
	t.Run("fib", TestFib)
	t.Run("minCostClimbingStairs", TestMinCostClimbingStairs)
	t.Run("rob", TestRob)
	t.Run("climbStairs", TestClimbStairs)
	fmt.Println("\nAll DP reflex drills passed.")
}
