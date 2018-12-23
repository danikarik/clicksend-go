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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/danikarik/clicksend-go/models"
)

// NewEmailCampaignPricePostParams creates a new EmailCampaignPricePostParams object
// with the default values initialized.
func NewEmailCampaignPricePostParams() *EmailCampaignPricePostParams {
	var ()
	return &EmailCampaignPricePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailCampaignPricePostParamsWithTimeout creates a new EmailCampaignPricePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailCampaignPricePostParamsWithTimeout(timeout time.Duration) *EmailCampaignPricePostParams {
	var ()
	return &EmailCampaignPricePostParams{

		timeout: timeout,
	}
}

// NewEmailCampaignPricePostParamsWithContext creates a new EmailCampaignPricePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailCampaignPricePostParamsWithContext(ctx context.Context) *EmailCampaignPricePostParams {
	var ()
	return &EmailCampaignPricePostParams{

		Context: ctx,
	}
}

// NewEmailCampaignPricePostParamsWithHTTPClient creates a new EmailCampaignPricePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailCampaignPricePostParamsWithHTTPClient(client *http.Client) *EmailCampaignPricePostParams {
	var ()
	return &EmailCampaignPricePostParams{
		HTTPClient: client,
	}
}

/*EmailCampaignPricePostParams contains all the parameters to send to the API endpoint
for the email campaign price post operation typically these are written to a http.Request
*/
type EmailCampaignPricePostParams struct {

	/*EmailCampaign
	  Email model

	*/
	EmailCampaign models.EmailCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email campaign price post params
func (o *EmailCampaignPricePostParams) WithTimeout(timeout time.Duration) *EmailCampaignPricePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email campaign price post params
func (o *EmailCampaignPricePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email campaign price post params
func (o *EmailCampaignPricePostParams) WithContext(ctx context.Context) *EmailCampaignPricePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email campaign price post params
func (o *EmailCampaignPricePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email campaign price post params
func (o *EmailCampaignPricePostParams) WithHTTPClient(client *http.Client) *EmailCampaignPricePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email campaign price post params
func (o *EmailCampaignPricePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailCampaign adds the emailCampaign to the email campaign price post params
func (o *EmailCampaignPricePostParams) WithEmailCampaign(emailCampaign models.EmailCampaign) *EmailCampaignPricePostParams {
	o.SetEmailCampaign(emailCampaign)
	return o
}

// SetEmailCampaign adds the emailCampaign to the email campaign price post params
func (o *EmailCampaignPricePostParams) SetEmailCampaign(emailCampaign models.EmailCampaign) {
	o.EmailCampaign = emailCampaign
}

// WriteToRequest writes these params to a swagger request
func (o *EmailCampaignPricePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.EmailCampaign); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
