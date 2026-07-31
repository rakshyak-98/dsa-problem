# study_play — write drill CLI

This folder runs the **write reflex** workflow. Practice files are in [`../drills/write/`](../drills/write/).

## Commands

```bash
go run .              # today's drill plan
go run . -- --run     # test today's specialty + log progress
go run . -- --weak    # show weakest functions
go run . -- --setup   # scaffold drills from blank templates
```

## Internal layout

| Path | Purpose |
|------|---------|
| `_support/blanks/` | source templates for `--setup` |
| `_support/templates/` | reference implementations (tested) |
| `_support/solutions_reference/` | annotated solutions (peek after attempt) |
| `_support/asks/` | question literacy prompts |
| `docs/` | study plans and guides |

Start with [`docs/START_HERE.md`](docs/START_HERE.md).
