---
id: faq
title: FAQ
sidebar_position: 5
---

# FAQ

### Q: Is the blackboard data of identical memory nodes in a subtree duplicated?

> It is a known issue, but because the next entry resets the node's blackboard data in `OnOpen`, it does not manifest as an error. If you hit a real bug, please report it.

### Q: How do I use `target` inside Tick?

> It is used in AI: `target` is usually the owner of the AI, and the owner holds a blackboard member. Built-in nodes do not use `target` — it is only accessed by custom nodes.

### Q: How do I interrupt a RUNNING state?

> See [issue #15](https://github.com/henrytien/behavior-tree/issues/15). The idea is to use a non-memory `Priority`: re-evaluate the higher-priority branches every frame, so when a higher-priority condition becomes true, the parent automatically closes the lower-priority subtree that was RUNNING (triggering its `OnClose`).

## TODO

- [ ] Typed parameters
- [ ] Parameters that read blackboard values via the `@varName` format
- [ ] SubTrees with custom parameter passing

## Further reading

- Node concepts and practical tips: [Design Notes](./design-notes)
- Behavior tree basics: [behaviac concepts](https://www.behaviac.com/concepts/)
- A community write-up of legacy behavior3go code: [CSDN blog](https://blog.csdn.net/u013272009/article/details/77131226)
- An Erlang implementation of behavior3 (with behavior-tree fundamentals): [behavior3erl](https://github.com/dong50252409/behavior3erl)

## Shipped projects

- [Jungle Battle (丛林大作战)](https://www.taptap.com/app/31608)
