---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/02_Eino_框架指南.md
tags:
  - source
  - eino
  - framework
  - orchestration
---

# Eino 框架指南

## 一句话摘要

这份资料把 Eino 拆成“组件接口 + 编排系统”两条主线，是本仓库理解 Go Agent 应用框架的基础入口。

## 这份资料回答了什么问题

- Eino 的核心抽象是什么
- `Chain`、`Graph`、`Workflow` 分别适合什么场景
- RAG、Tool Use、Agent 编排在 Eino 里如何落地

## 核心结论

- Eino 的稳定性来自统一的组件接口，而不是单一模板或脚手架。
- `Graph` 是实现 Agent 循环、分支、并行和状态注入的关键抽象。
- 在本仓库语境中，Eino 更像单 Agent 内部的运行编排层，而不是多 Agent 协议层。

## 与现有 Wiki 的关系

- 是 [[eino]] 的直接来源
- 为 [[agentic-rag]]、[[agent-platform]] 提供单 Agent 内部工作流实现视角
- 与 [[a2a]] 形成“内部编排 vs 外部协作”的分工关系

## 未解决问题

- 资料主要描述抽象和接口，缺少大规模生产经验对比
- 对 MCP 的接入方式没有展开，需要结合 [[eino-mcp-a2a-harness-项目设计稿]]

## 来源

- `raw/02_Eino_框架指南.md`
