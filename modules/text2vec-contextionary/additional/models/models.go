//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package models

type FeatureProjection struct {
	Vector []float32 `json:"vector"`
}

type NearestNeighbors struct {
	Neighbors []*NearestNeighbor `json:"neighbors"`
}

type NearestNeighbor struct {
	Concept  string    `json:"concept,omitempty"`
	Distance float32   `json:"distance,omitempty"`
	Vector   []float32 `json:"vector"`
}

type SemanticPath struct {
	Path []*SemanticPathElement `json:"path"`
}

type SemanticPathElement struct {
	Concept            string   `json:"concept,omitempty"`
	DistanceToNext     *float32 `json:"distanceToNext,omitempty"`
	DistanceToPrevious *float32 `json:"distanceToPrevious,omitempty"`
	DistanceToQuery    float32  `json:"distanceToQuery,omitempty"`
	DistanceToResult   float32  `json:"distanceToResult,omitempty"`
}
