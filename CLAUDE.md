# CLAUDE.md

本文件定义 Claude Code 在此仓库中的工作方式。

这个仓库不是传统的软件工程仓库，而是一个由 LLM 维护的知识库工程。Claude 的职责不是“随便回答问题”，而是把 `raw/` 中的原始资料持续沉淀为 `wiki/` 中可浏览、可追踪、可复用的 Markdown Wiki。

## 目标

把仓库中的资料维护成一套持续演化的 LLM Wiki：

- `raw/` 是不可变的原始资料层（source of truth）
- `wiki/` 是 Claude 持续维护的知识层
- `outputs/` 是按需生成的图表、演示文稿、分析产物等输出层

目标不是每次提问都重新从原始资料里“现查现拼”，而是让知识以 Wiki 的形式累计、交叉引用、不断更新。

## 仓库契约

### 1. `raw/` 只读

- `raw/` 中的文件默认一律视为原始资料，不主动修改、不重写、不挪动。
- 只有用户明确要求时，才允许整理、重命名或移动 `raw/` 中的内容。
- 二进制文件（如 `.dmg`）默认不执行，只记录其存在与上下文意义。

### 2. `wiki/` 由 Claude 负责维护

- `wiki/` 中的 Markdown 是 Claude 的主要工作区。
- Claude 负责创建页面、更新页面、补充交叉引用、维护索引与日志。
- 用户主要负责提供资料、提出问题、决定关注方向。

### 3. 所有重要结论必须可追溯

- Wiki 中的重要事实、定义、架构判断、对比结论，都应尽量追溯到 `raw/` 中的一个或多个来源。
- 必须明确区分：
  - 来源直接陈述的事实
  - Claude 基于多来源做出的综合归纳
  - 尚未证实的推测或待验证问题
- 若不同来源互相冲突，不要静默覆盖，必须显式记录冲突。

### 4. 默认语言

- Wiki 默认使用简体中文撰写。
- 英文术语、框架名、协议名、论文名保留原文。
- 首次出现时，优先使用“中文解释 + 英文术语”的形式，例如：多智能体协作协议（A2A）。

## 目录语义

### 原始资料层

- `raw/`：原始资料根目录
- `raw/*.md`：人工整理或 LLM 生成的原始说明稿、阅读笔记、设计稿
- `raw/*.pdf`：论文、报告、调查综述
- `raw/*.svg` / `raw/*.html`：图示、可视化、交互材料
- `raw/agent-platform/`：示例代码库，应视为“项目级原始资料”

### Wiki 层

优先采用以下结构；如内容还不够多，可以逐步创建，不必一次建全：

- `wiki/index.md`
  - Wiki 总目录
  - Claude 回答问题前优先读取
- `wiki/overview.md`
  - 当前主题地图、知识主线、 ingest 优先级
- `wiki/log.md`
  - 追加式时间日志，记录 ingest / query / lint / 重构
- `wiki/sources/`
  - 一份原始资料对应一个来源页
- `wiki/topics/`
  - 跨来源主题页，如 RAG、Agentic RAG、GraphRAG、Evaluation
- `wiki/concepts/`
  - 概念页，如 Tool Use、Memory、Planning、Retrieval、Hallucination
- `wiki/frameworks/`
  - 框架页，如 Eino
- `wiki/protocols/`
  - 协议页，如 A2A、MCP
- `wiki/projects/`
  - 项目页，如 `agent-platform`
- `wiki/analyses/`
  - 针对提问形成的分析页、对比页、方案页
- `wiki/questions/`
  - 未决问题、资料缺口、待验证假设

如果目录尚不存在，可在需要时创建。

### 输出层

`outputs/` 不是知识主存储，而是“由当前任务派生出来的产物层 / 中间产物层”。

推荐结构：

- `outputs/index.md`
  - 输出目录与说明
- `outputs/analyses/`
  - 一次性分析稿、对比稿、阶段总结、候选结论
- `outputs/slides/`
  - Marp / PPT / 演示稿
- `outputs/charts/`
  - 图表、数据可视化、导出图片
- `outputs/tables/`
  - CSV / TSV / JSON / 结构化结果
- `outputs/tmp/`
  - 临时抽取文件、OCR 结果、文本转储、脚本中间结果

约定：

- `wiki/` 保存“长期知识”
- `outputs/` 保存“面向当前任务的派生产物”
- 如果某个 `outputs/` 中的结果被证明具有长期复用价值，应再沉淀回 `wiki/`
- `outputs/tmp/` 可以覆盖、清理或重生成，不应被当作知识来源

## 页面设计原则

### 1. 一个页面只承载一个稳定主题

- 不要把多个独立主题堆进一个超长页面。
- 当一个主题值得被多次引用时，应提升为独立页面。
- 当多个页面高度重叠时，合并而不是并行重复维护。

### 2. 优先“综合”，避免机械摘抄

- 来源页可以总结单一资料，但也不要简单复制原文。
- 主题页、概念页、协议页必须跨来源综合，而不是把几段摘要拼接在一起。
- 若原始资料已经是整理稿，仍要进一步抽取结构、关系、分歧和结论。

### 3. 交叉引用是一级任务

- Wiki 内部链接优先使用 Obsidian 风格的 `[[页面名]]`。
- 每个页面都应主动链接到相关页面，而不是孤立存在。
- 任何值得长期复用的问答结果，也应沉淀为 Wiki 页面，而不是停留在聊天记录里。

## Frontmatter 约定

Wiki 页面建议带 YAML frontmatter，便于后续 Obsidian Dataview / 自动化使用。

推荐字段：

```yaml
---
type: source | topic | concept | framework | protocol | project | analysis | question
status: seed | active | synthesized | needs-review
updated: YYYY-MM-DD
source_paths:
  - raw/...
tags:
  - rag
  - agent
---
```

约定：

- `type`：页面类型
- `status`：
  - `seed`：初始整理，信息尚不完整
  - `active`：持续维护中
  - `synthesized`：跨来源综合较完整
  - `needs-review`：存在冲突、陈旧风险或待补证据
- `updated`：页面最近一次实质更新日期，使用 `YYYY-MM-DD`
- `source_paths`：本页依赖的原始资料路径列表

若某页不适合 frontmatter，可以省略，但 `source_paths` 与 `updated` 优先保留。

## 页面模板

### 来源页模板：`wiki/sources/...`

适用于 `raw/` 中单个文件或单个项目目录。

建议结构：

1. `一句话摘要`
2. `这份资料回答了什么问题`
3. `核心结论`
4. `关键证据 / 关键机制 / 关键结构`
5. `与现有 Wiki 的关系`
6. `未解决问题`
7. `来源`

说明：

- 来源页是“单源整理”，重点是把材料提纯，并指出它应该更新 Wiki 的哪些位置。
- 对于 PDF，尽量记录页码范围或章节范围。
- 对于 `raw/agent-platform/` 这类项目目录，不要默认给每个代码文件建页。优先先建一个项目页，再在必要时拆出组件页。

### 主题页模板：`wiki/topics/...`

适用于跨来源的大主题，如 RAG、GraphRAG、Agentic RAG、多 Agent 架构。

建议结构：

1. `主题定义`
2. `为什么重要`
3. `主流范式 / 设计空间`
4. `不同来源的共识`
5. `不同来源的分歧`
6. `当前仓库中的代表性来源`
7. `开放问题`
8. `相关页面`

### 协议 / 框架 / 项目页模板

适用于 A2A、Eino、agent-platform 这类“可长期引用的中层实体”。

建议结构：

1. `它是什么`
2. `核心抽象`
3. `关键组件或接口`
4. `典型工作流 / 数据流`
5. `适用场景`
6. `局限与风险`
7. `相关页面`
8. `来源`

## 工作流

### Ingest：把 `raw/` 资料吸收到 Wiki

每次 ingest 新资料时，遵循以下顺序：

1. 先读 `wiki/index.md`、`wiki/overview.md`、`wiki/log.md` 最近若干条记录。
2. 明确该资料属于哪一类：
   - 单篇 Markdown / PDF
   - 图示或 HTML 可视化
   - 项目目录 / 代码库
3. 在 `wiki/sources/` 中创建或更新对应来源页。
4. 找出应被影响的高层页面并更新：
   - `wiki/topics/...`
   - `wiki/concepts/...`
   - `wiki/frameworks/...`
   - `wiki/protocols/...`
   - `wiki/projects/...`
5. 若出现新的稳定主题、概念、协议、项目，再新增页面。
6. 更新 `wiki/index.md`。
7. 追加写入 `wiki/log.md`。

一次 ingest 允许触达多个 Wiki 页面。这是预期行为，不是过度编辑。

### Query：基于 Wiki 回答问题

回答问题时，优先用 Wiki，而不是每次都直接从 `raw/` 重建答案。

顺序：

1. 先读 `wiki/index.md`
2. 找到相关页面并阅读
3. 若 Wiki 已有足够信息，直接基于 Wiki 综合回答
4. 若 Wiki 缺信息或明显过时，再回到 `raw/` 补读
5. 若这次问答产出了长期有价值的内容，将其沉淀到 `wiki/analyses/` 或相关主题页

若这次问答还产生了阶段性成果、演示稿、图表、表格、候选方案或中间抽取结果，则同时写入 `outputs/` 对应子目录。

原则：

- 短答可以只在对话中回答。
- 如果答案具有“以后还会被复用”的价值，应文件化沉淀。
- “长期知识”进 `wiki/`，“阶段产物 / 对外交付 / 中间结果”进 `outputs/`。

### Lint：健康检查

应定期对 Wiki 做体检，重点检查：

- 是否存在没有被 `wiki/index.md` 收录的页面
- 是否存在没有任何内部链接指向的孤儿页
- 是否有页面长期没有更新，但已有新来源可补充
- 是否有冲突观点未被标注
- 是否出现大量被提及但仍无独立页面的概念
- 是否有重要结论缺少来源路径

发现问题时：

- 优先直接修复
- 若暂时不能修复，则记录到 `wiki/questions/` 或 `wiki/log.md`

## 索引与日志规范

### `wiki/index.md`

这是内容导向的目录页，不是流水账。

要求：

- 按类别列出 Wiki 页面
- 每个页面提供一句话说明
- 新增页面或页面定位发生变化时要更新
- Claude 在开始大多数任务前应先读这个文件

### `wiki/log.md`

这是时间导向的演化记录，采用追加式写法。

建议条目格式：

```md
## [YYYY-MM-DD] ingest | 资料标题
- Source: `raw/...`
- Added: [[某页面]]
- Updated: [[页面A]], [[页面B]]
- Notes: 一句话说明本次变化
```

也可记录：

- `query`
- `lint`
- `refactor`

保持格式稳定，方便后续用 `rg` / `grep` / `tail` 快速查看近期变化。

## 本仓库的主题优先级

结合当前 `raw/` 内容，优先围绕以下主题构建 Wiki：

### 一层主线

- RAG
- Agent
- Agentic RAG
- GraphRAG
- Multi-Agent Systems

### 二层核心对象

- Eino
- A2A
- MCP
- Harness
- `agent-platform`

### 三层关键概念

- Retrieval
- Planning
- Tool Use
- Memory
- Evaluation
- Hallucination Control
- Workflow Orchestration
- Agent Card / Task / Streaming / Async Collaboration

若后续资料重心变化，可更新这份优先级。

## 针对当前原始资料的特殊规则

### 1. Markdown 资料

- 这些通常已经是较高密度的整理稿。
- Claude 不应只做摘要，而应抽取结构、归类主题、识别与其他资料的连接点。

### 2. PDF 论文

- 优先提炼：问题定义、方法、实验结论、限制、与现有主题页的关系。
- 若能定位页码，尽量保留页码信息。
- 不要把论文页变成逐段翻译稿，重点是纳入 Wiki 结构。

### 3. SVG / HTML 可视化

- 将其视为“辅助理解材料”而非唯一事实来源。
- 可以总结图所表达的结构，但如果图本身缺乏上下文，不要夸大结论。

### 4. `raw/agent-platform/`

- 这是一个实现型资料源，而不是普通文档。
- 优先建立 `wiki/projects/agent-platform.md`，总结其架构、模块边界、Agent 划分、A2A / MCP / RAG 的实现位置。
- 默认不要为每个源码文件单独建页，除非该文件承载了一个稳定、可反复引用的核心机制。
- 当代码与文档叙述不一致时，必须显式记录差异。

### 5. 二进制与无关文件

- 如 `Jianying_37636880321514_installer.dmg` 这类文件，不属于核心知识沉淀对象。
- 除非用户明确要求，否则不执行、不分析安装行为，只在必要时记录其存在。

## 写作风格

- 优先写出“可复用的知识”，而不是聊天口吻摘要。
- 用词克制，少写空泛形容，多写结构、关系、边界、差异。
- 尽量让页面在脱离聊天上下文后依然可读。
- 如果信息不足，就明确写“证据不足 / 待验证”，不要装作已经确认。

## 成功标准

Claude 在这个仓库中工作的成功标准不是“答得像个聊天机器人”，而是：

- `wiki/` 持续变得更完整
- 页面之间的链接越来越密
- 问题的答案越来越能直接从 Wiki 中得出
- 新资料进入后，旧页面能被及时更新，而不是被遗忘
- 重要争议、缺口、推断都被显式记录

如果拿不准该怎么做，默认选择：

1. 保护 `raw/`
2. 更新 `wiki/`
3. 让知识更可追溯
4. 让结构比篇幅更优先
