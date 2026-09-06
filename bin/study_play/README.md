# study_play — write drill CLI

Internal tooling for the **write reflex** workflow. Practice files are in [`../../drills/write/`](../../drills/write/).

## Commands (from repo root)

```bash
go run .                          # today's drill plan: triggers → ask → drill → problems
go run . -- --refresh             # level-matched session (start here)
go run . -- --refresh --show      # same, with the recognition answers revealed
go run . -- --levels              # what each function has earned: L1 / L2 / L3
go run . -- --problems            # curated primary problem per function, by level
go run . -- --triggers            # full cross-topic pattern trigger table
go run . -- --run                 # test today's specialty + log progress
go run . -- --run=core            # test the Core 5 (was --run-core5)
go run . -- --weak                # weakest functions from the drill log
go run . -- --setup               # scaffold drills from blank templates
go run . -- --help                # full option reference (GNU-style: -h/-V, --opt=value, abbreviations)
```

## Understanding levels

A function is graded from `.drill_log.json`, not declared by hand, and the
session asks it in the form that matches its grade:

| Level | Earned when | You are asked to |
|-------|-------------|------------------|
| **L1 recall** | never passed, or failing often | restate the ask, then write the shape from memory |
| **L2 pattern** | ≥2 passes, under a third failing | read a cue with no function name in it and name the move |
| **L3 transfer** | ≥5 passes, near-clean, passed recently | apply the move to a variant that was never drilled |

A function is never graded above its intrinsic tier — typing `reverseInPlace`
fifty times does not turn it into a transfer question. Levels decay: an L2 function
untouched for 3 days, or an L3 for 7, comes back as due.

## What `--refresh` picks

Only the specialty block is chosen by weekday. Everything else comes from the log:

1. **Recognise** — 4 cues, one per pattern family, no code. The cross-topic mix is deliberate: recognising a move is only a skill when the topic is not given away.
2. **Rebuild** — the 3 weakest functions, whichever day owns them.
3. **New ground** — 2 never-attempted functions, highest-value family first (hashing → binary search → trees).
4. **Specialty** — today's weekday drill, with its triggers.
5. **Stretch** — L3 functions that are due, drilled on the twist rather than the original.

## Layout

| Path | Purpose |
|------|---------|
| `drills/write/` | your practice files |
| `drills/solutions/` | annotated solutions (peek after attempt) |
| `levels.go` | the cue table: ask / cue / move / twist per function |
| `primaries.go` | one curated LeetCode problem per function, plus the step up |
| `refresh.go` | session selection from the drill log |
| `_support/blanks/` | blank templates for `--setup` / `--reset` |
| `_support/asks/` | question literacy prompts |
| `doc/write/` | study plans and guides |
