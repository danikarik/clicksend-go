// Code generated by go-swagger; DO NOT EDIT.

package user_email_templates

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

// NewEmailTemplateGetParams creates a new EmailTemplateGetParams object
// with the default values initialized.
func NewEmailTemplateGetParams() *EmailTemplateGetParams {
	var ()
	return &EmailTemplateGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailTemplateGetParamsWithTimeout creates a new EmailTemplateGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailTemplateGetParamsWithTimeout(timeout time.Duration) *EmailTemplateGetParams {
	var ()
	return &EmailTemplateGetParams{

		timeout: timeout,
	}
}

// NewEmailTemplateGetParamsWithContext creates a new EmailTemplateGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailTemplateGetParamsWithContext(ctx context.Context) *EmailTemplateGetParams {
	var ()
	return &EmailTemplateGetParams{

		Context: ctx,
	}
}

// NewEmailTemplateGetParamsWithHTTPClient creates a new EmailTemplateGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailTemplateGetParamsWithHTTPClient(client *http.Client) *EmailTemplateGetParams {
	var ()
	return &EmailTemplateGetParams{
		HTTPClient: client,
	}
}

/*EmailTemplateGetParams contains all the parameters to send to the API endpoint
for the email template get operation typically these are written to a http.Request
*/
type EmailTemplateGetParams struct {

	/*TemplateID
	  Email template id

	*/
	TemplateID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email template get params
func (o *EmailTemplateGetParams) WithTimeout(timeout time.Duration) *EmailTemplateGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email template get params
func (o *EmailTemplateGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email template get params
func (o *EmailTemplateGetParams) WithContext(ctx context.Context) *EmailTemplateGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email template get params
func (o *EmailTemplateGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email template get params
func (o *EmailTemplateGetParams) WithHTTPClient(client *http.Client) *EmailTemplateGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email template get params
func (o *EmailTemplateGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateID adds the templateID to the email template get params
func (o *EmailTemplateGetParams) WithTemplateID(templateID int32) *EmailTemplateGetParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the email template get params
func (o *EmailTemplateGetParams) SetTemplateID(templateID int32) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *EmailTemplateGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param template_id
	if err := r.SetPathParam("template_id", swag.FormatInt32(o.TemplateID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
