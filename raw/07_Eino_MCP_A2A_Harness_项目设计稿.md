# 实战项目设计稿：基于 Eino + MCP + A2A + Harness 的 Agent 平台

> 目标：把前面的 `Eino`、`MCP`、`A2A` 和 `Harness` 四条线真正收束成一个可实现、可迭代、可上线的项目设计稿。

---

## 1. 项目定位

我们设计一个 **企业知识问答 + 研究分析 + 报告生成** 的 Agent 平台。

它要解决的不是“单次问答”，而是下面这类任务：

- 用户提问一个复杂问题
- 系统先查内部知识库
- 证据不足时补 Web 检索或调用专用 Agent
- 需要时调用结构化工具做解析、重排、引用整理
- 最后给出带出处、可追踪、可恢复的答案或报告

这类系统刚好适合下面的技术分工：

- `Eino`：做单个 Agent 内部的编排和执行循环
- `MCP`：把工具、资源、数据源标准化接进来
- `A2A`：把外部专业 Agent 接进来做协作
- `Harness`：为整个系统提供治理、状态、策略、安全和评估

---

## 2. 为什么这四个要一起上

### 2.1 只用 Eino 不够

只用 Eino，我们可以很快搭出：

- ReAct Agent
- Corrective RAG
- Graph 工作流

但如果任务开始变长、工具变多、状态要跨轮保存，就会遇到：

- 工具太多难治理
- 多 Agent 协作缺协议
- 失败后难恢复
- 没有统一的审批、审计和评估

### 2.2 加上 MCP 和 A2A 后，职责更清楚

```text
Eino 负责“内部怎么跑”
MCP 负责“工具怎么接”
A2A 负责“Agent 怎么协作”
Harness 负责“整个系统怎么受控地运行”
```

这也是这份设计稿的核心原则。

---

## 3. 项目目标

### 3.1 业务目标

- 支持企业内部知识库问答
- 支持复杂问题的多步分析
- 支持外部检索补充证据
- 支持生成结构化研究报告
- 支持多 Agent 专业分工

### 3.2 工程目标

- Agent 执行过程可恢复
- 工具调用可治理、可限权
- 支持人工审批节点
- 输出带引用和证据链
- 有统一 trace / 评估 / 回放能力

### 3.3 非目标

- 第一版不追求全自动闭环执行
- 第一版不追求开放任意第三方 Agent 接入
- 第一版不追求无限工具扩展

我们优先要的是：

```text
稳、可解释、可恢复、可扩展
```

---

## 4. 总体架构

```mermaid
flowchart TD
    U["用户 / Web / CLI"] --> API["API Gateway"]
    API --> GW["Gateway Agent (Eino Graph)"]

    GW --> POL["Policy / Approval"]
    GW --> STATE["State Store"]
    GW --> OBS["Tracing / Eval / Logs"]

    GW --> MCPHUB["MCP Tool Hub"]
    GW --> A2A["A2A Client"]

    MCPHUB --> RET["Retriever MCP Server"]
    MCPHUB --> RERANK["Reranker MCP Server"]
    MCPHUB --> PARSER["Doc Parser MCP Server"]
    MCPHUB --> SQL["SQL / Graph MCP Server"]

    A2A --> RAG["RAG Agent"]
    A2A --> SEARCH["Search Agent"]
    A2A --> REPORT["Report Agent"]

    RAG --> VDB["Milvus / ES"]
    SEARCH --> WEB["Web Search / Browser"]
    REPORT --> OBJ["Object Storage / Templates"]

    STATE --> PG["PostgreSQL"]
    STATE --> REDIS["Redis"]
    OBS --> LF["Langfuse / OpenTelemetry"]
```

---

## 5. 分层设计

### 5.1 接入层

负责：

- Web/CLI/API 接入
- 用户身份识别
- 请求限流
- 将用户请求转换成统一任务

组件建议：

- `cmd/api`
- Gin/Fiber/Chi 任一 Go Web 框架
- JWT / Bearer 鉴权

### 5.2 Orchestrator 层

核心是 `Gateway Agent`

职责：

- 任务分类
- 路由到合适的子 Agent
- 决定是走 MCP 工具还是 A2A 委派
- 聚合结果
- 维护执行状态机

实现建议：

- 用 `Eino Graph` 编排
- 用本地 `State` 管运行态
- 用持久化 `TaskStore` 管跨轮/跨请求状态

### 5.3 Tool 层

通过 MCP 接入工具和资源。

建议第一版只接 4 类：

- 检索工具：向量检索、BM25、Hybrid Search
- 处理工具：Rerank、Chunk Merge、Citation Builder
- 数据工具：SQL / Graph 查询
- 文档工具：PDF 解析、Markdown 清洗、元数据抽取

原则：

```text
先少而精
不要把所有能力都做成工具暴露给模型
```

### 5.4 Agent 协作层

通过 A2A 接入外部专业 Agent。

建议第一版的角色：

- `RAG Agent`：内部知识库问答
- `Search Agent`：外部检索和网页证据采集
- `Report Agent`：结构化摘要和报告生成

Gateway 不把所有能力都自己做完，而是负责：

- 判断哪个 Agent 更适合
- 组织子任务
- 统一汇总

### 5.5 Harness 治理层

这是最容易被忽略，但最关键的一层。

它不是单独一个服务，而是一组横切能力：

- 状态管理
- 生命周期拦截
- 策略控制
- 评估与回放
- 审计与可观测性

---

## 6. 用 Harness 六组件映射整个项目

| Harness 组件 | 在本项目里的落点 | 推荐实现 |
|-------------|------------------|---------|
| `E` Execution Loop | Gateway Agent 的 Graph、子 Agent 内部 Graph、终止/重试逻辑 | Eino `compose.Graph` |
| `T` Tool Registry | MCP Tool Hub、工具 schema、权限和工具白名单 | MCP client + registry wrapper |
| `C` Context Manager | 用户问题、历史摘要、检索证据、AGENTS/规则注入 | Prompt Builder + state summarizer |
| `S` State Store | task、checkpoint、artifact、citation、运行阶段 | PostgreSQL + Redis |
| `L` Lifecycle Hooks | tool 调用前审批、长任务超时、敏感工具拦截、审计 | hook middleware / policy engine |
| `V` Evaluation | trace、成本、成功率、引用完整率、失败原因 | Langfuse + 自定义评估表 |

### 6.1 这张表最重要的意思

不是说“我们也有日志，所以有 Harness”。

而是要落实成代码和模块：

- `E` 体现在 Graph 和状态机
- `T` 体现在工具注册和权限控制
- `C` 体现在上下文装配器
- `S` 体现在任务表和 checkpoint
- `L` 体现在统一 hook
- `V` 体现在 trace + evaluator

---

## 7. 推荐的业务场景

我们把第一版场景收敛成一个最有代表性的任务：

### 7.1 场景：研究型问答

用户提问：

> 请比较 GraphRAG、Agentic RAG 和传统 RAG，结合我们内部文档和网上资料，输出一份带引用的分析报告。

系统目标：

1. 先查内部知识库
2. 不够时查外部 Web
3. 对结果做重排和证据整合
4. 用报告模板生成输出
5. 返回引用和证据来源

这个场景天然需要：

- RAG
- 多步决策
- 外部 Agent 协作
- 工具治理
- 状态持久化
- 引用校验

---

## 8. 一次完整请求的执行流

```mermaid
sequenceDiagram
    participant User
    participant API as API Gateway
    participant GW as Gateway Agent
    participant MCP as MCP Tool Hub
    participant RAG as RAG Agent
    participant Search as Search Agent
    participant Report as Report Agent
    participant Store as State Store

    User->>API: 提问
    API->>GW: 创建任务
    GW->>Store: 写入 submitted
    GW->>MCP: 意图分类 / query normalize
    GW->>RAG: A2A 调用内部知识检索
    RAG-->>GW: 返回内部证据
    GW->>MCP: 相关性评分 / 引用候选整理
    alt 内部证据不足
        GW->>Search: A2A 调用外部搜索
        Search-->>GW: 返回网页证据
    end
    GW->>MCP: rerank / merge / citation build
    GW->>Report: A2A 调用报告生成
    Report-->>GW: 返回初稿
    GW->>MCP: answer groundedness check
    GW->>Store: 写入 artifacts / completed
    GW-->>API: 输出报告和引用
    API-->>User: 返回结果
```

---

## 9. 模块设计

### 9.1 Gateway Agent

这是平台的大脑，不直接包办所有专业能力。

它负责：

- 问题分类
- 是否要调用内部 RAG
- 是否要升级成“研究任务”
- 是否要转发给 Search Agent / Report Agent
- 控制结束条件和重试策略

### 9.2 RAG Agent

职责：

- 对接向量库和检索器
- 支持 query rewrite
- 返回候选文档和原始来源
- 返回置信度和命中信息

注意：

RAG Agent 不应该直接负责“最终漂亮回答”，它更像证据提供者。

### 9.3 Search Agent

职责：

- Web 搜索
- 网页抓取
- 页面清洗
- 来源可信度评分

建议：

- 外部检索和内部知识检索分开，避免一个 Agent 既做内网知识又做外网抓取，职责太杂

### 9.4 Report Agent

职责：

- 将证据转为结构化报告
- 控制输出格式
- 生成摘要、章节、引用列表

这样做的好处是：

- 把“证据收集”和“表达输出”分离
- 更容易替换不同写作策略或模板

---

## 10. Eino 内部编排建议

### 10.1 Gateway Graph

推荐节点：

```text
START
  -> normalize_query
  -> classify_intent
  -> call_rag_agent
  -> judge_need_search
      -> call_search_agent
  -> merge_evidence
  -> call_report_agent
  -> verify_grounding
  -> END
```

### 10.2 关键思想

- `A2A 调用` 也是 Graph 中的节点
- `MCP 工具调用` 也是 Graph 中的节点
- 不要把所有逻辑塞进一个大 prompt

### 10.3 示例伪代码

```go
type GatewayState struct {
    TaskID          string
    UserQuery       string
    NormalizedQuery string
    Intent          string
    NeedSearch      bool
    InternalDocs    []*schema.Document
    ExternalDocs    []*schema.Document
    Evidence        []*schema.Document
    DraftReport     string
    FinalReport     string
    Citations       []Citation
    RetryCount      int
}

func BuildGatewayGraph(ctx context.Context) (*compose.CompiledGraph, error) {
    g := compose.NewGraph[string, string](
        compose.WithGenLocalState(func(ctx context.Context) *GatewayState {
            return &GatewayState{}
        }),
    )

    g.AddLambdaNode("normalize_query", normalizeQuery)
    g.AddLambdaNode("classify_intent", classifyIntent)
    g.AddLambdaNode("call_rag_agent", callRAGAgent)
    g.AddLambdaNode("judge_need_search", judgeNeedSearch)
    g.AddLambdaNode("call_search_agent", callSearchAgent)
    g.AddLambdaNode("merge_evidence", mergeEvidence)
    g.AddLambdaNode("call_report_agent", callReportAgent)
    g.AddLambdaNode("verify_grounding", verifyGrounding)

    g.AddEdge(compose.START, "normalize_query")
    g.AddEdge("normalize_query", "classify_intent")
    g.AddEdge("classify_intent", "call_rag_agent")
    g.AddEdge("call_rag_agent", "judge_need_search")
    g.AddBranch("judge_need_search", needSearchBranch)
    g.AddEdge("call_search_agent", "merge_evidence")
    g.AddEdge("merge_evidence", "call_report_agent")
    g.AddEdge("call_report_agent", "verify_grounding")
    g.AddEdge("verify_grounding", compose.END)

    return g.Compile(ctx, compose.WithMaxRunSteps(12))
}
```

---

## 11. MCP 设计建议

### 11.1 不要把 MCP 当成“万物总线”

第一版建议只接下面 5 个 MCP Server：

1. `retriever-server`
2. `reranker-server`
3. `doc-parser-server`
4. `citation-server`
5. `sql-or-graph-server`

### 11.2 Tool Registry 设计

建议在业务层再包一层 registry，不要让模型直接看见所有 MCP tool：

```go
type ToolScope string

const (
    ScopeRetrieve ToolScope = "retrieve"
    ScopeCite     ToolScope = "cite"
    ScopeData     ToolScope = "data"
)

type RegisteredTool struct {
    Name        string
    Scope       ToolScope
    Description string
    RiskLevel   string
}

type ToolRegistry interface {
    ListByIntent(intent string) []RegisteredTool
    ValidateCall(tool string, args map[string]any) error
}
```

### 11.3 关键原则

- 工具按任务暴露，不按系统全量暴露
- schema 校验前置
- 高风险工具必须带 hook
- 检索、重排、引用构建优先走确定性工具

---

## 12. A2A 设计建议

### 12.1 Agent Card 规划

第一版至少有 3 个 Agent Card：

- `rag-agent`
- `search-agent`
- `report-agent`

Gateway 在启动时缓存它们的 card：

- 名称
- 描述
- skills
- 输入输出模式
- 是否支持 streaming
- 鉴权方式

### 12.2 A2A 适合什么，不适合什么

适合：

- 子任务委派
- 长任务执行
- 远端专业能力协作

不适合：

- 高频低延迟的本地小工具调用
- 简单 schema 转换
- 单步检索器访问

换句话说：

```text
MCP 负责“本地工具边界”
A2A 负责“远端 Agent 边界”
```

### 12.3 A2A Client 封装建议

```go
type A2AClient interface {
    Discover(ctx context.Context, baseURL string) (*AgentCard, error)
    SendTask(ctx context.Context, baseURL string, msg Message) (*Task, error)
    StreamTask(ctx context.Context, baseURL string, msg Message) (<-chan StreamEvent, error)
}
```

建议在业务侧增加：

- 超时控制
- 幂等 key
- retry policy
- task status 映射

---

## 13. 状态设计

### 13.1 为什么必须单独设计 State

这个系统不是一次 prompt 就结束。

它需要保存：

- 当前任务阶段
- 已检索证据
- 已调用过哪些 Agent
- 当前草稿
- 审批状态
- 失败原因

### 13.2 状态分层

#### 运行态

放 Redis 或内存：

- 流式中间结果
- 短时缓存
- 当前执行节点

#### 持久态

放 PostgreSQL：

- task
- context
- artifact
- citation
- audit log
- eval result

### 13.3 推荐表设计

```text
tasks
  id, context_id, state, intent, created_at, updated_at

task_steps
  id, task_id, step_name, step_state, input_json, output_json, error_text

artifacts
  id, task_id, type, content_uri, metadata_json

citations
  id, task_id, source_type, source_uri, chunk_id, span_text

approvals
  id, task_id, action, status, reviewer, reviewed_at

eval_runs
  id, task_id, grounded_score, citation_score, latency_ms, token_usage, passed
```

---

## 14. Lifecycle Hooks 设计

这一层最容易在 Demo 中被跳过，但上生产时必须补齐。

### 14.1 Hook 插入点

建议统一支持：

- `before_task_start`
- `before_tool_call`
- `after_tool_call`
- `before_a2a_dispatch`
- `after_a2a_dispatch`
- `before_finalize`
- `on_error`

### 14.2 第一版必须做的 Hook

- 敏感工具白名单检查
- 外部 Agent 域名白名单检查
- 任务超时和最大步数限制
- 高成本调用限流
- 输出前引用完整性校验

### 14.3 示例接口

```go
type HookContext struct {
    TaskID   string
    NodeName string
    Payload  map[string]any
}

type LifecycleHook interface {
    Before(ctx context.Context, hc HookContext) error
    After(ctx context.Context, hc HookContext, result any) error
    OnError(ctx context.Context, hc HookContext, err error) error
}
```

---

## 15. Evaluation 设计

### 15.1 第一版别只看“答得像不像”

至少要看 5 类指标：

| 维度 | 指标 |
|------|------|
| 结果质量 | groundedness、citation completeness、answer relevance |
| 过程效率 | total steps、latency、token |
| 工具质量 | tool success rate、tool timeout rate |
| 协作质量 | A2A success rate、delegation fallback rate |
| 运维质量 | manual takeover rate、failure distribution |

### 15.2 推荐最小评估链路

```text
每个任务结束
  -> 保存 trace
  -> 保存关键 artifact
  -> 运行 groundedness evaluator
  -> 运行 citation completeness evaluator
  -> 写入 eval_runs
```

### 15.3 第一版可接受的做法

- trace：Langfuse / OpenTelemetry
- 自动评估：rule-based + LLM-as-judge 混合
- 人工抽检：每周采样若干任务做复盘

---

## 16. 目录结构建议

```text
agent-platform/
├── cmd/
│   ├── api/                 # API Gateway
│   ├── gateway/             # Gateway Agent
│   ├── rag-agent/           # 内部知识检索 Agent
│   ├── search-agent/        # 外部搜索 Agent
│   └── report-agent/        # 报告生成 Agent
├── internal/
│   ├── agents/
│   │   ├── gateway/
│   │   ├── rag/
│   │   ├── search/
│   │   └── report/
│   ├── orchestrator/
│   │   ├── graph/
│   │   ├── state/
│   │   └── branch/
│   ├── mcp/
│   │   ├── client/
│   │   ├── registry/
│   │   └── tools/
│   ├── a2a/
│   │   ├── client/
│   │   ├── card/
│   │   └── transport/
│   ├── harness/
│   │   ├── hooks/
│   │   ├── policy/
│   │   ├── approval/
│   │   └── eval/
│   ├── store/
│   │   ├── postgres/
│   │   ├── redis/
│   │   └── objectstore/
│   ├── api/
│   │   ├── dto/
│   │   └── handler/
│   └── shared/
│       ├── schema/
│       ├── config/
│       └── tracing/
├── deployments/
│   ├── docker-compose.yml
│   └── k8s/
└── docs/
    ├── agent-cards/
    ├── prompt-policies/
    └── eval-specs/
```

---

## 17. 关键配置建议

### 17.1 环境变量

```bash
# Eino / LLM
CHAT_MODEL=ep-your-chat-model
EMBED_MODEL=ep-your-embed-model
ARK_API_KEY=your_api_key

# Vector DB
MILVUS_ADDRESS=localhost:19530
COLLECTION_NAME=knowledge_base

# State
POSTGRES_DSN=postgres://user:pass@localhost:5432/agent_platform
REDIS_ADDR=localhost:6379

# A2A Agents
RAG_AGENT_URL=http://localhost:8081
SEARCH_AGENT_URL=http://localhost:8082
REPORT_AGENT_URL=http://localhost:8083

# Observability
LANGFUSE_HOST=http://localhost:3000
LANGFUSE_PUBLIC_KEY=public_xxx
LANGFUSE_SECRET_KEY=secret_xxx
```

### 17.2 配置文件

建议补一个 `configs/dev.yaml`：

- 模型选择
- 各 Agent URL
- 工具白名单
- 超时策略
- 是否开启人工审批

---

## 18. 本地启动顺序

```text
1. 启动 PostgreSQL / Redis / Milvus / Langfuse
2. 启动 MCP Server 组
3. 启动 RAG Agent / Search Agent / Report Agent
4. 启动 Gateway Agent
5. 启动 API Gateway / Web UI
```

---

## 19. 三阶段落地计划

### Phase 1：MVP

范围：

- 单用户
- 内部知识库问答
- Gateway + RAG Agent
- 2~3 个 MCP tool
- 基础 trace 和 citation

目标：

- 能稳定回答带引用的问题

### Phase 2：多 Agent 扩展

范围：

- 接入 Search Agent
- 接入 Report Agent
- 支持 A2A 协作
- 引入任务状态持久化

目标：

- 能完成研究型问答和报告输出

### Phase 3：生产化

范围：

- 审批流
- 评估系统
- 权限策略
- 成本治理
- 异常恢复与回放

目标：

- 从“能跑”升级到“可运营”

---

## 20. 最容易踩的坑

### 20.1 一开始就做太多 Agent

建议先 3 个以内。

Agent 太多时，问题通常不是“更聪明”，而是：

- 路由难
- 状态乱
- 调试难
- 成本飙升

### 20.2 把 MCP 工具全暴露给模型

不建议。

应该通过 registry 做任务级可见性控制。

### 20.3 没有统一 State

如果 task、artifact、citation、approval 各自分散保存，后面很难回放和审计。

### 20.4 只看最终答案，不看执行过程

没有 `V`，调优会非常慢。

### 20.5 过度依赖 prompt，忽略 Graph 和 Hook

复杂系统的稳定性，更多来自：

- 明确的 Graph
- 明确的 state
- 明确的 hook
- 明确的 eval

---

## 21. 验收清单

在进入下一阶段前，建议至少确认下面这些点：

### 21.1 MVP 验收

- 能回答 20 个内部知识问题
- 每个答案都有来源引用
- trace 可回看
- 检索失败不会直接崩溃

### 21.2 多 Agent 验收

- Gateway 能正确调用 2 个以上外部 Agent
- A2A 超时和失败能降级
- 最终报告能区分内部证据和外部证据

### 21.3 生产化验收

- 有审批节点
- 有任务状态恢复
- 有失败分类统计
- 有成本和调用量监控

---

## 22. 一句话总结架构原则

```text
Eino 管内部编排，
MCP 管工具边界，
A2A 管协作边界，
Harness 管系统治理。
```

如果把这四层分清楚，这个项目就会非常稳；
如果把它们混成一个“大而全 Agent”，后面会越来越难维护。
