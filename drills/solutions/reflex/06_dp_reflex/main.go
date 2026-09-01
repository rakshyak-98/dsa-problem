// SOLUTION — Reflex 06 DP (peek after honest attempt)
package main

import "fmt"

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

func coinChange(coins []int, amount int) int {
	// dp[a] = fewest coins that make exactly a. Unbounded: each coin may be
	// reused, so the inner loop runs forward over amounts.
	const unreachable = 1 << 30
	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = unreachable
		for _, c := range coins {
			if c <= a && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}
	if dp[amount] >= unreachable {
		return -1
	}
	return dp[amount]
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	// minCostClimbingStairs — 3-step, 2-step, single, equal, cheap first
	assert("minCostClimbingStairs basic", minCostClimbingStairs([]int{10, 15, 20}) == 15)
	assert("minCostClimbingStairs two", minCostClimbingStairs([]int{1, 100}) == 1)
	assert("minCostClimbingStairs single", minCostClimbingStairs([]int{5}) == 0)
	assert("minCostClimbingStairs equal", minCostClimbingStairs([]int{5, 5, 5}) == 5)
	assert("minCostClimbingStairs cheap start", minCostClimbingStairs([]int{0, 1, 1, 1}) == 1)

	// rob — classic, single/two, empty, alternating, all same, skip middle
	assert("rob basic", rob([]int{2, 7, 9, 3, 1}) == 12)
	assert("rob single", rob([]int{5}) == 5)
	assert("rob two pick max", rob([]int{2, 1}) == 2)
	assert("rob empty", rob([]int{}) == 0)
	assert("rob alternating", rob([]int{5, 1, 5, 1}) == 10)
	assert("rob all same", rob([]int{3, 3, 3}) == 6)
	assert("rob endpoints", rob([]int{2, 1, 2}) == 4)

	// climbStairs — 0..6
	assert("climbStairs five", climbStairs(5) == 8)
	assert("climbStairs one", climbStairs(1) == 1)
	assert("climbStairs two", climbStairs(2) == 2)
	assert("climbStairs three", climbStairs(3) == 3)
	assert("climbStairs zero", climbStairs(0) == 0)
	assert("climbStairs four", climbStairs(4) == 5)
	assert("climbStairs six", climbStairs(6) == 13)

	// coinChange — exact, impossible, zero amount, single coin, greedy trap
	assert("coinChange basic", coinChange([]int{1, 2, 5}, 11) == 3)
	assert("coinChange impossible", coinChange([]int{2}, 3) == -1)
	assert("coinChange zero", coinChange([]int{1}, 0) == 0)
	assert("coinChange single coin", coinChange([]int{7}, 14) == 2)
	assert("coinChange greedy trap", coinChange([]int{1, 3, 4}, 6) == 2)
	assert("coinChange no coins", coinChange([]int{}, 5) == -1)

	fmt.Println("\nAll DP reflex drills passed.")
	fmt.Println("Primary: dynamic_programming/easy/fibonacci_number.js")
}
