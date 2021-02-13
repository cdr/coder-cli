// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewGetConfigEnvironmentsParams creates a new GetConfigEnvironmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigEnvironmentsParams() *GetConfigEnvironmentsParams {
	return &GetConfigEnvironmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigEnvironmentsParamsWithTimeout creates a new GetConfigEnvironmentsParams object
// with the ability to set a timeout on a request.
func NewGetConfigEnvironmentsParamsWithTimeout(timeout time.Duration) *GetConfigEnvironmentsParams {
	return &GetConfigEnvironmentsParams{
		timeout: timeout,
	}
}

// NewGetConfigEnvironmentsParamsWithContext creates a new GetConfigEnvironmentsParams object
// with the ability to set a context for a request.
func NewGetConfigEnvironmentsParamsWithContext(ctx context.Context) *GetConfigEnvironmentsParams {
	return &GetConfigEnvironmentsParams{
		Context: ctx,
	}
}

// NewGetConfigEnvironmentsParamsWithHTTPClient creates a new GetConfigEnvironmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigEnvironmentsParamsWithHTTPClient(client *http.Client) *GetConfigEnvironmentsParams {
	return &GetConfigEnvironmentsParams{
		HTTPClient: client,
	}
}

/* GetConfigEnvironmentsParams contains all the parameters to send to the API endpoint
   for the get config environments operation.

   Typically these are written to a http.Request.
*/
type GetConfigEnvironmentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get config environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigEnvironmentsParams) WithDefaults() *GetConfigEnvironmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get config environments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigEnvironmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get config environments params
func (o *GetConfigEnvironmentsParams) WithTimeout(timeout time.Duration) *GetConfigEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config environments params
func (o *GetConfigEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config environments params
func (o *GetConfigEnvironmentsParams) WithContext(ctx context.Context) *GetConfigEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config environments params
func (o *GetConfigEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get config environments params
func (o *GetConfigEnvironmentsParams) WithHTTPClient(client *http.Client) *GetConfigEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config environments params
func (o *GetConfigEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
