// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostAPIV1TagOKCode is the HTTP code returned for type PostAPIV1TagOK
const PostAPIV1TagOKCode int = 200

/*PostAPIV1TagOK Succesfull

swagger:response postApiV1TagOK
*/
type PostAPIV1TagOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewPostAPIV1TagOK creates PostAPIV1TagOK with default headers values
func NewPostAPIV1TagOK() *PostAPIV1TagOK {

	return &PostAPIV1TagOK{}
}

// WithPayload adds the payload to the post Api v1 tag o k response
func (o *PostAPIV1TagOK) WithPayload(payload io.ReadCloser) *PostAPIV1TagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api v1 tag o k response
func (o *PostAPIV1TagOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIV1TagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostAPIV1TagBadRequestCode is the HTTP code returned for type PostAPIV1TagBadRequest
const PostAPIV1TagBadRequestCode int = 400

/*PostAPIV1TagBadRequest Bad Request

swagger:response postApiV1TagBadRequest
*/
type PostAPIV1TagBadRequest struct {
}

// NewPostAPIV1TagBadRequest creates PostAPIV1TagBadRequest with default headers values
func NewPostAPIV1TagBadRequest() *PostAPIV1TagBadRequest {

	return &PostAPIV1TagBadRequest{}
}

// WriteResponse to the client
func (o *PostAPIV1TagBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostAPIV1TagInternalServerErrorCode is the HTTP code returned for type PostAPIV1TagInternalServerError
const PostAPIV1TagInternalServerErrorCode int = 500

/*PostAPIV1TagInternalServerError Internal error

swagger:response postApiV1TagInternalServerError
*/
type PostAPIV1TagInternalServerError struct {
}

// NewPostAPIV1TagInternalServerError creates PostAPIV1TagInternalServerError with default headers values
func NewPostAPIV1TagInternalServerError() *PostAPIV1TagInternalServerError {

	return &PostAPIV1TagInternalServerError{}
}

// WriteResponse to the client
func (o *PostAPIV1TagInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
