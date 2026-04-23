package postgres

import (
	"context"
	"errors"

	"agent-platform/internal/shared/schema"
)

type ArtifactRepo struct{}

func NewArtifactRepo() *ArtifactRepo {
	return &ArtifactRepo{}
}

func (*ArtifactRepo) Save(context.Context, *schema.Artifact) error {
	return errors.New("postgres artifact repository is not implemented")
}

func (*ArtifactRepo) ListByTask(context.Context, string) ([]schema.Artifact, error) {
	return nil, errors.New("postgres artifact repository is not implemented")
}
