# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Dependabot configuration for Go module and GitHub Actions updates ([7ed9cc3])
- zizmor linting of GitHub Actions workflows in CI ([7ed9cc3])
- Makefile, golangci-lint config, changelog script, AGENTS.md ([c42abd1])

### Changed

- Replaced the internal CI configuration with a GitHub Actions workflow ([7ed9cc3])
- Pinned all GitHub Actions to commit hashes and scoped workflow
  permissions ([7ed9cc3])
- Go directive bumped from 1.15 to 1.26, gofmt the codebase, and fix
  staticcheck findings ([a2ba36c])

## [1.3.1]

### Changed

- Removed old gitlab dependency ([6707244])

## [1.3.0]

### Added

- Optional tags filtering to plugin interface ([2e87fb5])

## [1.2.0]

### Added

- TLS session cache for tests ([5df4ca9])

## [1.1.0]

### Added

- Optional `Init()` function to plugin interface ([ec39523])

## [1.0.0]

### Added

- Initial stable release. L9 event schema, plugin interface,
  fingerprint helpers ([v1.0.0])

<!-- Commit links -->
[v1.0.0]: https://github.com/LeakIX/l9format/releases/tag/v1.0.0
[ec39523]: https://github.com/LeakIX/l9format/commit/ec39523
[5df4ca9]: https://github.com/LeakIX/l9format/commit/5df4ca9
[2e87fb5]: https://github.com/LeakIX/l9format/commit/2e87fb5
[6707244]: https://github.com/LeakIX/l9format/commit/6707244
[a2ba36c]: https://github.com/LeakIX/l9format/commit/a2ba36c
[c42abd1]: https://github.com/LeakIX/l9format/commit/c42abd1
[7ed9cc3]: https://github.com/LeakIX/l9format/commit/7ed9cc3