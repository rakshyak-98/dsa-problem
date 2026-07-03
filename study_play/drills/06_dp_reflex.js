/**
 * REFLEX DRILL 06 — Dynamic Programming (1D)
 *
 * RUN: node study_play/drills/06_dp_reflex.js
 */

// TODO: REFLEX — nth Fibonacci (O(n) tabulation)
function fib(n) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — min cost to climb stairs (cost[i] to step i)
function minCostClimbingStairs(cost) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — max money from house robber (no adjacent houses)
function rob(nums) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — count ways to climb n stairs (1 or 2 steps)
function climbStairs(n) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

assert("fib", fib(10) === 55);
assert("minCostClimbingStairs", minCostClimbingStairs([10, 15, 20]) === 15);
assert("rob", rob([2, 7, 9, 3, 1]) === 12);
assert("climbStairs", climbStairs(5) === 8);

console.log("\nAll DP reflex drills passed.");

module.exports = { fib, minCostClimbingStairs, rob, climbStairs };
