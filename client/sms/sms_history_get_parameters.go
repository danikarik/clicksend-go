// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSMSHistoryGetParams creates a new SMSHistoryGetParams object
// with the default values initialized.
func NewSMSHistoryGetParams() *SMSHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSHistoryGetParamsWithTimeout creates a new SMSHistoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSHistoryGetParamsWithTimeout(timeout time.Duration) *SMSHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewSMSHistoryGetParamsWithContext creates a new SMSHistoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSHistoryGetParamsWithContext(ctx context.Context) *SMSHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewSMSHistoryGetParamsWithHTTPClient creates a new SMSHistoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSHistoryGetParamsWithHTTPClient(client *http.Client) *SMSHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSHistoryGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*SMSHistoryGetParams contains all the parameters to send to the API endpoint
for the Sms history get operation typically these are written to a http.Request
*/
type SMSHistoryGetParams struct {

	/*DateFrom
	  Start date

	*/
	DateFrom *int32
	/*DateTo
	  End date

	*/
	DateTo *int32
	/*Limit
	  Number of records per page

	*/
	Limit *int32
	/*Page
	  Page number

	*/
	Page *int32
	/*Q
	  Custom query Example: from:{number},status_code:201.

	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms history get params
func (o *SMSHistoryGetParams) WithTimeout(timeout time.Duration) *SMSHistoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms history get params
func (o *SMSHistoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms history get params
func (o *SMSHistoryGetParams) WithContext(ctx context.Context) *SMSHistoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms history get params
func (o *SMSHistoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms history get params
func (o *SMSHistoryGetParams) WithHTTPClient(client *http.Client) *SMSHistoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms history get params
func (o *SMSHistoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateFrom adds the dateFrom to the Sms history get params
func (o *SMSHistoryGetParams) WithDateFrom(dateFrom *int32) *SMSHistoryGetParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the Sms history get params
func (o *SMSHistoryGetParams) SetDateFrom(dateFrom *int32) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the Sms history get params
func (o *SMSHistoryGetParams) WithDateTo(dateTo *int32) *SMSHistoryGetParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the Sms history get params
func (o *SMSHistoryGetParams) SetDateTo(dateTo *int32) {
	o.DateTo = dateTo
}

// WithLimit adds the limit to the Sms history get params
func (o *SMSHistoryGetParams) WithLimit(limit *int32) *SMSHistoryGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Sms history get params
func (o *SMSHistoryGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Sms history get params
func (o *SMSHistoryGetParams) WithPage(page *int32) *SMSHistoryGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Sms history get params
func (o *SMSHistoryGetParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the Sms history get params
func (o *SMSHistoryGetParams) WithQ(q *string) *SMSHistoryGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the Sms history get params
func (o *SMSHistoryGetParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *SMSHistoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateFrom != nil {

		// query param date_from
		var qrDateFrom int32
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := swag.FormatInt32(qrDateFrom)
		if qDateFrom != "" {
			if err := r.SetQueryParam("date_from", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param date_to
		var qrDateTo int32
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := swag.FormatInt32(qrDateTo)
		if qDateTo != "" {
			if err := r.SetQueryParam("date_to", qDateTo); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
