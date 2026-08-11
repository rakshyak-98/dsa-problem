# Backend interview prep — start here

Your interview is **tomorrow 4:00–5:00 PM**. This pack maps every concept on your resume to timed drills.

## First command

```bash
go run ./bin/study_backend -- --cram
```

Shows hour-by-hour schedule until interview time.

## Daily ritual (backend interview)

```
1. Core 5 EXPLAIN  → say answers out loud, fill TODO: EXPLAIN
2. Core 5 WRITE    → implement Go reflex functions from memory
3. Resume block    → one themed block (7 weekdays + bonus Go block)
4. Revision drill  → cross-topic recap (weekly cycle)
5. Revisit blocks  → re-run prior blocks from memory (see WEEKLY_REVISION.md)
6. STAR story      → rehearse one 90-second story from STAR_STORIES.md
7. Scenario        → 3 min answer to block scenario out loud
```

## Commands

```bash
go run ./bin/study_backend                    # today's plan + block + revision
go run ./bin/study_backend -- --run           # validate all answers (core + block)
go run ./bin/study_backend -- --run core      # Core 5 only (interview-day warm-up)
go run ./bin/study_backend -- --run reflex    # today's resume block
go run ./bin/study_backend -- --run revision  # today's weekly revision drill
go run ./bin/study_backend -- --drill revision  # revision drill path only
go run ./bin/study_backend -- --catalog       # blocks + revision cycle + topics
go run ./bin/study_backend -- --setup         # verify drills/backend/ layout
go run . -- --track backend                   # same plan via unified runner
```

## Files

| Path | Purpose |
|------|---------|
| `doc/backend/WEEKLY_REVISION.md` | Weekly cycle, topic index, revisit schedule |
| `doc/backend/INTERVIEW_CRAM_PLAN.md` | Full schedule Aug 6–7 |
| `doc/backend/DRILL_CONCEPTS.md` | Concept deep-dives from resume |
| `doc/backend/STAR_STORIES.md` | 90-sec stories with metrics |
| `drills/backend/explain/blocks/` | 8 resume-themed blocks |
| `drills/backend/explain/revision/` | 7 weekly revision drills |
| `drills/backend/write/core5/` | Go code reflex |
| `bin/study_backend/_support/answers/` | Peek after attempt |

## Resume → block map

| Resume highlight | Block |
|------------------|-------|
| MST REST APIs, JWT, MySQL schema | 01, 02 |
| iBind WebRTC, circuit breaker, 40% latency | 03, 04 |
| Opscale ZATCA, Airflow, signing | 05, 07 |
| Jenkins CI/CD, AWS | 06 |
| BitTorrent + HTTP server (Go) | 08 |

Combine with DSA drills if time allows: `go run .`
