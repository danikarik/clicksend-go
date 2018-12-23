// Code generated by go-swagger; DO NOT EDIT.

package fax_delivery_receipt_rules

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

// NewFAXDeliveryReceiptAutomationsGetParams creates a new FAXDeliveryReceiptAutomationsGetParams object
// with the default values initialized.
func NewFAXDeliveryReceiptAutomationsGetParams() *FAXDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXDeliveryReceiptAutomationsGetParamsWithTimeout creates a new FAXDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXDeliveryReceiptAutomationsGetParamsWithTimeout(timeout time.Duration) *FAXDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewFAXDeliveryReceiptAutomationsGetParamsWithContext creates a new FAXDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXDeliveryReceiptAutomationsGetParamsWithContext(ctx context.Context) *FAXDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXDeliveryReceiptAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewFAXDeliveryReceiptAutomationsGetParamsWithHTTPClient creates a new FAXDeliveryReceiptAutomationsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXDeliveryReceiptAutomationsGetParamsWithHTTPClient(client *http.Client) *FAXDeliveryReceiptAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXDeliveryReceiptAutomationsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*FAXDeliveryReceiptAutomationsGetParams contains all the parameters to send to the API endpoint
for the Fax delivery receipt automations get operation typically these are written to a http.Request
*/
type FAXDeliveryReceiptAutomationsGetParams struct {

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

// WithTimeout adds the timeout to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithTimeout(timeout time.Duration) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithContext(ctx context.Context) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithHTTPClient(client *http.Client) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithLimit(limit *int32) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithPage(page *int32) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) WithQ(q *string) *FAXDeliveryReceiptAutomationsGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the Fax delivery receipt automations get params
func (o *FAXDeliveryReceiptAutomationsGetParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *FAXDeliveryReceiptAutomationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
