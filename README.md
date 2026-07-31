# dsa-problem

Interview prep workspace for DSA drills, problem practice, and progress tracking.

## Prerequisites

| Tool | Version | Used for |
|------|---------|----------|
| [Go](https://go.dev/dl/) | 1.22+ | drill helpers (`study_play`, `study_code`, `daily`) |
| Browser | any modern | study tracker + visualizer |
| Node.js (optional) | 18+ | `npx serve` for local tracker hosting |

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

## Repository layout

```
dsa-problem/
├── drills/                  # ★ PRACTICE (front door)
│   ├── write/               # reflex drills: core5, reflex, variants
│   ├── read/                # reading drills: core, weekday, answers
│   ├── solutions/           # write drill solutions (after attempt)
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

## More docs

- [`study_play/docs/START_HERE.md`](study_play/docs/START_HERE.md) — writing reflex flow
- [`study_code/docs/START_HERE.md`](study_code/docs/START_HERE.md) — reading drill flow
- [`problems/CATEGORIES.md`](problems/CATEGORIES.md) — problem index by topic
