---
type: topic
status: active
updated: 2026-04-23
source_paths:
  - raw/06_Harness_实战架构版.md
  - raw/harness-agent-深度解读.md
  - raw/harness-agent.pdf
  - raw/harness_demo.py
tags:
  - topic
  - harness
  - agent
  - governance
---

# Agent Harness

## 主题定义

Agent Harness 指包裹模型与工具调用的运行时治理层，通常可拆成执行循环、工具注册表、上下文管理、状态存储、生命周期钩子、评估接口六部分。

## 为什么重要

- 它决定 Agent 是否可恢复、可限权、可审计、可评估。
- 复杂 Agent 的瓶颈往往不在模型本身，而在运行时设计是否合理。
- 本仓库的设计稿基本都把 Harness 视为平台化落地的必要层，而非附属功能。

## 当前仓库中的共识

- 模型不能直接操作外部世界，必须经由工具注册表和策略控制。
- 状态存储与上下文裁剪对长任务稳定性影响巨大。
- `L` 和 `V` 这类治理能力经常被忽略，但对生产系统尤为关键。

## 来自综述页的细化

- [[agent-harness-survey]] 给出了论文级别的主张：harness 是真实世界 Agent 性能的绑定约束之一。
- 该论文第 31 页的 completeness matrix 强化了本仓库已有判断：
  - `L` Lifecycle Hooks 经常缺失
  - `V` Evaluation Interface 经常缺失
- 这解释了为什么很多系统“能跑 demo”，却缺乏审计、限权、回放和稳定评估。
- 论文还把挑战显式组织成若干问题簇：
  - sandbox / security
  - evaluation / benchmarking
  - protocol standardization
  - context management
  - memory architecture
  - compute economics

## 术语说明

- 论文第 2 页摘要中的 completeness matrix 使用了更偏功能描述式的命名。
- 本 Wiki 统一采用第 3 页图示中的 `H = (E, T, C, S, L, V)`，因为它与现有中文深读稿和 demo 代码更一致。

## 六组件速记

- `E`：执行循环
- `T`：工具注册表
- `C`：上下文管理
- `S`：状态存储
- `L`：生命周期钩子
- `V`：评估接口

## 与其他主题的关系

- [[agentic-rag]] 需要 Harness 才能可靠执行多步检索
- [[a2a]] 与 [[eino]] 分别填补“跨 Agent 协作”和“单 Agent 编排”的位置
- [[agent-platform]] 是这些能力的系统化组合

## 代表性来源

- [[agent-harness-survey]]
- [[harness-实战架构版]]
- [[harness-agent-深度解读]]
- [[harness-demo-源码说明]]
- [[pdf-论文目录与首轮摘要]]
