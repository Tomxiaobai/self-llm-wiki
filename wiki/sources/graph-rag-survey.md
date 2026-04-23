---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/2408.08921v2.pdf
tags:
  - source
  - pdf
  - survey
  - graph-rag
---

# Graph Retrieval-Augmented Generation: A Survey

## 一句话摘要

这篇综述把 GraphRAG 正式抽象成一条完整流水线：Graph-Based Indexing、Graph-Guided Retrieval、Graph-Enhanced Generation，并进一步覆盖训练、评测、产业应用与未来方向。

## 论文关注的问题

- GraphRAG 与传统 RAG、LLMs on Graphs、KBQA 的边界分别是什么
- 图结构怎样进入索引、检索和生成各阶段
- GraphRAG 的检索粒度、检索器类型和增强手段有哪些
- 当前 benchmark、应用和未来瓶颈在哪里

## 论文结构速览

- 第 7 页起：GraphRAG 总体概览
- 第 8 页起：Graph-Based Indexing
- 第 10 页起：Graph-Guided Retrieval
- 第 17 页起：Graph-Enhanced Generation
- 第 23 页起：Training
- 第 24 页起：Applications and Evaluation
- 第 29 页起：Future Prospects

## 核心结论

- 论文把 GraphRAG 定义为一条端到端流程，而不是简单“把知识图谱接到 RAG 前面”。
- 检索阶段不只关心是否用图，还关心：
  - 使用什么 retriever
  - 做一次检索还是多阶段 / 迭代检索
  - 检索粒度是节点、三元组、路径、子图还是混合形式
- 生成阶段也不是被动消费图数据，还涉及图格式选择和 generation enhancement。
- GraphRAG 的未来压力点很明确：
  - 动态与自适应图
  - 多模态信息融合
  - 可扩展检索效率
  - 与 Graph Foundation Model 结合
  - 检索上下文的无损压缩
  - 标准 benchmark 缺失

## 对本仓库最有价值的增量

- 它把 [[graph-rag]] 从“知识图谱增强 RAG 的一个名字”升级成了结构化研究对象。
- 相比 `raw/GraphRAG.pdf` 这种介绍型材料，这篇综述更适合用来拆解系统设计空间。
- 对 [[agentic-rag]] 来说，它提供了 graph-based retrieval 能力在 Agent 系统中的清晰模块化入口。

## 与现有 Wiki 的关系

- 是 [[graph-rag]] 的核心综述来源
- 丰富了 [[rag]] 对“检索对象与结构”的理解
- 为 [[agentic-rag]] 提供 graph-based agent tools / memory substrate 的上游背景

## 仍需谨慎的地方

- 论文强调流程全景，但不同工作之间的实验条件并不天然可直接横向比较。
- 产业应用与 benchmark 部分提示了方向，但标准化仍明显不足。
- 如果后续要服务具体领域，还要回到客服、生物医学等场景论文继续细化。

## 来源

- `raw/2408.08921v2.pdf`
