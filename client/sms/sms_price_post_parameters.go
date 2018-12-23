// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/danikarik/clicksend-go/models"
)

// NewSMSPricePostParams creates a new SMSPricePostParams object
// with the default values initialized.
func NewSMSPricePostParams() *SMSPricePostParams {
	var ()
	return &SMSPricePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSPricePostParamsWithTimeout creates a new SMSPricePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSPricePostParamsWithTimeout(timeout time.Duration) *SMSPricePostParams {
	var ()
	return &SMSPricePostParams{

		timeout: timeout,
	}
}

// NewSMSPricePostParamsWithContext creates a new SMSPricePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSPricePostParamsWithContext(ctx context.Context) *SMSPricePostParams {
	var ()
	return &SMSPricePostParams{

		Context: ctx,
	}
}

// NewSMSPricePostParamsWithHTTPClient creates a new SMSPricePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSPricePostParamsWithHTTPClient(client *http.Client) *SMSPricePostParams {
	var ()
	return &SMSPricePostParams{
		HTTPClient: client,
	}
}

/*SMSPricePostParams contains all the parameters to send to the API endpoint
for the Sms price post operation typically these are written to a http.Request
*/
type SMSPricePostParams struct {

	/*SMSMessages
	  SmsMessageCollection model

	*/
	SMSMessages models.SMSMessageCollection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms price post params
func (o *SMSPricePostParams) WithTimeout(timeout time.Duration) *SMSPricePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms price post params
func (o *SMSPricePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms price post params
func (o *SMSPricePostParams) WithContext(ctx context.Context) *SMSPricePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms price post params
func (o *SMSPricePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms price post params
func (o *SMSPricePostParams) WithHTTPClient(client *http.Client) *SMSPricePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms price post params
func (o *SMSPricePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMSMessages adds the sMSMessages to the Sms price post params
func (o *SMSPricePostParams) WithSMSMessages(sMSMessages models.SMSMessageCollection) *SMSPricePostParams {
	o.SetSMSMessages(sMSMessages)
	return o
}

// SetSMSMessages adds the smsMessages to the Sms price post params
func (o *SMSPricePostParams) SetSMSMessages(sMSMessages models.SMSMessageCollection) {
	o.SMSMessages = sMSMessages
}

// WriteToRequest writes these params to a swagger request
func (o *SMSPricePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.SMSMessages); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
