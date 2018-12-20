// Code generated by go-swagger; DO NOT EDIT.

package account_recharge

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

// NewRechargeCreditCardGetParams creates a new RechargeCreditCardGetParams object
// with the default values initialized.
func NewRechargeCreditCardGetParams() *RechargeCreditCardGetParams {

	return &RechargeCreditCardGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRechargeCreditCardGetParamsWithTimeout creates a new RechargeCreditCardGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRechargeCreditCardGetParamsWithTimeout(timeout time.Duration) *RechargeCreditCardGetParams {

	return &RechargeCreditCardGetParams{

		timeout: timeout,
	}
}

// NewRechargeCreditCardGetParamsWithContext creates a new RechargeCreditCardGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRechargeCreditCardGetParamsWithContext(ctx context.Context) *RechargeCreditCardGetParams {

	return &RechargeCreditCardGetParams{

		Context: ctx,
	}
}

// NewRechargeCreditCardGetParamsWithHTTPClient creates a new RechargeCreditCardGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRechargeCreditCardGetParamsWithHTTPClient(client *http.Client) *RechargeCreditCardGetParams {

	return &RechargeCreditCardGetParams{
		HTTPClient: client,
	}
}

/*RechargeCreditCardGetParams contains all the parameters to send to the API endpoint
for the recharge credit card get operation typically these are written to a http.Request
*/
type RechargeCreditCardGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the recharge credit card get params
func (o *RechargeCreditCardGetParams) WithTimeout(timeout time.Duration) *RechargeCreditCardGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recharge credit card get params
func (o *RechargeCreditCardGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recharge credit card get params
func (o *RechargeCreditCardGetParams) WithContext(ctx context.Context) *RechargeCreditCardGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recharge credit card get params
func (o *RechargeCreditCardGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recharge credit card get params
func (o *RechargeCreditCardGetParams) WithHTTPClient(client *http.Client) *RechargeCreditCardGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recharge credit card get params
func (o *RechargeCreditCardGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RechargeCreditCardGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}