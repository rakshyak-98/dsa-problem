/**
 * PATTERN CHEAT SHEET — Muscle memory templates
 *
 * HOW TO USE:
 * 1. Read each template once with comments.
 * 2. Cover the file. Re-type from memory into a blank file.
 * 3. Repeat daily until you can write any section in under 2 minutes.
 *
 * Do NOT run this file for logic — it's reference + blank practice skeletons.
 */

// =============================================================================
// 1. HASH MAP — Two Sum / complement lookup
// Time: O(n)  |  Space: O(n)
// Trigger: "pair that sums to target", "complement exists"
// =============================================================================
function twoSumTemplate(nums, target) {
  const seen = new Map(); // value -> index
  for (let i = 0; i < nums.length; i++) {
    const need = target - nums[i];
    if (seen.has(need)) return [seen.get(need), i];
    seen.set(nums[i], i);
  }
  return [];
}

// =============================================================================
// 2. FREQUENCY MAP — count occurrences
// Trigger: "most frequent", "anagram", "duplicate"
// =============================================================================
function buildFreqMap(arr) {
  const freq = new Map();
  for (const x of arr) {
    freq.set(x, (freq.get(x) || 0) + 1);
  }
  return freq;
}

// =============================================================================
// 3. TWO POINTERS — opposite ends on sorted array
// Trigger: sorted + pair sum, palindrome, container area
// =============================================================================
function twoPointersOpposite(nums) {
  let left = 0;
  let right = nums.length - 1;
  while (left < right) {
    // use nums[left], nums[right]
    // move left++ or right-- based on condition
    left++;
    right--;
  }
}

// =============================================================================
// 4. TWO POINTERS — same direction (fast/slow or write index)
// Trigger: remove duplicates in-place, move zeroes, partition
// =============================================================================
function twoPointersSameDirection(nums) {
  let write = 0;
  for (let read = 0; read < nums.length; read++) {
    if (/* keep nums[read] */) {
      nums[write] = nums[read];
      write++;
    }
  }
  return write; // new length
}

// =============================================================================
// 5. SLIDING WINDOW — fixed size k
// Trigger: "subarray of size k", "maximum average"
// =============================================================================
function slidingWindowFixed(nums, k) {
  let windowSum = 0;
  for (let i = 0; i < k; i++) windowSum += nums[i];
  let best = windowSum;
  for (let i = k; i < nums.length; i++) {
    windowSum += nums[i] - nums[i - k];
    best = Math.max(best, windowSum);
  }
  return best;
}

// =============================================================================
// 6. SLIDING WINDOW — variable size
// Trigger: "longest/shortest subarray where condition"
// =============================================================================
function slidingWindowVariable(s) {
  let left = 0;
  let best = 0;
  const state = {}; // or Map — window state
  for (let right = 0; right < s.length; right++) {
    // expand: include s[right] in state
    while (/* window invalid */) {
      // shrink: remove s[left] from state
      left++;
    }
    best = Math.max(best, right - left + 1);
  }
  return best;
}

// =============================================================================
// 7. BINARY SEARCH — exact search
// Trigger: sorted array, find index, O(log n)
// =============================================================================
function binarySearchExact(nums, target) {
  let lo = 0;
  let hi = nums.length - 1;
  while (lo <= hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    if (nums[mid] === target) return mid;
    if (nums[mid] < target) lo = mid + 1;
    else hi = mid - 1;
  }
  return -1;
}

// =============================================================================
// 8. BINARY SEARCH — lower bound (first >= target)
// =============================================================================
function lowerBound(nums, target) {
  let lo = 0;
  let hi = nums.length; // exclusive
  while (lo < hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    if (nums[mid] < target) lo = mid + 1;
    else hi = mid;
  }
  return lo;
}

// =============================================================================
// 9. PREFIX SUM
// Trigger: range sum query, subarray sum equals k
// =============================================================================
function buildPrefix(nums) {
  const prefix = [0];
  for (let i = 0; i < nums.length; i++) {
    prefix.push(prefix[prefix.length - 1] + nums[i]);
  }
  return prefix;
  // sum(i..j) = prefix[j+1] - prefix[i]
}

// =============================================================================
// 10. BFS — grid shortest path / level order
// =============================================================================
function bfsGrid(grid, startR, startC) {
  const rows = grid.length;
  const cols = grid[0].length;
  const queue = [[startR, startC]];
  const visited = new Set([`${startR},${startC}`]);
  const dirs = [
    [0, 1],
    [0, -1],
    [1, 0],
    [-1, 0],
  ];
  while (queue.length) {
    const [r, c] = queue.shift();
    for (const [dr, dc] of dirs) {
      const nr = r + dr;
      const nc = c + dc;
      const key = `${nr},${nc}`;
      if (nr < 0 || nc < 0 || nr >= rows || nc >= cols) continue;
      if (visited.has(key)) continue;
      // if (grid[nr][nc] blocks) continue;
      visited.add(key);
      queue.push([nr, nc]);
    }
  }
}

// =============================================================================
// 11. DFS — grid / tree recursion
// =============================================================================
function dfsGrid(grid, r, c, visited) {
  const rows = grid.length;
  const cols = grid[0].length;
  if (r < 0 || c < 0 || r >= rows || c >= cols) return;
  const key = `${r},${c}`;
  if (visited.has(key)) return;
  visited.add(key);
  dfsGrid(grid, r + 1, c, visited);
  dfsGrid(grid, r - 1, c, visited);
  dfsGrid(grid, r, c + 1, visited);
  dfsGrid(grid, r, c - 1, visited);
}

// =============================================================================
// 12. 1D DP — bottom-up
// Trigger: "minimum cost", "max ways", Fibonacci-style
// =============================================================================
function dp1D(n) {
  const dp = new Array(n + 1).fill(0);
  dp[0] = /* base */;
  dp[1] = /* base */;
  for (let i = 2; i <= n; i++) {
    dp[i] = /* recurrence using dp[i-1], dp[i-2], etc. */;
  }
  return dp[n];
}

// =============================================================================
// 13. MONOTONIC STACK — next greater element pattern
// =============================================================================
function nextGreaterTemplate(nums) {
  const stack = []; // indices, decreasing values
  const result = new Array(nums.length).fill(-1);
  for (let i = 0; i < nums.length; i++) {
    while (stack.length && nums[i] > nums[stack[stack.length - 1]]) {
      const idx = stack.pop();
      result[idx] = nums[i];
    }
    stack.push(i);
  }
  return result;
}

module.exports = {
  twoSumTemplate,
  buildFreqMap,
  twoPointersOpposite,
  twoPointersSameDirection,
  slidingWindowFixed,
  slidingWindowVariable,
  binarySearchExact,
  lowerBound,
  buildPrefix,
  bfsGrid,
  dfsGrid,
  dp1D,
  nextGreaterTemplate,
};
