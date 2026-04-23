package client

import (
	"context"

	"agent-platform/internal/a2a/card"
)

type MessageRequest struct {
	Query string
}

type TaskResponse struct {
	TaskID string
	State  string
	Output string
}

type Client interface {
	Discover(ctx context.Context, baseURL string) (*card.AgentCard, error)
	Send(ctx context.Context, baseURL string, req MessageRequest) (*TaskResponse, error)
}

type NopClient struct{}

func (NopClient) Discover(_ context.Context, baseURL string) (*card.AgentCard, error) {
	return &card.AgentCard{
		Name:        "stub-agent",
		Description: "Local stub A2A client result",
		URL:         baseURL,
	}, nil
}

func (NopClient) Send(_ context.Context, _ string, req MessageRequest) (*TaskResponse, error) {
	return &TaskResponse{
		TaskID: "stub-task",
		State:  "completed",
		Output: req.Query,
	}, nil
}
