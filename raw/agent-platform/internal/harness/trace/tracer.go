package trace

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Span struct {
	TraceID string
	Name    string
	StartAt time.Time
}

type Tracer interface {
	Start(ctx context.Context, name string) (context.Context, Span)
	Event(ctx context.Context, span Span, key string, value any)
	End(ctx context.Context, span Span)
}

type NopTracer struct{}

var traceCounter uint64

func NewNopTracer() NopTracer {
	return NopTracer{}
}

func (NopTracer) Start(ctx context.Context, name string) (context.Context, Span) {
	return ctx, Span{
		TraceID: fmt.Sprintf("trace-%d-%d", time.Now().UnixNano(), atomic.AddUint64(&traceCounter, 1)),
		Name:    name,
		StartAt: time.Now(),
	}
}

func (NopTracer) Event(context.Context, Span, string, any) {}

func (NopTracer) End(context.Context, Span) {}
