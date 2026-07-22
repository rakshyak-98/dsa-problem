# Start Here

Open **`STUDY_PLAN.md`** for the full 12-week roadmap.
Open **`DAILY_30MIN_DRILL.md`** for the daily ritual.

## Today (Day 1)

1. Read **The consistency rule** and **Understand the question first** in `STUDY_PLAN.md` (10 min).
2. Run the daily helper:
   ```bash
   node study_play/daily_drill.js
   ```
3. Open today's drill and implement every `TODO: REFLEX` from memory.
4. Run tests:
   ```bash
   node study_play/daily_drill.js --run
   ```
5. Stuck after 15 min? Peek at `drills/_solutions_reference.js`, then close it and re-type.
6. Primary problem: `hashing/easy/two_sum.js` — **restate the ask in one sentence before coding**.

## Folder map

```
study_play/
├── START_HERE.md          ← you are here
├── STUDY_PLAN.md          ← full plan + weekly schedule
├── DAILY_30MIN_DRILL.md   ← daily ritual (tiers + clock)
├── daily_drill.js         ← prints today's drill + prompts
├── drills/
│   ├── 01_arrays_reflex.js
│   ├── 02_hashing_reflex.js
│   ├── 03_two_pointers_reflex.js
│   ├── 04_binary_search_reflex.js
│   ├── 05_trees_stacks_reflex.js
│   ├── 06_dp_reflex.js
│   ├── 07_graphs_reflex.js
│   └── _solutions_reference.js   ← only after honest attempt
└── templates/
    └── pattern_cheat_sheet.js    ← re-type from memory weekly
```

## Every study day

```bash
node study_play/daily_drill.js           # Core 5 + today's specialty
node study_play/daily_drill.js --run     # run specialty tests
node study_play/daily_drill.js --micro   # Core 5 only (low energy)
node study_play/daily_drill.js --catalog # full essential checklist
```

**Reflex tier:** Core 5 + specialty drill.  
**Minimum tier:** Core 5 only. Still builds reflexes.

Full pack: **`DAILY_30MIN_DRILL.md`**. Consistency beats marathon sessions.
