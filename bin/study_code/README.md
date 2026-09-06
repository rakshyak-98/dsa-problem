# study_code — reflex read drill CLI

Internal tooling for the **code reading** workflow. Practice files are in [`../../drills/read/weekday/`](../../drills/read/weekday/).

## Commands (from repo root)

```bash
go -C bin/study_code run .              # today's reflex reading plan
go -C bin/study_code run . -- --run     # check answers
go -C bin/study_code run . -- --catalog # full catalog
go -C bin/study_code run . -- --drill   # today's reflex read only ("reflex" is implied)
go -C bin/study_code run . -- --help    # option reference
```

Options follow the GNU conventions (`--help`, `--version`, `--opt=value`,
abbreviations). Guide: [`../../doc/read/START_HERE.md`](../../doc/read/START_HERE.md)
