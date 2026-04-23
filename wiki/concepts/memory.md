---
type: concept
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/07-Memory与RL训练.md
  - raw/2512.13564v2.pdf
tags:
  - concept
  - memory
  - agent
---

# Memory

## 概念定义

Memory 指 Agent 在多轮、多任务或跨会话运行中，用来保留事实、偏好、经验、工作流和可复用知识的机制。

## 为什么重要

- 没有记忆的 Agent 很难真正积累用户偏好和任务经验。
- 记忆并不等同于“存一份 transcript”；关键在于何时提取、如何路由、怎样和上下文压缩配合。
- 它同时牵涉 [[agent-harness]] 里的 `C` 和 `S` 两层能力。

## Hermes Agent 给出的实践要点

- 用 `MemoryManager` 统一协调内置与外部 provider
- 把 `prefetch`、`sync_turn`、`on_pre_compress`、`on_session_end` 这类时机做成生命周期钩子
- 让 session 过期前的 memory flush 成为消息型 Agent 的标准动作
- 把训练轨迹与 Skill 自主创建纳入“记忆与自我进化”大框架

## 当前仓库中的位置

- Memory 目前刚开始成型，Hermes 这组资料提供了很实在的工程实现视角
- 相关理论综述已在 `raw/2512.13564v2.pdf` 中出现，后续可继续拆独立来源页

## 代表性来源

- [[hermes-agent-memory与rl训练]]
- [[pdf-论文目录与首轮摘要]]
