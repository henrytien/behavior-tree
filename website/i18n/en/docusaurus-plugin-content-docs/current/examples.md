---
id: examples
title: Examples
sidebar_position: 4
---

# Examples

The `examples/` directory contains several loading and usage styles:

| Example | Description |
| --- | --- |
| [load_from_tree](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_tree) | Load from an exported **tree file** |
| [load_from_project](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_project) | Load from an exported **project file** |
| [load_from_rawproject](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_rawproject) | Load from a **raw project file** |
| [subtree](https://github.com/henrytien/behavior-tree/tree/master/examples/subtree) | SubTree usage (requires the dedicated editor branch `behavior3editor`) |
| [memsubtree](https://github.com/henrytien/behavior-tree/tree/master/examples/memsubtree) | Memory subtree example |
| [testwork](https://github.com/henrytien/behavior-tree/tree/master/examples/testwork) | Node worker test example |
| [mmoarpg](https://github.com/henrytien/behavior-tree/tree/master/examples/mmoarpg) | A behavior example for an MMOARPG-style game; open it in the desktop editor (Projects → open project) |

## Running an example

```bash
cd examples/load_from_tree
go run main.go
```

## Running the web editor locally

1. Download the `behavior3editor` source.
2. Install Node.js and npm.
3. Install bower: `npm install -g bower`
4. In the project directory: `npm install` then `bower install`
5. Install gulp: `npm install --global gulp`
6. Run: `gulp serve` in the project directory
7. Open `http://127.0.0.1:8000` in a browser

To deploy a web build yourself, just drop the generated `build` directory into your Tomcat / IIS directory.

## Full example

An io-style game example lives under `h5game/server`. `bin/b3.json` is the behavior tree data: create any new project in the editor, then choose **Project → Import → Tree as json** to restore the project.
