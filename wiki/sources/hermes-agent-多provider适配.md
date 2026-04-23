---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/04-多Provider适配.md
tags:
  - source
  - hermes-agent
  - provider
---

# Hermes Agent 多 Provider 适配

## 一句话摘要

这篇笔记展示了 Hermes Agent 如何在内部统一 OpenAI 风格消息格式，再在出口层面做 Anthropic、Codex、Nous Portal 等不同 Provider 的差异化适配。

## 这份资料回答了什么问题

- 为什么统一消息格式是必须的
- Anthropic、Codex、OpenAI 兼容 API 的差异在哪里
- 凭证池化、OAuth、辅助任务路由这些“非模型本身”的能力如何影响系统稳定性

## 核心结论

- 真正困难的不是“接多个模型”，而是维持统一内部表示和稳定出口转换。
- Anthropic thinking、tool 格式和系统消息规则会迫使适配层承担大量复杂度。
- Provider 适配不只是 API 调用问题，还包括认证、限流、冷却、路由和 fallback。

## 与现有 Wiki 的关系

- 是 [[multi-provider-adaptation]] 的主要来源
- 与 [[agent-harness]] 中的上下文、工具和运行时稳定性问题高度相关
- 在 [[hermes-agent-三方对比]] 中被用于对照 Hermes、OpenClaw、Claude Code 的 Provider 管理策略

## 来源

- `raw/hermes-agent/04-多Provider适配.md`
