// Code generated by go-swagger; DO NOT EDIT.

package contact

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

// NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams creates a new ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams object
// with the default values initialized.
func NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams() *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	var ()
	return &ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithTimeout creates a new ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithTimeout(timeout time.Duration) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	var ()
	return &ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams{

		timeout: timeout,
	}
}

// NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithContext creates a new ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithContext(ctx context.Context) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	var ()
	return &ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams{

		Context: ctx,
	}
}

// NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithHTTPClient creates a new ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParamsWithHTTPClient(client *http.Client) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	var ()
	return &ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams{
		HTTPClient: client,
	}
}

/*ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams contains all the parameters to send to the API endpoint
for the lists remove opted out contacts by list Id and opt out list Id put operation typically these are written to a http.Request
*/
type ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams struct {

	/*ListID
	  Your list id

	*/
	ListID int32
	/*OptOutListID
	  Your opt out list id

	*/
	OptOutListID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WithTimeout(timeout time.Duration) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WithContext(ctx context.Context) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WithHTTPClient(client *http.Client) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WithListID(listID int32) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) SetListID(listID int32) {
	o.ListID = listID
}

// WithOptOutListID adds the optOutListID to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WithOptOutListID(optOutListID int32) *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams {
	o.SetOptOutListID(optOutListID)
	return o
}

// SetOptOutListID adds the optOutListId to the lists remove opted out contacts by list Id and opt out list Id put params
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) SetOptOutListID(optOutListID int32) {
	o.OptOutListID = optOutListID
}

// WriteToRequest writes these params to a swagger request
func (o *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param list_id
	if err := r.SetPathParam("list_id", swag.FormatInt32(o.ListID)); err != nil {
		return err
	}

	// path param opt_out_list_id
	if err := r.SetPathParam("opt_out_list_id", swag.FormatInt32(o.OptOutListID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}