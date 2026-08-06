# STAR stories — rehearse at 90 seconds each

Use **Situation → Task → Action → Result** with metrics from your resume.

---

## 1. Reporting query performance (+35%) — MST

**Situation:** High-volume analytics queries on the no-code page builder were slow, blocking reporting dashboards.  
**Task:** Improve reporting query performance without breaking schema versioning or audit history.  
**Action:** Analyzed slow queries with EXPLAIN; added composite indexes on filter columns; refactored joins to reduce row scans; normalized hot paths while keeping denormalized read models where needed.  
**Result:** **35% faster** reporting queries; dashboards stayed accurate with governance and publishing workflows intact.

**Keywords:** MySQL, indexing, SQL refactoring, relational modeling, analytics.

---

## 2. API latency reduction (-40%) — iBind

**Situation:** Real-time video platform APIs were too slow under concurrent load.  
**Task:** Cut API latency for WebSocket-adjacent REST endpoints serving session and user state.  
**Action:** Profiled request path; optimized SQL (removed N+1); tuned connection pool size; added Redis caching for hot reads; trimmed middleware overhead.  
**Result:** **40% lower API latency**; smoother real-time experience alongside WebRTC media.

**Keywords:** query optimization, connection pooling, caching, real-time workloads.

---

## 3. Resilient inter-service communication — iBind

**Situation:** Transient downstream failures caused cascading errors in the video platform.  
**Task:** Redesign inter-service calls for resilience without hiding real outages.  
**Action:** Introduced exponential backoff retries for idempotent calls; circuit breakers to fail fast when downstream unhealthy; structured errors for clients; timeouts on every outbound call.  
**Result:** Isolated transient failures; improved maintainability and integration reliability.

**Keywords:** circuit breaker, exponential backoff, distributed systems, JWT REST APIs.

---

## 4. ZATCA invoice pipeline (-25% processing time) — Opscale

**Situation:** Saudi ZATCA e-invoicing required XML validation, cryptographic signing, and secure submission at high volume.  
**Task:** Keep compliance while reducing end-to-end invoice processing time.  
**Action:** Designed Airflow DAGs with retry policies and idempotent tasks; batched workloads; parallelized independent stages; automated recovery on failure.  
**Result:** **25% faster** invoice processing; fewer pipeline failures; compliance maintained.

**Keywords:** Apache Airflow, XML validation, cryptographic signing, idempotency, batching.

---

## 5. WebRTC across NAT — iBind

**Situation:** Users on NAT-restricted networks could not establish stable video calls.  
**Task:** Deliver P2P video with stable connectivity across corporate and mobile NATs.  
**Action:** Integrated WebSocket signaling; WebRTC with STUN for address discovery; TURN fallback when direct P2P blocked; tuned ICE gathering and connection state handling.  
**Result:** Stable conferencing across NAT-restricted networks.

**Keywords:** WebRTC, WebSocket, STUN, TURN, peer-to-peer.

---

## 6. Jenkins CI/CD — iBind

**Situation:** Manual releases were slow and error-prone.  
**Task:** Automate build, test, and deployment for faster repeatable releases.  
**Action:** Set up Jenkins pipelines: checkout → build → unit tests → artifact → deploy to staging/prod with gates; Docker images for consistent runtime.  
**Result:** Faster releases, repeatable deployments, improved team velocity.

**Keywords:** Jenkins, CI/CD, Docker, automation.

---

## Questions to ask them

1. What does the team's on-call/incident process look like?
2. How do you balance schema evolution vs zero-downtime deploys?
3. What's the primary datastore and expected query patterns?
4. How much of the stack is Go vs Node, and why?
