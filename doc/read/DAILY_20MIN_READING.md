# Daily Code Reading — 20 Minute Pack

> **Purpose:** Make pattern recognition and tracing automatic when *reading* code.  
> **Rule:** Fill answers blind. Open `drills/read/answers/` only after `--run` fails.  
> **Helper:** `go run ./bin/study_code`

---

## How every session works

```
1. Core Read 3   →  always (signature + trace + pattern)
2. Specialty     →  today’s weekday drill
3. Log           →  one miss + revisit in 3 days
```

| Tier | Time | Do |
|------|------|----|
| **Minimum** | ~10–12 min | Core Read 3 + log |
| **Standard** | ~20–25 min | Core Read 3 + today’s specialty |
| **Stretch** | ~30 min | Standard + re-read yesterday’s fails |

Missed a day? Do **not** catch up. Run today’s pack only.

```bash
go run ./bin/study_code
go run ./bin/study_code -- --run
go run ./bin/study_code -- --micro
```

---

## Part A — Core Read 3 (every day)

Open `drills/read/core/00_core_read/` (or use the three snippets printed by `go run .`).

For each snippet, do Pass 1–5 from `READING_PATTERNS.md` and fill answers.

| # | Skill | Target |
|---|--------|--------|
| 1 | Signature + state names | 90s |
| 2 | Trace one sample to the return value | 2 min |
| 3 | Name pattern + time complexity | 60s |

**Core checklist**

- [ ] Said signature aloud (in → out)
- [ ] Traced without running the program in your head as “vibes”
- [ ] Named pattern + O(…)

---

## Part B — Specialty (weekday rotation)

| Day | File | Reading skill |
|-----|------|----------------|
| **Mon** | `01_scan_structure` | Map I/O, loops, early exits, mutations |
| **Tue** | `02_trace_execution` | Hand-simulate; track every binder |
| **Wed** | `03_name_the_pattern` | Opaque code → pattern label |
| **Thu** | `04_find_the_bug` | Off-by-one, wrong pointer, missing visit |
| **Fri** | `05_complexity_glance` | Tight bounds from shape |
| **Sat** | `06_reconstruct_ask` | Code → one-sentence problem |
| **Sun** | `07_compare_variants` | Two solutions; name the tradeoff |

---

## Part C — 20 min clock

| Min | Block | Action |
|-----|-------|--------|
| 0–2 | Method flash | Recite 6 passes out loud |
| 2–10 | Core Read 3 | Fill answers in `00_core_read` |
| 10–18 | Specialty | All `TODO: READ` in today’s file |
| 18–20 | Run + log | `go run . -- --run`; note one miss |

---

## Part D — Failure recovery

| Situation | Do |
|-----------|----|
| Trace wrong | Rebuild state table; don’t peek yet |
| Pattern wrong | Re-do Pass 2 skeleton only, then guess again |
| Still stuck after run fails | Peek **that** answer only; re-trace blind |
| All specialty fails | Yesterday’s specialty + Core Read 3 |

---

## Part E — 30-day reading tracker

```
Week 1:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 2:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 3:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
Week 4:  Mon[ ] Tue[ ] Wed[ ] Thu[ ] Fri[ ] Sat[ ] Sun[ ]
```

Log line:

```
2026-07-29 | core OK | 03_name_the_pattern | mistook variable window for fixed | revisit Aug 1
```

Optional longer notes: `docs/reading_log.md`.
