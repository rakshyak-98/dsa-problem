# Code Reading Drills

Train **reading** DSA code the way `study_play` trains **writing** it.

Most interview and review time is spent understanding existing code — yours, an editorial, or a teammate’s. These drills build systematic reading habits: structure scan → trace → pattern → ask → complexity → bugs.

---

## Why this exists

Writing reflexes alone leave a gap:

- You recognize a pattern name but freeze when staring at someone else’s loops
- You skip tracing and misread off-by-one / pointer moves
- You cannot reconstruct the problem from a solution (weak “ask → code” link)

`study_code` closes that gap with short, timed reading exercises.

---

## How it fits with `study_play`

| Folder | Skill | Mode |
|--------|--------|------|
| `study_play/` | Write patterns blind | Fill `TODO: REFLEX` |
| `study_code/` | Read patterns fluently | Answer questions about opaque snippets |

Suggested order on a study day:

1. `study_play` Core 5 (write)
2. Today’s `study_code` specialty (read) — 15–20 min
3. Primary problem (understand → solve)

---

## Quick start

```bash
cd study_code
go run .              # today's reading drill + prompts
go run . -- --run     # run today's drill answers
go run . -- --catalog # all 7 drills
```

Day-one flow: [`START_HERE.md`](./START_HERE.md).  
Reading method: [`READING_PATTERNS.md`](./READING_PATTERNS.md).  
Daily clock: [`DAILY_20MIN_READING.md`](./DAILY_20MIN_READING.md).

---

## Folder map

```
drills/read/                   ← practice files (front door)
├── core/00_core_read/
├── weekday/01_scan_structure/ … 08_math_concepts/
└── answers/
bin/study_code/                ← CLI entry
doc/read/                      ← guides (START_HERE, patterns, ritual)
```

---

## Rule

**Read before you peek.** Fill answers in the drill file, run tests, then open `drills/read/answers/` only for fails — and re-read the snippet once more after correcting.
