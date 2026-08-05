# DSA Study Plan — From Reflex to Medium/Hard Mastery

> **Goal:** Make basic patterns automatic, then train your brain to *understand* a question quickly and *decompose* medium–hard problems without panic.
>
> **Repo:** You already have ~170 `.js` problems organized by topic. This plan tells you *what* to do, *in what order*, and *how* to think.

---

## How to use this folder

| File / folder | Purpose |
|---------------|---------|
| `STUDY_PLAN.md` (this file) | Full roadmap, weekly schedule, mental frameworks |
| `MATH_CONCEPTS.md` | Math formulas & theory for complexity, counting, number theory, graphs, geometry |
| `DRILL_CONCEPTS.md` | Pattern concepts behind each reflex drill |
| `DAILY_30MIN_DRILL.md` | The daily ritual you run every study day |
| `daily_drill.go` | Prints today's drill + understanding prompts |
| `drills/*/` | **Reflex drills** — rewrite from memory until automatic |
| `_support/templates/pattern_cheat_sheet.go` | Blank skeletons; cover the answers and re-type weekly |
| `../arrays/`, `../hashing/`, etc. | Real problems from your repo |

---

## The consistency rule (read this first)

**Show up every day with the same shape.** Do not invent a new session each time.

### Daily ritual (same order every day)

```
1. Understand  →  restate the question in your own words
2. Reflex      →  Core 5 + specialty (`DAILY_30MIN_DRILL.md`)
3. Primary     →  ONE problem from this week's table
4. Log         →  one sentence lesson + revisit date
```

### Energy tiers (pick one; all count as a win)

| Tier | Time | What you do |
|------|------|-------------|
| **Minimum** | ~20–30 min | Core 5 only + 1-line log |
| **Standard** | 45–60 min | Core 5 + specialty drill + **one** primary problem |
| **Stretch** | 75–90 min | Standard + optional second problem **or** re-solve yesterday's |

Rules that protect consistency:

1. **One primary problem per day.** The second listed problem is optional stretch — never required.
2. **Stuck after the time cap?** Stop coding. Write what you understood, peek editorial, then re-type the core idea from memory.
3. **Missed a day?** Do not catch up. Resume today's ritual. Consistency is forward-only.
4. **Sunday:** Rest from new problems. Optional: light reflex (`daily_drill.go`) if you want the streak.

**Default target:** Standard tier, 6 days/week.

---

## Understand the question first (before any code)

Most wrong answers start as a misunderstanding, not a bad algorithm. Spend 3–5 minutes here on every primary problem.

### Question literacy checklist

Write these as comments (or in the tracker) before coding:

```
Ask:        What is this problem asking in one sentence?
Input:      types, sorted?, duplicates?, negatives?
Output:     value / index / boolean / mutated array / count?
Forbidden:  what you must NOT do (extra space, modify input, etc.)
Example:    walk the sample by hand — input → expected
Edge:       empty, size 1, all same, already sorted, max n
```

### Translation test (pass this before Step 3 patterns)

Explain the problem to an imaginary junior teammate in **one sentence** without jargon.

- Pass: "Given numbers, return how many pairs add up to target."
- Fail: jumping straight to "I'll use a hash map."

If you cannot translate it, re-read the prompt — do not pattern-scan yet.

### Common misunderstanding traps

| Trap | What to check |
|------|----------------|
| Asking for **index** vs **value** | Return type in the signature / statement |
| **Subarray** (contiguous) vs **subsequence** | Contiguous ⇒ window / prefix; subsequence ⇒ DP / two pointers carefully |
| **In-place** vs new array | Space constraint in the prompt |
| **Any** valid answer vs **the** unique answer | Multiple answers allowed? |
| Sorted assumption | Is the array guaranteed sorted, or must you sort? |

Full solve process continues in [The problem-solving framework](#the-problem-solving-framework-use-on-every-mediumhard) below.

---

## The 5-phase roadmap (12 weeks)

```
Phase 1  Weeks 1–2   Foundations + reflex basics
Phase 2  Weeks 3–4   Two pointers, sliding window, binary search
Phase 3  Weeks 5–6   Hashing mastery + medium arrays/strings
Phase 4  Weeks 7–9   Trees, graphs, stacks, DP intro
Phase 5  Weeks 10–12 Medium consolidation + hard introduction
```

Each day lists: **Primary** (required on Standard tier) and **Stretch** (optional).

---

## Phase 1 — Foundations (Weeks 1–2)

**Objective:** Arrays, loops, and hashing feel like breathing. You can restate easy prompts in one sentence without rereading five times.

### Week 1 — Arrays & complexity

| Day | Reflex | Primary (understand → solve) | Stretch | Cap |
|-----|--------|------------------------------|---------|-----|
| Mon | `drills/write/reflex/01_arrays_reflex/` | `arrays/easy/plus_one.js` | `arrays/easy/concatenation_of_array.js` | 25 min |
| Tue | Re-drill arrays | `arrays/easy/find_closest_number_to_zero.js` | `arrays/easy/max_consecutive_ones.js` | 25 min |
| Wed | `drills/write/reflex/02_hashing_reflex/` | `hashing/easy/two_sum.js` | `hashing/easy/contains_duplicates.js` | 25 min |
| Thu | Re-drill hashing | `hashing/easy/fair_candy_swap.js` | `hashing/easy/degree_of_an_array.js` | 25 min |
| Fri | Mixed: 1 array + 1 hash drill fn | `arrays/easy/majority_element.js` | `hashing/easy/summary_ranges.js` | 30 min |
| Sat | Review: re-do failed primaries; blind drills | — | — | 60 min |
| Sun | Rest (optional light reflex) | — | — | — |

**Week 1 checklist**
- [ ] Can write `reverseArray`, `maxInArray`, `countFreq` from memory in &lt; 3 min each
- [ ] Can explain O(n) vs O(n²) for every solution you wrote
- [ ] Can restate each solved problem in one plain sentence
- [ ] Solved 6+ primary easies without peeking at old code

### Week 2 — Strings & simulation

| Day | Focus | Primary | Stretch |
|-----|-------|---------|---------|
| Mon | String basics | `strings/easy/find_words_containing_character.js` | `strings/easy/most_common_word.js` |
| Tue | Hashing on strings | `strings/easy/unique_morse_code_words.js` | `hashing/easy/find_resultant_array_after_removing_anagrams.js` |
| Wed | Simulation / step-by-step | `simulation/easy/baseball_game.js` | `simulation/easy/relative_ranks.js` |
| Thu | Mixed easy | 1 unseen from `strings/easy/` | 1 more if energy |
| Fri | **Timed mock:** 2 easy in 45 min | Random from weeks 1–2 | — |
| Sat | Re-drill `01` + `02`; fix weak spots | — | — |
| Sun | Rest | — | — |

**Week 2 checklist**
- [ ] Comfortable with `Map`, `Set`, and object freq maps
- [ ] 12+ total easy primaries solved in this plan
- [ ] For simulation problems: can list the operations in order before coding

---

## Phase 2 — Two pointers & binary search (Weeks 3–4)

**Objective:** See "sorted array", "pair sum", "palindrome" → instantly think two pointers or binary search — *after* you can state what the question wants.

### Week 3 — Two pointers

| Day | Reflex | Primary | Stretch |
|-----|--------|---------|---------|
| Mon | `drills/write/reflex/03_two_pointers_reflex/` | `two_pointers/easy/remove_duplicates_from_sorted_array.js` | `two_pointers/easy/move_zeroes.js` |
| Tue | Re-drill + opposite ends | `two_pointers/easy/best_time_to_buy_sell_stock.js` | `two_pointers/easy/squares_of_a_sorted_array.js` |
| Wed | Medium intro | `two_pointers/medium/container_with_most_water.js` | `two_pointers/medium/3sum.js` (stretch only) |
| Thu | Sliding window in drill | `misc/easy/maximum_average_subarray_1.js` | `two_pointers/easy/minimum_difference_between_highest_and_lowest_of_k_score.js` |
| Fri | Timed: 1 easy **or** 1 medium | Your choice from `two_pointers/` | — |
| Sat | Re-type `_support/templates/pattern_cheat_sheet.go` sections 1–3 | — | — |
| Sun | Rest | — | — |

### Week 4 — Binary search

| Day | Reflex | Primary | Stretch |
|-----|--------|---------|---------|
| Mon | `drills/write/reflex/04_binary_search_reflex/` | `binary_search/easy/search_insertion_position.js` | `binary_search/easy/find_smallest_letter_greater_than_target.js` |
| Tue | Binary search on answer | `binary_search/easy/longest_subsequence_with_limited_sum.js` | — |
| Wed | Mixed pointers + BS | `two_pointers/medium/find_the_duplicate_number.js` | — |
| Thu | Medium | `two_pointers/medium/sort_colors.js` | `two_pointers/medium/longest_palindromic_substring.js` |
| Fri | **Mock:** 1 medium, 45 min, talk out loud | Record steps in comments | — |
| Sat | Review mediums; tag mistake type | — | — |
| Sun | Rest | — | — |

**Phase 2 checklist**
- [ ] `left`, `right`, `while (left < right)` template is muscle memory
- [ ] `while (left <= right)` binary search template is muscle memory
- [ ] Before coding mediums: pass the translation test every time
- [ ] Solved at least 4 medium two-pointer primaries

---

## Phase 3 — Hashing mastery & medium arrays (Weeks 5–6)

**Objective:** Subarray problems, prefix sums, and hash maps become your default toolkit — and you know when the prompt is really a subarray vs subsequence question.

### Week 5

| Day | Focus | Primary | Stretch |
|-----|-------|---------|---------|
| Mon | `drills/write/reflex/02_hashing_reflex/` full rewrite | `hashing/medium/group_anagram.js` | `hashing/medium/top_k_ferquent_element.js` |
| Tue | Prefix sum + hash | `misc/easy/find_pivot_index.js` | `misc/medium/subarray_sum_divisible_by_k.js` |
| Wed | Array medium | `arrays/medium/find_all_duplicates_in_an_array.js` | `arrays/medium/max_product_subarray.js` |
| Thu | Set tricks | `hashing/medium/longest_consecutive_sequence.js` | — |
| Fri | Prefix / product | `misc/medium/product_of_array_except_self.js` | — |
| Sat | Re-drill `01`–`04` | — | — |
| Sun | Rest | — | — |

### Week 6 — Consolidation week

- Mon–Thu: Re-solve **1** medium/day you got wrong before (understand again → solve blind)
- Fri: Timed set: 2 medium in 90 min (not 3 — quality over volume)
- Sat: Write a one-page "pattern journal" (see template at bottom)
- Sun: Rest

**Phase 3 checklist**
- [ ] 8+ medium primaries solved
- [ ] Can explain prefix sum in one sentence with an example
- [ ] Group anagrams without looking at notes
- [ ] Can say "subarray vs subsequence" correctly for each Week 5 primary

---

## Phase 4 — Trees, graphs, stacks, DP (Weeks 7–9)

**Objective:** Recursive thinking and state definition — the bridge to hard problems.

### Week 7 — Trees & stacks

| Day | Reflex | Primary | Stretch |
|-----|--------|---------|---------|
| Mon | `drills/write/reflex/05_trees_stacks_reflex/` | `trees/easy/convert_sorted_array_to_binary_search_tree.js` | — |
| Tue | Tree BFS/DFS | 1 tree easy (repo or LeetCode) | — |
| Wed | Stack patterns | `stack_queue/easy/last_stone_weight.js` | — |
| Thu | Monotonic stack (study + 1) | 1 next-greater style problem | — |
| Fri | Matrix / graph easy | `matrix/easy/flood_fill.js` or `graphs/easy/flood_fill.js` | — |
| Sat | Graph easy | `graphs/easy/island_permeter.js` | — |
| Sun | Rest | — | — |

### Week 8 — Dynamic programming (easy)

| Day | Focus | Primary | Stretch |
|-----|-------|---------|---------|
| Mon | `drills/write/reflex/06_dp_reflex/` | `dynamic_programming/easy/fibonacci_number.js` | `dynamic_programming/easy/min_cost_climbing_staris.js` |
| Tue | 1D DP | `dynamic_programming/easy/pascale_triangle_1.js` | `dynamic_programming/easy/pascale_triangle_2.js` |
| Wed | DP on strings (study classic: LCS / edit distance) | Read + rewrite one idea | — |
| Thu | Greedy vs DP | `greedy/easy/can_place_flower.js` | `greedy/easy/lemonade_change.js` |
| Fri | Mixed | 1 DP easy | 1 greedy easy |
| Sat | Re-type DP templates from memory | — | — |
| Sun | Rest | — | — |

### Week 9 — Graphs BFS/DFS

| Day | Reflex | Focus |
|-----|--------|-------|
| Mon | `drills/write/reflex/07_graphs_reflex/` | BFS grid — primary: finish BFS section |
| Tue | Same drill | DFS grid + visited — primary: finish DFS section |
| Wed | — | 1 repo graph easy |
| Thu | — | 1 medium graph (attempt + understand editorial) |
| Fri | Mock: medium tree or graph | 45 min |
| Sat | Full reflex review: all 7 drill files | — |
| Sun | Rest | — |

---

## Phase 5 — Medium consolidation & hard intro (Weeks 10–12)

**Objective:** Medium problems feel *readable*. Hard problems feel *approachable*.

### Weekly rhythm (repeat weeks 10–12)

| Day | Activity |
|-----|----------|
| Mon | 1 new medium (90 min max). Spend first 5 min on question literacy only. |
| Tue | Re-solve yesterday's medium from scratch (no peeking at your old code) |
| Wed | 1 new medium + classify pattern used |
| Thu | 1 hard **study** problem: 30 min attempt → deep-read editorial → rewrite from memory |
| Fri | Timed: 2 medium in 90 min |
| Sat | Reflex: random drill file + `pattern_cheat_sheet.go` |
| Sun | Rest |

**Suggested hard intro (add to repo or use LeetCode):**
- `two_pointers/hard/trapping_rain_water.js` (you already have this!)
- Binary search on answer: Koko eating bananas
- DP: coin change, longest increasing subsequence (you have LIS in easy — upgrade to O(n log n))

**Phase 5 success criteria**
- [ ] 20+ medium problems total
- [ ] 3+ hard problems studied (attempt + editorial + rewrite)
- [ ] Can open a random medium, pass the translation test, and list 2–3 approaches in 5 minutes

---

## The problem-solving framework (use on EVERY medium/hard)

When you feel stuck, your brain freezes because there's no **process**. Follow these steps literally — write them in comments before coding.

### Step 0 — Understand (3–5 min) — do not skip

Use the [Question literacy checklist](#question-literacy-checklist). Pass the translation test.

### Step 1 — Restate (2 min)

```
Input:  ?
Output: ?
Constraints: size? sorted? duplicates? negative?
Edge:    empty, single element, all same, max size
```

### Step 2 — Brute force first (5 min)

Name the naive solution and its complexity. Interviewers and your future self only trust you if you know the slow way first.

> Example: 3Sum → three nested loops O(n³). Good. Now we have a baseline.

### Step 3 — Pattern scan (3 min)

Ask these triggers in order:

| If you see… | Think… |
|-------------|--------|
| Sorted array, pair/triplet sum | Two pointers |
| Subarray sum / count subarrays | Prefix sum + hash map |
| "Longest/shortest subarray with property" | Sliding window |
| Sorted + find position / min/max answer | Binary search |
| Frequency, anagram, two-sum variant | Hash map / set |
| Tree, levels, shortest path unweighted | BFS |
| Connected components, explore all paths | DFS + visited |
| Optimal substructure, choices at each step | DP |
| Local greedy choice seems optimal | Greedy (prove or test) |
| Next greater/smaller element | Monotonic stack |

### Step 4 — Sketch (5 min)

Draw a tiny example. Trace your chosen approach by hand. If trace fails, go back to Step 3.

### Step 5 — Code (15–25 min)

Use templates from `_support/templates/pattern_cheat_sheet.go`. Don't invent syntax under pressure.

### Step 6 — Post-mortem (5 min)

After every problem, log:

```
Problem:
Ask (1 sentence):
Pattern used:
Mistake type: [logic | edge case | pattern miss | understanding | syntax | timeout]
One sentence lesson:
Revisit date: (+3 days)
```

---

## Mistake taxonomy (stop repeating the same failure)

| Type | Symptom | Fix |
|------|---------|-----|
| **Understanding** | Solved a different problem than asked | Slow down on Step 0; rewrite ask in one sentence |
| **Pattern miss** | Didn't know where to start | Re-read pattern scan table; drill that pattern 3 days |
| **Edge case** | Wrong on empty/single/duplicates | Write edge cases in Step 1 every time |
| **Off-by-one** | WA on boundaries | Trace `left`, `right`, `mid` on paper |
| **State definition** | DP wrong | Define `dp[i]` meaning in English before coding |
| **Timeout** | TLE | Count nested loops; aim for O(n) or O(n log n) |

---

## Reflex training protocol

**Why reflex matters:** Medium/hard problems are slow if you're still remembering *how* to write two pointers. Basics must be automatic so working memory is free for *understanding* and *reasoning*.

### Daily reflex block (see `DAILY_30MIN_DRILL.md`)

```bash
cd study_play && go run .        # today's file + prompts
cd study_play && go run . -- --run  # run tests after implementing
```

1. Open today's drill with answers **covered**.
2. Re-implement every `TODO: REFLEX` function from memory.
3. Run the drill file (or `--run`).
4. If any function takes &gt; 3 min or fails tests → mark it; repeat tomorrow.

### Spaced repetition schedule

| After solving a problem | Review again on |
|-------------------------|-----------------|
| First solve | Day +1 |
| Second pass | Day +3 |
| Third pass | Day +7 |
| Fourth pass | Day +21 |

Keep a simple list in a notebook or append to post-mortem logs.

---

## Pattern journal template (end of Week 6, update monthly)

```markdown
## Patterns I own
- Two sum hash: [confidence 1-5]
- Sliding window: [1-5]
- ...

## Patterns I confuse
- DP vs greedy when: ...

## Questions I misunderstood (and the real ask)
1.
2.

## 3 problems that taught me the most
1.
2.
3.
```

---

## Mindset (read when motivation drops)

1. **Stuck for 30 minutes is training, not failure.** Medium problems are *designed* to stall you.
2. **One understood medium beats three rushed easies.** Deliberate difficulty grows skill.
3. **Reading a solution after honest effort is learning** — copying without attempt is not.
4. **Your repo is your gym.** The 170 files are equipment; this plan is the program.
5. **Consistency &gt; intensity.** The Minimum tier still counts. Showing up keeps the habit alive.

If you feel overwhelmed: do **only** the reflex block that day. That is enough.

---

## Quick reference — run drills

```bash
# From study_play/
go run .
go run -C drills/write/reflex/01_arrays_reflex .
go run -C drills/write/reflex/02_hashing_reflex .
go run -C drills/write/reflex/03_two_pointers_reflex .
go run -C drills/write/reflex/04_binary_search_reflex .
go run -C drills/write/reflex/05_trees_stacks_reflex .
go run -C drills/write/reflex/06_dp_reflex .
go run -C drills/write/reflex/07_graphs_reflex .
```

---

## After 12 weeks

- Maintain: 3 problems/week (1 medium, 1 review, 1 new)
- Monthly: re-type full `pattern_cheat_sheet.go` from memory
- Expand repo: add `hard/` folders per topic as you grow

You are building a skill, not racing a clock. Same ritual every day. Understand first. One primary. Log the lesson.
