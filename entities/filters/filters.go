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

package filters

import (
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
)

type Operator int

const (
	OperatorEqual            Operator = 1
	OperatorNotEqual         Operator = 2
	OperatorGreaterThan      Operator = 3
	OperatorGreaterThanEqual Operator = 4
	OperatorLessThan         Operator = 5
	OperatorLessThanEqual    Operator = 6
	OperatorAnd              Operator = 7
	OperatorOr               Operator = 8
	OperatorNot              Operator = 9
	OperatorWithinGeoRange   Operator = 10
	OperatorLike             Operator = 11
)

func (o Operator) OnValue() bool {
	switch o {
	case OperatorEqual,
		OperatorNotEqual,
		OperatorGreaterThan,
		OperatorGreaterThanEqual,
		OperatorLessThan,
		OperatorLessThanEqual,
		OperatorWithinGeoRange,
		OperatorLike:
		return true
	default:
		return false
	}
}

func (o Operator) Name() string {
	switch o {
	case OperatorEqual:
		return "Equal"
	case OperatorNotEqual:
		return "NotEqual"
	case OperatorGreaterThan:
		return "GreaterThan"
	case OperatorGreaterThanEqual:
		return "GreaterThanEqual"
	case OperatorLessThan:
		return "LessThan"
	case OperatorLessThanEqual:
		return "LessThanEqual"
	case OperatorAnd:
		return "And"
	case OperatorOr:
		return "Or"
	case OperatorNot:
		return "Not"
	case OperatorWithinGeoRange:
		return "WithinGeoRange"
	case OperatorLike:
		return "Like"
	default:
		panic("Unknown operator")
	}
}

type LocalFilter struct {
	Root *Clause `json:"root"`
}

type Value struct {
	Value interface{}     `json:"value"`
	Type  schema.DataType `json:"type"`
}

type Clause struct {
	Operator Operator `json:"operator"`
	On       *Path    `json:"on"`
	Value    *Value   `json:"value"`
	Operands []Clause `json:"operands"`
}

// GeoRange to be used with fields of type GeoCoordinates. Identifies a point
// and a maximum distance from that point.
type GeoRange struct {
	*models.GeoCoordinates
	Distance float32 `json:"distance"`
}
