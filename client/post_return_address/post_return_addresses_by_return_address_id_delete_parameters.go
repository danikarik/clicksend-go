// Code generated by go-swagger; DO NOT EDIT.

package post_return_address

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

// NewPostReturnAddressesByReturnAddressIDDeleteParams creates a new PostReturnAddressesByReturnAddressIDDeleteParams object
// with the default values initialized.
func NewPostReturnAddressesByReturnAddressIDDeleteParams() *PostReturnAddressesByReturnAddressIDDeleteParams {
	var ()
	return &PostReturnAddressesByReturnAddressIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostReturnAddressesByReturnAddressIDDeleteParamsWithTimeout creates a new PostReturnAddressesByReturnAddressIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostReturnAddressesByReturnAddressIDDeleteParamsWithTimeout(timeout time.Duration) *PostReturnAddressesByReturnAddressIDDeleteParams {
	var ()
	return &PostReturnAddressesByReturnAddressIDDeleteParams{

		timeout: timeout,
	}
}

// NewPostReturnAddressesByReturnAddressIDDeleteParamsWithContext creates a new PostReturnAddressesByReturnAddressIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostReturnAddressesByReturnAddressIDDeleteParamsWithContext(ctx context.Context) *PostReturnAddressesByReturnAddressIDDeleteParams {
	var ()
	return &PostReturnAddressesByReturnAddressIDDeleteParams{

		Context: ctx,
	}
}

// NewPostReturnAddressesByReturnAddressIDDeleteParamsWithHTTPClient creates a new PostReturnAddressesByReturnAddressIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostReturnAddressesByReturnAddressIDDeleteParamsWithHTTPClient(client *http.Client) *PostReturnAddressesByReturnAddressIDDeleteParams {
	var ()
	return &PostReturnAddressesByReturnAddressIDDeleteParams{
		HTTPClient: client,
	}
}

/*PostReturnAddressesByReturnAddressIDDeleteParams contains all the parameters to send to the API endpoint
for the post return addresses by return address Id delete operation typically these are written to a http.Request
*/
type PostReturnAddressesByReturnAddressIDDeleteParams struct {

	/*ReturnAddressID
	  Return address ID

	*/
	ReturnAddressID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) WithTimeout(timeout time.Duration) *PostReturnAddressesByReturnAddressIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) WithContext(ctx context.Context) *PostReturnAddressesByReturnAddressIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) WithHTTPClient(client *http.Client) *PostReturnAddressesByReturnAddressIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnAddressID adds the returnAddressID to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) WithReturnAddressID(returnAddressID int32) *PostReturnAddressesByReturnAddressIDDeleteParams {
	o.SetReturnAddressID(returnAddressID)
	return o
}

// SetReturnAddressID adds the returnAddressId to the post return addresses by return address Id delete params
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) SetReturnAddressID(returnAddressID int32) {
	o.ReturnAddressID = returnAddressID
}

// WriteToRequest writes these params to a swagger request
func (o *PostReturnAddressesByReturnAddressIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param return_address_id
	if err := r.SetPathParam("return_address_id", swag.FormatInt32(o.ReturnAddressID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
