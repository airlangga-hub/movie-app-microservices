package model

import model "github.com/airlangga-hub/movie-app-microservices/metadata/pkg/model"

type MovieDetails struct {
	Rating   *float64       `json:",omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
