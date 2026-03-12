package movie

import (
	"context"
	"errors"

	metadatamodel "github.com/airlangga-hub/movie-app-microservices/metadata/pkg/model"
	"github.com/airlangga-hub/movie-app-microservices/movie/internal/gateway"
	"github.com/airlangga-hub/movie-app-microservices/movie/pkg/model"
	ratingmodel "github.com/airlangga-hub/movie-app-microservices/rating/pkg/model"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordType ratingmodel.RecordType, recordID ratingmodel.RecordID) (float64, error)
	PutRating(ctx context.Context, recordType ratingmodel.RecordType, recordID ratingmodel.RecordID, rating *ratingmodel.Rating) error
}
type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway: ratingGateway, metadataGateway: metadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}

	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordTypeMovie, ratingmodel.RecordID(id))
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
