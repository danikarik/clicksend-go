// Code generated by go-swagger; DO NOT EDIT.

package voice

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

// NewVoiceCancelByMessageIDPutParams creates a new VoiceCancelByMessageIDPutParams object
// with the default values initialized.
func NewVoiceCancelByMessageIDPutParams() *VoiceCancelByMessageIDPutParams {
	var ()
	return &VoiceCancelByMessageIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoiceCancelByMessageIDPutParamsWithTimeout creates a new VoiceCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoiceCancelByMessageIDPutParamsWithTimeout(timeout time.Duration) *VoiceCancelByMessageIDPutParams {
	var ()
	return &VoiceCancelByMessageIDPutParams{

		timeout: timeout,
	}
}

// NewVoiceCancelByMessageIDPutParamsWithContext creates a new VoiceCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoiceCancelByMessageIDPutParamsWithContext(ctx context.Context) *VoiceCancelByMessageIDPutParams {
	var ()
	return &VoiceCancelByMessageIDPutParams{

		Context: ctx,
	}
}

// NewVoiceCancelByMessageIDPutParamsWithHTTPClient creates a new VoiceCancelByMessageIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoiceCancelByMessageIDPutParamsWithHTTPClient(client *http.Client) *VoiceCancelByMessageIDPutParams {
	var ()
	return &VoiceCancelByMessageIDPutParams{
		HTTPClient: client,
	}
}

/*VoiceCancelByMessageIDPutParams contains all the parameters to send to the API endpoint
for the voice cancel by message Id put operation typically these are written to a http.Request
*/
type VoiceCancelByMessageIDPutParams struct {

	/*MessageID
	  Your voice message id

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) WithTimeout(timeout time.Duration) *VoiceCancelByMessageIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) WithContext(ctx context.Context) *VoiceCancelByMessageIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) WithHTTPClient(client *http.Client) *VoiceCancelByMessageIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageID adds the messageID to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) WithMessageID(messageID string) *VoiceCancelByMessageIDPutParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the voice cancel by message Id put params
func (o *VoiceCancelByMessageIDPutParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *VoiceCancelByMessageIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
