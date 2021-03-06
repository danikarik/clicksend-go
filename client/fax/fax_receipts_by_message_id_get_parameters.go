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
)

// NewFAXReceiptsByMessageIDGetParams creates a new FAXReceiptsByMessageIDGetParams object
// with the default values initialized.
func NewFAXReceiptsByMessageIDGetParams() *FAXReceiptsByMessageIDGetParams {
	var ()
	return &FAXReceiptsByMessageIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFAXReceiptsByMessageIDGetParamsWithTimeout creates a new FAXReceiptsByMessageIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFAXReceiptsByMessageIDGetParamsWithTimeout(timeout time.Duration) *FAXReceiptsByMessageIDGetParams {
	var ()
	return &FAXReceiptsByMessageIDGetParams{

		timeout: timeout,
	}
}

// NewFAXReceiptsByMessageIDGetParamsWithContext creates a new FAXReceiptsByMessageIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFAXReceiptsByMessageIDGetParamsWithContext(ctx context.Context) *FAXReceiptsByMessageIDGetParams {
	var ()
	return &FAXReceiptsByMessageIDGetParams{

		Context: ctx,
	}
}

// NewFAXReceiptsByMessageIDGetParamsWithHTTPClient creates a new FAXReceiptsByMessageIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFAXReceiptsByMessageIDGetParamsWithHTTPClient(client *http.Client) *FAXReceiptsByMessageIDGetParams {
	var ()
	return &FAXReceiptsByMessageIDGetParams{
		HTTPClient: client,
	}
}

/*FAXReceiptsByMessageIDGetParams contains all the parameters to send to the API endpoint
for the Fax receipts by message Id get operation typically these are written to a http.Request
*/
type FAXReceiptsByMessageIDGetParams struct {

	/*MessageID
	  ID of the message receipt to retrieve

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) WithTimeout(timeout time.Duration) *FAXReceiptsByMessageIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) WithContext(ctx context.Context) *FAXReceiptsByMessageIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) WithHTTPClient(client *http.Client) *FAXReceiptsByMessageIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageID adds the messageID to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) WithMessageID(messageID string) *FAXReceiptsByMessageIDGetParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the Fax receipts by message Id get params
func (o *FAXReceiptsByMessageIDGetParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *FAXReceiptsByMessageIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param message_id
	if err := r.SetPathParam("message_id", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
