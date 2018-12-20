// Code generated by go-swagger; DO NOT EDIT.

package email_to_sms

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewSMSEmailSMSStrippedStringPutParams creates a new SMSEmailSMSStrippedStringPutParams object
// with the default values initialized.
func NewSMSEmailSMSStrippedStringPutParams() *SMSEmailSMSStrippedStringPutParams {
	var ()
	return &SMSEmailSMSStrippedStringPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSMSEmailSMSStrippedStringPutParamsWithTimeout creates a new SMSEmailSMSStrippedStringPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSMSEmailSMSStrippedStringPutParamsWithTimeout(timeout time.Duration) *SMSEmailSMSStrippedStringPutParams {
	var ()
	return &SMSEmailSMSStrippedStringPutParams{

		timeout: timeout,
	}
}

// NewSMSEmailSMSStrippedStringPutParamsWithContext creates a new SMSEmailSMSStrippedStringPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewSMSEmailSMSStrippedStringPutParamsWithContext(ctx context.Context) *SMSEmailSMSStrippedStringPutParams {
	var ()
	return &SMSEmailSMSStrippedStringPutParams{

		Context: ctx,
	}
}

// NewSMSEmailSMSStrippedStringPutParamsWithHTTPClient creates a new SMSEmailSMSStrippedStringPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSMSEmailSMSStrippedStringPutParamsWithHTTPClient(client *http.Client) *SMSEmailSMSStrippedStringPutParams {
	var ()
	return &SMSEmailSMSStrippedStringPutParams{
		HTTPClient: client,
	}
}

/*SMSEmailSMSStrippedStringPutParams contains all the parameters to send to the API endpoint
for the Sms email Sms stripped string put operation typically these are written to a http.Request
*/
type SMSEmailSMSStrippedStringPutParams struct {

	/*RuleID
	  Your rule id

	*/
	RuleID int32
	/*StrippedString
	  StrippedString model

	*/
	StrippedString models.StrippedString

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) WithTimeout(timeout time.Duration) *SMSEmailSMSStrippedStringPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) WithContext(ctx context.Context) *SMSEmailSMSStrippedStringPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) WithHTTPClient(client *http.Client) *SMSEmailSMSStrippedStringPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) WithRuleID(ruleID int32) *SMSEmailSMSStrippedStringPutParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) SetRuleID(ruleID int32) {
	o.RuleID = ruleID
}

// WithStrippedString adds the strippedString to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) WithStrippedString(strippedString models.StrippedString) *SMSEmailSMSStrippedStringPutParams {
	o.SetStrippedString(strippedString)
	return o
}

// SetStrippedString adds the strippedString to the Sms email Sms stripped string put params
func (o *SMSEmailSMSStrippedStringPutParams) SetStrippedString(strippedString models.StrippedString) {
	o.StrippedString = strippedString
}

// WriteToRequest writes these params to a swagger request
func (o *SMSEmailSMSStrippedStringPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rule_id
	if err := r.SetPathParam("rule_id", swag.FormatInt32(o.RuleID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.StrippedString); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}