// Code generated by go-swagger; DO NOT EDIT.

package sms_delivery_receipt_rules

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewSMSDeliveryReceiptAutomationPutParams creates a new SMSDeliveryReceiptAutomationPutParams object
// with the default values initialized.
func NewSMSDeliveryReceiptAutomationPutParams() *SMSDeliveryReceiptAutomationPutParams {
	var ()
	return &SMSDeliveryReceiptAutomationPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSDeliveryReceiptAutomationPutParamsWithTimeout creates a new SMSDeliveryReceiptAutomationPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSDeliveryReceiptAutomationPutParamsWithTimeout(timeout time.Duration) *SMSDeliveryReceiptAutomationPutParams {
	var ()
	return &SMSDeliveryReceiptAutomationPutParams{

		timeout: timeout,
	}
}

// NewSMSDeliveryReceiptAutomationPutParamsWithContext creates a new SMSDeliveryReceiptAutomationPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSDeliveryReceiptAutomationPutParamsWithContext(ctx context.Context) *SMSDeliveryReceiptAutomationPutParams {
	var ()
	return &SMSDeliveryReceiptAutomationPutParams{

		Context: ctx,
	}
}

// NewSMSDeliveryReceiptAutomationPutParamsWithHTTPClient creates a new SMSDeliveryReceiptAutomationPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSDeliveryReceiptAutomationPutParamsWithHTTPClient(client *http.Client) *SMSDeliveryReceiptAutomationPutParams {
	var ()
	return &SMSDeliveryReceiptAutomationPutParams{
		HTTPClient: client,
	}
}

/*SMSDeliveryReceiptAutomationPutParams contains all the parameters to send to the API endpoint
for the Sms delivery receipt automation put operation typically these are written to a http.Request
*/
type SMSDeliveryReceiptAutomationPutParams struct {

	/*DeliveryReceiptRule
	  Delivery receipt rule model

	*/
	DeliveryReceiptRule models.DeliveryReceiptRule
	/*ReceiptRuleID
	  Receipt rule id

	*/
	ReceiptRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) WithTimeout(timeout time.Duration) *SMSDeliveryReceiptAutomationPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) WithContext(ctx context.Context) *SMSDeliveryReceiptAutomationPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) WithHTTPClient(client *http.Client) *SMSDeliveryReceiptAutomationPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeliveryReceiptRule adds the deliveryReceiptRule to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) WithDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) *SMSDeliveryReceiptAutomationPutParams {
	o.SetDeliveryReceiptRule(deliveryReceiptRule)
	return o
}

// SetDeliveryReceiptRule adds the deliveryReceiptRule to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) SetDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) {
	o.DeliveryReceiptRule = deliveryReceiptRule
}

// WithReceiptRuleID adds the receiptRuleID to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) WithReceiptRuleID(receiptRuleID int32) *SMSDeliveryReceiptAutomationPutParams {
	o.SetReceiptRuleID(receiptRuleID)
	return o
}

// SetReceiptRuleID adds the receiptRuleId to the Sms delivery receipt automation put params
func (o *SMSDeliveryReceiptAutomationPutParams) SetReceiptRuleID(receiptRuleID int32) {
	o.ReceiptRuleID = receiptRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSDeliveryReceiptAutomationPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DeliveryReceiptRule); err != nil {
		return err
	}

	// path param receipt_rule_id
	if err := r.SetPathParam("receipt_rule_id", swag.FormatInt32(o.ReceiptRuleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
