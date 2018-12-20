// Code generated by go-swagger; DO NOT EDIT.

package inbound_fax_rules

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

// NewFAXInboundAutomationsGetParams creates a new FAXInboundAutomationsGetParams object
// with the default values initialized.
func NewFAXInboundAutomationsGetParams() *FAXInboundAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXInboundAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXInboundAutomationsGetParamsWithTimeout creates a new FAXInboundAutomationsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXInboundAutomationsGetParamsWithTimeout(timeout time.Duration) *FAXInboundAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXInboundAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewFAXInboundAutomationsGetParamsWithContext creates a new FAXInboundAutomationsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXInboundAutomationsGetParamsWithContext(ctx context.Context) *FAXInboundAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXInboundAutomationsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewFAXInboundAutomationsGetParamsWithHTTPClient creates a new FAXInboundAutomationsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXInboundAutomationsGetParamsWithHTTPClient(client *http.Client) *FAXInboundAutomationsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &FAXInboundAutomationsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*FAXInboundAutomationsGetParams contains all the parameters to send to the API endpoint
for the Fax inbound automations get operation typically these are written to a http.Request
*/
type FAXInboundAutomationsGetParams struct {

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

// WithTimeout adds the timeout to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithTimeout(timeout time.Duration) *FAXInboundAutomationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithContext(ctx context.Context) *FAXInboundAutomationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithHTTPClient(client *http.Client) *FAXInboundAutomationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithLimit(limit *int32) *FAXInboundAutomationsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithPage(page *int32) *FAXInboundAutomationsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WithQ adds the q to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) WithQ(q *string) *FAXInboundAutomationsGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the Fax inbound automations get params
func (o *FAXInboundAutomationsGetParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *FAXInboundAutomationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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