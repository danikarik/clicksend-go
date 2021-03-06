// Code generated by go-swagger; DO NOT EDIT.

package sms_campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sms campaign API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sms campaign API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SMSCampaignBySMSCampaignIDGet gets specific sms campaign

Get specific sms campaign
*/
func (a *Client) SMSCampaignBySMSCampaignIDGet(params *SMSCampaignBySMSCampaignIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignBySMSCampaignIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignBySMSCampaignIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignBySmsCampaignIdGet",
		Method:             "GET",
		PathPattern:        "/sms-campaigns/{sms_campaign_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignBySMSCampaignIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignBySMSCampaignIDGetOK), nil

}

/*
SMSCampaignsBySMSCampaignIDPut updates sms campaign

Update sms campaign
*/
func (a *Client) SMSCampaignsBySMSCampaignIDPut(params *SMSCampaignsBySMSCampaignIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignsBySMSCampaignIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignsBySMSCampaignIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignsBySmsCampaignIdPut",
		Method:             "PUT",
		PathPattern:        "/sms-campaigns/{sms_campaign_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignsBySMSCampaignIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignsBySMSCampaignIDPutOK), nil

}

/*
SMSCampaignsCancelBySMSCampaignIDPut cancels sms campaign

Cancel sms campaign
*/
func (a *Client) SMSCampaignsCancelBySMSCampaignIDPut(params *SMSCampaignsCancelBySMSCampaignIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignsCancelBySMSCampaignIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignsCancelBySMSCampaignIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignsCancelBySmsCampaignIdPut",
		Method:             "PUT",
		PathPattern:        "/sms-campaigns/{sms_campaign_id}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignsCancelBySMSCampaignIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignsCancelBySMSCampaignIDPutOK), nil

}

/*
SMSCampaignsGet gets list of sms campaigns

Get list of sms campaigns
*/
func (a *Client) SMSCampaignsGet(params *SMSCampaignsGetParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignsGet",
		Method:             "GET",
		PathPattern:        "/sms-campaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignsGetOK), nil

}

/*
SMSCampaignsPricePost calculates price for sms campaign

Calculate price for sms campaign
*/
func (a *Client) SMSCampaignsPricePost(params *SMSCampaignsPricePostParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignsPricePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignsPricePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignsPricePost",
		Method:             "POST",
		PathPattern:        "/sms-campaigns/price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignsPricePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignsPricePostOK), nil

}

/*
SMSCampaignsSendPost creates sms campaign

Create sms campaign
*/
func (a *Client) SMSCampaignsSendPost(params *SMSCampaignsSendPostParams, authInfo runtime.ClientAuthInfoWriter) (*SMSCampaignsSendPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSCampaignsSendPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsCampaignsSendPost",
		Method:             "POST",
		PathPattern:        "/sms-campaigns/send",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSCampaignsSendPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSCampaignsSendPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
