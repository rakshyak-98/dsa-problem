# Daily Reflex Practice — Essential Pack

> **Purpose:** Build automatic DSA reflexes so medium problems don’t stall on basics.
> **Rule:** Blind write only. No `drills/solutions/` until the blind block is done.
> **Helper:** `go run . -- --refresh`

This file is the full daily practice. The shape is the same every day; what it
asks you is not — the session is rebuilt each morning from what your drill log
says decayed.

---

## How every session works

```
1. Recognise    →  4 cues, no code. Which move does this want?
2. Rebuild      →  the weakest functions in your log, not today's weekday
3. New ground   →  something never attempted, highest-value family first
4. Specialty    →  today's weekday drill
5. Stretch      →  a move you own, applied to a variant you have not seen
```

One command builds that session from `.drill_log.json`:

```bash
go run . -- --refresh          # today's session
go run . -- --refresh --show   # reveal the recognition answers
```

Only step 4 is chosen by the weekday. Steps 1–3 and 5 are chosen by what you can
actually do today, which is why the session changes even when the day does not.

| Tier | Time | Do |
|------|------|----|
| **Minimum** | ~20–30 min | Recognise + Core 5 + log |
| **Reflex** | ~30–40 min | Recognise + Core 5 + Rebuild + today's specialty |
| **Standard** | 45–60 min | Reflex tier + the primary problem for one rebuilt function |

Missed a day? Do **not** catch up. Run `--refresh` — it already knows what
decayed.

---

## Understanding levels

Recall and recognition are different skills, and the second is the one an
interview tests. A function you can type from memory is not one you can *spot*
inside a statement that never names it. So every function carries a level, and
the session asks it in the matching form.

| Level | Earned when | Asked as |
|-------|-------------|----------|
| **L1 recall** | never passed, or failing often | "here is the ask — write the shape" |
| **L2 pattern** | ≥2 passes, under a third failing | "here is a cue — name the move" |
| **L3 transfer** | ≥5 passes, near-clean, passed recently | "here is a twist — apply the move" |

Levels are computed from your drill log, never declared:

```bash
go run . -- --levels     # every function, grouped by the level it has earned
go run . -- --problems   # the problem to solve next, ordered by that level
```

Two rules keep the grade honest:

- **Capped by tier.** `reverseInPlace` is a tier-1 function. Passing it fifty
  times leaves it at L1 — it is not a transfer question and never becomes one.
- **Decays.** An L2 function untouched for 3 days, or an L3 for 7, comes back
  as due. `→` in `--levels` marks it.

---

## The recognition round (why it is first)

```
1) you see: sorted but rotated — the pivot is the answer
   you write: ______________________________

2) you see: 'first non-repeating' anything
   you write: ______________________________
```

The cues never name the function, and the four are drawn from four *different*
pattern families on purpose. A themed round ("today is binary search day") lets
you answer from the heading instead of the statement, which trains nothing. Two
minutes, out loud, before any code.

Full table any time: `go run . -- --triggers`.

---

## Part A — Core 5 (every single day)

These five create the highest-leverage reflexes. Say the **ask** in one sentence, then write blind. Target: all five under **8 minutes**.

| # | Function | Ask (say aloud) | Target |
|---|----------|-----------------|--------|
| 1 | `twoSum(nums, target)` | Return indices of two values that sum to target | 90s |
| 2 | `binarySearch(nums, target)` | Return index of target in sorted array, or -1 | 60s |
| 3 | `removeDuplicates(nums)` | In-place unique prefix length on sorted array | 90s |
| 4 | `maxSumSubarrayK(nums, k)` | Max sum of any contiguous window of size k | 90s |
| 5 | `frequencyMap(arr)` / `buildFreq` | Map each value → count | 60s |

**Core 5 checklist (mark daily)**

- [ ] Said each ask before coding
- [ ] Wrote all five from memory
- [ ] Named complexity for each (usually O(n) or O(log n))
- [ ] Noted any fail for +1 day revisit

Blank skeletons (cover answers; re-type into a scratch file):

```go
// 1. HASH — complement lookup
func twoSum(nums []int, target int) []int {
  // map value -> index; need = target - nums[i]
}

// 2. BINARY SEARCH — exact
func binarySearch(nums []int, target int) int {
  // lo = 0, hi = n-1; for lo <= hi; mid; move lo/hi
}

// 3. TWO POINTERS — read/write
func removeDuplicates(nums []int) int {
  // write pointer; keep when nums[read] != nums[write-1]
}

// 4. SLIDING WINDOW — fixed k
func maxSumSubarrayK(nums []int, k int) int {
  // first window sum; slide: +nums[i] - nums[i-k]
}

// 5. FREQ MAP
func frequencyMap(arr []string) map[string]int {
  // m[x]++
}
```

---

## Part B — Specialty drill (weekday rotation)

After Core 5, open today’s file and implement every `TODO: REFLEX` from empty memory.

| Day | File | Essential functions you must own |
|-----|------|----------------------------------|
| **Mon** | `drills/write/reflex/01_arrays_reflex/` | `reverseInPlace`, `rotateRight`, `runningSum`, `subarraySumK`, `productExceptSelf`, `maxSubarraySum` |
| **Tue** | `drills/write/reflex/02_hashing_reflex/` | `twoSum`, `containsDuplicate`, `frequencyMap`, `firstUniqueChar`, `groupAnagrams` |
| **Wed** | `drills/write/reflex/03_two_pointers_reflex/` | `removeDuplicates`, `moveZeroes`, `maxArea`, `isPalindrome`, `maxSumSubarrayK`, `longestUniqueSubstring` |
| **Thu** | `drills/write/reflex/04_binary_search_reflex/` | `binarySearch`, `searchInsert`, `findMinRotated`, `minEatingSpeed` |
| **Fri** | `drills/write/reflex/05_trees_stacks_reflex/` | `inorderTraversal`, `preorderTraversal`, `postorderTraversal`, `levelOrderTraversal`, `maxDepth`, `isValidBST`, `isValidParentheses`, `dailyTemperatures` |
| **Sat** | `drills/write/reflex/06_dp_reflex/` | `climbStairs`, `minCostClimbingStairs`, `rob`, `coinChange` |
| **Sun** | `drills/write/reflex/07_graphs_reflex/` | `numIslands`, `floodFill`, `shortestPathGrid`, `canFinish` |

```bash
go run -C drills/write/reflex/0X_... .
# or
go run . -- --run reflex
```

**Sunday:** optional streak day (graphs). Rest from new problems is fine — still do Core 5 if you want the habit.

**Every 4th Sunday:** re-type `_support/templates/pattern_cheat_sheet.go` from memory instead of graphs (30 min).

---

## Part C — 30–40 min reflex clock

| Min | Block | Action |
|-----|-------|--------|
| 0–2 | **Recognise** | `--refresh`, answer the 4 cues out loud, then `--show` |
| 2–10 | **Core 5** | Blind write the five essentials |
| 10–20 | **Rebuild** | The weakest functions the refresh listed — blind, then run |
| 20–32 | **Specialty** | All `TODO: REFLEX` in today's file |
| 32–37 | **Run & fix** | `go run . -- --run reflex` — one fix pass, no solutions |
| 37–40 | **Log** | `--levels`: did anything move up a level? |

Low energy? Stop after Recognise + Core 5. That still counts as Minimum tier,
and the recognition round is the part that compounds.

---

## Part D — Essential pattern triggers (memorize)

Scan these every day until they fire automatically:

| When you see… | Hands write… | Drill that owns it |
|---------------|--------------|--------------------|
| pair sums to target | `Map` + complement | Core + Tue |
| seen before? / duplicates | `Set` or freq map | Tue |
| anagram / same letter counts | sorted key or count[26] | Tue |
| sorted + two values / area / palindrome | two pointers L/R | Wed |
| in-place filter / dedupe / move zeroes | read/write pointers | Wed + Core |
| subarray of size k | fixed sliding window | Wed + Core |
| longest/shortest subarray with property | variable window L expand/shrink | cheat sheet §6 |
| sorted + find index / insert pos | binary search `lo <= hi` or lower bound | Thu + Core |
| rotated sorted min | half-sorted decide which side | Thu |
| range sum / running total | prefix array | Mon |
| matching brackets | stack | Fri |
| next greater / warmer day | monotonic stack | Fri |
| tree order / depth | DFS / BFS / iterative stack | Fri |
| min cost / ways / fib-style | 1D DP — define `dp[i]` first | Sat |
| cannot take adjacent | DP `max(take, skip)` | Sat |
| grid components / fill | DFS or BFS + visited | Sun |
| shortest path unweighted grid | BFS queue | Sun |
| subarray (contiguous) vs subsequence | confirm ask first → window/prefix vs DP | always |

---

## Part E — Essential templates (type from memory weekly)

Keep these as muscle memory. Full versions live in `_support/templates/pattern_cheat_sheet.go`.

```go
// TWO POINTERS — opposite ends
left, right := 0, n-1
for left < right { /* move left++ or right-- */ }

// TWO POINTERS — read/write
write := 0
for read := 0; read < n; read++ {
  if true /* keep */ {
    nums[write] = nums[read]
    write++
  }
}

// SLIDING WINDOW — variable
left, best := 0, 0
for right := 0; right < n; right++ {
  // expand with right
  for false /* invalid */ { /* shrink left++ */ }
  if right-left+1 > best {
    best = right - left + 1
  }
}

// BINARY SEARCH
lo, hi := 0, n-1
for lo <= hi {
  mid := lo + (hi-lo)/2
  // compare nums[mid] with target; move lo/hi
}

// 1D DP
dp := make([]int, n+1)
// dp[0], dp[1] = bases; for i := 2; i <= n; i++ { dp[i] = ... }
```

---

## Part F — Reflex ownership criteria

You **own** a function when it is graded **L3 transfer** in `--levels`, which
takes all of:

- [ ] Wrote it blind (no peek)
- [ ] Tests pass (or hand-trace is correct for Core 5)
- [ ] Can state the ask in one sentence
- [ ] Can name time/space in one breath
- [ ] Named the move from its cue alone, with the function name hidden
- [ ] Specialty set finishes under **15 min**; Core 5 under **8 min**

Tier-1 functions cap at L1 by design. "Owning" them means the recognition round
never catches you out, not that the grade climbs.

**Speed Round** (after you own all 7 specialty files):

1. Delete implementations at top of today’s drill (keep tests).
2. Re-implement all specialty fns in **10 minutes**.
3. Any fail → drill that function again tomorrow (add to Core block).

---

## Part G — Failure recovery (still counts)

| Situation | Do this |
|-----------|---------|
| Stuck 10+ min on one specialty fn | Skip it, finish others; peek **only that one**; re-type blind in last 5 min |
| Core 5 fails | Stop specialty; re-do Core 5 only — protect the foundation |
| All specialty tests fail | Yesterday’s specialty file (recovery day) + Core 5 |
| No time | Core 5 only (`--drill core`) |

---

## Part H — 30-day reflex tracker

Check a box only after Core 5 (Minimum) or Core 5 + specialty (Reflex). Both improve reflexes.

```
Week 1:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 2:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 3:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 4:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
```

Log line:

```
2026-07-22 | core5 OK | 03_two_pointers | forgot maxArea move shorter side | revisit Jul 25
```

---

## Essential catalog (everything you must eventually write blind)

Use this as a master checklist. Specialty days cover these in rotation; Core 5 keeps the spine sharp.

### Arrays & prefix
- [ ] `reverseInPlace` · `rotateRight` · `runningSum` · `subarraySumK` · `productExceptSelf` · `maxSubarraySum`

### Hashing
- [ ] `twoSum` · `containsDuplicate` · `frequencyMap` · `firstUniqueChar` · `groupAnagrams`

### Two pointers & window
- [ ] `removeDuplicates` · `moveZeroes` · `maxArea` · `isPalindrome` · `maxSumSubarrayK` · `longestUniqueSubstring`

### Binary search
- [ ] `binarySearch` · `searchInsert` (lower bound) · `findMinRotated` · `minEatingSpeed` (search the answer)

### Trees & stacks
- [ ] `inorderTraversal` · `preorderTraversal` · `postorderTraversal` · `levelOrderTraversal` · `maxDepth` · `isValidBST` · `isValidParentheses` · `dailyTemperatures`

### DP
- [ ] `climbStairs` · `minCostClimbingStairs` · `rob` · `coinChange`

### Graphs
- [ ] `numIslands` · `floodFill` · `shortestPathGrid` · `canFinish`

### Linked lists (bonus)
- [ ] `reverseList` · `hasCycle` · `middleNode` · `mergeTwoLists` · `removeNthFromEnd`

---

## After reflex (Standard tier only)

One primary problem from `STUDY_PLAN.md`:

1. Restate ask in one sentence  
2. Trace the sample by hand  
3. Name brute force  
4. Pattern scan (Part D)  
5. Code → log ask + pattern + lesson  

Reflex first. Understanding before code. Core 5 every day — that is how reflexes stick.
