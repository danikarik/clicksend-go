// Code generated by go-swagger; DO NOT EDIT.

package referral_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new referral account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for referral account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ReferralAccountsGet gets all referral accounts

Get all referral accounts
*/
func (a *Client) ReferralAccountsGet(params *ReferralAccountsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ReferralAccountsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReferralAccountsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReferralAccountsGet",
		Method:             "GET",
		PathPattern:        "/referral/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReferralAccountsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReferralAccountsGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}