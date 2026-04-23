# Eino 框架指南

> Eino 是字节跳动开源的 Go 语言 LLM 应用开发框架，已捐赠至 CloudWeGo 社区。当前版本 v0.8.x。
> 官方文档：https://www.cloudwego.io/zh/docs/eino/
> GitHub：https://github.com/cloudwego/eino

---

## 1. 核心设计理念

Eino 的设计围绕两个核心概念展开：**组件（Component）** 和 **编排（Orchestration）**。

```
组件（Component）
  └── 标准化 interface：输入/输出类型 + 流处理范式 + CallOption

编排（Orchestration）
  └── 将组件连接成有向图（Graph / Chain / Workflow）
```

**五大设计优势：**

| 优势 | 说明 |
|------|------|
| 内核稳定 | 组件 interface + 编排范式构成稳定 API |
| 敏捷扩展 | 组件实现可插拔，最佳实践持续封装 |
| 高可靠 | Golang 强类型 + 编译期类型检查 |
| 实践驱动 | 豆包/抖音/扣子等数百个内部服务验证 |
| 工具生态 | 内置 tracing、APMPlus/Langfuse 集成、IDE 可视化编排插件 |

---

## 2. 架构分层

```
eino-examples        ← 最佳实践样例
      ↓
flow/                ← ReAct Agent、Multi-Agent 等高级封装
      ↓
compose/             ← Graph / Chain / Workflow 编排引擎
      ↓
schema/              ← Message、Tool、Document 等核心类型
      ↓
components/          ← ChatModel、Retriever、Embedding 等接口定义
      ↓
eino-ext/            ← 各组件的第三方具体实现
```

---

## 3. 组件系统（Components）

### 3.1 组件总览

| 组件 | 接口 | 职责 |
|------|------|------|
| **ChatModel** | `Generate(ctx, messages)` | 与 LLM 交互 |
| **ChatTemplate** | `Format(ctx, vars)` | 构建 Prompt |
| **Tool / ToolsNode** | `Run(ctx, input)` | 工具调用执行 |
| **Retriever** | `Retrieve(ctx, query)` | 向量检索 |
| **Document Loader** | `Load(ctx, source)` | 文档加载 |
| **Document Transformer** | `Transform(ctx, docs)` | 文档切分/处理 |
| **Embedding** | `EmbedStrings(ctx, texts)` | 文本向量化 |
| **Indexer** | `Store(ctx, docs)` | 向量存储 |
| **Lambda** | 任意函数 | 将自定义逻辑纳入编排图 |

### 3.2 支持的模型和后端

| 类型 | 支持后端 |
|------|--------|
| ChatModel | OpenAI、ARK（火山方舟/豆包）、Gemini、Claude、Ollama 等 |
| Embedding | ARK、OpenAI、其他兼容 API |
| Retriever | Milvus v2、Elasticsearch 7/8/9、OpenSearch、Redis、VikingDB、Dify |
| Indexer | Milvus v2、Elasticsearch、OpenSearch、Redis |

---

## 4. 编排系统（Orchestration）

### 4.1 三种编排范式对比

| 范式 | 适用场景 | 支持循环 | 复杂度 |
|------|---------|--------|-------|
| **Chain** | 线性顺序流程（RAG、简单问答）| 否 | 低 |
| **Graph** | 分支、并行、循环（Agent 推理）| 是 | 高 |
| **Workflow** | 字段级数据路由的有向无环图 | 部分 | 中 |

### 4.2 Chain — 链式编排

最简单的顺序组件连接方式：

```go
import "github.com/cloudwego/eino/compose"

chain := compose.NewChain[string, *schema.Message]()

chain.
    AppendRetriever(retriever).      // 检索
    AppendLambda(formatDocs).        // 格式化文档
    AppendChatTemplate(tmpl).        // 构建 Prompt
    AppendChatModel(chatModel)       // 生成答案

runner, err := chain.Compile(ctx)
result, err := runner.Invoke(ctx, "用户问题")
```

### 4.3 Graph — 全功能有向图

Graph 支持分支、并行、循环，是实现 Agent 的核心工具：

```go
graph := compose.NewGraph[[]*schema.Message, *schema.Message](
    // 可选：注入共享状态
    compose.WithGenLocalState(func(ctx context.Context) *MyState {
        return &MyState{}
    }),
)

// 添加节点
graph.AddChatModelNode("model", chatModel,
    compose.WithStatePreHandler(modelPreHandler))   // 节点前处理（读写 State）
graph.AddToolsNode("tools", toolsNode)

// 添加边
graph.AddEdge(compose.START, "model")
graph.AddEdge("tools", "model")                    // 循环边：工具结果回送模型

// 条件分支
graph.AddBranch("model", compose.NewStreamGraphBranch(
    func(ctx context.Context, msg *schema.Message) (string, error) {
        if msg.ToolCalls != nil {
            return "tools", nil    // 有工具调用 → 执行工具
        }
        return compose.END, nil   // 无工具调用 → 结束
    },
    map[string]bool{"tools": true, compose.END: true},
))

runner, err := graph.Compile(ctx, compose.WithMaxRunSteps(10))
result, err := runner.Invoke(ctx, messages)
```

#### Graph 节点类型

| 方法 | 对应组件 |
|------|--------|
| `AddChatModelNode` | ChatModel |
| `AddRetrieverNode` | Retriever |
| `AddEmbeddingNode` | Embedding |
| `AddToolsNode` | ToolsNode |
| `AddDocumentTransformerNode` | DocumentTransformer |
| `AddLambdaNode` | 任意自定义函数 |
| `AddGraphNode` | 嵌套子图 |

#### 拓扑模式

```
顺序：  START → A → B → END
并行：  START → A ─┬─ C → END
              └─ B ─┘
分支：  START → Router → A → END
                    └── B → END
循环：  START → Model ←─ Tools
                  └─────────→ END
```

### 4.4 流式处理

Eino 在编排层统一处理流式输出，调用方只需选择执行方式：

```go
// 非流式
result, err := runner.Invoke(ctx, input)

// 流式（逐 token 输出）
stream, err := runner.Stream(ctx, input)
for {
    chunk, err := stream.Recv()
    if err == io.EOF { break }
    fmt.Print(chunk.Content)
}
```

Eino 自动处理：流的复制与分发、多流合并、单流拼接，以及基于首包的路由判断（降低决策延迟）。

---

## 5. RAG Pipeline 完整实现

### 5.1 离线索引流程（文档入库）

```go
// 1. 加载文档
loader, _ := fileloader.NewFileLoader(ctx, &fileloader.Config{
    FilePath: "./data/knowledge.txt",
})
docs, _ := loader.Load(ctx, document.Source{})

// 2. 切分文档
splitter, _ := textsplitter.NewTextSplitter(ctx, &textsplitter.Config{
    ChunkSize:    512,
    ChunkOverlap: 50,
})
chunks, _ := splitter.Transform(ctx, docs)

// 3. Embedding + 存储
embedder, _ := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
    APIKey: os.Getenv("ARK_API_KEY"),
    Model:  "ep-your-embed-model",
})
indexer, _ := milvus2.NewIndexer(ctx, &milvus2.IndexerConfig{
    Address:        "localhost:19530",
    CollectionName: "my_knowledge_base",
    Embedder:       embedder,
})
indexer.Store(ctx, chunks)
```

### 5.2 在线检索生成流程（Graph 编排）

```go
// 构建 RAG Graph
graph := compose.NewGraph[string, *schema.Message]()

// 节点
graph.AddRetrieverNode("retriever", retriever)
graph.AddLambdaNode("format", compose.InvokableLambda(
    func(ctx context.Context, docs []*schema.Document) (map[string]any, error) {
        context := formatDocs(docs)
        return map[string]any{"context": context}, nil
    },
))
graph.AddChatTemplateNode("template", tmpl)
graph.AddChatModelNode("model", chatModel)

// 边
graph.AddEdge(compose.START, "retriever")
graph.AddEdge("retriever", "format")
graph.AddEdge("format", "template")
graph.AddEdge("template", "model")
graph.AddEdge("model", compose.END)

runner, _ := graph.Compile(ctx)
result, _ := runner.Invoke(ctx, "用户问题")
```

### 5.3 支持的 Retriever 后端

| 后端 | 包路径 | 特点 |
|------|-------|------|
| Milvus v2 | `eino-ext/components/retriever/milvus2` | 高性能，推荐 |
| Elasticsearch | `eino-ext/components/retriever/es8` | 支持混合检索 |
| Redis | `eino-ext/components/retriever/redis` | 轻量，低延迟 |
| VikingDB | `eino-ext/components/retriever/vikingdb` | 火山引擎 |
| Dify | `eino-ext/components/retriever/dify` | 对接 Dify 知识库 |

---

## 6. ReAct Agent 实现

ReAct（Reasoning + Acting）是最常用的 Agent 模式。Eino 通过 Graph 循环边实现：

```go
// ReAct Agent = Model 节点 + Tools 节点 + 循环
func buildReActAgent(ctx context.Context, model model.ChatModel, tools []tool.BaseTool) {
    toolsNode, _ := toolsnode.NewToolNode(ctx, &toolsnode.ToolsNodeConfig{
        Tools: tools,
    })

    graph := compose.NewGraph[[]*schema.Message, *schema.Message](
        compose.WithGenLocalState(func(ctx context.Context) *agentState {
            return &agentState{Messages: make([]*schema.Message, 0)}
        }),
    )

    graph.AddChatModelNode("model", model,
        compose.WithStatePreHandler(func(ctx context.Context, in []*schema.Message, state *agentState) ([]*schema.Message, error) {
            // 将历史消息注入
            return append(state.Messages, in...), nil
        }),
        compose.WithStatePostHandler(func(ctx context.Context, out *schema.Message, state *agentState) (*schema.Message, error) {
            // 保存模型输出到历史
            state.Messages = append(state.Messages, out)
            return out, nil
        }),
    )
    graph.AddToolsNode("tools", toolsNode)

    graph.AddEdge(compose.START, "model")
    graph.AddEdge("tools", "model")   // 工具结果回送模型（循环）

    graph.AddBranch("model", compose.NewStreamGraphBranch(
        func(ctx context.Context, msg *schema.Message) (string, error) {
            if len(msg.ToolCalls) > 0 {
                return "tools", nil
            }
            return compose.END, nil
        },
        map[string]bool{"tools": true, compose.END: true},
    ))

    agent, _ := graph.Compile(ctx,        compose.WithMaxRunSteps(10))
    return agent
}
```

---

## 7. ADK — Agent Development Kit

Eino ADK 是更高层次的多 Agent 封装，内置 Middleware 机制：

```go
agent, _ := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
    Model: chatModel,
    Tools: tools,
    Middleware: []adk.Middleware{
        middleware.NewSummarizationMiddleware(summaryConfig), // 历史压缩
        middleware.NewSkillMiddleware(skillConfig),           // 技能路由
    },
})
runner := adk.NewRunner(agent)
stream, _ := runner.Stream(ctx, session, messages)
```

### ADK 内置 Agent 类型

| Agent 类型 | 适用场景 |
|-----------|--------|
| `ChatModelAgent` | 单 Agent 多轮对话 |
| `WorkflowAgent` | 工作流驱动 |
| `SupervisorAgent` | 多 Agent 监管协调 |
| `PlanExecuteAgent` | 规划-执行模式 |
| `HostMultiAgent` | 多 Agent 托管协作 |

---

## 8. 快速上手

```bash
go get github.com/cloudwego/eino@latest
go get github.com/cloudwego/eino-ext/components/model/ark@latest
go get github.com/cloudwego/eino-ext/components/retriever/milvus2@latest
go get github.com/cloudwego/eino-ext/components/embedding/ark@latest
```

学习路径：ChatModel 基础 → RAG Pipeline → ReAct Agent → Multi-Agent ADK → A2A 集成
