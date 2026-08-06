# DSA Jargon — Plain English Glossary

**Purpose:** When you read a drill, problem, or solution and hit a term you do not know, look it up here. Each entry is one simplified sentence — no proofs, no code.

Use this alongside [`write/DRILL_CONCEPTS.md`](write/DRILL_CONCEPTS.md) (pattern details) and [`read/READING_PATTERNS.md`](read/READING_PATTERNS.md) (how to read code).

---

## Core terms

| Term | Plain English |
|------|---------------|
| **DSA** | Data Structures and Algorithms — how you store data and the step-by-step methods to solve problems with it. |
| **Data structure** | A way to organize data so common operations (look up, insert, delete) are fast. |
| **Algorithm** | A clear, finite set of steps that solves a problem. |
| **Pattern** | A reusable problem shape (like “two pointers” or “hash map”) that many different questions share. |
| **Reflex drill** | A short coding exercise you rewrite from memory until the pattern feels automatic. |
| **Invariant** | A rule that stays true after every step of your loop — if you break it, the answer is wrong. |
| **Brute force** | Try every possibility until you find the answer — correct but usually slow. |
| **Optimal solution** | The best-known approach for speed (and sometimes space) on that problem type. |
| **Editorial** | The official or community explanation of how to solve a problem. |
| **Template** | A skeleton version of a pattern you fill in for each new problem. |

---

## Problem statement words

| Term | Plain English |
|------|---------------|
| **Input** | What the function receives (array, string, number, etc.). |
| **Output** | What you must return (value, index, boolean, count, or a changed array). |
| **In-place** | Change the given array without making a new one — usually saves memory. |
| **Contiguous** | Items sitting next to each other with no gaps in between. |
| **Subarray** | A contiguous slice of an array (e.g. `[2,3]` inside `[1,2,3,4]`). |
| **Subsequence** | Items picked in order but not necessarily next to each other (e.g. `[1,3]` from `[1,2,3]`). |
| **Substring** | A contiguous piece of a string. |
| **Subsequence (string)** | Characters kept in order but with characters in between allowed to be skipped. |
| **Palindrome** | Reads the same forward and backward (`"racecar"`). |
| **Anagram** | Two strings with the same letters in a different order (`"listen"` / `"silent"`). |
| **Sorted** | Arranged in increasing (or decreasing) order. |
| **Duplicate** | The same value appearing more than once. |
| **Target** | The number or value you are trying to match, reach, or sum to. |
| **Index** | Position of an element in an array (usually starting at 0). |
| **Edge case** | Unusual but valid input (empty array, one element, all same values). |
| **Constraint** | A limit on input size or what you are allowed to do (e.g. “no extra array”). |

---

## Data structures

| Term | Plain English |
|------|---------------|
| **Array** | A fixed-order list of items accessed by position (index). |
| **Slice** | A flexible-length view of a sequence (Go’s dynamic array). |
| **String** | A sequence of characters. |
| **Hash map / hash table** | A lookup table that finds a key’s value in roughly constant average time. |
| **Map** | Same idea as a hash map — store key → value pairs. |
| **Set** | A collection that only tracks whether each value is present (no duplicates). |
| **Stack** | Last-in, first-out — like a stack of plates; push on top, pop from top. |
| **Queue** | First-in, first-out — like a line; add at back, remove from front. |
| **Deque (double-ended queue)** | A queue where you can add or remove from both ends. |
| **Heap / priority queue** | A structure that always gives you the smallest (min-heap) or largest (max-heap) item quickly. |
| **Binary tree** | Each node has at most two children: left and right. |
| **Binary search tree (BST)** | A binary tree where left child < node < right child — enables fast search. |
| **Graph** | Nodes (vertices) connected by edges — models relationships and paths. |
| **Adjacency list** | A graph stored as “each node → list of neighbors.” |
| **Adjacency matrix** | A graph stored as a 2D table of which pairs are connected. |
| **Matrix / grid** | A 2D array of rows and columns (common in graph-as-grid problems). |
| **Linked list** | Nodes chained by pointers; each node holds a value and a link to the next. |
| **Trie** | A tree built for fast prefix lookups on strings (like an autocomplete index). |

---

## Complexity & performance

| Term | Plain English |
|------|---------------|
| **Time complexity** | How the number of steps grows as input size `n` grows. |
| **Space complexity** | How much extra memory your solution uses as `n` grows. |
| **Big O / O(…)** | An upper bound — “at most about this slow” in the worst case. |
| **Ω (Omega)** | A lower bound — “at least this slow.” |
| **Θ (Theta)** | A tight bound — grows at exactly that rate (both upper and lower). |
| **O(1)** | Constant time — work does not grow with `n` (e.g. hash lookup average case). |
| **O(log n)** | Logarithmic — each step cuts the problem in half (binary search, balanced tree height). |
| **O(n)** | Linear — one pass through the data. |
| **O(n log n)** | Common for efficient sorting and divide-and-conquer (e.g. merge sort). |
| **O(n²)** | Quadratic — nested loops over `n` (e.g. check every pair). |
| **O(2ⁿ)** | Exponential — choices branch wildly (naive recursion without memo). |
| **Amortized** | Average cost per operation over many operations (e.g. dynamic array push is O(1) amortized). |
| **Worst case** | Slowest possible input for your algorithm. |
| **Average case** | Typical random input behavior. |
| **Recursion depth** | How many nested function calls stack up — counts toward space. |
| **Trade-off** | Using more memory to save time, or vice versa. |

---

## Array patterns (Drill 01)

| Term | Plain English |
|------|---------------|
| **Traversal / scan** | Walk through the array one element at a time. |
| **In-place reverse** | Swap elements from both ends moving inward until the array is backward. |
| **Two-pointer swap** | Use `left` and `right` indices to swap pairs while moving toward the center. |
| **Running sum / prefix sum** | Cumulative total up to each index (`[1,2,3]` → `[1,3,6]`). |
| **Prefix sum** | Sum of all elements from the start up to (and sometimes including) index `i`. |
| **Rotation** | Move elements circularly — last items wrap to the front. |
| **Modulo (`%`)** | Remainder after division; used to wrap indices safely (e.g. `k % n` for rotation). |

---

## Hashing patterns (Drill 02)

| Term | Plain English |
|------|---------------|
| **Hashing** | Turn a key into a table slot for fast lookup. |
| **Hash map complement** | Store values seen so far; check if `target - current` already exists (two-sum trick). |
| **Complement** | The other number needed to reach a target (`target - x`). |
| **Frequency map / frequency table** | Count how many times each key appears. |
| **Membership check** | “Have I seen this value before?” — answered quickly with a set or map. |
| **Canonical key** | A normalized form used to group equivalent items (sorted letters for anagrams). |
| **Collision** | Two different keys landing in the same hash slot — handled by the map implementation. |
| **Rolling hash** | Update a hash in O(1) when a window slides by adding/removing one element. |

---

## Two pointers & sliding window (Drill 03)

| Term | Plain English |
|------|---------------|
| **Two pointers** | Two indices moving through data to enforce a rule without nested loops. |
| **Opposite two pointers** | `left` at start, `right` at end — move inward (palindrome, container area). |
| **Read/write pointers** | `read` scans everything; `write` marks where the next kept value goes (compact in-place). |
| **Sliding window** | A contiguous range `[left…right]` that grows or shrinks to satisfy a condition. |
| **Fixed sliding window** | Window size stays at `k`; slide by adding one right element and removing one left. |
| **Variable sliding window** | Window size changes — expand `right` until invalid, then shrink `left` until valid again. |
| **Window invariant** | The condition your window must satisfy (e.g. “sum ≤ 10”, “all unique chars”). |
| **Kadane’s algorithm** | Track best subarray sum by choosing “extend current run” vs “start fresh here.” |

---

## Binary search (Drill 04)

| Term | Plain English |
|------|---------------|
| **Binary search** | Repeatedly guess the middle of a sorted range and discard the half that cannot contain the answer. |
| **Search space** | The range of indices or values you still consider possible. |
| **Lower bound** | First position where a condition becomes true (first index `≥ target`). |
| **Upper bound** | First position where a condition would stop being true (past the last match). |
| **Search insert position** | Index where `target` would sit if inserted into sorted order. |
| **Rotated sorted array** | A sorted array cut and shifted — one half is still sorted; binary search picks which half. |
| **Half-sorted** | In a rotated array, one side of `mid` is fully sorted — use that to decide where to search. |
| **`lo` / `hi` / `mid`** | Low bound, high bound, and middle index of the current search range. |

---

## Trees & stacks (Drill 05)

| Term | Plain English |
|------|---------------|
| **Tree node** | One item in a tree with a value and links to children. |
| **Root** | The top node of a tree — no parent above it. |
| **Leaf** | A node with no children. |
| **Subtree** | A node plus all descendants below it. |
| **Depth / height** | Longest path from a node down (height) or from root to a node (depth). |
| **Traversal** | Visiting every node in a defined order. |
| **Inorder** | Left subtree → current node → right subtree (gives sorted order in a BST). |
| **Preorder** | Current node → left → right (copy tree shape). |
| **Postorder** | Left → right → current node (delete or evaluate bottom-up). |
| **Level-order** | Visit row by row top to bottom — usually done with a queue (BFS on a tree). |
| **Recursion** | A function that calls itself on smaller subproblems (natural for trees). |
| **Call stack** | The runtime’s stack of active function calls — simulated with an explicit stack when iterating. |
| **LIFO** | Last in, first out — stack discipline. |
| **Valid parentheses** | Every closing bracket matches the most recent unmatched opener — classic stack use. |
| **Monotonic stack** | Stack kept in increasing or decreasing order to answer “next greater/smaller” in one pass. |
| **Next greater element** | For each position, the first warmer/larger value to its right. |

---

## Dynamic programming — DP (Drill 06)

| Term | Plain English |
|------|---------------|
| **Dynamic programming (DP)** | Build the answer from answers to smaller subproblems, reusing work instead of recomputing. |
| **Subproblem** | A smaller version of the main question (e.g. “best up to index `i`”). |
| **State** | What one DP entry represents — always define in plain English first (`dp[i]` = …). |
| **Base case** | Smallest subproblems with known answers (e.g. `n = 0` or `n = 1`). |
| **Transition / recurrence** | Rule to compute a state from previous states. |
| **Tabulation (bottom-up)** | Fill a table from small indices to large using loops. |
| **Memoization (top-down)** | Recurse with a cache so each subproblem is solved once. |
| **Take / skip** | At each step, choose include current item or exclude it (house robber pattern). |
| **1D DP** | States depend on a single index along a line (array, stairs, costs). |
| **2D DP** | States depend on two dimensions (two strings, grid paths). |
| **Overlapping subproblems** | Same subquestion asked many times — signal to use DP. |

---

## Graphs (Drill 07)

| Term | Plain English |
|------|---------------|
| **Vertex / node** | One point in a graph. |
| **Edge** | A connection between two nodes. |
| **Directed graph** | Edges have direction (A → B but not necessarily B → A). |
| **Undirected graph** | Edges work both ways. |
| **Weighted graph** | Each edge has a cost or distance. |
| **Neighbor** | A node directly connected by one edge. |
| **Degree** | Number of edges touching a node. |
| **Connected component** | A group of nodes reachable from each other. |
| **Island** | One connected blob of “land” cells in a grid metaphor. |
| **DFS (depth-first search)** | Explore as far as possible along one path before backtracking. |
| **BFS (breadth-first search)** | Explore layer by layer — queue-based; finds shortest steps in unweighted graphs. |
| **Visited set / visited array** | Marks cells or nodes already explored so you do not loop forever. |
| **Flood fill** | Recolor a connected region starting from one cell (paint-bucket idea). |
| **Shortest path (unweighted)** | Fewest edges between two nodes — BFS gives this on grids. |
| **4-directional** | Move up, down, left, right on a grid. |
| **Cycle** | A path that returns to a node already on the path. |
| **Topological sort** | Order nodes so every edge goes from earlier to later (task dependency order). |

---

## Heaps (Drill 08)

| Term | Plain English |
|------|---------------|
| **Heap** | A complete binary tree stored in an array where parent is always ≤ children (min) or ≥ (max). |
| **Min-heap** | Smallest element is always at the top — good for “kth largest” by keeping only k items. |
| **Max-heap** | Largest element at the top — good for “take two biggest” simulations. |
| **Priority queue** | Abstract idea: always extract the highest- or lowest-priority item — heaps implement this. |
| **Kth largest / kth smallest** | The element that would sit at position k if the array were sorted. |
| **Quickselect** | Partition array like quicksort but only recurse into the half containing k — average O(n). |
| **Merge k sorted lists** | Repeatedly pick the smallest front element among k lists — classic min-heap use. |

---

## Backtracking (Drill 09)

| Term | Plain English |
|------|---------------|
| **Backtracking** | Try a choice, recurse, then undo the choice and try the next option. |
| **Subset** | Any combination of elements (include or exclude each). |
| **Permutation** | Every ordering of all elements. |
| **Combination** | Choose exactly `k` items from `n` without caring about order. |
| **Branch** | One choice path in the search tree (include vs exclude, pick next number). |
| **Prune** | Stop exploring a branch early when it cannot lead to a valid answer. |
| **Path / partial solution** | Current list of choices built so far before backtracking. |
| **State space tree** | Tree of all choices — backtracking walks it without visiting every node blindly. |

---

## Math & bit tricks (Drill 10)

| Term | Plain English |
|------|---------------|
| **GCD (greatest common divisor)** | Largest number that divides both `a` and `b` evenly. |
| **LCM (least common multiple)** | Smallest positive number divisible by both `a` and `b`. |
| **Euclidean algorithm** | Repeatedly replace `(a,b)` with `(b, a mod b)` until `b` is 0 — fast GCD. |
| **Modulo / mod** | Remainder after division; keeps numbers bounded in competitive problems. |
| **Modular exponentiation** | Compute `a^n mod m` fast by squaring and multiplying only needed bits. |
| **nCr / binomial coefficient** | Number of ways to choose `k` items from `n` without order. |
| **Prime** | Integer greater than 1 with no divisors other than 1 and itself. |
| **Power of two** | A number like 1, 2, 4, 8 — exactly one bit set in binary. |
| **Bit manipulation** | Use `&`, `|`, `^`, `<<`, `>>` on binary representations for speed or tricks. |
| **XOR** | Bitwise exclusive or — useful for “cancel pairs” and parity problems. |

---

## Other common pattern names

| Term | Plain English |
|------|---------------|
| **Greedy** | Make the locally best choice at each step hoping it leads to a global optimum. |
| **Divide and conquer** | Split problem in half, solve halves, combine results (merge sort, binary search). |
| **Simulation** | Follow the problem’s rules step by step exactly as described. |
| **Sorting** | Rearrange elements into order — often O(n log n) for comparison sorts. |
| **Counting sort / bucket** | Count occurrences when value range is small — can beat O(n log n). |
| **Union-find (disjoint set)** | Track which elements belong to the same group with fast merge and lookup. |
| **Dijkstra** | Shortest path in a weighted graph with non-negative edges — uses a priority queue. |
| **Topological order** | Linear ordering of DAG nodes respecting all edge directions. |
| **DAG** | Directed acyclic graph — no cycles; good for dependencies. |
| **Bitmask DP** | DP where state is a subset encoded as bits in an integer. |
| **Meet in the middle** | Split input in half, enumerate both halves, combine — tames exponential blow-up. |

---

## Reading-drill vocabulary

| Term | Plain English |
|------|---------------|
| **Signature** | What goes in, what comes out, and whether input is mutated. |
| **Skeleton** | Control flow only — loops, recursion, early returns — before details. |
| **Trace** | Walk one concrete example step by step with variable values. |
| **Bound** | State the time and space complexity after you understand the shape. |
| **Pattern name** | Label the code shape (hash complement, sliding window, monotonic stack, etc.). |
| **Ask** | Restate the problem in one plain sentence a beginner would understand. |

---

## Quick “if you see this word…” index

| You read… | It usually means… |
|-----------|-------------------|
| contiguous | elements must be next to each other |
| in-place | reuse the input array, no new big array |
| complement | the partner value that completes a target sum |
| invariant | the rule your loop must never break |
| visited | mark so you do not process the same cell twice |
| window | a contiguous slice being adjusted with two pointers |
| state (`dp[i]`) | what one table cell represents in English |
| prune | skip a branch that cannot work |
| amortized | cheap on average over many operations |
| unweighted | every edge or step costs the same (BFS applies) |

---

## Related docs in this repo

| Doc | Use when… |
|-----|-----------|
| [`write/DRILL_CONCEPTS.md`](write/DRILL_CONCEPTS.md) | You want triggers and concepts per reflex drill |
| [`write/MATH_CONCEPTS.md`](write/MATH_CONCEPTS.md) | You need formulas, proofs, and deeper math |
| [`read/READING_PATTERNS.md`](read/READING_PATTERNS.md) | You are practicing the 6-pass code-reading method |
| [`write/STUDY_PLAN.md`](write/STUDY_PLAN.md) | You want the weekly schedule and problem-solving flow |
| [`../reference/problems/CATEGORIES.md`](../reference/problems/CATEGORIES.md) | You are browsing problems by topic name |
