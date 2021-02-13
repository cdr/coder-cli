// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditCountResponse audit count response
//
// swagger:model AuditCountResponse
type AuditCountResponse struct {

	// count
	Count int64 `json:"count,omitempty"`
}

// Validate validates this audit count response
func (m *AuditCountResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit count response based on context it is used
func (m *AuditCountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditCountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditCountResponse) UnmarshalBinary(b []byte) error {
	var res AuditCountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
