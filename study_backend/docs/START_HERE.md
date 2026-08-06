# Backend interview prep — start here

Your interview is **tomorrow 4:00–5:00 PM**. This pack maps every concept on your resume to timed drills.

## First command

```bash
cd study_backend && go run . -- --cram
```

Shows hour-by-hour schedule until interview time.

## Daily ritual (backend interview)

```
1. Core 5 EXPLAIN  → say answers out loud, fill TODO: EXPLAIN
2. Core 5 WRITE    → implement Go reflex functions from memory
3. Resume block    → one themed block (8 total, covers full resume)
4. STAR story      → rehearse one 90-second story from STAR_STORIES.md
5. Scenario        → 3 min answer to block scenario out loud
```

## Commands

```bash
cd study_backend && go run .                    # today's plan + block
cd study_backend && go run . -- --run           # validate your answers
cd study_backend && go run . -- --micro --run   # Core 5 only (interview-day warm-up)
cd study_backend && go run . -- --catalog       # all 8 resume blocks
cd study_backend && go run . -- --setup         # copy drills to drills/backend/
```

## Files

| Path | Purpose |
|------|---------|
| `docs/INTERVIEW_CRAM_PLAN.md` | Full schedule Aug 6–7 |
| `docs/DRILL_CONCEPTS.md` | Concept deep-dives from resume |
| `docs/STAR_STORIES.md` | 90-sec stories with metrics |
| `drills/backend/explain/` | Verbal concept drills |
| `drills/backend/write/core5/` | Go code reflex |
| `_support/answers/` | Peek after attempt |

## Resume → block map

| Resume highlight | Block |
|------------------|-------|
| MST REST APIs, JWT, MySQL schema | 01, 02 |
| iBind WebRTC, circuit breaker, 40% latency | 03, 04 |
| Opscale ZATCA, Airflow, signing | 05, 07 |
| Jenkins CI/CD, AWS | 06 |
| BitTorrent + HTTP server (Go) | 08 |

Combine with DSA drills if time allows: `cd daily && go run .`
