---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/harness_demo.py
tags:
  - source
  - harness
  - demo
  - python
---

# Harness Demo 源码说明

## 一句话摘要

`raw/harness_demo.py` 用一份简化 Python 代码，把 Harness 的六个组件直接实现成可阅读的最小示例。

## 这份资料回答了什么问题

- 六组件在代码层面长什么样
- 状态存储、工具注册表、上下文管理、生命周期钩子、评估接口如何最小实现
- “Harness 是运行时而不是 Prompt”的说法如何落到代码上

## 核心结论

- 这份 demo 不是完整框架，但很好地展示了 `S`、`T`、`C`、`L`、`V` 的职责边界。
- 代码层的组件化有助于把治理逻辑从单轮推理中拆出来。
- 它更适合当教学样例，而不是直接作为生产实现。

## 与现有 Wiki 的关系

- 是 [[agent-harness]] 的代码型佐证
- 与 [[harness-实战架构版]] 和 [[harness-agent-深度解读]] 互相印证
- 可作为未来 `agent-platform` 治理层的教学参照

## 未解决问题

- 依赖 OpenAI 兼容 API，但未绑定本仓库其他模块
- 只演示理念，没有完整安全隔离、协议协作和持久化存储

## 来源

- `raw/harness_demo.py`
