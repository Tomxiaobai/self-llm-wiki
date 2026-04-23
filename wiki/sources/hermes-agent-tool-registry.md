---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/03-Tool-Registry.md
tags:
  - source
  - hermes-agent
  - tools
---

# Hermes Agent Tool Registry

## 一句话摘要

这篇笔记解释了 Hermes Agent 如何用一个非常朴素的单例 Registry 模式管理工具注册、发现、校验、分发和 sync/async 桥接。

## 这份资料回答了什么问题

- 为什么 Hermes 没有选择复杂工具抽象
- `ToolEntry` 和 `ToolRegistry` 各自承担什么角色
- 工具发现、可用性检查和异步桥接是怎么完成的

## 核心结论

- Hermes 更重视低耦合和可维护的导入链，而不是在工具层叠很多抽象。
- 模块级延迟注册可以有效避开循环引用，同时让工具增删更直接。
- `check_fn`、结果大小控制和异步桥接这些细节，决定了工具系统是否能在真实任务里稳定工作。

## 与现有 Wiki 的关系

- 是 [[tool-registry]] 的核心来源
- 丰富 [[agent-harness]] 中的 `T` Tool Registry
- 可与 [[a2a]]、[[eino]]、[[agent-platform]] 的工具接入方式形成对照

## 来源

- `raw/hermes-agent/03-Tool-Registry.md`
