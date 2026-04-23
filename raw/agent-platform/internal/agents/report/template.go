package report

import (
	"fmt"
	"strings"

	"agent-platform/internal/shared/schema"
)

const DefaultTemplate = `# 研究结论

## 问题
%s

## 核心结论
%s

## 内部资料观点
%s

## 外部资料补充
%s

## 引用列表
%s
`

func RenderDefault(query string, internal []schema.Evidence, external []schema.Evidence, citations []schema.Citation) string {
	internalSection := joinEvidence(internal)
	externalSection := joinEvidence(external)
	conclusion := "基于当前证据，建议优先采用内部知识结论，并用外部资料做补充校验。"
	citationSection := joinCitations(citations)
	return fmt.Sprintf(DefaultTemplate, query, conclusion, internalSection, externalSection, citationSection)
}

func joinEvidence(items []schema.Evidence) string {
	if len(items) == 0 {
		return "暂无"
	}

	lines := make([]string, 0, len(items))
	for _, item := range items {
		lines = append(lines, fmt.Sprintf("- %s: %s", item.Title, item.Content))
	}
	return strings.Join(lines, "\n")
}

func joinCitations(items []schema.Citation) string {
	if len(items) == 0 {
		return "暂无"
	}

	lines := make([]string, 0, len(items))
	for _, item := range items {
		lines = append(lines, fmt.Sprintf("- %s (%s)", item.SourceURI, item.SpanText))
	}
	return strings.Join(lines, "\n")
}
