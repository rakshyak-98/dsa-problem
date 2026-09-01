# dsa-problem

A DSA drill workspace: write patterns from memory, read code for bugs, solve the
matching problem, and let the log decide what tomorrow looks like.

## Prerequisites

| Tool | Version | Used for |
|------|---------|----------|
| [Go](https://go.dev/dl/) | 1.22+ | every drill helper and the root `go run .` |

No other dependencies. No server, no database, no third-party Go modules.

## Start here

```bash
./setup.sh
```

Then open **[`doc/drills.md`](doc/drills.md)** — all practice files live under `drills/`.

## Quick commands (from repo root)

```bash
go run .                              # today's plan: read + write
go run . -- --refresh                 # level-matched write session (start here)
go run . -- --refresh --show          # same, with the recognition answers revealed
go run . -- --levels                  # what each function has earned: L1 / L2 / L3
go run . -- --problems                # curated primary problem per function, by level
go run . -- --triggers                # full cross-topic pattern trigger table
go run . -- --weak                    # weakest functions from the drill log
go run . -- --run                     # check today's reflex answers
go run . -- --core5                   # Core 5 reflex
go run . -- --track read              # reflex reading drills only
go run . -- --track write             # writing drills only
go run . -- --track leetcode          # daily 10 LeetCode problems (weekday topic)
go run . -- --run leetcode            # fetch today's 10 LeetCode problems
go run . -- --list-tracks             # show all available tracks
```

## Repository layout

```
dsa-problem/
├── doc/                     # ★ DOCUMENTATION (all guides)
│   ├── drills.md            # drills overview
│   ├── write/               # write reflex study plans
│   └── read/                # reading drill guides
├── drills/                  # ★ PRACTICE (by topic)
│   ├── write/               # reflex drills: core5, reflex, variants
│   ├── read/                # reading drills: core, weekday, answers
│   ├── leetcode/            # daily 10-question LeetCode practice sets
│   └── solutions/           # write drill solutions (after attempt)
├── bin/                     # internal CLI tooling
│   ├── study_play/          # write-drill CLI: levels, refresh, problem set
│   ├── study_code/          # read-drill CLI
│   ├── study_leetcode/      # daily LeetCode practice set CLI
│   └── scripts/             # test coverage gate
├── reference/problems/      # problem catalog by topic + solved simulations
├── main.go                  # unified daily command (go run .)
└── setup.sh
```

## Testing

```bash
./bin/scripts/test-coverage.sh
```

## More docs

- [`bin/study_play/README.md`](bin/study_play/README.md) — understanding levels and how `--refresh` picks a session
- [`doc/write/DAILY_30MIN_DRILL.md`](doc/write/DAILY_30MIN_DRILL.md) — the daily session, start to finish
- [`doc/write/START_HERE.md`](doc/write/START_HERE.md) — writing reflex flow
- [`doc/read/START_HERE.md`](doc/read/START_HERE.md) — reading drill flow
- [`doc/DSA_JARGON.md`](doc/DSA_JARGON.md) — plain-English glossary for DSA terms
- [`reference/problems/CATEGORIES.md`](reference/problems/CATEGORIES.md) — problem index by topic
