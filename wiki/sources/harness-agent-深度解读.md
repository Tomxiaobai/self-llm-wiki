---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/harness-agent-深度解读.md
tags:
  - source
  - harness
  - survey
  - evaluation
---

# Agent Harness for LLM Agents: A Survey — 深度解读

## 一句话摘要

这份中文深读稿把 Harness Survey 的主要论点、分类矩阵、实证结果和技术挑战重新组织成了更容易吸收的工程笔记。

## 这份资料回答了什么问题

- 为什么 Harness 被视为 Agent 可靠性的关键约束
- 22 个系统如何按 E/T/C/S/L/V 分类
- 安全、评估、协议标准化面临哪些核心挑战

## 核心结论

- Harness Engineering 是比 Prompt Engineering 和 Context Engineering 更宽的一层工程问题。
- 仅改变 Harness 设计，就可能显著改变系统性能、成本和安全性。
- `L` 与 `V` 往往是最被低估的组件，但对生产系统极重要。

## 与现有 Wiki 的关系

- 为 [[agent-harness]] 提供综述级证据
- 与 [[harness-实战架构版]] 形成“概念综述 + 中文实战解释”互补
- 对 [[agent-platform]] 的治理层设计有直接指导意义

## 未解决问题

- 这是一份深读稿，不等于原论文全文细节
- 仍应与 `raw/harness-agent.pdf` 的原文结合核对

## 来源

- `raw/harness-agent-深度解读.md`
