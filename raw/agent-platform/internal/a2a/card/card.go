package card

type AgentCard struct {
	Name        string
	Description string
	URL         string
	Skills      []Skill
}

type Skill struct {
	ID          string
	Name        string
	Description string
}
