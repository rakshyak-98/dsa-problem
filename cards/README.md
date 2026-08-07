# Spaced repetition cards

Flashcards for terminal review with a +1 / +3 / +7 / +21 day schedule (same ladder as [`doc/write/STUDY_PLAN.md`](../doc/write/STUDY_PLAN.md)).

**Sources**

| Source | Generator | Deck prefix |
|--------|-----------|-------------|
| `doc/` (this repo) | `generate_cards.py` | `jargon`, `patterns`, `math`, … |
| [Back2Basics](https://github.com/rakshyak-98/Back2Basics) vault | `generate_back2basics.py` | `b2b-*` |

## Quick start

```bash
go run . -- --track cards
go run . -- --track cards --due
go run . -- --track cards --stats
go run . -- --track cards --catalog
go run . -- --track=cards --deck=b2b-docker --limit=15
go run . -- --track cards --tag=triage --new=10
```

Or: `go run ./bin/study_cards -- …`

Note: at the root CLI, `-t` means `--track`. Use `--tag=…` to filter card tags.

During review: **Enter** shows the answer, then rate **1** again / **2** hard / **3** good / **4** easy (`q` quits and saves).

## DSA decks (`doc/`)

| Deck | Source | Practice |
|------|--------|----------|
| `jargon` | `doc/DSA_JARGON.md` | Plain-English DSA terms |
| `patterns` | `doc/write/DRILL_CONCEPTS.md` | Reflex triggers + core ideas |
| `math` | `doc/write/MATH_CONCEPTS.md` | Formulas, complexity, number theory |
| `reading` | `doc/read/READING_PATTERNS.md` | 6-pass method + traps |
| `backend` | `doc/backend/DRILL_CONCEPTS.md` | Interview concepts |
| `star` | `doc/backend/STAR_STORIES.md` | STAR rehearsals |
| `meta` | `doc/write/STUDY_PLAN.md` | Study-system facts |

## Back2Basics decks (`b2b-*`)

One deck per vault top-level folder, plus `b2b-oncall` from `INDEX.md` symptom routing.

Useful filters:

```bash
go run . -- --track cards --deck=b2b-database --tag=triage
go run . -- --track cards --deck=b2b-linux --tag=gotcha
go run . -- --track cards --deck=b2b-oncall
go run . -- --track cards --tag=mental-model --limit=20
go run . -- --track cards --deck=b2b-design-pattern --tag=decision
```

Card types extracted from each note: definition (one-liner), mental model, triage tables, decision tables, gotchas, when-not-to-use.

Counts: `manifest.json` / `go run . -- --track cards --catalog`.

## Layout

```
cards/
├── README.md
├── manifest.json
├── generate_cards.py           # from doc/
├── generate_back2basics.py     # from ../Back2Basics (or BACK2BASICS_ROOT)
├── decks/
│   ├── jargon.json             # DSA
│   ├── b2b-docker.json         # Back2Basics
│   └── …
└── .srs_progress.json          # local (gitignored)
```

## Regenerate

```bash
python3 cards/generate_cards.py              # DSA notes in doc/
python3 cards/generate_back2basics.py        # needs ../Back2Basics or BACK2BASICS_ROOT
```

Progress is keyed by card `id`, so regenerating keeps history for unchanged fronts.

## CLI reference

```
go run . -- --track cards --help
go run ./bin/study_cards -- --help
```
