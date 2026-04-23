---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/01-全景图.md
tags:
  - source
  - hermes-agent
  - architecture
---

# Hermes Agent 全景图

## 一句话摘要

这篇笔记给出了 Hermes Agent 的整体架构、代码规模和系统边界，是理解该项目“为什么像一个完整产品”的最好入口。

## 这份资料回答了什么问题

- Hermes Agent 的定位是什么
- 它的核心子系统有哪些
- 为什么它和普通 Agent library 的工程形态不同

## 核心结论

- Hermes Agent 覆盖 CLI、Gateway、Agent 核心循环、工具编排、Provider 适配、状态与记忆、RL 训练采样等多个层面。
- 它的双入口设计很关键：终端和消息平台共享一套 Agent 能力，但会话状态保持隔离。
- “训练数据采集内置”是它与很多其它框架的显著区别，也解释了其研究驱动色彩。

## 与现有 Wiki 的关系

- 是 [[hermes-agent]] 项目页的主入口来源
- 反向支撑 [[agent-harness]] 对“全栈 Agent 运行时”的理解
- 为 [[message-gateway]] 和 [[memory]] 提供系统级位置

## 来源

- `raw/hermes-agent/01-全景图.md`
