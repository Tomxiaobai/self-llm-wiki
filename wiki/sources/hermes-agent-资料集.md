---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/
tags:
  - source
  - hermes-agent
  - catalog
---

# Hermes Agent 资料集

## 一句话摘要

`raw/hermes-agent/` 是一组围绕 Hermes Agent 的专题式源码解读笔记，覆盖项目全景、主循环、工具系统、多 Provider、上下文压缩、消息网关、记忆系统与三方对比。

## 这组资料回答了什么问题

- Hermes Agent 为什么更像“完整产品”而不是“工具库”
- 一个全栈 Agent 系统在运行时到底由哪些子系统构成
- Hermes Agent 在工具、Provider、压缩、记忆、消息网关上的工程取舍是什么
- 它和 OpenClaw、Claude Code 这两类系统相比差异在哪里

## 子文档映射

- [[hermes-agent-全景图]]：项目全景、规模和总体分层
- [[hermes-agent-核心循环]]：`AIAgent`、预算控制、并行工具与主循环
- [[hermes-agent-tool-registry]]：工具注册与分发
- [[hermes-agent-多provider适配]]：统一消息格式与 Provider 出口适配
- [[hermes-agent-上下文压缩]]：长对话压缩策略
- [[hermes-agent-消息网关]]：多平台消息接入与会话管理
- [[hermes-agent-memory与rl训练]]：记忆、训练数据与 Skill 自主创建
- [[hermes-agent-三方对比]]：与 OpenClaw、Claude Code 的横向比较

## 与现有 Wiki 的关系

- 是 [[hermes-agent]] 项目页的直接来源
- 丰富了 [[agent-harness]] 的实践例子
- 推动建立了 [[tool-registry]]、[[context-compression]]、[[memory]]、[[message-gateway]]、[[multi-provider-adaptation]] 五个概念页

## 未解决问题

- 当前资料是二次源码解读，不是原始仓库快照
- 对某些实现细节的精确判断，后续最好结合 Hermes Agent 原仓库复核

## 来源

- `raw/hermes-agent/`
