# AI Agent Instructions — gitamix/lint

## Project

Go module `github.com/gitamix/lint` (Go 1.25.7, no vendor). A linting toolkit for git branches — loads config from YAML, inspects branches, detects issues.

## Quick commands

```
make check        # test + coverage + lint  (default)
make test         # unit-test + integration-test
make unit-test    # Go test ./... (excludes internal/test/integration)
make integration-test  # Go test ./internal/test/integration/... -tags=integration
make coverage     # checks threshold (default COVERAGE_THRESHOLD=90)
make lint         # golangci-lint run ./...
make setup        # installs git hooks + runs scripts/itsme.sh
make pr release <version>  # opens GitHub PR page with release template
```

## Architecture

```
cmd/          — entry-points for tooling (generator, 01)
config/       — YAML config types, loading, unmarshaling (branch, task, value subpackages)
branch/       — branch name parsing and detection
issue/        — linting issue types (message, type, severity level)
errs/         — custom error definitions
internal/test/ — test helpers: testcontainers, fakes, fixtures, testdata
scripts/      — automation: itsme.sh, coverage checks, PR creation
make/         — Makefile include fragments (test, lint, coverage, pr, setup)
.github/      — CI workflows (go_guard, tests, quality, release)
.githooks/    — pre-commit (unit-test), pre-push (make check)
docker/       — Dockerfile for test containers
```

## Style rules (enforced by .golangci.yml + STYLE.md)

- **All functions/methods must be exported** (unexported = forbidden)
- **All struct fields must be unexported**; use `New*` constructors
- **No nil** except `error` — model optionality with custom types or bool flags
- **No setters/getter prefixes** — accessors are `Name()` not `GetName()`
- **No logic in constructors** — pure field assignment + functional options only
- **Vertical formatting** — multi-line function calls / options must use one-item-per-line
- **No newlines inside function bodies** — forces small functions
- **No inline comments** inside functions — doc comments only
- **Tests in `*_test` package** — must import the implementation via constructor
- **Direct `t.Run()` per case** — each case is its own `t.Run(name, ...)` with inline setup, act, assert; `t.Parallel()` on the outer test and on every subtest. No `tests := []struct{...}` slices, no `args`/`want` structs
- **Testify** — `assert.Equal`, `assert.ErrorIs`; never `t.Error()`

### golangci-lint (v2)

Enabled linters: govet, staticcheck, unused, errcheck, revive, misspell, unparam, gocyclo (min 15), paralleltest, godoclint, godot, godox, dogsled, bodyclose, testpackage.
Exclusions in `*_test.go`: `errcheck`, `unparam`.
Formatters: gofmt + goimports.

## Testing

- **Unit tests**: `go test -race -count=1 ./...` excluding `internal/test/integration`
- **Integration tests**: tagged with `integration`, cover real packages via `-coverpkg`
- **Coverage**: unit profile → `tmp/coverage_unit.out`, integration → `tmp/coverage_integration.out`, combined → `tmp/coverage_total.out`
- Threshold: 90% (see `.github/ci.env`)
- Integration tests may require Docker (testcontainers — see `docker-compose.test.yml`)

## Git / CI / Release

- Branches: feature branches from `master`
- Commits: Conventional Commits (enforced by pre-push)
- PRs: merge with squash
- Changelog: Keep a Changelog (v2.0.0), update for user-visible changes
- Release: never tag manually; use `make pr release v1.0.0` (reads CHANGELOG block, opens PR) or `.github/workflows/release.yml` workflow dispatch
- CI on PR: go_guard → tests (unit + integration + coverage check) + quality (golangci-lint v2.8.0)

## Gotchas

- `make setup` registers git hooks (`.githooks/`) — run it after cloning
- Pre-commit runs `make unit-test`; pre-push runs `make check`
- Integration tests use testcontainers — Docker must be running
- `scripts/itsme.sh` replaces template placeholders (run after `make setup`)
- COVERAGE_THRESHOLD is read from `.github/ci.env` (default 90)
- Go version: `go.mod` specifies 1.25.7 — use `go-version-file: go.mod` in CI
