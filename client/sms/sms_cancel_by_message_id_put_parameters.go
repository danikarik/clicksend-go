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
)

// NewSMSCancelByMessageIDPutParams creates a new SMSCancelByMessageIDPutParams object
// with the default values initialized.
func NewSMSCancelByMessageIDPutParams() *SMSCancelByMessageIDPutParams {
	var ()
	return &SMSCancelByMessageIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSCancelByMessageIDPutParamsWithTimeout creates a new SMSCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSCancelByMessageIDPutParamsWithTimeout(timeout time.Duration) *SMSCancelByMessageIDPutParams {
	var ()
	return &SMSCancelByMessageIDPutParams{

		timeout: timeout,
	}
}

// NewSMSCancelByMessageIDPutParamsWithContext creates a new SMSCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSCancelByMessageIDPutParamsWithContext(ctx context.Context) *SMSCancelByMessageIDPutParams {
	var ()
	return &SMSCancelByMessageIDPutParams{

		Context: ctx,
	}
}

// NewSMSCancelByMessageIDPutParamsWithHTTPClient creates a new SMSCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSCancelByMessageIDPutParamsWithHTTPClient(client *http.Client) *SMSCancelByMessageIDPutParams {
	var ()
	return &SMSCancelByMessageIDPutParams{
		HTTPClient: client,
	}
}

/*SMSCancelByMessageIDPutParams contains all the parameters to send to the API endpoint
for the Sms cancel by message Id put operation typically these are written to a http.Request
*/
type SMSCancelByMessageIDPutParams struct {

	/*MessageID
	  The message ID you want to cancel

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) WithTimeout(timeout time.Duration) *SMSCancelByMessageIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) WithContext(ctx context.Context) *SMSCancelByMessageIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) WithHTTPClient(client *http.Client) *SMSCancelByMessageIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageID adds the messageID to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) WithMessageID(messageID string) *SMSCancelByMessageIDPutParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the Sms cancel by message Id put params
func (o *SMSCancelByMessageIDPutParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSCancelByMessageIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param message_id
	if err := r.SetPathParam("message_id", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
