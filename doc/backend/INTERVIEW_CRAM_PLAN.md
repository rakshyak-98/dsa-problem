# Interview cram plan — Aug 6–7, 2026

**Interview:** Thursday Aug 7, 4:00–5:00 PM  
**Role focus:** Backend Software Engineer (resume: Go, REST, MySQL, WebRTC, Airflow, ZATCA, Jenkins)

---

## Today (Aug 6) — cover full resume

### Session 1 · 2 hours · APIs + databases

| Min | Activity |
|-----|----------|
| 0–15 | Read `DRILL_CONCEPTS.md` § REST + § Databases (skim headers) |
| 15–30 | Core 5 EXPLAIN — out loud, then fill `explain/core5` |
| 30–50 | Core 5 WRITE — `write/core5` from memory |
| 50–75 | Block `01_rest_api_jwt` + scenario: page publish API |
| 75–105 | Block `02_databases_sql` + scenario: 35% reporting win |
| 105–120 | STAR: MST schema versioning story |

```bash
go run ./bin/study_backend -- --run
go run -C drills/backend/explain/blocks/01_rest_api_jwt .
go run -C drills/backend/explain/blocks/02_databases_sql .
```

### Session 2 · 2 hours · Distributed + real-time

| Min | Activity |
|-----|----------|
| 0–45 | Block `03_distributed_resilience` — draw circuit breaker states |
| 45–90 | Block `04_realtime_webrtc` — draw signaling vs media path |
| 90–105 | STAR: iBind 40% latency reduction |
| 105–120 | Mock Q: "Corporate NAT blocks video — debug steps?" |

### Session 3 · 2 hours · Pipelines + DevOps + compliance

| Min | Activity |
|-----|----------|
| 0–40 | Block `05_workflows_messaging` — Airflow idempotency |
| 40–75 | Block `06_devops_aws` — walk Jenkins pipeline stages |
| 75–110 | Block `07_compliance_security` — ZATCA signing flow |
| 110–120 | STAR: invoice pipeline 25% faster |

### Session 4 · 1.5 hours · Go systems + stories

| Min | Activity |
|-----|----------|
| 0–45 | Block `08_go_systems` — HTTP middleware chain + BitTorrent pieces |
| 45–75 | Rehearse all STAR stories at 90 sec each |
| 75–90 | List 5 questions to ask interviewer |

**Sleep by midnight.**

---

## Tomorrow (Aug 7) — interview day

### Morning · 45 min

```bash
go run ./bin/study_backend -- --micro --run
```

Skim `DRILL_CONCEPTS.md` trigger table only.

### Late morning · 45 min

- Block 03 + 04 explain drills
- WebRTC NAT debug checklist (say out loud)

### Pre-lunch · 30 min

- Block 05 + 07 quick review
- ZATCA rejection trace: validate → sign → submit → audit log

### 15:30–15:55 · warm-up

1. Core 5 EXPLAIN without peeking
2. One STAR with metrics (pick strongest: 35% reporting or 40% latency)
3. Water, quiet room, test audio if remote

### 16:00–17:00 · interview

**Answer structure:** Context (10s) → Your action (60s) → Result/metric (20s)

**System design:** Boxes + arrows; label sync vs async, failure modes.

**If stuck:** "Let me structure this…" then inputs/outputs/constraints.

---

## Concept checklist (tick before interview)

- [ ] JWT flow + where validated
- [ ] Circuit breaker 3 states
- [ ] 3 SQL optimization levers (index, rewrite, pool/cache)
- [ ] WebRTC: signaling vs media; STUN vs TURN
- [ ] Airflow task idempotency
- [ ] Jenkins pipeline stages you built
- [ ] ZATCA: XSD → sign → submit
- [ ] Go: goroutine + graceful shutdown
- [ ] REST error shape + status codes
- [ ] CAP during partition

---

## Optional DSA (if energy remains)

```bash
go run . -- --micro
```

One easy hashing problem only — don't burn out before interview.
