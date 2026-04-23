---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/03_A2A_协议详解.md
tags:
  - source
  - a2a
  - protocol
  - multi-agent
---

# A2A 协议详解

## 一句话摘要

这份资料系统解释了 A2A 作为 Agent 间任务委派协议的目标、交互流程、Agent Card 与 Task 状态机。

## 这份资料回答了什么问题

- 为什么需要 Agent-to-Agent 协议
- Agent Card、Task、Streaming、Webhook 在协议中扮演什么角色
- A2A 和 MCP 的边界在哪里

## 核心结论

- A2A 的核心价值不是“多一个 API”，而是跨框架 Agent 的能力发现与任务协作。
- Agent Card 是能力发现入口，Task 是长任务协作的核心状态单元。
- A2A 更像多 Agent 执行循环的跨服务延展，不等同于工具协议。

## 与现有 Wiki 的关系

- 是 [[a2a]] 的主来源
- 支撑 [[agentic-rag]] 和 [[agent-platform]] 的多 Agent 协作设计
- 与 [[eino]]、[[agent-harness]] 一起构成平台架构三角

## 未解决问题

- 资料强调协议结构，但没有展示完整生产级安全、鉴权和失败恢复实现
- 需要结合 [[agent-platform-项目骨架]] 观察 A2A 在代码中的实际落点

## 来源

- `raw/03_A2A_协议详解.md`
