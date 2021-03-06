// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCampaignRecipients get campaign recipients
// swagger:model getCampaignRecipients
type GetCampaignRecipients struct {

	// exclusion lists
	// Required: true
	ExclusionLists []int64 `json:"exclusionLists"`

	// lists
	// Required: true
	Lists []int64 `json:"lists"`
}

// Validate validates this get campaign recipients
func (m *GetCampaignRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionLists(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLists(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCampaignRecipients) validateExclusionLists(formats strfmt.Registry) error {

	if err := validate.Required("exclusionLists", "body", m.ExclusionLists); err != nil {
		return err
	}

	return nil
}

func (m *GetCampaignRecipients) validateLists(formats strfmt.Registry) error {

	if err := validate.Required("lists", "body", m.Lists); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCampaignRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCampaignRecipients) UnmarshalBinary(b []byte) error {
	var res GetCampaignRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
