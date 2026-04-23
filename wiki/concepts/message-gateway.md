---
type: concept
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/06-消息网关.md
tags:
  - concept
  - gateway
  - messaging
---

# Message Gateway

## 概念定义

Message Gateway 指把 Telegram、Discord、Slack、飞书等不同消息平台统一接入同一 Agent 能力层的运行时中枢。

## 为什么重要

- 对消息型 Agent 来说，真正困难的不只是 API 对接，而是会话隔离、命令分发、危险操作审批和断线恢复。
- 如果 Gateway 做不好，多平台接入很快就会把上下文与权限逻辑搅乱。

## Hermes Agent 给出的实践要点

- 用统一 `MessageEvent` 和 adapter 抹平平台差异
- 用 session key 明确区分 DM、群聊、线程、per-user / per-thread 场景
- 把 slash commands、approve/deny、queue、model override 都纳入一条统一分发管道
- 在 session reset 前先做 memory flush，避免长期运行中的知识损失

## 与其它能力的关系

- 与 [[memory]] 强关联，因为会话生命周期直接决定记忆什么时候写入
- 与 [[multi-provider-adaptation]] 不同，Gateway 关心的是消息入口，不是模型出口
- 对纯 CLI Agent 不是必需，但对个人助理型或企业消息型 Agent 极其关键

## 代表性来源

- [[hermes-agent-消息网关]]
