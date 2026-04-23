package rag

import (
	"context"
	"fmt"

	"agent-platform/internal/shared/schema"
)

type Retriever interface {
	Retrieve(ctx context.Context, query string) ([]schema.Evidence, error)
}

type StubRetriever struct{}

func (StubRetriever) Retrieve(_ context.Context, query string) ([]schema.Evidence, error) {
	return []schema.Evidence{
		{
			ID:        "doc-1",
			Title:     "GraphRAG Internal Note",
			SourceURI: "kb://graphrag/internal-note",
			Content:   fmt.Sprintf("Internal note related to: %s", query),
			Score:     0.93,
			External:  false,
		},
		{
			ID:        "doc-2",
			Title:     "Agentic RAG Internal Design",
			SourceURI: "kb://agentic-rag/design",
			Content:   "Agentic RAG combines retrieval with multi-step decision making.",
			Score:     0.88,
			External:  false,
		},
	}, nil
}
