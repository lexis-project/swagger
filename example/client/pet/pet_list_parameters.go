// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

/*PetListParams contains all the parameters to send to the API endpoint
for the pet list operation typically these are written to a http.Request
*/
type PetListParams struct {

	/*Status
	  Status values that need to be considered for filter

	*/
	Status []string

	Timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *PetListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.Timeout); err != nil {
		return err
	}
	var res []error

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
