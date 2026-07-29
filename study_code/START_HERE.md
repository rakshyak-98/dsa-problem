# Start Here — Code Reading

Open **`READING_PATTERNS.md`** for the 6-pass method.  
Open **`DAILY_20MIN_READING.md`** for the daily ritual.

## Today (Day 1)

1. Read **The 6-pass read** in `READING_PATTERNS.md` (8 min).
2. Run the helper:
   ```bash
   cd study_code && go run .
   ```
3. Open today’s drill. For each snippet:
   - Cover the answer variables
   - Do Pass 1–3 from the method
   - Fill every `TODO: READ` answer
4. Run:
   ```bash
   cd study_code && go run . -- --run
   ```
5. Misses → peek **only that snippet** in `drills/answers/`, then re-trace once without looking.

## Folder map

```
study_code/
├── START_HERE.md              ← you are here
├── READING_PATTERNS.md        ← how to read (memorize the passes)
├── DAILY_20MIN_READING.md     ← daily clock + weekday rotation
├── daily_read.go
├── drills/
│   ├── 01_scan_structure/     # map I/O, loops, state
│   ├── 02_trace_execution/    # hand-simulate one input
│   ├── 03_name_the_pattern/   # code → pattern name
│   ├── 04_find_the_bug/       # spot the wrong move
│   ├── 05_complexity_glance/  # time/space from shape
│   ├── 06_reconstruct_ask/    # code → one-sentence ask
│   ├── 07_compare_variants/   # two solutions, one difference
│   └── answers/
└── worksheets/
    └── reading_log.md         # optional session log
```

## Every reading day

```bash
cd study_code
go run .              # Core Read 3 + today’s specialty
go run . -- --run     # check your filled answers
go run . -- --micro   # Core Read 3 only
go run . -- --catalog # full checklist
```

**Minimum:** Core Read 3.  
**Standard:** Core Read 3 + specialty drill.
