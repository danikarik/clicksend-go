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

// NewSMSTemplatesPostParams creates a new SMSTemplatesPostParams object
// with the default values initialized.
func NewSMSTemplatesPostParams() *SMSTemplatesPostParams {
	var ()
	return &SMSTemplatesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSTemplatesPostParamsWithTimeout creates a new SMSTemplatesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSTemplatesPostParamsWithTimeout(timeout time.Duration) *SMSTemplatesPostParams {
	var ()
	return &SMSTemplatesPostParams{

		timeout: timeout,
	}
}

// NewSMSTemplatesPostParamsWithContext creates a new SMSTemplatesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSTemplatesPostParamsWithContext(ctx context.Context) *SMSTemplatesPostParams {
	var ()
	return &SMSTemplatesPostParams{

		Context: ctx,
	}
}

// NewSMSTemplatesPostParamsWithHTTPClient creates a new SMSTemplatesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSTemplatesPostParamsWithHTTPClient(client *http.Client) *SMSTemplatesPostParams {
	var ()
	return &SMSTemplatesPostParams{
		HTTPClient: client,
	}
}

/*SMSTemplatesPostParams contains all the parameters to send to the API endpoint
for the Sms templates post operation typically these are written to a http.Request
*/
type SMSTemplatesPostParams struct {

	/*SMSTemplate
	  SmsTemplate model

	*/
	SMSTemplate models.SMSTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms templates post params
func (o *SMSTemplatesPostParams) WithTimeout(timeout time.Duration) *SMSTemplatesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms templates post params
func (o *SMSTemplatesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms templates post params
func (o *SMSTemplatesPostParams) WithContext(ctx context.Context) *SMSTemplatesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms templates post params
func (o *SMSTemplatesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms templates post params
func (o *SMSTemplatesPostParams) WithHTTPClient(client *http.Client) *SMSTemplatesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms templates post params
func (o *SMSTemplatesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMSTemplate adds the sMSTemplate to the Sms templates post params
func (o *SMSTemplatesPostParams) WithSMSTemplate(sMSTemplate models.SMSTemplate) *SMSTemplatesPostParams {
	o.SetSMSTemplate(sMSTemplate)
	return o
}

// SetSMSTemplate adds the smsTemplate to the Sms templates post params
func (o *SMSTemplatesPostParams) SetSMSTemplate(sMSTemplate models.SMSTemplate) {
	o.SMSTemplate = sMSTemplate
}

// WriteToRequest writes these params to a swagger request
func (o *SMSTemplatesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.SMSTemplate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
