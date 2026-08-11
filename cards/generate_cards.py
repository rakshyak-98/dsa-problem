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
    return re.findall(r'assert\([^,]+,\s*[^,]+,\s*"([^"]+)"\)', text)


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
            front, back = parsed
            if hint_i < len(hints) and len(back) < 20:
                back = f"{back} — {hints[hint_i]}"
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
    blocks = re.split(r"\n\t\{", text)
    for block in blocks[1:]:
        title_m = re.search(r'title:\s*"([^"]+)"', block)
        prompt_m = re.search(r'prompt:\s*"([^"]+)"', block)
        hints = re.findall(r'"([^"]+)"', block.split("hints:", 1)[1] if "hints:" in block else "")
        if not title_m or not prompt_m or not hints:
            continue
        title, prompt = title_m.group(1), prompt_m.group(1)
        back = "; ".join(hints)
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
                            defin,
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
                defin,
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
