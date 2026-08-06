# Drills — start here

Everything you **practice** is exposed here. Files live under `study_play/practice/` and `study_code/practice/`; this folder is the front door.

## Layout

```
drills/
├── write/                 → study_play/practice/write/
│   ├── core5/             # daily Core 5 essentials
│   ├── reflex/            # weekday + bonus reflex drills
│   └── variants/          # medium pattern variants
├── read/                  → study_code/practice/read/
│   ├── core/              # Core Read 3 (every day)
│   ├── weekday/           # specialty reading drills
│   └── answers/           # peek only after honest attempt
├── solutions/             → study_play/_support/solutions_reference/
│   ├── reflex/            # runnable Go solution per reflex drill
│   ├── core5.md           # Core 5 quick reference
│   └── variants.md        # variant drill quick reference
└── tracker/               → study_play/practice/tracker/
    └── study_tracker.html
```

## Backend interview prep (resume drills)

```bash
cd study_backend && go run . -- --cram    # cram schedule until interview
cd study_backend && go run . -- --run     # validate answers
```

Guide: [`study_backend/docs/START_HERE.md`](../study_backend/docs/START_HERE.md)

## Daily flow (recommended)

```bash
cd daily && go run .           # unified plan: read + write
cd daily && go run . -- --run  # check today's answers
```

## Write drills

```bash
cd study_play && go run .                    # today's plan
go run -C study_play/practice/write/core5 .                    # Core 5
go run -C study_play/practice/write/reflex/02_hashing_reflex . # specialty
cd study_play && go run . -- --run           # test + log
```

Guide: [`study_play/docs/START_HERE.md`](../study_play/docs/START_HERE.md)  
Math reference: [`study_play/docs/MATH_CONCEPTS.md`](../study_play/docs/MATH_CONCEPTS.md)

## Read drills

```bash
cd study_code && go run .
go run -C study_code/practice/read/core/00_core_read .
go run -C study_code/practice/read/weekday/03_name_the_pattern .
cd study_code && go run . -- --run
```

Guide: [`study_code/docs/START_HERE.md`](../study_code/docs/START_HERE.md)

## Solutions (after honest attempt)

**Reflex Go solutions:** [`solutions/reflex/`](solutions/reflex/) — runnable `main.go` per drill  
**Quick notes:** [`solutions/*.md`](solutions/) — triggers and bugs  
Read answer keys: [`read/answers/`](read/answers/)

## Track progress

Open [`tracker/study_tracker.html`](tracker/study_tracker.html).

## Encapsulated (not in drills/)

| What | Where |
|------|--------|
| CLI tooling | `study_play/`, `study_code/`, `daily/` |
| Templates & blanks | `study_play/_support/` (solutions exposed at `drills/solutions/`) |
| Study plans | `study_play/docs/`, `study_code/docs/` |
| Problem catalog | `problems/CATEGORIES.md` |
| Visualizer | `visualizer/` |
