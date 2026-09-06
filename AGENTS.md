# AGENTS.md

This repo is a **pure Go (stdlib-only) DSA drill workspace**. There is no server,
database, or external service, and no third-party Go dependencies (no `go.sum`).
Go 1.22+ is the only required toolchain.

## Modules

Four separate Go modules with **no `go.work`**, so run `go` commands per module
(or use `go -C <dir> ...`):

| Module | Purpose |
|--------|---------|
| `.` (root) | unified daily runner — delegates to the others |
| `bin/study_play/` | write-reflex drills: levels, `--refresh` session, problem set |
| `bin/study_code/` | code-reading drills |
| `bin/study_leetcode/` | daily 10-question LeetCode set (hits the LeetCode API) |

Because they are separate modules, `go run ./bin/study_play` **fails** from the
repo root — the root module cannot resolve that package path. Use `go run . --
<flags>`, which delegates, or `go -C bin/study_play run .`.

## Everyday commands

```bash
go run .                    # today's plan (read + write)
go run . -- --refresh       # today's level-matched write session
go run . -- --levels        # what each function has earned: L1 / L2 / L3
go run . -- --run=reflex    # run today's specialty drill and log the result
```

Every entry point (root + the three `bin/` CLIs) shares one GNU-style option
front-end (`argutil.go`): `--help`/`-h`, `--version`/`-V`, `--opt=value`,
unambiguous long-option abbreviation, and an unknown `--option` is a usage
error (exit 2). Retired spellings still resolve: `--run-core5` → `--run=core`,
leetcode `--set` → `--show`, root `-r`/`-w`/`-l` → `--track`.

## Solving a drill (the core end-to-end flow)

Edit the `TODO: REFLEX` stubs in `drills/write/reflex/<NN>_*/main.go`, then:

```bash
go run -C drills/write/reflex/05_trees_stacks_reflex .
# or, to also update .drill_log.json:
go run . -- --run=reflex
```

Every assert prints `PASS:` / `FAIL:`; `--run` rolls those up into one result per
function and grades the function's level from the history.

## Lint / test

```bash
./bin/scripts/test-coverage.sh    # the real gate: the four source modules only
```

- `gofmt -l` and `go vet ./...` are clean in the root module and `bin/study_play`.
- **Do not judge the repo by `go test ./...` from the root.** That walks into
  `drills/**`, which are *practice files* — they are supposed to fail whenever a
  drill is blank or half-written. The coverage script excludes them.
- The coverage gate has a long-standing shortfall: it demands 80% and the tree
  sits in the mid-70s. That predates any recent work.

## Known issues

- `bin/study_leetcode` `TestRenderDailyGoMock` fails (`renderDailyGo missing
  "Search in sorted array"`) against the uncommitted change in `render_go.go`.
- `bin/study_leetcode/api.go` and `content.go` are not gofmt-clean.

## Conventions

- `levels.go` (cue table) and `primaries.go` (one LeetCode problem per function)
  must stay in sync: `levels_test.go` fails if a function has no cue, no primary,
  or a cue that names its own answer.
- `bin/study_leetcode/reflex_map.go` maps LeetCode slugs back to drill functions;
  keep it in step with `primaries.go` when adding problems.
- `argutil.go` is copied verbatim into all four modules (they are separate Go
  modules with no `go.work`). Edit one, then re-copy to the other three; each
  module supplies its own `gnuSpec` (program name, `canonical` option list,
  `aliases`). Add a new `--flag` to that module's `canonical` list or `gnuPre`
  rejects it as unrecognised. Bump `cliVersion` when the option surface changes.
