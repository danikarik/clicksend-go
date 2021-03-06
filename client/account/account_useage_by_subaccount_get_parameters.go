// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewAccountUseageBySubaccountGetParams creates a new AccountUseageBySubaccountGetParams object
// with the default values initialized.
func NewAccountUseageBySubaccountGetParams() *AccountUseageBySubaccountGetParams {
	var ()
	return &AccountUseageBySubaccountGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountUseageBySubaccountGetParamsWithTimeout creates a new AccountUseageBySubaccountGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountUseageBySubaccountGetParamsWithTimeout(timeout time.Duration) *AccountUseageBySubaccountGetParams {
	var ()
	return &AccountUseageBySubaccountGetParams{

		timeout: timeout,
	}
}

// NewAccountUseageBySubaccountGetParamsWithContext creates a new AccountUseageBySubaccountGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountUseageBySubaccountGetParamsWithContext(ctx context.Context) *AccountUseageBySubaccountGetParams {
	var ()
	return &AccountUseageBySubaccountGetParams{

		Context: ctx,
	}
}

// NewAccountUseageBySubaccountGetParamsWithHTTPClient creates a new AccountUseageBySubaccountGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountUseageBySubaccountGetParamsWithHTTPClient(client *http.Client) *AccountUseageBySubaccountGetParams {
	var ()
	return &AccountUseageBySubaccountGetParams{
		HTTPClient: client,
	}
}

/*AccountUseageBySubaccountGetParams contains all the parameters to send to the API endpoint
for the account useage by subaccount get operation typically these are written to a http.Request
*/
type AccountUseageBySubaccountGetParams struct {

	/*Month
	  Month to filter by (mm)

	*/
	Month int32
	/*Year
	  Year to filter by (yyyy)

	*/
	Year int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) WithTimeout(timeout time.Duration) *AccountUseageBySubaccountGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) WithContext(ctx context.Context) *AccountUseageBySubaccountGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) WithHTTPClient(client *http.Client) *AccountUseageBySubaccountGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonth adds the month to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) WithMonth(month int32) *AccountUseageBySubaccountGetParams {
	o.SetMonth(month)
	return o
}

// SetMonth adds the month to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) SetMonth(month int32) {
	o.Month = month
}

// WithYear adds the year to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) WithYear(year int32) *AccountUseageBySubaccountGetParams {
	o.SetYear(year)
	return o
}

// SetYear adds the year to the account useage by subaccount get params
func (o *AccountUseageBySubaccountGetParams) SetYear(year int32) {
	o.Year = year
}

// WriteToRequest writes these params to a swagger request
func (o *AccountUseageBySubaccountGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param month
	if err := r.SetPathParam("month", swag.FormatInt32(o.Month)); err != nil {
		return err
	}

	// path param year
	if err := r.SetPathParam("year", swag.FormatInt32(o.Year)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
