# Agent 设计模式

> 基于 Eino 框架的 Agent 设计模式详解，涵盖 ReAct、Plan-Execute、Multi-Agent 等核心范式。

---

## 1. Agent 核心要素

```
Agent = LLM（大脑）+ Memory（记忆）+ Tools（工具）+ Orchestration（编排）
```

| 要素 | Eino 对应 | 作用 |
|------|----------|------|
| LLM | `ChatModel` | 推理与决策 |
| Memory | `State` / Session | 短期/长期记忆 |
| Tools | `Tool` / `ToolsNode` | 与外部系统交互 |
| 编排 | `Graph` / ADK | 控制执行流程 |

---

## 2. ReAct 模式

最经典的 Agent 模式：**推理（Reasoning）+ 行动（Acting）** 交替循环。

### 执行流程

```
问题
  → [Think] LLM 分析：需要调用哪个工具？
  → [Act]   执行工具调用
  → [Observe] 获取工具结果
  → [Think] LLM 分析结果，是否需要继续？
  → 循环直到得出最终答案
```

### Eino Graph 实现

```go
graph := compose.NewGraph[[]*schema.Message, *schema.Message](
    compose.WithGenLocalState(func(ctx context.Context) *State {
        return &State{Messages: []*schema.Message{}}
    }),
)

// 节点
graph.AddChatModelNode("model", chatModel,
    compose.WithStatePreHandler(injectHistory),
    compose.WithStatePostHandler(saveToHistory),
)
graph.AddToolsNode("tools", toolsNode)

// 拓扑：循环结构
graph.AddEdge(compose.START, "model")
graph.AddEdge("tools", "model")  // 工具结果回送 → 形成循环

// 条件分支：有工具调用则执行工具，否则结束
graph.AddBranch("model", compose.NewStreamGraphBranch(
    func(ctx context.Context, msg *schema.Message) (string, error) {
        if len(msg.ToolCalls) > 0 {
            return "tools", nil
        }
        return compose.END, nil
    },
    map[string]bool{"tools": true, compose.END: true},
))
```

### 适用场景
- 需要多步工具调用的任务（搜索→分析→汇总）
- 任务步骤不确定，需要动态决策
- 单 Agent 可以处理的复杂问题

---

## 3. Plan-Execute 模式

**先规划，再执行**。适合任务结构相对固定、步骤可以预先确定的场景。

### 流程

```
问题
  → [Planner] LLM 生成执行计划（JSON 结构化步骤列表）
  → [Executor] 按计划逐步执行（每步可调用工具）
  → [Reviewer] 检查结果是否满足要求
  → 输出最终答案
```

### Eino 实现思路

```go
// 规划节点：输出结构化计划
plannerLambda := compose.InvokableLambda(func(ctx context.Context, q string) (*Plan, error) {
    resp, _ := plannerModel.Generate(ctx, buildPlanPrompt(q))
    return parsePlan(resp.Content), nil
})

// 执行节点：按步骤执行
executorLambda := compose.InvokableLambda(func(ctx context.Context, plan *Plan) (string, error) {
    results := []string{}
    for _, step := range plan.Steps {
        result, _ := executeTool(ctx, step.Tool, step.Input)
        results = append(results, result)
    }
    return synthesize(ctx, plan, results), nil
})

graph := compose.NewGraph[string, string]()
graph.AddLambdaNode("planner", plannerLambda)
graph.AddLambdaNode("executor", executorLambda)
graph.AddEdge(compose.START, "planner")
graph.AddEdge("planner", "executor")
graph.AddEdge("executor", compose.END)
```

### ReAct vs Plan-Execute

| 维度 | ReAct | Plan-Execute |
|------|-------|-------------|
| 规划方式 | 每步动态决策 | 一次性生成完整计划 |
| 灵活性 | 高 | 中 |
| 可解释性 | 中 | 高（计划可审查）|
| 适合场景 | 探索性任务 | 结构化流程任务 |
| 错误恢复 | 自动重试 | 需要重新规划 |

---

## 4. Multi-Agent 模式

将复杂任务分解给多个专业 Agent 协作完成。

### 4.1 Supervisor 模式（主管协调）

```
用户请求
    ↓
[Supervisor Agent]  ← 负责分解任务、分配给子 Agent、汇总结果
    ├──► [Research Agent]   负责信息检索
    ├──► [Analysis Agent]   负责数据分析
    └──► [Writer Agent]     负责内容生成
```

```go
// Eino ADK Supervisor 实现
supervisor, _ := adk.NewSupervisorAgent(ctx, &adk.SupervisorConfig{
    Model: orchestratorModel,
    SubAgents: []adk.SubAgent{
        {Name: "research", Agent: researchAgent},
        {Name: "analysis", Agent: analysisAgent},
        {Name: "writer",   Agent: writerAgent},
    },
})
```

### 4.2 Pipeline 模式（流水线）

```
Agent A（数据获取）→ Agent B（数据处理）→ Agent C（生成报告）
```

每个 Agent 的输出是下一个 Agent 的输入，适合有明确先后顺序的任务。

### 4.3 Peer-to-Peer 模式（对等协作）

通过 A2A 协议实现跨服务 Agent 协作：

```
Agent A（本地 Eino）
    │
    │  A2A 协议
    ▼
Agent B（远程 Python/LangChain）
Agent C（远程 Java/Spring AI）
```

---

## 5. Corrective RAG Agent

结合 RAG 和反思机制的 Agent 模式：

```
问题
  → 检索文档
  → [Grader] 评估文档相关性
    ├── 相关性高 → 生成答案
    ├── 相关性中 → 补充 Web 搜索 → 融合后生成
    └── 相关性低 → 重写查询 → 重新检索
  → [Hallucination Checker] 检测幻觉
    ├── 无幻觉 → 输出
    └── 有幻觉 → 重新生成
```

### Eino Graph 实现

```go
graph := compose.NewGraph[string, string]()
graph.AddRetrieverNode("retrieve", retriever)
graph.AddLambdaNode("grade", graderLambda)
graph.AddLambdaNode("rewrite", rewriteLambda)
graph.AddLambdaNode("web_search", webSearchLambda)
graph.AddChatModelNode("generate", chatModel)
graph.AddLambdaNode("check_hallucination", checkerLambda)

// 条件路由
graph.AddBranch("grade", relevanceBranch)    // 相关/不相关
graph.AddBranch("check_hallucination", hallucinationBranch) // 有/无幻觉
```

---

## 6. Human-in-the-Loop

Eino 通过 Checkpoint 机制实现 Agent 暂停等待人工介入：

```go
// 编译时注入 CheckpointStore
runner, _ := graph.Compile(ctx,
    compose.WithCheckPointStore(checkpointStore),
)

// 执行 → 遇到 interrupt 节点时暂停
result, err := runner.Invoke(ctx, input,
    compose.WithCheckPointID("thread-001"),
)
// err 包含中断信号

// 人工审查后继续执行
result, err = runner.Invoke(ctx, humanFeedback,
    compose.WithCheckPointID("thread-001"),
    compose.WithResume(),
)
```

对应 A2A 协议中的 `input-required` 状态。

---

## 7. 模式选型指南

| 场景 | 推荐模式 |
|------|--------|
| 通用问答、工具调用 | ReAct |
| 多步骤有序任务 | Plan-Execute |
| 需要人工审批 | ReAct + Checkpoint |
| 专业化分工的复杂任务 | Supervisor Multi-Agent |
| 跨框架/跨服务协作 | Multi-Agent + A2A |
| 知识库问答质量要求高 | Corrective RAG Agent |
