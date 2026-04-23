# Architecture

当前骨架采用单仓多入口模式：

- `cmd/api`：统一 API 入口
- `cmd/gateway`：Gateway 服务占位入口
- `cmd/rag-agent`：RAG Agent 占位入口
- `cmd/search-agent`：Search Agent 占位入口
- `cmd/report-agent`：Report Agent 占位入口

后续可以按服务化拆分部署，也可以先以内嵌依赖的方式本地开发。
