package schema

type Evidence struct {
	ID        string
	Title     string
	SourceURI string
	Content   string
	Score     float64
	External  bool
}

type Answer struct {
	Text             string
	Citations        []Citation
	InternalEvidence []Evidence
	ExternalEvidence []Evidence
}
