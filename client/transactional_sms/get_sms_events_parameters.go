// Code generated by go-swagger; DO NOT EDIT.

package transactional_sms

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

// NewGetSMSEventsParams creates a new GetSMSEventsParams object
// with the default values initialized.
func NewGetSMSEventsParams() *GetSMSEventsParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMSEventsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSMSEventsParamsWithTimeout creates a new GetSMSEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSMSEventsParamsWithTimeout(timeout time.Duration) *GetSMSEventsParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMSEventsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetSMSEventsParamsWithContext creates a new GetSMSEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSMSEventsParamsWithContext(ctx context.Context) *GetSMSEventsParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMSEventsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetSMSEventsParamsWithHTTPClient creates a new GetSMSEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSMSEventsParamsWithHTTPClient(client *http.Client) *GetSMSEventsParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMSEventsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetSMSEventsParams contains all the parameters to send to the API endpoint
for the get Sms events operation typically these are written to a http.Request
*/
type GetSMSEventsParams struct {

	/*Days
	  Number of days in the past including today (positive integer). Not compatible with 'startDate' and 'endDate'

	*/
	Days *int64
	/*EndDate
	  Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report

	*/
	EndDate *strfmt.Date
	/*Event
	  Filter the report for specific events

	*/
	Event *string
	/*Limit
	  Number of documents per page

	*/
	Limit *int64
	/*Offset
	  Index of the first document of the page

	*/
	Offset *int64
	/*PhoneNumber
	  Filter the report for a specific phone number

	*/
	PhoneNumber *string
	/*StartDate
	  Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report

	*/
	StartDate *strfmt.Date
	/*Tags
	  Filter the report for specific tags passed as a serialized urlencoded array

	*/
	Tags *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Sms events params
func (o *GetSMSEventsParams) WithTimeout(timeout time.Duration) *GetSMSEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Sms events params
func (o *GetSMSEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Sms events params
func (o *GetSMSEventsParams) WithContext(ctx context.Context) *GetSMSEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Sms events params
func (o *GetSMSEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Sms events params
func (o *GetSMSEventsParams) WithHTTPClient(client *http.Client) *GetSMSEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Sms events params
func (o *GetSMSEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDays adds the days to the get Sms events params
func (o *GetSMSEventsParams) WithDays(days *int64) *GetSMSEventsParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the get Sms events params
func (o *GetSMSEventsParams) SetDays(days *int64) {
	o.Days = days
}

// WithEndDate adds the endDate to the get Sms events params
func (o *GetSMSEventsParams) WithEndDate(endDate *strfmt.Date) *GetSMSEventsParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get Sms events params
func (o *GetSMSEventsParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithEvent adds the event to the get Sms events params
func (o *GetSMSEventsParams) WithEvent(event *string) *GetSMSEventsParams {
	o.SetEvent(event)
	return o
}

// SetEvent adds the event to the get Sms events params
func (o *GetSMSEventsParams) SetEvent(event *string) {
	o.Event = event
}

// WithLimit adds the limit to the get Sms events params
func (o *GetSMSEventsParams) WithLimit(limit *int64) *GetSMSEventsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get Sms events params
func (o *GetSMSEventsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get Sms events params
func (o *GetSMSEventsParams) WithOffset(offset *int64) *GetSMSEventsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get Sms events params
func (o *GetSMSEventsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPhoneNumber adds the phoneNumber to the get Sms events params
func (o *GetSMSEventsParams) WithPhoneNumber(phoneNumber *string) *GetSMSEventsParams {
	o.SetPhoneNumber(phoneNumber)
	return o
}

// SetPhoneNumber adds the phoneNumber to the get Sms events params
func (o *GetSMSEventsParams) SetPhoneNumber(phoneNumber *string) {
	o.PhoneNumber = phoneNumber
}

// WithStartDate adds the startDate to the get Sms events params
func (o *GetSMSEventsParams) WithStartDate(startDate *strfmt.Date) *GetSMSEventsParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get Sms events params
func (o *GetSMSEventsParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithTags adds the tags to the get Sms events params
func (o *GetSMSEventsParams) WithTags(tags *string) *GetSMSEventsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get Sms events params
func (o *GetSMSEventsParams) SetTags(tags *string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetSMSEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Event != nil {

		// query param event
		var qrEvent string
		if o.Event != nil {
			qrEvent = *o.Event
		}
		qEvent := qrEvent
		if qEvent != "" {
			if err := r.SetQueryParam("event", qEvent); err != nil {
				return err
			}
		}

	}

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

	if o.PhoneNumber != nil {

		// query param phoneNumber
		var qrPhoneNumber string
		if o.PhoneNumber != nil {
			qrPhoneNumber = *o.PhoneNumber
		}
		qPhoneNumber := qrPhoneNumber
		if qPhoneNumber != "" {
			if err := r.SetQueryParam("phoneNumber", qPhoneNumber); err != nil {
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

	if o.Tags != nil {

		// query param tags
		var qrTags string
		if o.Tags != nil {
			qrTags = *o.Tags
		}
		qTags := qrTags
		if qTags != "" {
			if err := r.SetQueryParam("tags", qTags); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}