// Code generated by go-swagger; DO NOT EDIT.

package inbound_sms_rules

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

// NewSMSInboundAutomationDeleteParams creates a new SMSInboundAutomationDeleteParams object
// with the default values initialized.
func NewSMSInboundAutomationDeleteParams() *SMSInboundAutomationDeleteParams {
	var ()
	return &SMSInboundAutomationDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSInboundAutomationDeleteParamsWithTimeout creates a new SMSInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSInboundAutomationDeleteParamsWithTimeout(timeout time.Duration) *SMSInboundAutomationDeleteParams {
	var ()
	return &SMSInboundAutomationDeleteParams{

		timeout: timeout,
	}
}

// NewSMSInboundAutomationDeleteParamsWithContext creates a new SMSInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSInboundAutomationDeleteParamsWithContext(ctx context.Context) *SMSInboundAutomationDeleteParams {
	var ()
	return &SMSInboundAutomationDeleteParams{

		Context: ctx,
	}
}

// NewSMSInboundAutomationDeleteParamsWithHTTPClient creates a new SMSInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSInboundAutomationDeleteParamsWithHTTPClient(client *http.Client) *SMSInboundAutomationDeleteParams {
	var ()
	return &SMSInboundAutomationDeleteParams{
		HTTPClient: client,
	}
}

/*SMSInboundAutomationDeleteParams contains all the parameters to send to the API endpoint
for the Sms inbound automation delete operation typically these are written to a http.Request
*/
type SMSInboundAutomationDeleteParams struct {

	/*InboundRuleID
	  Inbound rule id

	*/
	InboundRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) WithTimeout(timeout time.Duration) *SMSInboundAutomationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) WithContext(ctx context.Context) *SMSInboundAutomationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) WithHTTPClient(client *http.Client) *SMSInboundAutomationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundRuleID adds the inboundRuleID to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) WithInboundRuleID(inboundRuleID int32) *SMSInboundAutomationDeleteParams {
	o.SetInboundRuleID(inboundRuleID)
	return o
}

// SetInboundRuleID adds the inboundRuleId to the Sms inbound automation delete params
func (o *SMSInboundAutomationDeleteParams) SetInboundRuleID(inboundRuleID int32) {
	o.InboundRuleID = inboundRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSInboundAutomationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
