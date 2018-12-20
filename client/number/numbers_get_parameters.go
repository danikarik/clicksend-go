// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNumbersGetParams creates a new NumbersGetParams object
// with the default values initialized.
func NewNumbersGetParams() *NumbersGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &NumbersGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNumbersGetParamsWithTimeout creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNumbersGetParamsWithTimeout(timeout time.Duration) *NumbersGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &NumbersGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewNumbersGetParamsWithContext creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNumbersGetParamsWithContext(ctx context.Context) *NumbersGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &NumbersGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewNumbersGetParamsWithHTTPClient creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNumbersGetParamsWithHTTPClient(client *http.Client) *NumbersGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &NumbersGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*NumbersGetParams contains all the parameters to send to the API endpoint
for the numbers get operation typically these are written to a http.Request
*/
type NumbersGetParams struct {

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

// WithTimeout adds the timeout to the numbers get params
func (o *NumbersGetParams) WithTimeout(timeout time.Duration) *NumbersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the numbers get params
func (o *NumbersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the numbers get params
func (o *NumbersGetParams) WithContext(ctx context.Context) *NumbersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the numbers get params
func (o *NumbersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the numbers get params
func (o *NumbersGetParams) WithHTTPClient(client *http.Client) *NumbersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the numbers get params
func (o *NumbersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the numbers get params
func (o *NumbersGetParams) WithLimit(limit *int32) *NumbersGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the numbers get params
func (o *NumbersGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the numbers get params
func (o *NumbersGetParams) WithPage(page *int32) *NumbersGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the numbers get params
func (o *NumbersGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *NumbersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
