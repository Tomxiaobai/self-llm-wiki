package gateway

import (
	"context"
	"fmt"

	"agent-platform/internal/agents/rag"
	reportagent "agent-platform/internal/agents/report"
	searchagent "agent-platform/internal/agents/search"
	"agent-platform/internal/harness/hooks"
	"agent-platform/internal/harness/policy"
	"agent-platform/internal/harness/trace"
	"agent-platform/internal/mcp/registry"
	"agent-platform/internal/shared/schema"
)

type RAGAgent interface {
	Query(ctx context.Context, query string) ([]schema.Evidence, error)
}

type SearchAgent interface {
	Search(ctx context.Context, query string) ([]schema.Evidence, error)
}

type ReportAgent interface {
	Build(ctx context.Context, input reportagent.BuildInput) (*schema.Answer, error)
}

type Service struct {
	RAG      RAGAgent
	Search   SearchAgent
	Report   ReportAgent
	Registry registry.Registry
	Policy   policy.Engine
	Hooks    hooks.Chain
	Tracer   trace.Tracer
}

func NewService(
	ragAgent RAGAgent,
	searchAgent SearchAgent,
	reportAgent ReportAgent,
	reg registry.Registry,
	policyEngine policy.Engine,
	hookChain hooks.Chain,
	tracer trace.Tracer,
) *Service {
	if ragAgent == nil {
		ragAgent = rag.NewService(nil)
	}
	if searchAgent == nil {
		searchAgent = searchagent.NewService(nil)
	}
	if reportAgent == nil {
		reportAgent = reportagent.NewService("")
	}
	if reg == nil {
		defaultRegistry := registry.NewMemoryRegistry()
		reg = defaultRegistry
	}
	if policyEngine == nil {
		defaultPolicy := policy.NewAllowAllEngine()
		policyEngine = defaultPolicy
	}
	if tracer == nil {
		defaultTracer := trace.NewNopTracer()
		tracer = defaultTracer
	}
	if hookChain.Empty() {
		hookChain = hooks.NewChain(hooks.NopHook{})
	}

	return &Service{
		RAG:      ragAgent,
		Search:   searchAgent,
		Report:   reportAgent,
		Registry: reg,
		Policy:   policyEngine,
		Hooks:    hookChain,
		Tracer:   tracer,
	}
}

func (s *Service) Execute(ctx context.Context, task *schema.Task) (*schema.Answer, error) {
	ctx, span := s.Tracer.Start(ctx, "gateway.execute")
	defer s.Tracer.End(ctx, span)

	hookCtx := hooks.Context{TaskID: task.ID, Node: "gateway"}
	if err := s.Hooks.Before(ctx, hookCtx); err != nil {
		return nil, err
	}

	internalDocs, err := s.RAG.Query(ctx, task.Query)
	if err != nil {
		_ = s.Hooks.OnError(ctx, hookCtx, err)
		return nil, fmt.Errorf("rag query failed: %w", err)
	}

	externalDocs := make([]schema.Evidence, 0)
	if NeedExternalSearch(task.Query, internalDocs) {
		decision := s.Policy.CheckAgent("search-agent")
		if !decision.Allowed {
			return nil, fmt.Errorf("search-agent blocked: %s", decision.Reason)
		}

		externalDocs, err = s.Search.Search(ctx, task.Query)
		if err != nil {
			_ = s.Hooks.OnError(ctx, hookCtx, err)
			return nil, fmt.Errorf("search query failed: %w", err)
		}
	}

	answer, err := s.Report.Build(ctx, reportagent.BuildInput{
		Query:    task.Query,
		Internal: internalDocs,
		External: externalDocs,
	})
	if err != nil {
		_ = s.Hooks.OnError(ctx, hookCtx, err)
		return nil, fmt.Errorf("report build failed: %w", err)
	}

	if err := s.Hooks.After(ctx, hookCtx, answer); err != nil {
		return nil, err
	}

	return answer, nil
}
