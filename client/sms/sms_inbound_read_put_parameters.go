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

// NewSMSInboundReadPutParams creates a new SMSInboundReadPutParams object
// with the default values initialized.
func NewSMSInboundReadPutParams() *SMSInboundReadPutParams {
	var ()
	return &SMSInboundReadPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSInboundReadPutParamsWithTimeout creates a new SMSInboundReadPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSInboundReadPutParamsWithTimeout(timeout time.Duration) *SMSInboundReadPutParams {
	var ()
	return &SMSInboundReadPutParams{

		timeout: timeout,
	}
}

// NewSMSInboundReadPutParamsWithContext creates a new SMSInboundReadPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSInboundReadPutParamsWithContext(ctx context.Context) *SMSInboundReadPutParams {
	var ()
	return &SMSInboundReadPutParams{

		Context: ctx,
	}
}

// NewSMSInboundReadPutParamsWithHTTPClient creates a new SMSInboundReadPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSInboundReadPutParamsWithHTTPClient(client *http.Client) *SMSInboundReadPutParams {
	var ()
	return &SMSInboundReadPutParams{
		HTTPClient: client,
	}
}

/*SMSInboundReadPutParams contains all the parameters to send to the API endpoint
for the Sms inbound read put operation typically these are written to a http.Request
*/
type SMSInboundReadPutParams struct {

	/*DateBefore
	  DateBefore model

	*/
	DateBefore models.DateBefore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms inbound read put params
func (o *SMSInboundReadPutParams) WithTimeout(timeout time.Duration) *SMSInboundReadPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms inbound read put params
func (o *SMSInboundReadPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms inbound read put params
func (o *SMSInboundReadPutParams) WithContext(ctx context.Context) *SMSInboundReadPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms inbound read put params
func (o *SMSInboundReadPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms inbound read put params
func (o *SMSInboundReadPutParams) WithHTTPClient(client *http.Client) *SMSInboundReadPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms inbound read put params
func (o *SMSInboundReadPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateBefore adds the dateBefore to the Sms inbound read put params
func (o *SMSInboundReadPutParams) WithDateBefore(dateBefore models.DateBefore) *SMSInboundReadPutParams {
	o.SetDateBefore(dateBefore)
	return o
}

// SetDateBefore adds the dateBefore to the Sms inbound read put params
func (o *SMSInboundReadPutParams) SetDateBefore(dateBefore models.DateBefore) {
	o.DateBefore = dateBefore
}

// WriteToRequest writes these params to a swagger request
func (o *SMSInboundReadPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.DateBefore); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
