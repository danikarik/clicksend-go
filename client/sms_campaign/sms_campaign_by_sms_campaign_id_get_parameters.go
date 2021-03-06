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

// NewSMSCampaignBySMSCampaignIDGetParams creates a new SMSCampaignBySMSCampaignIDGetParams object
// with the default values initialized.
func NewSMSCampaignBySMSCampaignIDGetParams() *SMSCampaignBySMSCampaignIDGetParams {
	var ()
	return &SMSCampaignBySMSCampaignIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSCampaignBySMSCampaignIDGetParamsWithTimeout creates a new SMSCampaignBySMSCampaignIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSCampaignBySMSCampaignIDGetParamsWithTimeout(timeout time.Duration) *SMSCampaignBySMSCampaignIDGetParams {
	var ()
	return &SMSCampaignBySMSCampaignIDGetParams{

		timeout: timeout,
	}
}

// NewSMSCampaignBySMSCampaignIDGetParamsWithContext creates a new SMSCampaignBySMSCampaignIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSCampaignBySMSCampaignIDGetParamsWithContext(ctx context.Context) *SMSCampaignBySMSCampaignIDGetParams {
	var ()
	return &SMSCampaignBySMSCampaignIDGetParams{

		Context: ctx,
	}
}

// NewSMSCampaignBySMSCampaignIDGetParamsWithHTTPClient creates a new SMSCampaignBySMSCampaignIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSCampaignBySMSCampaignIDGetParamsWithHTTPClient(client *http.Client) *SMSCampaignBySMSCampaignIDGetParams {
	var ()
	return &SMSCampaignBySMSCampaignIDGetParams{
		HTTPClient: client,
	}
}

/*SMSCampaignBySMSCampaignIDGetParams contains all the parameters to send to the API endpoint
for the Sms campaign by Sms campaign Id get operation typically these are written to a http.Request
*/
type SMSCampaignBySMSCampaignIDGetParams struct {

	/*SMSCampaignID
	  ID of SMS campaign to retrieve

	*/
	SMSCampaignID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) WithTimeout(timeout time.Duration) *SMSCampaignBySMSCampaignIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) WithContext(ctx context.Context) *SMSCampaignBySMSCampaignIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) WithHTTPClient(client *http.Client) *SMSCampaignBySMSCampaignIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMSCampaignID adds the sMSCampaignID to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) WithSMSCampaignID(sMSCampaignID int32) *SMSCampaignBySMSCampaignIDGetParams {
	o.SetSMSCampaignID(sMSCampaignID)
	return o
}

// SetSMSCampaignID adds the smsCampaignId to the Sms campaign by Sms campaign Id get params
func (o *SMSCampaignBySMSCampaignIDGetParams) SetSMSCampaignID(sMSCampaignID int32) {
	o.SMSCampaignID = sMSCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSCampaignBySMSCampaignIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sms_campaign_id
	if err := r.SetPathParam("sms_campaign_id", swag.FormatInt32(o.SMSCampaignID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
