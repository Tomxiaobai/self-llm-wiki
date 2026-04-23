package schema

import (
	"fmt"
	"sync/atomic"
	"time"
)

type TaskState string

const (
	TaskSubmitted TaskState = "submitted"
	TaskWorking   TaskState = "working"
	TaskCompleted TaskState = "completed"
	TaskFailed    TaskState = "failed"
	TaskCanceled  TaskState = "canceled"
)

type Task struct {
	ID        string
	ContextID string
	Intent    string
	Query     string
	State     TaskState
	CreatedAt int64
	UpdatedAt int64
}

var taskCounter uint64

func NewTask(query string) *Task {
	now := time.Now().Unix()
	return &Task{
		ID:        nextID("task", &taskCounter),
		ContextID: nextID("ctx", &taskCounter),
		Intent:    "research",
		Query:     query,
		State:     TaskSubmitted,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func nextID(prefix string, counter *uint64) string {
	return fmt.Sprintf("%s-%d-%d", prefix, time.Now().UnixNano(), atomic.AddUint64(counter, 1))
}
