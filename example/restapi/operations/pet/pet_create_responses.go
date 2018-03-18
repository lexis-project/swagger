// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Stratoscale/swagger/example/models"
)

// PetCreateCreatedCode is the HTTP code returned for type PetCreateCreated
const PetCreateCreatedCode int = 201

/*PetCreateCreated Created

swagger:response petCreateCreated
*/
type PetCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewPetCreateCreated creates PetCreateCreated with default headers values
func NewPetCreateCreated() *PetCreateCreated {
	return &PetCreateCreated{}
}

// WithPayload adds the payload to the pet create created response
func (o *PetCreateCreated) WithPayload(payload *models.Pet) *PetCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pet create created response
func (o *PetCreateCreated) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PetCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PetCreateMethodNotAllowedCode is the HTTP code returned for type PetCreateMethodNotAllowed
const PetCreateMethodNotAllowedCode int = 405

/*PetCreateMethodNotAllowed Invalid input

swagger:response petCreateMethodNotAllowed
*/
type PetCreateMethodNotAllowed struct {
}

// NewPetCreateMethodNotAllowed creates PetCreateMethodNotAllowed with default headers values
func NewPetCreateMethodNotAllowed() *PetCreateMethodNotAllowed {
	return &PetCreateMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PetCreateMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
