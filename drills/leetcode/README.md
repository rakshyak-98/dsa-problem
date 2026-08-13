# LeetCode Daily Practice

**10 full LeetCode problems per weekday**, themed to match the reflex write drill rotation. This is **separate** from in-repo reflex drills — solve these on [leetcode.com](https://leetcode.com).

Problems are **fetched live from the LeetCode GraphQL API** when you run today's command and saved to [`daily.json`](daily.json).

## How it fits

| Layer | What | Where |
|-------|------|--------|
| Reflex drill | 4–6 tiny functions from memory | `drills/write/reflex/` |
| **LeetCode practice** | **10 API-fetched full problems** | **`daily.json` + `--track leetcode`** |
| Progress | Session log + LeetCode links | `drills/tracker/study_tracker.html` |

## Commands

```bash
go run . -- --track leetcode              # fetch (if needed) + show today's 10 problems
go run . -- --run leetcode                 # same from default DSA track
go run . -- --run -l                       # short flag (like -w for write reflex)
go run . -- --track leetcode -- --refresh # force re-fetch from LeetCode API
go run . -- --track leetcode -- --catalog # all weekday sets
go -C bin/study_leetcode run .            # direct CLI
```

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
3. Solve 2–3 problems on LeetCode (or all 10 if you have time)
4. Log URLs in the study tracker

## How fetching works

1. Includes today's [LeetCode daily challenge](https://leetcode.com/) when its tags match the weekday topic
2. Fills remaining slots from a topic seed pool using a date-based shuffle
3. Skips paid-only problems
4. Writes results to `drills/leetcode/daily.json` (cached for the rest of the day)
5. Falls back to cached `daily.json` if the API is unreachable

## Notes

- Requires network access to `leetcode.com`
- `daily.json` is local cache (gitignored); regenerate anytime with `--refresh`
- Each entry includes live title, difficulty, tags, URL, and reflex cross-ref when known
