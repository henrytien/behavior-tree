---
id: faq
title: FAQ
sidebar_position: 5
---

# 常见问题 FAQ

### Q：子树的相同记忆节点的黑板信息是重复的？

> 是个问题，但由于下一次进入会在 `OnOpen` 重置节点黑板信息，并不会有错误表现。如果遇到 BUG 表现欢迎提出。

### Q：Tick 里的 target 如何调用？

> 用在 AI 里，一般 `target` 就是这个 AI 的拥有者，拥有者持有 blackboard 成员。内建节点不使用 `target`，它只会被自定义节点访问。

### Q：如何设计打断一个进行中（RUNNING）的状态？

> 参考 [issue #15](https://github.com/henrytien/behavior-tree/issues/15)。核心思路是利用非记忆型的 `Priority`：每帧重新评估高优先级分支，当高优先级条件成立时，父节点会自动关闭正在 RUNNING 的低优先级子树（触发其 `OnClose`）。

## TODO

- [ ] 参数类型化
- [ ] 参数支持传递黑板值，利用格式 `@变量名`
- [ ] 子树支持自定义参数传递

## 其他参考

- 一些节点的概念：[有道云笔记](http://note.youdao.com/noteshare?id=4f46dc2144ea62b55f597630f5e666b4&sub=FF0B4E1D7916473E8DFC7242CFC62B69)
- 行为树基础知识：[behaviac concepts](https://www.behaviac.com/concepts/)
- 其他人写的旧版 behavior3go 代码介绍：[CSDN 博客](https://blog.csdn.net/u013272009/article/details/77131226)
- behavior3 的 Erlang 实现（含行为树基础分享）：[behavior3erl](https://github.com/dong50252409/behavior3erl)

## 上线项目

- [丛林大作战](https://www.taptap.com/app/31608)

## 联系

QQ 群：285728047
