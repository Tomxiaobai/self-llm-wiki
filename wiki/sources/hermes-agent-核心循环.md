---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/02-Agent核心循环.md
tags:
  - source
  - hermes-agent
  - agent-loop
---

# Hermes Agent 核心循环

## 一句话摘要

这篇笔记聚焦 `AIAgent` 主循环，解释了 Hermes Agent 如何通过同步 while 循环、预算控制、并行工具执行与错误重试组织整个对话生命周期。

## 这份资料回答了什么问题

- `AIAgent` 为什么会变成一个大而全的核心类
- `IterationBudget` 如何约束工具调用和死循环
- 并行工具执行与流式处理是怎么组合在一起的

## 核心结论

- Hermes Agent 的主循环是典型的“执行循环式 harness”，而不是事件回调式或纯声明式工作流。
- `IterationBudget` 不只是硬上限，还有 70% / 90% 的压力注入机制，引导模型主动收尾。
- 并行执行不是默认全开，而是在安全检查通过后才进入线程池分发。

## 与现有 Wiki 的关系

- 直接丰富 [[agent-harness]] 中的 `E` Execution Loop
- 与 [[context-compression]]、[[tool-registry]] 在运行时层面形成闭环
- 在 [[hermes-agent-三方对比]] 中被拿来与 Claude Code 和 OpenClaw 对照

## 来源

- `raw/hermes-agent/02-Agent核心循环.md`
