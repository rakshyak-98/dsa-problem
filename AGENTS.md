# AGENTS.md

## Cursor Cloud specific instructions

This repo is a **pure Go (stdlib-only) DSA interview-prep workspace** plus two static-HTML tools. There is no server, database, or external service, and there are no third-party Go dependencies (no `go.sum`). Go 1.22+ is the only required toolchain (already installed in the cloud image).

### Modules and how to run them

There are four separate Go modules with **no `go.work`**, so run `go` commands from each module directory (or use `go -C <module> ...`). Standard run commands are documented in `README.md`; briefly:

- `study_play/` — write-reflex drills + tracker. `go -C study_play run .` (plan), `... run . -- --run`, `... run . -- --catalog`.
- `study_code/` — code-reading drills. `go -C study_code run .`.
- `daily/` — unified daily runner. `go -C daily run .`.
- `practice/` — scratch problem module (not part of the test suite).

**Solving a drill (the core end-to-end flow):** edit the TODO stubs in `study_play/drills/<NN>_*/main.go`, then run that drill directly, e.g. `go -C study_play run ./drills/05_trees_stacks_reflex`. All asserts print `PASS:` and it ends with `All ... reflex drills passed.` when correct.

**Static GUIs:** open `study_play/study_tracker.html` and `visualizer/index.html` directly in a browser. The visualizer loads CodeMirror from the Cloudflare CDN, so it needs internet access.

### Lint / test

- Tests: `study_code` and `daily` pass with plain `go test .`. The coverage helper is `tools/scripts/test-coverage.sh`.
- **Known caveat — `study_play` `go test` fails on `go vet`:** `daily_drill.go` has an unreachable duplicate `--setup` block whose `fmt.Println("...\n")` trips vet ("redundant newline"). Since `go test` runs vet by default, run study_play tests with `go -C study_play test . ./_support/templates -vet=off`.
- **Known caveat — 4 study_play tests fail regardless of vet:** `TestBlankContent`, `TestWriteDrillFromBlank`, `TestSetupAllDrills`, `TestResetTodayDrill` fail because `reset.go` embeds `_support/blanks/*` but reads `blanks/<file>.go`. The `_support/templates` package tests and all other study_play tests pass.

### Setup / drill-generation caveats (pre-existing, non-blocking)

- The generated drills are **already committed** under `study_play/drills/`, so no setup is needed to use the product.
- `study_play --setup` and `--reset` are broken by the same `reset.go` embed-path bug above (`no blank template ...`). Do not rely on them to (re)generate drills.
- `setup.sh` contains an unresolved git merge conflict near the end and will not complete cleanly; the useful setup steps are already covered by the committed drills.
- `practice/go.mod` declares `go 1.26.4` (higher than the installed toolchain), so any `go` command in `practice/` triggers a `go1.26.4` toolchain download. It is a scratch module and is excluded from `tools/scripts/test-coverage.sh`.
