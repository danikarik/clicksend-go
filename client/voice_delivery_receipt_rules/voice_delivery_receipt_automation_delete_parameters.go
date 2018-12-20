// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

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

// NewVoiceDeliveryReceiptAutomationDeleteParams creates a new VoiceDeliveryReceiptAutomationDeleteParams object
// with the default values initialized.
func NewVoiceDeliveryReceiptAutomationDeleteParams() *VoiceDeliveryReceiptAutomationDeleteParams {
	var ()
	return &VoiceDeliveryReceiptAutomationDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoiceDeliveryReceiptAutomationDeleteParamsWithTimeout creates a new VoiceDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoiceDeliveryReceiptAutomationDeleteParamsWithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationDeleteParams {
	var ()
	return &VoiceDeliveryReceiptAutomationDeleteParams{

		timeout: timeout,
	}
}

// NewVoiceDeliveryReceiptAutomationDeleteParamsWithContext creates a new VoiceDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoiceDeliveryReceiptAutomationDeleteParamsWithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationDeleteParams {
	var ()
	return &VoiceDeliveryReceiptAutomationDeleteParams{

		Context: ctx,
	}
}

// NewVoiceDeliveryReceiptAutomationDeleteParamsWithHTTPClient creates a new VoiceDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoiceDeliveryReceiptAutomationDeleteParamsWithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationDeleteParams {
	var ()
	return &VoiceDeliveryReceiptAutomationDeleteParams{
		HTTPClient: client,
	}
}

/*VoiceDeliveryReceiptAutomationDeleteParams contains all the parameters to send to the API endpoint
for the voice delivery receipt automation delete operation typically these are written to a http.Request
*/
type VoiceDeliveryReceiptAutomationDeleteParams struct {

	/*ReceiptRuleID
	  Receipt rule id

	*/
	ReceiptRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) WithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) WithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) WithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReceiptRuleID adds the receiptRuleID to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) WithReceiptRuleID(receiptRuleID int32) *VoiceDeliveryReceiptAutomationDeleteParams {
	o.SetReceiptRuleID(receiptRuleID)
	return o
}

// SetReceiptRuleID adds the receiptRuleId to the voice delivery receipt automation delete params
func (o *VoiceDeliveryReceiptAutomationDeleteParams) SetReceiptRuleID(receiptRuleID int32) {
	o.ReceiptRuleID = receiptRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *VoiceDeliveryReceiptAutomationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param receipt_rule_id
	if err := r.SetPathParam("receipt_rule_id", swag.FormatInt32(o.ReceiptRuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}