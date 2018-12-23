// Code generated by go-swagger; DO NOT EDIT.

package mms_campaign

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewMMSCampaignsByMMSCampaignIDPutParams creates a new MMSCampaignsByMMSCampaignIDPutParams object
// with the default values initialized.
func NewMMSCampaignsByMMSCampaignIDPutParams() *MMSCampaignsByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsByMMSCampaignIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMMSCampaignsByMMSCampaignIDPutParamsWithTimeout creates a new MMSCampaignsByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMMSCampaignsByMMSCampaignIDPutParamsWithTimeout(timeout time.Duration) *MMSCampaignsByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsByMMSCampaignIDPutParams{

		timeout: timeout,
	}
}

// NewMMSCampaignsByMMSCampaignIDPutParamsWithContext creates a new MMSCampaignsByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewMMSCampaignsByMMSCampaignIDPutParamsWithContext(ctx context.Context) *MMSCampaignsByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsByMMSCampaignIDPutParams{

		Context: ctx,
	}
}

// NewMMSCampaignsByMMSCampaignIDPutParamsWithHTTPClient creates a new MMSCampaignsByMMSCampaignIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMMSCampaignsByMMSCampaignIDPutParamsWithHTTPClient(client *http.Client) *MMSCampaignsByMMSCampaignIDPutParams {
	var ()
	return &MMSCampaignsByMMSCampaignIDPutParams{
		HTTPClient: client,
	}
}

/*MMSCampaignsByMMSCampaignIDPutParams contains all the parameters to send to the API endpoint
for the Mms campaigns by Mms campaign Id put operation typically these are written to a http.Request
*/
type MMSCampaignsByMMSCampaignIDPutParams struct {

	/*Campaign
	  MmsCampaign model

	*/
	Campaign models.MMSCampaign
	/*MMSCampaignID
	  ID of MMS campaign to update

	*/
	MMSCampaignID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) WithTimeout(timeout time.Duration) *MMSCampaignsByMMSCampaignIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) WithContext(ctx context.Context) *MMSCampaignsByMMSCampaignIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) WithHTTPClient(client *http.Client) *MMSCampaignsByMMSCampaignIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaign adds the campaign to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) WithCampaign(campaign models.MMSCampaign) *MMSCampaignsByMMSCampaignIDPutParams {
	o.SetCampaign(campaign)
	return o
}

// SetCampaign adds the campaign to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) SetCampaign(campaign models.MMSCampaign) {
	o.Campaign = campaign
}

// WithMMSCampaignID adds the mMSCampaignID to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) WithMMSCampaignID(mMSCampaignID int32) *MMSCampaignsByMMSCampaignIDPutParams {
	o.SetMMSCampaignID(mMSCampaignID)
	return o
}

// SetMMSCampaignID adds the mmsCampaignId to the Mms campaigns by Mms campaign Id put params
func (o *MMSCampaignsByMMSCampaignIDPutParams) SetMMSCampaignID(mMSCampaignID int32) {
	o.MMSCampaignID = mMSCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *MMSCampaignsByMMSCampaignIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Campaign); err != nil {
		return err
	}

	// path param mms_campaign_id
	if err := r.SetPathParam("mms_campaign_id", swag.FormatInt32(o.MMSCampaignID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
