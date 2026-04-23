---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/05-上下文压缩.md
tags:
  - source
  - hermes-agent
  - context
---

# Hermes Agent 上下文压缩

## 一句话摘要

这篇笔记详细拆解了 Hermes Agent 的上下文压缩器：如何在不丢关键信息的前提下，靠头尾保护、工具输出裁剪和结构化摘要延长对话寿命。

## 这份资料回答了什么问题

- 什么时候开始压缩上下文
- 头部、尾部和中间区为什么要区别处理
- 为什么压缩必须用结构化摘要，而不是随便总结一段话

## 核心结论

- Hermes 把压缩触发阈值设得较早，体现出“保留余量优先”的策略。
- 先裁剪工具输出、再调用 LLM 摘要，是成本与保真之间的渐进式折中。
- 结构化摘要模板能比自由摘要更稳定地保留目标、约束、进度、决策和关键上下文。

## 与现有 Wiki 的关系

- 是 [[context-compression]] 的核心来源
- 丰富 [[agent-harness]] 中 `C` Context Manager 的实践部分
- 与 [[memory]] 存在直接联动，因为压缩前的记忆抢救会影响最终保留的信息

## 来源

- `raw/hermes-agent/05-上下文压缩.md`
