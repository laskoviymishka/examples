// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListUser list user
//
// swagger:model ListUser
type ListUser struct {

	// accepted
	Accepted bool `json:"accepted,omitempty"`

	// access mode
	AccessMode AccessMode `json:"access_mode,omitempty"`

	// me
	Me bool `json:"me,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this list user
func (m *ListUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUser) validateAccessMode(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessMode) { // not required
		return nil
	}

	if err := m.AccessMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access_mode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this list user based on the context it is used
func (m *ListUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUser) contextValidateAccessMode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccessMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access_mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListUser) UnmarshalBinary(b []byte) error {
	var res ListUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
