// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PetCreateHandlerFunc turns a function with the right signature into a pet create handler
type PetCreateHandlerFunc func(PetCreateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PetCreateHandlerFunc) Handle(params PetCreateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PetCreateHandler interface for that can handle valid pet create params
type PetCreateHandler interface {
	Handle(PetCreateParams, interface{}) middleware.Responder
}

// NewPetCreate creates a new http.Handler for the pet create operation
func NewPetCreate(ctx *middleware.Context, handler PetCreateHandler) *PetCreate {
	return &PetCreate{Context: ctx, Handler: handler}
}

/*PetCreate swagger:route POST /pet pet petCreate

Add a new pet to the store

*/
type PetCreate struct {
	Context *middleware.Context
	Handler PetCreateHandler
}

func (o *PetCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPetCreateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
