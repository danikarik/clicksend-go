// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewAccountVerifySendPutParams creates a new AccountVerifySendPutParams object
// with the default values initialized.
func NewAccountVerifySendPutParams() *AccountVerifySendPutParams {
	var ()
	return &AccountVerifySendPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAccountVerifySendPutParamsWithTimeout creates a new AccountVerifySendPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountVerifySendPutParamsWithTimeout(timeout time.Duration) *AccountVerifySendPutParams {
	var ()
	return &AccountVerifySendPutParams{

		timeout: timeout,
	}
}

// NewAccountVerifySendPutParamsWithContext creates a new AccountVerifySendPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountVerifySendPutParamsWithContext(ctx context.Context) *AccountVerifySendPutParams {
	var ()
	return &AccountVerifySendPutParams{

		Context: ctx,
	}
}

// NewAccountVerifySendPutParamsWithHTTPClient creates a new AccountVerifySendPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAccountVerifySendPutParamsWithHTTPClient(client *http.Client) *AccountVerifySendPutParams {
	var ()
	return &AccountVerifySendPutParams{
		HTTPClient: client,
	}
}

/*AccountVerifySendPutParams contains all the parameters to send to the API endpoint
for the account verify send put operation typically these are written to a http.Request
*/
type AccountVerifySendPutParams struct {

	/*AccountVerify
	  Account details

	*/
	AccountVerify models.AccountVerify

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the account verify send put params
func (o *AccountVerifySendPutParams) WithTimeout(timeout time.Duration) *AccountVerifySendPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account verify send put params
func (o *AccountVerifySendPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account verify send put params
func (o *AccountVerifySendPutParams) WithContext(ctx context.Context) *AccountVerifySendPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account verify send put params
func (o *AccountVerifySendPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account verify send put params
func (o *AccountVerifySendPutParams) WithHTTPClient(client *http.Client) *AccountVerifySendPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account verify send put params
func (o *AccountVerifySendPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountVerify adds the accountVerify to the account verify send put params
func (o *AccountVerifySendPutParams) WithAccountVerify(accountVerify models.AccountVerify) *AccountVerifySendPutParams {
	o.SetAccountVerify(accountVerify)
	return o
}

// SetAccountVerify adds the accountVerify to the account verify send put params
func (o *AccountVerifySendPutParams) SetAccountVerify(accountVerify models.AccountVerify) {
	o.AccountVerify = accountVerify
}

// WriteToRequest writes these params to a swagger request
func (o *AccountVerifySendPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.AccountVerify); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}