---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/agent-platform/
tags:
  - source
  - project
  - go
  - agent-platform
---

# agent-platform 项目骨架

## 一句话摘要

`raw/agent-platform/` 是本仓库里最接近“可运行系统”的实现样例：一个基于 Go 的 Agent 平台骨架，预留了 Gateway、RAG、Search、Report、MCP、A2A、Harness 与 Store 边界。

## 这份资料回答了什么问题

- 平台设计稿如何落到实际目录和模块
- 目前项目已经具备哪些最小能力
- 哪些模块仍是 stub 或占位实现

## 核心结论

- 项目采用单仓多入口结构，当前已具备 `/v1/chat` 和本地 stub Agent 的最小演示能力。
- `docs/architecture.md`、`docs/agent-cards.md`、`docs/eval-spec.md` 表明代码结构已经开始围绕平台治理和协议边界组织。
- 当前 A2A、MCP、Harness 是“边界先行”的骨架实现，重点在模块拆分与后续扩展空间。

## 关键结构

- `cmd/`：`api`、`gateway`、`rag-agent`、`search-agent`、`report-agent`
- `internal/agents/`：各 Agent 的 graph / service / state 骨架
- `internal/mcp/`、`internal/a2a/`、`internal/harness/`：协议与治理边界
- `internal/store/`：memory / postgres / redis 三类存储抽象

## 与现有 Wiki 的关系

- 是 [[agent-platform]] 项目页的直接来源
- 验证 [[项目脚手架与代码骨架]] 与 [[eino-mcp-a2a-harness-项目设计稿]] 已部分进入实现
- 强关联 [[eino]]、[[a2a]]、[[agent-harness]]

## 未解决问题

- 仍以骨架和占位实现为主，尚非完整产品
- A2A 还没有完整 JSON-RPC / SSE 实现，Harness 评估层也偏轻量

## 来源

- `raw/agent-platform/README.md`
- `raw/agent-platform/docs/architecture.md`
- `raw/agent-platform/docs/agent-cards.md`
- `raw/agent-platform/docs/eval-spec.md`
- `raw/agent-platform/go.mod`
