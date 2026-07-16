# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.2] - 2026-07-15

### Added

- Go Guard reusable workflow was added to check for the presence of `go.mod`.
- Integration tests were added to the default `make test` flow.

### Fixed

- CI test workflow was corrected to run the proper coverage target.
- Set correct patterns in all files to be replaced with `make setup`
so that is why any file is not changed on the command called

### Changed

- Tests and Quality workflows now require Go Guard to run first and skip Go-specific jobs when `go.mod` is missing.
- Coverage generation and reporting were refined for unit and integration tests.
- golangci-lint configuration was migrated to v2 format, including updated linter, exclusions, formatter, and output sections.

## [1.0.1] - 2026-07-04

### Fixed

- Workflow 'Release' failed on 'Make scripts executable' cause of incorrect path to scripts used in the workflow

## [1.0.0] - 2026-07-04

### Added

- Initial release of the Go repository template.
- Project style guide and reusable section template.
- Code of Conduct, contribution guide, issue templates, and pull request template.
- GitHub Actions workflows for tests, linting, and releases.
- Release, setup, coverage, and repository bootstrap scripts.
- Git hooks and Make targets for local quality checks.

### Changed

- Standardized local validation around `make check`.
- Formalized the release process with semantic versioning,
automated tags, and changelog-based release notes.
- Improved contributor onboarding and repository workflow documentation.
- Extended the release workflow to run automatically after merging a pull request
with a `release: vX.Y.Z` title.
- Simplified `make pr` to support release pull requests only.
- `make pr release <version>` now opens the GitHub Pull Request page with prefilled title and body.
- Release PR summary is now populated from the matching `CHANGELOG.md` version entry.
- Release PR titles are normalized to the `release: vX.Y.Z` format.
