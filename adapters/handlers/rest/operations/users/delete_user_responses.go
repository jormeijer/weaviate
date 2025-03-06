//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// DeleteUserNoContentCode is the HTTP code returned for type DeleteUserNoContent
const DeleteUserNoContentCode int = 204

/*
DeleteUserNoContent Successfully deleted.

swagger:response deleteUserNoContent
*/
type DeleteUserNoContent struct{}

// NewDeleteUserNoContent creates DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*
DeleteUserBadRequest Malformed request.

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

// WithPayload adds the payload to the delete user bad request response
func (o *DeleteUserBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user bad request response
func (o *DeleteUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserUnauthorizedCode is the HTTP code returned for type DeleteUserUnauthorized
const DeleteUserUnauthorizedCode int = 401

/*
DeleteUserUnauthorized Unauthorized or invalid credentials.

swagger:response deleteUserUnauthorized
*/
type DeleteUserUnauthorized struct{}

// NewDeleteUserUnauthorized creates DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteUserForbiddenCode is the HTTP code returned for type DeleteUserForbidden
const DeleteUserForbiddenCode int = 403

/*
DeleteUserForbidden Forbidden

swagger:response deleteUserForbidden
*/
type DeleteUserForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserForbidden creates DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {
	return &DeleteUserForbidden{}
}

// WithPayload adds the payload to the delete user forbidden response
func (o *DeleteUserForbidden) WithPayload(payload *models.ErrorResponse) *DeleteUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user forbidden response
func (o *DeleteUserForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*
DeleteUserNotFound user not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct{}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteUserUnprocessableEntityCode is the HTTP code returned for type DeleteUserUnprocessableEntity
const DeleteUserUnprocessableEntityCode int = 422

/*
DeleteUserUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous.

swagger:response deleteUserUnprocessableEntity
*/
type DeleteUserUnprocessableEntity struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserUnprocessableEntity creates DeleteUserUnprocessableEntity with default headers values
func NewDeleteUserUnprocessableEntity() *DeleteUserUnprocessableEntity {
	return &DeleteUserUnprocessableEntity{}
}

// WithPayload adds the payload to the delete user unprocessable entity response
func (o *DeleteUserUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *DeleteUserUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user unprocessable entity response
func (o *DeleteUserUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserInternalServerErrorCode is the HTTP code returned for type DeleteUserInternalServerError
const DeleteUserInternalServerErrorCode int = 500

/*
DeleteUserInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response deleteUserInternalServerError
*/
type DeleteUserInternalServerError struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserInternalServerError creates DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

// WithPayload adds the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
