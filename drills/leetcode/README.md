# LeetCode Daily Practice

**10 full LeetCode problems per weekday**, themed to match the reflex write drill rotation. This is **separate** from in-repo reflex drills — solve these on [leetcode.com](https://leetcode.com).

Problems are **fetched live from the LeetCode GraphQL API** when you run today's command and saved to [`daily.json`](daily.json).

## How it fits

| Layer | What | Where |
|-------|------|--------|
| Reflex drill | 4–6 tiny functions from memory | `drills/write/reflex/` |
| **LeetCode practice** | **10 API-fetched full problems** | **`daily.json` + `--track leetcode`** |
| Progress | First-try result and next review per function | `.drill_log.json` + `--weak` |
| **Cards** | **Per solved problem: your pattern guess, the signal, the insight, the next re-solve** | **`cards.json` + `--log` / `--gaps`** |

## Commands

```bash
go run . -- --track leetcode              # fetch (if needed) + show today's 10 problems
go run . -- --run leetcode                 # same from default DSA track
go run . -- --run -l                       # short flag (like -w for write reflex)
go run . -- --track leetcode -- --refresh # force re-fetch from LeetCode API
go run . -- --track leetcode -- --catalog # all weekday sets
go -C bin/study_leetcode run .            # direct CLI
go run . -- --track leetcode -- --answers # also show topic, tags and reflex links
go run . -- --log=two-sum                 # record (or review) a solved problem
go run . -- --gaps                        # patterns on your cards with no reflex drill
```

Topic, tags and reflex links are **hidden** in the terminal listing and moved
to an *Answers* section at the end of `daily.md` / `daily.go`: naming the
pattern yourself is the first thing to practise.

## Weekday topics

| Day | Topic | LeetCode tags | Reflex drill |
|-----|-------|---------------|--------------|
| Mon | Arrays | `array` | `01_arrays_reflex` |
| Tue | Hashing | `hash-table` | `02_hashing_reflex` |
| Wed | Two Pointers | `two-pointers`, `sliding-window` | `03_two_pointers_reflex` |
| Thu | Binary Search | `binary-search` | `04_binary_search_reflex` |
| Fri | Trees & Stacks | `tree`, `stack`, `binary-tree` | `05_trees_stacks_reflex` |
| Sat | Dynamic Programming | `dynamic-programming` | `06_dp_reflex` |
| Sun | Graphs | `graph`, `dfs`, `bfs` | `07_graphs_reflex` |

## Daily flow (recommended)

1. Run reflex drill: `go run . -- --drill reflex` then implement in `drills/write/reflex/`
2. Run LeetCode set: `go run . -- --track leetcode` (updates `daily.json`)
3. **Re-solve the due cards first** (listed under RE-SOLVE): read the statement,
   name the pattern within 60 s, write the skeleton, then `--log=SLUG`
4. For 2–3 new problems: write your pattern guess and the words that point to
   it **before** reading the Answers section, solve, then `--log=SLUG`
5. Every week or so: `--gaps` — a pattern on 2+ cards with no reflex drill is
   the next drill file to write

A card moves up a box (re-solve in 3, 7, 14, 30, 60 days) when you named the
pattern **and** solved it clean; one of the two keeps its box; neither, or
reading the solution, brings it back tomorrow. Cards also appear in
`--recognize`.

## How fetching works

1. Includes today's [LeetCode daily challenge](https://leetcode.com/) when its tags match the weekday topic
2. Randomly picks the rest from LeetCode's live topic catalog (same tags as the weekday), using a date-based shuffle so the set is stable for the day
3. Falls back to the curated seed pool only if the topic API cannot fill all 10 slots
4. Skips paid-only problems
5. Writes results to `drills/leetcode/daily.json` (cached for the rest of the day)
6. Falls back to cached `daily.json` if the API is unreachable

## Notes

- Requires network access to `leetcode.com`
- `daily.json` is local cache (gitignored); regenerate anytime with `--refresh`
- Each entry includes live title, difficulty, tags, URL, and reflex cross-ref when known
