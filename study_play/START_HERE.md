# Start Here

Open **`STUDY_PLAN.md`** for the full 12-week roadmap.
Open **`DAILY_30MIN_DRILL.md`** for the daily ritual.

## Today (Day 1)

1. Read **The consistency rule** and **Understand the question first** in `STUDY_PLAN.md` (10 min).
2. Run the daily helper:
   ```bash
   cd study_play && go run .
   ```
3. Open today's drill and implement every `TODO: REFLEX` from memory.
4. Run tests:
   ```bash
   cd study_play && go run . -- --run
   ```
5. Stuck after 15 min? Peek at `drills/solutions_reference/`, then close it and re-type.
6. Primary problem: `hashing/easy/two_sum.js` — **restate the ask in one sentence before coding**.

## Folder map

```
study_play/
├── START_HERE.md          ← you are here
├── STUDY_PLAN.md          ← full plan + weekly schedule
├── DAILY_30MIN_DRILL.md   ← daily ritual (tiers + clock)
├── daily_drill.go         ← prints today's drill + prompts
├── go.mod
├── drills/
│   ├── 01_arrays_reflex/
│   ├── 02_hashing_reflex/
│   ├── 03_two_pointers_reflex/
│   ├── 04_binary_search_reflex/
│   ├── 05_trees_stacks_reflex/
│   ├── 06_dp_reflex/
│   ├── 07_graphs_reflex/
│   └── solutions_reference/   ← only after honest attempt
└── templates/
    └── pattern_cheat_sheet.go ← re-type from memory weekly
```

## Every study day

```bash
cd study_play
go run .              # Core 5 + today's specialty
go run . -- --run     # run specialty tests
go run . -- --reset   # wipe today's drill back to TODO stubs
go run . -- --micro   # Core 5 only (low energy)
go run . -- --catalog # full essential checklist
```

**Reflex tier:** Core 5 + specialty drill.  
**Minimum tier:** Core 5 only. Still builds reflexes.

Full pack: **`DAILY_30MIN_DRILL.md`**. Consistency beats marathon sessions.
