# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.3.2]

### Changed

- Go directive bumped from 1.15 to 1.26 ([a249033])
- gofmt the entire codebase, convert length vars to consts ([7dfa1a3])
- Added Makefile, .drone.yml, .golangci.yml, AGENTS.md

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
[7dfa1a3]: https://github.com/LeakIX/l9format/commit/7dfa1a3
[a249033]: https://github.com/LeakIX/l9format/commit/a249033