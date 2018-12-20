// Code generated by go-swagger; DO NOT EDIT.

package contact_list

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

// NewListsByListIDDeleteParams creates a new ListsByListIDDeleteParams object
// with the default values initialized.
func NewListsByListIDDeleteParams() *ListsByListIDDeleteParams {
	var ()
	return &ListsByListIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListsByListIDDeleteParamsWithTimeout creates a new ListsByListIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListsByListIDDeleteParamsWithTimeout(timeout time.Duration) *ListsByListIDDeleteParams {
	var ()
	return &ListsByListIDDeleteParams{

		timeout: timeout,
	}
}

// NewListsByListIDDeleteParamsWithContext creates a new ListsByListIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewListsByListIDDeleteParamsWithContext(ctx context.Context) *ListsByListIDDeleteParams {
	var ()
	return &ListsByListIDDeleteParams{

		Context: ctx,
	}
}

// NewListsByListIDDeleteParamsWithHTTPClient creates a new ListsByListIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListsByListIDDeleteParamsWithHTTPClient(client *http.Client) *ListsByListIDDeleteParams {
	var ()
	return &ListsByListIDDeleteParams{
		HTTPClient: client,
	}
}

/*ListsByListIDDeleteParams contains all the parameters to send to the API endpoint
for the lists by list Id delete operation typically these are written to a http.Request
*/
type ListsByListIDDeleteParams struct {

	/*ListID
	  List ID

	*/
	ListID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) WithTimeout(timeout time.Duration) *ListsByListIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) WithContext(ctx context.Context) *ListsByListIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) WithHTTPClient(client *http.Client) *ListsByListIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) WithListID(listID int32) *ListsByListIDDeleteParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists by list Id delete params
func (o *ListsByListIDDeleteParams) SetListID(listID int32) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *ListsByListIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param list_id
	if err := r.SetPathParam("list_id", swag.FormatInt32(o.ListID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
