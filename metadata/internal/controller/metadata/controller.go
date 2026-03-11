package metadata

import (
	"context"
	"errors"

	"github.com/airlangga-hub/movie-app-microservices/metadata/internal/repository"
	model "github.com/airlangga-hub/movie-app-microservices/metadata/pkg"
)

var ErrNotFound = errors.New("not found.")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

type Controller struct {
	repo metadataRepository
}

func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}