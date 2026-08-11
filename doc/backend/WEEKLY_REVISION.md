# Backend weekly revision cycle

Cross-topic revision drills that pair with the 7 weekday resume blocks. Repeat every week until interview answers are automatic.

## Daily ritual

```
1. Core 5 EXPLAIN  →  drills/backend/explain/core5/
2. Core 5 WRITE    →  drills/backend/write/core5/
3. Resume block    →  today's weekday block (Mon–Sun)
4. Revision drill  →  today's revision file (cross-topic recap)
5. Revisit blocks  →  re-run prior blocks from memory (no peeking)
6. STAR story      →  doc/backend/STAR_STORIES.md
```

## Weekly cycle

| Day | Resume block | Revision drill | Revisit (no peeking) |
|-----|--------------|----------------|----------------------|
| Mon | `01_rest_api_jwt` | `01_api_auth_recap` | — |
| Tue | `02_databases_sql` | `02_data_resilience_recap` | block 01 |
| Wed | `03_distributed_resilience` | `03_distributed_realtime` | block 02 |
| Thu | `04_realtime_webrtc` | `04_realtime_messaging` | block 03 |
| Fri | `05_workflows_messaging` | `05_devops_orchestration` | blocks 04–05 |
| Sat | `06_devops_aws` | `06_full_week_sweep` | blocks 01–06 |
| Sun | `07_compliance_security` | `07_go_compliance_mix` | blocks 07 + 08 |

**Bonus block:** `08_go_systems` — Go HTTP server + BitTorrent (Sunday revision + cram eve).

## Topic index (collected)

### Core 5 (every session)

| Topic | Explain | Write |
|-------|---------|-------|
| JWT auth flow | access vs refresh, exp/iss/aud | `parseBearerToken` |
| Circuit breaker | closed/open/half-open | `circuitState` |
| SQL slow query | EXPLAIN → index → pool | — |
| REST vs WebSocket | polling vs duplex | `httpStatusClass` |
| Retry + backoff | jitter, idempotency | `exponentialBackoff`, `isIdempotentMethod` |

### Resume blocks (weekday)

| Block | Topics |
|-------|--------|
| 01 REST/JWT | resource URLs, versioning, pagination, structured errors, JWT |
| 02 SQL | B-tree indexes, composite order, N+1, pooling, EXPLAIN |
| 03 Distributed | timeouts, retries, circuit breaker, CAP, saga vs 2PC |
| 04 WebRTC | WebSocket upgrade, SDP/ICE, STUN/TURN, NAT checklist |
| 05 Workflows | Airflow DAGs, idempotent tasks, RabbitMQ, DLQ |
| 06 DevOps | Docker health checks, Jenkins stages, EC2/Lambda/CDN |
| 07 Compliance | XSD validation, signing, audit trail, replay protection |
| 08 Go systems | goroutines/channels, HTTP server, BitTorrent, context |

### Trigger table (say out loud on Saturday)

| If they ask… | Lead with… |
|--------------|------------|
| Slow SQL | EXPLAIN → index/join → pool/cache |
| API design | Resources, verbs, status codes, errors, versioning |
| Auth | JWT claims; validate signature; refresh flow |
| Outage / flakes | Timeouts → retry/backoff → circuit breaker → fail fast |
| Real-time | WebSocket signaling; WebRTC P2P; STUN/TURN |
| Batch job | Airflow DAG, idempotent tasks, DLQ, audit |
| Deploy | Jenkins stages; Docker; health checks; rollback |
| Go concurrency | Goroutines + channels; context; worker pools |
| Invoice rejected | XSD → sign → submit → audit + payload hash |

### Flashcard decks (spaced repetition)

```bash
go run . -- --track cards --deck=backend --due    # concepts from DRILL_CONCEPTS
go run . -- --track cards --deck=star --due       # STAR stories
go run . -- --track cards --deck=b2b-docker --due # Back2Basics add-ons
```

Relevant b2b decks: `b2b-docker`, `b2b-database`, `b2b-networking`, `b2b-golang`, `b2b-devops`, `b2b-aws`, `b2b-nginx`, `b2b-kubernates`, `b2b-system-design`, `b2b-messaging`, `b2b-security`.

## Commands

```bash
go run . -- --track backend                         # today's block + revision plan
go run . -- --track backend -- --drill revision     # revision drill only
go run . -- --track backend -- --run revision       # check revision answers
go run . -- --track backend -- --run core           # Core 5 explain + write
go run . -- --track backend -- --run reflex         # today's resume block
go run ./bin/study_backend -- --catalog             # blocks + revision cycle + topics
go run -C drills/backend/explain/revision/01_api_auth_recap .
```

## Files

| Path | Purpose |
|------|---------|
| `drills/backend/explain/blocks/` | 8 resume-themed blocks |
| `drills/backend/explain/revision/` | 7 weekly revision drills |
| `drills/backend/explain/core5/` | daily Core 5 verbal |
| `drills/backend/write/core5/` | Go reflex implementations |
| `doc/backend/DRILL_CONCEPTS.md` | concept deep-dives |
| `doc/backend/STAR_STORIES.md` | 90-sec rehearsed stories |
