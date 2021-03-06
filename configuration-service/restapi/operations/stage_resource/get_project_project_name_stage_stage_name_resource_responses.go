// Code generated by go-swagger; DO NOT EDIT.

package stage_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/configuration-service/models"
)

// GetProjectProjectNameStageStageNameResourceOKCode is the HTTP code returned for type GetProjectProjectNameStageStageNameResourceOK
const GetProjectProjectNameStageStageNameResourceOKCode int = 200

/*GetProjectProjectNameStageStageNameResourceOK Success

swagger:response getProjectProjectNameStageStageNameResourceOK
*/
type GetProjectProjectNameStageStageNameResourceOK struct {

	/*
	  In: Body
	*/
	Payload *models.Resources `json:"body,omitempty"`
}

// NewGetProjectProjectNameStageStageNameResourceOK creates GetProjectProjectNameStageStageNameResourceOK with default headers values
func NewGetProjectProjectNameStageStageNameResourceOK() *GetProjectProjectNameStageStageNameResourceOK {

	return &GetProjectProjectNameStageStageNameResourceOK{}
}

// WithPayload adds the payload to the get project project name stage stage name resource o k response
func (o *GetProjectProjectNameStageStageNameResourceOK) WithPayload(payload *models.Resources) *GetProjectProjectNameStageStageNameResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project project name stage stage name resource o k response
func (o *GetProjectProjectNameStageStageNameResourceOK) SetPayload(payload *models.Resources) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectProjectNameStageStageNameResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectProjectNameStageStageNameResourceNotFoundCode is the HTTP code returned for type GetProjectProjectNameStageStageNameResourceNotFound
const GetProjectProjectNameStageStageNameResourceNotFoundCode int = 404

/*GetProjectProjectNameStageStageNameResourceNotFound Failed. Stage could not be found.

swagger:response getProjectProjectNameStageStageNameResourceNotFound
*/
type GetProjectProjectNameStageStageNameResourceNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectProjectNameStageStageNameResourceNotFound creates GetProjectProjectNameStageStageNameResourceNotFound with default headers values
func NewGetProjectProjectNameStageStageNameResourceNotFound() *GetProjectProjectNameStageStageNameResourceNotFound {

	return &GetProjectProjectNameStageStageNameResourceNotFound{}
}

// WithPayload adds the payload to the get project project name stage stage name resource not found response
func (o *GetProjectProjectNameStageStageNameResourceNotFound) WithPayload(payload *models.Error) *GetProjectProjectNameStageStageNameResourceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project project name stage stage name resource not found response
func (o *GetProjectProjectNameStageStageNameResourceNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectProjectNameStageStageNameResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProjectProjectNameStageStageNameResourceDefault Error

swagger:response getProjectProjectNameStageStageNameResourceDefault
*/
type GetProjectProjectNameStageStageNameResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectProjectNameStageStageNameResourceDefault creates GetProjectProjectNameStageStageNameResourceDefault with default headers values
func NewGetProjectProjectNameStageStageNameResourceDefault(code int) *GetProjectProjectNameStageStageNameResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProjectProjectNameStageStageNameResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get project project name stage stage name resource default response
func (o *GetProjectProjectNameStageStageNameResourceDefault) WithStatusCode(code int) *GetProjectProjectNameStageStageNameResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get project project name stage stage name resource default response
func (o *GetProjectProjectNameStageStageNameResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get project project name stage stage name resource default response
func (o *GetProjectProjectNameStageStageNameResourceDefault) WithPayload(payload *models.Error) *GetProjectProjectNameStageStageNameResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project project name stage stage name resource default response
func (o *GetProjectProjectNameStageStageNameResourceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectProjectNameStageStageNameResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
