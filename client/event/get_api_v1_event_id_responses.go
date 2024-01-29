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

// GetAPIV1EventIDReader is a Reader for the GetAPIV1EventID structure.
type GetAPIV1EventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV1EventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV1EventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIV1EventIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/event/{id}] GetAPIV1EventID", response, response.Code())
	}
}

// NewGetAPIV1EventIDOK creates a GetAPIV1EventIDOK with default headers values
func NewGetAPIV1EventIDOK() *GetAPIV1EventIDOK {
	return &GetAPIV1EventIDOK{}
}

/*
GetAPIV1EventIDOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV1EventIDOK struct {
	Payload *models.ServicesEventOutput
}

// IsSuccess returns true when this get Api v1 event Id o k response has a 2xx status code
func (o *GetAPIV1EventIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v1 event Id o k response has a 3xx status code
func (o *GetAPIV1EventIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 event Id o k response has a 4xx status code
func (o *GetAPIV1EventIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v1 event Id o k response has a 5xx status code
func (o *GetAPIV1EventIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 event Id o k response a status code equal to that given
func (o *GetAPIV1EventIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v1 event Id o k response
func (o *GetAPIV1EventIDOK) Code() int {
	return 200
}

func (o *GetAPIV1EventIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/event/{id}][%d] getApiV1EventIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1EventIDOK) String() string {
	return fmt.Sprintf("[GET /api/v1/event/{id}][%d] getApiV1EventIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV1EventIDOK) GetPayload() *models.ServicesEventOutput {
	return o.Payload
}

func (o *GetAPIV1EventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServicesEventOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1EventIDNotFound creates a GetAPIV1EventIDNotFound with default headers values
func NewGetAPIV1EventIDNotFound() *GetAPIV1EventIDNotFound {
	return &GetAPIV1EventIDNotFound{}
}

/*
GetAPIV1EventIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV1EventIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this get Api v1 event Id not found response has a 2xx status code
func (o *GetAPIV1EventIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api v1 event Id not found response has a 3xx status code
func (o *GetAPIV1EventIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v1 event Id not found response has a 4xx status code
func (o *GetAPIV1EventIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api v1 event Id not found response has a 5xx status code
func (o *GetAPIV1EventIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v1 event Id not found response a status code equal to that given
func (o *GetAPIV1EventIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api v1 event Id not found response
func (o *GetAPIV1EventIDNotFound) Code() int {
	return 404
}

func (o *GetAPIV1EventIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/event/{id}][%d] getApiV1EventIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1EventIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/event/{id}][%d] getApiV1EventIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIV1EventIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetAPIV1EventIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
