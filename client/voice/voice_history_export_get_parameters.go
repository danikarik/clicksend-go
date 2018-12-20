// Code generated by go-swagger; DO NOT EDIT.

package voice

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

// NewVoiceHistoryExportGetParams creates a new VoiceHistoryExportGetParams object
// with the default values initialized.
func NewVoiceHistoryExportGetParams() *VoiceHistoryExportGetParams {
	var ()
	return &VoiceHistoryExportGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoiceHistoryExportGetParamsWithTimeout creates a new VoiceHistoryExportGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoiceHistoryExportGetParamsWithTimeout(timeout time.Duration) *VoiceHistoryExportGetParams {
	var ()
	return &VoiceHistoryExportGetParams{

		timeout: timeout,
	}
}

// NewVoiceHistoryExportGetParamsWithContext creates a new VoiceHistoryExportGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoiceHistoryExportGetParamsWithContext(ctx context.Context) *VoiceHistoryExportGetParams {
	var ()
	return &VoiceHistoryExportGetParams{

		Context: ctx,
	}
}

// NewVoiceHistoryExportGetParamsWithHTTPClient creates a new VoiceHistoryExportGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoiceHistoryExportGetParamsWithHTTPClient(client *http.Client) *VoiceHistoryExportGetParams {
	var ()
	return &VoiceHistoryExportGetParams{
		HTTPClient: client,
	}
}

/*VoiceHistoryExportGetParams contains all the parameters to send to the API endpoint
for the voice history export get operation typically these are written to a http.Request
*/
type VoiceHistoryExportGetParams struct {

	/*Filename
	  Filename to export to

	*/
	Filename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the voice history export get params
func (o *VoiceHistoryExportGetParams) WithTimeout(timeout time.Duration) *VoiceHistoryExportGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the voice history export get params
func (o *VoiceHistoryExportGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the voice history export get params
func (o *VoiceHistoryExportGetParams) WithContext(ctx context.Context) *VoiceHistoryExportGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the voice history export get params
func (o *VoiceHistoryExportGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the voice history export get params
func (o *VoiceHistoryExportGetParams) WithHTTPClient(client *http.Client) *VoiceHistoryExportGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the voice history export get params
func (o *VoiceHistoryExportGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the voice history export get params
func (o *VoiceHistoryExportGetParams) WithFilename(filename string) *VoiceHistoryExportGetParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the voice history export get params
func (o *VoiceHistoryExportGetParams) SetFilename(filename string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *VoiceHistoryExportGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filename
	qrFilename := o.Filename
	qFilename := qrFilename
	if qFilename != "" {
		if err := r.SetQueryParam("filename", qFilename); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
