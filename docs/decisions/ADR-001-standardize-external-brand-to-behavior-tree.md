# ADR-001: Standardize External Brand to behavior-tree

## Status
Accepted

## Date
2026-06-10

## Context
The public repository and Go module path are `github.com/henrytien/behavior-tree`. Existing documentation used `behavior3go` as the visible project name in several places, which made the external brand inconsistent with the repository, documentation URL, and module path.

Go package declarations are identifiers and cannot contain hyphens. That means the external brand `behavior-tree` cannot also be the literal Go package declaration.

## Decision
Use `behavior-tree` as the external repository, documentation, and product brand.

Keep `github.com/henrytien/behavior-tree` as the Go module path. Use `behaviortree` for the root Go package declaration. Documentation examples may continue to alias imports as `b3` where that keeps behavior3-style constants and examples readable.

## Alternatives Considered

### behavior3go
Rejected as the external brand because it reflects upstream-derived project history rather than the current repository and module identity.

### behavior-tree-go
Rejected because it implies a repository or module rename from `github.com/henrytien/behavior-tree`, which is unnecessary for the current Go module.

### package behavior-tree
Rejected because it is not a valid Go package declaration. Go identifiers cannot contain hyphens.

## Consequences
- Public documentation, website metadata, and page titles should say `behavior-tree`.
- The root package declaration should be `behaviortree`, not `behavior-tree`.
- Historical references to `magicsea/behavior3go` may remain only when explaining project provenance or legacy material.
- Code examples may use `b3 "github.com/henrytien/behavior-tree"` as an explicit import alias.
