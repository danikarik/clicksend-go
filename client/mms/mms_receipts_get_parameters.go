// Code generated by go-swagger; DO NOT EDIT.

package mms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMMSReceiptsGetParams creates a new MMSReceiptsGetParams object
// with the default values initialized.
func NewMMSReceiptsGetParams() *MMSReceiptsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MMSReceiptsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMMSReceiptsGetParamsWithTimeout creates a new MMSReceiptsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMMSReceiptsGetParamsWithTimeout(timeout time.Duration) *MMSReceiptsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MMSReceiptsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewMMSReceiptsGetParamsWithContext creates a new MMSReceiptsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMMSReceiptsGetParamsWithContext(ctx context.Context) *MMSReceiptsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MMSReceiptsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewMMSReceiptsGetParamsWithHTTPClient creates a new MMSReceiptsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMMSReceiptsGetParamsWithHTTPClient(client *http.Client) *MMSReceiptsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MMSReceiptsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*MMSReceiptsGetParams contains all the parameters to send to the API endpoint
for the Mms receipts get operation typically these are written to a http.Request
*/
type MMSReceiptsGetParams struct {

	/*Limit
	  Number of records per page

	*/
	Limit *int32
	/*Page
	  Page number

	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Mms receipts get params
func (o *MMSReceiptsGetParams) WithTimeout(timeout time.Duration) *MMSReceiptsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Mms receipts get params
func (o *MMSReceiptsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Mms receipts get params
func (o *MMSReceiptsGetParams) WithContext(ctx context.Context) *MMSReceiptsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Mms receipts get params
func (o *MMSReceiptsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Mms receipts get params
func (o *MMSReceiptsGetParams) WithHTTPClient(client *http.Client) *MMSReceiptsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Mms receipts get params
func (o *MMSReceiptsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the Mms receipts get params
func (o *MMSReceiptsGetParams) WithLimit(limit *int32) *MMSReceiptsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the Mms receipts get params
func (o *MMSReceiptsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the Mms receipts get params
func (o *MMSReceiptsGetParams) WithPage(page *int32) *MMSReceiptsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the Mms receipts get params
func (o *MMSReceiptsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *MMSReceiptsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
