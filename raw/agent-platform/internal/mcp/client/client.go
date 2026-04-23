package client

import "context"

type Client interface {
	Call(ctx context.Context, tool string, args map[string]any) (map[string]any, error)
}

type LocalClient struct{}

func (LocalClient) Call(_ context.Context, tool string, args map[string]any) (map[string]any, error) {
	return map[string]any{
		"tool": tool,
		"args": args,
	}, nil
}
