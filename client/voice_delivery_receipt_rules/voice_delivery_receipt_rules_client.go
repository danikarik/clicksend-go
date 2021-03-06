// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new voice delivery receipt rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voice delivery receipt rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
VoiceDeliveryReceiptAutomationDelete deletes voice delivery receipt automation

Delete voice delivery receipt automation
*/
func (a *Client) VoiceDeliveryReceiptAutomationDelete(params *VoiceDeliveryReceiptAutomationDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*VoiceDeliveryReceiptAutomationDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoiceDeliveryReceiptAutomationDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoiceDeliveryReceiptAutomationDelete",
		Method:             "DELETE",
		PathPattern:        "/automations/voice/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoiceDeliveryReceiptAutomationDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VoiceDeliveryReceiptAutomationDeleteOK), nil

}

/*
VoiceDeliveryReceiptAutomationGet gets specific voice delivery receipt automation

Get specific voice delivery receipt automation
*/
func (a *Client) VoiceDeliveryReceiptAutomationGet(params *VoiceDeliveryReceiptAutomationGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoiceDeliveryReceiptAutomationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoiceDeliveryReceiptAutomationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoiceDeliveryReceiptAutomationGet",
		Method:             "GET",
		PathPattern:        "/automations/voice/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoiceDeliveryReceiptAutomationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VoiceDeliveryReceiptAutomationGetOK), nil

}

/*
VoiceDeliveryReceiptAutomationPost creates voice delivery receipt automations

Create voice delivery receipt automations
*/
func (a *Client) VoiceDeliveryReceiptAutomationPost(params *VoiceDeliveryReceiptAutomationPostParams, authInfo runtime.ClientAuthInfoWriter) (*VoiceDeliveryReceiptAutomationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoiceDeliveryReceiptAutomationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoiceDeliveryReceiptAutomationPost",
		Method:             "POST",
		PathPattern:        "/automations/voice/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoiceDeliveryReceiptAutomationPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VoiceDeliveryReceiptAutomationPostOK), nil

}

/*
VoiceDeliveryReceiptAutomationPut updates voice delivery receipt automation

Update voice delivery receipt automation
*/
func (a *Client) VoiceDeliveryReceiptAutomationPut(params *VoiceDeliveryReceiptAutomationPutParams, authInfo runtime.ClientAuthInfoWriter) (*VoiceDeliveryReceiptAutomationPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoiceDeliveryReceiptAutomationPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoiceDeliveryReceiptAutomationPut",
		Method:             "PUT",
		PathPattern:        "/automations/voice/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoiceDeliveryReceiptAutomationPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VoiceDeliveryReceiptAutomationPutOK), nil

}

/*
VoiceDeliveryReceiptAutomationsGet gets all voice delivery receipt automations

Get all voice delivery receipt automations
*/
func (a *Client) VoiceDeliveryReceiptAutomationsGet(params *VoiceDeliveryReceiptAutomationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*VoiceDeliveryReceiptAutomationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVoiceDeliveryReceiptAutomationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VoiceDeliveryReceiptAutomationsGet",
		Method:             "GET",
		PathPattern:        "/automations/voice/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VoiceDeliveryReceiptAutomationsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VoiceDeliveryReceiptAutomationsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
