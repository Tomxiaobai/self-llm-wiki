package llm

import "context"

type ChatModel interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

type EchoModel struct{}

func (EchoModel) Generate(_ context.Context, prompt string) (string, error) {
	return prompt, nil
}
