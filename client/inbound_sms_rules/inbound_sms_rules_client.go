// Code generated by go-swagger; DO NOT EDIT.

package inbound_sms_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new inbound sms rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inbound sms rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SMSInboundAutomationDelete deletes inbound sms automation

Delete inbound sms automation
*/
func (a *Client) SMSInboundAutomationDelete(params *SMSInboundAutomationDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SMSInboundAutomationDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSInboundAutomationDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsInboundAutomationDelete",
		Method:             "DELETE",
		PathPattern:        "/automations/sms/inbound/{inbound_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSInboundAutomationDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSInboundAutomationDeleteOK), nil

}

/*
SMSInboundAutomationGet gets specific inbound sms automation

Get specific inbound sms automation
*/
func (a *Client) SMSInboundAutomationGet(params *SMSInboundAutomationGetParams, authInfo runtime.ClientAuthInfoWriter) (*SMSInboundAutomationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSInboundAutomationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsInboundAutomationGet",
		Method:             "GET",
		PathPattern:        "/automations/sms/inbound/{inbound_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSInboundAutomationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSInboundAutomationGetOK), nil

}

/*
SMSInboundAutomationPost creates new inbound sms automation

Create new inbound sms automation
*/
func (a *Client) SMSInboundAutomationPost(params *SMSInboundAutomationPostParams, authInfo runtime.ClientAuthInfoWriter) (*SMSInboundAutomationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSInboundAutomationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsInboundAutomationPost",
		Method:             "POST",
		PathPattern:        "/automations/sms/inbound",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSInboundAutomationPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSInboundAutomationPostOK), nil

}

/*
SMSInboundAutomationPut updates inbound sms automation

Update inbound sms automation
*/
func (a *Client) SMSInboundAutomationPut(params *SMSInboundAutomationPutParams, authInfo runtime.ClientAuthInfoWriter) (*SMSInboundAutomationPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSInboundAutomationPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsInboundAutomationPut",
		Method:             "PUT",
		PathPattern:        "/automations/sms/inbound/{inbound_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSInboundAutomationPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSInboundAutomationPutOK), nil

}

/*
SMSInboundAutomationsGet gets all inbound sms automations

Get all inbound sms automations
*/
func (a *Client) SMSInboundAutomationsGet(params *SMSInboundAutomationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*SMSInboundAutomationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSMSInboundAutomationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SmsInboundAutomationsGet",
		Method:             "GET",
		PathPattern:        "/automations/sms/inbound",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SMSInboundAutomationsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SMSInboundAutomationsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
