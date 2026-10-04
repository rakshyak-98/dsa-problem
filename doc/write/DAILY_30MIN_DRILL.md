# Daily Reflex Practice — Essential Pack

> **Purpose:** Build automatic DSA reflexes so medium problems don’t stall on basics.
> **Rule:** Blind write only. No `drills/solutions/` until the blind block is done.
> **Helper:** `go run .`

This file is the full daily practice: today's weekday specialty drill.

```bash
go run .                 # today's plan
go run . -- --weak       # the weakest functions from your drill log
```

| Tier | Time | Do |
|------|------|----|
| **Minimum** | ~20–30 min | Half of today's specialty + log |
| **Reflex** | ~30–40 min | Today's specialty |
| **Standard** | 45–60 min | Reflex tier + the primary problem for one function (`--problems`) |

Missed a day? Do **not** catch up. Run today's plan; `--weak` shows what to redo.

---

## Part A — Specialty drill (weekday rotation)

Open today’s file and implement every `TODO: REFLEX` from empty memory.

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

**Sunday:** optional streak day (graphs). Rest from new problems is fine.

**Every 4th Sunday:** re-type `_support/templates/pattern_cheat_sheet.go` from memory instead of graphs (30 min).

---

## Part B — 30–40 min reflex clock

| Min | Block | Action |
|-----|-------|--------|
| 0–2 | **Warm-up** | Read today's warm-up ask out loud |
| 2–5 | **Weak** | `--weak`: redo the worst function blind |
| 5–30 | **Specialty** | All `TODO: REFLEX` in today's file |
| 30–37 | **Run & fix** | `go run . -- --run reflex` — one fix pass, no solutions |
| 37–40 | **Log** | Note any fail for a +1 day revisit |

Low energy? Do half of today's specialty. That still counts as Minimum tier.

---

## Part C — Essential templates (type from memory weekly)

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

## Part D — Reflex ownership criteria

You **own** a function when all of these hold:

- [ ] Wrote it blind (no peek)
- [ ] Tests pass
- [ ] Can state the ask in one sentence
- [ ] Can name time/space in one breath
- [ ] Specialty set finishes under **15 min**

**Speed Round** (after you own all 7 specialty files):

1. Delete implementations at top of today’s drill (keep tests).
2. Re-implement all specialty fns in **10 minutes**.
3. Any fail → drill that function again tomorrow (redo it first tomorrow).

---

## Part E — Failure recovery (still counts)

| Situation | Do this |
|-----------|---------|
| Stuck 10+ min on one specialty fn | Skip it, finish others; peek **only that one**; re-type blind in last 5 min |
| All specialty tests fail | Yesterday’s specialty file (recovery day) |
| No time | One specialty function, blind |

---

## Part F — 30-day reflex tracker

Check a box only after a specialty drill (Minimum: half, Reflex: full). Both improve reflexes.

```
Week 1:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 2:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 3:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 4:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
```

Log line:

```
2026-07-22 | reflex OK | 03_two_pointers | forgot maxArea move shorter side | revisit Jul 25
```

---

## Essential catalog (everything you must eventually write blind)

Use this as a master checklist. Specialty days cover these in rotation.

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

---

## After reflex (Standard tier only)

One primary problem from `STUDY_PLAN.md`:

1. Restate ask in one sentence  
2. Trace the sample by hand  
3. Name brute force  
4. Pattern scan  
5. Code → log ask + pattern + lesson  

Reflex first. Understanding before code. Daily reps are how reflexes stick.
