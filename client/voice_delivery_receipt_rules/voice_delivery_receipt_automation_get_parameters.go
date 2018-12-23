// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

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

// NewVoiceDeliveryReceiptAutomationGetParams creates a new VoiceDeliveryReceiptAutomationGetParams object
// with the default values initialized.
func NewVoiceDeliveryReceiptAutomationGetParams() *VoiceDeliveryReceiptAutomationGetParams {
	var ()
	return &VoiceDeliveryReceiptAutomationGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoiceDeliveryReceiptAutomationGetParamsWithTimeout creates a new VoiceDeliveryReceiptAutomationGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoiceDeliveryReceiptAutomationGetParamsWithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationGetParams {
	var ()
	return &VoiceDeliveryReceiptAutomationGetParams{

		timeout: timeout,
	}
}

// NewVoiceDeliveryReceiptAutomationGetParamsWithContext creates a new VoiceDeliveryReceiptAutomationGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoiceDeliveryReceiptAutomationGetParamsWithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationGetParams {
	var ()
	return &VoiceDeliveryReceiptAutomationGetParams{

		Context: ctx,
	}
}

// NewVoiceDeliveryReceiptAutomationGetParamsWithHTTPClient creates a new VoiceDeliveryReceiptAutomationGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoiceDeliveryReceiptAutomationGetParamsWithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationGetParams {
	var ()
	return &VoiceDeliveryReceiptAutomationGetParams{
		HTTPClient: client,
	}
}

/*VoiceDeliveryReceiptAutomationGetParams contains all the parameters to send to the API endpoint
for the voice delivery receipt automation get operation typically these are written to a http.Request
*/
type VoiceDeliveryReceiptAutomationGetParams struct {

	/*ReceiptRuleID
	  Receipt rule id

	*/
	ReceiptRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) WithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) WithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) WithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReceiptRuleID adds the receiptRuleID to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) WithReceiptRuleID(receiptRuleID int32) *VoiceDeliveryReceiptAutomationGetParams {
	o.SetReceiptRuleID(receiptRuleID)
	return o
}

// SetReceiptRuleID adds the receiptRuleId to the voice delivery receipt automation get params
func (o *VoiceDeliveryReceiptAutomationGetParams) SetReceiptRuleID(receiptRuleID int32) {
	o.ReceiptRuleID = receiptRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *VoiceDeliveryReceiptAutomationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
