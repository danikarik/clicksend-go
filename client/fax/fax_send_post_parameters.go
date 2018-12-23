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

// NewFAXSendPostParams creates a new FAXSendPostParams object
// with the default values initialized.
func NewFAXSendPostParams() *FAXSendPostParams {
	var ()
	return &FAXSendPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXSendPostParamsWithTimeout creates a new FAXSendPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXSendPostParamsWithTimeout(timeout time.Duration) *FAXSendPostParams {
	var ()
	return &FAXSendPostParams{

		timeout: timeout,
	}
}

// NewFAXSendPostParamsWithContext creates a new FAXSendPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXSendPostParamsWithContext(ctx context.Context) *FAXSendPostParams {
	var ()
	return &FAXSendPostParams{

		Context: ctx,
	}
}

// NewFAXSendPostParamsWithHTTPClient creates a new FAXSendPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXSendPostParamsWithHTTPClient(client *http.Client) *FAXSendPostParams {
	var ()
	return &FAXSendPostParams{
		HTTPClient: client,
	}
}

/*FAXSendPostParams contains all the parameters to send to the API endpoint
for the Fax send post operation typically these are written to a http.Request
*/
type FAXSendPostParams struct {

	/*FAXMessage
	  FaxMessageCollection model

	*/
	FAXMessage models.FAXMessageCollection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Fax send post params
func (o *FAXSendPostParams) WithTimeout(timeout time.Duration) *FAXSendPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax send post params
func (o *FAXSendPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax send post params
func (o *FAXSendPostParams) WithContext(ctx context.Context) *FAXSendPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax send post params
func (o *FAXSendPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax send post params
func (o *FAXSendPostParams) WithHTTPClient(client *http.Client) *FAXSendPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax send post params
func (o *FAXSendPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFAXMessage adds the fAXMessage to the Fax send post params
func (o *FAXSendPostParams) WithFAXMessage(fAXMessage models.FAXMessageCollection) *FAXSendPostParams {
	o.SetFAXMessage(fAXMessage)
	return o
}

// SetFAXMessage adds the faxMessage to the Fax send post params
func (o *FAXSendPostParams) SetFAXMessage(fAXMessage models.FAXMessageCollection) {
	o.FAXMessage = fAXMessage
}

// WriteToRequest writes these params to a swagger request
func (o *FAXSendPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.FAXMessage); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
