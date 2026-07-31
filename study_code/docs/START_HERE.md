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
5. Misses → peek **only that snippet** in `drills/read/answers/`, then re-trace once without looking.

## Folder map

```
drills/read/                   ← practice files (start here)
study_code/
├── docs/                      ← START_HERE, READING_PATTERNS, DAILY_20MIN_READING
├── practice/read/             ← core/, weekday/, answers/
└── daily_read.go
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
