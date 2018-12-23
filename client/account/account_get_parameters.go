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
)

// NewAccountGetParams creates a new AccountGetParams object
// with the default values initialized.
func NewAccountGetParams() *AccountGetParams {

	return &AccountGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountGetParamsWithTimeout creates a new AccountGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountGetParamsWithTimeout(timeout time.Duration) *AccountGetParams {

	return &AccountGetParams{

		timeout: timeout,
	}
}

// NewAccountGetParamsWithContext creates a new AccountGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountGetParamsWithContext(ctx context.Context) *AccountGetParams {

	return &AccountGetParams{

		Context: ctx,
	}
}

// NewAccountGetParamsWithHTTPClient creates a new AccountGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountGetParamsWithHTTPClient(client *http.Client) *AccountGetParams {

	return &AccountGetParams{
		HTTPClient: client,
	}
}

/*AccountGetParams contains all the parameters to send to the API endpoint
for the account get operation typically these are written to a http.Request
*/
type AccountGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account get params
func (o *AccountGetParams) WithTimeout(timeout time.Duration) *AccountGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account get params
func (o *AccountGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account get params
func (o *AccountGetParams) WithContext(ctx context.Context) *AccountGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account get params
func (o *AccountGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account get params
func (o *AccountGetParams) WithHTTPClient(client *http.Client) *AccountGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account get params
func (o *AccountGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AccountGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
