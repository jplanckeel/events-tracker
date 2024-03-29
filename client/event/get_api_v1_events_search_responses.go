// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jplanckeel/events-tracker/models"
)

// GetAPIV1EventsSearchReader is a Reader for the GetAPIV1EventsSearch structure.
type GetAPIV1EventsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV1EventsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV1EventsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIV1EventsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/events/search] GetAPIV1EventsSearch", response, response.Code())
	}
}

// NewGetAPIV1EventsSearchOK creates a GetAPIV1EventsSearchOK with default headers values
func NewGetAPIV1EventsSearchOK() *GetAPIV1EventsSearchOK {
	return &GetAPIV1EventsSearchOK{}
}

/*
GetAPIV1EventsSearchOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV1EventsSearchOK struct {
	Payload []*models.ServicesEventOutput
}

// IsSuccess returns true when this get Api v1 events search o k response has a 2xx status code
func (o *GetAPIV1EventsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v1 events search o k response has a 3xx status code
func (o *GetAPIV1EventsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 events search o k response has a 4xx status code
func (o *GetAPIV1EventsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v1 events search o k response has a 5xx status code
func (o *GetAPIV1EventsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 events search o k response a status code equal to that given
func (o *GetAPIV1EventsSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v1 events search o k response
func (o *GetAPIV1EventsSearchOK) Code() int {
	return 200
}

func (o *GetAPIV1EventsSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/events/search][%d] getApiV1EventsSearchOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1EventsSearchOK) String() string {
	return fmt.Sprintf("[GET /api/v1/events/search][%d] getApiV1EventsSearchOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1EventsSearchOK) GetPayload() []*models.ServicesEventOutput {
	return o.Payload
}

func (o *GetAPIV1EventsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1EventsSearchNotFound creates a GetAPIV1EventsSearchNotFound with default headers values
func NewGetAPIV1EventsSearchNotFound() *GetAPIV1EventsSearchNotFound {
	return &GetAPIV1EventsSearchNotFound{}
}

/*
GetAPIV1EventsSearchNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV1EventsSearchNotFound struct {
	Payload string
}

// IsSuccess returns true when this get Api v1 events search not found response has a 2xx status code
func (o *GetAPIV1EventsSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v1 events search not found response has a 3xx status code
func (o *GetAPIV1EventsSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 events search not found response has a 4xx status code
func (o *GetAPIV1EventsSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v1 events search not found response has a 5xx status code
func (o *GetAPIV1EventsSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 events search not found response a status code equal to that given
func (o *GetAPIV1EventsSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api v1 events search not found response
func (o *GetAPIV1EventsSearchNotFound) Code() int {
	return 404
}

func (o *GetAPIV1EventsSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/events/search][%d] getApiV1EventsSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1EventsSearchNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/events/search][%d] getApiV1EventsSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1EventsSearchNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetAPIV1EventsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
