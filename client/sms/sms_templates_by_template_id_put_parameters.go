// Code generated by go-swagger; DO NOT EDIT.

package sms

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewSMSTemplatesByTemplateIDPutParams creates a new SMSTemplatesByTemplateIDPutParams object
// with the default values initialized.
func NewSMSTemplatesByTemplateIDPutParams() *SMSTemplatesByTemplateIDPutParams {
	var ()
	return &SMSTemplatesByTemplateIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSTemplatesByTemplateIDPutParamsWithTimeout creates a new SMSTemplatesByTemplateIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSTemplatesByTemplateIDPutParamsWithTimeout(timeout time.Duration) *SMSTemplatesByTemplateIDPutParams {
	var ()
	return &SMSTemplatesByTemplateIDPutParams{

		timeout: timeout,
	}
}

// NewSMSTemplatesByTemplateIDPutParamsWithContext creates a new SMSTemplatesByTemplateIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSTemplatesByTemplateIDPutParamsWithContext(ctx context.Context) *SMSTemplatesByTemplateIDPutParams {
	var ()
	return &SMSTemplatesByTemplateIDPutParams{

		Context: ctx,
	}
}

// NewSMSTemplatesByTemplateIDPutParamsWithHTTPClient creates a new SMSTemplatesByTemplateIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSTemplatesByTemplateIDPutParamsWithHTTPClient(client *http.Client) *SMSTemplatesByTemplateIDPutParams {
	var ()
	return &SMSTemplatesByTemplateIDPutParams{
		HTTPClient: client,
	}
}

/*SMSTemplatesByTemplateIDPutParams contains all the parameters to send to the API endpoint
for the Sms templates by template Id put operation typically these are written to a http.Request
*/
type SMSTemplatesByTemplateIDPutParams struct {

	/*SMSTemplate
	  Template item

	*/
	SMSTemplate models.SMSTemplate
	/*TemplateID
	  Template id

	*/
	TemplateID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) WithTimeout(timeout time.Duration) *SMSTemplatesByTemplateIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) WithContext(ctx context.Context) *SMSTemplatesByTemplateIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) WithHTTPClient(client *http.Client) *SMSTemplatesByTemplateIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMSTemplate adds the sMSTemplate to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) WithSMSTemplate(sMSTemplate models.SMSTemplate) *SMSTemplatesByTemplateIDPutParams {
	o.SetSMSTemplate(sMSTemplate)
	return o
}

// SetSMSTemplate adds the smsTemplate to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) SetSMSTemplate(sMSTemplate models.SMSTemplate) {
	o.SMSTemplate = sMSTemplate
}

// WithTemplateID adds the templateID to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) WithTemplateID(templateID int32) *SMSTemplatesByTemplateIDPutParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the Sms templates by template Id put params
func (o *SMSTemplatesByTemplateIDPutParams) SetTemplateID(templateID int32) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *SMSTemplatesByTemplateIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.SMSTemplate); err != nil {
		return err
	}

	// path param template_id
	if err := r.SetPathParam("template_id", swag.FormatInt32(o.TemplateID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
