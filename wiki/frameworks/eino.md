---
type: framework
status: active
updated: 2026-04-23
source_paths:
  - raw/02_Eino_框架指南.md
  - raw/04_Agent_设计模式.md
tags:
  - framework
  - eino
  - go
---

# Eino

## 它是什么

Eino 是一个面向 Go 的 LLM 应用开发框架，核心特点是标准化组件接口与多种编排范式。

## 核心抽象

- Components：`ChatModel`、`Retriever`、`Embedding`、`Tool`、`Lambda` 等
- Orchestration：`Chain`、`Graph`、`Workflow`
- State / Call Options：用于把运行时控制逻辑注入编排

## 在本仓库中的角色

- 作为单 Agent 内部编排引擎，支撑 RAG、ReAct、Corrective RAG 等工作流
- 为 [[agent-platform]] 的 Gateway 和各子 Agent 提供实现思路
- 与 [[a2a]] 构成“内部怎么跑 / 外部怎么协作”的分工

## 适用场景

- Go 技术栈中的 RAG、Tool Use、Multi-Agent 原型与服务
- 需要把图结构、重试、分支、状态管理纳入执行逻辑的 Agent 系统

## 局限与风险

- 资料中更强调框架能力，较少覆盖生产治理细节
- 真正的平台化仍需引入 [[agent-harness]]、[[a2a]]、存储与评估层

## 相关页面

- [[eino-框架指南]]
- [[agent-设计模式]]
- [[agentic-rag]]
- [[agent-platform]]
