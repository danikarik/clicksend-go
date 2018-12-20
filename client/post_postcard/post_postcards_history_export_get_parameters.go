// Code generated by go-swagger; DO NOT EDIT.

package post_postcard

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

// NewPostPostcardsHistoryExportGetParams creates a new PostPostcardsHistoryExportGetParams object
// with the default values initialized.
func NewPostPostcardsHistoryExportGetParams() *PostPostcardsHistoryExportGetParams {
	var ()
	return &PostPostcardsHistoryExportGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPostcardsHistoryExportGetParamsWithTimeout creates a new PostPostcardsHistoryExportGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPostcardsHistoryExportGetParamsWithTimeout(timeout time.Duration) *PostPostcardsHistoryExportGetParams {
	var ()
	return &PostPostcardsHistoryExportGetParams{

		timeout: timeout,
	}
}

// NewPostPostcardsHistoryExportGetParamsWithContext creates a new PostPostcardsHistoryExportGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPostcardsHistoryExportGetParamsWithContext(ctx context.Context) *PostPostcardsHistoryExportGetParams {
	var ()
	return &PostPostcardsHistoryExportGetParams{

		Context: ctx,
	}
}

// NewPostPostcardsHistoryExportGetParamsWithHTTPClient creates a new PostPostcardsHistoryExportGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPostcardsHistoryExportGetParamsWithHTTPClient(client *http.Client) *PostPostcardsHistoryExportGetParams {
	var ()
	return &PostPostcardsHistoryExportGetParams{
		HTTPClient: client,
	}
}

/*PostPostcardsHistoryExportGetParams contains all the parameters to send to the API endpoint
for the post postcards history export get operation typically these are written to a http.Request
*/
type PostPostcardsHistoryExportGetParams struct {

	/*Filename
	  Filename to export to

	*/
	Filename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) WithTimeout(timeout time.Duration) *PostPostcardsHistoryExportGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) WithContext(ctx context.Context) *PostPostcardsHistoryExportGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) WithHTTPClient(client *http.Client) *PostPostcardsHistoryExportGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) WithFilename(filename string) *PostPostcardsHistoryExportGetParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the post postcards history export get params
func (o *PostPostcardsHistoryExportGetParams) SetFilename(filename string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *PostPostcardsHistoryExportGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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