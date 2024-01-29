// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new event API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for event API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV1EventID(params *GetAPIV1EventIDParams, opts ...ClientOption) (*GetAPIV1EventIDOK, error)

	GetAPIV1EventsList(params *GetAPIV1EventsListParams, opts ...ClientOption) (*GetAPIV1EventsListOK, error)

	GetAPIV1EventsSearch(params *GetAPIV1EventsSearchParams, opts ...ClientOption) (*GetAPIV1EventsSearchOK, error)

	PostAPIV1Event(params *PostAPIV1EventParams, opts ...ClientOption) (*PostAPIV1EventOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV1EventID gets an events by id

Get an Events by id
*/
func (a *Client) GetAPIV1EventID(params *GetAPIV1EventIDParams, opts ...ClientOption) (*GetAPIV1EventIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1EventIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1EventID",
		Method:             "GET",
		PathPattern:        "/api/v1/event/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1EventIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV1EventIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1EventID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV1EventsList lists all events

List all Events
*/
func (a *Client) GetAPIV1EventsList(params *GetAPIV1EventsListParams, opts ...ClientOption) (*GetAPIV1EventsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1EventsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1EventsList",
		Method:             "GET",
		PathPattern:        "/api/v1/events/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1EventsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV1EventsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1EventsList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV1EventsSearch searches events

Search Events
*/
func (a *Client) GetAPIV1EventsSearch(params *GetAPIV1EventsSearchParams, opts ...ClientOption) (*GetAPIV1EventsSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV1EventsSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV1EventsSearch",
		Method:             "GET",
		PathPattern:        "/api/v1/events/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV1EventsSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV1EventsSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV1EventsSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV1Event creates an event

Create an Event
*/
func (a *Client) PostAPIV1Event(params *PostAPIV1EventParams, opts ...ClientOption) (*PostAPIV1EventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV1EventParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV1Event",
		Method:             "POST",
		PathPattern:        "/api/v1/event/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIV1EventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPIV1EventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV1Event: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}