# Spaced repetition cards

Backend interview flashcards for terminal review with a +1 / +3 / +7 / +21 day schedule.

**Sources:** `doc/backend/DRILL_CONCEPTS.md`, `doc/backend/STAR_STORIES.md`, `doc/backend/WEEKLY_REVISION.md`

## Quick start

```bash
go run . -- --track cards
go run . -- --track cards --due
go run . -- --track cards --stats
go run . -- --track cards --catalog
go run . -- --track cards --deck=backend --limit=15
go run . -- --track cards --deck=star --due
go run . -- --track cards --tag=trigger --new=10
```

Or: `go run ./bin/study_cards -- …`

Note: at the root CLI, `-t` means `--track`. Use `--tag=…` to filter card tags.

During review: **Enter** shows the answer, then rate **1** again / **2** hard / **3** good / **4** easy (`q` quits and saves).

## Decks

| Deck | Source | Practice |
|------|--------|----------|
| `backend` | `doc/backend/DRILL_CONCEPTS.md` + weekly revision | Interview triggers, concept Q&A |
| `star` | `doc/backend/STAR_STORIES.md` | STAR story rehearsals |

Cards are **question-style only** — no full-note dumps from vault markdown.

## Layout

```
cards/
├── README.md
├── manifest.json
├── generate_cards.py           # backend + star from doc/backend/
├── decks/
│   ├── backend.json
│   └── star.json
└── .srs_progress.json          # local (gitignored)
```

## Regenerate

```bash
python3 cards/generate_cards.py
```

Progress is keyed by card `id`, so regenerating keeps history for unchanged fronts.

## CLI reference

```
go run . -- --track cards --help
go run ./bin/study_cards -- --help
```
