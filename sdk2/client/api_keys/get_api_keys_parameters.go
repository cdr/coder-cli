// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewGetAPIKeysParams creates a new GetAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIKeysParams() *GetAPIKeysParams {
	return &GetAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKeysParamsWithTimeout creates a new GetAPIKeysParams object
// with the ability to set a timeout on a request.
func NewGetAPIKeysParamsWithTimeout(timeout time.Duration) *GetAPIKeysParams {
	return &GetAPIKeysParams{
		timeout: timeout,
	}
}

// NewGetAPIKeysParamsWithContext creates a new GetAPIKeysParams object
// with the ability to set a context for a request.
func NewGetAPIKeysParamsWithContext(ctx context.Context) *GetAPIKeysParams {
	return &GetAPIKeysParams{
		Context: ctx,
	}
}

// NewGetAPIKeysParamsWithHTTPClient creates a new GetAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIKeysParamsWithHTTPClient(client *http.Client) *GetAPIKeysParams {
	return &GetAPIKeysParams{
		HTTPClient: client,
	}
}

/* GetAPIKeysParams contains all the parameters to send to the API endpoint
   for the get api keys operation.

   Typically these are written to a http.Request.
*/
type GetAPIKeysParams struct {

	/* ID.

	   API Key ID
	*/
	ID string

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeysParams) WithDefaults() *GetAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get api keys params
func (o *GetAPIKeysParams) WithTimeout(timeout time.Duration) *GetAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get api keys params
func (o *GetAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get api keys params
func (o *GetAPIKeysParams) WithContext(ctx context.Context) *GetAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get api keys params
func (o *GetAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get api keys params
func (o *GetAPIKeysParams) WithHTTPClient(client *http.Client) *GetAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get api keys params
func (o *GetAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get api keys params
func (o *GetAPIKeysParams) WithID(id string) *GetAPIKeysParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get api keys params
func (o *GetAPIKeysParams) SetID(id string) {
	o.ID = id
}

// WithUserID adds the userID to the get api keys params
func (o *GetAPIKeysParams) WithUserID(userID string) *GetAPIKeysParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get api keys params
func (o *GetAPIKeysParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
