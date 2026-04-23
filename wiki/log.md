# Wiki Log

按时间追加记录 Wiki 的 ingest、query、lint、refactor 行为。

## [2026-04-23] setup | 初始化 LLM Wiki 结构

- Added: [[overview]]
- Updated: [[index]]
- Notes: 创建根级 `CLAUDE.md`，初始化 `wiki/index.md`、`wiki/overview.md`、`wiki/log.md`，约定 `raw/` 为原始资料层、`wiki/` 为 LLM 维护层。

## [2026-04-23] ingest | 首轮构建 raw -> wiki

- Added: [[raw-source-catalog]], [[eino-框架指南]], [[a2a-协议详解]], [[agent-设计模式]], [[eino-a2a-rag-实战项目蓝图]], [[harness-实战架构版]], [[eino-mcp-a2a-harness-项目设计稿]], [[项目脚手架与代码骨架]], [[harness-agent-深度解读]], [[harness-demo-源码说明]], [[agent-platform-项目骨架]], [[pdf-论文目录与首轮摘要]], [[可视化与辅助资料]]
- Added: [[rag]], [[graph-rag]], [[agentic-rag]], [[agent-harness]], [[eino]], [[a2a]], [[agent-platform]]
- Updated: [[index]], [[overview]]
- Notes: 基于 `raw/` 的 Markdown、项目骨架、PDF 标题/首页文本与可视化材料，完成第一轮可导航 Wiki 构建。

## [2026-04-23] ingest | 三篇核心综述深化

- Added: [[agentic-rag-survey]], [[graph-rag-survey]], [[agent-harness-survey]]
- Updated: [[agentic-rag]], [[graph-rag]], [[agent-harness]], [[pdf-论文目录与首轮摘要]], [[raw-source-catalog]], [[index]]
- Notes: 基于 PDF 摘要、章节起始页和关键段落抽取，补齐 Agentic RAG、GraphRAG、Agent Harness 三篇核心综述的独立来源页，并反向增强主题页。

## [2026-04-23] refactor | 初始化 outputs 产出层

- Added: `outputs/index.md`, `outputs/analyses/2026-04-23-raw-to-wiki-build-summary.md`
- Added: `outputs/slides/`, `outputs/charts/`, `outputs/tables/`, `outputs/tmp/`
- Notes: 将 `outputs/` 从空目录补成正式产出层，用于存放中间产物、分析稿、图表、演示稿和结构化导出结果。

## [2026-04-23] ingest | Hermes Agent 资料集接入

- Added: [[hermes-agent]], [[hermes-agent-资料集]], [[hermes-agent-全景图]], [[hermes-agent-核心循环]], [[hermes-agent-tool-registry]], [[hermes-agent-多provider适配]], [[hermes-agent-上下文压缩]], [[hermes-agent-消息网关]], [[hermes-agent-memory与rl训练]], [[hermes-agent-三方对比]]
- Added: [[tool-registry]], [[context-compression]], [[memory]], [[message-gateway]], [[multi-provider-adaptation]]
- Updated: [[agent-harness]], [[raw-source-catalog]], [[overview]], [[index]]
- Notes: 将 `raw/hermes-agent/` 的 8 篇结构化源码解读接入 Wiki，建立 Hermes Agent 项目页，并补齐工具注册、上下文压缩、记忆、消息网关和多 Provider 适配等概念页。
