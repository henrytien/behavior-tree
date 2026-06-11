# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Optional library logger (`Logf` / `SetLogger` / `SetLogOutput`), silent by
  default so the library never writes to stdout unasked.
- `loader.CreateBehaviorTreeFromConfigSafe`: an error-returning loader that
  recovers from malformed-config panics instead of crashing the process.
- Core unit tests for composites (Sequence, Priority), the Inverter decorator,
  and `RUNNING` tick propagation.
- Bilingual (中文 / English) documentation site built with Docusaurus, plus a
  "Design Notes" page covering node concepts and implementation-specific tips.
- CI workflows: cross-platform Go build/vet/test, and automatic deployment of
  the documentation site to GitHub Pages.
- `LICENSE`, `CONTRIBUTING.md`, and this `CHANGELOG.md`.

### Changed
- Renamed the import alias `b3` to `bt` across the codebase to match the
  behavior-tree naming (pure alias rename, no behavior change).
- Standardized naming: external brand / module path is `behavior-tree`; the Go
  package declaration is `behaviortree`.
- Upgraded the Go directive to 1.23.
- Rewrote the README with installation, a quick-start example, and a built-in
  node overview; removed outdated local-editor setup steps.

### Fixed
- Corrected copy-paste panic messages in the `Limiter`, `MaxTime`, `Repeater`,
  `RepeatUntilFailure`, and `RepeatUntilSuccess` decorators, which all named
  the wrong decorator type.
- Removed dead `unreachable code` after `panic` in the config package and
  replaced the deprecated `io/ioutil` with `os`.
- Replaced dead external links (behaviac.com, the Youdao note) with stable
  references.

## [0.2.0]

Earlier history, inherited from `magicsea/behavior3go`:

### Added
- `SubTree` node support (requires the editor to export the node `category`).
- `RandWait` action: random wait supporting `min_ms`/`max_ms`, compatible with
  `timemini`/`timemax` and a fixed `milliseconds`.
- `Probability` decorator: supports `probability`/`rate`/`percent` and
  `skip_status`.

[Unreleased]: https://github.com/henrytien/behavior-tree/compare/master...HEAD
