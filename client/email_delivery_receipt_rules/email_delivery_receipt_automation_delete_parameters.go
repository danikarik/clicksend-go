// Code generated by go-swagger; DO NOT EDIT.

package email_delivery_receipt_rules

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

// NewEmailDeliveryReceiptAutomationDeleteParams creates a new EmailDeliveryReceiptAutomationDeleteParams object
// with the default values initialized.
func NewEmailDeliveryReceiptAutomationDeleteParams() *EmailDeliveryReceiptAutomationDeleteParams {
	var ()
	return &EmailDeliveryReceiptAutomationDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailDeliveryReceiptAutomationDeleteParamsWithTimeout creates a new EmailDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailDeliveryReceiptAutomationDeleteParamsWithTimeout(timeout time.Duration) *EmailDeliveryReceiptAutomationDeleteParams {
	var ()
	return &EmailDeliveryReceiptAutomationDeleteParams{

		timeout: timeout,
	}
}

// NewEmailDeliveryReceiptAutomationDeleteParamsWithContext creates a new EmailDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailDeliveryReceiptAutomationDeleteParamsWithContext(ctx context.Context) *EmailDeliveryReceiptAutomationDeleteParams {
	var ()
	return &EmailDeliveryReceiptAutomationDeleteParams{

		Context: ctx,
	}
}

// NewEmailDeliveryReceiptAutomationDeleteParamsWithHTTPClient creates a new EmailDeliveryReceiptAutomationDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailDeliveryReceiptAutomationDeleteParamsWithHTTPClient(client *http.Client) *EmailDeliveryReceiptAutomationDeleteParams {
	var ()
	return &EmailDeliveryReceiptAutomationDeleteParams{
		HTTPClient: client,
	}
}

/*EmailDeliveryReceiptAutomationDeleteParams contains all the parameters to send to the API endpoint
for the email delivery receipt automation delete operation typically these are written to a http.Request
*/
type EmailDeliveryReceiptAutomationDeleteParams struct {

	/*ReceiptRuleID
	  Receipt rule id

	*/
	ReceiptRuleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) WithTimeout(timeout time.Duration) *EmailDeliveryReceiptAutomationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) WithContext(ctx context.Context) *EmailDeliveryReceiptAutomationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) WithHTTPClient(client *http.Client) *EmailDeliveryReceiptAutomationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReceiptRuleID adds the receiptRuleID to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) WithReceiptRuleID(receiptRuleID int32) *EmailDeliveryReceiptAutomationDeleteParams {
	o.SetReceiptRuleID(receiptRuleID)
	return o
}

// SetReceiptRuleID adds the receiptRuleId to the email delivery receipt automation delete params
func (o *EmailDeliveryReceiptAutomationDeleteParams) SetReceiptRuleID(receiptRuleID int32) {
	o.ReceiptRuleID = receiptRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *EmailDeliveryReceiptAutomationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
