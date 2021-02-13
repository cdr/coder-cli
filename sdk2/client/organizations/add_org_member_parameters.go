// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"cdr.dev/coder-cli/sdk2/models"
)

// NewAddOrgMemberParams creates a new AddOrgMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddOrgMemberParams() *AddOrgMemberParams {
	return &AddOrgMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrgMemberParamsWithTimeout creates a new AddOrgMemberParams object
// with the ability to set a timeout on a request.
func NewAddOrgMemberParamsWithTimeout(timeout time.Duration) *AddOrgMemberParams {
	return &AddOrgMemberParams{
		timeout: timeout,
	}
}

// NewAddOrgMemberParamsWithContext creates a new AddOrgMemberParams object
// with the ability to set a context for a request.
func NewAddOrgMemberParamsWithContext(ctx context.Context) *AddOrgMemberParams {
	return &AddOrgMemberParams{
		Context: ctx,
	}
}

// NewAddOrgMemberParamsWithHTTPClient creates a new AddOrgMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddOrgMemberParamsWithHTTPClient(client *http.Client) *AddOrgMemberParams {
	return &AddOrgMemberParams{
		HTTPClient: client,
	}
}

/* AddOrgMemberParams contains all the parameters to send to the API endpoint
   for the add org member operation.

   Typically these are written to a http.Request.
*/
type AddOrgMemberParams struct {

	/* ID.

	   Org ID
	*/
	ID string

	/* Request.

	   Request body
	*/
	Request *models.AddOrganizationMemberRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrgMemberParams) WithDefaults() *AddOrgMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrgMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add org member params
func (o *AddOrgMemberParams) WithTimeout(timeout time.Duration) *AddOrgMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add org member params
func (o *AddOrgMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add org member params
func (o *AddOrgMemberParams) WithContext(ctx context.Context) *AddOrgMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add org member params
func (o *AddOrgMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add org member params
func (o *AddOrgMemberParams) WithHTTPClient(client *http.Client) *AddOrgMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add org member params
func (o *AddOrgMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add org member params
func (o *AddOrgMemberParams) WithID(id string) *AddOrgMemberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add org member params
func (o *AddOrgMemberParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the add org member params
func (o *AddOrgMemberParams) WithRequest(request *models.AddOrganizationMemberRequest) *AddOrgMemberParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the add org member params
func (o *AddOrgMemberParams) SetRequest(request *models.AddOrganizationMemberRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrgMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
