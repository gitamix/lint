# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-09-02

### Added

- Implemented issue with its message and type,
indicating problems in the project that need to be linted
- Implemented config string value with linting issue level to report
- Lint configuration with branch-related setting
including its name and task the branch related with
- Lint configuration with commit-related setting including its message, scope and types
- Loading and unmarshaling config from YAML file
- Integration tests to load and unmarshal the whole lint config
- Added linting branch with its name and task
- Added linting commit messages with its type, scope, subject and body
- Filled project README with getting started guide,
full YAML configuration example and rules reference table

### Changed

- Issue type parsing to accept string input or default value
- Replaced table-driven tests templating in `STYLE.md`
with direct `t.Run()` per case
