// Code generated by go-swagger; DO NOT EDIT.

package inbound_fax_rules

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

// NewFAXInboundAutomationDeleteParams creates a new FAXInboundAutomationDeleteParams object
// with the default values initialized.
func NewFAXInboundAutomationDeleteParams() *FAXInboundAutomationDeleteParams {
	var ()
	return &FAXInboundAutomationDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXInboundAutomationDeleteParamsWithTimeout creates a new FAXInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXInboundAutomationDeleteParamsWithTimeout(timeout time.Duration) *FAXInboundAutomationDeleteParams {
	var ()
	return &FAXInboundAutomationDeleteParams{

		timeout: timeout,
	}
}

// NewFAXInboundAutomationDeleteParamsWithContext creates a new FAXInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXInboundAutomationDeleteParamsWithContext(ctx context.Context) *FAXInboundAutomationDeleteParams {
	var ()
	return &FAXInboundAutomationDeleteParams{

		Context: ctx,
	}
}

// NewFAXInboundAutomationDeleteParamsWithHTTPClient creates a new FAXInboundAutomationDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXInboundAutomationDeleteParamsWithHTTPClient(client *http.Client) *FAXInboundAutomationDeleteParams {
	var ()
	return &FAXInboundAutomationDeleteParams{
		HTTPClient: client,
	}
}

/*FAXInboundAutomationDeleteParams contains all the parameters to send to the API endpoint
for the Fax inbound automation delete operation typically these are written to a http.Request
*/
type FAXInboundAutomationDeleteParams struct {

	/*InboundRuleID
	  Inbound rule id

	*/
	InboundRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) WithTimeout(timeout time.Duration) *FAXInboundAutomationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) WithContext(ctx context.Context) *FAXInboundAutomationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) WithHTTPClient(client *http.Client) *FAXInboundAutomationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundRuleID adds the inboundRuleID to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) WithInboundRuleID(inboundRuleID int32) *FAXInboundAutomationDeleteParams {
	o.SetInboundRuleID(inboundRuleID)
	return o
}

// SetInboundRuleID adds the inboundRuleId to the Fax inbound automation delete params
func (o *FAXInboundAutomationDeleteParams) SetInboundRuleID(inboundRuleID int32) {
	o.InboundRuleID = inboundRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *FAXInboundAutomationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
