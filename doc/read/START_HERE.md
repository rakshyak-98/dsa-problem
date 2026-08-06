# Start Here — Code Reading

Open **`READING_PATTERNS.md`** for the 6-pass method.  
Open **`DAILY_20MIN_READING.md`** for the daily ritual.

## Today (Day 1)

1. Read **The 6-pass read** in `READING_PATTERNS.md` (8 min).
2. Run the helper:
   ```bash
   go run ./bin/study_code
   ```
3. Open today’s drill. For each snippet:
   - Cover the answer variables
   - Do Pass 1–3 from the method
   - Fill every `TODO: READ` answer
4. Run:
   ```bash
   go run ./bin/study_code -- --run
   ```
5. Misses → peek **only that snippet** in `drills/read/answers/`, then re-trace once without looking.

## Folder map

```
drills/read/                   ← practice files (start here)
doc/read/                      ← guides (this folder)
bin/study_code/                ← read drill CLI
```

## Every reading day

```bash
go run ./bin/study_code              # Core Read 3 + today’s specialty
go run ./bin/study_code -- --run     # check your filled answers
go run ./bin/study_code -- --micro   # Core Read 3 only
go run ./bin/study_code -- --catalog # full checklist
```

**Minimum:** Core Read 3.  
**Standard:** Core Read 3 + specialty drill.
