// Code generated by go-swagger; DO NOT EDIT.

package mms_campaign

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

// NewMMSCampaignsCancelByMMSCampaignIDPutParams creates a new MMSCampaignsCancelByMMSCampaignIDPutParams object
// with the default values initialized.
func NewMMSCampaignsCancelByMMSCampaignIDPutParams() *MMSCampaignsCancelByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsCancelByMMSCampaignIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithTimeout creates a new MMSCampaignsCancelByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithTimeout(timeout time.Duration) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsCancelByMMSCampaignIDPutParams{

		timeout: timeout,
	}
}

// NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithContext creates a new MMSCampaignsCancelByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithContext(ctx context.Context) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsCancelByMMSCampaignIDPutParams{

		Context: ctx,
	}
}

// NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithHTTPClient creates a new MMSCampaignsCancelByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMMSCampaignsCancelByMMSCampaignIDPutParamsWithHTTPClient(client *http.Client) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsCancelByMMSCampaignIDPutParams{
		HTTPClient: client,
	}
}

/*MMSCampaignsCancelByMMSCampaignIDPutParams contains all the parameters to send to the API endpoint
for the Mms campaigns cancel by Mms campaign Id put operation typically these are written to a http.Request
*/
type MMSCampaignsCancelByMMSCampaignIDPutParams struct {

	/*MMSCampaignID
	  ID of MMS Campaign to cancel

	*/
	MMSCampaignID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) WithTimeout(timeout time.Duration) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) WithContext(ctx context.Context) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) WithHTTPClient(client *http.Client) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMMSCampaignID adds the mMSCampaignID to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) WithMMSCampaignID(mMSCampaignID int32) *MMSCampaignsCancelByMMSCampaignIDPutParams {
	o.SetMMSCampaignID(mMSCampaignID)
	return o
}

// SetMMSCampaignID adds the mmsCampaignId to the Mms campaigns cancel by Mms campaign Id put params
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) SetMMSCampaignID(mMSCampaignID int32) {
	o.MMSCampaignID = mMSCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *MMSCampaignsCancelByMMSCampaignIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mms_campaign_id
	if err := r.SetPathParam("mms_campaign_id", swag.FormatInt32(o.MMSCampaignID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
