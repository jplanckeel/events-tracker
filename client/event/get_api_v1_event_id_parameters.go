// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewGetAPIV1EventIDParams creates a new GetAPIV1EventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV1EventIDParams() *GetAPIV1EventIDParams {
	return &GetAPIV1EventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV1EventIDParamsWithTimeout creates a new GetAPIV1EventIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIV1EventIDParamsWithTimeout(timeout time.Duration) *GetAPIV1EventIDParams {
	return &GetAPIV1EventIDParams{
		timeout: timeout,
	}
}

// NewGetAPIV1EventIDParamsWithContext creates a new GetAPIV1EventIDParams object
// with the ability to set a context for a request.
func NewGetAPIV1EventIDParamsWithContext(ctx context.Context) *GetAPIV1EventIDParams {
	return &GetAPIV1EventIDParams{
		Context: ctx,
	}
}

// NewGetAPIV1EventIDParamsWithHTTPClient creates a new GetAPIV1EventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV1EventIDParamsWithHTTPClient(client *http.Client) *GetAPIV1EventIDParams {
	return &GetAPIV1EventIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIV1EventIDParams contains all the parameters to send to the API endpoint

	for the get API v1 event ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIV1EventIDParams struct {

	/* ID.

	   xxxxx-xxxxxx-xxxxx-xxxxxx
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v1 event ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1EventIDParams) WithDefaults() *GetAPIV1EventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v1 event ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1EventIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) WithTimeout(timeout time.Duration) *GetAPIV1EventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) WithContext(ctx context.Context) *GetAPIV1EventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) WithHTTPClient(client *http.Client) *GetAPIV1EventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) WithID(id string) *GetAPIV1EventIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v1 event ID params
func (o *GetAPIV1EventIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV1EventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
