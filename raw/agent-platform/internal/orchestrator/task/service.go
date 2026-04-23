package task

import (
	"context"
	"fmt"

	"agent-platform/internal/harness/eval"
	"agent-platform/internal/harness/trace"
	"agent-platform/internal/shared/schema"
)

type Service struct {
	Tasks     TaskRepository
	Artifacts ArtifactRepository
	Gateway   Gateway
	Evaluator eval.Evaluator
	Tracer    trace.Tracer
}

func NewService(
	tasks TaskRepository,
	artifacts ArtifactRepository,
	gateway Gateway,
	evaluator eval.Evaluator,
	tracer trace.Tracer,
) *Service {
	if evaluator == nil {
		evaluator = eval.SimpleEvaluator{}
	}
	if tracer == nil {
		defaultTracer := trace.NewNopTracer()
		tracer = defaultTracer
	}

	return &Service{
		Tasks:     tasks,
		Artifacts: artifacts,
		Gateway:   gateway,
		Evaluator: evaluator,
		Tracer:    tracer,
	}
}

func (s *Service) HandleQuery(ctx context.Context, query string) (*QueryResult, error) {
	task := schema.NewTask(query)
	if err := s.Tasks.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	if err := s.Tasks.UpdateState(ctx, task.ID, schema.TaskWorking); err != nil {
		return nil, fmt.Errorf("mark task working: %w", err)
	}

	ctx, span := s.Tracer.Start(ctx, "task.handle_query")
	defer s.Tracer.End(ctx, span)

	answer, err := s.Gateway.Execute(ctx, task)
	if err != nil {
		_ = s.Tasks.UpdateState(ctx, task.ID, schema.TaskFailed)
		return nil, fmt.Errorf("gateway execution failed: %w", err)
	}

	if err := s.Artifacts.Save(ctx, &schema.Artifact{
		ID:      fmt.Sprintf("artifact-%s-answer", task.ID),
		TaskID:  task.ID,
		Type:    schema.ArtifactAnswer,
		Content: answer.Text,
		Metadata: map[string]any{
			"citations": len(answer.Citations),
		},
	}); err != nil {
		return nil, fmt.Errorf("save artifact: %w", err)
	}

	evalResult, err := s.Evaluator.Evaluate(ctx, task.ID, answer.Text)
	if err != nil {
		return nil, fmt.Errorf("evaluate answer: %w", err)
	}

	if err := s.Tasks.UpdateState(ctx, task.ID, schema.TaskCompleted); err != nil {
		return nil, fmt.Errorf("mark task completed: %w", err)
	}

	return &QueryResult{
		TaskID:     task.ID,
		Answer:     answer.Text,
		Citations:  answer.Citations,
		TraceID:    span.TraceID,
		EvalPassed: evalResult.Passed,
	}, nil
}
