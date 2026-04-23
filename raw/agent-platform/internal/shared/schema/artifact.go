package schema

type ArtifactType string

const (
	ArtifactAnswer ArtifactType = "answer"
	ArtifactReport ArtifactType = "report"
	ArtifactTrace  ArtifactType = "trace"
	ArtifactEval   ArtifactType = "eval"
)

type Artifact struct {
	ID       string
	TaskID   string
	Type     ArtifactType
	Content  string
	Metadata map[string]any
}
