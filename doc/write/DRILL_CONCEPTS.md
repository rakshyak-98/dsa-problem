# Drill Concepts Guide

This file explains the concepts behind each reflex drill in `drills/write/reflex/` (source: `drills/write/reflex/`).

## Drill 01: Arrays

Functions: `reverseInPlace`, `rotateRight`, `runningSum`, `subarraySumK`, `productExceptSelf`, `maxSubarraySum`, `mergeSort`, `quickSort`, `sieve`

- **Core idea:** Precompute along the array so each query is O(1).
- **When this pattern appears:** Range sums, "count the subarrays that…", best contiguous run.
- **Key concepts:**
  - **Two-pointer swap:** Use `left` and `right` for in-place reverse.
  - **Rotation logic:** Reduce `k` with modulo and transform indices safely.
  - **Prefix array:** `sum(i..j) = pre[j] - pre[i-1]` — the foundation of the rest.
  - **Prefix + map:** A prefix seen before means the gap between the two sums to k.
    This turns an O(n²) subarray scan into one pass.
  - **Two-directional prefix:** Product-except-self needs everything left *and*
    right of i, so it is a prefix pass followed by a suffix pass.
  - **Kadane:** `cur = max(x, cur+x)` — extend the run or restart at x.
- **Also drilled here — the array algorithms you hand-write from scratch:**
  - **Merge sort:** Split in half, recurse both, merge with two read pointers.
    `<=` in the merge keeps equal keys stable (and lets the merge step count
    inversions). Returning the input slice on the base case aliases the caller.
  - **Quick sort:** Lomuto partition — pivot last, `i` tracks the `< pivot`
    frontier, swap pivot home, recurse `lo..i-1` and `i+1..hi`. O(n log n)
    average, O(n²) when the pivot is always an extreme.
  - **Sieve of Eratosthenes:** Cross out multiples of each prime from `p*p`;
    survivors are prime. O(n log log n).

## Drill 02: Hashing

Functions: `twoSum`, `containsDuplicate`, `frequencyMap`, `firstUniqueChar`, `groupAnagrams`, `singleNumber`

- **Core idea:** Trade memory for fast lookup.
- **When this pattern appears:** Need membership checks, counting, grouping by signature.
- **Key concepts:**
  - **Complement lookup:** For target pair questions, check map before inserting.
  - **Seen set:** Detect duplicates in O(1) average lookup per element.
  - **Frequency table:** Count occurrences with `map[key]++`.
  - **Canonical key:** Group equivalent strings (anagrams) by a normalized form.
  - **XOR instead of a set:** `singleNumber` finds the one unpaired value by
    XORing the whole array — equal values cancel to 0. O(1) space, the no-hash
    counterpart to a seen-set. (Rest appear 3× → bit-count mod 3, not plain XOR.)

## Drill 03: Two Pointers and Sliding Window

Functions: `removeDuplicates`, `moveZeroes`, `maxArea`, `isPalindrome`, `maxSumSubarrayK`, `longestUniqueSubstring`, `dutchFlag`

- **Core idea:** Keep constraints with pointer movement instead of nested loops.
- **When this pattern appears:** Sorted arrays, in-place filtering, contiguous window metrics.
- **Key concepts:**
  - **Read/write pointers:** Compact valid values to the front in one pass.
  - **Opposite-end pointers:** Move boundaries based on rule (area/palindrome logic).
  - **Fixed window:** Build first window, then slide by add-right/remove-left.
  - **Variable window:** Right always advances; left moves only to restore the
    property. "Subarray of size k" is fixed; "longest such that…" is variable.
  - **Invariant thinking:** Each pointer move must preserve correctness condition.
  - **Three-way partition (Dutch national flag):** `low`/`mid`/`high` sort an
    array of 0/1/2 in one pass, O(1) space. After a swap with `high`, do **not**
    advance `mid` — the value pulled in is unexamined.

## Drill 04: Binary Search

Functions: `binarySearch`, `searchInsert`, `findMinRotated`, `minEatingSpeed`, `quickSelect`, `fastPow`, `gcd`

- **Core idea:** Use sorted structure to discard half each step.
- **When this pattern appears:** Exact search, insertion point, rotated sorted arrays.
- **Key concepts:**
  - **Search space boundaries:** `lo`, `hi`, and loop condition define correctness.
  - **Midpoint safety:** `mid := lo + (hi-lo)/2`.
  - **Lower bound:** Find first index where condition becomes true.
  - **Half-sorted decision:** In rotated arrays, detect which half is ordered.
  - **Searching the answer, not the array:** When the input is unsorted but
    "does x work?" is monotonic, binary search over x. This is the leap that
    turns binary search from a lookup into a general technique.
- **Also drilled here — same "discard a half / halve each step" idea:**
  - **Quickselect:** Partition, then recurse only into the side holding rank `k`.
    Average O(n); the answer to "kth largest without a full sort".
  - **Binary exponentiation (`fastPow`):** Fold `n` bit by bit — square the base
    each step, multiply it in on a set bit. Negative `n` → invert the base first.
    Same shape works for `x^n mod m` and matrix power.
  - **Euclid's gcd:** `for b != 0 { a, b = b, a%b }`. `lcm(a,b) = a/gcd*b`.

## Drill 05: Trees and Stacks

Functions: `inorderTraversal`, `preorderTraversal`, `postorderTraversal`, `levelOrderTraversal`, `maxDepth`, `isValidBST`, `isValidParentheses`, `dailyTemperatures`

- **Core idea:** Use stack discipline for nested structure and next-greater relations.
- **When this pattern appears:** Tree traversal, bracket matching, nearest next condition.
- **Key concepts:**
  - **Traversal state:** Recursion or explicit stack to simulate call stack.
  - **Order picks the use:** inorder for sorted BST output, preorder for copy/serialize, postorder when children must resolve first, level-order (BFS queue) for row-by-row.
  - **Depth recurrence:** Tree answers often combine left and right subtree results.
  - **LIFO matching:** Parentheses validation uses most recent unmatched opener.
  - **Monotonic stack:** Maintain decreasing/increasing stack for next-greater problems.
  - **Ancestor bounds:** BST validity needs `(lo, hi)` carried down the recursion.
    Comparing a node only with its parent is the classic wrong answer.

## Drill 06: Dynamic Programming (1D)

Functions: `climbStairs`, `minCostClimbingStairs`, `rob`, `coinChange`

- **Core idea:** Build answers from solved smaller subproblems.
- **When this pattern appears:** "Min cost", "max value", "count ways" on linear states.
- **Key concepts:**
  - **State definition first:** Write what `dp[i]` means before coding.
  - **Base cases:** Initialize smallest known states correctly.
  - **Transition rule:** Express current answer from prior states.
  - **Take/skip pattern:** For adjacency constraints, compare include vs exclude.
  - **Unbounded choice:** When an item may be reused, loop amounts outward and
    try every item at each one. Greedy is wrong: coins {1,3,4}, amount 6.

## Drill 07: Graphs (Grids and Dependencies)

Functions: `numIslands`, `floodFill`, `shortestPathGrid`, `canFinish`, `dfs`, `bfs`, `bfsShortestPath`, `topoSort`, `dijkstra`

- **Core idea:** Explore connected cells with traversal rules.
- **When this pattern appears:** Components in grids, fill regions, shortest unweighted path.
- **Key concepts:**
  - **DFS/BFS component scan:** Start traversal from each unvisited valid cell.
  - **Visited control:** Prevent revisits and infinite loops.
  - **Neighbor generation:** Standard 4-direction moves with boundary checks.
  - **BFS distance layers:** First time reaching a node gives shortest path in unweighted graphs.
  - **Adjacency list + in-degrees:** Not every graph is a grid. Prerequisites,
    ordering, and "is there a cycle" all mean topological sort (Kahn's algorithm).
- **Also drilled here — the same algorithms on a plain adjacency list (`graph[u]` = neighbours):**
  - **`dfs`:** Recursive preorder — mark on entry, append, recurse into unseen
    neighbours. Marking *after* the recursion revisits nodes on a cycle.
  - **`bfs`:** Queue, marking each node **on enqueue** — mark on dequeue and a
    node enters the queue twice.
  - **`bfsShortestPath`:** Level-by-level BFS; the count of levels crossed is the
    edge count. Never DFS for a shortest path.
  - **`topoSort` (Kahn):** Queue zero-in-degree nodes (smallest index on ties),
    pop, decrement targets. Fewer than `n` emitted → a cycle → return nil.
  - **`dijkstra`:** Min-heap of `(node, dist)`; pop the closest, relax its edges,
    skip a stale `(v, d)` where `d > dist[v]`. Non-negative weights only.

## Drill 10: Linked Lists

Functions: `reverseList`, `hasCycle`, `middleNode`, `mergeTwoLists`, `removeNthFromEnd`

- **Core idea:** Manipulate pointers when you cannot index and cannot look back.
- **When this pattern appears:** Reversal, loop detection, merging, positional deletes.
- **Key concepts:**
  - **prev/cur/next rewire:** Save `next` before overwriting `cur.Next`, or the
    rest of the list is lost. Return `prev`, which ends on the new head.
  - **Slow and fast pointers:** Two speeds over one list answers "is there a
    cycle" and "where is the middle" in O(1) space.
  - **Dummy head:** A sentinel before the head makes deleting the first node the
    same code as deleting any other — it removes the special case entirely.
  - **Gap pointers:** To find the nth from the end in one pass, run a lead
    pointer n ahead and advance both until it falls off.

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
- Longest/shortest run with a property -> **Sliding window (variable)**
- Count subarrays summing to k -> **Prefix sum + map**
- Min rate/capacity that still works -> **Binary search the answer**
- Prerequisites / ordering / cycle -> **Topological sort**
- Reverse, loop, or nth-from-end on a list -> **Pointer rewiring (Drill 10)**
- Implement a sort, or kth largest in O(n) -> **Merge/quick sort, quickselect (Drills 01 + 04)**
- Array of only 0/1/2, sort in one pass -> **Dutch national flag (Drill 03)**
- gcd/lcm, fast exponent, primes up to n, one unpaired value -> **Euclid / binary exp / sieve / XOR (Drills 04 + 01 + 02)**
- Traverse a graph, reach every node -> **DFS / BFS on an adjacency list (Drill 07)**
- Fewest edges between two nodes -> **BFS shortest path (Drill 07)**
- Directed order or cycle check, weighted shortest path -> **Topological sort / Dijkstra (Drill 07)**

## How to Study This File

1. Read one drill section.
2. Say the trigger out loud.
3. Rewrite one function from memory.
4. Run the matching drill tests.
5. Log one mistake and the corrected invariant.

**Math behind the patterns:** See `MATH_CONCEPTS.md` for complexity formulas, combinatorics, recurrences, and geometry. The runnable number-theory functions (`gcd`, `fastPow`, `sieve`) live in Drills 04 and 01; that page is the wider reference behind them.
