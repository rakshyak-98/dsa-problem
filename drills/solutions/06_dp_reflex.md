# Drill 06 — DP (1D)

## State first
Write `dp[i]` meaning in English before coding.

## fib / climbStairs
- **Transition:** dp[i] = dp[i-1] + dp[i-2]
- **Space opt:** two variables rolling

## minCostClimbingStairs
- **dp[i]** = min cost to stand on step i
- **Answer:** min(dp[n-1], dp[n-2]) to step past top

## rob
- **Transition:** max(rob[i]+dp[i-2], dp[i-1])
- **Bug:** take adjacent houses
