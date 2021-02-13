// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHKeyPair SSH key pair
//
// swagger:model SSHKeyPair
type SSHKeyPair struct {

	// public key
	PublicKey string `json:"public_key,omitempty"`
}

// Validate validates this SSH key pair
func (m *SSHKeyPair) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH key pair based on context it is used
func (m *SSHKeyPair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeyPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyPair) UnmarshalBinary(b []byte) error {
	var res SSHKeyPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
