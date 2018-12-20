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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/danikarik/clicksend-go/models"
)

// NewEmailDeliveryReceiptAutomationPostParams creates a new EmailDeliveryReceiptAutomationPostParams object
// with the default values initialized.
func NewEmailDeliveryReceiptAutomationPostParams() *EmailDeliveryReceiptAutomationPostParams {
	var ()
	return &EmailDeliveryReceiptAutomationPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailDeliveryReceiptAutomationPostParamsWithTimeout creates a new EmailDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailDeliveryReceiptAutomationPostParamsWithTimeout(timeout time.Duration) *EmailDeliveryReceiptAutomationPostParams {
	var ()
	return &EmailDeliveryReceiptAutomationPostParams{

		timeout: timeout,
	}
}

// NewEmailDeliveryReceiptAutomationPostParamsWithContext creates a new EmailDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailDeliveryReceiptAutomationPostParamsWithContext(ctx context.Context) *EmailDeliveryReceiptAutomationPostParams {
	var ()
	return &EmailDeliveryReceiptAutomationPostParams{

		Context: ctx,
	}
}

// NewEmailDeliveryReceiptAutomationPostParamsWithHTTPClient creates a new EmailDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailDeliveryReceiptAutomationPostParamsWithHTTPClient(client *http.Client) *EmailDeliveryReceiptAutomationPostParams {
	var ()
	return &EmailDeliveryReceiptAutomationPostParams{
		HTTPClient: client,
	}
}

/*EmailDeliveryReceiptAutomationPostParams contains all the parameters to send to the API endpoint
for the email delivery receipt automation post operation typically these are written to a http.Request
*/
type EmailDeliveryReceiptAutomationPostParams struct {

	/*DeliveryReceiptRule
	  Email delivery receipt rule model

	*/
	DeliveryReceiptRule models.DeliveryReceiptRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) WithTimeout(timeout time.Duration) *EmailDeliveryReceiptAutomationPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) WithContext(ctx context.Context) *EmailDeliveryReceiptAutomationPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) WithHTTPClient(client *http.Client) *EmailDeliveryReceiptAutomationPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeliveryReceiptRule adds the deliveryReceiptRule to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) WithDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) *EmailDeliveryReceiptAutomationPostParams {
	o.SetDeliveryReceiptRule(deliveryReceiptRule)
	return o
}

// SetDeliveryReceiptRule adds the deliveryReceiptRule to the email delivery receipt automation post params
func (o *EmailDeliveryReceiptAutomationPostParams) SetDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) {
	o.DeliveryReceiptRule = deliveryReceiptRule
}

// WriteToRequest writes these params to a swagger request
func (o *EmailDeliveryReceiptAutomationPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DeliveryReceiptRule); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}