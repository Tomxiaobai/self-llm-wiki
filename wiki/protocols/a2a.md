---
type: protocol
status: active
updated: 2026-04-23
source_paths:
  - raw/03_A2A_协议详解.md
  - raw/07_Eino_MCP_A2A_Harness_项目设计稿.md
  - raw/agent-platform/docs/agent-cards.md
tags:
  - protocol
  - a2a
  - multi-agent
---

# A2A

## 它是什么

A2A（Agent2Agent）是一个面向 Agent 间能力发现、任务委派和长任务协作的开放协议。

## 核心抽象

- Agent Card：能力发现入口
- Task：有状态任务单元
- Streaming / Push：支持长任务进度回传
- Transport：跨服务传输与协作机制

## 在本仓库中的角色

- 用于让 Gateway 把任务转交给 RAG、Search、Report 等专业 Agent
- 让多 Agent 协作从“手写调用关系”提升为协议化边界
- 与 MCP 形成互补：MCP 偏工具接入，A2A 偏 Agent 协作

## 典型工作流

用户请求进入 Gateway 后，Gateway 根据任务类型决定：

- 直接在本地编排中处理
- 通过 MCP 调工具
- 通过 A2A 委派给专门 Agent，再汇总结果

## 局限与风险

- 仅有协议不等于有完整系统，还需要权限、策略、状态与评估
- 当前 `agent-platform` 只预留了 A2A 结构，未实现完整 JSON-RPC / SSE 能力

## 相关页面

- [[a2a-协议详解]]
- [[eino-mcp-a2a-harness-项目设计稿]]
- [[agentic-rag]]
- [[agent-platform]]
