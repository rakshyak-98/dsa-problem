# Backend interview drills

Resume-driven prep for **Rakshyak Satpathy — Backend Software Engineer**.

## Quick start

```bash
go run ./bin/study_backend                    # today's block + revision plan
go run ./bin/study_backend -- --drill revision  # weekly revision drill
go run ./bin/study_backend -- --run revision    # check revision answers
go run ./bin/study_backend -- --catalog         # full topic + cycle catalog
```

## Layout

```
drills/backend/
├── explain/
│   ├── core5/           # daily Core 5 verbal (JWT, circuit breaker, SQL, etc.)
│   ├── blocks/          # 8 resume-themed blocks (Mon–Sun + bonus Go)
│   └── revision/        # 7 weekly cross-topic revision drills
├── write/
│   └── core5/           # Go reflex implementations
└── scenario/
    └── mock_scenarios/  # system design + behavioral prompts
```

## Weekly revision cycle

| Day | Block | Revision |
|-----|-------|----------|
| Mon | 01 REST/JWT | `01_api_auth_recap` |
| Tue | 02 SQL | `02_data_resilience_recap` |
| Wed | 03 Distributed | `03_distributed_realtime` |
| Thu | 04 WebRTC | `04_realtime_messaging` |
| Fri | 05 Workflows | `05_devops_orchestration` |
| Sat | 06 DevOps | `06_full_week_sweep` |
| Sun | 07 Compliance | `07_go_compliance_mix` (+ block 08 Go) |

Full schedule: [`doc/backend/WEEKLY_REVISION.md`](../doc/backend/WEEKLY_REVISION.md)

## Guides

- [`doc/backend/START_HERE.md`](../doc/backend/START_HERE.md)
- [`doc/backend/WEEKLY_REVISION.md`](../doc/backend/WEEKLY_REVISION.md)
- [`doc/backend/INTERVIEW_CRAM_PLAN.md`](../doc/backend/INTERVIEW_CRAM_PLAN.md)
- [`doc/backend/DRILL_CONCEPTS.md`](../doc/backend/DRILL_CONCEPTS.md)
- [`doc/backend/STAR_STORIES.md`](../doc/backend/STAR_STORIES.md)

Answers (after attempt): `bin/study_backend/_support/answers/`
