// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewGetImageTagParams creates a new GetImageTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetImageTagParams() *GetImageTagParams {
	return &GetImageTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageTagParamsWithTimeout creates a new GetImageTagParams object
// with the ability to set a timeout on a request.
func NewGetImageTagParamsWithTimeout(timeout time.Duration) *GetImageTagParams {
	return &GetImageTagParams{
		timeout: timeout,
	}
}

// NewGetImageTagParamsWithContext creates a new GetImageTagParams object
// with the ability to set a context for a request.
func NewGetImageTagParamsWithContext(ctx context.Context) *GetImageTagParams {
	return &GetImageTagParams{
		Context: ctx,
	}
}

// NewGetImageTagParamsWithHTTPClient creates a new GetImageTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetImageTagParamsWithHTTPClient(client *http.Client) *GetImageTagParams {
	return &GetImageTagParams{
		HTTPClient: client,
	}
}

/* GetImageTagParams contains all the parameters to send to the API endpoint
   for the get image tag operation.

   Typically these are written to a http.Request.
*/
type GetImageTagParams struct {

	/* ID.

	   Image ID
	*/
	ID string

	/* Tag.

	   Image Tag ID
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get image tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImageTagParams) WithDefaults() *GetImageTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get image tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImageTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get image tag params
func (o *GetImageTagParams) WithTimeout(timeout time.Duration) *GetImageTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image tag params
func (o *GetImageTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image tag params
func (o *GetImageTagParams) WithContext(ctx context.Context) *GetImageTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image tag params
func (o *GetImageTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image tag params
func (o *GetImageTagParams) WithHTTPClient(client *http.Client) *GetImageTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image tag params
func (o *GetImageTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get image tag params
func (o *GetImageTagParams) WithID(id string) *GetImageTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get image tag params
func (o *GetImageTagParams) SetID(id string) {
	o.ID = id
}

// WithTag adds the tag to the get image tag params
func (o *GetImageTagParams) WithTag(tag string) *GetImageTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get image tag params
func (o *GetImageTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
