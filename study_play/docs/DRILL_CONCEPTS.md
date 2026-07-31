# Drill Concepts Guide

This file explains the concepts behind each reflex drill in `drills/write/reflex/` (source: `study_play/practice/write/reflex/`).

## Drill 01: Arrays

Functions: `reverseInPlace`, `indexOfMax`, `arraySum`, `rotateRight`, `runningSum`

- **Core idea:** Traverse arrays with clear index control.
- **When this pattern appears:** In-place mutation, basic scanning, prefix accumulation.
- **Key concepts:**
  - **Two-pointer swap:** Use `left` and `right` for in-place reverse.
  - **Single pass selection:** Track best index/value while scanning.
  - **Accumulation:** Maintain running total for sums and prefix output.
  - **Rotation logic:** Reduce `k` with modulo and transform indices safely.

## Drill 02: Hashing

Functions: `twoSum`, `containsDuplicate`, `frequencyMap`, `firstUniqueChar`, `groupAnagrams`

- **Core idea:** Trade memory for fast lookup.
- **When this pattern appears:** Need membership checks, counting, grouping by signature.
- **Key concepts:**
  - **Complement lookup:** For target pair questions, check map before inserting.
  - **Seen set:** Detect duplicates in O(1) average lookup per element.
  - **Frequency table:** Count occurrences with `map[key]++`.
  - **Canonical key:** Group equivalent strings (anagrams) by a normalized form.

## Drill 03: Two Pointers and Sliding Window

Functions: `removeDuplicates`, `moveZeroes`, `maxArea`, `isPalindrome`, `maxSumSubarrayK`

- **Core idea:** Keep constraints with pointer movement instead of nested loops.
- **When this pattern appears:** Sorted arrays, in-place filtering, contiguous window metrics.
- **Key concepts:**
  - **Read/write pointers:** Compact valid values to the front in one pass.
  - **Opposite-end pointers:** Move boundaries based on rule (area/palindrome logic).
  - **Fixed window:** Build first window, then slide by add-right/remove-left.
  - **Invariant thinking:** Each pointer move must preserve correctness condition.

## Drill 04: Binary Search

Functions: `binarySearch`, `searchInsert`, `findMinRotated`, `isTargetPresent`

- **Core idea:** Use sorted structure to discard half each step.
- **When this pattern appears:** Exact search, insertion point, rotated sorted arrays.
- **Key concepts:**
  - **Search space boundaries:** `lo`, `hi`, and loop condition define correctness.
  - **Midpoint safety:** `mid := lo + (hi-lo)/2`.
  - **Lower bound:** Find first index where condition becomes true.
  - **Half-sorted decision:** In rotated arrays, detect which half is ordered.

## Drill 05: Trees and Stacks

Functions: `inorderTraversal`, `maxDepth`, `isValidParentheses`, `dailyTemperatures`

- **Core idea:** Use stack discipline for nested structure and next-greater relations.
- **When this pattern appears:** Tree traversal, bracket matching, nearest next condition.
- **Key concepts:**
  - **Traversal state:** Recursion or explicit stack to simulate call stack.
  - **Depth recurrence:** Tree answers often combine left and right subtree results.
  - **LIFO matching:** Parentheses validation uses most recent unmatched opener.
  - **Monotonic stack:** Maintain decreasing/increasing stack for next-greater problems.

## Drill 06: Dynamic Programming (1D)

Functions: `fib`, `minCostClimbingStairs`, `rob`, `climbStairs`

- **Core idea:** Build answers from solved smaller subproblems.
- **When this pattern appears:** "Min cost", "max value", "count ways" on linear states.
- **Key concepts:**
  - **State definition first:** Write what `dp[i]` means before coding.
  - **Base cases:** Initialize smallest known states correctly.
  - **Transition rule:** Express current answer from prior states.
  - **Take/skip pattern:** For adjacency constraints, compare include vs exclude.

## Drill 07: Graphs (Grid BFS/DFS)

Functions: `numIslands`, `floodFill`, `shortestPathGrid`

- **Core idea:** Explore connected cells with traversal rules.
- **When this pattern appears:** Components in grids, fill regions, shortest unweighted path.
- **Key concepts:**
  - **DFS/BFS component scan:** Start traversal from each unvisited valid cell.
  - **Visited control:** Prevent revisits and infinite loops.
  - **Neighbor generation:** Standard 4-direction moves with boundary checks.
  - **BFS distance layers:** First time reaching a node gives shortest path in unweighted graphs.

## Fast Pattern Picker

Use this quick mapping when solving:

- Pair sum or quick lookup -> **Hash map**
- Sorted + index/position -> **Binary search**
- In-place compact/reorder -> **Two pointers**
- Contiguous fixed-size segment -> **Sliding window (fixed)**
- Count/group occurrences -> **Frequency map**
- Next greater element -> **Monotonic stack**
- Min/max ways over sequence -> **1D DP**
- Grid region count/fill -> **DFS/BFS**
- Shortest path in unweighted grid -> **BFS**

## How to Study This File

1. Read one drill section.
2. Say the trigger out loud.
3. Rewrite one function from memory.
4. Run the matching drill tests.
5. Log one mistake and the corrected invariant.
