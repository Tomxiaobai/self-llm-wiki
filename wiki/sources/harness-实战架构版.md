---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/06_Harness_实战架构版.md
tags:
  - source
  - harness
  - agent
  - governance
---

# Harness 实战架构版

## 一句话摘要

这份资料把 Harness 六组件框架翻译成更贴近 Coding Agent 和平台工程的语言，是本仓库理解 Agent 治理层的桥梁文档。

## 这份资料回答了什么问题

- 为什么 Agent 不等于“模型 + Prompt”
- Execution、Tool Registry、Context、State、Hooks、Evaluation 各自解决什么问题
- MCP、A2A、AGENTS/Skills 分别落在 Harness 的哪一层

## 核心结论

- Harness 决定 Agent 是否可恢复、可控、可审计，而不只是是否“会回答”。
- 模型不能直接操作世界，必须经由工具注册表与生命周期钩子。
- 对复杂 Agent 系统而言，治理层不是附属品，而是平台可用性的前提条件。

## 与现有 Wiki 的关系

- 是 [[agent-harness]] 的重要来源
- 解释了 [[a2a]]、[[eino]]、[[agent-platform]] 在整体系统里的分层关系
- 为理解 `raw/harness_demo.py` 提供概念框架

## 未解决问题

- 资料偏框架解释，缺少量化评估方法
- 需要与 [[pdf-论文目录与首轮摘要]] 中的 Harness Survey 对读

## 来源

- `raw/06_Harness_实战架构版.md`
