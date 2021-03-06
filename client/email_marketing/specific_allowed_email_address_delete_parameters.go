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

// NewSpecificAllowedEmailAddressDeleteParams creates a new SpecificAllowedEmailAddressDeleteParams object
// with the default values initialized.
func NewSpecificAllowedEmailAddressDeleteParams() *SpecificAllowedEmailAddressDeleteParams {
	var ()
	return &SpecificAllowedEmailAddressDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpecificAllowedEmailAddressDeleteParamsWithTimeout creates a new SpecificAllowedEmailAddressDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpecificAllowedEmailAddressDeleteParamsWithTimeout(timeout time.Duration) *SpecificAllowedEmailAddressDeleteParams {
	var ()
	return &SpecificAllowedEmailAddressDeleteParams{

		timeout: timeout,
	}
}

// NewSpecificAllowedEmailAddressDeleteParamsWithContext creates a new SpecificAllowedEmailAddressDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpecificAllowedEmailAddressDeleteParamsWithContext(ctx context.Context) *SpecificAllowedEmailAddressDeleteParams {
	var ()
	return &SpecificAllowedEmailAddressDeleteParams{

		Context: ctx,
	}
}

// NewSpecificAllowedEmailAddressDeleteParamsWithHTTPClient creates a new SpecificAllowedEmailAddressDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpecificAllowedEmailAddressDeleteParamsWithHTTPClient(client *http.Client) *SpecificAllowedEmailAddressDeleteParams {
	var ()
	return &SpecificAllowedEmailAddressDeleteParams{
		HTTPClient: client,
	}
}

/*SpecificAllowedEmailAddressDeleteParams contains all the parameters to send to the API endpoint
for the specific allowed email address delete operation typically these are written to a http.Request
*/
type SpecificAllowedEmailAddressDeleteParams struct {

	/*EmailAddressID
	  Allowed email address id

	*/
	EmailAddressID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) WithTimeout(timeout time.Duration) *SpecificAllowedEmailAddressDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) WithContext(ctx context.Context) *SpecificAllowedEmailAddressDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) WithHTTPClient(client *http.Client) *SpecificAllowedEmailAddressDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailAddressID adds the emailAddressID to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) WithEmailAddressID(emailAddressID int32) *SpecificAllowedEmailAddressDeleteParams {
	o.SetEmailAddressID(emailAddressID)
	return o
}

// SetEmailAddressID adds the emailAddressId to the specific allowed email address delete params
func (o *SpecificAllowedEmailAddressDeleteParams) SetEmailAddressID(emailAddressID int32) {
	o.EmailAddressID = emailAddressID
}

// WriteToRequest writes these params to a swagger request
func (o *SpecificAllowedEmailAddressDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email_address_id
	if err := r.SetPathParam("email_address_id", swag.FormatInt32(o.EmailAddressID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
