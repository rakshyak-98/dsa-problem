/**
 * Daily reflex practice helper — Core 5 + specialty drill
 *
 * RUN:              node study_play/daily_drill.js
 * RUN with tests:   node study_play/daily_drill.js --run
 * Core 5 only:      node study_play/daily_drill.js --micro
 * Full catalog:     node study_play/daily_drill.js --catalog
 */

const { execSync } = require("child_process");
const path = require("path");

const CORE5 = [
  {
    name: "twoSum(nums, target)",
    ask: "indices of two values that sum to target",
    pattern: "Map + complement",
    sec: 90,
  },
  {
    name: "binarySearch(nums, target)",
    ask: "index of target in sorted array, or -1",
    pattern: "lo <= hi binary search",
    sec: 60,
  },
  {
    name: "removeDuplicates(nums)",
    ask: "in-place unique prefix length on sorted array",
    pattern: "read/write two pointers",
    sec: 90,
  },
  {
    name: "maxSumSubarrayK(nums, k)",
    ask: "max sum of any contiguous window of size k",
    pattern: "fixed sliding window",
    sec: 90,
  },
  {
    name: "frequencyMap(arr)",
    ask: "Map each value to its count",
    pattern: "freq Map loop",
    sec: 60,
  },
];

const DRILLS = [
  {
    day: "Monday",
    file: "01_arrays_reflex.js",
    patterns: "reverse, max index, sum, rotate, prefix",
    functions: [
      "reverseInPlace",
      "indexOfMax",
      "arraySum",
      "rotateRight",
      "runningSum",
    ],
    triggers: [
      "in-place mutate / reverse → two pointers L/R swap",
      "running total → prefix / accumulate",
      "rotate by k → k %= n, then reverse sections (or copy with modulo)",
    ],
    understandWarmup:
      "Rotate right by k: same-length array; last k elements move to the front.",
  },
  {
    day: "Tuesday",
    file: "02_hashing_reflex.js",
    patterns: "two-sum, dup, freq map, unique char, anagrams",
    functions: [
      "twoSum",
      "containsDuplicate",
      "frequencyMap",
      "firstUniqueChar",
      "groupAnagrams",
    ],
    triggers: [
      "pair sums to target → Map + complement",
      "seen before? → Set",
      "same letter multiset → sorted key or count signature",
    ],
    understandWarmup:
      "Two-sum: return indices of two numbers that add to target (usually one valid pair).",
  },
  {
    day: "Wednesday",
    file: "03_two_pointers_reflex.js",
    patterns: "dedupe, zeroes, container, palindrome, window",
    functions: [
      "removeDuplicates",
      "moveZeroes",
      "maxArea",
      "isPalindrome",
      "maxSumSubarrayK",
    ],
    triggers: [
      "sorted + two values → L/R two pointers",
      "in-place filter / dedupe → read/write pointers",
      "subarray length k → fixed sliding window",
    ],
    understandWarmup:
      "Container with most water: max area between two lines = width × min(height).",
  },
  {
    day: "Thursday",
    file: "04_binary_search_reflex.js",
    patterns: "exact BS, lower bound, rotated min, present?",
    functions: [
      "binarySearch",
      "searchInsert",
      "findMinRotated",
      "isTargetPresent",
    ],
    triggers: [
      "sorted + find exact → lo <= hi, mid compare",
      "first position ≥ target → lower bound / searchInsert",
      "rotated sorted min → decide which half is sorted",
    ],
    understandWarmup:
      "Lower bound: first index where value is ≥ target (insertion point).",
  },
  {
    day: "Friday",
    file: "05_trees_stacks_reflex.js",
    patterns: "inorder, depth, parens, mono stack",
    functions: [
      "inorderTraversal",
      "maxDepth",
      "isValidParentheses",
      "dailyTemperatures",
    ],
    triggers: [
      "tree order without recursion → stack / iterative inorder",
      "matching brackets → stack of opens",
      "next greater element → monotonic decreasing stack",
    ],
    understandWarmup:
      "Valid parentheses: every closer must match the latest unmatched opener.",
  },
  {
    day: "Saturday",
    file: "06_dp_reflex.js",
    patterns: "fib, min cost, rob, climb stairs",
    functions: ["fib", "minCostClimbingStairs", "rob", "climbStairs"],
    triggers: [
      "min cost / max ways on a line → 1D DP",
      "define dp[i] in English before coding",
      "rob houses → cannot take adjacent → max(take, skip)",
    ],
    understandWarmup:
      "Min cost climbing: from i you pay cost[i], then jump 1 or 2 steps; reach top with min total.",
  },
  {
    day: "Sunday",
    file: "07_graphs_reflex.js",
    patterns: "islands, flood fill, BFS path",
    functions: ["numIslands", "floodFill", "shortestPathGrid"],
    triggers: [
      "grid regions / components → DFS or BFS + visited",
      "shortest path unweighted grid → BFS",
      "flood fill → DFS/BFS from start, recolor connected cells",
    ],
    understandWarmup:
      "Island count: each unvisited land cell starts one DFS/BFS component.",
  },
];

const ESSENTIAL_CATALOG = {
  "Arrays & prefix": [
    "reverseInPlace",
    "indexOfMax",
    "arraySum",
    "rotateRight",
    "runningSum",
  ],
  Hashing: [
    "twoSum",
    "containsDuplicate",
    "frequencyMap",
    "firstUniqueChar",
    "groupAnagrams",
  ],
  "Two pointers & window": [
    "removeDuplicates",
    "moveZeroes",
    "maxArea",
    "isPalindrome",
    "maxSumSubarrayK",
  ],
  "Binary search": [
    "binarySearch",
    "searchInsert",
    "findMinRotated",
    "isTargetPresent",
  ],
  "Trees & stacks": [
    "inorderTraversal",
    "maxDepth",
    "isValidParentheses",
    "dailyTemperatures",
  ],
  DP: ["fib", "climbStairs", "minCostClimbingStairs", "rob"],
  Graphs: ["numIslands", "floodFill", "shortestPathGrid"],
};

const ALL_TRIGGERS = [
  "pair sums to target → Map + complement",
  "duplicates / seen before → Set or freq map",
  "sorted + two values / area → L/R two pointers",
  "in-place filter / dedupe → read/write pointers",
  "subarray size k → fixed sliding window",
  "sorted + find / insert → binary search or lower bound",
  "matching brackets → stack",
  "next greater → monotonic stack",
  "min cost / ways → 1D DP (define dp[i] first)",
  "grid regions / fill → DFS or BFS + visited",
  "shortest unweighted path → BFS",
];

const dayIndex = new Date().getDay(); // 0 = Sun
const drillIndex = dayIndex === 0 ? 6 : dayIndex - 1;
const today = DRILLS[drillIndex];
const drillPath = path.join(__dirname, "drills", today.file);
const runTests = process.argv.includes("--run");
const microOnly = process.argv.includes("--micro");
const showCatalog = process.argv.includes("--catalog");

function printBanner() {
  console.log(`
╔══════════════════════════════════════════════════════════╗
║         DAILY REFLEX PRACTICE — ESSENTIAL PACK           ║
╚══════════════════════════════════════════════════════════╝
`);
}

function printCore5() {
  console.log(`── CORE 5 (every day — target < 8 min) ───────────────
  Say the ask, then write blind.
`);
  CORE5.forEach((fn, i) => {
    console.log(`  ${i + 1}. ${fn.name}`);
    console.log(`     Ask:     ${fn.ask}`);
    console.log(`     Pattern: ${fn.pattern}`);
    console.log(`     Target:  ${fn.sec}s\n`);
  });
}

function printCatalog() {
  printBanner();
  console.log(`Essential catalog — own every function blind.\n`);
  for (const [group, fns] of Object.entries(ESSENTIAL_CATALOG)) {
    console.log(`  ${group}`);
    fns.forEach((fn) => console.log(`    [ ] ${fn}`));
    console.log();
  }
  console.log(`Guide: study_play/DAILY_30MIN_DRILL.md
`);
}

function printMicro() {
  printBanner();
  console.log(`Minimum / micro day (${today.day}) — Core 5 only.
`);
  printCore5();
  console.log(`When ready for specialty:
  Open: study_play/drills/${today.file}
  Run:  node study_play/daily_drill.js --run
`);
}

function printToday() {
  printBanner();
  console.log(`Today: ${today.day}
Specialty file: study_play/drills/${today.file}
Specialty set:  ${today.patterns}
`);

  if (today.day === "Sunday") {
    console.log(
      `Note: Sunday is rest from new problems. Graphs specialty is optional; still do Core 5.\n`
    );
  }

  console.log(`── TIERS ──────────────────────────────────────────────
  Minimum   ~20-30 min   Core 5 + log
  Reflex    ~30-40 min   Core 5 + today's specialty drill
  Standard  45-60 min    Reflex + ONE primary from STUDY_PLAN.md
`);

  printCore5();

  console.log(`── SPECIALTY FUNCTIONS (blind write after Core 5) ──`);
  today.functions.forEach((fn) => console.log(`  • ${fn}`));

  console.log(`
── REFLEX CLOCK ────────────────────────────────────────
  0-2 min    Full trigger scan (out loud)
  2-10 min   Core 5 blind write
  10-12 min  Understand warm-up (specialty)
  12-32 min  Specialty TODO: REFLEX functions
  32-37 min  Run tests & fix once (no solutions)
  37-40 min  Log fails + revisit (+3 days)

── TODAY'S TRIGGERS ────────────────────────────────────`);
  today.triggers.forEach((t) => console.log(`  • ${t}`));

  console.log(`
── ALWAYS-ON TRIGGERS (scan daily) ─────────────────────`);
  ALL_TRIGGERS.forEach((t) => console.log(`  • ${t}`));

  console.log(`
── UNDERSTAND WARM-UP (say aloud) ──────────────────────
  ${today.understandWarmup}

── COMMANDS ────────────────────────────────────────────
  Open specialty:  study_play/drills/${today.file}
  Run specialty:   node study_play/daily_drill.js --run
  Core 5 only:     node study_play/daily_drill.js --micro
  Full catalog:    node study_play/daily_drill.js --catalog
  Full guide:      study_play/DAILY_30MIN_DRILL.md
`);
}

if (showCatalog) {
  printCatalog();
  process.exit(0);
}

if (microOnly) {
  printMicro();
  process.exit(0);
}

printToday();

if (runTests) {
  console.log("Running specialty tests...\n");
  try {
    execSync(`node "${drillPath}"`, { stdio: "inherit" });
  } catch {
    console.log("\nTests failed — good data. Fix blind, then re-run.");
    process.exit(1);
  }
} else {
  console.log(
    "Tip: after Core 5 + specialty, run  node study_play/daily_drill.js --run\n"
  );
}
