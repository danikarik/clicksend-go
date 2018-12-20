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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/danikarik/clicksend-go/models"
)

// NewEmailTemplatePostParams creates a new EmailTemplatePostParams object
// with the default values initialized.
func NewEmailTemplatePostParams() *EmailTemplatePostParams {
	var ()
	return &EmailTemplatePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailTemplatePostParamsWithTimeout creates a new EmailTemplatePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailTemplatePostParamsWithTimeout(timeout time.Duration) *EmailTemplatePostParams {
	var ()
	return &EmailTemplatePostParams{

		timeout: timeout,
	}
}

// NewEmailTemplatePostParamsWithContext creates a new EmailTemplatePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailTemplatePostParamsWithContext(ctx context.Context) *EmailTemplatePostParams {
	var ()
	return &EmailTemplatePostParams{

		Context: ctx,
	}
}

// NewEmailTemplatePostParamsWithHTTPClient creates a new EmailTemplatePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailTemplatePostParamsWithHTTPClient(client *http.Client) *EmailTemplatePostParams {
	var ()
	return &EmailTemplatePostParams{
		HTTPClient: client,
	}
}

/*EmailTemplatePostParams contains all the parameters to send to the API endpoint
for the email template post operation typically these are written to a http.Request
*/
type EmailTemplatePostParams struct {

	/*EmailTemplate
	  Email template model

	*/
	EmailTemplate models.EmailTemplateNew

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email template post params
func (o *EmailTemplatePostParams) WithTimeout(timeout time.Duration) *EmailTemplatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email template post params
func (o *EmailTemplatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email template post params
func (o *EmailTemplatePostParams) WithContext(ctx context.Context) *EmailTemplatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email template post params
func (o *EmailTemplatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email template post params
func (o *EmailTemplatePostParams) WithHTTPClient(client *http.Client) *EmailTemplatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email template post params
func (o *EmailTemplatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailTemplate adds the emailTemplate to the email template post params
func (o *EmailTemplatePostParams) WithEmailTemplate(emailTemplate models.EmailTemplateNew) *EmailTemplatePostParams {
	o.SetEmailTemplate(emailTemplate)
	return o
}

// SetEmailTemplate adds the emailTemplate to the email template post params
func (o *EmailTemplatePostParams) SetEmailTemplate(emailTemplate models.EmailTemplateNew) {
	o.EmailTemplate = emailTemplate
}

// WriteToRequest writes these params to a swagger request
func (o *EmailTemplatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.EmailTemplate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
