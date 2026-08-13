# LeetCode Daily Practice

**10 full LeetCode problems per weekday**, themed to match the reflex write drill rotation. This is **separate** from in-repo reflex drills — solve these on [leetcode.com](https://leetcode.com).

## How it fits

| Layer | What | Where |
|-------|------|--------|
| Reflex drill | 4–6 tiny functions from memory | `drills/write/reflex/` |
| **LeetCode practice** | **10 curated full problems** | **This folder + `--track leetcode`** |
| Progress | Session log + LeetCode links | `drills/tracker/study_tracker.html` |

## Commands

```bash
go run . -- --track leetcode              # today's 10 problems
go run . -- --track leetcode -- --catalog # all weekday sets
go run ./bin/study_leetcode               # direct CLI (or: go -C bin/study_leetcode run .)
```

## Weekday topics

| Day | Topic | Reflex drill | Problem count |
|-----|-------|--------------|---------------|
| Monday | Arrays | `01_arrays_reflex` | 10 |
| Tuesday | Hashing | `02_hashing_reflex` | 10 |
| Wednesday | Two Pointers | `03_two_pointers_reflex` | 10 |
| Thursday | Binary Search | `04_binary_search_reflex` | 10 |
| Friday | Trees & Stacks | `05_trees_stacks_reflex` | 10 |
| Saturday | Dynamic Programming | `06_dp_reflex` | 10 |
| Sunday | Graphs | `07_graphs_reflex` | 10 |

## Daily flow (recommended)

1. Run reflex drill: `go run . -- --drill reflex` then implement in `drills/write/reflex/`
2. Run LeetCode set: `go run . -- --track leetcode`
3. Solve 2–3 problems on LeetCode (or all 10 if you have time)
4. Log URLs in the study tracker

## Notes

- Problems are classic interview picks from LeetCode, not generated stubs.
- Each entry links to the matching reflex function when one exists.
- Difficulty mix: mostly Easy/Medium with 1–2 Hard stretch problems on some days.
