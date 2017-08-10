/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package locations

 
 

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateLocationsDeleteParams creates a new WeaviateLocationsDeleteParams object
// with the default values initialized.
func NewWeaviateLocationsDeleteParams() WeaviateLocationsDeleteParams {
	var ()
	return WeaviateLocationsDeleteParams{}
}

// WeaviateLocationsDeleteParams contains all the bound params for the weaviate locations delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.locations.delete
type WeaviateLocationsDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	LocationID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaviateLocationsDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rLocationID, rhkLocationID, _ := route.Params.GetOK("locationId")
	if err := o.bindLocationID(rLocationID, rhkLocationID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateLocationsDeleteParams) bindLocationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("locationId", "path", "strfmt.UUID", raw)
	}
	o.LocationID = *(value.(*strfmt.UUID))

	return nil
}
