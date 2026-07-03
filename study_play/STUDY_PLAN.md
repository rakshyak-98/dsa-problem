# DSA Study Plan — From Reflex to Medium/Hard Mastery

> **Goal:** Make basic patterns automatic (reflex), then train your brain to *recognize* and *decompose* medium–hard problems without panic.
>
> **Repo:** You already have ~170 `.js` problems organized by topic. This plan tells you *what* to do, *in what order*, and *how* to think.

---

## How to use this folder

| File / folder | Purpose |
|---------------|---------|
| `STUDY_PLAN.md` (this file) | Full roadmap, weekly schedule, mental frameworks |
| `drills/*.js` | **Reflex drills** — rewrite from memory until automatic |
| `templates/pattern_cheat_sheet.js` | Blank skeletons; cover the answers and re-type daily |
| `../arrays/`, `../hashing/`, etc. | Real problems from your repo |

**Daily minimum (non-negotiable):** 45–90 minutes, 6 days/week. Rest 1 day.

---

## The 5-phase roadmap (12 weeks)

```
Phase 1  Weeks 1–2   Foundations + reflex basics
Phase 2  Weeks 3–4   Two pointers, sliding window, binary search
Phase 3  Weeks 5–6   Hashing mastery + medium arrays/strings
Phase 4  Weeks 7–9   Trees, graphs, stacks, DP intro
Phase 5  Weeks 10–12 Medium consolidation + hard introduction
```

Track progress by checking boxes in each week's section below.

---

## Phase 1 — Foundations (Weeks 1–2)

**Objective:** Arrays, loops, and hashing feel like breathing. No lookup needed for two-sum or frequency count.

### Week 1 — Arrays & complexity

| Day | Reflex drill | Repo problems (solve without solution first) | Time cap |
|-----|--------------|-----------------------------------------------|----------|
| Mon | `drills/01_arrays_reflex.js` — all 5 functions | `arrays/easy/plus_one.js`, `arrays/easy/concatenation_of_array.js` | 25 min each |
| Tue | Re-drill arrays (from memory) | `arrays/easy/find_closest_number_to_zero.js`, `arrays/easy/max_consecutive_ones.js` | 25 min |
| Wed | `drills/02_hashing_reflex.js` — twoSum, freqMap | `hashing/easy/two_sum.js`, `hashing/easy/contains_duplicates.js` | 25 min |
| Thu | Re-drill hashing | `hashing/easy/fair_candy_swap.js`, `hashing/easy/degree_of_an_array.js` | 25 min |
| Fri | Mixed: 1 array + 1 hash | `arrays/easy/majority_element.js`, `hashing/easy/summary_ranges.js` | 30 min |
| Sat | **Review day:** re-do any failed problems; re-type drills blind | — | 60 min |
| Sun | Rest | Walk, sleep, no guilt | — |

**Week 1 checklist**
- [ ] Can write `reverseArray`, `maxInArray`, `countFreq` from memory in &lt; 3 min each
- [ ] Can explain O(n) vs O(n²) for every solution you wrote
- [ ] Solved 8+ easy problems without peeking at old code

### Week 2 — Strings & simulation

| Day | Focus | Repo problems |
|-----|-------|---------------|
| Mon | String basics (char codes, split, join) | `strings/easy/find_words_containing_character.js`, `strings/easy/most_common_word.js` |
| Tue | Reflex: hashing on strings | `strings/easy/unique_morse_code_words.js`, `hashing/easy/find_resultant_array_after_removing_anagrams.js` |
| Wed | Simulation / step-by-step | `simulation/easy/baseball_game.js`, `simulation/easy/relative_ranks.js` |
| Thu | Mixed easy batch (3 problems, 20 min each) | Pick 3 from `strings/easy/` you haven't done |
| Fri | **Timed mock:** 2 easy in 45 min total | Random from weeks 1–2 |
| Sat | Re-drill `01` + `02`; fix weak spots | — |
| Sun | Rest | — |

**Week 2 checklist**
- [ ] Comfortable with `Map`, `Set`, and object freq maps
- [ ] 16+ total easy problems solved lifetime in this repo

---

## Phase 2 — Two pointers & binary search (Weeks 3–4)

**Objective:** See "sorted array", "pair sum", "palindrome" → instantly think two pointers or binary search.

### Week 3 — Two pointers

| Day | Reflex drill | Repo problems |
|-----|--------------|---------------|
| Mon | `drills/03_two_pointers_reflex.js` | `two_pointers/easy/remove_duplicates_from_sorted_array.js`, `two_pointers/easy/move_zeroes.js` |
| Tue | Re-drill + opposite ends | `two_pointers/easy/best_time_to_buy_sell_stock.js`, `two_pointers/easy/squares_of_a_sorted_array.js` |
| Wed | Medium intro | `two_pointers/medium/container_with_most_water.js`, `two_pointers/medium/3sum.js` |
| Thu | Sliding window in drill file | `misc/easy/maximum_average_subarray_1.js`, `two_pointers/easy/minimum_difference_between_highest_and_lowest_of_k_score.js` |
| Fri | Timed: 1 easy + 1 medium | Your choice from `two_pointers/` |
| Sat | Re-type `templates/pattern_cheat_sheet.js` sections 1–3 | — |
| Sun | Rest | — |

### Week 4 — Binary search

| Day | Reflex drill | Repo problems |
|-----|--------------|---------------|
| Mon | `drills/04_binary_search_reflex.js` | `binary_search/easy/search_insertion_position.js`, `binary_search/easy/find_smallest_letter_greater_than_target.js` |
| Tue | Binary search on answer (concept) | `binary_search/easy/longest_subsequence_with_limited_sum.js` |
| Wed | Mixed pointers + BS | `two_pointers/medium/find_the_duplicate_number.js` |
| Thu | Medium batch | `two_pointers/medium/sort_colors.js`, `two_pointers/medium/longest_palindromic_substring.js` |
| Fri | **Mock interview:** 1 medium, 45 min, talk out loud | Record yourself or write steps in comments |
| Sat | Review all medium attempts; categorize mistake type (see framework below) | — |
| Sun | Rest | — |

**Phase 2 checklist**
- [ ] `left`, `right`, `while (left < right)` template is muscle memory
- [ ] `while (left <= right)` binary search template is muscle memory
- [ ] Solved at least 5 medium two-pointer problems

---

## Phase 3 — Hashing mastery & medium arrays (Weeks 5–6)

**Objective:** Subarray problems, prefix sums, and hash maps become your default toolkit.

### Week 5

| Day | Focus | Repo problems |
|-----|-------|---------------|
| Mon | `drills/02_hashing_reflex.js` (full rewrite) | `hashing/medium/group_anagram.js`, `hashing/medium/top_k_ferquent_element.js` |
| Tue | Prefix sum + hash | `misc/medium/subarray_sum_divisible_by_k.js`, `misc/easy/find_pivot_index.js` |
| Wed | Array medium | `arrays/medium/max_product_subarray.js`, `arrays/medium/find_all_duplicates_in_an_array.js` |
| Thu | Longest consecutive / set tricks | `hashing/medium/longest_consecutive_sequence.js` |
| Fri | Product except self (no division) | `misc/medium/product_of_array_except_self.js` |
| Sat | Re-drill all files `01`–`04` | — |
| Sun | Rest | — |

### Week 6 — Consolidation week

- Mon–Thu: Re-solve 2 medium/day you got wrong before (spaced repetition)
- Fri: Timed set: 3 medium in 2 hours
- Sat: Write a one-page "pattern journal" (see template at bottom)
- Sun: Rest

**Phase 3 checklist**
- [ ] 10+ medium problems solved
- [ ] Can explain prefix sum in one sentence with an example
- [ ] Group anagrams & top-k without looking at notes

---

## Phase 4 — Trees, graphs, stacks, DP (Weeks 7–9)

**Objective:** Recursive thinking and state definition — the bridge to hard problems.

### Week 7 — Trees & stacks

| Day | Reflex drill | Repo problems |
|-----|--------------|---------------|
| Mon | `drills/05_trees_stacks_reflex.js` — traversals | `trees/easy/convert_sorted_array_to_binary_search_tree.js` |
| Tue | BFS/DFS tree | Add tree problems as you expand repo |
| Wed | Stack patterns | `stack_queue/easy/last_stone_weight.js` |
| Thu | Monotonic stack concept (study + 1 problem) | — |
| Fri–Sat | Matrix + graph easy | `matrix/easy/flood_fill.js` → `graphs/easy/flood_fill.js`, `graphs/easy/island_permeter.js` |
| Sun | Rest | — |

### Week 8 — Dynamic programming (easy)

| Day | Focus | Repo problems |
|-----|-------|---------------|
| Mon | `drills/06_dp_reflex.js` | `dynamic_programming/easy/fibonacci_number.js`, `dynamic_programming/easy/min_cost_climbing_staris.js` |
| Tue | 1D DP | `dynamic_programming/easy/pascale_triangle_1.js`, `dynamic_programming/easy/pascale_triangle_2.js` |
| Wed | DP on strings (study classic: LCS, edit distance — add to repo) | — |
| Thu | Greedy vs DP | `greedy/easy/can_place_flower.js`, `greedy/easy/lemonade_change.js` |
| Fri | Mixed | 1 DP easy + 1 greedy easy |
| Sat | Re-type DP templates from memory | — |
| Sun | Rest | — |

### Week 9 — Graphs BFS/DFS

| Day | Reflex drill | Focus |
|-----|--------------|-------|
| Mon–Tue | `drills/07_graphs_reflex.js` | BFS grid, DFS grid, visited set |
| Wed–Thu | Repo graph easies + 1 medium graph (add or external) | — |
| Fri | Mock: medium tree or graph | 45 min |
| Sat | Full reflex review: all 7 drill files | — |
| Sun | Rest | — |

---

## Phase 5 — Medium consolidation & hard intro (Weeks 10–12)

**Objective:** Medium problems feel *readable*. Hard problems feel *approachable* (not solved instantly — that's normal).

### Weekly rhythm (repeat weeks 10–12)

| Day | Activity |
|-----|----------|
| Mon | 1 new medium (90 min max, then read solution if stuck) |
| Tue | Re-solve yesterday's medium from scratch |
| Wed | 1 new medium + classify pattern used |
| Thu | 1 hard **study** problem: 30 min attempt, then deep-read editorial, rewrite from memory |
| Fri | Timed: 2 medium in 90 min |
| Sat | Reflex: random drill file + `pattern_cheat_sheet.js` |
| Sun | Rest |

**Suggested hard intro (add to repo or use LeetCode):**
- `two_pointers/hard/trapping_rain_water.js` (you already have this!)
- Binary search on answer: Koko eating bananas
- DP: coin change, longest increasing subsequence (you have LIS in easy — upgrade to O(n log n))

**Phase 5 success criteria**
- [ ] 25+ medium problems total
- [ ] 3+ hard problems studied (attempt + editorial + rewrite)
- [ ] Can open a random medium and list 2–3 possible approaches in 5 minutes

---

## The problem-solving framework (use on EVERY medium/hard)

When you feel stuck, your brain freezes because there's no **process**. Follow these steps literally — write them in comments before coding.

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

Use templates from `templates/pattern_cheat_sheet.js`. Don't invent syntax under pressure.

### Step 6 — Post-mortem (5 min)

After every problem, log:

```
Problem:
Pattern used:
Mistake type: [logic | edge case | pattern miss | syntax | timeout]
One sentence lesson:
Revisit date: (+3 days)
```

---

## Mistake taxonomy (stop repeating the same failure)

| Type | Symptom | Fix |
|------|---------|-----|
| **Pattern miss** | Didn't know where to start | Re-read pattern scan table; drill that pattern 3 days |
| **Edge case** | Wrong on empty/single/duplicates | Write edge cases in Step 1 every time |
| **Off-by-one** | WA on boundaries | Trace `left`, `right`, `mid` on paper |
| **State definition** | DP wrong | Define `dp[i]` meaning in English before coding |
| **Timeout** | TLE | Count nested loops; aim for O(n) or O(n log n) |

---

## Reflex training protocol

**Why reflex matters:** Medium/hard problems are slow if you're still remembering *how* to write two pointers. Basics must be automatic so working memory is free for *reasoning*.

### Daily reflex block (10–15 min, before new problems)

1. Open a drill file with answers **covered**.
2. Re-implement every `TODO: REFLEX` function from memory.
3. Run: `node study_play/drills/01_arrays_reflex.js` (etc.)
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

## 3 problems that taught me the most
1.
2.
3.
```

---

## Mindset (read when motivation drops)

1. **Stuck for 30 minutes is training, not failure.** Medium problems are *designed* to stall you.
2. **One medium a day beats five easies you already know.** Deliberate difficulty grows skill.
3. **Reading a solution after honest effort is learning** — copying without attempt is not.
4. **Your repo is your gym.** The 170 files are equipment; this plan is the program.
5. **Consistency &gt; intensity.** 45 minutes daily for 12 weeks changes what your brain reaches for automatically.

If you feel overwhelmed: do **only** the reflex block + **one** easy problem that day. Showing up keeps the habit alive.

---

## Quick reference — run drills

```bash
# From repo root
node study_play/drills/01_arrays_reflex.js
node study_play/drills/02_hashing_reflex.js
node study_play/drills/03_two_pointers_reflex.js
node study_play/drills/04_binary_search_reflex.js
node study_play/drills/05_trees_stacks_reflex.js
node study_play/drills/06_dp_reflex.js
node study_play/drills/07_graphs_reflex.js
```

---

## After 12 weeks

- Maintain: 3 problems/week (1 medium, 1 review, 1 new)
- Monthly: re-type full `pattern_cheat_sheet.js` from memory
- Expand repo: add `hard/` folders per topic as you grow

You are building a skill, not racing a clock. Follow the weeks in order, check boxes honestly, and let reflex drills run every single study day.
