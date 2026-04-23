---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/08-三方对比.md
tags:
  - source
  - hermes-agent
  - comparison
---

# Hermes Agent 三方对比

## 一句话摘要

这篇笔记把 Hermes Agent、OpenClaw、Claude Code 放在同一张桌子上，从定位、架构、主循环、工具系统、Provider 管理、消息网关和 Memory 等多个维度进行横向解剖。

## 这份资料回答了什么问题

- 三类 Agent 系统的产品哲学有何根本差异
- 为什么同样叫“AI Agent”，架构会走向完全不同的方向
- Hermes Agent 在单体全栈、Provider 管理和消息网关上的取舍意味着什么

## 核心结论

- Hermes Agent 更偏研究驱动的全栈产品，OpenClaw 更偏平台和插件驱动，Claude Code 更偏终端工程体验驱动。
- 架构差异并不是实现风格问题，而是目标用户、运行环境和组织目标的映射。
- 这种三方比较非常有助于把 Hermes 的设计选择放回更大的 Agent 产品谱系中理解。

## 与现有 Wiki 的关系

- 让 [[hermes-agent]] 不再只是孤立项目页，而是能与其它 Agent 形态进行定位比较
- 强化 [[agent-harness]] 对不同运行时工程风格的感知
- 也有助于理解为何 [[agent-platform]] 与 Hermes Agent 在工程组织上差异很大

## 来源

- `raw/hermes-agent/08-三方对比.md`
