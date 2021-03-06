// Code generated by go-swagger; DO NOT EDIT.

package contact_list

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

// NewListsPostParams creates a new ListsPostParams object
// with the default values initialized.
func NewListsPostParams() *ListsPostParams {
	var ()
	return &ListsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListsPostParamsWithTimeout creates a new ListsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListsPostParamsWithTimeout(timeout time.Duration) *ListsPostParams {
	var ()
	return &ListsPostParams{

		timeout: timeout,
	}
}

// NewListsPostParamsWithContext creates a new ListsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewListsPostParamsWithContext(ctx context.Context) *ListsPostParams {
	var ()
	return &ListsPostParams{

		Context: ctx,
	}
}

// NewListsPostParamsWithHTTPClient creates a new ListsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListsPostParamsWithHTTPClient(client *http.Client) *ListsPostParams {
	var ()
	return &ListsPostParams{
		HTTPClient: client,
	}
}

/*ListsPostParams contains all the parameters to send to the API endpoint
for the lists post operation typically these are written to a http.Request
*/
type ListsPostParams struct {

	/*ContactList
	  Contact list model

	*/
	ContactList models.ContactList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lists post params
func (o *ListsPostParams) WithTimeout(timeout time.Duration) *ListsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists post params
func (o *ListsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists post params
func (o *ListsPostParams) WithContext(ctx context.Context) *ListsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists post params
func (o *ListsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists post params
func (o *ListsPostParams) WithHTTPClient(client *http.Client) *ListsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists post params
func (o *ListsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactList adds the contactList to the lists post params
func (o *ListsPostParams) WithContactList(contactList models.ContactList) *ListsPostParams {
	o.SetContactList(contactList)
	return o
}

// SetContactList adds the contactList to the lists post params
func (o *ListsPostParams) SetContactList(contactList models.ContactList) {
	o.ContactList = contactList
}

// WriteToRequest writes these params to a swagger request
func (o *ListsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.ContactList); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
