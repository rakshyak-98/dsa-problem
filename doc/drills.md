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
├── leetcode/              # daily 10-question LeetCode practice sets
└── solutions/             # write drill solutions (peek after attempt)
    ├── reflex/            # runnable Go solution per reflex drill
    ├── core5.md
    └── variants.md
```

## Daily flow (recommended)

```bash
go run .                       # daily drill: read + write (DSA track)
go run . -- --refresh          # today's level-matched write session
go run . -- --levels           # what each function has earned: L1 / L2 / L3
go run . -- --problems         # curated primary problem per function
go run . -- --run=core         # check core answers
go run . -- --run=reflex       # check reflex specialty answers
go run . -- --track=leetcode --run   # fetch today's 10 LeetCode problems
go run . -- --drill=core        # core only (Core Read 3 + Core 5)
go run . -- --drill=reflex      # today's specialty only
go run . -- --track=read        # reading only
go run . -- --track=write       # writing only
go run . -- --track=leetcode    # 10 LeetCode problems matching today's topic
go run . -- --help              # full option reference
```

Options are GNU-style: `-h`/`--help`, `-V`/`--version`, `--option=value` (space
still works), any unambiguous abbreviation, and an unknown option is an error.
The older `--run leetcode` and the `-r`/`-w`/`-l` run-side flags still work but
`--track` is the preferred selector.

## LeetCode practice (separate from reflex drills)

```bash
go run . -- --track=leetcode              # fetch + show today's 10 problems
go run . -- --track=leetcode --run        # same, explicit fetch
go run . -- --track=leetcode --refresh    # force re-fetch from LeetCode API
go run . -- --track=leetcode --catalog    # all weekday sets
go -C bin/study_leetcode run .            # direct CLI
```

Guide: [`drills/leetcode/README.md`](leetcode/README.md)

## Write drills

```bash
go run .                    # today's plan
go run -C drills/write/core5 .             # Core 5
go run -C drills/write/reflex/02_hashing_reflex .
go run . -- --run           # test + log
```

Guide: [`doc/write/START_HERE.md`](../doc/write/START_HERE.md)  
Math reference: [`doc/write/MATH_CONCEPTS.md`](../doc/write/MATH_CONCEPTS.md)

## Read drills

```bash
go -C bin/study_code run .
go run -C drills/read/core/00_core_read .
go run -C drills/read/weekday/03_name_the_pattern .
go -C bin/study_code run . -- --run
```

Guide: [`doc/read/START_HERE.md`](../doc/read/START_HERE.md)

## Solutions (after honest attempt)

**Reflex Go solutions:** [`solutions/reflex/`](solutions/reflex/) — runnable `main.go` per drill  
**Quick notes:** [`solutions/*.md`](solutions/) — triggers and bugs  
Read answer keys: [`read/answers/`](read/answers/)

## Track progress

```bash
go run . -- --levels    # every function, grouped by the level it has earned
go run . -- --weak      # the five worst, worst first
```

Progress is derived from `.drill_log.json`, which every `--run` updates.

## Where things live

| What | Where |
|------|--------|
| Documentation | `doc/` |
| Practice files | `drills/` (this folder) |
| CLI tooling | `bin/` |
| Problem catalog | `reference/problems/CATEGORIES.md` |
| Progress log | `.drill_log.json` |
