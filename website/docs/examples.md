---
id: examples
title: 示例
sidebar_position: 4
---

# 示例

仓库 `examples/` 目录下提供了多种加载与使用方式的示例：

| 示例 | 说明 |
| --- | --- |
| [load_from_tree](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_tree) | 从导出的**树文件**加载 |
| [load_from_project](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_project) | 从导出的**工程文件**加载 |
| [load_from_rawproject](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_rawproject) | 从**原生工程文件**加载 |
| [subtree](https://github.com/henrytien/behavior-tree/tree/master/examples/subtree) | 子树的使用示例（需专用编辑器分支版本 `behavior3editor`）|
| [memsubtree](https://github.com/henrytien/behavior-tree/tree/master/examples/memsubtree) | 记忆子树示例 |
| [testwork](https://github.com/henrytien/behavior-tree/tree/master/examples/testwork) | 节点 worker 测试示例 |
| [mmoarpg](https://github.com/henrytien/behavior-tree/tree/master/examples/mmoarpg) | 一个 MMOARPG 类型游戏的行为示例，可用桌面版编辑器打开查看（Projects → open project）|

## 运行示例

```bash
cd examples/load_from_tree
go run main.go
```

## 网页版编辑器本地搭建

1. 下载 `behavior3editor` 源码到本地。
2. 安装 Node.js、npm。
3. 安装 bower：`npm install -g bower`
4. 进入工程目录：`npm install` 然后 `bower install`
5. 安装 gulp：`npm install --global gulp`
6. 运行：在工程目录下 `gulp serve`
7. 浏览器打开 `http://127.0.0.1:8000`

自行部署 web 版本时，只需把生成的 `build` 目录放到 Tomcat / IIS 目录即可通过浏览器访问。

## 完整示例

io 类游戏示例位于 `h5game/server`。`bin/b3.json` 为行为树数据：在编辑器中新建任意工程，选择 **Project → Import → Tree as json** 导入树即可还原工程。
