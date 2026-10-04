# study_play — write drill CLI

Internal tooling for the **write reflex** workflow. Practice files are in [`../../drills/write/`](../../drills/write/).

## Commands (from repo root)

```bash
go run .                          # today's drill plan: ask → drill → problems
go run . -- --show                # today's plan with the small-skill answers revealed
go run . -- --problems            # curated primary problem per function
go run . -- --run                 # test today's specialty + log progress
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
| `_support/blanks/` | blank templates for `--setup` / `--reset` |
| `_support/asks/` | question literacy prompts |
| `doc/write/` | study plans and guides |
