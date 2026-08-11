#!/usr/bin/env python3
"""Rebuild backend interview flashcards from doc/backend/ and drills/backend/.

Writes question-style cards to cards/decks/backend.json and star.json.
Run from repo root: python3 cards/generate_cards.py
"""
from __future__ import annotations

import hashlib
import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
DOC = ROOT / "doc"
DRILLS = ROOT / "drills" / "backend"
OUT = ROOT / "cards" / "decks"
cards: list[dict] = []

EXPLAIN_RE = re.compile(r"//\s*TODO:\s*EXPLAIN\s*—\s*(.+)$")
BLOCK_HEADER_RE = re.compile(r"^//\s*BLOCK\s+(\d+)\s*—\s*(.+)$")
REVISION_HEADER_RE = re.compile(r"^//\s*REVISION\s+(\d+)\s*—\s*(.+)$")
CORE5_HEADER_RE = re.compile(r"^//\s*CORE\s+5\s+EXPLAIN")
SECTION_COMMENT_RE = re.compile(r"^//\s*Q(\d+)\s*—\s*(.+)$")

BLOCK_TOPIC_TAGS: dict[str, list[str]] = {
    "01_rest_api_jwt": ["api", "jwt", "block01"],
    "02_databases_sql": ["db", "block02"],
    "03_distributed_resilience": ["resilience", "block03"],
    "04_realtime_webrtc": ["webrtc", "block04"],
    "05_workflows_messaging": ["airflow", "messaging", "block05"],
    "06_devops_aws": ["devops", "aws", "block06"],
    "07_compliance_security": ["zatca", "compliance", "block07"],
    "08_go_systems": ["go", "block08"],
}

REVISION_TOPIC_TAGS: dict[str, list[str]] = {
    "01_api_auth_recap": ["api", "jwt", "revision"],
    "02_data_resilience_recap": ["db", "resilience", "revision"],
    "03_distributed_realtime": ["resilience", "webrtc", "revision"],
    "04_realtime_messaging": ["webrtc", "messaging", "revision"],
    "05_devops_orchestration": ["devops", "revision"],
    "06_full_week_sweep": ["revision", "trigger"],
    "07_go_compliance_mix": ["go", "zatca", "revision"],
}


def cid(*parts: str) -> str:
    return hashlib.sha1("|".join(parts).encode()).hexdigest()[:12]


def clean(s: str) -> str:
    s = re.sub(r"\*\*([^*]+)\*\*", r"\1", s)
    s = re.sub(r"`([^`]+)`", r"\1", s)
    return re.sub(r"\s+", " ", s).strip()


def add(deck: str, tags: list[str], front: str, back: str, source: str, section: str = "") -> None:
    front, back = clean(front), clean(back)
    if not front or not back or front == back or len(front) < 8 or len(back) < 2:
        return
    if re.fullmatch(r"What is \d+\?", front):
        return
    tag_set = sorted(set(tags + [deck]))
    cards.append(
        {
            "id": cid(deck, front),
            "deck": deck,
            "tags": tag_set,
            "front": front,
            "back": back,
            "source": source,
            "section": section,
        }
    )


def parse_explain_prompt(body: str) -> tuple[str, str] | None:
    quotes = re.findall(r'"([^"]+)"', body)
    if not quotes:
        return None

    # Question: text before the first colon that introduces quoted answer(s)
    colon_idx = body.find(":")
    if colon_idx > 0 and quotes:
        q = body[:colon_idx].strip()
        if not q.endswith("?"):
            q = q + "?"
    else:
        q = re.sub(r"\([^)]*\)", "", body)
        q = re.sub(r'"[^"]*"', "", q).strip()
        if not q.endswith("?"):
            q = q + "?"

    back = " or ".join(quotes)
    for note in re.findall(r"\(([^)]+)\)", body):
        low = note.lower()
        if low.startswith("e.g.") or low.startswith("either ok"):
            continue
        if note not in back:
            back = f"{back} ({note})"

    return q, back


def parse_assert_hints(text: str) -> list[str]:
    """Extract the hint string (last quoted literal) from each assert line."""
    hints: list[str] = []
    for line in text.splitlines():
        stripped = line.strip()
        if not stripped.startswith("assert("):
            continue
        quoted = re.findall(r'"([^"]*)"', line)
        if quoted:
            hints.append(quoted[-1])
    return hints


def prompt_key(body: str) -> str:
    m = re.match(r'^(.+?):\s*"', body)
    if m:
        key = clean(m.group(1))
    else:
        key = re.sub(r'"([^"]*)"', r"\1", body)
        key = re.sub(r"\([^)]*\)", "", key)
        key = clean(key)
    key = re.sub(r"\bkeyword\b", "", key, flags=re.I)
    key = re.sub(r"\bone word:?\b", "", key, flags=re.I)
    return re.sub(r"\s+", " ", key).lower().rstrip("?:").strip()


def cap_sentence(s: str) -> str:
    s = clean(s)
    if not s:
        return s
    if s[0].islower():
        s = s[0].upper() + s[1:]
    if s[-1] not in ".!?":
        s += "."
    return s


# Drill prompt key → (interview question, definition-style answer).
DRILL_QA: dict[str, tuple[str, str]] = {
    "http status for resource created": (
        "What HTTP status code indicates successful resource creation?",
        "HTTP 201 Created. Returned when a new resource is successfully created on the server.",
    ),
    "safe retry for update without side effects": (
        "Which REST verb is safe to retry for updates without side effects?",
        "PUT is idempotent; the same request can be retried without duplicate side effects.",
    ),
    "jwt segments count": (
        "How many segments does a JWT contain?",
        "Three segments — header, payload, and signature — separated by dots.",
    ),
    "validate jwt signature with what key type for rs256": (
        "What key type verifies an RS256 JWT signature?",
        "The issuer's public key. RS256 is asymmetric — sign with private key, verify with public key.",
    ),
    "api versioning strategy you used": (
        "How do you version a REST API?",
        "URL path prefix (/v1/...) or Accept-Version header. Both keep clients on stable contracts.",
    ),
    "default mysql innodb index structure": (
        "What is the default index structure in MySQL InnoDB?",
        "B-tree. Supports equality and range scans on WHERE, JOIN, and ORDER BY columns.",
    ),
    "composite index column order matters for left-prefix rule": (
        "Does column order matter in a composite index?",
        "Yes — composite indexes follow the left-prefix rule; leading columns must match the query filter order.",
    ),
    "n+1 problem fixed by": (
        "What is the N+1 query problem? How do you fix it?",
        "One query fetches parents, then N queries fetch each child. Fix with JOINs, eager loading, or batch fetching.",
    ),
    "connection pool avoids repeated": (
        "What does a database connection pool avoid?",
        "Repeated TCP and authentication handshakes. Reuses open connections across requests.",
    ),
    "explain shows": (
        "What does EXPLAIN show?",
        "The query execution plan — which indexes are used, join order, and estimated row counts.",
    ),
    "exponential backoff multiplier base commonly": (
        "What is the common base multiplier for exponential backoff?",
        "2 — double the delay each retry attempt. Add jitter to prevent synchronized retries.",
    ),
    "circuit breaker prevents": (
        "What does a circuit breaker prevent?",
        "Cascading failures. Halts requests to an unhealthy downstream after a failure threshold.",
    ),
    "idempotency key sent in": (
        "Where is an idempotency key sent?",
        "Request header (Idempotency-Key). Ensures duplicate POST retries produce the same result.",
    ),
    "cap": (
        "During a network partition, what does CAP force you to sacrifice?",
        "Consistency or availability. Most web APIs favor availability with eventual consistency for reads.",
    ),
    "saga pattern coordinates via": (
        "How does the saga pattern coordinate distributed transactions?",
        "Local transactions plus compensating actions, coordinated via events — choreography or orchestration.",
    ),
    "webrtc media path is": (
        "What path does WebRTC media take?",
        "Peer-to-peer over encrypted UDP. Media does not flow through your signaling server.",
    ),
    "signaling carries": (
        "What does WebRTC signaling carry?",
        "SDP offer/answer and ICE candidates, typically over WebSocket.",
    ),
    "ice gathers": (
        "What does ICE gather?",
        "Connection candidates — host, server-reflexive (STUN), and relay (TURN) — for NAT traversal.",
    ),
    "stun discovers public": (
        "What does STUN discover?",
        "The client's public IP and port for NAT hole punching.",
    ),
    "turn used when direct p2p": (
        "When is TURN used in WebRTC?",
        "When direct peer-to-peer fails — symmetric NAT, corporate VPNs, or strict firewalls.",
    ),
    "airflow unit of work": (
        "What is the unit of work in Apache Airflow?",
        "A task (operator). Tasks are wired together as nodes in a DAG.",
    ),
    "dag means": (
        "What does DAG mean in Airflow?",
        "Directed acyclic graph — tasks with dependency edges and no cycles.",
    ),
    "failed task safe rerun needs": (
        "What must Airflow tasks be for safe reruns after failure?",
        "Idempotent. Clearing and rerunning a failed node must not duplicate side effects.",
    ),
    "rabbitmq dead letter queue for": (
        "What is a RabbitMQ dead-letter queue for?",
        "Poison messages that fail after max retries. Isolates them without blocking the main queue.",
    ),
    "work queue pattern": (
        "What is the RabbitMQ work-queue pattern?",
        "Competing consumers pull jobs from a shared queue; ack only after successful processing.",
    ),
    "docker image built from": (
        "What is a Docker image built from?",
        "Stacked filesystem layers. Multi-stage builds keep production images small.",
    ),
    "ci stage before deploy": (
        "What runs in CI before deploy?",
        "Tests, lint, and security scans. Artifact only ships after the pipeline passes.",
    ),
    "lambda best for": (
        "When is AWS Lambda the best fit?",
        "Short, event-driven handlers. Not suited for long-running or always-on services.",
    ),
    "cdn caches": (
        "What does a CDN cache?",
        "Static assets at the edge. Cuts latency and reduces load on the origin server.",
    ),
    "container health check probes": (
        "What are liveness vs readiness health probes?",
        "Liveness restarts a dead container; readiness gates traffic until the instance can serve requests.",
    ),
    "xsd validates": (
        "What does XSD validation check?",
        "XML structure against the authority schema. Required before signing and submission.",
    ),
    "digital signature proves": (
        "What does a digital signature prove?",
        "Integrity (payload unchanged) and authenticity (trusted signer).",
    ),
    "signing hash algorithm example": (
        "What hash algorithm is commonly used for digital signing?",
        "SHA-256. Hash the payload, then sign the digest with the private key.",
    ),
    "replay attack prevented by": (
        "How do you prevent replay attacks on signed payloads?",
        "Unique ID, timestamp, and nonce in the signed payload.",
    ),
    "audit trail stores": (
        "What does an audit trail store?",
        "Who changed what and when — actor, timestamp, and payload hash.",
    ),
    "goroutine scheduling model": (
        "What is Go's goroutine scheduling model?",
        "M:N multiplexing — many goroutines scheduled onto fewer OS threads by the runtime.",
    ),
    "context.context used for": (
        "What is context.Context used for in Go?",
        "Cancellation, deadlines, and request-scoped values propagated across goroutines.",
    ),
    "bittorrent metadata encoding": (
        "How is BitTorrent metadata encoded?",
        "Bencode — a compact binary format for .torrent file structures.",
    ),
    "http keep-alive reuses": (
        "What does HTTP keep-alive reuse?",
        "The same TCP connection for multiple HTTP requests on one socket.",
    ),
    "graceful shutdown closes": (
        "What does graceful HTTP shutdown do?",
        "Stops accepting new connections, drains in-flight requests, then exits.",
    ),
    "three parts of a jwt, separated by what character": (
        "What separates the three parts of a JWT?",
        "Dots. Format: header.payload.signature (Base64URL-encoded).",
    ),
    "where is jwt usually sent? one word": (
        "Where is a JWT typically sent in HTTP?",
        "Authorization header as Bearer <token>.",
    ),
    "claim that enforces expiry": (
        "Which JWT claim enforces token expiry?",
        "exp. Reject any token whose exp timestamp is in the past.",
    ),
    "three states": (
        "What are the three circuit breaker states?",
        "Closed (normal) → open (fail fast) → half-open (probe). Success closes; failure reopens.",
    ),
    "open circuit behavior": (
        "What is open-circuit behavior?",
        "Fail fast — return an error immediately without calling the unhealthy downstream.",
    ),
    "index helps which clause? keyword \"where\"": (
        "Which SQL clause does an index primarily accelerate?",
        "WHERE (and JOIN keys). Composite indexes also follow the left-prefix rule.",
    ),
    "pooling reduces what? keyword \"connection\"": (
        "What does connection pooling reduce?",
        "Per-request connection overhead. Reuses open DB connections across requests.",
    ),
    "websocket starts as http \"upgrade\"": (
        "How does a WebSocket connection start?",
        "HTTP Upgrade handshake, then switches to full-duplex TCP.",
    ),
    "server-push chat transport": (
        "What transport enables server-push for real-time chat?",
        "WebSocket — full-duplex TCP for signaling, chat, and live notifications.",
    ),
    "jitter prevents \"thundering herd\"": (
        "What does jitter prevent in retry logic?",
        "Thundering herd — randomizing delays stops all clients from retrying simultaneously.",
    ),
    "retries need \"idempotent\" operations": (
        "What kind of operations are safe to retry?",
        "Idempotent operations. For writes, use an Idempotency-Key header.",
    ),
    "http status for successful get": (
        "What HTTP status code indicates a successful GET?",
        "200 OK. Standard response for a successful read with a response body.",
    ),
    "idempotent update verb": (
        "Which HTTP verb is idempotent for updates?",
        "PUT. Safe to retry without creating duplicate side effects.",
    ),
    "jwt sent in http": (
        "Where is a JWT sent in an HTTP request?",
        "Authorization header as Bearer <token>.",
    ),
    "safe post retry header": (
        "Which header makes POST retries safe?",
        "Idempotency-Key. Duplicate requests with the same key produce the same result.",
    ),
    "api version in url prefix": (
        "How is API versioning commonly done in the URL?",
        "Path prefix — e.g. /v1/pages. Keeps version visible and easy to route.",
    ),
    "first step for slow sql": (
        "What is the first step when debugging a slow SQL query?",
        "EXPLAIN ANALYZE. Inspect the execution plan before changing schema or queries.",
    ),
    "index helps which clause": (
        "Which SQL clause does an index primarily accelerate?",
        "WHERE and JOIN keys. Mind column order on composite indexes (left-prefix rule).",
    ),
    "pool reuses": (
        "What does a connection pool reuse?",
        "Open database connections. Avoids TCP and auth handshake on every query.",
    ),
    "circuit breaker open behavior": (
        "What happens when a circuit breaker is open?",
        "Fail fast — reject calls immediately without hitting the failing dependency.",
    ),
    "backoff multiplier base": (
        "What is the typical exponential backoff multiplier?",
        "Base 2 — double the delay each attempt. Add jitter to spread retry timing.",
    ),
    "during partition sacrifice consistency or availability": (
        "During a network partition, what do you sacrifice?",
        "Consistency or availability (CAP). Most web APIs favor availability.",
    ),
    "microservices distributed tx pattern": (
        "What pattern handles distributed transactions in microservices?",
        "Saga — local transactions plus compensating actions. Prefer over 2PC under partitions.",
    ),
    "jitter prevents": (
        "What does jitter prevent in retry logic?",
        "Thundering herd — randomized delays stop synchronized retries across clients.",
    ),
    "websocket starts as http": (
        "How does a WebSocket connection begin?",
        "HTTP Upgrade request, then full-duplex TCP for bidirectional messaging.",
    ),
    "idempotency key location": (
        "Where is an idempotency key sent?",
        "Request header (Idempotency-Key). Critical for safe POST retries.",
    ),
    "relay when p2p fails": (
        "What relays WebRTC media when P2P fails?",
        "TURN server. Required when NAT or firewall blocks direct peer paths.",
    ),
    "signaling transport": (
        "What transport carries WebRTC signaling?",
        "WebSocket — exchanges SDP offer/answer and ICE candidates.",
    ),
    "airflow graph type": (
        "What graph type does Airflow use?",
        "DAG — directed acyclic graph of dependent tasks.",
    ),
    "poison messages go to": (
        "Where do poison messages go after max retries?",
        "Dead-letter queue (DLQ). Isolates failures without blocking the main queue.",
    ),
    "competing consumers pattern": (
        "What is the competing consumers pattern?",
        "Multiple workers pull from one queue; each message processed by exactly one consumer.",
    ),
    "docker health: traffic gate uses": (
        "Which Docker/K8s probe gates traffic to a container?",
        "Readiness probe. Endpoint removed from load balancing until ready.",
    ),
    "docker health: restart probe": (
        "Which Docker/K8s probe restarts a container?",
        "Liveness probe. Restarts the container if the process is hung or dead.",
    ),
    "short event handler on aws": (
        "What AWS service fits short event-driven handlers?",
        "Lambda. Billed per invocation; not for long-running workloads.",
    ),
    "static asset edge cache": (
        "What caches static assets at the edge?",
        "CDN (e.g. CloudFront). Serves files close to users, reducing origin load.",
    ),
    "pipeline runs tests before deploy": (
        "What runs before deploy in a typical pipeline?",
        "CI — tests, lint, security scan. Deploy only after the pipeline passes.",
    ),
    "zatca step before submit": (
        "What ZATCA step happens before API submission?",
        "Cryptographic signing — hash payload with SHA-256, sign with private key.",
    ),
    "schema validation format": (
        "What format validates invoice XML schema?",
        "XSD. Validate structure against the authority schema before signing.",
    ),
    "go cancellation primitive": (
        "What is Go's cancellation primitive?",
        "context.Context — propagates cancel signals and deadlines across goroutines.",
    ),
    "http graceful stop method": (
        "How do you gracefully stop an HTTP server in Go?",
        "http.Server.Shutdown(ctx) — stop accepting, drain in-flight, then exit.",
    ),
    "index helps which clause? where": (
        "Which SQL clause does an index primarily accelerate?",
        "WHERE (and JOIN keys). Composite indexes follow the left-prefix rule.",
    ),
    "pooling reduces what? connection": (
        "What does connection pooling reduce?",
        "Per-request connection overhead. Reuses open DB connections across requests.",
    ),
    "retries need idempotent operations": (
        "What kind of operations are safe to retry?",
        "Idempotent operations. For writes, use an Idempotency-Key header.",
    ),
}


def lookup_drill_qa(body: str, hint: str = "") -> tuple[str, str]:
    key = prompt_key(body)
    if key in DRILL_QA:
        return DRILL_QA[key]
    # Fallback: build question from stem, answer from hint
    colon = body.find(":")
    stem = body[:colon].strip() if colon > 0 else body
    question = format_fallback_question(stem)
    answer = hint or re.findall(r'"([^"]+)"', body)[0] if re.findall(r'"([^"]+)"', body) else stem
    return question, cap_sentence(answer)


def format_fallback_question(stem: str) -> str:
    s = clean(stem).rstrip("?")
    low = s.lower()
    if low.startswith(("what ", "how ", "when ", "which ", "define ", "why ")):
        return s + "?" if not s.endswith("?") else s
    if low.startswith("explain "):
        return cap_sentence(s).rstrip(".") + "?"
    return f"What is {s}?"


def drill_card(body: str, hint: str = "") -> tuple[str, str]:
    return lookup_drill_qa(body, hint)


TRIGGER_QA: dict[str, tuple[str, str]] = {
    "slow sql query": (
        "How would you debug a slow SQL query?",
        "EXPLAIN plan first → check for missing index or bad join → tune pool or add cache.",
    ),
    "api design": (
        "How do you approach REST API design?",
        "Resource URLs, HTTP verbs, status codes, structured error contract, and versioning.",
    ),
    "auth": (
        "How do you implement API authentication with JWT?",
        "Validate signature and claims (exp, iss, aud); support refresh token flow.",
    ),
    "outage / flakes": (
        "How do you handle downstream outages and flaky services?",
        "Timeouts → retry with exponential backoff → circuit breaker → fail fast.",
    ),
    "real-time chat/video": (
        "How would you architect real-time chat or video?",
        "WebSocket for signaling; WebRTC for P2P media; STUN/TURN for NAT traversal.",
    ),
    "batch compliance job": (
        "How do you design a batch compliance pipeline?",
        "Airflow DAG with idempotent tasks, dead-letter queue, and full audit trail.",
    ),
    "deploy process": (
        "What does your deploy process look like?",
        "Jenkins pipeline stages; Docker image; health checks; rollback plan.",
    ),
    "go concurrency": (
        "What Go concurrency primitives do you use in servers?",
        "Goroutines and channels; context for cancellation; worker pools for bounded parallelism.",
    ),
    "invoice rejected": (
        "An invoice was rejected by the authority — how do you trace it?",
        "XSD validate → sign → submit → correlate audit log and payload hash.",
    ),
}


def trigger_card(term: str, defin: str) -> tuple[str, str]:
    key = term.lower().strip()
    if key in TRIGGER_QA:
        return TRIGGER_QA[key]
    return f"How would you approach {term}?", cap_sentence(defin)


CONCEPT_QA: dict[str, tuple[str, str]] = {
    "resources": (
        "What are REST resource naming conventions?",
        "Nouns, plural paths — e.g. /pages, /pages/{id}/versions.",
    ),
    "status codes": (
        "What HTTP status codes should a REST API use?",
        "200 OK, 201 Created, 400 client error, 401 unauth, 403 forbidden, 404 not found, 409 conflict, 500 server error.",
    ),
    "errors": (
        "What shape should structured API errors take?",
        '{ "code": "PAGE_NOT_FOUND", "message": "...", "details": {} } — machine-readable code plus human message.',
    ),
    "versioning": (
        "How do you version a REST API?",
        "/v1/... URL prefix or Accept-Version header.",
    ),
    "pagination": (
        "What are cursor vs offset pagination trade-offs?",
        "Cursor: stable under inserts, better for large datasets. Offset: simpler but slow on large offsets.",
    ),
    "idempotency": (
        "Define idempotency in APIs.",
        "Operation produces the same result regardless of execution count. PUT/DELETE are idempotent; POST needs Idempotency-Key for retries.",
    ),
    "dag": (
        "What is an Airflow DAG?",
        "Directed acyclic graph of tasks with dependency edges. Defines workflow execution order.",
    ),
    "operators": (
        "What are Airflow operators?",
        "Python task definitions with dependencies. Each operator is one unit of work in the DAG.",
    ),
    "retries": (
        "How should Airflow tasks handle retries?",
        "Configure retries and retry_delay; tasks must be idempotent so reruns are safe.",
    ),
    "sensors": (
        "What is an Airflow sensor?",
        "Task that waits for an external condition before downstream tasks run.",
    ),
    "work queue": (
        "What is the RabbitMQ work-queue pattern?",
        "Competing consumers pull jobs; acknowledge only after successful processing.",
    ),
    "pub/sub": (
        "What is RabbitMQ pub/sub?",
        "Fanout to many subscribers. One message delivered to all bound queues.",
    ),
    "dlq": (
        "What is a dead-letter queue for?",
        "Poison messages after max retries — isolate without blocking the main queue.",
    ),
    "tracker (bep 15)": (
        "What does a BitTorrent tracker do (BEP 15)?",
        "UDP announce protocol — clients exchange peer lists to discover download sources.",
    ),
    "wire protocol": (
        "What does the BitTorrent wire protocol handle?",
        "Peer handshake, choke/unchoke, and block requests between clients.",
    ),
    "concurrency model": (
        "What is BitTorrent's concurrency model?",
        "Per-peer goroutines with a piece bitmap and work scheduling across the swarm.",
    ),
}


def concept_card(term: str, defin: str, section: str = "") -> tuple[str, str]:
    key = term.lower().strip()
    if key in CONCEPT_QA:
        return CONCEPT_QA[key]
    low_term = term.lower()
    if low_term in ("b-tree default for `where`, `join`, `order by` left-prefix",):
        pass
    # Generic patterns
    if "step" in section.lower() or re.search(r"step \d", section, re.I):
        return f"{section}: {term}?", cap_sentence(defin)
    return f"What is {term}?", cap_sentence(defin)


def service_card(svc: str, use: str) -> tuple[str, str]:
    return (
        f"When would you choose AWS {svc}?",
        cap_sentence(use),
    )


def mock_card(prompt: str, hints: list[str]) -> tuple[str, str]:
    parts = [cap_sentence(h).rstrip(".") for h in hints if h.strip()]
    if len(parts) == 1:
        return prompt, cap_sentence(parts[0])
    back = ". ".join(f"{i + 1}) {p}" for i, p in enumerate(parts)) + "."
    return prompt, back


# General fundamentals — definition + trade-off/fix pattern (complements resume-specific cards).
FUNDAMENTALS: list[tuple[list[str], str, str, str]] = [
    (
        ["backend", "go", "concept"],
        "What is a mutex? When would you use it over a channel in Go?",
        "Mutual exclusion lock preventing concurrent access to shared resource. Mutex for protecting state; channels for goroutine communication.",
        "Concurrency",
    ),
    (
        ["backend", "go", "concept"],
        "Define deadlock. How do you prevent it?",
        "Circular wait where goroutines block indefinitely. Prevent via lock ordering, timeouts, and avoiding nested locks.",
        "Concurrency",
    ),
    (
        ["backend", "go", "concept"],
        "What is a race condition?",
        "Multiple goroutines access shared data without synchronization, causing unpredictable results. Fix with mutex, channels, or atomic ops.",
        "Concurrency",
    ),
    (
        ["backend", "db", "concept"],
        "What is a database index? What are the trade-offs?",
        "Data structure (B-tree, hash) enabling faster query lookups. Trade-off: faster reads, slower writes, increased storage.",
        "Databases",
    ),
    (
        ["backend", "db", "concept"],
        "Define normalization vs. denormalization.",
        "Normalization reduces redundancy via separate tables; denormalization duplicates data for faster reads at the cost of consistency complexity.",
        "Databases",
    ),
    (
        ["backend", "db", "concept"],
        "What is the N+1 query problem?",
        "One query fetches parent records, then N additional queries per child. Fix: JOIN or batch loading.",
        "Databases",
    ),
    (
        ["backend", "concept", "resilience"],
        "What is cache invalidation?",
        "Removing or updating stale cache entries when source data changes. Strategies: TTL, event-based, LRU.",
        "Caching",
    ),
    (
        ["backend", "concept", "resilience"],
        "Define cache stampede.",
        "Multiple requests hit expired cache simultaneously, all query the database. Fix: lock-based refresh or probabilistic early expiration.",
        "Caching",
    ),
    (
        ["backend", "api", "concept"],
        "Define idempotency in APIs.",
        "Operation produces the same result regardless of execution count. Critical for safe retries on payments and transfers.",
        "API",
    ),
    (
        ["backend", "resilience", "concept"],
        "What is eventual consistency?",
        "Distributed replicas converge to the same state eventually, not immediately. Trade-off that favors availability over strong consistency.",
        "Architecture",
    ),
    (
        ["backend", "resilience", "concept"],
        "What is a circuit breaker?",
        "Pattern halting requests to a failing service after a threshold. States: closed (normal) → open (failing) → half-open (testing).",
        "Architecture",
    ),
    (
        ["backend", "db", "concept"],
        "Define sharding.",
        "Horizontal partitioning splitting data across multiple database instances by shard key (user ID, region).",
        "Architecture",
    ),
    (
        ["backend", "go", "concept"],
        "What is a goroutine?",
        "Lightweight thread managed by the Go runtime. Thousands run concurrently on few OS threads (M:N scheduling).",
        "Go",
    ),
    (
        ["backend", "go", "concept"],
        "What is an empty interface (any) in Go?",
        "Type that accepts any value. Used for dynamic typing; requires type assertion or switch to use the concrete value.",
        "Go",
    ),
]


def infer_drill_meta(path: Path, lines: list[str]) -> tuple[str, list[str], str]:
    rel = path.relative_to(DRILLS)
    parts = rel.parts
    section = "/".join(parts[:-1])
    tags: list[str] = ["backend", "drill"]

    if "core5" in parts:
        tags.append("core5")
        return "Core 5", tags, section

    if "blocks" in parts and len(parts) >= 2:
        block = parts[1]
        tags.extend(BLOCK_TOPIC_TAGS.get(block, ["block"]))
        title = block.replace("_", " ")
        return title, tags, section

    if "revision" in parts and len(parts) >= 2:
        rev = parts[1]
        tags.extend(REVISION_TOPIC_TAGS.get(rev, ["revision"]))
        return rev.replace("_", " "), tags, section

    return section, tags, section


def parse_explain_drills() -> None:
    for path in sorted(DRILLS.rglob("main.go")):
        if "write" in path.parts or "scenario" in path.parts:
            continue
        text = path.read_text()
        lines = text.splitlines()
        section_title, base_tags, section_path = infer_drill_meta(path, lines)
        hints = parse_assert_hints(text)
        hint_i = 0
        source = str(path.relative_to(ROOT)).replace("\\", "/")

        for line in lines:
            m = EXPLAIN_RE.search(line)
            if not m:
                continue
            parsed = parse_explain_prompt(m.group(1))
            if not parsed:
                continue
            body = m.group(1)
            hint = hints[hint_i] if hint_i < len(hints) else ""
            front, back = drill_card(body, hint)
            hint_i += 1
            add(
                "backend",
                base_tags,
                front,
                back,
                source,
                section_title,
            )


def parse_mock_scenarios() -> None:
    path = DRILLS / "scenario" / "mock_scenarios" / "main.go"
    if not path.exists():
        return
    text = path.read_text()
    source = str(path.relative_to(ROOT)).replace("\\", "/")
    start = text.find("var scenarios")
    end = text.find("\nfunc main()")
    scenario_text = text[start:end] if start >= 0 and end > start else text
    blocks = re.split(r"\n\t\{", scenario_text)
    for block in blocks[1:]:
        title_m = re.search(r'title:\s*"([^"]+)"', block)
        prompt_m = re.search(r'prompt:\s*"([^"]+)"', block)
        hints = re.findall(r'"([^"]+)"', block.split("hints:", 1)[1] if "hints:" in block else "")
        if not title_m or not prompt_m or not hints:
            continue
        title, prompt = title_m.group(1), prompt_m.group(1)
        front, back = mock_card(prompt, hints)
        add(
            "backend",
            ["backend", "scenario", "mock"],
            front,
            back,
            source,
            title,
        )


def parse_trigger_tables(text: str, source: str) -> None:
    """Interview trigger tables (if they ask X → lead with Y)."""
    section = ""
    lines = text.splitlines()
    i = 0
    while i < len(lines):
        line = lines[i]
        if line.startswith("## "):
            section = line[3:].strip()
            i += 1
            continue
        if "|" in line and i + 1 < len(lines) and re.match(r"\|\s*-+", lines[i + 1]):
            headers = [h.strip().lower() for h in line.strip("|").split("|")]
            if "if they ask" not in " ".join(headers):
                i += 1
                continue
            term_idx, def_idx = 0, 1
            for idx, h in enumerate(headers):
                if "if they ask" in h:
                    term_idx = idx
                if "lead with" in h:
                    def_idx = idx
            i += 2
            while i < len(lines) and lines[i].strip().startswith("|"):
                cols = [c.strip() for c in lines[i].strip("|").split("|")]
                if term_idx < len(cols) and def_idx < len(cols):
                    term, defin = cols[term_idx], cols[def_idx]
                    if term and defin and not term.startswith("-"):
                        front, back = trigger_card(term, defin)
                        add(
                            "backend",
                            ["backend", "trigger"],
                            front,
                            back,
                            source,
                            section,
                        )
                i += 1
            continue
        i += 1


def parse_service_tables(text: str, source: str) -> None:
    """AWS-style service | use tables."""
    section = ""
    lines = text.splitlines()
    i = 0
    while i < len(lines):
        line = lines[i]
        if line.startswith("## "):
            section = line[3:].strip()
        if "|" in line and i + 1 < len(lines) and re.match(r"\|\s*-+", lines[i + 1]):
            headers = [h.strip().lower() for h in line.strip("|").split("|")]
            if "service" in headers and "use" in headers:
                svc_idx = headers.index("service")
                use_idx = headers.index("use")
                i += 2
                while i < len(lines) and lines[i].strip().startswith("|"):
                    cols = [c.strip() for c in lines[i].strip("|").split("|")]
                    if svc_idx < len(cols) and use_idx < len(cols):
                        svc, use = cols[svc_idx], cols[use_idx]
                        if svc and use and not svc.startswith("-"):
                            front, back = service_card(svc, use)
                            add(
                                "backend",
                                ["backend", "aws"],
                                front,
                                back,
                                source,
                                section,
                            )
                    i += 1
                continue
        i += 1


def parse_concept_bullets(text: str, source: str) -> None:
    section = ""
    subsection = ""
    for line in text.splitlines():
        if line.startswith("## "):
            section = line[3:].strip()
            subsection = ""
            continue
        if line.startswith("### "):
            subsection = line[4:].strip()
            continue
        if line.startswith("|") or line.startswith("```"):
            continue
        m = re.match(r"^- \*\*([^*]+):\*\*\s*(.+)$", line)
        if m:
            term, defin = m.group(1), m.group(2)
            sec = subsection or section
            tags = ["backend", "concept"]
            low = (section + subsection + term).lower()
            for key, t in [
                ("rest", "api"),
                ("jwt", "jwt"),
                ("database", "db"),
                ("index", "db"),
                ("query", "db"),
                ("pool", "db"),
                ("distributed", "resilience"),
                ("circuit", "resilience"),
                ("retry", "resilience"),
                ("websocket", "webrtc"),
                ("webrtc", "webrtc"),
                ("stun", "webrtc"),
                ("airflow", "airflow"),
                ("rabbit", "messaging"),
                ("jenkins", "devops"),
                ("docker", "devops"),
                ("aws", "aws"),
                ("zatca", "zatca"),
                ("compliance", "zatca"),
                ("go ", "go"),
                ("bittorrent", "go"),
            ]:
                if key in low:
                    tags.append(t)
            front, back = concept_card(term, defin, sec)
            add(
                "backend",
                tags,
                front,
                back,
                source,
                sec,
            )
        m_num = re.match(r"^(\d+)\.\s+(.+)$", line.strip())
        if m_num and subsection:
            step, body = m_num.group(1), m_num.group(2)
            front, back = step_card(subsection, step, body)
            add(
                "backend",
                ["backend", "concept", subsection.lower().replace(" ", "_")],
                front,
                back,
                source,
                subsection,
            )


STEP_QUESTIONS: dict[str, list[str]] = {
    "jwt": [
        "What is the first step in JWT authentication?",
        "How does the client send a JWT on each request?",
        "What does JWT middleware validate before handling the request?",
    ],
    "query optimization workflow": [
        "Query optimization — what is step 1?",
        "After EXPLAIN, how do you rewrite a slow query?",
        "When do you add or adjust indexes?",
        "What do you do after indexing for still-hot queries?",
    ],
    "nat debug checklist": [
        "WebRTC NAT debug — what do you check first?",
        "What ICE candidate types should be gathered?",
        "What connection state confirms success vs failure?",
        "When should you deploy TURN?",
        "What firewall rules matter for WebRTC?",
    ],
    "invoice submission flow": [
        "ZATCA invoice flow — what is generated first?",
        "What validation runs before signing?",
        "How is the invoice signed?",
        "How is the invoice submitted to the authority?",
        "What do you store after submission?",
    ],
}


def step_card(subsection: str, step: str, body: str) -> tuple[str, str]:
    key = subsection.lower().strip()
    idx = int(step) - 1
    if key in STEP_QUESTIONS and idx < len(STEP_QUESTIONS[key]):
        return STEP_QUESTIONS[key][idx], cap_sentence(body)
    return f"{subsection} — what is step {step}?", cap_sentence(body)


def build_synthesis_cards() -> None:
    """High-level Q&A not covered by one-line drill prompts."""
    source = "doc/backend/DRILL_CONCEPTS.md"
    curated = [
        (["backend", "jwt"], "Walk through JWT auth flow end-to-end.", "Client logs in → server issues access (+ optional refresh) token. Client sends Authorization: Bearer <token>. Middleware validates signature, exp, iss, aud; attaches claims to context.", "JWT"),
        (["backend", "resilience"], "What are the circuit breaker states?", "Closed (normal) → failures ≥ threshold → open (fail fast) → after timeout → half-open (probe) → success closes; failure reopens.", "Circuit breaker"),
        (["backend", "resilience"], "What is the exponential backoff formula? When should you retry?", "delay = min(cap, base × 2^attempt) + jitter. Only retry transient errors (timeouts, 503). Require idempotent ops or Idempotency-Key.", "Retry"),
        (["backend", "cap"], "During a network partition, what do most web APIs favor?", "Availability over strong consistency — eventual consistency for reads is the common trade-off.", "CAP"),
        (["backend", "saga"], "Saga vs two-phase commit (2PC)?", "2PC: strong consistency, fragile under partitions. Saga: local transactions + compensating actions; common in microservices.", "Saga"),
        (["backend", "webrtc"], "WebRTC: what travels over signaling vs media?", "Signaling (WebSocket): SDP offer/answer, ICE candidates. Media (UDP): encrypted RTP peer-to-peer, not through your server.", "WebRTC"),
        (["backend", "webrtc"], "STUN vs TURN — when do you use each?", "STUN discovers public IP/port (low cost). TURN relays when P2P fails (higher cost; symmetric NAT, strict firewalls).", "STUN/TURN"),
        (["backend", "api"], "Which REST verbs are idempotent?", "PUT and DELETE are idempotent. POST needs Idempotency-Key header for safe retries.", "REST"),
        (["backend", "db"], "What is the composite index column order rule?", "Equality filter columns first, then range columns. Follows the left-prefix rule.", "Indexing"),
        (["backend", "db"], "What is the query optimization workflow?", "EXPLAIN ANALYZE → rewrite query → add/adjust indexes → cache hot reads and tune connection pool.", "Query opt"),
        (["backend", "db"], "How do you size a database connection pool?", "Rule of thumb: ≈ (cores × 2) + spindle. Measure under load; watch for pool exhaustion causing latency spikes.", "Pooling"),
        (["backend", "nat"], "NAT debug checklist for WebRTC?", "1) Signaling connected? 2) ICE candidates gathered (host, srflx, relay)? 3) Connection state connected vs failed? 4) Try TURN if only host candidates. 5) Firewall allows UDP/TURN TLS.", "NAT"),
        (["backend", "airflow"], "How do you recover when an Airflow task fails?", "Clear the failed task and rerun from that node. Tasks must be idempotent; configure retries and retry_delay.", "Airflow"),
        (["backend", "messaging"], "RabbitMQ work queue vs pub/sub?", "Work queue: competing consumers, ack after processing. Pub/sub: fanout to many subscribers.", "RabbitMQ"),
        (["backend", "messaging"], "What is a dead-letter queue for?", "Poison messages after max retries — isolate without blocking the main queue.", "DLQ"),
        (["backend", "jenkins"], "What are typical Jenkins pipeline stages?", "checkout → install → lint → unit test → build → integration → security scan → deploy staging → smoke → deploy prod", "Jenkins"),
        (["backend", "docker"], "Liveness vs readiness health checks?", "Liveness: restart if dead. Readiness: whether to send traffic.", "Docker"),
        (["backend", "aws"], "When to use EC2 vs Lambda vs CDN?", "EC2: long-running services. Lambda: short event handlers. CDN: static assets at the edge.", "AWS"),
        (["backend", "zatca"], "What is the ZATCA invoice submission flow?", "Generate XML → XSD validate → cryptographic sign → submit over TLS → store response + audit (who/when/payload hash).", "ZATCA"),
        (["backend", "go"], "What Go concurrency primitives do you use in servers?", "Goroutines and channels for worker pools; context.Context for cancel/deadlines; sync.Mutex for shared state.", "Go"),
        (["backend", "go"], "How do you implement HTTP graceful shutdown in Go?", "Call Shutdown(ctx) to stop accepting new connections, drain in-flight requests, then exit.", "Go HTTP"),
        (["backend", "go"], "How does BitTorrent piece scheduling work?", "Per-peer goroutines, piece bitmap, request rarest pieces first to improve swarm health.", "BitTorrent"),
        (["backend", "db"], "What is a covering index? When does it help?", "Index includes all columns in SELECT → index-only scan with no table lookup.", "Indexing"),
        (["backend", "db"], "What is the page-builder schema pattern for versions?", "Versioned entities: pages, page_versions, audit_log; separate draft vs published snapshots.", "Schema"),
        (["backend", "api"], "What shape should a structured API error take?", "code (e.g. PAGE_NOT_FOUND), message, optional details object.", "REST errors"),
        (["backend", "api"], "Cursor vs offset pagination — trade-offs?", "Cursor: stable under inserts. Offset: simple but degrades on large offsets.", "Pagination"),
        (["backend", "webrtc"], "What are WebSocket use cases in real-time apps?", "Signaling, chat, live notifications. Use heartbeats/ping to detect dead connections.", "WebSocket"),
        (["backend", "zatca"], "How do you protect signed invoices from replay attacks?", "Unique invoice ID, timestamp, and nonce in the signed payload.", "Replay"),
        (["backend", "go"], "What order should HTTP middleware run in?", "Logging → recovery → auth → handler. Context carries request ID and user.", "Go HTTP"),
        (["backend", "go"], "What are BitTorrent wire protocol basics?", "Handshake, choke/unchoke, request blocks. Tracker (BEP 15) UDP announce for peer discovery.", "BitTorrent"),
    ]
    for tags, front, back, sec in curated:
        add("backend", tags, front, back, source, sec)


def build_fundamentals_cards() -> None:
    source = "cards/generate_cards.py"
    for tags, front, back, sec in FUNDAMENTALS:
        add("backend", tags, front, back, source, sec)


def build_backend_cards() -> None:
    concepts_path = DOC / "backend" / "DRILL_CONCEPTS.md"
    concepts = concepts_path.read_text()

    parse_trigger_tables(concepts, "doc/backend/DRILL_CONCEPTS.md")
    # WEEKLY_REVISION trigger table overlaps DRILL_CONCEPTS — skip duplicate triggers
    parse_service_tables(concepts, "doc/backend/DRILL_CONCEPTS.md")
    parse_concept_bullets(concepts, "doc/backend/DRILL_CONCEPTS.md")
    build_fundamentals_cards()
    build_synthesis_cards()
    parse_explain_drills()
    parse_mock_scenarios()


def build_star_cards() -> None:
    star = (DOC / "backend" / "STAR_STORIES.md").read_text()
    for s in re.split(r"\n## \d+\. ", star)[1:]:
        title = s.split("\n", 1)[0].strip()
        sit = re.search(r"\*\*Situation:\*\* (.+)", s)
        task = re.search(r"\*\*Task:\*\* (.+)", s)
        action = re.search(r"\*\*Action:\*\* (.+)", s)
        result = re.search(r"\*\*Result:\*\* (.+)", s)
        kw = re.search(r"\*\*Keywords:\*\* (.+)", s)
        if sit and result:
            back = f"S: {sit.group(1)} T: {task.group(1) if task else ''} A: {action.group(1) if action else ''} R: {result.group(1)}"
            add("star", ["backend", "star"], f"STAR: {title} — rehearse S→T→A→R", back, "doc/backend/STAR_STORIES.md", title)
            if kw:
                add("star", ["backend", "star"], f"STAR keywords: {title}?", kw.group(1), "doc/backend/STAR_STORIES.md", title)


def build() -> None:
    cards.clear()
    build_backend_cards()
    build_star_cards()

    seen: set[str] = set()
    uniq: list[dict] = []
    for c in cards:
        if c["id"] in seen:
            continue
        seen.add(c["id"])
        uniq.append(c)

    OUT.mkdir(parents=True, exist_ok=True)

    keep = {"backend", "star"}
    for old in OUT.glob("*.json"):
        if old.stem not in keep:
            old.unlink()

    by: dict[str, list] = {}
    for c in uniq:
        by.setdefault(c["deck"], []).append(c)

    manifest = []
    total = 0
    for deck, items in sorted(by.items()):
        (OUT / f"{deck}.json").write_text(json.dumps(items, indent=2, ensure_ascii=False) + "\n")
        manifest.append({"deck": deck, "count": len(items), "file": f"decks/{deck}.json"})
        total += len(items)
        print(f"{deck}: {len(items)}")

    (ROOT / "cards" / "manifest.json").write_text(
        json.dumps({"total": total, "decks": manifest}, indent=2) + "\n"
    )
    print("TOTAL", total)


if __name__ == "__main__":
    build()
