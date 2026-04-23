---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/04_Agent_设计模式.md
tags:
  - source
  - agent
  - react
  - plan-execute
  - multi-agent
---

# Agent 设计模式

## 一句话摘要

这份资料把 Agent 拆解为 LLM、Memory、Tools、Orchestration 四要素，并给出 ReAct、Plan-Execute、Multi-Agent、Corrective RAG 等常见模式。

## 这份资料回答了什么问题

- Agent 系统由哪些基础能力构成
- 常见 Agent 模式如何在 Eino Graph 中编码
- Corrective RAG 和多 Agent 如何进入工程实现

## 核心结论

- Agent 模式差异，主要体现在“规划方式、工具使用方式、状态管理方式”。
- Corrective RAG 本质上是把检索、评估、重写、重试放进一个受控执行循环。
- 多 Agent 需要上层协调者或协议层，而不是简单并列多个模型调用。

## 与现有 Wiki 的关系

- 是 [[agentic-rag]] 和 [[agent-harness]] 的直接来源
- 为 [[eino]] 提供模式级使用场景
- 为 [[a2a]] 和 [[agent-platform]] 提供多 Agent 协作动机

## 未解决问题

- 资料重在模式解释，缺少失败模式与治理细节
- Memory 的长期演化和压缩策略需要结合 PDF 综述继续补强

## 来源

- `raw/04_Agent_设计模式.md`
