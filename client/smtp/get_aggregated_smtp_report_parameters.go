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

// NewGetAggregatedSMTPReportParams creates a new GetAggregatedSMTPReportParams object
// with the default values initialized.
func NewGetAggregatedSMTPReportParams() *GetAggregatedSMTPReportParams {
	var ()
	return &GetAggregatedSMTPReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAggregatedSMTPReportParamsWithTimeout creates a new GetAggregatedSMTPReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAggregatedSMTPReportParamsWithTimeout(timeout time.Duration) *GetAggregatedSMTPReportParams {
	var ()
	return &GetAggregatedSMTPReportParams{

		timeout: timeout,
	}
}

// NewGetAggregatedSMTPReportParamsWithContext creates a new GetAggregatedSMTPReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAggregatedSMTPReportParamsWithContext(ctx context.Context) *GetAggregatedSMTPReportParams {
	var ()
	return &GetAggregatedSMTPReportParams{

		Context: ctx,
	}
}

// NewGetAggregatedSMTPReportParamsWithHTTPClient creates a new GetAggregatedSMTPReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAggregatedSMTPReportParamsWithHTTPClient(client *http.Client) *GetAggregatedSMTPReportParams {
	var ()
	return &GetAggregatedSMTPReportParams{
		HTTPClient: client,
	}
}

/*GetAggregatedSMTPReportParams contains all the parameters to send to the API endpoint
for the get aggregated Smtp report operation typically these are written to a http.Request
*/
type GetAggregatedSMTPReportParams struct {

	/*Days
	  Number of days in the past including today (positive integer). Not compatible with 'startDate' and 'endDate'

	*/
	Days *int64
	/*EndDate
	  Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate

	*/
	EndDate *strfmt.Date
	/*StartDate
	  Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate

	*/
	StartDate *strfmt.Date
	/*Tag
	  Tag of the emails

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithTimeout(timeout time.Duration) *GetAggregatedSMTPReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithContext(ctx context.Context) *GetAggregatedSMTPReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithHTTPClient(client *http.Client) *GetAggregatedSMTPReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDays adds the days to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithDays(days *int64) *GetAggregatedSMTPReportParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetDays(days *int64) {
	o.Days = days
}

// WithEndDate adds the endDate to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithEndDate(endDate *strfmt.Date) *GetAggregatedSMTPReportParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithStartDate(startDate *strfmt.Date) *GetAggregatedSMTPReportParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithTag adds the tag to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) WithTag(tag *string) *GetAggregatedSMTPReportParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get aggregated Smtp report params
func (o *GetAggregatedSMTPReportParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetAggregatedSMTPReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Days != nil {

		// query param days
		var qrDays int64
		if o.Days != nil {
			qrDays = *o.Days
		}
		qDays := swag.FormatInt64(qrDays)
		if qDays != "" {
			if err := r.SetQueryParam("days", qDays); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}