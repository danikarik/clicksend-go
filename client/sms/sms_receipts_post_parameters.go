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

// NewSMSReceiptsPostParams creates a new SMSReceiptsPostParams object
// with the default values initialized.
func NewSMSReceiptsPostParams() *SMSReceiptsPostParams {
	var ()
	return &SMSReceiptsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSReceiptsPostParamsWithTimeout creates a new SMSReceiptsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSReceiptsPostParamsWithTimeout(timeout time.Duration) *SMSReceiptsPostParams {
	var ()
	return &SMSReceiptsPostParams{

		timeout: timeout,
	}
}

// NewSMSReceiptsPostParamsWithContext creates a new SMSReceiptsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSReceiptsPostParamsWithContext(ctx context.Context) *SMSReceiptsPostParams {
	var ()
	return &SMSReceiptsPostParams{

		Context: ctx,
	}
}

// NewSMSReceiptsPostParamsWithHTTPClient creates a new SMSReceiptsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSReceiptsPostParamsWithHTTPClient(client *http.Client) *SMSReceiptsPostParams {
	var ()
	return &SMSReceiptsPostParams{
		HTTPClient: client,
	}
}

/*SMSReceiptsPostParams contains all the parameters to send to the API endpoint
for the Sms receipts post operation typically these are written to a http.Request
*/
type SMSReceiptsPostParams struct {

	/*URL
	  Url model

	*/
	URL models.URL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms receipts post params
func (o *SMSReceiptsPostParams) WithTimeout(timeout time.Duration) *SMSReceiptsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms receipts post params
func (o *SMSReceiptsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms receipts post params
func (o *SMSReceiptsPostParams) WithContext(ctx context.Context) *SMSReceiptsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms receipts post params
func (o *SMSReceiptsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms receipts post params
func (o *SMSReceiptsPostParams) WithHTTPClient(client *http.Client) *SMSReceiptsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms receipts post params
func (o *SMSReceiptsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithURL adds the url to the Sms receipts post params
func (o *SMSReceiptsPostParams) WithURL(url models.URL) *SMSReceiptsPostParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the Sms receipts post params
func (o *SMSReceiptsPostParams) SetURL(url models.URL) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *SMSReceiptsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.URL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
