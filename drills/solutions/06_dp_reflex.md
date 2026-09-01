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

## coinChange
- **Trigger:** fewest items making an exact total, coins reusable
- **Pattern:** `dp[a] = min(dp[a-c] + 1)` over every coin c ≤ a
- **Greedy fails:** coins {1,3,4}, amount 6 → greedy gives 3 (4+1+1), answer is 2 (3+3)
- **Bug:** using 0 as "unreachable" — it collides with the real `dp[0] = 0`
