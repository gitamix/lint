# lint

[![Go Reference](https://pkg.go.dev/badge/github.com/gitamix/lint.svg)](https://pkg.go.dev/github.com/gitamix/lint)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gitamix/lint)
[![Release](https://img.shields.io/github/v/release/gitamix/lint?style=flat)](https://github.com/gitamix/lint/releases)

A configurable Go library that lints git branch names and commit messages
against rules defined in YAML.
Every rule carries its own severity level — `critical`, `warning`, or `info` —
so you control how each violation is reported.

> [!NOTE]
> The project is not a CLI tool —
> it is a Go library you can embed into any tool,
> including your own CLI or another Go project.

## Features

- **Branch linting** — validate branch names against a regex pattern
  and enforce the task/ticket ID format, e.g. `feature/TASK-123`.
- **Commit linting** — validate the commit subject: allowed types, scope
  pattern, description length and task ID; and the commit body: length
  and mandatory body per commit type.
- **YAML-driven configuration** — declare rules in a single `lint.yml`;
  every key is optional, and rules are strictly opt-in:
  an empty config reports no issues.
- **Per-rule severity** — every rule is paired with an issue level:
  `critical`, `warning`, or `info`.
- **Programmatic configuration** — build the same configuration in code
  with functional options, no YAML required.
- **Pure library** — the linter never runs git commands itself:
  you pass in values built with
  [github.com/gitamix/types](https://github.com/gitamix/types),
  so it fits any workflow.

## Installation

Run the command in terminal:

```sh
go get github.com/gitamix/lint
go get github.com/gitamix/types
```

The `github.com/gitamix/types` module provides the git domain types
(`Branch`, `Name`, `Commit`, `Message`, …)
consumed by the linters — you need it to construct the values you want to lint.

## Getting started

Describe your rules in a `gitamix.yml` file.
Every key is optional — only the rules you configure are enforced.

```yaml
# Git branch rules.
branch:
  # task rules to enforce a specific format
  # for task identifiers in branch names.
  task:
    issue:
      # Level for task rules: critical, warning, info
      level: "info"
    # Regular expression pattern for the task identifier.
    pattern: (TASK|PROJ|BUG)-[0-9]+
  name:
    issue:
      # Level for branch name rules: critical, warning, info
      level: "warning"
    # Regular expression pattern for branch names.
    # This pattern allows only feature, bugfix and hotfix
    # branches with an uppercase task ID.
    pattern: ^(feature|bugfix|hotfix)/[A-Z]+-\d+

# Git commit rules.
commit:
  # Commit message rules, grouping subject and body configuration.
  message:
    # Subject rules (the first line of the commit message).
    subject:
      # Description rules for the subject text,
      # excluding the task identifier, type, and scope.
      description:
        # Length interval for the subject description text.
        length:
          issue:
            # Level for description length rules: critical, warning, info
            level: warning
          # Minimum number of characters allowed in the subject description.
          min: 10
          # Maximum number of characters allowed in the subject description.
          max: 72
      # Task integration rules for the commit subject.
      task:
        # Identifier rules for the task reference in the commit subject.
        id:
          issue:
            # Level for task id rules: critical, warning, info
            level: critical
          # Regular expression pattern for the task identifier.
          pattern: (TASK|PROJ|BUG)-[0-9]+
    # Body rules (the text after the subject, separated by a blank line).
    body:
      # Mandate rules: commit types for which the body is mandatory.
      mandate:
        # Commit types for which the body is mandatory.
        types:
          issue:
            # Level for body mandate rules: critical, warning, info
            level: critical
          # List of commit types that require a body.
          list:
            - fix
            - chore
            - refactor
            - perf
      # Length interval for the body text.
      length:
        issue:
          # Level for body length rules: critical, warning, info
          level: info
        # Minimum number of characters allowed in the body.
        min: 10
        # Maximum number of characters allowed in the body.
        max: 255
  # Scope rules for the commit subject
  # (the part in parentheses after the type).
  scope:
    issue:
      # Level for scope rules: critical, warning, info
      level: warning
    # Regular expression pattern for the commit scope.
    pattern: ^[A-Za-z _-]+$
  # Allowed commit types (e.g. feat, fix, chore).
  types:
    issue:
      # Level for commit type rules: critical, warning, info
      level: critical
    # List of allowed commit types.
    list:
      - feat
      - fix
      - chore
      - refactor
      - perf
      - test
      - docs
```

Load the config and lint branches and commits:

```go
package main

import (
    "fmt"
    "log"

    typesbranch "github.com/gitamix/types/branch"
    typescommit "github.com/gitamix/types/commit"

    "github.com/gitamix/lint/branch"
    "github.com/gitamix/lint/commit"
    "github.com/gitamix/lint/config"
)

func main() {
    cfg, err := config.Load("path/to/gitamix.yml")
    if err != nil {
        log.Fatal(err)
    }
    brs := branch.NewBranch(
        typesbranch.NewBranch(typesbranch.NewName("release/TASK-123")),
        cfg.Branch(),
    )
    for _, iss := range brs.Issues() {
        fmt.Println(iss.Type(), iss.Message()) 
    }
    cmts := commit.NewCommits(
        []typescommit.Commit{
            typescommit.NewCommit(
                typescommit.NewHash("abc1234567"),
                typescommit.ParseMessage([]byte("[TASK-123] fix: tiny fix\n")),
            ),
        },
        cfg.Commit(),
    )
    for _, iss := range cmts.Issues() {
        fmt.Println(iss.Type(), iss.Message())
    }
}
```

Output:

```text
warning branch name 'release/TASK-123' doesn't match the required pattern '^(feature|bugfix|hotfix)/[A-Z]+-\d+'
warning abc1234: scope not found in subject by expression '^[A-Za-z _-]+$'
warning abc1234: subject description length is not in range [10-72]
critical abc1234: body is required for type 'fix'
```

### Rules reference

| YAML path | Rule | Default level |
|---|---|---|
| `branch.name` | Branch name matches `pattern` | — set `level` explicitly |
| `branch.task` | Branch name contains a task ID matching `pattern` | `critical` |
| `commit.types` | Commit type is one of `list` | `critical` |
| `commit.scope` | Commit scope matches `pattern` | `warning` |
| `commit.message.subject.description.length` | Subject description length is within `[min, max]` (inclusive) | `warning` |
| `commit.message.subject.task.id` | Subject contains a task ID matching `pattern` | `critical` |
| `commit.message.body.length` | Body length is within `[min, max]` (inclusive) | — set `level` explicitly |
| `commit.message.body.mandate.types` | Body is required for the commit types in `list` | always `critical` |

When a rule is configured but `level` is omitted,
the default level from the table applies.
Rules without a default report with an unspecified level,
so set `level` explicitly for them.

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/gitamix/lint)

## Requirements

- [Go](https://go.dev/) 1.25.7 or later.
- [github.com/gitamix/types](https://github.com/gitamix/types) — git domain types
  used to construct the values being linted.
- [gopkg.in/yaml.v3](https://github.com/go-yaml/yaml) — YAML parsing
  for `config.Load` and `config.Unmarshal` (pulled in automatically).

The library does not run git commands itself.
Obtain branch names and commit messages in your own workflow
(for example, with `git branch` and `git log`),
wrap them into [github.com/gitamix/types](https://github.com/gitamix/types) values,
and pass them to the linters.

## Contributing

Want to contribute?
Read [CONTRIBUTING.md](CONTRIBUTING.md)
for the full workflow, repository requirements, and Pull Request process.

Please open an issue to discuss large or breaking changes before implementing.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Author / Contact

Maintained by [Ilya Sitnikov](https://github.com/gitamix)
