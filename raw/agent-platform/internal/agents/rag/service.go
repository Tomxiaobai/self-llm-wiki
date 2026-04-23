package rag

import (
	"context"

	"agent-platform/internal/shared/schema"
)

type Service struct {
	Retriever Retriever
}

func NewService(retriever Retriever) *Service {
	if retriever == nil {
		retriever = StubRetriever{}
	}
	return &Service{Retriever: retriever}
}

func (s *Service) Query(ctx context.Context, query string) ([]schema.Evidence, error) {
	return s.Retriever.Retrieve(ctx, query)
}
