// Code generated by go-swagger; DO NOT EDIT.

package sms_delivery_receipt_rules

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

// NewSMSDeliveryReceiptAutomationsGetParams creates a new SMSDeliveryReceiptAutomationsGetParams object
// with the default values initialized.
func NewSMSDeliveryReceiptAutomationsGetParams() *SMSDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSDeliveryReceiptAutomationsGetParamsWithTimeout creates a new SMSDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSDeliveryReceiptAutomationsGetParamsWithTimeout(timeout time.Duration) *SMSDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewSMSDeliveryReceiptAutomationsGetParamsWithContext creates a new SMSDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSDeliveryReceiptAutomationsGetParamsWithContext(ctx context.Context) *SMSDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewSMSDeliveryReceiptAutomationsGetParamsWithHTTPClient creates a new SMSDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSDeliveryReceiptAutomationsGetParamsWithHTTPClient(client *http.Client) *SMSDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSDeliveryReceiptAutomationsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*SMSDeliveryReceiptAutomationsGetParams contains all the parameters to send to the API endpoint
for the Sms delivery receipt automations get operation typically these are written to a http.Request
*/
type SMSDeliveryReceiptAutomationsGetParams struct {

	/*Limit
	  Number of records per page

	*/
	Limit *int32
	/*Page
	  Page number

	*/
	Page *int32
	/*Q
	  Your keyword or query.

	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithTimeout(timeout time.Duration) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithContext(ctx context.Context) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithHTTPClient(client *http.Client) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithLimit(limit *int32) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithPage(page *int32) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) WithQ(q *string) *SMSDeliveryReceiptAutomationsGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the Sms delivery receipt automations get params
func (o *SMSDeliveryReceiptAutomationsGetParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *SMSDeliveryReceiptAutomationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
