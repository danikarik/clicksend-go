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

// NewRechargePurchaseByPackageIDPutParams creates a new RechargePurchaseByPackageIDPutParams object
// with the default values initialized.
func NewRechargePurchaseByPackageIDPutParams() *RechargePurchaseByPackageIDPutParams {
	var ()
	return &RechargePurchaseByPackageIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRechargePurchaseByPackageIDPutParamsWithTimeout creates a new RechargePurchaseByPackageIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRechargePurchaseByPackageIDPutParamsWithTimeout(timeout time.Duration) *RechargePurchaseByPackageIDPutParams {
	var ()
	return &RechargePurchaseByPackageIDPutParams{

		timeout: timeout,
	}
}

// NewRechargePurchaseByPackageIDPutParamsWithContext creates a new RechargePurchaseByPackageIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewRechargePurchaseByPackageIDPutParamsWithContext(ctx context.Context) *RechargePurchaseByPackageIDPutParams {
	var ()
	return &RechargePurchaseByPackageIDPutParams{

		Context: ctx,
	}
}

// NewRechargePurchaseByPackageIDPutParamsWithHTTPClient creates a new RechargePurchaseByPackageIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRechargePurchaseByPackageIDPutParamsWithHTTPClient(client *http.Client) *RechargePurchaseByPackageIDPutParams {
	var ()
	return &RechargePurchaseByPackageIDPutParams{
		HTTPClient: client,
	}
}

/*RechargePurchaseByPackageIDPutParams contains all the parameters to send to the API endpoint
for the recharge purchase by package Id put operation typically these are written to a http.Request
*/
type RechargePurchaseByPackageIDPutParams struct {

	/*PackageID
	  ID of package to purchase

	*/
	PackageID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) WithTimeout(timeout time.Duration) *RechargePurchaseByPackageIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) WithContext(ctx context.Context) *RechargePurchaseByPackageIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) WithHTTPClient(client *http.Client) *RechargePurchaseByPackageIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) WithPackageID(packageID int32) *RechargePurchaseByPackageIDPutParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the recharge purchase by package Id put params
func (o *RechargePurchaseByPackageIDPutParams) SetPackageID(packageID int32) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *RechargePurchaseByPackageIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param package_id
	if err := r.SetPathParam("package_id", swag.FormatInt32(o.PackageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
