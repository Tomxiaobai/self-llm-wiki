---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/07_Eino_MCP_A2A_Harness_项目设计稿.md
tags:
  - source
  - design
  - mcp
  - a2a
  - harness
---

# 实战项目设计稿：基于 Eino + MCP + A2A + Harness 的 Agent 平台

## 一句话摘要

这份资料把 Eino、MCP、A2A、Harness 四条线收束成一个平台级分层方案，回答“这些能力为什么要同时出现”。

## 这份资料回答了什么问题

- 为什么单独用 Eino 不够
- MCP、A2A、Harness 分别补哪一层能力
- 企业知识问答与研究分析平台的推荐分层是什么

## 核心结论

- Eino 负责 Agent 内部怎么跑，MCP 负责工具怎么接，A2A 负责 Agent 怎么协作，Harness 负责系统怎么受控地运行。
- 平台化设计需要把治理能力作为横切层，而不是揉进业务代码。
- 第一版应优先追求“稳、可解释、可恢复、可扩展”，而不是全自动和全能力开放。

## 与现有 Wiki 的关系

- 是 [[agent-platform]] 的上层设计来源
- 强关联 [[eino]]、[[a2a]]、[[agent-harness]]
- 为 [[项目脚手架与代码骨架]] 提供架构前提

## 未解决问题

- 文档假设了 MCP 角色，但当前仓库还没有完整 MCP 资料沉淀
- 仍需通过 `raw/agent-platform/` 验证哪些设计已经进入代码，哪些仍停留在蓝图

## 来源

- `raw/07_Eino_MCP_A2A_Harness_项目设计稿.md`
