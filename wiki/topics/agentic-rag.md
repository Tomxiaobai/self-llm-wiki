---
type: topic
status: active
updated: 2026-04-23
source_paths:
  - raw/04_Agent_设计模式.md
  - raw/05_实战项目蓝图.md
  - raw/2501.09136v3.pdf
tags:
  - topic
  - agentic-rag
  - agent
---

# Agentic RAG

## 主题定义

Agentic RAG 把检索从一次性前处理，升级为 Agent 在执行循环中的一等能力。模型可以按任务需要动态决定检索、重写、补搜、调用工具和验证答案。

## 为什么重要

- 它更适合复杂问题、长任务和多来源综合，而不是只回答简单事实问答。
- 它天然要求规划、工具治理、状态管理和错误恢复，因此与 Harness 问题紧密相连。
- 本仓库大多数设计稿都把 Agentic RAG 视为企业级研究分析平台的核心范式。

## 当前仓库中的共识

- 检索结果需要被评分，而不是无条件输入生成模型。
- 查询重写、补充 Web 搜索、幻觉检查和多步重试应被纳入执行循环。
- Gateway / Supervisor 角色经常需要把 RAG 任务拆分给不同子 Agent 或工具。

## 来自综述页的细化

- [[agentic-rag-survey]] 把 Agentic RAG 的基础能力收束为四项：
  - Reflection
  - Planning
  - Tool Use
  - Multi-Agent
- 它还把系统形态分成多种稳定 workflow / architecture：
  - Prompt Chaining
  - Routing
  - Parallelization
  - Orchestrator-Workers
  - Evaluator-Optimizer
  - Single-Agent / Multi-Agent / Hierarchical / Corrective / Adaptive / Graph-Based
- 这意味着本仓库里的 Agentic RAG 不应只理解为“会重写 query 的 RAG”，而应理解为一整套可组合的执行模式族。

## 设计要点

- 内部编排：由 [[eino]] 或类似框架实现 Graph / Branch / Retry
- 跨 Agent 协作：由 [[a2a]] 等协议支持
- 运行时治理：由 [[agent-harness]] 提供上下文、状态、审批和评估

## 与相邻主题的关系

- 基于 [[rag]]，但不止于一次检索
- 可叠加 [[graph-rag]] 作为结构化检索能力
- 在系统层面通常落进 [[agent-platform]]

## 代表性来源

- [[agent-设计模式]]
- [[eino-a2a-rag-实战项目蓝图]]
- [[agentic-rag-survey]]
- [[pdf-论文目录与首轮摘要]]
- [[可视化与辅助资料]]
