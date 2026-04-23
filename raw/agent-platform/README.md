# Agent Platform

一个基于 Go 的最小 Agent 平台骨架，面向 `Eino + MCP + A2A + Harness` 方向的后续扩展。

## 当前能力

- 提供 `/v1/chat` API
- 内置一个最小的 Gateway 编排链
- 内置 RAG / Search / Report 三个本地 stub Agent
- 预留 MCP、A2A、Harness、Store 扩展接口
- 可以直接 `go build ./...`

## 快速启动

```bash
go run ./cmd/api
```

然后访问：

```bash
curl -X POST http://localhost:8080/v1/chat \
  -H "Content-Type: application/json" \
  -d '{"query":"比较 GraphRAG 和 Agentic RAG"}'
```

## 目录说明

- `cmd/`：各服务入口
- `internal/agents/`：Gateway / RAG / Search / Report 的业务骨架
- `internal/mcp/`：MCP 工具接入边界
- `internal/a2a/`：A2A 协作边界
- `internal/harness/`：Hook / Policy / Eval / Trace
- `internal/store/`：memory / postgres / redis 骨架
- `configs/`：环境配置示例
- `deployments/`：本地依赖启动文件

## 备注

这个版本优先追求“骨架正确且可编译”，不是完整业务实现。
