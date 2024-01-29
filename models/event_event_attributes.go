// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventEventAttributes event event attributes
//
// swagger:model event.EventAttributes
type EventEventAttributes struct {

	// message
	// Example: deployment service lambda version v0.0.0
	// Required: true
	Message *string `json:"message"`

	// priority
	// Example: P1
	// Required: true
	Priority *string `json:"priority"`

	// related id
	// Example: 53aa35c8-e659-44b2-882f-f6056e443c99
	RelatedID string `json:"related_id,omitempty"`

	// service
	// Example: service-event
	// Required: true
	Service *string `json:"service"`

	// source
	// Example: github_action
	// Required: true
	Source *string `json:"source"`

	// status
	// Example: success
	// Required: true
	Status *string `json:"status"`

	// type
	// Example: deployment
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this event event attributes
func (m *EventEventAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventEventAttributes) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *EventEventAttributes) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *EventEventAttributes) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

func (m *EventEventAttributes) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *EventEventAttributes) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *EventEventAttributes) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event event attributes based on context it is used
func (m *EventEventAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventEventAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventEventAttributes) UnmarshalBinary(b []byte) error {
	var res EventEventAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}