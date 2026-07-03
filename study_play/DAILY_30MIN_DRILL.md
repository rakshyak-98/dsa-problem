# Daily 30-Minute Reflex Drill

> **Purpose:** Make core DSA patterns automatic so your brain is free to *think* on medium/hard problems.
> **Rule:** No copy-paste. No `_solutions_reference.js` until the full 30 minutes is done.

---

## The 30-minute clock (every day)

| Min | Block | What to do |
|-----|-------|------------|
| 0–2 | **Trigger scan** | Read today's pattern triggers (table below). Say them out loud once. |
| 2–22 | **Blind write** | Open today's drill file. Implement every `TODO: REFLEX` from **empty memory**. No peeking. |
| 22–27 | **Run & fix** | `node study_play/drills/XX_....js` — fix failures once without solutions. |
| 27–30 | **Log** | Write 1 line: what you forgot + revisit in 3 days. |

```bash
# Quick start — prints today's drill and runs tests
node study_play/daily_drill.js
```

---

## Weekly rotation (repeat forever)

| Day | Drill file | Patterns you're drilling |
|-----|------------|--------------------------|
| **Mon** | `drills/01_arrays_reflex.js` | reverse, max index, sum, rotate, prefix |
| **Tue** | `drills/02_hashing_reflex.js` | two-sum, dup check, freq map, anagrams |
| **Wed** | `drills/03_two_pointers_reflex.js` | dedupe, move zeroes, container, window |
| **Thu** | `drills/04_binary_search_reflex.js` | exact search, lower bound, rotated min |
| **Fri** | `drills/05_trees_stacks_reflex.js` | inorder, depth, parens, monotonic stack |
| **Sat** | `drills/06_dp_reflex.js` | fib, min cost stairs, rob, climb stairs |
| **Sun** | `drills/07_graphs_reflex.js` | islands, flood fill, BFS shortest path |

**Every 4th Sunday:** swap Sun for re-typing `templates/pattern_cheat_sheet.js` from memory (30 min).

---

## Pattern triggers (2-min scan)

Memorize these pairs — they are your reflex shortcuts:

| When you see… | Your hands write… |
|---------------|-------------------|
| pair sums to target | `Map` + complement loop |
| sorted + two values | two pointers L/R |
| in-place filter / dedupe | read/write two pointers |
| subarray size k | fixed sliding window |
| sorted + find position | binary search `lo <= hi` or lower bound |
| tree order / stack | iterative inorder or DFS |
| next greater element | monotonic stack |
| min cost / max ways | 1D DP tabulation |
| grid regions / fill | DFS or BFS + visited |

---

## Reflex pass criteria

You **own** a drill when:

- [ ] All tests pass
- [ ] You wrote every function without opening solutions
- [ ] Total blind-write time under **15 minutes** (goal; ok if slower at first)

When you pass all 7 drills under 15 min each → add **Speed Round** (below).

---

## Speed Round (after basics are automatic)

Same 30 min, harder rules:

1. **Delete** your implementations at the top of the drill file (keep tests).
2. Re-implement all functions in **10 minutes**.
3. Run tests. Any fail → that function gets drilled **tomorrow** again (swap with next day's file).

**Micro-drill (2 min each, any day with 5 spare minutes):**

```javascript
// Write on paper or blank file — no tests needed
twoSum(nums, target)           // target: 90 sec
binarySearch(nums, target)     // target: 60 sec
removeDuplicates(nums)         // target: 90 sec
maxArea(heights)               // target: 3 min
```

---

## If you fail a day

| Situation | Do this (still 30 min) |
|-----------|------------------------|
| Stuck 10+ min on one function | Skip it, finish others, peek solution **only for that one**, close file, re-type from memory in last 5 min |
| All tests fail | Re-do **yesterday's** drill file instead (recovery day) |
| No time / low energy | Do **only** `twoSum` + `binarySearch` + `removeDuplicates` blind (10 min) — show up beats perfect |

---

## 30-day reflex tracker

Copy and check off:

```
Week 1:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 2:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 3:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 4:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
```

Log line format (keep in notes or commit message):

```
2026-06-30 | 02_hashing | forgot groupAnagrams key sort | revisit Jul 3
```

---

## How this connects to the full plan

- **30 min/day** = reflex (this file) — non-negotiable baseline
- **+15–45 min** optional = one repo problem from `STUDY_PLAN.md` when you have energy

Reflex first. Problem-solving second. Both together make medium questions readable.
