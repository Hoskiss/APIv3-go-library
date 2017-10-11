// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSMTPTemplatesParams creates a new GetSMTPTemplatesParams object
// with the default values initialized.
func NewGetSMTPTemplatesParams() *GetSMTPTemplatesParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPTemplatesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSMTPTemplatesParamsWithTimeout creates a new GetSMTPTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSMTPTemplatesParamsWithTimeout(timeout time.Duration) *GetSMTPTemplatesParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPTemplatesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetSMTPTemplatesParamsWithContext creates a new GetSMTPTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSMTPTemplatesParamsWithContext(ctx context.Context) *GetSMTPTemplatesParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPTemplatesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetSMTPTemplatesParamsWithHTTPClient creates a new GetSMTPTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSMTPTemplatesParamsWithHTTPClient(client *http.Client) *GetSMTPTemplatesParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPTemplatesParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetSMTPTemplatesParams contains all the parameters to send to the API endpoint
for the get Smtp templates operation typically these are written to a http.Request
*/
type GetSMTPTemplatesParams struct {

	/*Limit
	  Number of documents returned per page

	*/
	Limit *int64
	/*Offset
	  Index of the first document in the page

	*/
	Offset *int64
	/*TemplateStatus
	  Filter on the status of the template. Active = true, inactive = false

	*/
	TemplateStatus *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithTimeout(timeout time.Duration) *GetSMTPTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithContext(ctx context.Context) *GetSMTPTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithHTTPClient(client *http.Client) *GetSMTPTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithLimit(limit *int64) *GetSMTPTemplatesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithOffset(offset *int64) *GetSMTPTemplatesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithTemplateStatus adds the templateStatus to the get Smtp templates params
func (o *GetSMTPTemplatesParams) WithTemplateStatus(templateStatus *bool) *GetSMTPTemplatesParams {
	o.SetTemplateStatus(templateStatus)
	return o
}

// SetTemplateStatus adds the templateStatus to the get Smtp templates params
func (o *GetSMTPTemplatesParams) SetTemplateStatus(templateStatus *bool) {
	o.TemplateStatus = templateStatus
}

// WriteToRequest writes these params to a swagger request
func (o *GetSMTPTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.TemplateStatus != nil {

		// query param templateStatus
		var qrTemplateStatus bool
		if o.TemplateStatus != nil {
			qrTemplateStatus = *o.TemplateStatus
		}
		qTemplateStatus := swag.FormatBool(qrTemplateStatus)
		if qTemplateStatus != "" {
			if err := r.SetQueryParam("templateStatus", qTemplateStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}