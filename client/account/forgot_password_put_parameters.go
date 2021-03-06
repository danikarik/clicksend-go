// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewForgotPasswordPutParams creates a new ForgotPasswordPutParams object
// with the default values initialized.
func NewForgotPasswordPutParams() *ForgotPasswordPutParams {
	var ()
	return &ForgotPasswordPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewForgotPasswordPutParamsWithTimeout creates a new ForgotPasswordPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForgotPasswordPutParamsWithTimeout(timeout time.Duration) *ForgotPasswordPutParams {
	var ()
	return &ForgotPasswordPutParams{

		timeout: timeout,
	}
}

// NewForgotPasswordPutParamsWithContext creates a new ForgotPasswordPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewForgotPasswordPutParamsWithContext(ctx context.Context) *ForgotPasswordPutParams {
	var ()
	return &ForgotPasswordPutParams{

		Context: ctx,
	}
}

// NewForgotPasswordPutParamsWithHTTPClient creates a new ForgotPasswordPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForgotPasswordPutParamsWithHTTPClient(client *http.Client) *ForgotPasswordPutParams {
	var ()
	return &ForgotPasswordPutParams{
		HTTPClient: client,
	}
}

/*ForgotPasswordPutParams contains all the parameters to send to the API endpoint
for the forgot password put operation typically these are written to a http.Request
*/
type ForgotPasswordPutParams struct {

	/*ForgotPassword*/
	ForgotPassword *models.ForgotPassword

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forgot password put params
func (o *ForgotPasswordPutParams) WithTimeout(timeout time.Duration) *ForgotPasswordPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forgot password put params
func (o *ForgotPasswordPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forgot password put params
func (o *ForgotPasswordPutParams) WithContext(ctx context.Context) *ForgotPasswordPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forgot password put params
func (o *ForgotPasswordPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forgot password put params
func (o *ForgotPasswordPutParams) WithHTTPClient(client *http.Client) *ForgotPasswordPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forgot password put params
func (o *ForgotPasswordPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForgotPassword adds the forgotPassword to the forgot password put params
func (o *ForgotPasswordPutParams) WithForgotPassword(forgotPassword *models.ForgotPassword) *ForgotPasswordPutParams {
	o.SetForgotPassword(forgotPassword)
	return o
}

// SetForgotPassword adds the forgotPassword to the forgot password put params
func (o *ForgotPasswordPutParams) SetForgotPassword(forgotPassword *models.ForgotPassword) {
	o.ForgotPassword = forgotPassword
}

// WriteToRequest writes these params to a swagger request
func (o *ForgotPasswordPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForgotPassword != nil {
		if err := r.SetBodyParam(o.ForgotPassword); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
