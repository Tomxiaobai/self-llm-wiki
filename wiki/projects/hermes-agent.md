---
type: project
status: active
updated: 2026-04-23
source_paths:
  - raw/hermes-agent/
tags:
  - project
  - hermes-agent
  - python
  - agent
---

# Hermes Agent

## 它是什么

`Hermes Agent` 是 Nous Research 推出的一个偏“全栈 Agent 产品”形态的开源项目。与很多只提供库级抽象的 Agent 框架不同，它同时覆盖了 CLI、消息网关、工具系统、多 Provider 适配、上下文压缩、记忆系统、状态持久化和 RL 训练数据采集。

## 当前仓库里它以什么形式出现

当前 `raw/hermes-agent/` 不是原始源码仓库，而是一组围绕 Hermes Agent 的结构化源码解读笔记。这些笔记比 README 更细，已经把项目按关键子系统拆成 8 个主题。

## 它体现了哪些设计选择

- 用单体 `AIAgent` 承接主循环，而不是把运行时拆成很多小对象
- 工具系统采用朴素的 Registry 模式，而不是重抽象工具框架
- 多 Provider 适配统一到 OpenAI 风格消息格式上
- 用渐进式上下文压缩保证长会话稳定
- 用 `MemoryManager` 协调内置记忆、外部记忆和 RL 轨迹生成
- 内置多平台消息网关，把 IM 接入当成一等公民

## 为什么它重要

- 它给 [[agent-harness]] 提供了一个非常具体的、可读性强的工程样例
- 它让 [[tool-registry]]、[[context-compression]]、[[memory]]、[[message-gateway]] 这些概念在同一个系统里形成闭环
- 它和 [[agent-platform]] 形成很好的对照：
  - `agent-platform` 更像 Go 平台骨架
  - `hermes-agent` 更像 Python 全栈单体产品

## 当前能力与限制

- 从笔记描述看，Hermes Agent 的覆盖面非常广，接近 full-stack agent runtime
- 但它也明显承担了单体架构代价：核心文件很大、复杂度高、旋钮很多
- 当前仓库中的资料主要来自二次解读笔记，因此在做精确事实判断时，最好再回到原项目源码核对

## 相关页面

- [[hermes-agent-资料集]]
- [[hermes-agent-全景图]]
- [[hermes-agent-核心循环]]
- [[hermes-agent-tool-registry]]
- [[hermes-agent-多provider适配]]
- [[hermes-agent-上下文压缩]]
- [[hermes-agent-消息网关]]
- [[hermes-agent-memory与rl训练]]
- [[hermes-agent-三方对比]]
- [[agent-harness]]
- [[tool-registry]]
- [[context-compression]]
- [[memory]]
- [[message-gateway]]
- [[multi-provider-adaptation]]
