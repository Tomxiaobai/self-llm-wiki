---
type: project
status: active
updated: 2026-04-23
source_paths:
  - raw/agent-platform/
  - raw/07_Eino_MCP_A2A_Harness_项目设计稿.md
  - raw/08_项目脚手架与代码骨架.md
tags:
  - project
  - agent-platform
  - go
---

# agent-platform

## 它是什么

`agent-platform` 是一个基于 Go 的最小 Agent 平台骨架，用来承接本仓库关于 Eino、MCP、A2A 与 Harness 的设计思路。

## 当前已具备的结构

- `cmd/api` 提供统一 API 入口
- `cmd/gateway`、`cmd/rag-agent`、`cmd/search-agent`、`cmd/report-agent` 预留了多角色入口
- `internal/agents/`、`internal/mcp/`、`internal/a2a/`、`internal/harness/`、`internal/store/` 完成了边界切分
- `docs/` 补充了 architecture、agent cards、eval spec 等说明

## 它体现了哪些设计选择

- 单仓多入口，而不是先拆成多个仓库
- 先把边界立稳，再逐步填充业务实现
- 将治理、协议、存储与 Agent 逻辑并列，而不是事后缝补

## 当前能力与限制

- 已能作为“结构正确、可编译”的最小骨架
- 当前 RAG / Search / Report 更偏 stub，占位多于真实业务
- A2A、MCP、Harness 仍是边界先行，尚未成为成熟运行时

## 为什么它重要

- 它是本仓库里最接近“把理论落成系统”的样例
- 可以用来对照设计稿，判断哪些理念已经进入代码
- 未来如果把 Wiki 继续做成工程化平台，这里很可能是直接起点

## 相关页面

- [[agent-platform-项目骨架]]
- [[项目脚手架与代码骨架]]
- [[eino-mcp-a2a-harness-项目设计稿]]
- [[eino]]
- [[a2a]]
- [[agent-harness]]
