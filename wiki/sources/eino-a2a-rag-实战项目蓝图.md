---
type: source
status: active
updated: 2026-04-23
source_paths:
  - raw/05_实战项目蓝图.md
tags:
  - source
  - rag
  - a2a
  - project
  - blueprint
---

# 实战项目蓝图：基于 Eino + A2A 的 RAG Agent 系统

## 一句话摘要

这是一份把 Eino、A2A、RAG Agent 和基础存储设施拼成完整系统蓝图的过渡文档，重点在“系统怎么跑起来”。

## 这份资料回答了什么问题

- 一个 RAG Agent 系统需要哪些角色与服务
- Corrective RAG 的在线图如何编排
- A2A 如何把 Gateway、RAG、Search、Analysis 等角色接起来

## 核心结论

- Gateway Agent 是统一入口，负责路由、汇总和治理。
- RAG Agent 不该只是“检索 + 生成”，而应具备评分、重写、补检索、幻觉检查等环节。
- 向量数据库、知识图谱、Agent 协议、编排框架必须一起考虑，才能从单点 demo 走向系统设计。

## 与现有 Wiki 的关系

- 是 [[agentic-rag]] 与 [[agent-platform]] 的重要过桥资料
- 把 [[eino]] 的单 Agent 编排与 [[a2a]] 的多 Agent 协作连起来
- 为 [[rag]]、[[graph-rag]] 提供系统级落点

## 未解决问题

- 还是设计稿，不等于已验证实现
- 评估、审批、状态恢复等治理能力需要结合 Harness 补全

## 来源

- `raw/05_实战项目蓝图.md`
