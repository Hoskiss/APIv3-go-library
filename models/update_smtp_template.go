// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateSMTPTemplate update Smtp template
// swagger:model updateSmtpTemplate

type UpdateSMTPTemplate struct {

	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff and rtf
	AttachmentURL string `json:"attachmentUrl,omitempty"`

	// Required if htmlUrl is empty. Body of the message (HTML must have more than 10 characters)
	HTMLContent string `json:"htmlContent,omitempty"`

	// Required if htmlContent is empty. URL to the body of the email (HTML)
	HTMLURL string `json:"htmlUrl,omitempty"`

	// Status of the template. isActive = false means template is inactive, isActive = true means template is active
	IsActive bool `json:"isActive,omitempty"`

	// Email on which campaign recipients will be able to reply to
	ReplyTo strfmt.Email `json:"replyTo,omitempty"`

	// sender
	Sender *UpdateSMTPTemplateSender `json:"sender,omitempty"`

	// Subject of the email
	Subject string `json:"subject,omitempty"`

	// Tag of the template
	Tag string `json:"tag,omitempty"`

	// Name of the template
	TemplateName string `json:"templateName,omitempty"`

	// To personalize the «To» Field, e.g. if you want to include the first name and last name of your recipient, add [FNAME] [LNAME]. These attributes must already exist in contacts database
	ToField string `json:"toField,omitempty"`
}

/* polymorph updateSmtpTemplate attachmentUrl false */

/* polymorph updateSmtpTemplate htmlContent false */

/* polymorph updateSmtpTemplate htmlUrl false */

/* polymorph updateSmtpTemplate isActive false */

/* polymorph updateSmtpTemplate replyTo false */

/* polymorph updateSmtpTemplate sender false */

/* polymorph updateSmtpTemplate subject false */

/* polymorph updateSmtpTemplate tag false */

/* polymorph updateSmtpTemplate templateName false */

/* polymorph updateSmtpTemplate toField false */

// Validate validates this update Smtp template
func (m *UpdateSMTPTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSMTPTemplate) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {

		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSMTPTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSMTPTemplate) UnmarshalBinary(b []byte) error {
	var res UpdateSMTPTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}