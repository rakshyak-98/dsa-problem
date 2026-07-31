# dsa-problem

Interview prep workspace: problem catalog, reflex drills, code-reading drills, study tracker, and algorithm visualizer.

## Prerequisites

| Tool | Version | Used for |
|------|---------|----------|
| [Go](https://go.dev/dl/) | 1.22+ | `study_play` and `study_code` drill helpers |
| Browser | any modern | study tracker + visualizer |
| Node.js (optional) | 18+ | `npx serve` for local tracker hosting |

## Quick setup

```bash
git clone https://github.com/rakshyak-98/dsa-problem.git
cd dsa-problem
./setup.sh
```

`setup.sh` creates `study_play/drills/` from the embedded blank templates and verifies both Go helpers run.

Manual setup (equivalent):

```bash
cd study_play && go run . -- --setup
```

## Repository map

```
dsa-problem/
├── setup.sh              # first-time setup
├── CATEGORIES.md         # problem catalog by topic
├── study_play/           # write reflex drills + study tracker
│   ├── START_HERE.md
│   ├── study_tracker.html
│   └── drills/           # created by setup (from blanks/)
├── study_code/           # read drills (answers in drills/)
│   └── START_HERE.md
├── simulation/           # JS simulation problems
└── visualizer/           # algorithm visualizer (static HTML/JS)
```

## Daily workflow

**Write reflexes** (`study_play`):

```bash
cd study_play
go run .              # today's drill plan
go run . -- --run     # run today's specialty tests
go run . -- --reset   # restore today's drill to TODO stubs
```

**Read code fluently** (`study_code`):

```bash
cd study_code
go run .              # today's reading drill
go run . -- --run     # check filled answers
```

**Track progress**: open `study_play/study_tracker.html` in a browser.

**Visualize algorithms**: open `visualizer/index.html`.

## More docs

- [`study_play/README.md`](study_play/README.md) — tracker, 12-week plan, reflex drills
- [`study_play/START_HERE.md`](study_play/START_HERE.md) — day-one writing flow
- [`study_code/README.md`](study_code/README.md) — code-reading drills
- [`study_code/START_HERE.md`](study_code/START_HERE.md) — day-one reading flow
