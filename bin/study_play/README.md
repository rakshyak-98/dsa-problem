# study_play — write drill CLI

Internal tooling for the **write reflex** workflow. Practice files are in [`../../drills/write/`](../../drills/write/).

## Commands (from repo root)

```bash
go run .                          # today's drill plan: ask → drill → problems
go run . -- --show                # today's plan with the small-skill answers revealed
go run . -- --problems            # curated primary problem per function
go run . -- --run                 # grade today's specialty + log progress
go run . -- --fn=A,B              # blank just these functions; --run --fn=A,B grades them
go run . -- --hint=A              # pattern + technique for a function
go run . -- --recognize           # name the pattern for 5 unlabeled problems
go run . -- --weak                # weakest functions from the drill log
go run . -- --setup               # scaffold drills from blank templates
go run . -- --help                # full option reference (GNU-style: -h/-V, --opt=value, abbreviations)
```

## Layout

| Path | Purpose |
|------|---------|
| `drills/write/` | your practice files |
| `drills/solutions/` | annotated solutions (peek after attempt) |
| `primaries.go` | one curated LeetCode problem per function, plus a harder follow-up |
| `patterns.go` | pattern ids, function → pattern, review intervals, LeetCode cards (copied verbatim into `bin/study_leetcode`) |
| `grade.go` | per-function grading (time + memory capped) and first-try scoring |
| `schedule.go` | spaced-repetition review queue shown in the daily plan |
| `single.go` | `--fn`: swap single functions back to their blank stubs |
| `hints.go` | technique per function (`--hint`, and after a first-try miss) |
| `recognize.go` | `--recognize` question bank and round |
| `_support/blanks/` | blank templates for `--setup` / `--reset` |
| `_support/asks/` | question literacy prompts |
| `doc/write/` | study plans and guides |
