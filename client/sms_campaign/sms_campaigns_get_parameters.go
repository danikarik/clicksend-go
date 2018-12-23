// Code generated by go-swagger; DO NOT EDIT.

package sms_campaign

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

// NewSMSCampaignsGetParams creates a new SMSCampaignsGetParams object
// with the default values initialized.
func NewSMSCampaignsGetParams() *SMSCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSCampaignsGetParamsWithTimeout creates a new SMSCampaignsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSCampaignsGetParamsWithTimeout(timeout time.Duration) *SMSCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewSMSCampaignsGetParamsWithContext creates a new SMSCampaignsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSCampaignsGetParamsWithContext(ctx context.Context) *SMSCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewSMSCampaignsGetParamsWithHTTPClient creates a new SMSCampaignsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSCampaignsGetParamsWithHTTPClient(client *http.Client) *SMSCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &SMSCampaignsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*SMSCampaignsGetParams contains all the parameters to send to the API endpoint
for the Sms campaigns get operation typically these are written to a http.Request
*/
type SMSCampaignsGetParams struct {

	/*Limit
	  Number of records per page

	*/
	Limit *int32
	/*Page
	  Page number

	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms campaigns get params
func (o *SMSCampaignsGetParams) WithTimeout(timeout time.Duration) *SMSCampaignsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms campaigns get params
func (o *SMSCampaignsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms campaigns get params
func (o *SMSCampaignsGetParams) WithContext(ctx context.Context) *SMSCampaignsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms campaigns get params
func (o *SMSCampaignsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms campaigns get params
func (o *SMSCampaignsGetParams) WithHTTPClient(client *http.Client) *SMSCampaignsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms campaigns get params
func (o *SMSCampaignsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the Sms campaigns get params
func (o *SMSCampaignsGetParams) WithLimit(limit *int32) *SMSCampaignsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Sms campaigns get params
func (o *SMSCampaignsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Sms campaigns get params
func (o *SMSCampaignsGetParams) WithPage(page *int32) *SMSCampaignsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Sms campaigns get params
func (o *SMSCampaignsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *SMSCampaignsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
