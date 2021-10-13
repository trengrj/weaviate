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

// Code generated by go generate; DO NOT EDIT.
// This file was generated by go generate ./deprecations at 2021-10-13
package deprecations

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
)

func timeMust(t time.Time, err error) strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	return strfmt.DateTime(t)
}

func timeMustPtr(t time.Time, err error) *strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	parsed := strfmt.DateTime(t)
	return &parsed
}

func ptString(in string) *string {
	return &in
}

var ByID = map[string]models.Deprecation{
	"rest-meta-prop": models.Deprecation{
		ID:           "rest-meta-prop",
		Status:       "deprecated",
		APIType:      "REST",
		Mitigation:   "Use ?include=<propName>, e.g. ?include=_classification for classification meta or ?include=_vector to show the vector position or ?include=_classification,_vector for both. When consuming the response use the additional fields such as _vector, as the meta object in the reponse, such as meta.vector will be removed.",
		Msg:          "use of deprecated property ?meta=true/false",
		SinceVersion: "0.22.8",
		SinceTime:    timeMust(time.Parse(time.RFC3339, "2020-06-15T16:18:06.000Z")),
		RemovedIn:    ptString("0.23.0"),
		RemovedTime:  timeMustPtr(time.Parse(time.RFC3339, "2020-06-15T16:18:06.000Z")),
	},
	"config-files": models.Deprecation{
		ID:           "config-files",
		Status:       "deprecated",
		APIType:      "Configuration",
		Mitigation:   "Configure Weaviate using environment variables.",
		Msg:          "use of deprecated command line argument --config-file",
		SinceVersion: "0.22.16",
		SinceTime:    timeMust(time.Parse(time.RFC3339, "2020-09-08T09:46:00.000Z")),
	},
	"cardinality": models.Deprecation{
		ID:           "cardinality",
		Status:       "deprecated",
		APIType:      "REST",
		Mitigation:   "Omit this field. Starting in 0.22.7 it no longer has any effect.",
		Msg:          "use of deprecated property option 'cardinality'",
		SinceVersion: "0.22.17",
		SinceTime:    timeMust(time.Parse(time.RFC3339, "2020-09-16T09:06:00.000Z")),
		RemovedIn:    ptString("0.23.0"),
		RemovedTime:  timeMustPtr(time.Parse(time.RFC3339, "2020-09-16T09:06:00.000Z")),
	},
	"ref-meta-deprecated-fields": models.Deprecation{
		ID:           "ref-meta-deprecated-fields",
		Status:       "deprecated",
		APIType:      "REST",
		Mitigation:   "when using _classification the reference meta after a successful\nclassification contains various counts and distances. Starting in 0.22.20\nthe fields winningDistance and losingDistance are considered deprecated.\nNew fields were added and they have more descriptive names. User\nmeanWinningDistance instead of winningDistance and use meanLosingDistance\ninstead of losingDistance",
		Msg:          "response contains deprecated fields winningDistance and losingDistance",
		SinceVersion: "0.22.20",
		SinceTime:    timeMust(time.Parse(time.RFC3339, "2020-11-26T14:58:00.000Z")),
		RemovedIn:    ptString("0.23.0"),
		RemovedTime:  timeMustPtr(time.Parse(time.RFC3339, "2020-11-26T14:58:00.000Z")),
	},
}
