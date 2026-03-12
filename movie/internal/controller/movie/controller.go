package movie

import (
	"context"
	"errors"
	
	ratingmodel "github.com/airlangga-hub/movie-app-microservices/rating/pkg/model"
	metadatamodel "github.com/airlangga-hub/movie-app-microservices/metadata/pkg/model"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordType ratingmodel.RecordType, recordID ratingmodel.RecordID) (float64, error)
	PutRating(ctx context.Context, recordType ratingmodel.RecordType,  recordID ratingmodel.RecordID, rating *ratingmodel.Rating) error
}
type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}
