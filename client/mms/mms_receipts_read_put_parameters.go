// Code generated by go-swagger; DO NOT EDIT.

package mms

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

// NewMMSReceiptsReadPutParams creates a new MMSReceiptsReadPutParams object
// with the default values initialized.
func NewMMSReceiptsReadPutParams() *MMSReceiptsReadPutParams {
	var ()
	return &MMSReceiptsReadPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMMSReceiptsReadPutParamsWithTimeout creates a new MMSReceiptsReadPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMMSReceiptsReadPutParamsWithTimeout(timeout time.Duration) *MMSReceiptsReadPutParams {
	var ()
	return &MMSReceiptsReadPutParams{

		timeout: timeout,
	}
}

// NewMMSReceiptsReadPutParamsWithContext creates a new MMSReceiptsReadPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewMMSReceiptsReadPutParamsWithContext(ctx context.Context) *MMSReceiptsReadPutParams {
	var ()
	return &MMSReceiptsReadPutParams{

		Context: ctx,
	}
}

// NewMMSReceiptsReadPutParamsWithHTTPClient creates a new MMSReceiptsReadPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMMSReceiptsReadPutParamsWithHTTPClient(client *http.Client) *MMSReceiptsReadPutParams {
	var ()
	return &MMSReceiptsReadPutParams{
		HTTPClient: client,
	}
}

/*MMSReceiptsReadPutParams contains all the parameters to send to the API endpoint
for the Mms receipts read put operation typically these are written to a http.Request
*/
type MMSReceiptsReadPutParams struct {

	/*DateBefore
	  DateBefore model

	*/
	DateBefore models.DateBefore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) WithTimeout(timeout time.Duration) *MMSReceiptsReadPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) WithContext(ctx context.Context) *MMSReceiptsReadPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) WithHTTPClient(client *http.Client) *MMSReceiptsReadPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateBefore adds the dateBefore to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) WithDateBefore(dateBefore models.DateBefore) *MMSReceiptsReadPutParams {
	o.SetDateBefore(dateBefore)
	return o
}

// SetDateBefore adds the dateBefore to the Mms receipts read put params
func (o *MMSReceiptsReadPutParams) SetDateBefore(dateBefore models.DateBefore) {
	o.DateBefore = dateBefore
}

// WriteToRequest writes these params to a swagger request
func (o *MMSReceiptsReadPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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