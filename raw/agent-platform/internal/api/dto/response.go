package dto

import "agent-platform/internal/shared/schema"

type ChatResponse struct {
	TaskID     string            `json:"task_id"`
	Answer     string            `json:"answer"`
	Citations  []schema.Citation `json:"citations"`
	TraceID    string            `json:"trace_id"`
	EvalPassed bool              `json:"eval_passed"`
}
