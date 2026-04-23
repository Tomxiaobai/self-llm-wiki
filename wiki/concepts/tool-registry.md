---
type: concept
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/03-Tool-Registry.md
  - raw/06_Harness_实战架构版.md
tags:
  - concept
  - tools
  - registry
---

# Tool Registry

## 概念定义

Tool Registry 指一个 Agent 运行时用于登记、发现、筛选、分发工具调用的统一机制。它不仅决定“有哪些工具”，还决定“谁现在可用、怎么调用、结果怎么约束”。

## 为什么重要

- 模型不能直接操作外部世界，工具注册表就是连接模型与环境的闸门。
- 一个真实可用的 Tool Registry 必须处理 schema、可用性检查、权限、结果大小、异步桥接和错误标准化。
- 这正对应 [[agent-harness]] 里的 `T` Tool Registry。

## Hermes Agent 给出的实践要点

- 用单例 `registry` + 模块级 `register()` 实现低耦合发现链
- 用 `ToolEntry` 显式保存 schema、handler、check_fn、requires_env 等元信息
- 用 `check_fn` 和环境检查把“看得见的工具”和“装得好的工具”区分开
- 用持久化 event loop 处理 sync/async 阻抗
- 用结果大小控制避免工具输出把上下文打爆

## 与其他系统的关系

- Hermes 代表的是朴素 Registry 风格
- [[a2a]] 解决的是 Agent 间协作，不是工具注册
- MCP 则可以看作更标准化、跨系统化的工具暴露协议

## 代表性来源

- [[hermes-agent-tool-registry]]
- [[harness-实战架构版]]
