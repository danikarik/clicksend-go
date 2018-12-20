// Code generated by go-swagger; DO NOT EDIT.

package reseller_account

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

// NewResellerAccountsByClientUserIDGetParams creates a new ResellerAccountsByClientUserIDGetParams object
// with the default values initialized.
func NewResellerAccountsByClientUserIDGetParams() *ResellerAccountsByClientUserIDGetParams {
	var ()
	return &ResellerAccountsByClientUserIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResellerAccountsByClientUserIDGetParamsWithTimeout creates a new ResellerAccountsByClientUserIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResellerAccountsByClientUserIDGetParamsWithTimeout(timeout time.Duration) *ResellerAccountsByClientUserIDGetParams {
	var ()
	return &ResellerAccountsByClientUserIDGetParams{

		timeout: timeout,
	}
}

// NewResellerAccountsByClientUserIDGetParamsWithContext creates a new ResellerAccountsByClientUserIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewResellerAccountsByClientUserIDGetParamsWithContext(ctx context.Context) *ResellerAccountsByClientUserIDGetParams {
	var ()
	return &ResellerAccountsByClientUserIDGetParams{

		Context: ctx,
	}
}

// NewResellerAccountsByClientUserIDGetParamsWithHTTPClient creates a new ResellerAccountsByClientUserIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResellerAccountsByClientUserIDGetParamsWithHTTPClient(client *http.Client) *ResellerAccountsByClientUserIDGetParams {
	var ()
	return &ResellerAccountsByClientUserIDGetParams{
		HTTPClient: client,
	}
}

/*ResellerAccountsByClientUserIDGetParams contains all the parameters to send to the API endpoint
for the reseller accounts by client user Id get operation typically these are written to a http.Request
*/
type ResellerAccountsByClientUserIDGetParams struct {

	/*ClientUserID
	  User ID of client

	*/
	ClientUserID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) WithTimeout(timeout time.Duration) *ResellerAccountsByClientUserIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) WithContext(ctx context.Context) *ResellerAccountsByClientUserIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) WithHTTPClient(client *http.Client) *ResellerAccountsByClientUserIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientUserID adds the clientUserID to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) WithClientUserID(clientUserID int32) *ResellerAccountsByClientUserIDGetParams {
	o.SetClientUserID(clientUserID)
	return o
}

// SetClientUserID adds the clientUserId to the reseller accounts by client user Id get params
func (o *ResellerAccountsByClientUserIDGetParams) SetClientUserID(clientUserID int32) {
	o.ClientUserID = clientUserID
}

// WriteToRequest writes these params to a swagger request
func (o *ResellerAccountsByClientUserIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_user_id
	if err := r.SetPathParam("client_user_id", swag.FormatInt32(o.ClientUserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
