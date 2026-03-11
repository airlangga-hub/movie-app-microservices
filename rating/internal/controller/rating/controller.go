package rating

import (
	"context"
	"errors"

	"github.com/airlangga-hub/movie-app-microservices/rating/internal/repository"
	"github.com/airlangga-hub/movie-app-microservices/rating/pkg/model"
)

var ErrNotFound = errors.New("ratings not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordType model.RecordType, recordID model.RecordID) ([]model.Rating, error)
	Put(ctx context.Context, recordType model.RecordType, recordID model.RecordID, rating *model.Rating) error
}

type Controller struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, recordType model.RecordType, recordID model.RecordID) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordType, recordID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, err
	}
	
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	
	return sum / float64(len(ratings)), nil
}

func (c *Controller) PutRating(ctx context.Context, recordType model.RecordType, recordID model.RecordID, rating *model.Rating) error {
	return c.repo.Put(ctx, recordType, recordID, rating)
}