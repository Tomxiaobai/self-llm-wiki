---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/2408.08921v2.pdf
  - raw/2501.09136v3.pdf
  - raw/2509.04716v1.pdf
  - raw/2512.13564v2.pdf
  - raw/2601.16746v2.pdf
  - raw/3626772.3661370.pdf
  - raw/GraphRAG.pdf
  - raw/btae353.pdf
  - raw/harness-agent.pdf
  - raw/s41598-025-21222-z.pdf
tags:
  - source
  - pdf
  - survey
  - papers
---

# PDF 论文目录与首轮摘要

这是一份基于 PDF 元数据和首页文本抽取形成的首轮目录页，用来给论文批次建立可导航入口。

## 图谱 / 知识图谱增强 RAG

| 文件 | 标题 | 首轮判断 |
|---|---|---|
| `raw/2408.08921v2.pdf` | *Graph Retrieval-Augmented Generation: A Survey* | 已拆出 [[graph-rag-survey]]，是 [[graph-rag]] 的核心综述来源 |
| `raw/GraphRAG.pdf` | GraphRAG 中文介绍型材料 | 更像产品/方案介绍稿，适合补充直觉和对比表述 |
| `raw/3626772.3661370.pdf` | *Retrieval-Augmented Generation with Knowledge Graphs for Customer Service Question Answering* | 客服问答场景下的 KG-RAG 短文，适合做案例材料 |
| `raw/btae353.pdf` | *KRAGEN: a knowledge graph-enhanced RAG framework for biomedical problem solving using large language models* | 生物医学场景的 KG-RAG 框架，适合看领域化落地 |
| `raw/s41598-025-21222-z.pdf` | *Research on the construction and application of retrieval enhanced generation (RAG) model based on knowledge graph* | 知识图谱增强 RAG 的另一份场景型论文，可补实验与结构视角 |
| `raw/2509.04716v1.pdf` | *KERAG: Knowledge-Enhanced Retrieval-Augmented Generation for Advanced Question Answering* | 知识增强问答，适合理解 KG 与高级 QA 的结合方式 |

## Agent / Agentic RAG / Harness

| 文件 | 标题 | 首轮判断 |
|---|---|---|
| `raw/2501.09136v3.pdf` | *Agentic Retrieval-Augmented Generation: A Survey on Agentic RAG* | 已拆出 [[agentic-rag-survey]]，是 [[agentic-rag]] 的核心综述来源 |
| `raw/harness-agent.pdf` | *Agent Harness for Large Language Model Agents: A Survey* | 已拆出 [[agent-harness-survey]]，是 [[agent-harness]] 的原始论文来源 |
| `raw/2512.13564v2.pdf` | *Memory in the Age of AI Agents: A Survey Forms, Functions and Dynamics* | Agent Memory 综述，可作为后续扩展概念页的高价值来源 |
| `raw/2601.16746v2.pdf` | *SWE-Pruner: Self-Adaptive Context Pruning for Coding Agents* | Coding Agent 上下文裁剪论文，适合作为 Harness / Context 管理的补充证据 |

## 当前吸收方式

- 当前已为 3 篇核心综述建立独立来源页：
  - [[agentic-rag-survey]]
  - [[graph-rag-survey]]
  - [[agent-harness-survey]]
- 其余 PDF 仍保留在目录页层级，后续按主题优先级继续拆分。

## 推荐下一步

1. 为 `2512.13564v2.pdf` 拆出 Memory 综述页
2. 为 `2601.16746v2.pdf` 拆出 Context Pruning / Coding Agent 上下文页
3. 为 `2509.04716v1.pdf` 或 `3626772.3661370.pdf` 增补具体场景型 Graph / KG-RAG 页
4. 补 `MCP` 相关主题页，把协议层拼齐

## 相关页面

- [[rag]]
- [[graph-rag]]
- [[agentic-rag]]
- [[agent-harness]]
