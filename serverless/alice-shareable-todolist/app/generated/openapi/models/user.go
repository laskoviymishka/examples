// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// User user
//
// swagger:model User
type User struct {

	// avatar id
	AvatarID string `json:"avatar_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
