package movie

import (
	"context"
	"errors"

	metadatamodel "github.com/airlangga-hub/movie-app-microservices/metadata/pkg/model"
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
