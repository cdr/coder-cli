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
	"github.com/go-openapi/swag"

	"cdr.dev/coder-cli/sdk2/models"
)

// NewPatchImageParams creates a new PatchImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchImageParams() *PatchImageParams {
	return &PatchImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchImageParamsWithTimeout creates a new PatchImageParams object
// with the ability to set a timeout on a request.
func NewPatchImageParamsWithTimeout(timeout time.Duration) *PatchImageParams {
	return &PatchImageParams{
		timeout: timeout,
	}
}

// NewPatchImageParamsWithContext creates a new PatchImageParams object
// with the ability to set a context for a request.
func NewPatchImageParamsWithContext(ctx context.Context) *PatchImageParams {
	return &PatchImageParams{
		Context: ctx,
	}
}

// NewPatchImageParamsWithHTTPClient creates a new PatchImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchImageParamsWithHTTPClient(client *http.Client) *PatchImageParams {
	return &PatchImageParams{
		HTTPClient: client,
	}
}

/* PatchImageParams contains all the parameters to send to the API endpoint
   for the patch image operation.

   Typically these are written to a http.Request.
*/
type PatchImageParams struct {

	/* Envs.

	   Populate 'environments' with environments using this image
	*/
	Envs *bool

	/* ID.

	   Image ID
	*/
	ID string

	/* Request.

	   Request body
	*/
	Request *models.UpdateImageRequest

	/* UserIds.

	   Populate 'user_ids' with User IDs using this image
	*/
	UserIds *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchImageParams) WithDefaults() *PatchImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch image params
func (o *PatchImageParams) WithTimeout(timeout time.Duration) *PatchImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch image params
func (o *PatchImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch image params
func (o *PatchImageParams) WithContext(ctx context.Context) *PatchImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch image params
func (o *PatchImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch image params
func (o *PatchImageParams) WithHTTPClient(client *http.Client) *PatchImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch image params
func (o *PatchImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvs adds the envs to the patch image params
func (o *PatchImageParams) WithEnvs(envs *bool) *PatchImageParams {
	o.SetEnvs(envs)
	return o
}

// SetEnvs adds the envs to the patch image params
func (o *PatchImageParams) SetEnvs(envs *bool) {
	o.Envs = envs
}

// WithID adds the id to the patch image params
func (o *PatchImageParams) WithID(id string) *PatchImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch image params
func (o *PatchImageParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the patch image params
func (o *PatchImageParams) WithRequest(request *models.UpdateImageRequest) *PatchImageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the patch image params
func (o *PatchImageParams) SetRequest(request *models.UpdateImageRequest) {
	o.Request = request
}

// WithUserIds adds the userIds to the patch image params
func (o *PatchImageParams) WithUserIds(userIds *bool) *PatchImageParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the patch image params
func (o *PatchImageParams) SetUserIds(userIds *bool) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *PatchImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Envs != nil {

		// query param envs
		var qrEnvs bool

		if o.Envs != nil {
			qrEnvs = *o.Envs
		}
		qEnvs := swag.FormatBool(qrEnvs)
		if qEnvs != "" {

			if err := r.SetQueryParam("envs", qEnvs); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if o.UserIds != nil {

		// query param user_ids
		var qrUserIds bool

		if o.UserIds != nil {
			qrUserIds = *o.UserIds
		}
		qUserIds := swag.FormatBool(qrUserIds)
		if qUserIds != "" {

			if err := r.SetQueryParam("user_ids", qUserIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
