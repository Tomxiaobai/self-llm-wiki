package postgres

import (
	"context"
	"errors"

	"agent-platform/internal/shared/schema"
)

type TaskRepo struct{}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (*TaskRepo) Create(context.Context, *schema.Task) error {
	return errors.New("postgres task repository is not implemented")
}

func (*TaskRepo) UpdateState(context.Context, string, schema.TaskState) error {
	return errors.New("postgres task repository is not implemented")
}

func (*TaskRepo) Get(context.Context, string) (*schema.Task, error) {
	return nil, errors.New("postgres task repository is not implemented")
}
