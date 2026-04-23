package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"agent-platform/internal/api/dto"
	"agent-platform/internal/orchestrator/task"
)

type QueryService interface {
	HandleQuery(ctx context.Context, query string) (*task.QueryResult, error)
}

func NewChatHandler(svc QueryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := svc.HandleQuery(r.Context(), req.Query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := dto.ChatResponse{
			TaskID:     result.TaskID,
			Answer:     result.Answer,
			Citations:  result.Citations,
			TraceID:    result.TraceID,
			EvalPassed: result.EvalPassed,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}
