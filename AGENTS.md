# AGENTS.md

This repo is a **pure Go (stdlib-only) DSA drill workspace**. There is no server,
database, or external service, and no third-party Go dependencies (no `go.sum`).
Go 1.22+ is the only required toolchain.

## Modules

Three separate Go modules with **no `go.work`**, so run `go` commands per module
(or use `go -C <dir> ...`):

| Module | Purpose |
|--------|---------|
| `.` (root) | unified daily runner — delegates to the others |
| `bin/study_play/` | write-reflex drills: drill log, problem set |
| `bin/study_leetcode/` | daily 10-question LeetCode set (hits the LeetCode API) |

Because they are separate modules, `go run ./bin/study_play` **fails** from the
repo root — the root module cannot resolve that package path. Use `go run . --
<flags>`, which delegates, or `go -C bin/study_play run .`.

## Everyday commands

```bash
go run .                    # today's plan (weekday reflex write)
go run . -- --problems      # curated LeetCode problem per function
go run . -- --run=reflex    # run today's specialty drill and log the result
go run . -- --reset         # restore today's reflex drill to its blank template
```

Every entry point (root + the three `bin/` CLIs) shares one GNU-style option
front-end (`argutil.go`): `--help`/`-h`, `--version`/`-V`, `--opt=value`,
unambiguous long-option abbreviation, and an unknown `--option` is a usage
error (exit 2). Retired spellings still resolve:
leetcode `--set` → `--show`, root `-w`/`-l` → `--track`.

## Solving a drill (the core end-to-end flow)

Edit the `TODO: REFLEX` stubs in `drills/write/reflex/<NN>_*/main.go`, then:

```bash
go run -C drills/write/reflex/05_trees_stacks_reflex .
# or, to also update .drill_log.json:
go run . -- --run=reflex
```

Every assert prints `PASS:` / `FAIL:`; `--run` rolls those up into one result per
function in `.drill_log.json`.

## Lint / test

```bash
./bin/scripts/test-coverage.sh    # the real gate: the three source modules only
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

- `primaries.go` (one LeetCode problem per function) is also the function
  registry the drill log uses to map asserts to functions; `primaries_test.go`
  fails if a catalog function has no primary.
- Small-skill questions: `skillBank` in `bin/study_play/asks.go`. Compute answers; do not
  eyeball them. `--missed`/`--got` store misses under `skills` in `.drill_log.json`.
- `bin/study_leetcode/reflex_map.go` maps LeetCode slugs back to drill functions;
  keep it in step with `primaries.go` when adding problems.
- `argutil.go` is copied verbatim into all three modules (they are separate Go
  modules with no `go.work`). Edit one, then re-copy to the other three; each
  module supplies its own `gnuSpec` (program name, `canonical` option list,
  `aliases`). Add a new `--flag` to that module's `canonical` list or `gnuPre`
  rejects it as unrecognised. Bump `cliVersion` when the option surface changes.
