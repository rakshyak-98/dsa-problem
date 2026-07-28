# Daily Reflex Practice — Essential Pack

> **Purpose:** Build automatic DSA reflexes so medium problems don’t stall on basics.
> **Rule:** Blind write only. No `solutions_reference/` until the blind block is done.
> **Helper:** `cd study_play && go run .`

This file is the full daily practice. Same shape every day. Do the **Core 5** every session, then today’s specialty drill.

---

## How every session works

```
1. Core 5        →  always (builds permanent reflexes)
2. Specialty     →  today’s weekday drill file
3. Trigger scan  →  say the pattern table out loud
4. Log           →  what broke + revisit in 3 days
```

| Tier | Time | Do |
|------|------|----|
| **Minimum** | ~20–30 min | Core 5 + log |
| **Reflex** | ~30–40 min | Core 5 + today’s specialty drill |
| **Standard** | 45–60 min | Reflex tier + one primary from `STUDY_PLAN.md` |

Missed a day? Do **not** catch up. Run today’s pack only.

```bash
cd study_play && go run .             # prints Core 5 + today’s specialty
cd study_play && go run . -- --run    # run today’s specialty tests
cd study_play && go run . -- --reset  # reset today’s drill to blank TODOs
cd study_play && go run . -- --micro  # Core 5 only (paper / blank file)
```

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
| **Mon** | `drills/01_arrays_reflex/` | `reverseInPlace`, `indexOfMax`, `arraySum`, `rotateRight`, `runningSum` |
| **Tue** | `drills/02_hashing_reflex/` | `twoSum`, `containsDuplicate`, `frequencyMap`, `firstUniqueChar`, `groupAnagrams` |
| **Wed** | `drills/03_two_pointers_reflex/` | `removeDuplicates`, `moveZeroes`, `maxArea`, `isPalindrome`, `maxSumSubarrayK` |
| **Thu** | `drills/04_binary_search_reflex/` | `binarySearch`, `searchInsert`, `findMinRotated`, `isTargetPresent` |
| **Fri** | `drills/05_trees_stacks_reflex/` | `inorderTraversal`, `maxDepth`, `isValidParentheses`, `dailyTemperatures` |
| **Sat** | `drills/06_dp_reflex/` | `fib`, `minCostClimbingStairs`, `rob`, `climbStairs` |
| **Sun** | `drills/07_graphs_reflex/` | `numIslands`, `floodFill`, `shortestPathGrid` |

```bash
go run ./drills/0X_...
# or
cd study_play && go run . -- --run
```

**Sunday:** optional streak day (graphs). Rest from new problems is fine — still do Core 5 if you want the habit.

**Every 4th Sunday:** re-type `templates/pattern_cheat_sheet.go` from memory instead of graphs (30 min).

---

## Part C — 30–40 min reflex clock

| Min | Block | Action |
|-----|-------|--------|
| 0–2 | **Trigger scan** | Read the full trigger table below out loud |
| 2–10 | **Core 5** | Blind write the five essentials |
| 10–12 | **Understand warm-up** | One sentence ask for today’s hardest specialty fn |
| 12–32 | **Specialty blind write** | All `TODO: REFLEX` in today’s file |
| 32–37 | **Run & fix** | `cd study_play && go run . -- --run` — one fix pass, no solutions |
| 37–40 | **Log** | What failed + revisit date (+3 days) |

Low energy? Stop after Core 5. That still counts as Minimum tier.

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

Keep these as muscle memory. Full versions live in `templates/pattern_cheat_sheet.go`.

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

You **own** a function when all are true:

- [ ] Wrote it blind (no peek)
- [ ] Tests pass (or hand-trace is correct for Core 5)
- [ ] Can state the ask in one sentence
- [ ] Can name time/space in one breath
- [ ] Specialty set finishes under **15 min**; Core 5 under **8 min**

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
| No time | Core 5 only (`--micro`) |

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
- [ ] `reverseInPlace` · `indexOfMax` · `arraySum` · `rotateRight` · `runningSum`

### Hashing
- [ ] `twoSum` · `containsDuplicate` · `frequencyMap` · `firstUniqueChar` · `groupAnagrams`

### Two pointers & window
- [ ] `removeDuplicates` · `moveZeroes` · `maxArea` · `isPalindrome` · `maxSumSubarrayK`

### Binary search
- [ ] `binarySearch` · `searchInsert` (lower bound) · `findMinRotated` · `isTargetPresent`

### Trees & stacks
- [ ] `inorderTraversal` · `maxDepth` · `isValidParentheses` · `dailyTemperatures`

### DP
- [ ] `fib` · `climbStairs` · `minCostClimbingStairs` · `rob`

### Graphs
- [ ] `numIslands` · `floodFill` · `shortestPathGrid`

---

## After reflex (Standard tier only)

One primary problem from `STUDY_PLAN.md`:

1. Restate ask in one sentence  
2. Trace the sample by hand  
3. Name brute force  
4. Pattern scan (Part D)  
5. Code → log ask + pattern + lesson  

Reflex first. Understanding before code. Core 5 every day — that is how reflexes stick.
