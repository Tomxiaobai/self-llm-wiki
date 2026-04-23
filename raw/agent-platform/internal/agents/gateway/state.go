package gateway

import "agent-platform/internal/shared/schema"

type State struct {
	TaskID          string
	Query           string
	NormalizedQuery string
	NeedSearch      bool
	InternalDocs    []schema.Evidence
	ExternalDocs    []schema.Evidence
}
