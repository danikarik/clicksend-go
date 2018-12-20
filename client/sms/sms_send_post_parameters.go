// Code generated by go-swagger; DO NOT EDIT.

package sms

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

// NewSMSSendPostParams creates a new SMSSendPostParams object
// with the default values initialized.
func NewSMSSendPostParams() *SMSSendPostParams {
	var ()
	return &SMSSendPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSSendPostParamsWithTimeout creates a new SMSSendPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSSendPostParamsWithTimeout(timeout time.Duration) *SMSSendPostParams {
	var ()
	return &SMSSendPostParams{

		timeout: timeout,
	}
}

// NewSMSSendPostParamsWithContext creates a new SMSSendPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSSendPostParamsWithContext(ctx context.Context) *SMSSendPostParams {
	var ()
	return &SMSSendPostParams{

		Context: ctx,
	}
}

// NewSMSSendPostParamsWithHTTPClient creates a new SMSSendPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSSendPostParamsWithHTTPClient(client *http.Client) *SMSSendPostParams {
	var ()
	return &SMSSendPostParams{
		HTTPClient: client,
	}
}

/*SMSSendPostParams contains all the parameters to send to the API endpoint
for the Sms send post operation typically these are written to a http.Request
*/
type SMSSendPostParams struct {

	/*SMSMessages
	  SmsMessageCollection model

	*/
	SMSMessages models.SMSMessageCollection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms send post params
func (o *SMSSendPostParams) WithTimeout(timeout time.Duration) *SMSSendPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms send post params
func (o *SMSSendPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms send post params
func (o *SMSSendPostParams) WithContext(ctx context.Context) *SMSSendPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms send post params
func (o *SMSSendPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms send post params
func (o *SMSSendPostParams) WithHTTPClient(client *http.Client) *SMSSendPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms send post params
func (o *SMSSendPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMSMessages adds the sMSMessages to the Sms send post params
func (o *SMSSendPostParams) WithSMSMessages(sMSMessages models.SMSMessageCollection) *SMSSendPostParams {
	o.SetSMSMessages(sMSMessages)
	return o
}

// SetSMSMessages adds the smsMessages to the Sms send post params
func (o *SMSSendPostParams) SetSMSMessages(sMSMessages models.SMSMessageCollection) {
	o.SMSMessages = sMSMessages
}

// WriteToRequest writes these params to a swagger request
func (o *SMSSendPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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