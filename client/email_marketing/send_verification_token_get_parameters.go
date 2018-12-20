// Code generated by go-swagger; DO NOT EDIT.

package email_marketing

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

// NewSendVerificationTokenGetParams creates a new SendVerificationTokenGetParams object
// with the default values initialized.
func NewSendVerificationTokenGetParams() *SendVerificationTokenGetParams {
	var ()
	return &SendVerificationTokenGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendVerificationTokenGetParamsWithTimeout creates a new SendVerificationTokenGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendVerificationTokenGetParamsWithTimeout(timeout time.Duration) *SendVerificationTokenGetParams {
	var ()
	return &SendVerificationTokenGetParams{

		timeout: timeout,
	}
}

// NewSendVerificationTokenGetParamsWithContext creates a new SendVerificationTokenGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendVerificationTokenGetParamsWithContext(ctx context.Context) *SendVerificationTokenGetParams {
	var ()
	return &SendVerificationTokenGetParams{

		Context: ctx,
	}
}

// NewSendVerificationTokenGetParamsWithHTTPClient creates a new SendVerificationTokenGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendVerificationTokenGetParamsWithHTTPClient(client *http.Client) *SendVerificationTokenGetParams {
	var ()
	return &SendVerificationTokenGetParams{
		HTTPClient: client,
	}
}

/*SendVerificationTokenGetParams contains all the parameters to send to the API endpoint
for the send verification token get operation typically these are written to a http.Request
*/
type SendVerificationTokenGetParams struct {

	/*EmailAddressID
	  Allowed email address id

	*/
	EmailAddressID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send verification token get params
func (o *SendVerificationTokenGetParams) WithTimeout(timeout time.Duration) *SendVerificationTokenGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send verification token get params
func (o *SendVerificationTokenGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send verification token get params
func (o *SendVerificationTokenGetParams) WithContext(ctx context.Context) *SendVerificationTokenGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send verification token get params
func (o *SendVerificationTokenGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send verification token get params
func (o *SendVerificationTokenGetParams) WithHTTPClient(client *http.Client) *SendVerificationTokenGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send verification token get params
func (o *SendVerificationTokenGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailAddressID adds the emailAddressID to the send verification token get params
func (o *SendVerificationTokenGetParams) WithEmailAddressID(emailAddressID int32) *SendVerificationTokenGetParams {
	o.SetEmailAddressID(emailAddressID)
	return o
}

// SetEmailAddressID adds the emailAddressId to the send verification token get params
func (o *SendVerificationTokenGetParams) SetEmailAddressID(emailAddressID int32) {
	o.EmailAddressID = emailAddressID
}

// WriteToRequest writes these params to a swagger request
func (o *SendVerificationTokenGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email_address_id
	if err := r.SetPathParam("email_address_id", swag.FormatInt32(o.EmailAddressID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
