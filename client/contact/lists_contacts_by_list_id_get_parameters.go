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

// NewListsContactsByListIDGetParams creates a new ListsContactsByListIDGetParams object
// with the default values initialized.
func NewListsContactsByListIDGetParams() *ListsContactsByListIDGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &ListsContactsByListIDGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListsContactsByListIDGetParamsWithTimeout creates a new ListsContactsByListIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListsContactsByListIDGetParamsWithTimeout(timeout time.Duration) *ListsContactsByListIDGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &ListsContactsByListIDGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewListsContactsByListIDGetParamsWithContext creates a new ListsContactsByListIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewListsContactsByListIDGetParamsWithContext(ctx context.Context) *ListsContactsByListIDGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &ListsContactsByListIDGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewListsContactsByListIDGetParamsWithHTTPClient creates a new ListsContactsByListIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListsContactsByListIDGetParamsWithHTTPClient(client *http.Client) *ListsContactsByListIDGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &ListsContactsByListIDGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*ListsContactsByListIDGetParams contains all the parameters to send to the API endpoint
for the lists contacts by list Id get operation typically these are written to a http.Request
*/
type ListsContactsByListIDGetParams struct {

	/*Limit
	  Number of records per page

	*/
	Limit *int32
	/*ListID
	  Contact list ID

	*/
	ListID int32
	/*Page
	  Page number

	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithTimeout(timeout time.Duration) *ListsContactsByListIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithContext(ctx context.Context) *ListsContactsByListIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithHTTPClient(client *http.Client) *ListsContactsByListIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithLimit(limit *int32) *ListsContactsByListIDGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithListID adds the listID to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithListID(listID int32) *ListsContactsByListIDGetParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetListID(listID int32) {
	o.ListID = listID
}

// WithPage adds the page to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) WithPage(page *int32) *ListsContactsByListIDGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the lists contacts by list Id get params
func (o *ListsContactsByListIDGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListsContactsByListIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param list_id
	if err := r.SetPathParam("list_id", swag.FormatInt32(o.ListID)); err != nil {
		return err
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
