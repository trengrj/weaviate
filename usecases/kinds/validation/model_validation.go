//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package validation

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/config"
)

type exists func(context.Context, kind.Kind, strfmt.UUID) (bool, error)

const (
	// ErrorMissingActionThings message
	ErrorMissingActionThings string = "no things, object and subject, are added. Add 'things' by using the 'things' key in the root of the JSON"
	// ErrorMissingActionThingsObject message
	ErrorMissingActionThingsObject string = "no object-thing is added. Add the 'object' inside the 'things' part of the JSON"
	// ErrorMissingActionThingsSubject message
	ErrorMissingActionThingsSubject string = "no subject-thing is added. Add the 'subject' inside the 'things' part of the JSON"
	// ErrorMissingActionThingsObjectLocation message
	ErrorMissingActionThingsObjectLocation string = "no 'locationURL' is found in the object-thing. Add the 'locationURL' inside the 'object-thing' part of the JSON"
	// ErrorMissingActionThingsObjectType message
	ErrorMissingActionThingsObjectType string = "no 'type' is found in the object-thing. Add the 'type' inside the 'object-thing' part of the JSON"
	// ErrorInvalidActionThingsObjectType message
	ErrorInvalidActionThingsObjectType string = "object-thing requires one of the following values in 'type': '%s', '%s' or '%s'"
	// ErrorMissingActionThingsSubjectLocation message
	ErrorMissingActionThingsSubjectLocation string = "no 'locationURL' is found in the subject-thing. Add the 'locationURL' inside the 'subject-thing' part of the JSON"
	// ErrorMissingActionThingsSubjectType message
	ErrorMissingActionThingsSubjectType string = "no 'type' is found in the subject-thing. Add the 'type' inside the 'subject-thing' part of the JSON"
	// ErrorInvalidActionThingsSubjectType message
	ErrorInvalidActionThingsSubjectType string = "subject-thing requires one of the following values in 'type': '%s', '%s' or '%s'"
	// ErrorMissingClass message
	ErrorMissingClass string = "the given class is empty"
	// ErrorMissingContext message
	ErrorMissingContext string = "the given context is empty"
	// ErrorNoExternalCredentials message
	ErrorNoExternalCredentials string = "no credentials available for the Weaviate instance for %s given in the %s"
	// ErrorExternalNotFound message
	ErrorExternalNotFound string = "given statuscode of '%s' is '%d', but 200 was expected for LocationURL given in the %s"
	// ErrorInvalidCRefType message
	ErrorInvalidCRefType string = "'cref' type '%s' does not exists"
	// ErrorNotFoundInDatabase message
	ErrorNotFoundInDatabase string = "%s: no %s with id %s found"
)

type Validator struct {
	schema schema.Schema
	exists exists
	config *config.WeaviateConfig
}

func New(schema schema.Schema, exists exists,
	config *config.WeaviateConfig) *Validator {
	return &Validator{
		schema: schema,
		exists: exists,
		config: config,
	}
}

func (v *Validator) Thing(ctx context.Context, thing *models.Thing) error {
	if err := validateClass(thing.Class); err != nil {
		return err
	}

	return v.properties(ctx, kind.Thing, thing)
}

func (v *Validator) Action(ctx context.Context, action *models.Action) error {
	if err := validateClass(action.Class); err != nil {
		return err
	}

	return v.properties(ctx, kind.Action, action)
}

func validateClass(class string) error {
	// If the given class is empty, return an error
	if class == "" {
		return fmt.Errorf(ErrorMissingClass)
	}

	// No error
	return nil
}

// ValidateSingleRef validates a single ref based on location URL and existence of the object in the database
func (v *Validator) ValidateSingleRef(ctx context.Context, cref *models.SingleRef,
	errorVal string) error {
	ref, err := crossref.ParseSingleRef(cref)
	if err != nil {
		return fmt.Errorf("invalid reference: %s", err)
	}

	if !ref.Local {
		return fmt.Errorf("unrecognized cross-ref ref format")
	}

	return v.validateLocalRef(ctx, ref, errorVal)
}

func (v *Validator) validateLocalRef(ctx context.Context, ref *crossref.Ref, errorVal string) error {
	// Check whether the given Object exists in the DB
	var err error

	ok, err := v.exists(ctx, ref.Kind, ref.TargetID)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf(ErrorNotFoundInDatabase, errorVal, ref.Kind.Name(), ref.TargetID)
	}

	return nil
}

func (v *Validator) ValidateMultipleRef(ctx context.Context, refs models.MultipleRef,
	errorVal string) error {
	if refs == nil {
		return nil
	}

	for _, ref := range refs {
		err := v.ValidateSingleRef(ctx, ref, errorVal)
		if err != nil {
			return err
		}
	}
	return nil
}
