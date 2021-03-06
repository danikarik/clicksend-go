// Code generated by go-swagger; DO NOT EDIT.

package fax_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fax delivery receipt rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fax delivery receipt rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FAXDeliveryReceiptAutomationDelete deletes fax delivery receipt automation

Delete fax delivery receipt automation
*/
func (a *Client) FAXDeliveryReceiptAutomationDelete(params *FAXDeliveryReceiptAutomationDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*FAXDeliveryReceiptAutomationDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFAXDeliveryReceiptAutomationDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FaxDeliveryReceiptAutomationDelete",
		Method:             "DELETE",
		PathPattern:        "/automations/fax/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FAXDeliveryReceiptAutomationDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FAXDeliveryReceiptAutomationDeleteOK), nil

}

/*
FAXDeliveryReceiptAutomationGet gets specific fax delivery receipt automation

Get specific fax delivery receipt automation
*/
func (a *Client) FAXDeliveryReceiptAutomationGet(params *FAXDeliveryReceiptAutomationGetParams, authInfo runtime.ClientAuthInfoWriter) (*FAXDeliveryReceiptAutomationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFAXDeliveryReceiptAutomationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FaxDeliveryReceiptAutomationGet",
		Method:             "GET",
		PathPattern:        "/automations/fax/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FAXDeliveryReceiptAutomationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FAXDeliveryReceiptAutomationGetOK), nil

}

/*
FAXDeliveryReceiptAutomationPost creates fax delivery receipt automations

Create fax delivery receipt automations
*/
func (a *Client) FAXDeliveryReceiptAutomationPost(params *FAXDeliveryReceiptAutomationPostParams, authInfo runtime.ClientAuthInfoWriter) (*FAXDeliveryReceiptAutomationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFAXDeliveryReceiptAutomationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FaxDeliveryReceiptAutomationPost",
		Method:             "POST",
		PathPattern:        "/automations/fax/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FAXDeliveryReceiptAutomationPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FAXDeliveryReceiptAutomationPostOK), nil

}

/*
FAXDeliveryReceiptAutomationPut updates fax delivery receipt automation

Update fax delivery receipt automation
*/
func (a *Client) FAXDeliveryReceiptAutomationPut(params *FAXDeliveryReceiptAutomationPutParams, authInfo runtime.ClientAuthInfoWriter) (*FAXDeliveryReceiptAutomationPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFAXDeliveryReceiptAutomationPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FaxDeliveryReceiptAutomationPut",
		Method:             "PUT",
		PathPattern:        "/automations/fax/receipts/{receipt_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FAXDeliveryReceiptAutomationPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FAXDeliveryReceiptAutomationPutOK), nil

}

/*
FAXDeliveryReceiptAutomationsGet gets all fax delivery receipt automations

Get all fax delivery receipt automations
*/
func (a *Client) FAXDeliveryReceiptAutomationsGet(params *FAXDeliveryReceiptAutomationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*FAXDeliveryReceiptAutomationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFAXDeliveryReceiptAutomationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FaxDeliveryReceiptAutomationsGet",
		Method:             "GET",
		PathPattern:        "/automations/fax/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FAXDeliveryReceiptAutomationsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FAXDeliveryReceiptAutomationsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
