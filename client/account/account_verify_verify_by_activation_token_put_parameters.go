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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAccountVerifyVerifyByActivationTokenPutParams creates a new AccountVerifyVerifyByActivationTokenPutParams object
// with the default values initialized.
func NewAccountVerifyVerifyByActivationTokenPutParams() *AccountVerifyVerifyByActivationTokenPutParams {
	var ()
	return &AccountVerifyVerifyByActivationTokenPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountVerifyVerifyByActivationTokenPutParamsWithTimeout creates a new AccountVerifyVerifyByActivationTokenPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountVerifyVerifyByActivationTokenPutParamsWithTimeout(timeout time.Duration) *AccountVerifyVerifyByActivationTokenPutParams {
	var ()
	return &AccountVerifyVerifyByActivationTokenPutParams{

		timeout: timeout,
	}
}

// NewAccountVerifyVerifyByActivationTokenPutParamsWithContext creates a new AccountVerifyVerifyByActivationTokenPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountVerifyVerifyByActivationTokenPutParamsWithContext(ctx context.Context) *AccountVerifyVerifyByActivationTokenPutParams {
	var ()
	return &AccountVerifyVerifyByActivationTokenPutParams{

		Context: ctx,
	}
}

// NewAccountVerifyVerifyByActivationTokenPutParamsWithHTTPClient creates a new AccountVerifyVerifyByActivationTokenPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountVerifyVerifyByActivationTokenPutParamsWithHTTPClient(client *http.Client) *AccountVerifyVerifyByActivationTokenPutParams {
	var ()
	return &AccountVerifyVerifyByActivationTokenPutParams{
		HTTPClient: client,
	}
}

/*AccountVerifyVerifyByActivationTokenPutParams contains all the parameters to send to the API endpoint
for the account verify verify by activation token put operation typically these are written to a http.Request
*/
type AccountVerifyVerifyByActivationTokenPutParams struct {

	/*ActivationToken*/
	ActivationToken int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) WithTimeout(timeout time.Duration) *AccountVerifyVerifyByActivationTokenPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) WithContext(ctx context.Context) *AccountVerifyVerifyByActivationTokenPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) WithHTTPClient(client *http.Client) *AccountVerifyVerifyByActivationTokenPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivationToken adds the activationToken to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) WithActivationToken(activationToken int32) *AccountVerifyVerifyByActivationTokenPutParams {
	o.SetActivationToken(activationToken)
	return o
}

// SetActivationToken adds the activationToken to the account verify verify by activation token put params
func (o *AccountVerifyVerifyByActivationTokenPutParams) SetActivationToken(activationToken int32) {
	o.ActivationToken = activationToken
}

// WriteToRequest writes these params to a swagger request
func (o *AccountVerifyVerifyByActivationTokenPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param activation_token
	if err := r.SetPathParam("activation_token", swag.FormatInt32(o.ActivationToken)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
