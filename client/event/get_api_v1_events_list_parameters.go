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

// NewGetAPIV1EventsListParams creates a new GetAPIV1EventsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV1EventsListParams() *GetAPIV1EventsListParams {
	return &GetAPIV1EventsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV1EventsListParamsWithTimeout creates a new GetAPIV1EventsListParams object
// with the ability to set a timeout on a request.
func NewGetAPIV1EventsListParamsWithTimeout(timeout time.Duration) *GetAPIV1EventsListParams {
	return &GetAPIV1EventsListParams{
		timeout: timeout,
	}
}

// NewGetAPIV1EventsListParamsWithContext creates a new GetAPIV1EventsListParams object
// with the ability to set a context for a request.
func NewGetAPIV1EventsListParamsWithContext(ctx context.Context) *GetAPIV1EventsListParams {
	return &GetAPIV1EventsListParams{
		Context: ctx,
	}
}

// NewGetAPIV1EventsListParamsWithHTTPClient creates a new GetAPIV1EventsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV1EventsListParamsWithHTTPClient(client *http.Client) *GetAPIV1EventsListParams {
	return &GetAPIV1EventsListParams{
		HTTPClient: client,
	}
}

/*
GetAPIV1EventsListParams contains all the parameters to send to the API endpoint

	for the get API v1 events list operation.

	Typically these are written to a http.Request.
*/
type GetAPIV1EventsListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v1 events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1EventsListParams) WithDefaults() *GetAPIV1EventsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v1 events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV1EventsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v1 events list params
func (o *GetAPIV1EventsListParams) WithTimeout(timeout time.Duration) *GetAPIV1EventsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v1 events list params
func (o *GetAPIV1EventsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v1 events list params
func (o *GetAPIV1EventsListParams) WithContext(ctx context.Context) *GetAPIV1EventsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v1 events list params
func (o *GetAPIV1EventsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v1 events list params
func (o *GetAPIV1EventsListParams) WithHTTPClient(client *http.Client) *GetAPIV1EventsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v1 events list params
func (o *GetAPIV1EventsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV1EventsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
