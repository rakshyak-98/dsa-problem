#!/usr/bin/env python3
"""Extract spaced-repetition cards from the Back2Basics Obsidian vault.

Default vault path: sibling ../Back2Basics (override with BACK2BASICS_ROOT).
Writes decks as cards/decks/b2b-<folder>.json and merges into manifest.json.
Does not delete existing DSA decks from doc/.
"""
from __future__ import annotations

import hashlib
import json
import os
import re
from collections import defaultdict
from pathlib import Path

HERE = Path(__file__).resolve().parent
DSA_ROOT = HERE.parent
OUT = HERE / "decks"
MANIFEST = HERE / "manifest.json"

SKIP_NAMES = {
    "readme.md",
    "index.md",
    "notes_standard.md",
    "agents.md",
    "worklog.md",
    "interview.md",
    "useful prompt for learning with ai chat.md",
}

SKIP_DIRS = {".git", ".obsidian", "Canvas", "Errors"}

cards: list[dict] = []


def cid(*parts: str) -> str:
    return hashlib.sha1("|".join(parts).encode()).hexdigest()[:12]


def clean(s: str) -> str:
    s = re.sub(r"\[\[([^\]|]+)(?:\|[^\]]+)?\]\]", r"\1", s)  # wikilinks
    s = re.sub(r"\*\*([^*]+)\*\*", r"\1", s)
    s = re.sub(r"`([^`]+)`", r"\1", s)
    s = re.sub(r"!\[.*?\]\(.*?\)", "", s)
    s = re.sub(r"\[([^\]]+)\]\([^)]+\)", r"\1", s)
    s = re.sub(r">\s*\[!WARNING\]\s*", "", s, flags=re.I)
    s = re.sub(r">\s*\[!\w+\]\s*", "", s)
    s = re.sub(r"^>\s?", "", s, flags=re.M)
    s = re.sub(r"\s+", " ", s).strip(" -\t\n")
    return s


def slug_deck(top: str) -> str:
    s = top.strip().lower()
    s = s.replace("&", " and ")
    s = re.sub(r"[^a-z0-9]+", "-", s).strip("-")
    if not s:
        s = "misc"
    return f"b2b-{s}"


def add(deck: str, tags: list[str], front: str, back: str, source: str, section: str = "") -> None:
    front, back = clean(front), clean(back)
    if not front or not back or front == back:
        return
    if len(front) < 10 or len(back) < 12:
        return
    # Generic table dumps with tiny answers are low value
    if front.startswith("What about") and len(back) < 24:
        return
    if len(back) > 600:
        back = back[:597].rstrip() + "…"
    if len(front) > 240:
        front = front[:237].rstrip() + "…"
    # skip checklist fluff / empty stubs
    if front.lower().startswith("what is related"):
        return
    if front.startswith("What is ") and back.startswith("http"):
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


def parse_tables(text: str) -> list[tuple[str, list[str], list[list[str]]]]:
    """Return list of (section, headers, rows) for markdown tables."""
    lines = text.splitlines()
    section = ""
    out = []
    i = 0
    while i < len(lines):
        line = lines[i]
        if line.startswith("## "):
            section = line[3:].strip()
            i += 1
            continue
        if line.startswith("### "):
            section = line[4:].strip()
            i += 1
            continue
        if "|" in line and i + 1 < len(lines) and re.match(r"\|\s*-+", lines[i + 1]):
            headers = [h.strip() for h in line.strip("|").split("|")]
            i += 2
            rows = []
            while i < len(lines) and lines[i].strip().startswith("|"):
                cols = [c.strip() for c in lines[i].strip("|").split("|")]
                if any(cols) and not all(re.fullmatch(r":?-+:?", c or "") for c in cols):
                    rows.append(cols)
                i += 1
            if rows:
                out.append((section, headers, rows))
            continue
        i += 1
    return out


def first_heading(text: str) -> str:
    for line in text.splitlines():
        if line.startswith("# "):
            return line[2:].strip()
    return ""


def one_liner(text: str) -> str:
    """First blockquote after the H1, or first non-link prose line."""
    lines = text.splitlines()
    past_title = False
    buf = []
    for line in lines:
        if line.startswith("# "):
            past_title = True
            continue
        if not past_title:
            continue
        if line.startswith("[[") and not buf:
            continue
        if line.startswith(">"):
            # gather consecutive quote lines
            q = re.sub(r"^>\s?", "", line).strip()
            if q.startswith("[!"):
                continue
            if q:
                buf.append(q)
            continue
        if buf:
            break
        if line.startswith("## "):
            break
        if line.strip() and not line.startswith("```"):
            # only use if short definition-like
            s = line.strip()
            if len(s) < 220 and not s.startswith("|") and not s.startswith("- ["):
                return s
            break
    return " ".join(buf)


def extract_mental_model(text: str) -> str:
    m = re.search(
        r"##\s+Mental model\s*\n+(.*?)(?=\n##\s|\Z)",
        text,
        re.S | re.I,
    )
    if not m:
        return ""
    body = m.group(1)
    # strip code fences
    body = re.sub(r"```.*?```", "", body, flags=re.S)
    # take first 1–3 prose paragraphs / bullets before a table
    chunks = []
    for line in body.splitlines():
        if line.strip().startswith("|"):
            break
        if line.startswith("###"):
            break
        s = line.strip()
        if not s or s.startswith("[["):
            if chunks:
                break
            continue
        if s.startswith("- ") or s.startswith("* "):
            chunks.append(s[2:])
        elif not s.startswith("#"):
            chunks.append(s)
        if len(" ".join(chunks)) > 350:
            break
    return " ".join(chunks)


def extract_gotchas(text: str, deck: str, tags: list[str], source: str, title: str) -> None:
    m = re.search(r"##\s+Gotchas\s*\n+(.*?)(?=\n##\s|\Z)", text, re.S | re.I)
    if not m:
        return
    body = m.group(1)
    # callout blocks and bullets
    for block in re.split(r"\n(?=>\s*\[!|\n- )", body):
        s = clean(block)
        if len(s) < 20:
            continue
        add(
            deck,
            tags + ["gotcha"],
            f"Gotcha — {title}: what goes wrong?",
            s,
            source,
            "Gotchas",
        )


def extract_when_not(text: str, deck: str, tags: list[str], source: str, title: str) -> None:
    m = re.search(
        r"##\s+When NOT to use\s*\n+(.*?)(?=\n##\s|\Z)",
        text,
        re.S | re.I,
    )
    if not m:
        return
    body = clean(re.sub(r"```.*?```", "", m.group(1), flags=re.S))
    if len(body) < 15:
        return
    add(
        deck,
        tags + ["when-not"],
        f"When should you NOT use {title}?",
        body,
        source,
        "When NOT to use",
    )


def extract_definition_bullets(text: str, deck: str, tags: list[str], source: str, section_hint: str) -> None:
    section = section_hint
    for line in text.splitlines():
        if line.startswith("## "):
            section = line[3:].strip()
            continue
        m = re.match(r"- \*\*([^*]+)\*\*\s*[—–\-:]\s*(.+)", line)
        if m:
            term, defin = m.group(1).strip(), m.group(2).strip()
            add(deck, tags, f"What is {term}?", defin, source, section)
            continue
        m = re.match(r"- \*\*([^*]+)\*\*\s+—\s*(.+)", line)
        if m:
            add(deck, tags, f"What is {m.group(1).strip()}?", m.group(2).strip(), source, section)


def table_to_cards(deck: str, tags: list[str], source: str, section: str, headers: list[str], rows: list[list[str]], title: str) -> None:
    h = [x.lower() for x in headers]
    hjoin = " ".join(h)

    def col(*names: str) -> int | None:
        for i, header in enumerate(h):
            for n in names:
                if n in header:
                    return i
        return None

    # Triage: Symptom | Check | Fix
    if "symptom" in hjoin and ("check" in hjoin or "fix" in hjoin):
        si = col("symptom")
        ci = col("check", "diag")
        fi = col("fix")
        for row in rows:
            if si is None or si >= len(row):
                continue
            symptom = row[si]
            parts = []
            if ci is not None and ci < len(row) and row[ci]:
                parts.append(f"Check: {row[ci]}")
            if fi is not None and fi < len(row) and row[fi]:
                parts.append(f"Fix: {row[fi]}")
            if not parts:
                continue
            add(
                deck,
                tags + ["triage"],
                f"Triage ({title}): {symptom} — what do you do?",
                " · ".join(parts),
                source,
                section or "Triage",
            )
        return

    # Decision: You see… | Reach for | Not this
    if ("you see" in hjoin or "variation" in hjoin or "when" in hjoin) and (
        "reach" in hjoin or "pattern" in hjoin or "use" in hjoin or "go to" in hjoin
    ):
        ai = 0
        bi = 1 if len(headers) > 1 else 0
        for i, header in enumerate(h):
            if any(k in header for k in ("see", "variation", "when", "need", "symptom")):
                ai = i
            if any(k in header for k in ("reach", "pattern", "use", "go to", "note")):
                bi = i
        for row in rows:
            if ai >= len(row) or bi >= len(row):
                continue
            extra = ""
            if len(row) > 2 and row[2]:
                extra = f" (not: {row[2]})"
            add(
                deck,
                tags + ["decision"],
                f"When you see “{row[ai]}” → what do you reach for?",
                row[bi] + extra,
                source,
                section or "Decision",
            )
        return

    # Property / Question / Mechanism style (ACID)
    if "property" in hjoin or "question" in hjoin:
        pi = col("property", "term", "level", "principle", "name") or 0
        qi = col("question", "meaning", "answers")
        mi = col("mechanism", "typical", "how")
        for row in rows:
            if pi >= len(row):
                continue
            term = row[pi]
            bits = []
            if qi is not None and qi < len(row) and row[qi]:
                bits.append(row[qi])
            if mi is not None and mi < len(row) and row[mi]:
                bits.append(row[mi])
            if not bits and len(row) > 1:
                bits.append(row[1])
            add(deck, tags, f"In {title}: what is {term}?", " — ".join(bits), source, section)
        return

    # Generic 2-col glossary — skip very wide dump tables (config knobs with tiny cells)
    if len(headers) >= 2:
        for row in rows:
            if len(row) < 2 or not row[0] or not row[1]:
                continue
            if row[0].startswith("✅") or row[0].startswith("[ ]"):
                continue
            answer = " — ".join(c for c in row[1:] if c)
            if len(answer) < 24:
                continue
            add(
                deck,
                tags,
                f"What about “{row[0]}” ({title})?",
                answer,
                source,
                section,
            )


def process_note(path: Path, vault: Path) -> None:
    rel = path.relative_to(vault)
    if path.name.lower() in SKIP_NAMES:
        return
    if any(p in SKIP_DIRS for p in rel.parts):
        return

    text = path.read_text(encoding="utf-8", errors="ignore")
    if len(text.strip()) < 80:
        return

    top = rel.parts[0] if len(rel.parts) > 1 else "general"
    deck = slug_deck(top if len(rel.parts) > 1 else "general")
    tag_top = re.sub(r"[^a-z0-9]+", "-", top.lower()).strip("-") or "general"
    tags = ["back2basics", tag_top]
    source = f"Back2Basics/{rel.as_posix()}"
    title = first_heading(text) or path.stem
    title = clean(title)

    # Core definition card
    blurb = one_liner(text)
    if blurb:
        add(deck, tags + ["definition"], f"What is {title}?", blurb, source, title)

    mental = extract_mental_model(text)
    if mental and clean(mental) != clean(blurb or ""):
        add(
            deck,
            tags + ["mental-model"],
            f"Mental model: {title}?",
            mental,
            source,
            "Mental model",
        )

    for section, headers, rows in parse_tables(text):
        table_to_cards(deck, tags, source, section, headers, rows, title)

    extract_definition_bullets(text, deck, tags, source, title)
    extract_gotchas(text, deck, tags, source, title)
    extract_when_not(text, deck, tags, source, title)


def process_index(vault: Path) -> None:
    idx = vault / "INDEX.md"
    if not idx.exists():
        return
    text = idx.read_text(encoding="utf-8", errors="ignore")
    deck = "b2b-oncall"
    for section, headers, rows in parse_tables(text):
        h = [x.lower() for x in headers]
        if "symptom" in " ".join(h) or "need" in " ".join(h):
            for row in rows:
                if len(row) < 2:
                    continue
                add(
                    deck,
                    ["back2basics", "oncall", "index"],
                    f"On-call: {row[0]} — which notes?",
                    row[1],
                    "Back2Basics/INDEX.md",
                    section or "On-call",
                )


def vault_root() -> Path:
    env = os.environ.get("BACK2BASICS_ROOT")
    if env:
        return Path(env).expanduser().resolve()
    sibling = DSA_ROOT.parent / "Back2Basics"
    if sibling.is_dir():
        return sibling
    raise SystemExit(
        "Back2Basics vault not found. Set BACK2BASICS_ROOT or place it next to dsa-problem."
    )


def write_decks() -> None:
    OUT.mkdir(parents=True, exist_ok=True)

    # Remove previous b2b-* decks only
    for old in OUT.glob("b2b-*.json"):
        old.unlink()

    by: dict[str, list] = defaultdict(list)
    seen: set[str] = set()
    for c in cards:
        if c["id"] in seen:
            continue
        seen.add(c["id"])
        by[c["deck"]].append(c)

    # Preserve non-b2b decks in manifest
    existing = []
    if MANIFEST.exists():
        data = json.loads(MANIFEST.read_text())
        for d in data.get("decks", []):
            if not str(d.get("deck", "")).startswith("b2b-"):
                existing.append(d)

    new_manifest = list(existing)
    total_b2b = 0
    for deck, items in sorted(by.items()):
        (OUT / f"{deck}.json").write_text(json.dumps(items, indent=2, ensure_ascii=False) + "\n")
        new_manifest.append({"deck": deck, "count": len(items), "file": f"decks/{deck}.json"})
        total_b2b += len(items)
        print(f"{deck}: {len(items)}")

    # Recount totals from files on disk
    total = 0
    final = []
    for entry in sorted(OUT.glob("*.json")):
        items = json.loads(entry.read_text())
        name = entry.stem
        final.append({"deck": name, "count": len(items), "file": f"decks/{name}.json"})
        total += len(items)
        if not name.startswith("b2b-"):
            print(f"(keep) {name}: {len(items)}")

    MANIFEST.write_text(json.dumps({"total": total, "decks": final}, indent=2) + "\n")
    print(f"Back2Basics cards: {total_b2b}")
    print(f"TOTAL all decks: {total}")


def main() -> None:
    cards.clear()
    vault = vault_root()
    print(f"Vault: {vault}")
    process_index(vault)
    for path in sorted(vault.rglob("*.md")):
        process_note(path, vault)
    write_decks()


if __name__ == "__main__":
    main()
