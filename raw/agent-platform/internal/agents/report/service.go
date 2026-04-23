package report

import (
	"context"

	"agent-platform/internal/shared/schema"
)

type BuildInput struct {
	Query    string
	Internal []schema.Evidence
	External []schema.Evidence
}

type Service struct {
	Template string
}

func NewService(template string) *Service {
	if template == "" {
		template = DefaultTemplate
	}
	return &Service{Template: template}
}

func (s *Service) Build(_ context.Context, input BuildInput) (*schema.Answer, error) {
	citations := make([]schema.Citation, 0, len(input.Internal)+len(input.External))
	for _, item := range input.Internal {
		citations = append(citations, schema.Citation{
			SourceType: "internal",
			SourceURI:  item.SourceURI,
			ChunkID:    item.ID,
			SpanText:   item.Title,
		})
	}
	for _, item := range input.External {
		citations = append(citations, schema.Citation{
			SourceType: "external",
			SourceURI:  item.SourceURI,
			ChunkID:    item.ID,
			SpanText:   item.Title,
		})
	}

	return &schema.Answer{
		Text:             RenderDefault(input.Query, input.Internal, input.External, citations),
		Citations:        citations,
		InternalEvidence: input.Internal,
		ExternalEvidence: input.External,
	}, nil
}
