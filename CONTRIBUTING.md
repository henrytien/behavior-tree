# Contributing

Thanks for your interest in improving **behavior-tree**! This project is a Go
implementation of the [behavior3](https://github.com/behavior3) behavior tree,
maintained independently at `github.com/henrytien/behavior-tree`.

## Getting started

Requires Go 1.23 or later.

```bash
git clone https://github.com/henrytien/behavior-tree
cd behavior-tree
go build ./...
go test ./...
```

## Before opening a pull request

Please make sure the following all pass locally:

```bash
gofmt -l .        # should print nothing
go vet ./...      # should report nothing
go test ./...     # should pass
```

- **Format**: code must be `gofmt`-clean.
- **Tests**: add tests for new behavior. Core behavior tests live in
  `loader/` (the loader package imports every node package, so trees can be
  assembled from JSON configs in tests without import cycles).
- **No stray output**: the library must stay silent by default. Route any
  diagnostic output through `Logf` rather than `fmt.Println`.
- **Backward compatibility**: avoid breaking the existing public API. Prefer
  additive changes (for example, the error-returning `...Safe` loader sits
  alongside the panic-based one).

## Commit messages

Use short, imperative, English commit subjects with a type prefix, e.g.:

```
feat: add CreateBehaviorTreeFromConfigSafe returning an error
fix: correct copy-paste panic messages in decorators
docs: add Design Notes page
test: add core behavior tests for composites
```

## Naming conventions

- External brand, repository, module path, docs: `behavior-tree`.
- Go package declaration: `behaviortree` (Go package names cannot contain
  hyphens).
- Import alias used throughout the code and examples: `bt`.

## Documentation

User-facing docs live under `website/` (Docusaurus, bilingual 中文 / English).
Chinese is the default locale; English translations live under
`website/i18n/en/`. When you change behavior, update the matching doc page in
both languages.

## Reporting issues

When filing a bug, please include the behavior tree JSON (or a minimal
reproduction), the Go version, and what you expected versus what happened.
