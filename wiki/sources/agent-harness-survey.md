---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/harness-agent.pdf
tags:
  - source
  - pdf
  - survey
  - harness
---

# Agent Harness for Large Language Model Agents: A Survey

## 一句话摘要

这篇综述试图把“Agent 运行时基础设施”提升为独立研究对象，论点非常鲜明：现实世界里 Agent 系统的绑定约束，往往不是模型本身，而是包裹模型的 harness。

## 论文关注的问题

- 什么叫 agent harness，为什么它不是 prompt 或框架的附属品
- 如何用统一框架分析不同 Agent 系统的完整性
- 当前 Agent 生态在哪些治理能力上系统性欠缺
- 沙箱、安全、评估、协议标准化、上下文与记忆等难题应该如何组织讨论

## 论文结构速览

- 第 2 页：五项主要贡献与挑战全景
- 第 3 页：`H = (E, T, C, S, L, V)` 六组件框架
- 第 31 页：22 个系统的 completeness matrix
- 第 41 页起：Sandboxing and Security
- 第 50 页起：Evaluation and Benchmarking
- 第 61 页起：Runtime Context Management

## 核心结论

- 论文明确主张 harness 是独立研究对象，而不是“工具箱”或“执行细节”。
- 它提出的六组件框架是：
  - `E` Execution Loop
  - `T` Tool Registry
  - `C` Context Manager
  - `S` State Store
  - `L` Lifecycle Hooks
  - `V` Evaluation Interface
- 第 31 页的 completeness matrix 显示，生态里最常被低估和缺失的往往是 `L` 与 `V`。
- 论文把九类横切挑战串到一个统一问题域里，尤其强调：
  - sandboxing / security
  - evaluation / benchmarking
  - protocol standardization
  - context / memory management
  - compute economics

## 一个值得记录的细节

- 论文第 2 页摘要中的 completeness matrix 用了另一组更“功能描述式”的命名：
  - Execution environment
  - Tool integration
  - Context management
  - Scope negotiation
  - Loop management
  - Verification
- 但第 3 页正文图示清晰给出了本仓库当前采用的 `H = (E, T, C, S, L, V)` 表达。
- 在 Wiki 里，后者更适合作为稳定术语表，因为它与现有中文深读稿和 demo 代码更一致。

## 对本仓库最有价值的增量

- 它给 [[agent-harness]] 提供了真正的论文级锚点，而不仅是中文解释稿。
- 它解释了为什么 [[agent-platform]] 必须把治理层、协议层、评估层独立出来。
- 它还能反过来帮助校验 [[agentic-rag]]：一个 Agentic RAG 系统如果没有状态、审批、评估，通常只是“带工具的 RAG”，还谈不上稳定的 agent system。

## 与现有 Wiki 的关系

- 是 [[agent-harness]] 的原始论文来源
- 与 [[harness-agent-深度解读]] 形成“原文 + 中文深读”对应关系
- 与 [[harness-demo-源码说明]] 形成“理论框架 + 最小实现”对应关系

## 仍需谨慎的地方

- 当前版本是 preprint，尚未同行评审。
- 论文的历史叙事和分类框架很强，但具体系统评分仍受作者选择的定义和样本集影响。
- 若后续要做更严格的系统对标，最好把矩阵定义和评价标准再单独拆页。

## 来源

- `raw/harness-agent.pdf`
