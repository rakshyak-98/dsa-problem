# Drills — start here

Everything you **practice** lives under `drills/`, organized by topic. CLI helpers live in `bin/`.

**New to DSA words?** See [`DSA_JARGON.md`](DSA_JARGON.md) for one-sentence plain-English definitions.

## Layout

```
drills/
├── write/                 # DSA reflex writing
│   ├── core5/             # daily Core 5 essentials
│   ├── reflex/            # weekday + bonus reflex drills
│   └── variants/          # medium pattern variants
├── read/                  # code reading
│   ├── core/              # Core Read 3 (every day)
│   ├── weekday/           # specialty reading drills
│   └── answers/           # peek only after honest attempt
├── backend/               # backend interview prep
│   ├── explain/           # verbal concept drills
│   ├── write/             # Go reflex implementations
│   └── scenario/          # STAR / mock interview prompts
├── solutions/             # write drill solutions (peek after attempt)
│   ├── reflex/            # runnable Go solution per reflex drill
│   ├── core5.md
│   └── variants.md
└── tracker/               # browser study tracker
    └── study_tracker.html
```

## Daily flow (recommended)

```bash
go run .                      # daily drill: read + write (DSA track)
go run . -- --run core          # check core answers
go run . -- --run reflex        # check reflex specialty answers
go run . -- --drill core        # core only (Core Read 3 + Core 5)
go run . -- --drill reflex       # today's specialty only
go run . -- --track backend   # backend interview track
go run . -- --track read      # reading only
go run . -- --track write     # writing only
```

## Write drills

```bash
go run ./bin/study_play                    # today's plan
go run -C drills/write/core5 .             # Core 5
go run -C drills/write/reflex/02_hashing_reflex .
go run ./bin/study_play -- --run           # test + log
```

Guide: [`doc/write/START_HERE.md`](../doc/write/START_HERE.md)  
Math reference: [`doc/write/MATH_CONCEPTS.md`](../doc/write/MATH_CONCEPTS.md)

## Read drills

```bash
go run ./bin/study_code
go run -C drills/read/core/00_core_read .
go run -C drills/read/weekday/03_name_the_pattern .
go run ./bin/study_code -- --run
```

Guide: [`doc/read/START_HERE.md`](../doc/read/START_HERE.md)

## Backend interview prep

```bash
go run ./bin/study_backend -- --cram    # cram schedule until interview
go run ./bin/study_backend -- --run     # validate answers
```

Guide: [`doc/backend/START_HERE.md`](../doc/backend/START_HERE.md)

## Solutions (after honest attempt)

**Reflex Go solutions:** [`solutions/reflex/`](solutions/reflex/) — runnable `main.go` per drill  
**Quick notes:** [`solutions/*.md`](solutions/) — triggers and bugs  
Read answer keys: [`read/answers/`](read/answers/)

## Track progress

Open [`tracker/study_tracker.html`](tracker/study_tracker.html).

## Where things live

| What | Where |
|------|--------|
| Documentation | `doc/` |
| Practice files | `drills/` (this folder) |
| CLI tooling | `bin/` |
| Problem catalog | `reference/problems/CATEGORIES.md` |
| Visualizer | `reference/visualizer/` |
