// Code generated by go-swagger; DO NOT EDIT.

package upload

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

// NewUploadsPostParams creates a new UploadsPostParams object
// with the default values initialized.
func NewUploadsPostParams() *UploadsPostParams {
	var ()
	return &UploadsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadsPostParamsWithTimeout creates a new UploadsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadsPostParamsWithTimeout(timeout time.Duration) *UploadsPostParams {
	var ()
	return &UploadsPostParams{

		timeout: timeout,
	}
}

// NewUploadsPostParamsWithContext creates a new UploadsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadsPostParamsWithContext(ctx context.Context) *UploadsPostParams {
	var ()
	return &UploadsPostParams{

		Context: ctx,
	}
}

// NewUploadsPostParamsWithHTTPClient creates a new UploadsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadsPostParamsWithHTTPClient(client *http.Client) *UploadsPostParams {
	var ()
	return &UploadsPostParams{
		HTTPClient: client,
	}
}

/*UploadsPostParams contains all the parameters to send to the API endpoint
for the uploads post operation typically these are written to a http.Request
*/
type UploadsPostParams struct {

	/*Convert*/
	Convert string
	/*UploadFile
	  Your file to be uploaded

	*/
	UploadFile models.UploadFile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the uploads post params
func (o *UploadsPostParams) WithTimeout(timeout time.Duration) *UploadsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uploads post params
func (o *UploadsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uploads post params
func (o *UploadsPostParams) WithContext(ctx context.Context) *UploadsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uploads post params
func (o *UploadsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uploads post params
func (o *UploadsPostParams) WithHTTPClient(client *http.Client) *UploadsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uploads post params
func (o *UploadsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConvert adds the convert to the uploads post params
func (o *UploadsPostParams) WithConvert(convert string) *UploadsPostParams {
	o.SetConvert(convert)
	return o
}

// SetConvert adds the convert to the uploads post params
func (o *UploadsPostParams) SetConvert(convert string) {
	o.Convert = convert
}

// WithUploadFile adds the uploadFile to the uploads post params
func (o *UploadsPostParams) WithUploadFile(uploadFile models.UploadFile) *UploadsPostParams {
	o.SetUploadFile(uploadFile)
	return o
}

// SetUploadFile adds the uploadFile to the uploads post params
func (o *UploadsPostParams) SetUploadFile(uploadFile models.UploadFile) {
	o.UploadFile = uploadFile
}

// WriteToRequest writes these params to a swagger request
func (o *UploadsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param convert
	qrConvert := o.Convert
	qConvert := qrConvert
	if qConvert != "" {
		if err := r.SetQueryParam("convert", qConvert); err != nil {
			return err
		}
	}

	if err := r.SetBodyParam(o.UploadFile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
