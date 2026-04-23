---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/
tags:
  - catalog
  - raw
  - ingest
---

# Raw Source Catalog

这是 `raw/` 的首轮映射目录，用来说明每份原始资料目前在 Wiki 中的落点。

## 已建来源页

| Raw 文件或目录 | 对应 Wiki 页面 | 说明 |
|---|---|---|
| `raw/02_Eino_框架指南.md` | [[eino-框架指南]] | Eino 组件与编排入门稿 |
| `raw/03_A2A_协议详解.md` | [[a2a-协议详解]] | A2A 协议结构化说明 |
| `raw/04_Agent_设计模式.md` | [[agent-设计模式]] | ReAct / Plan-Execute / Multi-Agent / Corrective RAG |
| `raw/05_实战项目蓝图.md` | [[eino-a2a-rag-实战项目蓝图]] | Eino + A2A RAG 系统蓝图 |
| `raw/06_Harness_实战架构版.md` | [[harness-实战架构版]] | Harness 六组件实战解释 |
| `raw/07_Eino_MCP_A2A_Harness_项目设计稿.md` | [[eino-mcp-a2a-harness-项目设计稿]] | 平台级项目设计稿 |
| `raw/08_项目脚手架与代码骨架.md` | [[项目脚手架与代码骨架]] | 可开工的目录与接口骨架 |
| `raw/harness-agent-深度解读.md` | [[harness-agent-深度解读]] | Agent Harness Survey 中文深读 |
| `raw/harness_demo.py` | [[harness-demo-源码说明]] | Harness 六组件的最小代码示例 |
| `raw/agent-platform/` | [[agent-platform-项目骨架]] | Go 项目骨架、文档与代码结构 |
| `raw/hermes-agent/` | [[hermes-agent-资料集]] | Hermes Agent 8 篇结构化源码解读与项目映射 |
| `raw/2501.09136v3.pdf` | [[agentic-rag-survey]] | Agentic RAG 核心综述，已拆独立论文页 |
| `raw/2408.08921v2.pdf` | [[graph-rag-survey]] | GraphRAG 核心综述，已拆独立论文页 |
| `raw/harness-agent.pdf` | [[agent-harness-survey]] | Harness 核心综述原文页 |

## 批量目录页

| Raw 文件集合 | 对应 Wiki 页面 | 说明 |
|---|---|---|
| `raw/*.pdf` | [[pdf-论文目录与首轮摘要]] | 10 份 PDF 的标题、主题和首轮用途判断 |
| `raw/*.svg` + `raw/agentic_rag_workflow_interactive.html` | [[可视化与辅助资料]] | 图示、taxonomy、交互流程图的首轮整理 |

## 元信息与非核心资产

| Raw 文件 | 当前处理方式 | 说明 |
|---|---|---|
| `raw/AGENTS.md` | 作为历史元说明保留 | 早期仓库说明，已被根级 [[CLAUDE]] 取代 |
| `raw/CLAUDE.md` | 作为历史元说明保留 | 说明旧的仓库定位，不再作为主规范 |
| `raw/Jianying_37636880321514_installer.dmg` | 不纳入知识主线 | 二进制安装包，默认只记录存在 |
| `raw/.DS_Store` | 忽略 | 系统生成文件 |

## 相关页面

- [[overview]]
- [[index]]
- [[rag]]
- [[graph-rag]]
- [[agentic-rag]]
- [[agent-harness]]
- [[hermes-agent]]
