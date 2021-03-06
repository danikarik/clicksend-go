// Code generated by go-swagger; DO NOT EDIT.

package master_email_templates

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

// NewMasterEmailTemplatesInCategoryGetParams creates a new MasterEmailTemplatesInCategoryGetParams object
// with the default values initialized.
func NewMasterEmailTemplatesInCategoryGetParams() *MasterEmailTemplatesInCategoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MasterEmailTemplatesInCategoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMasterEmailTemplatesInCategoryGetParamsWithTimeout creates a new MasterEmailTemplatesInCategoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMasterEmailTemplatesInCategoryGetParamsWithTimeout(timeout time.Duration) *MasterEmailTemplatesInCategoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MasterEmailTemplatesInCategoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewMasterEmailTemplatesInCategoryGetParamsWithContext creates a new MasterEmailTemplatesInCategoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMasterEmailTemplatesInCategoryGetParamsWithContext(ctx context.Context) *MasterEmailTemplatesInCategoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MasterEmailTemplatesInCategoryGetParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewMasterEmailTemplatesInCategoryGetParamsWithHTTPClient creates a new MasterEmailTemplatesInCategoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMasterEmailTemplatesInCategoryGetParamsWithHTTPClient(client *http.Client) *MasterEmailTemplatesInCategoryGetParams {
	var (
		limitDefault = int32(10)
		pageDefault  = int32(1)
	)
	return &MasterEmailTemplatesInCategoryGetParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*MasterEmailTemplatesInCategoryGetParams contains all the parameters to send to the API endpoint
for the master email templates in category get operation typically these are written to a http.Request
*/
type MasterEmailTemplatesInCategoryGetParams struct {

	/*CategoryID
	  Email category id

	*/
	CategoryID int32
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

// WithTimeout adds the timeout to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithTimeout(timeout time.Duration) *MasterEmailTemplatesInCategoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithContext(ctx context.Context) *MasterEmailTemplatesInCategoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithHTTPClient(client *http.Client) *MasterEmailTemplatesInCategoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithCategoryID(categoryID int32) *MasterEmailTemplatesInCategoryGetParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetCategoryID(categoryID int32) {
	o.CategoryID = categoryID
}

// WithLimit adds the limit to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithLimit(limit *int32) *MasterEmailTemplatesInCategoryGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) WithPage(page *int32) *MasterEmailTemplatesInCategoryGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the master email templates in category get params
func (o *MasterEmailTemplatesInCategoryGetParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *MasterEmailTemplatesInCategoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category_id
	if err := r.SetPathParam("category_id", swag.FormatInt32(o.CategoryID)); err != nil {
		return err
	}

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
