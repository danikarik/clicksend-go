// Code generated by go-swagger; DO NOT EDIT.

package email_to_sms

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

// NewSMSEmailSMSPostParams creates a new SMSEmailSMSPostParams object
// with the default values initialized.
func NewSMSEmailSMSPostParams() *SMSEmailSMSPostParams {
	var ()
	return &SMSEmailSMSPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSEmailSMSPostParamsWithTimeout creates a new SMSEmailSMSPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSEmailSMSPostParamsWithTimeout(timeout time.Duration) *SMSEmailSMSPostParams {
	var ()
	return &SMSEmailSMSPostParams{

		timeout: timeout,
	}
}

// NewSMSEmailSMSPostParamsWithContext creates a new SMSEmailSMSPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSEmailSMSPostParamsWithContext(ctx context.Context) *SMSEmailSMSPostParams {
	var ()
	return &SMSEmailSMSPostParams{

		Context: ctx,
	}
}

// NewSMSEmailSMSPostParamsWithHTTPClient creates a new SMSEmailSMSPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSEmailSMSPostParamsWithHTTPClient(client *http.Client) *SMSEmailSMSPostParams {
	var ()
	return &SMSEmailSMSPostParams{
		HTTPClient: client,
	}
}

/*SMSEmailSMSPostParams contains all the parameters to send to the API endpoint
for the Sms email Sms post operation typically these are written to a http.Request
*/
type SMSEmailSMSPostParams struct {

	/*EmailSMSAddress
	  EmailSMSAddress model

	*/
	EmailSMSAddress models.EmailSMSAddress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) WithTimeout(timeout time.Duration) *SMSEmailSMSPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) WithContext(ctx context.Context) *SMSEmailSMSPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) WithHTTPClient(client *http.Client) *SMSEmailSMSPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailSMSAddress adds the emailSMSAddress to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) WithEmailSMSAddress(emailSMSAddress models.EmailSMSAddress) *SMSEmailSMSPostParams {
	o.SetEmailSMSAddress(emailSMSAddress)
	return o
}

// SetEmailSMSAddress adds the emailSmsAddress to the Sms email Sms post params
func (o *SMSEmailSMSPostParams) SetEmailSMSAddress(emailSMSAddress models.EmailSMSAddress) {
	o.EmailSMSAddress = emailSMSAddress
}

// WriteToRequest writes these params to a swagger request
func (o *SMSEmailSMSPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.EmailSMSAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
