/**
 * SOLUTIONS — Reference implementations for reflex drills
 *
 * USE ONLY AFTER honest attempt (15+ min per file).
 * Cover this file and re-implement in the drill files until automatic.
 *
 * Do not import this in drills during practice.
 */

// ---------- 01 arrays ----------
function reverseInPlace(arr) {
  let l = 0;
  let r = arr.length - 1;
  while (l < r) {
    [arr[l], arr[r]] = [arr[r], arr[l]];
    l++;
    r--;
  }
  return arr;
}

function indexOfMax(arr) {
  let best = 0;
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] > arr[best]) best = i;
  }
  return best;
}

function arraySum(arr) {
  let sum = 0;
  for (const x of arr) sum += x;
  return sum;
}

function rotateRight(arr, k) {
  const n = arr.length;
  if (n === 0) return arr;
  k = k % n;
  return arr.slice(n - k).concat(arr.slice(0, n - k));
}

function runningSum(arr) {
  const out = [];
  let sum = 0;
  for (const x of arr) {
    sum += x;
    out.push(sum);
  }
  return out;
}

// ---------- 02 hashing ----------
function twoSum(nums, target) {
  const seen = new Map();
  for (let i = 0; i < nums.length; i++) {
    const need = target - nums[i];
    if (seen.has(need)) return [seen.get(need), i];
    seen.set(nums[i], i);
  }
  return [];
}

function containsDuplicate(nums) {
  const set = new Set();
  for (const n of nums) {
    if (set.has(n)) return true;
    set.add(n);
  }
  return false;
}

function frequencyMap(arr) {
  const map = new Map();
  for (const x of arr) map.set(x, (map.get(x) || 0) + 1);
  return map;
}

function firstUniqueChar(s) {
  const freq = frequencyMap(s.split(""));
  for (let i = 0; i < s.length; i++) {
    if (freq.get(s[i]) === 1) return s[i];
  }
  return "";
}

function groupAnagrams(strs) {
  const map = new Map();
  for (const w of strs) {
    const key = w.split("").sort().join("");
    if (!map.has(key)) map.set(key, []);
    map.get(key).push(w);
  }
  return [...map.values()];
}

// ---------- 03 two pointers ----------
function removeDuplicates(nums) {
  if (nums.length === 0) return 0;
  let write = 1;
  for (let read = 1; read < nums.length; read++) {
    if (nums[read] !== nums[read - 1]) {
      nums[write] = nums[read];
      write++;
    }
  }
  return write;
}

function moveZeroes(nums) {
  let write = 0;
  for (let read = 0; read < nums.length; read++) {
    if (nums[read] !== 0) {
      nums[write++] = nums[read];
    }
  }
  while (write < nums.length) nums[write++] = 0;
}

function maxArea(heights) {
  let l = 0;
  let r = heights.length - 1;
  let best = 0;
  while (l < r) {
    best = Math.max(best, Math.min(heights[l], heights[r]) * (r - l));
    if (heights[l] < heights[r]) l++;
    else r--;
  }
  return best;
}

function isPalindrome(s) {
  const clean = s.toLowerCase().replace(/[^a-z0-9]/g, "");
  let l = 0;
  let r = clean.length - 1;
  while (l < r) {
    if (clean[l] !== clean[r]) return false;
    l++;
    r--;
  }
  return true;
}

function maxSumSubarrayK(nums, k) {
  let sum = 0;
  for (let i = 0; i < k; i++) sum += nums[i];
  let best = sum;
  for (let i = k; i < nums.length; i++) {
    sum += nums[i] - nums[i - k];
    best = Math.max(best, sum);
  }
  return best;
}

// ---------- 04 binary search ----------
function binarySearch(nums, target) {
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

function searchInsert(nums, target) {
  let lo = 0;
  let hi = nums.length;
  while (lo < hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    if (nums[mid] < target) lo = mid + 1;
    else hi = mid;
  }
  return lo;
}

function findMinRotated(nums) {
  let lo = 0;
  let hi = nums.length - 1;
  while (lo < hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    if (nums[mid] > nums[hi]) lo = mid + 1;
    else hi = mid;
  }
  return lo;
}

function isTargetPresent(nums, target) {
  return binarySearch(nums, target) !== -1;
}

// ---------- 05 trees/stacks ----------
class TreeNode {
  constructor(val, left = null, right = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

function inorderTraversal(root) {
  const out = [];
  const stack = [];
  let cur = root;
  while (cur || stack.length) {
    while (cur) {
      stack.push(cur);
      cur = cur.left;
    }
    cur = stack.pop();
    out.push(cur.val);
    cur = cur.right;
  }
  return out;
}

function maxDepth(root) {
  if (!root) return 0;
  return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
}

function isValidParentheses(s) {
  const stack = [];
  const match = { ")": "(", "]": "[", "}": "{" };
  for (const ch of s) {
    if (ch === "(" || ch === "[" || ch === "{") stack.push(ch);
    else {
      if (stack.pop() !== match[ch]) return false;
    }
  }
  return stack.length === 0;
}

function dailyTemperatures(temps) {
  const stack = [];
  const out = new Array(temps.length).fill(0);
  for (let i = 0; i < temps.length; i++) {
    while (stack.length && temps[i] > temps[stack[stack.length - 1]]) {
      const j = stack.pop();
      out[j] = i - j;
    }
    stack.push(i);
  }
  return out;
}

// ---------- 06 dp ----------
function fib(n) {
  if (n <= 1) return n;
  let a = 0;
  let b = 1;
  for (let i = 2; i <= n; i++) {
    const c = a + b;
    a = b;
    b = c;
  }
  return b;
}

function minCostClimbingStairs(cost) {
  let a = 0;
  let b = 0;
  for (let i = 2; i <= cost.length; i++) {
    const c = Math.min(b + cost[i - 1], a + cost[i - 2]);
    a = b;
    b = c;
  }
  return b;
}

function rob(nums) {
  let prev2 = 0;
  let prev1 = 0;
  for (const x of nums) {
    const cur = Math.max(prev1, prev2 + x);
    prev2 = prev1;
    prev1 = cur;
  }
  return prev1;
}

function climbStairs(n) {
  if (n <= 2) return n;
  let a = 1;
  let b = 2;
  for (let i = 3; i <= n; i++) {
    const c = a + b;
    a = b;
    b = c;
  }
  return b;
}

// ---------- 07 graphs ----------
function numIslands(grid) {
  const rows = grid.length;
  const cols = grid[0].length;
  let count = 0;

  function dfs(r, c) {
    if (r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] !== "1") return;
    grid[r][c] = "0";
    dfs(r + 1, c);
    dfs(r - 1, c);
    dfs(r, c + 1);
    dfs(r, c - 1);
  }

  for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
      if (grid[r][c] === "1") {
        count++;
        dfs(r, c);
      }
    }
  }
  return count;
}

function floodFill(image, sr, sc, color) {
  const start = image[sr][sc];
  if (start === color) return image;

  const rows = image.length;
  const cols = image[0].length;
  const stack = [[sr, sc]];

  while (stack.length) {
    const [r, c] = stack.pop();
    if (r < 0 || c < 0 || r >= rows || c >= cols || image[r][c] !== start) continue;
    image[r][c] = color;
    stack.push([r + 1, c], [r - 1, c], [r, c + 1], [r, c - 1]);
  }
  return image;
}

function shortestPathGrid(grid) {
  const rows = grid.length;
  const cols = grid[0].length;
  if (grid[0][0] === 1 || grid[rows - 1][cols - 1] === 1) return -1;

  const queue = [[0, 0, 1]];
  grid[0][0] = 1;
  const dirs = [
    [0, 1],
    [0, -1],
    [1, 0],
    [-1, 0],
  ];

  while (queue.length) {
    const [r, c, dist] = queue.shift();
    if (r === rows - 1 && c === cols - 1) return dist;
    for (const [dr, dc] of dirs) {
      const nr = r + dr;
      const nc = c + dc;
      if (nr < 0 || nc < 0 || nr >= rows || nc >= cols || grid[nr][nc] === 1) continue;
      grid[nr][nc] = 1;
      queue.push([nr, nc, dist + 1]);
    }
  }
  return -1;
}

module.exports = {
  reverseInPlace,
  indexOfMax,
  arraySum,
  rotateRight,
  runningSum,
  twoSum,
  containsDuplicate,
  frequencyMap,
  firstUniqueChar,
  groupAnagrams,
  removeDuplicates,
  moveZeroes,
  maxArea,
  isPalindrome,
  maxSumSubarrayK,
  binarySearch,
  searchInsert,
  findMinRotated,
  isTargetPresent,
  TreeNode,
  inorderTraversal,
  maxDepth,
  isValidParentheses,
  dailyTemperatures,
  fib,
  minCostClimbingStairs,
  rob,
  climbStairs,
  numIslands,
  floodFill,
  shortestPathGrid,
};
