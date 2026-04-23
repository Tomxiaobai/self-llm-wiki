package hooks

import "context"

type Context struct {
	TaskID string
	Node   string
}

type Hook interface {
	Before(context.Context, Context) error
	After(context.Context, Context, any) error
	OnError(context.Context, Context, error) error
}

type NopHook struct{}

func (NopHook) Before(context.Context, Context) error         { return nil }
func (NopHook) After(context.Context, Context, any) error     { return nil }
func (NopHook) OnError(context.Context, Context, error) error { return nil }

type Chain struct {
	hooks []Hook
}

func NewChain(hooks ...Hook) Chain {
	return Chain{hooks: hooks}
}

func (c Chain) Empty() bool {
	return len(c.hooks) == 0
}

func (c Chain) Before(ctx context.Context, hc Context) error {
	for _, hook := range c.hooks {
		if err := hook.Before(ctx, hc); err != nil {
			return err
		}
	}
	return nil
}

func (c Chain) After(ctx context.Context, hc Context, result any) error {
	for _, hook := range c.hooks {
		if err := hook.After(ctx, hc, result); err != nil {
			return err
		}
	}
	return nil
}

func (c Chain) OnError(ctx context.Context, hc Context, err error) error {
	for _, hook := range c.hooks {
		if hookErr := hook.OnError(ctx, hc, err); hookErr != nil {
			return hookErr
		}
	}
	return nil
}
