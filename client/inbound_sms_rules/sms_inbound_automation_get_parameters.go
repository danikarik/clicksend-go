// Code generated by go-swagger; DO NOT EDIT.

package inbound_sms_rules

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

// NewSMSInboundAutomationGetParams creates a new SMSInboundAutomationGetParams object
// with the default values initialized.
func NewSMSInboundAutomationGetParams() *SMSInboundAutomationGetParams {
	var ()
	return &SMSInboundAutomationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSInboundAutomationGetParamsWithTimeout creates a new SMSInboundAutomationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSInboundAutomationGetParamsWithTimeout(timeout time.Duration) *SMSInboundAutomationGetParams {
	var ()
	return &SMSInboundAutomationGetParams{

		timeout: timeout,
	}
}

// NewSMSInboundAutomationGetParamsWithContext creates a new SMSInboundAutomationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSInboundAutomationGetParamsWithContext(ctx context.Context) *SMSInboundAutomationGetParams {
	var ()
	return &SMSInboundAutomationGetParams{

		Context: ctx,
	}
}

// NewSMSInboundAutomationGetParamsWithHTTPClient creates a new SMSInboundAutomationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSInboundAutomationGetParamsWithHTTPClient(client *http.Client) *SMSInboundAutomationGetParams {
	var ()
	return &SMSInboundAutomationGetParams{
		HTTPClient: client,
	}
}

/*SMSInboundAutomationGetParams contains all the parameters to send to the API endpoint
for the Sms inbound automation get operation typically these are written to a http.Request
*/
type SMSInboundAutomationGetParams struct {

	/*InboundRuleID
	  Inbound rule id

	*/
	InboundRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) WithTimeout(timeout time.Duration) *SMSInboundAutomationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) WithContext(ctx context.Context) *SMSInboundAutomationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) WithHTTPClient(client *http.Client) *SMSInboundAutomationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundRuleID adds the inboundRuleID to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) WithInboundRuleID(inboundRuleID int32) *SMSInboundAutomationGetParams {
	o.SetInboundRuleID(inboundRuleID)
	return o
}

// SetInboundRuleID adds the inboundRuleId to the Sms inbound automation get params
func (o *SMSInboundAutomationGetParams) SetInboundRuleID(inboundRuleID int32) {
	o.InboundRuleID = inboundRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSInboundAutomationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inbound_rule_id
	if err := r.SetPathParam("inbound_rule_id", swag.FormatInt32(o.InboundRuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
