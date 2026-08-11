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


def cap_sentence(s: str) -> str:
    s = clean(s)
    if not s:
        return s
    if s[0].islower():
        s = s[0].upper() + s[1:]
    if s[-1] not in ".!?":
        s += "."
    return s


def either_phrase(options: str) -> str:
    parts = [p.strip() for p in re.split(r"\s+or\s+", options, flags=re.I) if p.strip()]
    if len(parts) == 2:
        return f"Either {parts[0]} or {parts[1]}"
    if len(parts) > 2:
        return ", or ".join(parts[:-1]) + f", or {parts[-1]}"
    return options


# Assert hints from explain drills → spoken interview phrasing.
HINT_INTERVIEW: dict[str, str] = {
    "200 ok": "For a successful GET, I'd return 200 OK.",
    "201 created": "When a resource is created successfully, return 201 Created.",
    "authorization header": "Send the JWT in the Authorization header as Bearer <token>.",
    "b-tree indexes": "InnoDB uses B-tree indexes by default — good for equality and range scans.",
    "cdn at edge": "Use a CDN like CloudFront to cache static assets at the edge, close to users.",
    "ci before cd": "Run CI — tests, lint, and security scans — before any deploy to production.",
    "cp or ap": "During a partition you choose CP or AP — consistency or availability.",
    "dag definition": "A DAG is a directed acyclic graph of tasks with explicit dependencies.",
    "docker image": "A Docker image is a layered filesystem snapshot used to run containers.",
    "explain analyze first": "I'd start with EXPLAIN ANALYZE to see the execution plan and where time is spent.",
    "explain plan": "Lead with the EXPLAIN plan — look for full table scans, missing indexes, or bad joins.",
    "http upgrade": "WebSocket begins as an HTTP Upgrade request, then switches to full-duplex TCP.",
    "ice candidates": "ICE gathers host, server-reflexive (STUN), and relay (TURN) candidates.",
    "idempotency-key": "Send an Idempotency-Key header so duplicate POST retries are safe.",
    "idempotency-key header": "Put the idempotency key in a request header — typically Idempotency-Key.",
    "lambda for events": "AWS Lambda fits short, event-driven handlers — not long-running services.",
    "m goroutines on n threads": "Go multiplexes many goroutines onto fewer OS threads (M:N scheduling).",
    "nat blocks direct path": "NAT often blocks direct peer paths — that's when you need STUN/TURN.",
    "put is idempotent": "PUT is idempotent — safe to retry the same update without duplicate side effects.",
    "rest resources": "Design around REST resources — nouns, plural URLs, proper verbs and status codes.",
    "rs256 uses public key": "For RS256, verify the JWT signature with the issuer's public key.",
    "sdp offer/answer": "Signaling exchanges SDP offer/answer to negotiate codecs and media parameters.",
    "sha-256 common": "Sign with a strong hash like SHA-256 before applying the private key.",
    "tcp reuse": "HTTP keep-alive reuses the same TCP connection for multiple requests on one socket.",
    "tcp/auth handshake": "Pooling avoids repeating the TCP and database authentication handshake on every query.",
    "turn relays media": "TURN relays media when direct peer-to-peer fails — common behind VPNs or strict firewalls.",
    "url versioning": "Version the API in the URL path — e.g. /v1/pages — or via an Accept-Version header.",
    "where / join": "Indexes accelerate WHERE filters and JOIN keys — mind the left-prefix rule on composites.",
    "where/join": "Indexes accelerate WHERE filters and JOIN keys — mind the left-prefix rule on composites.",
    "websocket": "WebSocket gives full-duplex TCP — ideal for signaling, chat, and live notifications.",
    "websocket signaling": "Use WebSocket for signaling — SDP exchange and ICE trickling.",
    "xml schema": "Validate invoice XML against the authority XSD schema before signing.",
    "xsd validation": "Run XSD validation against the authority schema before signing and submission.",
    "actor + timestamp": "Audit trails record who changed what and when — actor, timestamp, and payload hash.",
    "batch/join": "Fix N+1 queries with JOINs, eager loading, or batched fetches instead of per-row lookups.",
    "bencode format": "BitTorrent metadata uses bencode — a compact binary encoding for .torrent structures.",
    "bencode metadata": "BitTorrent .torrent files encode metadata in bencode format.",
    "cancel/deadline": "Use context.Context to propagate cancellation and deadlines across goroutines.",
    "cascading failures": "Circuit breakers stop cascading failures by failing fast when a dependency is unhealthy.",
    "closed, open, half-open": "Circuit breaker states: CLOSED (normal), OPEN (fail fast), HALF-OPEN (probing).",
    "competing consumers": "Work-queue pattern — multiple competing consumers, ack after successful processing.",
    "connection reuse": "Connection pooling reuses open DB connections instead of opening one per request.",
    "context.context": "context.Context carries cancellation, deadlines, and request-scoped values.",
    "cryptographic signing": "Cryptographically sign the payload — hash with SHA-256, then sign with the private key.",
    "dead letter queue": "Route poison messages to a dead-letter queue after max retries so the main queue isn't blocked.",
    "directed acyclic graph": "Airflow models workflows as a DAG — directed acyclic graph of dependent tasks.",
    "distributed tx pattern": "For distributed transactions in microservices, I'd use a saga instead of 2PC.",
    "dots separate parts": "JWT has three Base64URL parts separated by dots: header, payload, signature.",
    "double each attempt": "Exponential backoff doubles the delay each attempt (base 2) and adds jitter.",
    "duplex push": "WebSocket enables server-push — full-duplex, unlike plain HTTP request/response.",
    "execution plan": "EXPLAIN shows the query execution plan — which indexes are used and estimated row counts.",
    "exp claim": "The exp claim enforces token expiry — reject any JWT past that timestamp.",
    "fail fast": "When the circuit is open, fail fast — return an error immediately without calling the sick service.",
    "goroutines": "Use goroutines for concurrency — lightweight threads scheduled by the Go runtime.",
    "header.payload.signature": "A JWT has three dot-separated parts: header, payload, and signature.",
    "http.server.shutdown": "Call http.Server.Shutdown(ctx) — stop accepting new connections and drain in-flight requests.",
    "idempotent ops": "Only retry idempotent operations, or attach an Idempotency-Key for write retries.",
    "idempotent tasks": "Airflow tasks must be idempotent so clearing and rerunning a failed node is safe.",
    "integrity/authenticity": "A digital signature proves integrity (unchanged payload) and authenticity (trusted signer).",
    "k8s/docker health": "Use liveness (restart if dead) and readiness (gate traffic) probes in Kubernetes or Docker.",
    "layered filesystem": "Docker images are stacked layers — multi-stage builds keep production images small.",
    "leftmost prefix rule": "Yes — composite indexes follow the leftmost-prefix rule, so column order matters.",
    "liveness restarts pod": "Liveness probe restarts the container if the process is hung or dead.",
    "open circuit fails fast": "An open circuit fails fast — callers get an immediate error without hitting the dependency.",
    "operators/tasks": "In Airflow, operators define tasks — Python units of work wired together in a DAG.",
    "path or header version": "Either URL path versioning (/v1/...) or an Accept-Version header works.",
    "peer-to-peer media": "WebRTC media travels peer-to-peer over encrypted UDP — not through your API server.",
    "pick cp or ap": "During a partition, pick CP or AP — you sacrifice either consistency or availability.",
    "poison messages": "Poison messages go to a dead-letter queue after repeated failures.",
    "public ip": "STUN discovers the client's public IP and port for NAT hole punching.",
    "readiness gates traffic": "Readiness probe gates traffic — the load balancer skips instances that aren't ready.",
    "retry/backoff": "On outages I'd use timeouts, retries with exponential backoff, then circuit breakers.",
    "safe retries": "Tasks need to be idempotent so retries and reruns from a failed node are safe.",
    "saga over 2pc": "I'd choose a saga over 2PC — local transactions plus compensating actions, more resilient to partitions.",
    "short-lived/event-driven": "Lambda suits short-lived, event-driven workloads — not always-on services.",
    "static content edge cache": "CDN caches static content at the edge to cut latency and origin load.",
    "stop accepting, drain": "Graceful shutdown stops accepting new connections and drains in-flight requests before exit.",
    "tests in pipeline": "The CI pipeline runs tests (and lint/security scans) before any deploy artifact ships.",
    "thundering herd": "Jitter randomizes retry timing so clients don't all retry at once — avoiding a thundering herd.",
    "unique request id": "Prevent replays with a unique invoice ID, timestamp, and nonce in the signed payload.",
    "work queue pattern": "Work-queue pattern — competing consumers pull jobs and ack only after processing.",
}


def hint_interview_line(hint: str) -> str | None:
    if not hint:
        return None
    key = hint.lower().strip().rstrip(".")
    return HINT_INTERVIEW.get(key)


def expand_terse_answer(front: str, raw: str, hint: str) -> str:
    """Turn drill keywords into spoken interview phrasing."""
    fl = front.lower().rstrip("?").strip()
    base = hint or raw
    al = base.lower()
    rl = raw.lower()

    if re.fullmatch(r"\d{3}", base):
        label = hint if hint and not re.fullmatch(r"\d{3}", hint) else ""
        if label:
            return f"I'd return HTTP {base} ({label})."
        return f"I'd return HTTP {base}."

    if "http status" in fl and re.search(r"\d{3}", base):
        return cap_sentence(base)

    if "idempotent" in fl and ("put" in al or "put" in rl):
        return "PUT is idempotent, so it's safe to retry updates without duplicate side effects."

    if ("retry" in fl or "safe" in fl) and ("update" in fl or "side effects" in fl):
        return "PUT is idempotent, so it's safe to retry updates without duplicate side effects."

    if "retry" in fl and "post" in fl:
        return "Use an Idempotency-Key header so POST retries are safe."

    if "jwt" in fl and ("sent" in fl or "header" in fl or "location" in fl):
        return "Send it in the Authorization header as Bearer <token>."

    if "jwt" in fl and ("segment" in fl or "parts" in fl or raw == "3" or base == "3"):
        return "Three dot-separated Base64URL parts: header, payload, and signature."

    if "jwt" in fl and ("rs256" in fl or "signature" in fl or "key" in fl):
        return "Verify RS256 signatures with the issuer's public key."

    if "version" in fl and ("api" in fl or "url" in fl):
        return "Either URL path versioning (/v1/...) or an Accept-Version header."

    if "slow sql" in fl or ("first step" in fl and "sql" in fl):
        return "Start with EXPLAIN ANALYZE to inspect the execution plan."

    if "index" in fl and "clause" in fl:
        return "Indexes help WHERE, JOIN, and ORDER BY — composite indexes follow the left-prefix rule."

    if "pool" in fl and ("reuses" in fl or "reduces" in fl or "avoids" in fl):
        return "Connection pooling reuses TCP connections and avoids repeated handshake overhead."

    if "circuit breaker" in fl and "open" in fl:
        return "When open, the circuit breaker fails fast instead of hammering a sick dependency."

    if "backoff" in fl or "multiplier" in fl:
        return "Use exponential backoff with base multiplier 2 — double each attempt and add jitter."

    if "partition" in fl and ("consistency" in rl or "availability" in rl):
        return (
            "During a partition you choose CP or AP — sacrifice consistency or availability. "
            "Most web APIs favor availability with eventual consistency."
        )

    if "saga" in fl:
        return "Sagas coordinate distributed transactions via events — choreography or orchestration."

    if "jitter" in fl:
        return "Jitter spreads retry delays so clients don't retry in sync (thundering herd)."

    if "retries need" in fl or ("retry" in fl and "operation" in fl):
        return "Only retry idempotent operations, or use an Idempotency-Key for writes."

    if "idempotency key" in fl:
        return "Send the idempotency key in a request header (e.g. Idempotency-Key)."

    if "websocket" in fl and "upgrade" in fl:
        return "WebSocket starts as an HTTP Upgrade handshake, then runs full-duplex over TCP."

    if "websocket" in fl or ("chat" in fl and "transport" in fl):
        return "WebSocket — full-duplex TCP for signaling, chat, and live push notifications."

    if "webrtc" in fl and "media" in fl:
        return "Media flows peer-to-peer over encrypted UDP — not through your signaling server."

    if "signaling" in fl:
        return "Signaling carries SDP offer/answer and ICE candidates over WebSocket."

    if "ice" in fl:
        return "ICE gathers host, server-reflexive, and relay candidates for NAT traversal."

    if "stun" in fl:
        return "STUN discovers the client's public IP and port for NAT hole punching."

    if "turn" in fl or ("p2p" in fl and "relay" in fl):
        return "TURN relays media when direct peer-to-peer fails (symmetric NAT, strict firewalls)."

    if "dag" in fl:
        return "DAG — a directed acyclic graph of tasks with dependency edges."

    if "airflow" in fl and "unit" in fl:
        return "A task is the unit of work; tasks are wired together in a DAG."

    if "dead letter" in fl or "dlq" in al or "poison" in rl:
        return "Dead-letter queues isolate poison messages after max retries so the main queue keeps moving."

    if "work queue" in fl or ("competing" in fl and "consumer" in fl):
        return "Work queue — competing consumers pull jobs; ack only after successful processing."

    if "liveness" in fl or ("restart" in fl and "health" in fl):
        return "Liveness probe — restart the container if the process is dead."

    if "readiness" in fl or ("traffic" in fl and "health" in fl):
        return "Readiness probe — remove the instance from load balancing until it's ready for traffic."

    if "lambda" in fl:
        return "Lambda fits short, event-driven handlers — not long-running services."

    if "cdn" in fl or "cloudfront" in al:
        return "Use a CDN (e.g. CloudFront) to cache static assets at the edge."

    if "docker" in fl and "layer" in fl:
        return "Images are built from stacked layers; multi-stage builds keep production images small."

    if "xsd" in fl:
        return "XSD validates invoice XML structure against the authority schema before signing."

    if "sign" in fl and ("zatca" in fl or "submit" in fl):
        return "Cryptographically sign the payload (hash + private key) before submission."

    if "context" in fl and "go" in fl:
        return "context.Context carries cancellation, deadlines, and request-scoped values."

    if "graceful" in fl and "shutdown" in fl:
        return "Call http.Server.Shutdown(ctx) — stop accepting new connections and drain in-flight requests."

    if "bencode" in fl:
        return "BitTorrent metadata is bencode — a compact binary encoding for .torrent files."

    if "keep-alive" in fl:
        return "HTTP keep-alive reuses the same TCP connection for multiple requests."

    if "goroutine" in fl and "scheduling" in fl:
        return "Go uses an M:N scheduler — many goroutines multiplexed onto fewer OS threads."

    if "exp" in al and "claim" in fl:
        return "The exp claim enforces JWT expiry — reject tokens past that timestamp."

    if "three states" in fl or "closed,open,half-open" in al.replace(" ", ""):
        return "CLOSED (normal) → OPEN (fail fast) → HALF-OPEN (probe) → back to CLOSED on success."

    if "n+1" in fl:
        return "Fix N+1 with eager loading, JOINs, or batched queries instead of per-row lookups."

    if "explain" in al and len(base) <= 10:
        return "EXPLAIN shows the query execution plan — use EXPLAIN ANALYZE to measure cost."

    if "b-tree" in al or "btree" in al:
        return "InnoDB's default index structure is a B-tree — great for range scans and equality filters."

    if "left-prefix" in fl or ("composite" in fl and "order" in fl):
        return "Yes — composite indexes follow the left-prefix rule; column order matters."

    if re.search(r"\s+or\s+", base, re.I) and len(base) < 40:
        return cap_sentence(either_phrase(base) + " — both are acceptable")

    if len(base) <= 12 and base.islower() and " or " not in base:
        return cap_sentence(base)

    return base


def interview_answer(front: str, raw: str, hint: str = "") -> str:
    raw = clean(raw)
    hint = clean(hint)
    scripted = hint_interview_line(hint)
    if scripted:
        return cap_sentence(scripted)
    answer = expand_terse_answer(front, raw, hint or raw)
    return cap_sentence(answer)


def format_trigger_answer(defin: str) -> str:
    defin = clean(defin)
    if defin.lower().startswith("i'd"):
        return cap_sentence(defin)
    return cap_sentence(f"I'd lead with: {defin}")


def format_mock_answer(hints: list[str]) -> str:
    parts = [cap_sentence(h).rstrip(".") for h in hints if h.strip()]
    if not parts:
        return ""
    if len(parts) == 1:
        return cap_sentence(parts[0])
    return "I'd structure my answer: " + "; ".join(f"{i + 1}) {p}" for i, p in enumerate(parts)) + "."


def format_concept_answer(term: str, defin: str) -> str:
    defin = clean(defin)
    if len(defin) >= 72 or defin.count(".") >= 2:
        return defin
    low = defin.lower()
    if low.startswith(term.lower()):
        return cap_sentence(defin)
    if defin.startswith("{") or defin.startswith("/"):
        return cap_sentence(f"{term}: {defin}")
    return cap_sentence(f"{term} — {defin}")


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
            front, raw = parsed
            hint = hints[hint_i] if hint_i < len(hints) else ""
            back = interview_answer(front, raw, hint)
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
        back = format_mock_answer(hints)
        add(
            "backend",
            ["backend", "scenario", "mock"],
            f"Mock interview: {prompt}",
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
                        add(
                            "backend",
                            ["backend", "trigger"],
                            f"Interview trigger: {term}?",
                            format_trigger_answer(defin),
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
                            add(
                                "backend",
                                ["backend", "aws"],
                                f"When to use AWS {svc}?",
                                use,
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
            add(
                "backend",
                tags,
                f"{sec}: {term}?",
                format_concept_answer(term, defin),
                source,
                sec,
            )
        m_num = re.match(r"^(\d+)\.\s+(.+)$", line.strip())
        if m_num and subsection:
            step, body = m_num.group(1), m_num.group(2)
            add(
                "backend",
                ["backend", "concept", subsection.lower().replace(" ", "_")],
                f"{subsection} — step {step}?",
                body,
                source,
                subsection,
            )


def build_synthesis_cards() -> None:
    """High-level interview answers not covered by one-line drill prompts."""
    source = "doc/backend/DRILL_CONCEPTS.md"
    curated = [
        (["backend", "jwt"], "Walk through JWT auth flow end-to-end.", "Client logs in → server issues access (+ optional refresh). Client sends Authorization: Bearer <token>. Middleware validates signature, exp, iss, aud; attaches claims to context.", "JWT"),
        (["backend", "resilience"], "Circuit breaker state machine?", "CLOSED → failures ≥ threshold → OPEN (fail fast) → after timeout → HALF-OPEN (probe) → success → CLOSED; failure → OPEN", "Circuit breaker"),
        (["backend", "resilience"], "Exponential backoff formula and retry rules?", "delay = min(cap, base * 2^attempt) + jitter. Only retry transient errors. Require idempotent ops or Idempotency-Key.", "Retry"),
        (["backend", "cap"], "CAP during partition — what do most web APIs favor?", "Choose CP or AP. Most web APIs favor availability + eventual consistency for reads.", "CAP"),
        (["backend", "saga"], "Saga vs 2PC?", "2PC: strong consistency, fragile under partitions. Saga: local txs + compensating actions; common in microservices.", "Saga"),
        (["backend", "webrtc"], "WebRTC: signaling vs media paths?", "Signaling (WebSocket): SDP offer/answer, ICE. Media (UDP): encrypted RTP peer-to-peer, not through your server.", "WebRTC"),
        (["backend", "webrtc"], "STUN vs TURN?", "STUN discovers public IP/port (cheap). TURN relays when P2P fails (costly; symmetric NAT / strict firewalls).", "STUN/TURN"),
        (["backend", "api"], "Idempotency rules for REST verbs?", "PUT/DELETE idempotent; POST needs Idempotency-Key for retries.", "REST"),
        (["backend", "db"], "Composite index column order rule?", "Equality filters first, then range.", "Indexing"),
        (["backend", "db"], "Query optimization workflow (4 steps)?", "EXPLAIN ANALYZE → rewrite query → add/adjust indexes → cache hot reads / tune pool.", "Query opt"),
        (["backend", "db"], "Connection pool sizing rule of thumb?", "≈ (cores × 2) + spindle; measure under load; watch pool exhaustion.", "Pooling"),
        (["backend", "nat"], "NAT debug checklist for WebRTC (5 steps)?", "1) Signaling open? 2) ICE candidates gathered? 3) Connection connected vs failed? 4) Try TURN if only host. 5) Firewall UDP/TURN TLS.", "NAT"),
        (["backend", "airflow"], "Airflow recovery when a task fails?", "Clear failed task, rerun from failed node. Tasks must be idempotent; use retries + retry_delay.", "Airflow"),
        (["backend", "messaging"], "RabbitMQ work queue vs pub/sub?", "Work queue: competing consumers, ack after processing. Pub/sub: fanout to many subscribers.", "RabbitMQ"),
        (["backend", "messaging"], "What is a dead letter queue for?", "Poison messages after max retries — isolate without blocking the main queue.", "DLQ"),
        (["backend", "jenkins"], "Typical Jenkins pipeline stages?", "checkout → install → lint → unit test → build → integration → security scan → deploy staging → smoke → deploy prod", "Jenkins"),
        (["backend", "docker"], "Liveness vs readiness health checks?", "Liveness: restart if dead. Readiness: whether to send traffic.", "Docker"),
        (["backend", "aws"], "When to use EC2 vs Lambda vs CDN?", "EC2: long-running services. Lambda: short event handlers. CDN: static assets at the edge.", "AWS"),
        (["backend", "zatca"], "ZATCA invoice submission flow?", "Generate XML → XSD validate → cryptographic sign → submit over TLS → store response + audit (who/when/payload hash).", "ZATCA"),
        (["backend", "go"], "Go concurrency primitives for servers?", "goroutine + channel for worker pools; context.Context for cancel/deadlines; sync.Mutex for shared state (prefer channels for ownership).", "Go"),
        (["backend", "go"], "HTTP graceful shutdown in Go?", "Stop accepting new connections (Shutdown(ctx)), drain in-flight requests, then exit.", "Go HTTP"),
        (["backend", "go"], "BitTorrent piece scheduling at a high level?", "Per-peer goroutines, piece bitmap, request rarest pieces first to improve swarm health.", "BitTorrent"),
        (["backend", "db"], "Covering index — when does it help?", "Index includes all columns in SELECT → index-only scan, no table lookup.", "Indexing"),
        (["backend", "db"], "Page builder schema pattern for versions?", "Versioned entities: pages, page_versions, audit_log; draft vs published snapshots.", "Schema"),
        (["backend", "api"], "Structured API error response shape?", "code (e.g. PAGE_NOT_FOUND), message, optional details object.", "REST errors"),
        (["backend", "api"], "Cursor vs offset pagination tradeoff?", "Cursor: stable under inserts; offset: simple but slow on large offsets.", "Pagination"),
        (["backend", "webrtc"], "WebSocket use cases in real-time apps?", "Signaling, chat, live notifications; heartbeats/ping detect dead connections.", "WebSocket"),
        (["backend", "airflow"], "What is an Airflow sensor?", "Waits for an external condition before downstream tasks run.", "Airflow"),
        (["backend", "zatca"], "Replay protection for signed invoices?", "Unique invoice ID, timestamp, nonce in signed payload.", "Replay"),
        (["backend", "go"], "HTTP server middleware chain order?", "Logging → recovery → auth → handler; context carries request ID + user.", "Go HTTP"),
        (["backend", "go"], "BitTorrent wire protocol basics?", "Handshake, choke/unchoke, request blocks; tracker (BEP 15) UDP announce for peers.", "BitTorrent"),
    ]
    for tags, front, back, sec in curated:
        add("backend", tags, front, back, source, sec)


def build_backend_cards() -> None:
    concepts_path = DOC / "backend" / "DRILL_CONCEPTS.md"
    concepts = concepts_path.read_text()

    parse_trigger_tables(concepts, "doc/backend/DRILL_CONCEPTS.md")
    # WEEKLY_REVISION trigger table overlaps DRILL_CONCEPTS — skip duplicate triggers
    parse_service_tables(concepts, "doc/backend/DRILL_CONCEPTS.md")
    parse_concept_bullets(concepts, "doc/backend/DRILL_CONCEPTS.md")
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
