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

// NewFAXReceiptsReadPutParams creates a new FAXReceiptsReadPutParams object
// with the default values initialized.
func NewFAXReceiptsReadPutParams() *FAXReceiptsReadPutParams {
	var ()
	return &FAXReceiptsReadPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXReceiptsReadPutParamsWithTimeout creates a new FAXReceiptsReadPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXReceiptsReadPutParamsWithTimeout(timeout time.Duration) *FAXReceiptsReadPutParams {
	var ()
	return &FAXReceiptsReadPutParams{

		timeout: timeout,
	}
}

// NewFAXReceiptsReadPutParamsWithContext creates a new FAXReceiptsReadPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXReceiptsReadPutParamsWithContext(ctx context.Context) *FAXReceiptsReadPutParams {
	var ()
	return &FAXReceiptsReadPutParams{

		Context: ctx,
	}
}

// NewFAXReceiptsReadPutParamsWithHTTPClient creates a new FAXReceiptsReadPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXReceiptsReadPutParamsWithHTTPClient(client *http.Client) *FAXReceiptsReadPutParams {
	var ()
	return &FAXReceiptsReadPutParams{
		HTTPClient: client,
	}
}

/*FAXReceiptsReadPutParams contains all the parameters to send to the API endpoint
for the Fax receipts read put operation typically these are written to a http.Request
*/
type FAXReceiptsReadPutParams struct {

	/*DateBefore
	  DateBefore model

	*/
	DateBefore models.DateBefore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) WithTimeout(timeout time.Duration) *FAXReceiptsReadPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) WithContext(ctx context.Context) *FAXReceiptsReadPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) WithHTTPClient(client *http.Client) *FAXReceiptsReadPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateBefore adds the dateBefore to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) WithDateBefore(dateBefore models.DateBefore) *FAXReceiptsReadPutParams {
	o.SetDateBefore(dateBefore)
	return o
}

// SetDateBefore adds the dateBefore to the Fax receipts read put params
func (o *FAXReceiptsReadPutParams) SetDateBefore(dateBefore models.DateBefore) {
	o.DateBefore = dateBefore
}

// WriteToRequest writes these params to a swagger request
func (o *FAXReceiptsReadPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
