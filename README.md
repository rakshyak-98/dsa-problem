# dsa-problem

Interview prep workspace: problem catalog, reflex drills, code-reading drills, study tracker, and algorithm visualizer.

## Prerequisites

| Tool | Version | Used for |
|------|---------|----------|
| [Go](https://go.dev/dl/) | 1.22+ | drill helpers, unified daily command |
| Browser | any modern | study tracker + visualizer |
| Node.js (optional) | 18+ | `npx serve` for local tracker hosting |

## Quick setup

```bash
git clone https://github.com/rakshyak-98/dsa-problem.git
cd dsa-problem
./setup.sh
```

## Unified daily command (recommended)

```bash
cd daily && go run .           # read + write plan for today
cd daily && go run . -- --run  # run today's reading + writing tests
cd daily && go run . -- --micro  # core tiers only
```

## Daily workflow

**1. Question literacy** — printed in `study_play` daily output (`asks/`)

**2. Core 5** (every day, under 8 min):

```bash
cd study_play && go run ./core5
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
cd study_play && go run ./variants
```

**6. Primary problem** — mapped in daily output (e.g. `hashing/easy/two_sum.js`)

**7. Track progress**: `study_play/study_tracker.html`

**8. Visualize**: `visualizer/index.html` (link printed in daily output)

## Repository map

```
dsa-problem/
├── setup.sh
├── daily/                    # unified read + write command
├── CATEGORIES.md
├── study_play/
│   ├── core5/                # executable Core 5 drill
│   ├── variants/             # pattern variant drills
│   ├── asks/                 # question literacy
│   ├── solutions_reference/  # peek after honest attempt
│   ├── drills/               # weekday reflex drills (+ bonus heap/backtrack)
│   └── study_tracker.html
├── study_code/
└── visualizer/
```

## Features

| Feature | Command / location |
|---------|-------------------|
| Executable Core 5 | `study_play/core5/` |
| Drill → problem mapping | printed by `go run .` in study_play |
| Question literacy asks | `study_play/asks/` |
| Spaced repetition log | `study_play/.drill_log.json` via `--run` |
| Weak function review | `go run . -- --weak` |
| Annotated solutions | `study_play/solutions_reference/` |
| Bonus drills | `08_heap_reflex`, `09_backtrack_reflex` |
| Visualizer links | printed in daily output |

## More docs

- [`study_play/START_HERE.md`](study_play/START_HERE.md) — day-one writing flow
- [`study_play/DAILY_30MIN_DRILL.md`](study_play/DAILY_30MIN_DRILL.md) — full reflex ritual
- [`study_code/START_HERE.md`](study_code/START_HERE.md) — day-one reading flow
