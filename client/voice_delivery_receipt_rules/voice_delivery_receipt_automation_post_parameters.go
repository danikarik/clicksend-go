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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/danikarik/clicksend-go/models"
)

// NewVoiceDeliveryReceiptAutomationPostParams creates a new VoiceDeliveryReceiptAutomationPostParams object
// with the default values initialized.
func NewVoiceDeliveryReceiptAutomationPostParams() *VoiceDeliveryReceiptAutomationPostParams {
	var ()
	return &VoiceDeliveryReceiptAutomationPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoiceDeliveryReceiptAutomationPostParamsWithTimeout creates a new VoiceDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoiceDeliveryReceiptAutomationPostParamsWithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationPostParams {
	var ()
	return &VoiceDeliveryReceiptAutomationPostParams{

		timeout: timeout,
	}
}

// NewVoiceDeliveryReceiptAutomationPostParamsWithContext creates a new VoiceDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoiceDeliveryReceiptAutomationPostParamsWithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationPostParams {
	var ()
	return &VoiceDeliveryReceiptAutomationPostParams{

		Context: ctx,
	}
}

// NewVoiceDeliveryReceiptAutomationPostParamsWithHTTPClient creates a new VoiceDeliveryReceiptAutomationPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoiceDeliveryReceiptAutomationPostParamsWithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationPostParams {
	var ()
	return &VoiceDeliveryReceiptAutomationPostParams{
		HTTPClient: client,
	}
}

/*VoiceDeliveryReceiptAutomationPostParams contains all the parameters to send to the API endpoint
for the voice delivery receipt automation post operation typically these are written to a http.Request
*/
type VoiceDeliveryReceiptAutomationPostParams struct {

	/*DeliveryReceiptRule
	  voice delivery receipt rule model

	*/
	DeliveryReceiptRule models.DeliveryReceiptRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) WithTimeout(timeout time.Duration) *VoiceDeliveryReceiptAutomationPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) WithContext(ctx context.Context) *VoiceDeliveryReceiptAutomationPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) WithHTTPClient(client *http.Client) *VoiceDeliveryReceiptAutomationPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeliveryReceiptRule adds the deliveryReceiptRule to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) WithDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) *VoiceDeliveryReceiptAutomationPostParams {
	o.SetDeliveryReceiptRule(deliveryReceiptRule)
	return o
}

// SetDeliveryReceiptRule adds the deliveryReceiptRule to the voice delivery receipt automation post params
func (o *VoiceDeliveryReceiptAutomationPostParams) SetDeliveryReceiptRule(deliveryReceiptRule models.DeliveryReceiptRule) {
	o.DeliveryReceiptRule = deliveryReceiptRule
}

// WriteToRequest writes these params to a swagger request
func (o *VoiceDeliveryReceiptAutomationPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
