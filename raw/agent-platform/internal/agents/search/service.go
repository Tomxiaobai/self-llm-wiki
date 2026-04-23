package search

import (
	"context"
	"fmt"

	"agent-platform/internal/shared/schema"
)

type Service struct {
	Cleaner Cleaner
}

func NewService(cleaner Cleaner) *Service {
	if cleaner == nil {
		cleaner = BasicCleaner{}
	}
	return &Service{Cleaner: cleaner}
}

func (s *Service) Search(_ context.Context, query string) ([]schema.Evidence, error) {
	content := s.Cleaner.Clean(fmt.Sprintf("External web result related to %s", query))
	return []schema.Evidence{
		{
			ID:        "web-1",
			Title:     "External Search Result",
			SourceURI: "https://example.com/search-result",
			Content:   content,
			Score:     0.79,
			External:  true,
		},
	}, nil
}
