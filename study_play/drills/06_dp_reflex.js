/**
 * REFLEX DRILL 06 — Dynamic Programming (1D)
 *
 * RUN: node study_play/drills/06_dp_reflex.js
 */

// TODO: REFLEX — nth Fibonacci (O(n) tabulation)
function fib(n) {
  if(n <= 1 ) return 1;
  let prev = 0;
  let curr = 1;
  for(let i = 2; i <= n; i++){
    const next = prev + curr;
    prev = curr;
    curr = next;
  }
  return curr; 
}

// TODO: REFLEX — min cost to climb stairs (cost[i] to step i)
function minCostClimbingStairs(cost) {
  const n = cost.length;
  let prev2 = cost[0];
  let prev1 = cost[1];

  for (let i = 2; i < n; i++) {
    const curr = cost[i] + Math.min(prev1, prev2);
    prev2 = prev1;
    prev1 = curr;
  }
  return Math.min(prev1, prev2);
}

// TODO: REFLEX — max money from house robber (no adjacent houses)
function rob(nums) {
  const n = nums.length;
  if (n === 0) return 0;
  if (n === 1) return nums[0]

  let prev2 = 0;
  let prev1 = 0;

  for (const num of nums ){
    const current = Math.max(prev1, prev2 + nums);
    prev2 - prev1;
    prev1 = current
  }
  return prev1;
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
