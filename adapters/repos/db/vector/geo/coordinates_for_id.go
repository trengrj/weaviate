package geo

import (
	"context"
	"fmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// CoordinatesForID must provide the geo coordinates for the specified index
// id
type CoordinatesForID func(ctx context.Context, id int32) (models.GeoCoordinates, error)

// VectorForID transforms the geo coordinates into a "vector" of fixed length
// two, where element 0 represents the latitude and element 1 represents the
// longitude. This way it is usable by a generic vector index such as HNSW
func (cfid CoordinatesForID) VectorForID(ctx context.Context, id int32) ([]float32, error) {
	coordinates, err := cfid(ctx, id)
	if err != nil {
		return nil, err
	}

	return geoCoordiantesToVector(coordinates)
}

func geoCoordiantesToVector(in models.GeoCoordinates) ([]float32, error) {
	if in.Latitude == nil {
		return nil, fmt.Errorf("latitude must be set")
	}

	if in.Longitude == nil {
		return nil, fmt.Errorf("longitude must be set")
	}

	return []float32{*in.Latitude, *in.Longitude}, nil
}
