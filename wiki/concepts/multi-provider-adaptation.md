---
type: concept
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/04-多Provider适配.md
tags:
  - concept
  - provider
  - adaptation
---

# Multi-Provider Adaptation

## 概念定义

Multi-Provider Adaptation 指在同一个 Agent 运行时内部维持统一消息表示和调用流程，再针对不同模型供应商的协议、认证、thinking、tool 格式和限额规则做出口适配。

## 为什么重要

- 现实里很少有两个 Provider 的 API 真的完全一样
- 真正的复杂度常常来自 thinking 语义、tool block、OAuth、冷却、凭证池化和 fallback，而不是“换个 base_url”
- 它直接影响系统的稳定性、成本和可恢复性

## Hermes Agent 给出的实践要点

- 内部统一用 OpenAI 风格消息格式
- 在出口适配 Anthropic、Codex 和其它 OpenAI 兼容 Provider
- 通过 credential pool、auxiliary routing 和 fallback 链处理资源耗尽与辅助任务路由
- 把认证与模型能力映射视作适配层一部分，而不是外围脚本

## 与其它能力的关系

- 与 [[tool-registry]] 一起决定 Agent 能做什么、怎么做
- 与 [[context-compression]]、[[memory]] 共同影响不同模型窗口和成本下的运行方式
- 与 MCP、A2A 不同，它不解决协议协作，而是解决模型出口兼容

## 代表性来源

- [[hermes-agent-多provider适配]]
