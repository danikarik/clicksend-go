// Code generated by go-swagger; DO NOT EDIT.

package post_postcard

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

// NewPostPostcardsHistoryGetParams creates a new PostPostcardsHistoryGetParams object
// with the default values initialized.
func NewPostPostcardsHistoryGetParams() *PostPostcardsHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &PostPostcardsHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPostcardsHistoryGetParamsWithTimeout creates a new PostPostcardsHistoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPostcardsHistoryGetParamsWithTimeout(timeout time.Duration) *PostPostcardsHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &PostPostcardsHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewPostPostcardsHistoryGetParamsWithContext creates a new PostPostcardsHistoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPostcardsHistoryGetParamsWithContext(ctx context.Context) *PostPostcardsHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &PostPostcardsHistoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewPostPostcardsHistoryGetParamsWithHTTPClient creates a new PostPostcardsHistoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPostcardsHistoryGetParamsWithHTTPClient(client *http.Client) *PostPostcardsHistoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &PostPostcardsHistoryGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*PostPostcardsHistoryGetParams contains all the parameters to send to the API endpoint
for the post postcards history get operation typically these are written to a http.Request
*/
type PostPostcardsHistoryGetParams struct {

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

// WithTimeout adds the timeout to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) WithTimeout(timeout time.Duration) *PostPostcardsHistoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) WithContext(ctx context.Context) *PostPostcardsHistoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) WithHTTPClient(client *http.Client) *PostPostcardsHistoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) WithLimit(limit *int32) *PostPostcardsHistoryGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) WithPage(page *int32) *PostPostcardsHistoryGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the post postcards history get params
func (o *PostPostcardsHistoryGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *PostPostcardsHistoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
