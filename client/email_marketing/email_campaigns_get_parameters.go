// Code generated by go-swagger; DO NOT EDIT.

package email_marketing

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

// NewEmailCampaignsGetParams creates a new EmailCampaignsGetParams object
// with the default values initialized.
func NewEmailCampaignsGetParams() *EmailCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailCampaignsGetParamsWithTimeout creates a new EmailCampaignsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailCampaignsGetParamsWithTimeout(timeout time.Duration) *EmailCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewEmailCampaignsGetParamsWithContext creates a new EmailCampaignsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailCampaignsGetParamsWithContext(ctx context.Context) *EmailCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailCampaignsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewEmailCampaignsGetParamsWithHTTPClient creates a new EmailCampaignsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailCampaignsGetParamsWithHTTPClient(client *http.Client) *EmailCampaignsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &EmailCampaignsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*EmailCampaignsGetParams contains all the parameters to send to the API endpoint
for the email campaigns get operation typically these are written to a http.Request
*/
type EmailCampaignsGetParams struct {

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

// WithTimeout adds the timeout to the email campaigns get params
func (o *EmailCampaignsGetParams) WithTimeout(timeout time.Duration) *EmailCampaignsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email campaigns get params
func (o *EmailCampaignsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email campaigns get params
func (o *EmailCampaignsGetParams) WithContext(ctx context.Context) *EmailCampaignsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email campaigns get params
func (o *EmailCampaignsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email campaigns get params
func (o *EmailCampaignsGetParams) WithHTTPClient(client *http.Client) *EmailCampaignsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email campaigns get params
func (o *EmailCampaignsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the email campaigns get params
func (o *EmailCampaignsGetParams) WithLimit(limit *int32) *EmailCampaignsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the email campaigns get params
func (o *EmailCampaignsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the email campaigns get params
func (o *EmailCampaignsGetParams) WithPage(page *int32) *EmailCampaignsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the email campaigns get params
func (o *EmailCampaignsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *EmailCampaignsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
