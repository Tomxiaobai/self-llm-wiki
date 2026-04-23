# Outputs Index

`outputs/` 是这个仓库的产出层，用来放当前任务派生出的中间产物、分析稿、图表、演示稿和结构化导出结果。

它和 `wiki/` 的分工不同：

- `wiki/`：长期知识、主题页、来源页、稳定结论
- `outputs/`：当前任务的派生产物、阶段性交付物、中间抽取结果

## 目录约定

- `analyses/`：一次性分析稿、阶段总结、候选结论
- `slides/`：Marp / PPT / 演示稿
- `charts/`：图表和可视化导出
- `tables/`：CSV / TSV / JSON / 结构化结果
- `tmp/`：临时文本抽取、OCR、脚本中间文件

## 当前已有产物

- [analyses/2026-04-23-raw-to-wiki-build-summary.md](./analyses/2026-04-23-raw-to-wiki-build-summary.md)

## 使用规则

- 若某份产物具有长期复用价值，应把其结论再沉淀回 `wiki/`
- `tmp/` 内文件允许覆盖和清理，不应视为最终知识资产
