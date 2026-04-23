package policy

type Decision struct {
	Allowed bool
	Reason  string
}

type Engine interface {
	CheckTool(tool string) Decision
	CheckAgent(target string) Decision
}

type StaticEngine struct {
	AllowedTools  map[string]bool
	AllowedAgents map[string]bool
}

func NewAllowAllEngine() StaticEngine {
	return StaticEngine{}
}

func (e StaticEngine) CheckTool(tool string) Decision {
	if len(e.AllowedTools) == 0 || e.AllowedTools[tool] {
		return Decision{Allowed: true, Reason: "tool allowed"}
	}
	return Decision{Allowed: false, Reason: "tool blocked by policy"}
}

func (e StaticEngine) CheckAgent(target string) Decision {
	if len(e.AllowedAgents) == 0 || e.AllowedAgents[target] {
		return Decision{Allowed: true, Reason: "agent allowed"}
	}
	return Decision{Allowed: false, Reason: "agent blocked by policy"}
}
