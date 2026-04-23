package runner

import (
	"context"

	"agent-platform/internal/orchestrator/task"
)

type Runner struct {
	Tasks *task.Service
}

func (r Runner) RunQuery(ctx context.Context, query string) (*task.QueryResult, error) {
	return r.Tasks.HandleQuery(ctx, query)
}
