# study_play — write drill CLI

Internal tooling for the **write reflex** workflow. Practice files are in [`../../drills/write/`](../../drills/write/).

## Commands (from repo root)

```bash
go run ./bin/study_play              # today's drill plan
go run ./bin/study_play -- --run     # test today's specialty + log progress
go run ./bin/study_play -- --weak    # show weakest functions
go run ./bin/study_play -- --setup   # scaffold drills from blank templates
```

## Layout

| Path | Purpose |
|------|---------|
| `drills/write/` | your practice files |
| `drills/solutions/` | annotated solutions (peek after attempt) |
| `_support/blanks/` | blank templates for `--setup` / `--reset` |
| `_support/asks/` | question literacy prompts |
| `doc/write/` | study plans and guides |
