package memory

import (
	"context"

	"github.com/airlangga-hub/movie-app-microservices/rating/internal/repository"
	"github.com/airlangga-hub/movie-app-microservices/rating/pkg/model"
)

type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

func New() *Repository {
	return &Repository{data: map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

func (r *Repository) Get(ctx context.Context, recordType model.RecordType, recordID model.RecordID) ([]model.Rating, error) {
	ratings, ok := r.data[recordType][recordID]
	if !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return ratings, nil
}