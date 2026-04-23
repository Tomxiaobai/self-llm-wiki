# Wiki Index

这是 `wiki/` 的内容目录。Claude 在执行 ingest、query、lint 之前，应先读这里，再决定要深入哪些页面。

## Core

- [[overview]] - 顶层主题地图、当前关注方向、首批 ingest 优先级
- [[log]] - Wiki 的时间演化记录
- [[raw-source-catalog]] - `raw/` 到 Wiki 的映射目录

## Topics

- [[rag]] - 本仓库对 RAG 主线的总览页
- [[graph-rag]] - 图结构 / 知识图谱增强检索的主题页
- [[agentic-rag]] - 把检索纳入 Agent 执行循环的主题页
- [[agent-harness]] - Agent 运行时治理层主题页

## Frameworks, Protocols, Projects

- [[eino]] - Go Agent 编排框架页
- [[a2a]] - Agent-to-Agent 协议页
- [[agent-platform]] - Go Agent 平台骨架项目页
- [[hermes-agent]] - Python 全栈 Agent 项目页

## Concepts

- [[tool-registry]] - 工具注册、发现、可用性与分发
- [[context-compression]] - 长对话压缩与上下文保真
- [[memory]] - 跨轮与跨会话记忆
- [[message-gateway]] - 多平台消息接入与会话隔离
- [[multi-provider-adaptation]] - 多模型 Provider 出口适配

## Sources

- [[eino-框架指南]]
- [[a2a-协议详解]]
- [[agent-设计模式]]
- [[eino-a2a-rag-实战项目蓝图]]
- [[harness-实战架构版]]
- [[eino-mcp-a2a-harness-项目设计稿]]
- [[项目脚手架与代码骨架]]
- [[harness-agent-深度解读]]
- [[harness-demo-源码说明]]
- [[agent-platform-项目骨架]]
- [[hermes-agent-资料集]]
- [[hermes-agent-全景图]]
- [[hermes-agent-核心循环]]
- [[hermes-agent-tool-registry]]
- [[hermes-agent-多provider适配]]
- [[hermes-agent-上下文压缩]]
- [[hermes-agent-消息网关]]
- [[hermes-agent-memory与rl训练]]
- [[hermes-agent-三方对比]]
- [[agentic-rag-survey]]
- [[graph-rag-survey]]
- [[agent-harness-survey]]
- [[pdf-论文目录与首轮摘要]]
- [[可视化与辅助资料]]

## Current Status

- 当前 Wiki 已完成两轮构建，已有主题页、项目页、来源页，以及 3 篇关键综述的独立论文页。
- PDF 不再只是目录入口；`Agentic RAG`、`GraphRAG`、`Agent Harness` 已具备更细粒度来源页。

## First Ingest Batch

第一轮已优先吸收以下资料：

- `raw/02_Eino_框架指南.md`
- `raw/03_A2A_协议详解.md`
- `raw/04_Agent_设计模式.md`
- `raw/05_实战项目蓝图.md`
- `raw/06_Harness_实战架构版.md`
- `raw/agent-platform/`

## Next Priority

- 补 `MCP` 相关主题页
- 为 `Memory` 理论综述和 `SWE-Pruner` 拆更细的来源页
- 继续把问答沉淀到 `wiki/analyses/`
