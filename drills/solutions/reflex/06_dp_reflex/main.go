// SOLUTION — Reflex 06 DP (peek after honest attempt)
package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		c := b + cost[i-1]
		if a+cost[i-2] < c {
			c = a + cost[i-2]
		}
		a, b = b, c
	}
	return b
}

func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, x := range nums {
		cur := prev1
		if prev2+x > cur {
			cur = prev2 + x
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
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
}
