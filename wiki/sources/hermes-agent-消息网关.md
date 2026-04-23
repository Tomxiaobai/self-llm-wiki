---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/06-消息网关.md
tags:
  - source
  - hermes-agent
  - gateway
---

# Hermes Agent 消息网关

## 一句话摘要

这篇笔记把 Hermes 的消息网关解释成一个独立运行时：统一多平台适配器、SessionStore、命令分发、审批流程和会话过期策略。

## 这份资料回答了什么问题

- 多平台消息接入为什么不只是“多写几个 bot”
- 如何为不同平台、群聊、线程和用户建立正确的会话隔离
- Slash 命令、危险审批、Session reset 与 memory flush 怎样接进同一个消息主管道

## 核心结论

- Hermes 把 IM 接入当成核心运行时，而不是边缘插件。
- Session key 设计和 per-user / per-thread 隔离策略，是多平台 Agent 是否可用的基础设施。
- 会话过期前的主动 memory flush，非常适合长期运行的消息型 Agent。

## 与现有 Wiki 的关系

- 是 [[message-gateway]] 的主要来源
- 让 [[hermes-agent]] 与纯 CLI Agent 形成清晰对照
- 丰富 [[memory]] 在会话生命周期管理中的角色

## 来源

- `raw/hermes-agent/06-消息网关.md`
