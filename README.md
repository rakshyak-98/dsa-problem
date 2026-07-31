# dsa-problem

Interview prep workspace for DSA drills, problem practice, and progress tracking.

## Start here

```bash
./setup.sh
```

Then open **[`drills/README.md`](drills/README.md)** — all practice files live under `drills/`.

## Quick commands

```bash
cd daily && go run .                    # unified daily plan
go run -C study_play/practice/write/core5 .  # Core 5 reflex
cd study_play && go run . -- --run        # test today's write drill
cd study_code && go run . -- --run      # check reading answers
open drills/tracker/study_tracker.html  # progress tracker
```

## Daily workflow

**1. Question literacy** — printed in `study_play` daily output (`_support/asks/`)

**2. Core 5** (every day, under 8 min):

```bash
go run -C study_play/practice/write/core5 .
cd study_play && go run . -- --run-core5   # test + log progress
```

**3. Write reflexes** (`study_play`):

```bash
cd study_play
go run .              # today's drill plan + problem map + visualizer link
go run . -- --run     # run specialty tests + spaced-repetition log
go run . -- --weak    # show your weakest functions
go run . -- --reset   # restore today's drill to TODO stubs
```

**4. Read code** (`study_code`):

```bash
cd study_code && go run .
cd study_code && go run . -- --run
```

**5. Variants** (medium prep):

```bash
go run -C study_play/practice/write/variants .
```

**6. Primary problem** — mapped in daily output (e.g. `hashing/easy/two_sum.js`)

**7. Track progress**: `drills/tracker/study_tracker.html`

**8. Visualize**: `visualizer/index.html` (link printed in daily output)

## Repository layout

```
dsa-problem/
├── drills/                  # ★ PRACTICE (front door)
│   ├── write/               # reflex drills: core5, reflex, variants
│   ├── read/                # reading drills: core, weekday, answers
│   └── tracker/             # browser study tracker
├── daily/                   # unified read + write command
├── study_play/              # write-drill CLI + internal support
│   ├── _support/            # blanks, templates, solutions (hidden)
│   └── docs/                # study plans and guides
├── study_code/              # read-drill CLI
│   └── docs/
├── problems/                # problem catalog + simulation
├── visualizer/              # algorithm visualizer
└── tools/scripts/           # test coverage gate
```

## Testing

```bash
./tools/scripts/test-coverage.sh
```

Runs Go tests across `study_play`, `study_play/_support/templates`, `daily`, and `study_code` with an **80% average coverage** gate. Reference implementations in `_support/templates/` are fully tested; student drill stubs in `drills/` are excluded (they start as `panic` stubs).

## More docs

- [`study_play/docs/START_HERE.md`](study_play/docs/START_HERE.md) — writing reflex flow
- [`study_code/docs/START_HERE.md`](study_code/docs/START_HERE.md) — reading drill flow
- [`problems/CATEGORIES.md`](problems/CATEGORIES.md) — problem index by topic
