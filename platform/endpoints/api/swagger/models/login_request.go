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

// LoginRequest login request
//
// swagger:model login_request
type LoginRequest struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this login request
func (m *LoginRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginRequest) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login request based on context it is used
func (m *LoginRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginRequest) UnmarshalBinary(b []byte) error {
	var res LoginRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
