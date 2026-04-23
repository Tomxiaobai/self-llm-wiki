---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/2501.09136v3.pdf
tags:
  - source
  - pdf
  - survey
  - agentic-rag
---

# Agentic Retrieval-Augmented Generation: A Survey on Agentic RAG

## 一句话摘要

这篇综述把 Agentic RAG 从“传统 RAG 的增强版”提升为一个独立范式，系统整理了其核心能力、工作流模式、体系分类、应用场景与评测资源。

## 论文关注的问题

- 传统 RAG 为什么难以应对多步推理、动态检索和复杂任务管理
- Agent 能力如何改变检索与生成的耦合方式
- Agentic RAG 可以被分成哪些稳定架构
- 真实落地时有哪些工具、数据集、应用领域和挑战

## 论文结构速览

- 第 1 页：摘要与研究定位
- 第 9 页起：Agentic Intelligence 的四个基础能力
- 第 11 页起：工作流模式（workflow patterns）
- 第 14 页起：Agentic RAG taxonomy
- 第 29 页起：框架对比与应用
- 第 31 页起：工具与框架
- 第 32 页起：benchmarks 与 datasets

## 核心结论

- 论文把 Agentic RAG 的关键增量概括为四类能力：`Reflection`、`Planning`、`Tool Use`、`Multi-Agent`。
- Agentic RAG 不是单一架构，而是一组工作流与系统形态的集合。
- 工作流模式层面，论文明确区分了：
  - Prompt Chaining
  - Routing
  - Parallelization
  - Orchestrator-Workers
  - Evaluator-Optimizer
- 体系分类层面，论文给出了从单 Agent 到图增强、纠错式、自适应、多 Agent 的连续谱系：
  - Single-Agent Router
  - Multi-Agent
  - Hierarchical
  - Corrective
  - Adaptive
  - Graph-Based Agentic RAG
  - Agentic Document Workflows

## 对本仓库最有价值的增量

- 它把 [[agent-设计模式]] 中零散出现的 Agent 模式，提升成了可直接用于系统分类的 taxonomy。
- 它把 [[graph-rag]] 与 [[agentic-rag]] 的关系说清了：图结构可以不是独立方向，也可以被吸收进 Agentic RAG 作为检索与推理底座。
- 它让 [[agent-platform]] 的多角色设计更有理论锚点，因为 Orchestrator、Workers、Evaluator 这些结构在综述中都有明确位置。

## 与现有 Wiki 的关系

- 是 [[agentic-rag]] 的核心综述来源
- 反向支撑 [[rag]] 对范式演化的叙述
- 与 [[graph-rag]] 存在明显交叉，尤其在 graph-based agentic systems 一节
- 需要 [[agent-harness]] 来承接其长任务、状态、治理与评估需求

## 仍需谨慎的地方

- 这篇论文更强于分类与全景回顾，弱于严格的统一实验结论。
- 应用、工具和 benchmark 部分提供了方向，但不等于给出了成熟的工程基线。
- 如果后续要做更细的系统实现对比，仍需回到具体框架和代码仓库。

## 来源

- `raw/2501.09136v3.pdf`
