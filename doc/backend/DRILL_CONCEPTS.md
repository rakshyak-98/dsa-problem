# Backend drill concepts — mapped to your resume

Quick reference for interview answers. Pair with explain drills in `drills/backend/explain/`.

---

## Trigger table (if they ask X → you say Y)

| If they ask… | Lead with… |
|--------------|------------|
| Slow SQL query | EXPLAIN plan → missing index / bad join → pool or cache |
| API design | Resource URLs, verbs, status codes, error contract, versioning |
| Auth | JWT claims (exp, iss, aud); validate signature; refresh flow |
| Outage / flakes | Timeouts → retry w/ backoff → circuit breaker → fail fast |
| Real-time chat/video | WebSocket signaling; WebRTC P2P media; STUN/TURN for NAT |
| Batch compliance job | Airflow DAG, idempotent tasks, DLQ, audit trail |
| Deploy process | Jenkins stages; Docker image; health checks; rollback |
| Go concurrency | Goroutines + channels; context for cancel; worker pools |
| Invoice rejected | XSD validate → sign → submit → correlate audit + payload hash |

---

## REST API design (MST, iBind)

- **Resources:** nouns, plural (`/pages`, `/pages/{id}/versions`)
- **Status codes:** 200 OK, 201 Created, 400 client, 401 unauth, 403 forbidden, 404 not found, 409 conflict, 500 server
- **Errors:** `{ "code": "PAGE_NOT_FOUND", "message": "...", "details": {} }`
- **Versioning:** `/v1/...` or `Accept-Version` header
- **Pagination:** `limit` + `cursor` (stable) vs offset (simple)
- **Idempotency:** PUT/DELETE idempotent; POST needs `Idempotency-Key` for retries

### JWT

1. Client logs in → server issues access (+ optional refresh) token
2. Client sends `Authorization: Bearer <token>`
3. Middleware validates signature, `exp`, `iss`, `aud`; attaches claims to request context

---

## Databases (MST MySQL, iBind pooling)

### Indexing

- B-tree default for `WHERE`, `JOIN`, `ORDER BY` left-prefix
- Composite index column order: equality filters first, then range
- Covering index: includes all columns in SELECT → index-only scan

### Query optimization workflow

1. `EXPLAIN ANALYZE` — find full table scans
2. Rewrite query (subquery → join, remove SELECT *)
3. Add/adjust indexes
4. Cache hot reads (Redis); tune pool

### Connection pooling

- Reuses TCP + auth handshakes
- Size ≈ `(cores * 2) + spindle` rule of thumb; measure under load
- Watch pool exhaustion → latency spikes

### Schema design (page builder)

- Versioned entities: `pages`, `page_versions`, `audit_log`
- RBAC: roles → permissions → resource scopes
- Publishing: draft vs published snapshots

---

## Distributed systems (iBind)

### Retry + exponential backoff

```
delay = min(cap, base * 2^attempt) + jitter
```

- Only retry **transient** errors (timeouts, 503)
- Require **idempotent** operations or idempotency keys

### Circuit breaker states

```
CLOSED (normal) → failures ≥ threshold → OPEN (fail fast)
OPEN → after timeout → HALF-OPEN (probe)
HALF-OPEN → success → CLOSED; failure → OPEN
```

### CAP (during network partition)

- Choose **CP** (consistency) or **AP** (availability)
- Most web APIs favor availability + eventual consistency for reads

### Saga vs 2PC

- 2PC: strong consistency, fragile under partitions
- Saga: local transactions + compensating actions; common in microservices

---

## WebSocket & WebRTC (iBind)

### WebSocket

- HTTP Upgrade handshake → full-duplex TCP
- Use for signaling, chat, live notifications
- Heartbeats / ping to detect dead connections

### WebRTC architecture

```
Signaling (WebSocket): SDP offer/answer, ICE candidates
Media (UDP): encrypted RTP between peers (not through your server)
```

### STUN vs TURN

| | STUN | TURN |
|---|------|------|
| Purpose | Discover public IP/port | Relay when P2P fails |
| Cost | Low | Higher (bandwidth on server) |
| When | Most home networks | Symmetric NAT, strict corporate firewalls |

### NAT debug checklist

1. Signaling connected? (WebSocket open)
2. ICE candidates gathered? (host, srflx, relay)
3. Connection state `connected` vs `failed`?
4. Try TURN if only host candidates
5. Firewall allows UDP ports / TURN TLS

---

## Workflows & messaging (Opscale, RabbitMQ)

### Apache Airflow

- **DAG:** directed acyclic graph of tasks
- **Operators:** Python tasks with dependencies
- **Retries:** `retries`, `retry_delay`; tasks must be **idempotent**
- **Sensors:** wait for external condition
- Recovery: clear failed task, rerun from failed node

### RabbitMQ patterns

- **Work queue:** competing consumers, ack after processing
- **Pub/sub:** fanout to many subscribers
- **DLQ:** poison messages after max retries

---

## DevOps & AWS (iBind Jenkins)

### Jenkins pipeline stages

```
checkout → install deps → lint → unit test → build artifact
→ integration test → security scan → deploy staging → smoke test → deploy prod
```

### Docker

- Image layers; multi-stage builds for small images
- Health checks: liveness (restart) vs readiness (traffic)

### AWS quick map

| Service | Use |
|---------|-----|
| EC2 | Long-running services, full control |
| Lambda | Short event-driven handlers |
| CDN (CloudFront) | Static assets, edge caching |
| S3 | Artifacts, backups |

---

## Compliance & security (ZATCA)

### Invoice submission flow

1. Generate invoice XML from business data
2. **XSD validation** against authority schema
3. **Cryptographic signing** (hash + private key)
4. Submit to ZATCA API over TLS
5. Store response + audit log (who, when, payload hash)

### Replay protection

- Unique invoice ID, timestamp, nonce in signed payload

---

## Go systems (projects)

### HTTP server from scratch

- `net/http` or raw: accept loop → goroutine per conn (or pool)
- Parse request → middleware chain → handler → response
- Keep-alive: reuse connection for multiple requests on same TCP
- Graceful shutdown: `Shutdown(ctx)` stops listener, drains in-flight

### BitTorrent client

- **Tracker (BEP 15):** UDP announce → peer list
- **bencode:** binary encoding for torrent metadata
- **Wire protocol:** handshake, choke/unchoke, request blocks
- **Concurrency:** per-peer goroutines, piece bitmap, work scheduling

### Concurrency primitives

- `goroutine` + `channel` for worker pools
- `context.Context` for cancellation and deadlines
- `sync.Mutex` for shared state; prefer channels for ownership transfer

---

## How to study this file

1. Read one section → run matching explain block
2. Say the trigger table row out loud without reading
3. Draw architecture for WebRTC and ZATCA flows on paper
4. Rehearse STAR story for that section
