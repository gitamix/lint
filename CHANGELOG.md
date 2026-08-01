# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Implemented issue with its message and type,
indicating problems in the project that need to be linted
- Implemented config string value with linting issue level to report
- Lint configuration with branch-related setting
including its name and ticket the branch related with
- Loading and unmarshaling config from YAML file
- Integration tests to load and unmarhal the whole lint config
- Added linting current branch with its name and ticket
- Implemented fake git client for testing purposes.
- Added Docker test container with fake git tree to lint it in tests
- Added integration tests for linting current branch

### Changed

- Issue type parsing to accept string input or default value
