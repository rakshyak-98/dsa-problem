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
└── tracker/               → study_play/practice/tracker/
    └── study_tracker.html
```

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

## Read drills

```bash
cd study_code && go run .
go run -C study_code/practice/read/core/00_core_read .
go run -C study_code/practice/read/weekday/03_name_the_pattern .
cd study_code && go run . -- --run
```

Guide: [`study_code/docs/START_HERE.md`](../study_code/docs/START_HERE.md)

## Track progress

Open [`tracker/study_tracker.html`](tracker/study_tracker.html).

## Encapsulated (not in drills/)

| What | Where |
|------|--------|
| CLI tooling | `study_play/`, `study_code/`, `daily/` |
| Templates & solutions | `study_play/_support/` |
| Study plans | `study_play/docs/`, `study_code/docs/` |
| Problem catalog | `problems/CATEGORIES.md` |
| Visualizer | `visualizer/` |
