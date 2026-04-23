package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"agent-platform/internal/shared/schema"
)

type Store struct {
	mu        sync.RWMutex
	tasks     map[string]*schema.Task
	artifacts map[string][]schema.Artifact
}

func NewStore() *Store {
	return &Store{
		tasks:     make(map[string]*schema.Task),
		artifacts: make(map[string][]schema.Artifact),
	}
}

func (s *Store) Create(_ context.Context, task *schema.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[task.ID]; exists {
		return fmt.Errorf("task %q already exists", task.ID)
	}
	copyTask := *task
	s.tasks[task.ID] = &copyTask
	return nil
}

func (s *Store) UpdateState(_ context.Context, taskID string, state schema.TaskState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return fmt.Errorf("task %q not found", taskID)
	}
	task.State = state
	task.UpdatedAt = time.Now().Unix()
	return nil
}

func (s *Store) Get(_ context.Context, taskID string) (*schema.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return nil, fmt.Errorf("task %q not found", taskID)
	}
	copyTask := *task
	return &copyTask, nil
}

func (s *Store) Save(_ context.Context, artifact *schema.Artifact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.artifacts[artifact.TaskID] = append(s.artifacts[artifact.TaskID], *artifact)
	return nil
}

func (s *Store) ListByTask(_ context.Context, taskID string) ([]schema.Artifact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := s.artifacts[taskID]
	out := make([]schema.Artifact, len(items))
	copy(out, items)
	return out, nil
}
