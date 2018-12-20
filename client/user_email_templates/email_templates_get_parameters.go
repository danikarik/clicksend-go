// Code generated by go-swagger; DO NOT EDIT.

package user_email_templates

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

// NewEmailTemplatesGetParams creates a new EmailTemplatesGetParams object
// with the default values initialized.
func NewEmailTemplatesGetParams() *EmailTemplatesGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailTemplatesGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailTemplatesGetParamsWithTimeout creates a new EmailTemplatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailTemplatesGetParamsWithTimeout(timeout time.Duration) *EmailTemplatesGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailTemplatesGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewEmailTemplatesGetParamsWithContext creates a new EmailTemplatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailTemplatesGetParamsWithContext(ctx context.Context) *EmailTemplatesGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailTemplatesGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewEmailTemplatesGetParamsWithHTTPClient creates a new EmailTemplatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailTemplatesGetParamsWithHTTPClient(client *http.Client) *EmailTemplatesGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailTemplatesGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*EmailTemplatesGetParams contains all the parameters to send to the API endpoint
for the email templates get operation typically these are written to a http.Request
*/
type EmailTemplatesGetParams struct {

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

// WithTimeout adds the timeout to the email templates get params
func (o *EmailTemplatesGetParams) WithTimeout(timeout time.Duration) *EmailTemplatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email templates get params
func (o *EmailTemplatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email templates get params
func (o *EmailTemplatesGetParams) WithContext(ctx context.Context) *EmailTemplatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email templates get params
func (o *EmailTemplatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email templates get params
func (o *EmailTemplatesGetParams) WithHTTPClient(client *http.Client) *EmailTemplatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email templates get params
func (o *EmailTemplatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the email templates get params
func (o *EmailTemplatesGetParams) WithLimit(limit *int32) *EmailTemplatesGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the email templates get params
func (o *EmailTemplatesGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the email templates get params
func (o *EmailTemplatesGetParams) WithPage(page *int32) *EmailTemplatesGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the email templates get params
func (o *EmailTemplatesGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *EmailTemplatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
