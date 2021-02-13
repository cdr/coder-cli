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

// NewDeleteImageTagParams creates a new DeleteImageTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteImageTagParams() *DeleteImageTagParams {
	return &DeleteImageTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageTagParamsWithTimeout creates a new DeleteImageTagParams object
// with the ability to set a timeout on a request.
func NewDeleteImageTagParamsWithTimeout(timeout time.Duration) *DeleteImageTagParams {
	return &DeleteImageTagParams{
		timeout: timeout,
	}
}

// NewDeleteImageTagParamsWithContext creates a new DeleteImageTagParams object
// with the ability to set a context for a request.
func NewDeleteImageTagParamsWithContext(ctx context.Context) *DeleteImageTagParams {
	return &DeleteImageTagParams{
		Context: ctx,
	}
}

// NewDeleteImageTagParamsWithHTTPClient creates a new DeleteImageTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteImageTagParamsWithHTTPClient(client *http.Client) *DeleteImageTagParams {
	return &DeleteImageTagParams{
		HTTPClient: client,
	}
}

/* DeleteImageTagParams contains all the parameters to send to the API endpoint
   for the delete image tag operation.

   Typically these are written to a http.Request.
*/
type DeleteImageTagParams struct {

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

// WithDefaults hydrates default values in the delete image tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImageTagParams) WithDefaults() *DeleteImageTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete image tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteImageTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete image tag params
func (o *DeleteImageTagParams) WithTimeout(timeout time.Duration) *DeleteImageTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image tag params
func (o *DeleteImageTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image tag params
func (o *DeleteImageTagParams) WithContext(ctx context.Context) *DeleteImageTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image tag params
func (o *DeleteImageTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image tag params
func (o *DeleteImageTagParams) WithHTTPClient(client *http.Client) *DeleteImageTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image tag params
func (o *DeleteImageTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete image tag params
func (o *DeleteImageTagParams) WithID(id string) *DeleteImageTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete image tag params
func (o *DeleteImageTagParams) SetID(id string) {
	o.ID = id
}

// WithTag adds the tag to the delete image tag params
func (o *DeleteImageTagParams) WithTag(tag string) *DeleteImageTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the delete image tag params
func (o *DeleteImageTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
