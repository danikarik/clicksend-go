// Code generated by go-swagger; DO NOT EDIT.

package account_recharge

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

// NewRechargeTransactionsGetParams creates a new RechargeTransactionsGetParams object
// with the default values initialized.
func NewRechargeTransactionsGetParams() *RechargeTransactionsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &RechargeTransactionsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRechargeTransactionsGetParamsWithTimeout creates a new RechargeTransactionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRechargeTransactionsGetParamsWithTimeout(timeout time.Duration) *RechargeTransactionsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &RechargeTransactionsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewRechargeTransactionsGetParamsWithContext creates a new RechargeTransactionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRechargeTransactionsGetParamsWithContext(ctx context.Context) *RechargeTransactionsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &RechargeTransactionsGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewRechargeTransactionsGetParamsWithHTTPClient creates a new RechargeTransactionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRechargeTransactionsGetParamsWithHTTPClient(client *http.Client) *RechargeTransactionsGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &RechargeTransactionsGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*RechargeTransactionsGetParams contains all the parameters to send to the API endpoint
for the recharge transactions get operation typically these are written to a http.Request
*/
type RechargeTransactionsGetParams struct {

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

// WithTimeout adds the timeout to the recharge transactions get params
func (o *RechargeTransactionsGetParams) WithTimeout(timeout time.Duration) *RechargeTransactionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recharge transactions get params
func (o *RechargeTransactionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recharge transactions get params
func (o *RechargeTransactionsGetParams) WithContext(ctx context.Context) *RechargeTransactionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recharge transactions get params
func (o *RechargeTransactionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recharge transactions get params
func (o *RechargeTransactionsGetParams) WithHTTPClient(client *http.Client) *RechargeTransactionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recharge transactions get params
func (o *RechargeTransactionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the recharge transactions get params
func (o *RechargeTransactionsGetParams) WithLimit(limit *int32) *RechargeTransactionsGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the recharge transactions get params
func (o *RechargeTransactionsGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the recharge transactions get params
func (o *RechargeTransactionsGetParams) WithPage(page *int32) *RechargeTransactionsGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the recharge transactions get params
func (o *RechargeTransactionsGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *RechargeTransactionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
