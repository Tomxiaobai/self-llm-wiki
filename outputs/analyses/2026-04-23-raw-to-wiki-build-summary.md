# Raw to Wiki Build Summary

日期：2026-04-23

## 本次产出目标

把 `raw/` 中的资料首轮构建为可导航的 LLM Wiki，并补齐关键综述页的深度来源页。

## 已完成的主要产物

### 1. 仓库规范

- 根级 `CLAUDE.md`
- `wiki/index.md`
- `wiki/overview.md`
- `wiki/log.md`

### 2. 首轮来源页

- Eino
- A2A
- Agent 设计模式
- Eino + A2A 实战蓝图
- Harness 实战架构版
- Eino + MCP + A2A + Harness 项目设计稿
- 项目脚手架与代码骨架
- Harness 深度解读
- Harness Demo 源码说明
- `agent-platform` 项目骨架
- PDF 论文目录页
- 可视化资料目录页

### 3. 主题 / 中层页面

- `rag`
- `graph-rag`
- `agentic-rag`
- `agent-harness`
- `eino`
- `a2a`
- `agent-platform`

### 4. 第二轮深化页

- `agentic-rag-survey`
- `graph-rag-survey`
- `agent-harness-survey`

## 当前仓库状态

- `raw/` 已被映射到 `wiki/` 的主要知识骨架
- `outputs/` 已初始化为正式产出层
- 当前最值得继续深化的是：
  - `Memory in the Age of AI Agents`
  - `SWE-Pruner`
  - `MCP`

## 建议下一步

1. 为 Memory 综述拆独立页，并补 `wiki/concepts/memory.md`
2. 为 SWE-Pruner 拆独立页，并补 `context pruning` / `coding agent context` 相关概念页
3. 建立 `wiki/protocols/mcp.md`
4. 视需要在 `outputs/slides/` 下生成对外可展示的 Marp 演示稿
