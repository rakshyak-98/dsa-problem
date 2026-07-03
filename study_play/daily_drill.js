/**
 * Daily 30-min reflex drill helper
 *
 * RUN: node study_play/daily_drill.js
 * RUN with tests: node study_play/daily_drill.js --run
 */

const { execSync } = require("child_process");
const path = require("path");

const DRILLS = [
  { day: "Monday", file: "01_arrays_reflex.js", patterns: "reverse, max, sum, rotate, prefix" },
  { day: "Tuesday", file: "02_hashing_reflex.js", patterns: "two-sum, dup, freq map, anagrams" },
  { day: "Wednesday", file: "03_two_pointers_reflex.js", patterns: "dedupe, zeroes, container, window" },
  { day: "Thursday", file: "04_binary_search_reflex.js", patterns: "exact BS, lower bound, rotated min" },
  { day: "Friday", file: "05_trees_stacks_reflex.js", patterns: "inorder, depth, parens, mono stack" },
  { day: "Saturday", file: "06_dp_reflex.js", patterns: "fib, min cost, rob, climb stairs" },
  { day: "Sunday", file: "07_graphs_reflex.js", patterns: "islands, flood fill, BFS path" },
];

const dayIndex = new Date().getDay(); // 0 = Sun
const drillIndex = dayIndex === 0 ? 6 : dayIndex - 1;
const today = DRILLS[drillIndex];
const drillPath = path.join(__dirname, "drills", today.file);
const runTests = process.argv.includes("--run");

console.log(`
╔══════════════════════════════════════════════════════════╗
║           DAILY 30-MIN REFLEX DRILL                      ║
╚══════════════════════════════════════════════════════════╝

Today: ${today.day}
File:  study_play/drills/${today.file}
Drill: ${today.patterns}

── 30-MINUTE CLOCK ──────────────────────────────────────
  0-2 min   Trigger scan (read DAILY_30MIN_DRILL.md triggers)
  2-22 min  Blind write all TODO: REFLEX functions
  22-27 min Run tests & fix (no solutions yet)
  27-30 min Log what you forgot + revisit date

── COMMANDS ─────────────────────────────────────────────
  Open drill:  study_play/drills/${today.file}
  Run tests:   node study_play/drills/${today.file}
`);

if (runTests) {
  console.log("Running tests...\n");
  try {
    execSync(`node "${drillPath}"`, { stdio: "inherit" });
  } catch {
    console.log("\nTests failed — good data. Fix blind, then re-run.");
    process.exit(1);
  }
} else {
  console.log("Tip: after implementing, run  node study_play/daily_drill.js --run\n");
}
