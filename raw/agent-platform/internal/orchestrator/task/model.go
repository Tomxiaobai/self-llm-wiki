package task

import (
	"context"

	"agent-platform/internal/shared/schema"
)

type Gateway interface {
	Execute(ctx context.Context, task *schema.Task) (*schema.Answer, error)
}

type TaskRepository interface {
	Create(ctx context.Context, task *schema.Task) error
	UpdateState(ctx context.Context, taskID string, state schema.TaskState) error
	Get(ctx context.Context, taskID string) (*schema.Task, error)
}

type ArtifactRepository interface {
	Save(ctx context.Context, artifact *schema.Artifact) error
	ListByTask(ctx context.Context, taskID string) ([]schema.Artifact, error)
}

type QueryResult struct {
	TaskID     string
	Answer     string
	Citations  []schema.Citation
	TraceID    string
	EvalPassed bool
}
