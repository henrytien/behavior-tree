# 设计方案：AI 友好的行为树编辑格式

> 状态：**草案 / 待评审**　·　作者：henrytien　·　目标读者：维护者

## 1. 背景与问题

当前编辑器（`behavior-tree-editor`）导出的 JSON 是为"可视化编辑器读写"设计的，**不适合人或 AI 直接编写**。实测发现，让 AI 编辑行为树很容易产出"坏树"。根因不是编辑器有 bug，也不是运行时逻辑有问题，而是**数据格式本身对手写不友好**。

### 现有格式（编辑器导出）

经核对编辑器源码（`ExportManager.js`）与实际样本（`examples/load_from_project/project.json`、`examples/subtree/example.b3`），格式如下：

- **Project → 多棵树**：数组 `trees: [ ... ]`
- **单棵树 → 节点**：Map，`nodes: { "<uuid>": {...} }`，key 是节点 UUID
- **连线**：composite 存 `children: ["<uuid>", ...]`；decorator 存 `child: "<uuid>"`；`root: "<uuid>"`
- **子树引用**：节点 `category: "tree"`，其 `name` 字段 = 被引用那棵树的 **UUID**；主树与子树同在一个 project 文件里
- 每个节点还带 `display: {x, y}` 坐标和相机状态

### 为什么 AI 容易写坏

| 风险点 | 说明 |
| --- | --- |
| **UUID 交叉引用** | `children` / `child` / `root` / 子树 `name` 全靠 36 位 UUID 互指。AI 极易引用不存在的 UUID、或漏改某一处引用，导致孤儿节点或加载 panic |
| **信息冗余** | 节点既是 `nodes` 的 key，内部又重复一个 `id` 字段，两者必须一致 |
| **坐标噪声** | 每个节点的 `display:{x,y}` 与 AI 的逻辑意图无关，AI 要么瞎编、要么漏写 |
| **category 与连线强耦合** | decorator 只能 `child`、composite 只能 `children`、action/condition 不能有子节点；配错即坏 |

结论：这套格式当年的设计目标就是"机器生成、机器读取"，从未打算给人/AI 手写。

## 2. 设计目标

1. **AI/人能直接读写**：无 UUID、无坐标、用缩进或嵌套表达父子关系。
2. **引用按名字**，不按 UUID。
3. **可双向转换**：DSL ⇄ 编辑器 JSON，保证现有编辑器与运行时不受影响。
4. **强校验**：转换时即报错（引用缺失、category 非法、子树未定义），把"坏树"挡在加载之前。
5. **零运行时改动**：Go 运行时和编辑器都不需要改；这是一层独立的工具。

## 3. 提议方案：中间 DSL + 转换器

### 3.1 DSL 形态（YAML 示例）

```yaml
# 一个 project 可含多棵树；用 name 互相引用，没有 UUID
trees:
  main:
    root:
      MemSequence:
        - Wait: { milliseconds: 1000 }
        - Priority:
            - SubTree: patrol        # 按名字引用另一棵树
            - Log: { info: "idle" }

  patrol:
    root:
      Sequence:
        - Log: { info: "move" }
        - Wait: { milliseconds: 500 }
```

要点：
- 父子关系用**嵌套/缩进**表达，不需要任何 id。
- 节点类型直接写名字（`Sequence`/`Wait`/...），属性写在 `{ }` 里。
- 子树用 `SubTree: <treeName>`，转换器负责把 `treeName` 解析成对应树的 UUID。
- composite 的子节点是列表，decorator 只接受单个子节点（列表长度必须为 1，否则报错）。

> YAML 仅为示例。也可考虑：① 仍用 JSON 但简化（嵌套式、无 UUID）；② 自定义紧凑 DSL。YAML 对 AI 最友好、可读性最好，倾向首选。

### 3.2 转换器（DSL → 编辑器 JSON）

一个独立工具（建议 Go，复用现有 `config` 包的结构体；或 Node 脚本以贴近编辑器）。职责：

1. 遍历 DSL 树，为每个节点**生成 UUID**（可复用根包的 `CreateUUID`）。
2. 把嵌套关系翻译成 `children` / `child` / `root` 的 UUID 引用。
3. **自动布局**：按树深度/兄弟序生成合理的 `display:{x,y}`，无需人工干预。
4. 把 `SubTree: <name>` 解析为 `category:"tree"` + `name:"<被引用树的 UUID>"`。
5. 组装成 project 结构（`trees` 数组、`custom_nodes` 等），输出编辑器可直接打开的 `.b3` / project JSON。

### 3.3 校验器（独立可用）

转换前先做静态校验，**这是最高性价比、可以最先单独落地的一步**：

- 所有被引用的子树名/节点都存在；
- decorator 恰好 1 个子节点，action/condition 无子节点；
- 节点类型已注册（在 `loader` 的内建表 + 自定义表里）；
- 属性类型符合各节点要求（如 `Wait.milliseconds` 是整数）。

校验器也可以**反向用在现有编辑器导出的 JSON 上**：检查 AI 直接改过的 `.b3` 是否引用完整，立刻定位坏在哪。

### 3.4 反向转换（编辑器 JSON → DSL，可选）

便于把已有的可视化树导出成 DSL 给 AI 阅读/修改。优先级低于正向。

## 4. 放在哪里

建议**独立目录或独立小工具**，不污染库本身：

- 选项 A：本仓库新增 `cmd/btc/`（behavior-tree compiler），`go run ./cmd/btc compile main.yaml -o main.b3`。复用 `config`/根包，零新依赖。
- 选项 B：放在编辑器仓库，作为编辑器的"导入 DSL"功能。
- 倾向 **A**：Go 实现、与运行时同仓、CI 可直接测。

## 5. 与"学习 UE"的关系

UE 行为树真正值钱的是两点，**与数据格式无关**，可单独借鉴：

1. **条件中断 / Observer Aborts**：高优先级分支条件变化时自动打断正在跑的低优先级子树。本实现**已能做到**（非记忆 `Priority` 每帧重评 + `OnClose`），只是缺文档和语法糖。→ 可作为后续增强：在 `design-notes` 补"如何实现可打断行为"，或加一个显式的"条件守卫"装饰器。
2. **Service 节点**（挂在分支上、按间隔刷新黑板）：本实现**没有**，是可新增的能力。

> 注意：UE 用二进制资产 + 专用编辑器，并不解决"AI 友好文本格式"问题。所以"学 UE" ≠ "让 AI 更好编辑"，两者是不同的方向，可并行。

## 6. 分阶段落地建议

1. **阶段一（先做、收益最高）**：校验器。输入编辑器 JSON，输出"引用完整性 / category 合法性 / 类型"报告。立刻能接住 AI 生成的坏树。
2. **阶段二**：DSL → JSON 正向转换器（含自动布局）。
3. **阶段三**：反向转换 + 与编辑器/CI 集成。
4. **并行增强（独立）**：文档化"可打断行为"，评估新增 Service 节点。

## 7. 待决问题

- DSL 选 YAML / 简化 JSON / 自定义？（倾向 YAML）
- 转换器放本仓库 `cmd/` 还是编辑器仓库？（倾向本仓库）
- 自动布局算法的精度要求（能打开即可 vs 美观）？
- 是否需要保留往返保真（round-trip：JSON→DSL→JSON 不丢信息）？

---

*本文件为设计草案，供评审。确认方向后再进入实现。*
