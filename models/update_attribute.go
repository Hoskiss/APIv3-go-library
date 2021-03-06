// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateAttribute update attribute
// swagger:model updateAttribute
type UpdateAttribute struct {

	// enumeration
	Enumeration UpdateAttributeEnumeration `json:"enumeration"`

	// Value of the attribute. Use only if the attribute's category is calculated or global
	Value string `json:"value,omitempty"`
}

// Validate validates this update attribute
func (m *UpdateAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAttribute) UnmarshalBinary(b []byte) error {
	var res UpdateAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
