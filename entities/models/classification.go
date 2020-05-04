//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Classification Manage classifications, trigger them and view status of past classifications.
// swagger:model Classification
type Classification struct {

	// base the text-based classification on these fields (of type text)
	BasedOnProperties []string `json:"basedOnProperties"`

	// class (name) which is used in this classification
	Class string `json:"class,omitempty"`

	// which ref-property to set as part of the classification
	ClassifyProperties []string `json:"classifyProperties"`

	// error message if status == failed
	Error string `json:"error,omitempty"`

	// ID to uniquely identify this classification run
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Only available on type=contextual. All words in a source corpus are ranked by their information gain against the possible target objects. A cutoff percentile of 40 implies that the top 40% are used and the bottom 60% are cut-off.
	InformationGainCutoffPercentile *int32 `json:"informationGainCutoffPercentile,omitempty"`

	// Only available on type=contextual. Words in a corpus will receive an additional boost based on how high they are ranked according to information gain. Setting this value to 3 implies that the top-ranked word will be ranked 3 times as high as the bottom ranked word. The curve in between is logarithmic. A maximum boost of 1 implies that no boosting occurs.
	InformationGainMaximumBoost *int32 `json:"informationGainMaximumBoost,omitempty"`

	// k-value when using k-Neareast-Neighbor
	K *int32 `json:"k,omitempty"`

	// additional meta information about the classification
	Meta *ClassificationMeta `json:"meta,omitempty"`

	// Only available on type=contextual. Both IG and tf-idf are mechanisms to remove words from the corpora. However, on very short corpora this could lead to a removal of all words, or all but a single word. This value guarantees that - regardless of tf-idf and IG score - always at least n words are used.
	MinimumUsableWords *int32 `json:"minimumUsableWords,omitempty"`

	// limit the objects to be classified
	SourceWhere *WhereFilter `json:"sourceWhere,omitempty"`

	// status of this classification
	// Enum: [running completed failed]
	Status string `json:"status,omitempty"`

	// Limit the possible sources when using an algorithm which doesn't really on trainig data, e.g. 'contextual'. When using an algorithm with a training set, such as 'knn', limit the training set instead
	TargetWhere *WhereFilter `json:"targetWhere,omitempty"`

	// Only available on type=contextual. All words in a corpus are ranked by their tf-idf score. A cutoff percentile of 80 implies that the top 80% are used and the bottom 20% are cut-off. This is very effective to remove words that occur in almost all objects, such as filler and stop words.
	TfidfCutoffPercentile *int32 `json:"tfidfCutoffPercentile,omitempty"`

	// Limit the training objects to be considered during the classification. Can only be used on types with explicit training sets, such as 'knn'
	TrainingSetWhere *WhereFilter `json:"trainingSetWhere,omitempty"`

	// which algorythim to use for classifications
	// Enum: [knn contextual]
	Type *string `json:"type,omitempty"`
}

// Validate validates this classification
func (m *Classification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceWhere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetWhere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainingSetWhere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Classification) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Classification) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *Classification) validateSourceWhere(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceWhere) { // not required
		return nil
	}

	if m.SourceWhere != nil {
		if err := m.SourceWhere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceWhere")
			}
			return err
		}
	}

	return nil
}

var classificationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["running","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		classificationTypeStatusPropEnum = append(classificationTypeStatusPropEnum, v)
	}
}

const (

	// ClassificationStatusRunning captures enum value "running"
	ClassificationStatusRunning string = "running"

	// ClassificationStatusCompleted captures enum value "completed"
	ClassificationStatusCompleted string = "completed"

	// ClassificationStatusFailed captures enum value "failed"
	ClassificationStatusFailed string = "failed"
)

// prop value enum
func (m *Classification) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, classificationTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Classification) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Classification) validateTargetWhere(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetWhere) { // not required
		return nil
	}

	if m.TargetWhere != nil {
		if err := m.TargetWhere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetWhere")
			}
			return err
		}
	}

	return nil
}

func (m *Classification) validateTrainingSetWhere(formats strfmt.Registry) error {

	if swag.IsZero(m.TrainingSetWhere) { // not required
		return nil
	}

	if m.TrainingSetWhere != nil {
		if err := m.TrainingSetWhere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trainingSetWhere")
			}
			return err
		}
	}

	return nil
}

var classificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["knn","contextual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		classificationTypeTypePropEnum = append(classificationTypeTypePropEnum, v)
	}
}

const (

	// ClassificationTypeKnn captures enum value "knn"
	ClassificationTypeKnn string = "knn"

	// ClassificationTypeContextual captures enum value "contextual"
	ClassificationTypeContextual string = "contextual"
)

// prop value enum
func (m *Classification) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, classificationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Classification) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Classification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Classification) UnmarshalBinary(b []byte) error {
	var res Classification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
