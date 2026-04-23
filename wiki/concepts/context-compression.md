---
type: concept
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/05-上下文压缩.md
  - raw/2601.16746v2.pdf
tags:
  - concept
  - context
  - compression
---

# Context Compression

## 概念定义

Context Compression 指在不破坏任务连续性的前提下，对长对话上下文进行裁剪、摘要和重组的策略集合。

## 为什么重要

- 长任务的失败，很多时候不是模型不会做，而是上下文窗口先爆了。
- 压缩做得差，会让 Agent 忘记目标、约束、关键决策和错误上下文。
- 它是 [[agent-harness]] 中 `C` Context Manager 的核心能力之一。

## Hermes Agent 给出的实践要点

- 提前在阈值较低时触发压缩，保留后续操作余量
- 严格保护头部和尾部上下文，不做均匀裁剪
- 先裁剪旧工具输出，再做 LLM 结构化摘要
- 摘要模板按 Goal / Constraints / Progress / Decisions / Files / Next Steps / Critical Context 组织

## 与记忆的关系

- 压缩不应等于遗忘
- 在 Hermes 中，`on_pre_compress` 让 [[memory]] 系统先从即将丢弃的内容中提炼关键信息，再交给压缩器

## 代表性来源

- [[hermes-agent-上下文压缩]]
- [[pdf-论文目录与首轮摘要]]
