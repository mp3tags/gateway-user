// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "mp3tags/gateway-user/models"
)

// NewPostAPIV1TagParams creates a new PostAPIV1TagParams object
// no default values defined in spec.
func NewPostAPIV1TagParams() PostAPIV1TagParams {

	return PostAPIV1TagParams{}
}

// PostAPIV1TagParams contains all the bound params for the post API v1 tag operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostAPIV1Tag
type PostAPIV1TagParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Tags *models.ChangeTagsRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostAPIV1TagParams() beforehand.
func (o *PostAPIV1TagParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ChangeTagsRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("tags", "body"))
			} else {
				res = append(res, errors.NewParseError("tags", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Tags = &body
			}
		}
	} else {
		res = append(res, errors.Required("tags", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
