#!/usr/bin/env python3
"""Rebuild cards/decks/*.json from doc/ notes. Run from repo root or cards/."""
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


def parse_glossary_tables(text: str, source: str, deck: str, base_tags: list[str]) -> None:
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
            hjoin = " ".join(headers)
            interesting = any(
                k in hjoin
                for k in (
                    "term",
                    "plain english",
                    "meaning",
                    "intuition",
                    "lead with",
                    "use",
                    "you read",
                    "shape you see",
                    "pattern name",
                    "trap",
                    "fix",
                    "service",
                    "notation",
                    "object",
                    "formula",
                    "sum",
                    "recurrence",
                    "code shape",
                    "time",
                    "op",
                    "relation",
                    "variant",
                    "problem",
                    "result",
                    "when you see",
                    "math tool",
                )
            )
            if not interesting:
                i += 1
                continue
            term_idx, def_idx = 0, 1
            for idx, h in enumerate(headers):
                if any(
                    k in h
                    for k in (
                        "term",
                        "notation",
                        "op",
                        "object",
                        "you read",
                        "if they ask",
                        "service",
                        "shape you see",
                        "trap",
                        "sum",
                        "recurrence",
                        "code shape",
                        "when you see",
                        "problem",
                        "variant",
                    )
                ):
                    term_idx = idx
                if any(
                    k in h
                    for k in (
                        "plain english",
                        "meaning",
                        "intuition",
                        "lead with",
                        "use",
                        "formula",
                        "notes",
                        "solution",
                        "example",
                        "time",
                        "typical complexity",
                        "it usually means",
                        "closed form",
                        "pattern name",
                        "fix",
                        "math tool",
                        "result",
                    )
                ):
                    def_idx = idx
            i += 2
            while i < len(lines) and lines[i].strip().startswith("|"):
                cols = [c.strip() for c in lines[i].strip("|").split("|")]
                if term_idx < len(cols) and def_idx < len(cols):
                    term, defin = cols[term_idx], cols[def_idx]
                    if term and defin and not term.startswith("-"):
                        if "if they ask" in hjoin:
                            add(deck, base_tags + ["trigger"], f"Interview trigger: {term}?", defin, source, section)
                        elif "you read" in hjoin:
                            add(deck, base_tags, f"When you read “{term}”, what does it usually mean?", defin, source, section)
                        elif "shape you see" in hjoin:
                            add(deck, base_tags, f"Code shape “{term}” → which pattern?", defin, source, section)
                        elif "trap" in hjoin:
                            add(deck, base_tags, f"Reading trap: {term} — how do you fix it?", defin, source, section)
                        elif "service" in hjoin:
                            add(deck, base_tags, f"What is AWS {term} typically used for?", defin, source, section)
                        elif "when you see" in hjoin or "math tool" in hjoin:
                            add(deck, base_tags + ["picker"], f"Math picker: when you see “{term}” → which tool?", defin, source, section)
                        elif "recurrence" in hjoin and "solution" in hjoin:
                            extra = cols[2] if len(cols) > 2 else ""
                            add(deck, base_tags, f"Solve recurrence: {term}", defin + (f" (e.g. {extra})" if extra else ""), source, section)
                        elif "code shape" in hjoin:
                            add(deck, base_tags, f"Time complexity of: {term}?", defin, source, section)
                        elif "sum" in hjoin and "closed form" in hjoin:
                            notes = cols[2] if len(cols) > 2 else ""
                            add(deck, base_tags, f"Closed form of {term}?", defin + (f" ({notes})" if notes else ""), source, section)
                        elif "formula" in hjoin:
                            add(deck, base_tags, f"Formula for {term}?", defin, source, section)
                        elif "problem" in hjoin and "result" in hjoin:
                            add(deck, base_tags, f"Classic result: {term}?", defin, source, section)
                        else:
                            add(deck, base_tags, f"What is {term}?", defin, source, section)
                i += 1
            continue
        i += 1


def build() -> None:
    cards.clear()
    parse_glossary_tables((DOC / "DSA_JARGON.md").read_text(), "doc/DSA_JARGON.md", "jargon", ["jargon", "dsa"])

    drill = (DOC / "write" / "DRILL_CONCEPTS.md").read_text()
    section = ""
    for line in drill.splitlines():
        if line.startswith("## "):
            section = line[3:].strip()
            continue
        m = re.match(r"- \*\*Core idea:\*\* (.+)", line)
        if m and section.startswith("Drill"):
            add("patterns", ["patterns", "write", "core-idea"], f"Core idea of {section}?", m.group(1), "doc/write/DRILL_CONCEPTS.md", section)
        m = re.match(r"- \*\*When this pattern appears:\*\* (.+)", line)
        if m and section.startswith("Drill"):
            add("patterns", ["patterns", "write", "trigger"], f"When does “{section}” appear?", m.group(1), "doc/write/DRILL_CONCEPTS.md", section)
        m = re.match(r"  - \*\*([^*]+):\*\* (.+)", line)
        if m and section.startswith("Drill"):
            add("patterns", ["patterns", "write"], f"In {section}: what is “{m.group(1)}”?", m.group(2), "doc/write/DRILL_CONCEPTS.md", section)
        m = re.match(r"- (.+) -> \*\*(.+)\*\*", line)
        if m:
            add("patterns", ["patterns", "picker"], f"Pattern picker: {m.group(1).strip()} → ?", m.group(2).strip(), "doc/write/DRILL_CONCEPTS.md", "Fast Pattern Picker")

    reading = (DOC / "read" / "READING_PATTERNS.md").read_text()
    for line in reading.splitlines():
        m = re.match(r"\|\s*(\d+)\s*\|\s*\*\*([^*]+)\*\*\s*\|\s*([^|]+)\|\s*([^|]+)\|", line)
        if m:
            add(
                "reading",
                ["reading", "6-pass"],
                f"6-pass method Pass {m.group(1)} ({m.group(2).strip()}, ~{m.group(3).strip()}): what question?",
                m.group(4).strip(),
                "doc/read/READING_PATTERNS.md",
                "The 6-pass read",
            )
    parse_glossary_tables(reading, "doc/read/READING_PATTERNS.md", "reading", ["reading"])
    add(
        "reading",
        ["reading", "ownership"],
        "When do you “own” a reading snippet? List the 5 checks.",
        "Signature+mutation in 15s; name the pattern in one breath; trace one sample; restate ask in one sentence; give tight time/space bounds",
        "doc/read/READING_PATTERNS.md",
        "Ownership",
    )

    backend = (DOC / "backend" / "DRILL_CONCEPTS.md").read_text()
    parse_glossary_tables(backend, "doc/backend/DRILL_CONCEPTS.md", "backend", ["backend"])
    section = subsection = ""
    for line in backend.splitlines():
        if line.startswith("## "):
            section, subsection = line[3:].strip(), ""
            continue
        if line.startswith("### "):
            subsection = line[4:].strip()
            continue
        m = re.match(r"- \*\*([^:*]+):\*\* (.+)", line)
        if m:
            topic = subsection or section
            add("backend", ["backend"], f"Backend — {topic}: {m.group(1)}?", m.group(2), "doc/backend/DRILL_CONCEPTS.md", topic)

    for deck, tags, front, back, sec in [
        ("backend", ["backend", "jwt"], "Walk through JWT auth flow.", "Client logs in → server issues access (+ optional refresh). Client sends Authorization: Bearer <token>. Middleware validates signature, exp, iss, aud; attaches claims to context.", "JWT"),
        ("backend", ["backend", "resilience"], "Circuit breaker state machine?", "CLOSED → failures ≥ threshold → OPEN (fail fast) → after timeout → HALF-OPEN (probe) → success → CLOSED; failure → OPEN", "Circuit breaker"),
        ("backend", ["backend", "resilience"], "Exponential backoff formula + rules?", "delay = min(cap, base * 2^attempt) + jitter. Only retry transient errors. Require idempotent ops or Idempotency-Key.", "Retry"),
        ("backend", ["backend", "cap"], "CAP during partition — what do most web APIs favor?", "Choose CP or AP. Most web APIs favor availability + eventual consistency for reads.", "CAP"),
        ("backend", ["backend", "saga"], "Saga vs 2PC?", "2PC: strong consistency, fragile under partitions. Saga: local txs + compensating actions; common in microservices.", "Saga"),
        ("backend", ["backend", "webrtc"], "WebRTC: signaling vs media paths?", "Signaling (WebSocket): SDP offer/answer, ICE. Media (UDP): encrypted RTP peer-to-peer, not through your server.", "WebRTC"),
        ("backend", ["backend", "webrtc"], "STUN vs TURN?", "STUN discovers public IP/port (cheap). TURN relays when P2P fails (costly; symmetric NAT / strict firewalls).", "STUN/TURN"),
        ("backend", ["backend", "api"], "Idempotency rules for REST verbs?", "PUT/DELETE idempotent; POST needs Idempotency-Key for retries.", "REST"),
        ("backend", ["backend", "db"], "Composite index column order rule?", "Equality filters first, then range.", "Indexing"),
        ("backend", ["backend", "db"], "Query optimization workflow (4 steps)?", "EXPLAIN ANALYZE → rewrite query → add/adjust indexes → cache hot reads / tune pool.", "Query opt"),
        ("backend", ["backend", "nat"], "NAT debug checklist for WebRTC (5 steps)?", "1) Signaling open? 2) ICE candidates gathered? 3) Connection connected vs failed? 4) Try TURN if only host. 5) Firewall UDP/TURN TLS.", "NAT"),
        ("backend", ["backend", "airflow"], "Airflow recovery when a task fails?", "Clear failed task, rerun from failed node. Tasks must be idempotent; use retries + retry_delay.", "Airflow"),
        ("backend", ["backend", "jenkins"], "Typical Jenkins pipeline stages?", "checkout → install → lint → unit test → build → integration → security scan → deploy staging → smoke → deploy prod", "Jenkins"),
        ("backend", ["backend", "docker"], "Liveness vs readiness health checks?", "Liveness: restart if dead. Readiness: whether to send traffic.", "Docker"),
        ("backend", ["backend", "zatca"], "ZATCA invoice submission flow?", "Generate XML → XSD validate → cryptographic sign → submit over TLS → store response + audit (who/when/payload hash).", "ZATCA"),
        ("backend", ["backend", "go"], "Go concurrency primitives for servers?", "goroutine + channel for worker pools; context.Context for cancel/deadlines; sync.Mutex for shared state (prefer channels for ownership).", "Go"),
    ]:
        add(deck, tags, front, back, "doc/backend/DRILL_CONCEPTS.md", sec)

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

    math = (DOC / "write" / "MATH_CONCEPTS.md").read_text()
    parse_glossary_tables(math, "doc/write/MATH_CONCEPTS.md", "math", ["math"])
    fs = re.search(r"## Formula Sheet.*?\n```(.*?)```", math, re.S)
    if fs:
        for line in fs.group(1).splitlines():
            line = line.strip()
            if ":" not in line:
                continue
            name, formula = line.split(":", 1)
            add("math", ["math", "formula"], f"Formula: {name.strip()}?", formula.strip(), "doc/write/MATH_CONCEPTS.md", "Formula Sheet")
    section = ""
    for line in math.splitlines():
        if re.match(r"^## \d+", line):
            section = re.sub(r"^\d+\s*[—-]\s*", "", line[3:].strip())
            continue
        m = re.match(r"- \*\*Core idea:\*\* (.+)", line)
        if m and section:
            add("math", ["math", "core-idea"], f"Math core idea — {section}?", m.group(1), "doc/write/MATH_CONCEPTS.md", section)
        m = re.match(r"- \*\*When this appears:\*\* (.+)", line)
        if m and section:
            add("math", ["math", "trigger"], f"When does “{section}” appear in problems?", m.group(1), "doc/write/MATH_CONCEPTS.md", section)

    for expr, meaning in [
        ("n & (n - 1)", "clears lowest set bit"),
        ("n & (-n)", "isolates lowest set bit (two's complement)"),
        ("1 << k", "equals 2^k"),
        ("n > 0 && (n & (n-1)) == 0", "tests whether n is a power of two"),
    ]:
        add("math", ["math", "bits"], f"Bit trick: what does {expr} do?", meaning, "doc/write/MATH_CONCEPTS.md", "Bits")

    add("math", ["math", "master"], "Master theorem case 1 (leaf-heavy)?", "If f(n) = O(n^(log_b a - ε)) for ε>0 → T(n) = Θ(n^(log_b a))", "doc/write/MATH_CONCEPTS.md", "Master")
    add("math", ["math", "master"], "Master theorem case 2 (balanced)?", "If f(n) = Θ(n^(log_b a) log^k n) → T(n) = Θ(n^(log_b a) log^(k+1) n)", "doc/write/MATH_CONCEPTS.md", "Master")
    add("math", ["math", "master"], "Master theorem case 3 (root-heavy)?", "If f(n) = Ω(n^(log_b a + ε)) and regularity → T(n) = Θ(f(n))", "doc/write/MATH_CONCEPTS.md", "Master")
    add("math", ["math", "heap"], "Heap array indexing (0-based): parent / left / right?", "parent (i-1)/2; left 2i+1; right 2i+2", "doc/write/MATH_CONCEPTS.md", "Heaps")
    add("math", ["math", "gcd"], "Euclidean GCD algorithm?", "gcd(a,b) = gcd(b, a mod b) until b==0; gcd(a,0)=a. O(log min(a,b))", "doc/write/MATH_CONCEPTS.md", "Number theory")
    add("math", ["math", "catalan"], "Catalan number formula?", "C_n = (1/(n+1)) * C(2n, n)", "doc/write/MATH_CONCEPTS.md", "Combinatorics")

    add("meta", ["meta", "srs"], "Spaced repetition schedule after solving a problem?", "Day +1, then +3, then +7, then +21", "doc/write/STUDY_PLAN.md", "SRS")

    seen: set[str] = set()
    uniq: list[dict] = []
    for c in cards:
        if c["id"] in seen:
            continue
        seen.add(c["id"])
        uniq.append(c)

    OUT.mkdir(parents=True, exist_ok=True)
    by: dict[str, list] = {}
    for c in uniq:
        by.setdefault(c["deck"], []).append(c)
    manifest = []
    for deck, items in sorted(by.items()):
        (OUT / f"{deck}.json").write_text(json.dumps(items, indent=2, ensure_ascii=False) + "\n")
        manifest.append({"deck": deck, "count": len(items), "file": f"decks/{deck}.json"})
        print(f"{deck}: {len(items)}")
    (ROOT / "cards" / "manifest.json").write_text(json.dumps({"total": len(uniq), "decks": manifest}, indent=2) + "\n")
    print("TOTAL", len(uniq))


if __name__ == "__main__":
    build()
