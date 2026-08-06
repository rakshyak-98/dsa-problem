# dsa-problem

Interview prep workspace for DSA drills, problem practice, and progress tracking.

## Prerequisites

| Tool | Version | Used for |
|------|---------|----------|
| [Go](https://go.dev/dl/) | 1.22+ | drill helpers (`bin/study_play`, `bin/study_code`, root `go run`) |
| Browser | any modern | study tracker + visualizer |
| Node.js (optional) | 18+ | `npx serve` for local tracker hosting |

## Start here

```bash
./setup.sh
```

Then open **[`doc/drills.md`](doc/drills.md)** — all practice files live under `drills/`.

## Quick commands (from repo root)

```bash
go run .                              # unified daily plan (DSA: read + write)
go run . -- --track backend           # backend interview prep only
go run . -- --track read              # DSA reading drills only
go run . -- --track write             # DSA writing drills only
go run . -- --list-tracks             # show all available tracks
go run . -- --run                     # check today's answers
go run . -- --core5                   # Core 5 reflex
open drills/tracker/study_tracker.html
```

## Repository layout

```
dsa-problem/
├── doc/                     # ★ DOCUMENTATION (all guides)
│   ├── drills.md            # drills overview
│   ├── write/               # write reflex study plans
│   ├── read/                # reading drill guides
│   └── backend/             # backend interview prep
├── drills/                  # ★ PRACTICE (by topic)
│   ├── write/               # reflex drills: core5, reflex, variants
│   ├── read/                # reading drills: core, weekday, answers
│   ├── backend/             # backend interview drills
│   ├── solutions/           # write drill solutions (after attempt)
│   └── tracker/             # browser study tracker
├── bin/                     # internal CLI tooling
│   ├── study_play/          # write-drill CLI + templates
│   ├── study_code/          # read-drill CLI
│   ├── study_backend/       # backend interview CLI
│   └── scripts/             # test coverage gate
├── reference/               # unrelated reference material
│   ├── problems/            # problem catalog + simulation
│   └── visualizer/          # algorithm visualizer
├── main.go                  # unified daily command (go run .)
└── setup.sh
```

## Testing

```bash
./bin/scripts/test-coverage.sh
```

## More docs

- [`doc/DSA_JARGON.md`](doc/DSA_JARGON.md) — plain-English glossary for DSA terms
- [`doc/write/START_HERE.md`](doc/write/START_HERE.md) — writing reflex flow
- [`doc/read/START_HERE.md`](doc/read/START_HERE.md) — reading drill flow
- [`doc/backend/START_HERE.md`](doc/backend/START_HERE.md) — backend interview cram
- [`reference/problems/CATEGORIES.md`](reference/problems/CATEGORIES.md) — problem index by topic
