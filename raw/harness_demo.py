"""
Agent Harness Demo — H = (E, T, C, S, L, V)
=============================================
基于论文 "Agent Harness for LLM Agents: A Survey" 的简易实现。
展示六个核心组件如何协同工作，构建一个可靠的 Agent 系统。

依赖: pip install openai  (使用 OpenAI 兼容 API)
用法: python harness_demo.py
"""

from __future__ import annotations

import json
import time
import hashlib
from datetime import datetime
from dataclasses import dataclass, field
from typing import Any, Callable, Optional
from pathlib import Path

# ============================================================
# S — State Store (状态存储)
# 跨轮次持久化任务状态，支持故障恢复
# ============================================================

class StateStore:
    """S 组件：管理跨轮次/跨会话的持久化状态"""

    def __init__(self, persist_path: str | None = None):
        self._state: dict[str, Any] = {}
        self._history: list[dict] = []       # 写操作审计日志
        self._persist_path = persist_path

    def get(self, key: str, default=None):
        return self._state.get(key, default)

    def set(self, key: str, value: Any):
        """写入状态，同时记录审计日志（写入可审计性）"""
        self._history.append({
            "op": "set", "key": key, "value": value,
            "timestamp": datetime.now().isoformat()
        })
        self._state[key] = value

    def snapshot(self) -> dict:
        """获取完整状态快照，用于故障恢复"""
        return {"state": dict(self._state), "history": list(self._history)}

    def restore(self, snapshot: dict):
        """从快照恢复状态"""
        self._state = snapshot["state"]
        self._history = snapshot["history"]


# ============================================================
# T — Tool Registry (工具注册表)
# 维护工具接口目录，路由和监控工具调用
# ============================================================

@dataclass
class ToolSpec:
    """工具规范：名称、描述、参数 schema、执行函数"""
    name: str
    description: str
    parameters: dict          # JSON Schema
    function: Callable
    scope: str = "general"    # 能力范围: general / file / network / execute


class ToolRegistry:
    """T 组件：管理工具注册、发现和调用路由"""

    def __init__(self):
        self._tools: dict[str, ToolSpec] = {}
        self._call_log: list[dict] = []  # 工具调用监控日志

    def register(self, tool: ToolSpec):
        """注册工具（带 schema 验证）"""
        self._tools[tool.name] = tool
        print(f"  [T] 注册工具: {tool.name} (scope={tool.scope})")

    def list_tools(self) -> list[dict]:
        """返回工具清单（供 LLM 选择）"""
        return [
            {
                "type": "function",
                "function": {
                    "name": t.name,
                    "description": t.description,
                    "parameters": t.parameters
                }
            }
            for t in self._tools.values()
        ]

    def call(self, name: str, arguments: dict) -> str:
        """路由并执行工具调用，记录日志"""
        if name not in self._tools:
            return json.dumps({"error": f"未知工具: {name}"})

        tool = self._tools[name]
        start = time.time()
        try:
            result = tool.function(**arguments)
            elapsed = time.time() - start
            self._call_log.append({
                "tool": name, "args": arguments,
                "result": result, "elapsed_ms": round(elapsed * 1000),
                "status": "success"
            })
            return json.dumps(result, ensure_ascii=False) if not isinstance(result, str) else result
        except Exception as e:
            self._call_log.append({
                "tool": name, "args": arguments,
                "error": str(e), "status": "error"
            })
            return json.dumps({"error": str(e)})


# ============================================================
# C — Context Manager (上下文管理器)
# 控制每一轮进入模型上下文窗口的信息
# ============================================================

class ContextManager:
    """C 组件：管理上下文窗口，防止上下文膨胀"""

    def __init__(self, max_messages: int = 20, system_prompt: str = ""):
        self.system_prompt = system_prompt
        self.max_messages = max_messages
        self._messages: list[dict] = []

    def add_message(self, role: str, content: str, **kwargs):
        """添加消息到上下文"""
        msg = {"role": role, "content": content, **kwargs}
        self._messages.append(msg)

    def add_tool_result(self, tool_call_id: str, content: str):
        """添加工具调用结果"""
        self._messages.append({
            "role": "tool", "tool_call_id": tool_call_id, "content": content
        })

    def get_context(self) -> list[dict]:
        """
        构建发送给 LLM 的上下文。
        策略：保留 system prompt + 最近 N 条消息（截断策略）
        生产系统中可替换为摘要/检索增强策略。
        """
        context = [{"role": "system", "content": self.system_prompt}]

        # 截断：只保留最近 max_messages 条
        if len(self._messages) > self.max_messages:
            truncated = len(self._messages) - self.max_messages
            print(f"  [C] 上下文截断: 丢弃最早 {truncated} 条消息")
            context += self._messages[-self.max_messages:]
        else:
            context += self._messages

        return context

    @property
    def message_count(self) -> int:
        return len(self._messages)


# ============================================================
# L — Lifecycle Hooks (生命周期钩子)
# 认证、日志、策略执行的前/后拦截点
# ============================================================

class LifecycleHooks:
    """L 组件：在关键节点执行拦截逻辑"""

    def __init__(self):
        self._hooks: dict[str, list[Callable]] = {
            "before_llm_call": [],
            "after_llm_call": [],
            "before_tool_call": [],
            "after_tool_call": [],
            "on_error": [],
            "on_step_complete": [],
        }
        self.logs: list[str] = []

    def register(self, event: str, hook: Callable):
        if event in self._hooks:
            self._hooks[event].append(hook)

    def trigger(self, event: str, **kwargs) -> bool:
        """
        触发钩子，返回 True 表示允许继续，False 表示拦截。
        任一钩子返回 False 即拦截。
        """
        for hook in self._hooks.get(event, []):
            result = hook(**kwargs)
            if result is False:
                self.logs.append(f"[L] {event} 被钩子拦截")
                return False
        return True


# ============================================================
# V — Evaluation Interface (评估接口)
# 捕获动作轨迹、中间状态、成功信号
# ============================================================

class EvaluationInterface:
    """V 组件：记录完整执行轨迹用于离线分析"""

    def __init__(self):
        self.trajectory: list[dict] = []
        self.start_time: float = 0
        self.total_tokens: int = 0
        self.total_cost: float = 0

    def record_step(self, step: int, action: str, detail: dict):
        self.trajectory.append({
            "step": step, "action": action,
            "timestamp": datetime.now().isoformat(),
            **detail
        })

    def record_tokens(self, prompt_tokens: int, completion_tokens: int):
        self.total_tokens += prompt_tokens + completion_tokens
        # 粗略成本估算 (以 GPT-4o-mini 为例)
        self.total_cost += prompt_tokens * 0.15 / 1e6 + completion_tokens * 0.6 / 1e6

    def summary(self) -> dict:
        elapsed = time.time() - self.start_time
        return {
            "total_steps": len(self.trajectory),
            "total_tokens": self.total_tokens,
            "estimated_cost_usd": round(self.total_cost, 6),
            "elapsed_seconds": round(elapsed, 2),
            "trajectory": self.trajectory
        }


# ============================================================
# E — Execution Loop (执行循环)
# 管理 观察-思考-行动 循环、终止条件、错误恢复
# ============================================================

class ExecutionLoop:
    """
    E 组件：Agent 的核心执行引擎。
    实现 Observe → Think → Act 循环，
    并通过 L/V 组件实现治理和可观测性。
    """

    def __init__(
        self,
        llm_client,               # OpenAI 兼容客户端
        model: str,
        tool_registry: ToolRegistry,
        context_manager: ContextManager,
        state_store: StateStore,
        lifecycle_hooks: LifecycleHooks,
        eval_interface: EvaluationInterface,
        max_steps: int = 10,      # 防止执行失控
    ):
        self.llm = llm_client
        self.model = model
        self.T = tool_registry
        self.C = context_manager
        self.S = state_store
        self.L = lifecycle_hooks
        self.V = eval_interface
        self.max_steps = max_steps

    def run(self, user_input: str) -> str:
        """执行完整的 Agent 任务循环"""

        self.V.start_time = time.time()
        self.S.set("status", "running")
        self.S.set("user_input", user_input)
        self.C.add_message("user", user_input)

        step = 0
        final_answer = ""

        while step < self.max_steps:
            step += 1
            self.S.set("current_step", step)
            print(f"\n{'='*50}")
            print(f"  Step {step}/{self.max_steps}")
            print(f"{'='*50}")

            # ---- Phase 1: Think (LLM 调用) ----
            if not self.L.trigger("before_llm_call", step=step):
                final_answer = "[被生命周期钩子拦截]"
                break

            context = self.C.get_context()
            tools = self.T.list_tools()

            try:
                response = self.llm.chat.completions.create(
                    model=self.model,
                    messages=context,
                    tools=tools if tools else None,
                    tool_choice="auto" if tools else None,
                )
            except Exception as e:
                self.L.trigger("on_error", error=e, step=step)
                self.V.record_step(step, "llm_error", {"error": str(e)})
                # 错误恢复：重试一次
                print(f"  [E] LLM 调用失败，尝试恢复: {e}")
                self.S.set("status", "error")
                final_answer = f"[执行错误: {e}]"
                break

            msg = response.choices[0].message

            # 记录 token 使用
            if response.usage:
                self.V.record_tokens(
                    response.usage.prompt_tokens,
                    response.usage.completion_tokens
                )

            self.L.trigger("after_llm_call", step=step, response=msg)

            # ---- Phase 2: Decide — 是否有工具调用？ ----
            if msg.tool_calls:
                # 将 assistant 消息（含 tool_calls）加入上下文
                self.C._messages.append(msg.to_dict() if hasattr(msg, 'to_dict') else {
                    "role": "assistant",
                    "content": msg.content,
                    "tool_calls": [
                        {
                            "id": tc.id,
                            "type": "function",
                            "function": {
                                "name": tc.function.name,
                                "arguments": tc.function.arguments
                            }
                        }
                        for tc in msg.tool_calls
                    ]
                })

                for tc in msg.tool_calls:
                    tool_name = tc.function.name
                    tool_args = json.loads(tc.function.arguments)
                    print(f"  [E] 调用工具: {tool_name}({tool_args})")

                    # ---- Phase 3: Act (工具执行) ----
                    if not self.L.trigger("before_tool_call",
                                          tool=tool_name, args=tool_args, step=step):
                        result = "[工具调用被策略拦截]"
                    else:
                        result = self.T.call(tool_name, tool_args)
                        self.L.trigger("after_tool_call",
                                       tool=tool_name, result=result, step=step)

                    print(f"  [E] 工具结果: {result[:200]}")
                    self.C.add_tool_result(tc.id, result)
                    self.V.record_step(step, "tool_call", {
                        "tool": tool_name, "args": tool_args,
                        "result": result[:500]
                    })

                self.L.trigger("on_step_complete", step=step)

            else:
                # 没有工具调用 → Agent 认为任务完成
                final_answer = msg.content or ""
                self.C.add_message("assistant", final_answer)
                self.V.record_step(step, "final_answer", {
                    "answer": final_answer[:500]
                })
                print(f"  [E] Agent 回复: {final_answer[:200]}")
                break

        # 终止条件检查
        if step >= self.max_steps and not final_answer:
            final_answer = "[达到最大步数限制，执行终止]"
            self.V.record_step(step, "max_steps_reached", {})

        self.S.set("status", "completed")
        self.S.set("final_answer", final_answer)
        return final_answer


# ============================================================
# Harness — 六组件统一封装
# H = (E, T, C, S, L, V)
# ============================================================

class AgentHarness:
    """
    Agent Harness: 完整的运行时治理系统。
    统一管理 E/T/C/S/L/V 六个组件。
    """

    def __init__(self, llm_client, model: str, system_prompt: str = "",
                 max_steps: int = 10, max_context_messages: int = 20):

        # 初始化六个组件
        self.S = StateStore()
        self.T = ToolRegistry()
        self.C = ContextManager(max_messages=max_context_messages,
                                system_prompt=system_prompt)
        self.L = LifecycleHooks()
        self.V = EvaluationInterface()
        self.E = ExecutionLoop(
            llm_client=llm_client,
            model=model,
            tool_registry=self.T,
            context_manager=self.C,
            state_store=self.S,
            lifecycle_hooks=self.L,
            eval_interface=self.V,
            max_steps=max_steps,
        )

        print(f"[Harness] 初始化完成 — H = (E, T, C, S, L, V)")
        print(f"[Harness] 模型: {model}, 最大步数: {max_steps}")

    def register_tool(self, name: str, description: str,
                      parameters: dict, function: Callable,
                      scope: str = "general"):
        """注册工具到 T 组件"""
        self.T.register(ToolSpec(
            name=name, description=description,
            parameters=parameters, function=function, scope=scope
        ))

    def add_hook(self, event: str, hook: Callable):
        """注册生命周期钩子到 L 组件"""
        self.L.register(event, hook)

    def run(self, user_input: str) -> str:
        """执行 Agent 任务"""
        print(f"\n{'#'*60}")
        print(f"  用户输入: {user_input}")
        print(f"{'#'*60}")
        # 每次任务重置上下文（保留 system prompt）
        self.C._messages.clear()
        result = self.E.run(user_input)
        return result

    def report(self) -> dict:
        """输出评估报告（V 组件）"""
        return self.V.summary()


# ============================================================
# Demo: 构建一个带工具的 Agent
# ============================================================

def demo_tools():
    """定义一组演示用工具"""

    def calculate(expression: str) -> dict:
        """安全的数学计算器"""
        allowed = set("0123456789+-*/().% ")
        if not all(c in allowed for c in expression):
            return {"error": "不安全的表达式"}
        try:
            result = eval(expression)  # demo only; 生产环境应使用 ast.literal_eval 或专用解析器
            return {"expression": expression, "result": result}
        except Exception as e:
            return {"error": str(e)}

    def get_weather(city: str) -> dict:
        """模拟天气查询"""
        weather_db = {
            "北京": {"temp": 22, "condition": "晴", "humidity": 35},
            "上海": {"temp": 26, "condition": "多云", "humidity": 72},
            "深圳": {"temp": 30, "condition": "雷阵雨", "humidity": 85},
            "杭州": {"temp": 24, "condition": "阴", "humidity": 60},
        }
        if city in weather_db:
            return {"city": city, **weather_db[city]}
        return {"city": city, "error": "暂无该城市数据"}

    def search_knowledge(query: str) -> dict:
        """模拟知识库检索"""
        kb = {
            "harness": "Agent Harness 是包裹 LLM 的运行时治理层，包含 E/T/C/S/L/V 六个组件。"
                       "其核心论点是：Harness 基础设施（而非模型能力）是 Agent 系统可靠性的绑定约束。",
            "mcp": "MCP (Model Context Protocol) 是 Anthropic 提出的 Agent→Tool 协议，"
                   "使用 JSON-RPC 传输，本地延迟 2-15ms。",
            "a2a": "A2A (Agent-to-Agent) 是 Google 提出的 Agent→Agent 协议，"
                   "使用 HTTPS+SSE 传输，支持多步协调和有状态会话。",
            "rag": "RAG (Retrieval-Augmented Generation) 通过检索外部知识增强 LLM 生成，"
                   "是解决幻觉和知识时效性的核心范式。",
        }
        results = []
        for key, content in kb.items():
            if key in query.lower() or any(w in content for w in query):
                results.append({"topic": key, "content": content})
        if not results:
            return {"query": query, "results": [], "message": "未找到相关内容"}
        return {"query": query, "results": results}

    return [
        ("calculate", "数学表达式计算器，支持加减乘除",
         {"type": "object", "properties": {"expression": {"type": "string", "description": "数学表达式"}},
          "required": ["expression"]},
         calculate, "general"),

        ("get_weather", "查询城市天气信息",
         {"type": "object", "properties": {"city": {"type": "string", "description": "城市名称"}},
          "required": ["city"]},
         get_weather, "network"),

        ("search_knowledge", "搜索知识库中的技术概念",
         {"type": "object", "properties": {"query": {"type": "string", "description": "搜索关键词"}},
          "required": ["query"]},
         search_knowledge, "general"),
    ]


def demo_hooks():
    """定义演示用生命周期钩子"""

    def log_llm_call(step, **kw):
        print(f"  [L] Hook: LLM 调用开始 (step={step})")
        return True

    def log_tool_call(tool, args, step, **kw):
        print(f"  [L] Hook: 工具调用 {tool} (step={step})")
        # 策略示例：禁止调用危险工具
        if tool == "dangerous_tool":
            print(f"  [L] Hook: 策略拦截! 禁止调用 {tool}")
            return False
        return True

    def log_step_complete(step, **kw):
        print(f"  [L] Hook: Step {step} 完成")
        return True

    return [
        ("before_llm_call", log_llm_call),
        ("before_tool_call", log_tool_call),
        ("on_step_complete", log_step_complete),
    ]


# ============================================================
# MockLLM — 无需 API Key 的本地模拟
# ============================================================

class MockChatCompletion:
    """模拟 OpenAI chat completion 响应"""

    def __init__(self):
        self._step = 0

    def create(self, model, messages, tools=None, tool_choice=None):
        """根据对话内容模拟 LLM 的工具调用和回复"""
        self._step += 1
        user_msg = ""
        for m in reversed(messages):
            if m.get("role") == "user":
                user_msg = m.get("content", "")
                break

        # 检查是否是工具结果后的回复
        last_msg = messages[-1] if messages else {}
        if last_msg.get("role") == "tool":
            # 工具结果已返回，生成最终回复
            tool_results = []
            for m in messages:
                if m.get("role") == "tool":
                    tool_results.append(m.get("content", ""))
            summary = "; ".join(tool_results[-3:])
            return self._text_response(
                f"根据查询结果：{summary}\n\n以上就是我为您查到的信息。"
            )

        # 根据用户输入决定调用哪些工具
        if "天气" in user_msg or "weather" in user_msg.lower():
            city = "北京"
            for c in ["北京", "上海", "深圳", "杭州"]:
                if c in user_msg:
                    city = c
                    break
            return self._tool_response("get_weather", {"city": city})

        if "计算" in user_msg or "算" in user_msg:
            import re
            expr_match = re.search(r'[\d+\-*/().\s]+', user_msg)
            expr = expr_match.group().strip() if expr_match else "1+1"
            # 确保提取完整表达式
            for candidate in re.findall(r'[(\d][\d+\-*/().\s]+[\d)]', user_msg):
                if len(candidate) > len(expr):
                    expr = candidate.strip()
            return self._tool_response("calculate", {"expression": expr})

        if any(kw in user_msg.lower() for kw in ["harness", "mcp", "a2a", "rag", "知识"]):
            query = user_msg
            for kw in ["harness", "mcp", "a2a", "rag"]:
                if kw in user_msg.lower():
                    query = kw
                    break
            return self._tool_response("search_knowledge", {"query": query})

        # 默认：直接回复
        return self._text_response(
            f"你好！我是基于 Harness 框架运行的 Agent。\n"
            f"我可以帮你：查天气、做计算、搜索知识库。\n"
            f"试试问我：'北京天气怎么样？' 或 '帮我计算 123*456'"
        )

    def _tool_response(self, name, args):
        return _MockResponse(tool_calls=[_MockToolCall(name, args)])

    def _text_response(self, content):
        return _MockResponse(content=content)


class _MockToolCall:
    def __init__(self, name, args):
        self.id = f"call_{hashlib.md5(name.encode()).hexdigest()[:8]}"
        self.function = _MockFunction(name, json.dumps(args, ensure_ascii=False))

class _MockFunction:
    def __init__(self, name, arguments):
        self.name = name
        self.arguments = arguments

class _MockMessage:
    def __init__(self, content=None, tool_calls=None):
        self.content = content
        self.tool_calls = tool_calls
        self.role = "assistant"

class _MockUsage:
    def __init__(self):
        self.prompt_tokens = 150
        self.completion_tokens = 50

class _MockResponse:
    def __init__(self, content=None, tool_calls=None):
        self.choices = [type("Choice", (), {"message": _MockMessage(content, tool_calls)})()]
        self.usage = _MockUsage()

class MockLLMClient:
    def __init__(self):
        self.chat = type("Chat", (), {"completions": MockChatCompletion()})()


# ============================================================
# Main
# ============================================================

def main():
    print("""
╔══════════════════════════════════════════════════════════════╗
║          Agent Harness Demo — H = (E, T, C, S, L, V)       ║
║                                                              ║
║  基于论文 "Agent Harness for LLM Agents: A Survey" 实现     ║
╚══════════════════════════════════════════════════════════════╝
""")

    # ----- 选择 LLM 后端 -----
    use_real_llm = False
    llm_client = None
    model = "mock"

    try:
        import openai
        import os
        api_key = os.environ.get("OPENAI_API_KEY")
        base_url = os.environ.get("OPENAI_BASE_URL")
        if api_key:
            llm_client = openai.OpenAI(api_key=api_key, base_url=base_url)
            model = os.environ.get("MODEL_NAME", "gpt-4o-mini")
            # 验证 API key 是否有效
            llm_client.models.list()
            use_real_llm = True
            print(f"[Setup] 使用真实 LLM: {model}")
        else:
            raise ValueError("no key")
    except Exception:
        llm_client = MockLLMClient()
        model = "mock-llm"
        print("[Setup] 使用 MockLLM (设置 OPENAI_API_KEY 环境变量可使用真实模型)")

    # ----- 构建 Harness -----
    harness = AgentHarness(
        llm_client=llm_client,
        model=model,
        system_prompt=(
            "你是一个有用的AI助手，运行在 Agent Harness 治理框架之下。\n"
            "你可以使用提供的工具来回答问题。\n"
            "回答要简洁、准确。如果需要使用工具，直接调用即可。"
        ),
        max_steps=5,
        max_context_messages=20,
    )

    # ----- 注册工具 (T 组件) -----
    print("\n[Setup] 注册工具...")
    for name, desc, params, func, scope in demo_tools():
        harness.register_tool(name, desc, params, func, scope)

    # ----- 注册钩子 (L 组件) -----
    print("\n[Setup] 注册生命周期钩子...")
    for event, hook in demo_hooks():
        harness.add_hook(event, hook)
        print(f"  [L] 注册钩子: {event}")

    # ----- 运行示例任务 -----
    test_queries = [
        "北京和上海的天气怎么样？",
        "帮我计算 (15 + 27) * 3",
        "什么是 Agent Harness？帮我查一下",
    ]

    for i, query in enumerate(test_queries, 1):
        print(f"\n\n{'*'*60}")
        print(f"  Demo {i}/{len(test_queries)}")
        print(f"{'*'*60}")
        answer = harness.run(query)
        print(f"\n  最终回答: {answer[:300]}")

    # ----- 评估报告 (V 组件) -----
    print(f"\n\n{'='*60}")
    print("  评估报告 (V 组件)")
    print(f"{'='*60}")
    report = harness.report()
    print(f"  总步数:     {report['total_steps']}")
    print(f"  总 Tokens:  {report['total_tokens']}")
    print(f"  估算成本:   ${report['estimated_cost_usd']}")
    print(f"  总耗时:     {report['elapsed_seconds']}s")

    # ----- 状态快照 (S 组件) -----
    print(f"\n  状态快照 (S 组件):")
    snapshot = harness.S.snapshot()
    for k, v in snapshot["state"].items():
        val_str = str(v)[:80]
        print(f"    {k}: {val_str}")

    print(f"\n  执行轨迹 (V 组件):")
    for t in report["trajectory"]:
        print(f"    Step {t['step']}: {t['action']} @ {t['timestamp']}")

    print(f"\n{'='*60}")
    print("  Demo 完成! ")
    print(f"{'='*60}")
    print("""
架构回顾:
  E (执行循环)    — 驱动 Observe→Think→Act 循环，防止失控
  T (工具注册表)  — 管理工具注册、路由和监控
  C (上下文管理)  — 控制上下文窗口大小，防止膨胀
  S (状态存储)    — 跨轮次持久化状态，支持故障恢复
  L (生命周期钩子) — 在关键节点执行策略拦截
  V (评估接口)    — 记录完整轨迹用于离线分析
""")


if __name__ == "__main__":
    main()
