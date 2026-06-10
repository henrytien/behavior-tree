---
title: behavior3go
---

# behavior3go

## Behavior Tree for Go

A Go implementation of the [behavior3](https://github.com/behavior3) behavior tree, fully compatible with the official online editor's data format.

### Quick Links

- [📖 Documentation](/docs/intro) — concepts, node reference, and examples
- [🚀 Getting Started](/docs/getting-started)
- [💻 GitHub Repository](https://github.com/henrytien/behavior-tree)

### Features

- **Stateless tree** — only one copy of the tree structure exists; state lives in the Blackboard, so one tree can drive hundreds of objects.
- **Editor compatible** — translated from the JS version; identical data format to the official behavior3 online editor.
- **SubTree support** — compose and reuse logic with `SubTree` nodes.
- **Rich built-in nodes** — four categories (Composite, Decorator, Action, Condition), including random wait and probability decorators.

> This repository is derived from `magicsea/behavior3go` and is now maintained independently as `github.com/henrytien/behavior-tree`.
