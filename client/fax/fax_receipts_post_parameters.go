// Code generated by go-swagger; DO NOT EDIT.

package fax

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

// NewFAXReceiptsPostParams creates a new FAXReceiptsPostParams object
// with the default values initialized.
func NewFAXReceiptsPostParams() *FAXReceiptsPostParams {
	var ()
	return &FAXReceiptsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXReceiptsPostParamsWithTimeout creates a new FAXReceiptsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXReceiptsPostParamsWithTimeout(timeout time.Duration) *FAXReceiptsPostParams {
	var ()
	return &FAXReceiptsPostParams{

		timeout: timeout,
	}
}

// NewFAXReceiptsPostParamsWithContext creates a new FAXReceiptsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXReceiptsPostParamsWithContext(ctx context.Context) *FAXReceiptsPostParams {
	var ()
	return &FAXReceiptsPostParams{

		Context: ctx,
	}
}

// NewFAXReceiptsPostParamsWithHTTPClient creates a new FAXReceiptsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXReceiptsPostParamsWithHTTPClient(client *http.Client) *FAXReceiptsPostParams {
	var ()
	return &FAXReceiptsPostParams{
		HTTPClient: client,
	}
}

/*FAXReceiptsPostParams contains all the parameters to send to the API endpoint
for the Fax receipts post operation typically these are written to a http.Request
*/
type FAXReceiptsPostParams struct {

	/*URL
	  Url model

	*/
	URL models.URL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Fax receipts post params
func (o *FAXReceiptsPostParams) WithTimeout(timeout time.Duration) *FAXReceiptsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax receipts post params
func (o *FAXReceiptsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax receipts post params
func (o *FAXReceiptsPostParams) WithContext(ctx context.Context) *FAXReceiptsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax receipts post params
func (o *FAXReceiptsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax receipts post params
func (o *FAXReceiptsPostParams) WithHTTPClient(client *http.Client) *FAXReceiptsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax receipts post params
func (o *FAXReceiptsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithURL adds the url to the Fax receipts post params
func (o *FAXReceiptsPostParams) WithURL(url models.URL) *FAXReceiptsPostParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the Fax receipts post params
func (o *FAXReceiptsPostParams) SetURL(url models.URL) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *FAXReceiptsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
