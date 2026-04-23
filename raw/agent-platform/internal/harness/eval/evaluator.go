package eval

import (
	"context"
	"strings"
)

type Result struct {
	GroundedScore float64
	CitationScore float64
	Passed        bool
}

type Evaluator interface {
	Evaluate(ctx context.Context, taskID string, answer string) (Result, error)
}

type SimpleEvaluator struct{}

func (SimpleEvaluator) Evaluate(_ context.Context, _ string, answer string) (Result, error) {
	passed := strings.TrimSpace(answer) != ""
	score := 0.0
	if passed {
		score = 1.0
	}
	return Result{
		GroundedScore: score,
		CitationScore: score,
		Passed:        passed,
	}, nil
}
