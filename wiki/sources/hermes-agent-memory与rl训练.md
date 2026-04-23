---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/07-Memory与RL训练.md
tags:
  - source
  - hermes-agent
  - memory
  - rl
---

# Hermes Agent Memory 与 RL 训练

## 一句话摘要

这篇笔记展示了 Hermes 如何把记忆、上下文压缩、外部 memory provider、训练轨迹和 Skill 自主创建连成一条完整的“自我积累”链路。

## 这份资料回答了什么问题

- `MemoryManager` 如何统一协调内置和外部记忆后端
- 生命周期钩子如何影响记忆写入、召回和压缩前抢救
- 为什么 RL 训练数据采集和 Agent 运行时会被放进同一个系统

## 核心结论

- 记忆系统的关键不只是存储后端，而是统一调度与生命周期注入。
- `on_pre_compress` 这种钩子让上下文压缩与长期记忆不再互相冲突。
- 把轨迹压缩和 Skill 创建纳入同一体系，体现出 Hermes 的研究驱动与自我进化导向。

## 与现有 Wiki 的关系

- 是 [[memory]] 的重要来源
- 补足了当前仓库里 `Memory` 这条原本还偏空缺的主线
- 与 [[agent-harness]] 中的 `S` State Store 和 `C` Context Manager 紧密相关

## 来源

- `raw/hermes-agent/07-Memory与RL训练.md`
