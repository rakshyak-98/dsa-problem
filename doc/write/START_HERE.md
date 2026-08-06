# Start Here

Open **`STUDY_PLAN.md`** for the full 12-week roadmap.
Open **`DAILY_30MIN_DRILL.md`** for the daily ritual.

## Today (Day 1)

1. Read **The consistency rule** and **Understand the question first** in `STUDY_PLAN.md` (10 min).
2. Run the daily helper:
   ```bash
   go run ./bin/study_play
   ```
3. Open today's drill and implement every `TODO: REFLEX` from memory.
4. Run tests:
   ```bash
   go run ./bin/study_play -- --run
   ```
5. Stuck after 15 min? Peek at `drills/solutions/`, then close it and re-type.
6. Primary problem: `hashing/easy/two_sum.js` — **restate the ask in one sentence before coding**.

## Folder map

```
drills/write/              ← practice files (start here)
drills/solutions/          ← write solutions (peek after attempt)
doc/write/                 ← guides (this folder, incl. MATH_CONCEPTS.md)
bin/study_play/            ← CLI + blank templates (_support/)
```

## Every study day

```bash
go run ./bin/study_play              # Core 5 + today's specialty
go run ./bin/study_play -- --run     # run specialty tests
go run ./bin/study_play -- --reset   # wipe today's drill back to TODO stubs
go run ./bin/study_play -- --drill   # Core 5 only (low energy)
go run ./bin/study_play -- --catalog # full essential checklist
```

**Reflex tier:** Core 5 + specialty drill.  
**Minimum tier:** Core 5 only. Still builds reflexes.

Full pack: **`DAILY_30MIN_DRILL.md`**. Consistency beats marathon sessions.
