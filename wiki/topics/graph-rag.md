---
type: topic
status: active
updated: 2026-04-23
source_paths:
  - raw/2408.08921v2.pdf
  - raw/GraphRAG.pdf
  - raw/3626772.3661370.pdf
  - raw/btae353.pdf
  - raw/s41598-025-21222-z.pdf
tags:
  - topic
  - graph-rag
  - kg-rag
---

# Graph RAG

## 主题定义

Graph RAG 指把文本或知识先组织成图结构，再基于实体、关系、社区或图遍历完成检索与证据组合的 RAG 形态。

## 为什么重要

- 它更适合多跳推理、跨文档实体关联和全局结构理解。
- 相比“按 chunk 检索”，图结构更容易表达长距离关系和主题聚类。
- 本仓库中的多篇论文都把知识图谱视为缓解幻觉、提升问答质量的关键手段。

## 当前仓库中的共识

- 图结构能提升复杂问题和关系推理场景下的检索质量。
- GraphRAG 和 KG-RAG 通常都强调实体、关系、路径或社区层次，而不是单纯相似度检索。
- 这类方法很适合客服问答、生物医学、企业知识等结构关系明显的领域。

## 来自综述页的细化

- [[graph-rag-survey]] 明确把 GraphRAG 分成三段流水线：
  - Graph-Based Indexing
  - Graph-Guided Retrieval
  - Graph-Enhanced Generation
- 检索阶段的设计空间不只是“是否用图”，还包括：
  - retriever 类型
  - 一次检索 / 多阶段检索 / 迭代检索
  - 粒度是节点、三元组、路径、子图还是混合形式
- 未来瓶颈也更清楚了：
  - 动态图
  - 多模态图
  - 可扩展检索
  - 标准 benchmark 缺失
  - 检索上下文压缩

## 与传统 RAG 的差异

- 检索对象不只是文本块，而是实体、边、子图、社区
- 优势在复杂关联和多跳推理
- 代价是图构建和更新成本更高，细节召回可能依赖补充机制

## 与 Agentic RAG 的关系

- Graph RAG 关注“知识如何组织与检索”
- [[agentic-rag]] 关注“检索如何被纳入多步执行循环”
- 两者可以叠加：Agent 用图检索作为一个工具或中间记忆层

## 代表性来源

- [[graph-rag-survey]]
- [[pdf-论文目录与首轮摘要]]
- [[可视化与辅助资料]]
- [[eino-a2a-rag-实战项目蓝图]]
