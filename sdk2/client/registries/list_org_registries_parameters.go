// Code generated by go-swagger; DO NOT EDIT.

package registries

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
)

// NewListOrgRegistriesParams creates a new ListOrgRegistriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOrgRegistriesParams() *ListOrgRegistriesParams {
	return &ListOrgRegistriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOrgRegistriesParamsWithTimeout creates a new ListOrgRegistriesParams object
// with the ability to set a timeout on a request.
func NewListOrgRegistriesParamsWithTimeout(timeout time.Duration) *ListOrgRegistriesParams {
	return &ListOrgRegistriesParams{
		timeout: timeout,
	}
}

// NewListOrgRegistriesParamsWithContext creates a new ListOrgRegistriesParams object
// with the ability to set a context for a request.
func NewListOrgRegistriesParamsWithContext(ctx context.Context) *ListOrgRegistriesParams {
	return &ListOrgRegistriesParams{
		Context: ctx,
	}
}

// NewListOrgRegistriesParamsWithHTTPClient creates a new ListOrgRegistriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOrgRegistriesParamsWithHTTPClient(client *http.Client) *ListOrgRegistriesParams {
	return &ListOrgRegistriesParams{
		HTTPClient: client,
	}
}

/* ListOrgRegistriesParams contains all the parameters to send to the API endpoint
   for the list org registries operation.

   Typically these are written to a http.Request.
*/
type ListOrgRegistriesParams struct {

	/* Org.

	   Organization ID
	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list org registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrgRegistriesParams) WithDefaults() *ListOrgRegistriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list org registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrgRegistriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list org registries params
func (o *ListOrgRegistriesParams) WithTimeout(timeout time.Duration) *ListOrgRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list org registries params
func (o *ListOrgRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list org registries params
func (o *ListOrgRegistriesParams) WithContext(ctx context.Context) *ListOrgRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list org registries params
func (o *ListOrgRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list org registries params
func (o *ListOrgRegistriesParams) WithHTTPClient(client *http.Client) *ListOrgRegistriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list org registries params
func (o *ListOrgRegistriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the list org registries params
func (o *ListOrgRegistriesParams) WithOrg(org string) *ListOrgRegistriesParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the list org registries params
func (o *ListOrgRegistriesParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrgRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param org
	qrOrg := o.Org
	qOrg := qrOrg
	if qOrg != "" {

		if err := r.SetQueryParam("org", qOrg); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
