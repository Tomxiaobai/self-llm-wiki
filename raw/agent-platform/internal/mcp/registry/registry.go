package registry

import "fmt"

type Tool struct {
	Name        string
	Description string
	Scope       string
	RiskLevel   string
}

type Registry interface {
	List(scope string) []Tool
	Exists(name string) bool
	Validate(name string, args map[string]any) error
}

type MemoryRegistry struct {
	tools []Tool
}

func NewMemoryRegistry() MemoryRegistry {
	return MemoryRegistry{
		tools: []Tool{
			{Name: "retrieve_documents", Description: "Retrieve internal documents", Scope: "retrieve", RiskLevel: "low"},
			{Name: "rerank_documents", Description: "Rerank candidate evidence", Scope: "retrieve", RiskLevel: "low"},
			{Name: "parse_document", Description: "Parse raw document content", Scope: "document", RiskLevel: "low"},
			{Name: "build_citations", Description: "Build citations from evidence", Scope: "cite", RiskLevel: "low"},
		},
	}
}

func (r MemoryRegistry) List(scope string) []Tool {
	if scope == "" {
		return append([]Tool(nil), r.tools...)
	}

	out := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		if tool.Scope == scope {
			out = append(out, tool)
		}
	}
	return out
}

func (r MemoryRegistry) Exists(name string) bool {
	for _, tool := range r.tools {
		if tool.Name == name {
			return true
		}
	}
	return false
}

func (r MemoryRegistry) Validate(name string, _ map[string]any) error {
	if !r.Exists(name) {
		return fmt.Errorf("tool %q is not registered", name)
	}
	return nil
}
