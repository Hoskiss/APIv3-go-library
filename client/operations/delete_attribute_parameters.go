// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDeleteAttributeParams creates a new DeleteAttributeParams object
// with the default values initialized.
func NewDeleteAttributeParams() *DeleteAttributeParams {
	var ()
	return &DeleteAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAttributeParamsWithTimeout creates a new DeleteAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAttributeParamsWithTimeout(timeout time.Duration) *DeleteAttributeParams {
	var ()
	return &DeleteAttributeParams{

		timeout: timeout,
	}
}

// NewDeleteAttributeParamsWithContext creates a new DeleteAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAttributeParamsWithContext(ctx context.Context) *DeleteAttributeParams {
	var ()
	return &DeleteAttributeParams{

		Context: ctx,
	}
}

// NewDeleteAttributeParamsWithHTTPClient creates a new DeleteAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAttributeParamsWithHTTPClient(client *http.Client) *DeleteAttributeParams {
	var ()
	return &DeleteAttributeParams{
		HTTPClient: client,
	}
}

/*DeleteAttributeParams contains all the parameters to send to the API endpoint
for the delete attribute operation typically these are written to a http.Request
*/
type DeleteAttributeParams struct {

	/*AttributeID
	  id of the attribute

	*/
	AttributeID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete attribute params
func (o *DeleteAttributeParams) WithTimeout(timeout time.Duration) *DeleteAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete attribute params
func (o *DeleteAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete attribute params
func (o *DeleteAttributeParams) WithContext(ctx context.Context) *DeleteAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete attribute params
func (o *DeleteAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete attribute params
func (o *DeleteAttributeParams) WithHTTPClient(client *http.Client) *DeleteAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete attribute params
func (o *DeleteAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeID adds the attributeID to the delete attribute params
func (o *DeleteAttributeParams) WithAttributeID(attributeID int64) *DeleteAttributeParams {
	o.SetAttributeID(attributeID)
	return o
}

// SetAttributeID adds the attributeId to the delete attribute params
func (o *DeleteAttributeParams) SetAttributeID(attributeID int64) {
	o.AttributeID = attributeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeId
	if err := r.SetPathParam("attributeId", swag.FormatInt64(o.AttributeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}