#!/usr/bin/env python3
"""Rebuild backend interview flashcards from doc/backend/ notes.

Writes only question-style cards to cards/decks/backend.json and star.json.
Run from repo root: python3 cards/generate_cards.py
"""
from __future__ import annotations

import hashlib
import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
DOC = ROOT / "doc"
OUT = ROOT / "cards" / "decks"
cards: list[dict] = []


def cid(*parts: str) -> str:
    return hashlib.sha1("|".join(parts).encode()).hexdigest()[:12]


def clean(s: str) -> str:
    s = re.sub(r"\*\*([^*]+)\*\*", r"\1", s)
    s = re.sub(r"`([^`]+)`", r"\1", s)
    return re.sub(r"\s+", " ", s).strip()


def add(deck: str, tags: list[str], front: str, back: str, source: str, section: str = "") -> None:
    front, back = clean(front), clean(back)
    if not front or not back or front == back or len(front) < 8 or len(back) < 3:
        return
    if re.fullmatch(r"What is \d+\?", front):
        return
    cards.append(
        {
            "id": cid(deck, front),
            "deck": deck,
            "tags": tags,
            "front": front,
            "back": back,
            "source": source,
            "section": section,
        }
    )


def parse_backend_triggers(text: str, source: str) -> None:
    """Interview trigger table only — not full note bullets."""
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
                            defin,
                            source,
                            section,
                        )
                i += 1
            continue
        i += 1


def build_backend_cards() -> None:
    backend = (DOC / "backend" / "DRILL_CONCEPTS.md").read_text()
    source = "doc/backend/DRILL_CONCEPTS.md"
    parse_backend_triggers(backend, source)

    curated = [
        (["backend", "jwt"], "Walk through JWT auth flow.", "Client logs in → server issues access (+ optional refresh). Client sends Authorization: Bearer <token>. Middleware validates signature, exp, iss, aud; attaches claims to context.", "JWT"),
        (["backend", "resilience"], "Circuit breaker state machine?", "CLOSED → failures ≥ threshold → OPEN (fail fast) → after timeout → HALF-OPEN (probe) → success → CLOSED; failure → OPEN", "Circuit breaker"),
        (["backend", "resilience"], "Exponential backoff formula + rules?", "delay = min(cap, base * 2^attempt) + jitter. Only retry transient errors. Require idempotent ops or Idempotency-Key.", "Retry"),
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
    ]
    for tags, front, back, sec in curated:
        add("backend", tags, front, back, source, sec)

    # Weekly revision drill prompts (doc/backend/WEEKLY_REVISION.md)
    revision = [
        (["backend", "revision"], "Revision Mon: API auth recap — idempotent update verb?", "PUT (safe to retry without side effects).", "01_api_auth_recap"),
        (["backend", "revision"], "Revision Tue: slow SQL — first diagnostic step?", "EXPLAIN (or EXPLAIN ANALYZE) the query plan.", "02_data_resilience_recap"),
        (["backend", "revision"], "Revision Wed: circuit breaker prevents what?", "Cascading failures across downstream services.", "03_distributed_realtime"),
        (["backend", "revision"], "Revision Thu: when P2P WebRTC fails, use what?", "TURN relay server.", "04_realtime_messaging"),
        (["backend", "revision"], "Revision Fri: readiness probe gates what?", "Whether the pod receives traffic.", "05_devops_orchestration"),
        (["backend", "revision"], "Revision Sat: all 9 triggers — invoice rejected?", "XSD validate → sign → submit → correlate audit + payload hash.", "06_full_week_sweep"),
        (["backend", "revision"], "Revision Sun: ZATCA step before API submit?", "Cryptographic signing after XSD validation.", "07_go_compliance_mix"),
    ]
    for tags, front, back, sec in revision:
        add("backend", tags, front, back, "doc/backend/WEEKLY_REVISION.md", sec)


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

    # Remove legacy DSA and Back2Basics full-note decks.
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
