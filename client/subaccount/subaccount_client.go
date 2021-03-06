// Code generated by go-swagger; DO NOT EDIT.

package subaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new subaccount API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subaccount API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SubaccountsBySubaccountIDDelete deletes a subaccount

Delete a subaccount
*/
func (a *Client) SubaccountsBySubaccountIDDelete(params *SubaccountsBySubaccountIDDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsBySubaccountIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsBySubaccountIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsBySubaccountIdDelete",
		Method:             "DELETE",
		PathPattern:        "/subaccounts/{subaccount_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsBySubaccountIDDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsBySubaccountIDDeleteOK), nil

}

/*
SubaccountsBySubaccountIDGet gets specific subaccount

Get specific subaccount
*/
func (a *Client) SubaccountsBySubaccountIDGet(params *SubaccountsBySubaccountIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsBySubaccountIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsBySubaccountIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsBySubaccountIdGet",
		Method:             "GET",
		PathPattern:        "/subaccounts/{subaccount_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsBySubaccountIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsBySubaccountIDGetOK), nil

}

/*
SubaccountsBySubaccountIDPut updates subaccount

Update subaccount
*/
func (a *Client) SubaccountsBySubaccountIDPut(params *SubaccountsBySubaccountIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsBySubaccountIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsBySubaccountIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsBySubaccountIdPut",
		Method:             "PUT",
		PathPattern:        "/subaccounts/{subaccount_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsBySubaccountIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsBySubaccountIDPutOK), nil

}

/*
SubaccountsGet gets all subaccounts

Get all subaccounts
*/
func (a *Client) SubaccountsGet(params *SubaccountsGetParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsGet",
		Method:             "GET",
		PathPattern:        "/subaccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsGetOK), nil

}

/*
SubaccountsPost creates new subaccount

Create new subaccount
*/
func (a *Client) SubaccountsPost(params *SubaccountsPostParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsPost",
		Method:             "POST",
		PathPattern:        "/subaccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsPostOK), nil

}

/*
SubaccountsRegenAPIKeyBySubaccountIDPut regenerates an API key

Regenerate an API Key
*/
func (a *Client) SubaccountsRegenAPIKeyBySubaccountIDPut(params *SubaccountsRegenAPIKeyBySubaccountIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*SubaccountsRegenAPIKeyBySubaccountIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubaccountsRegenAPIKeyBySubaccountIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubaccountsRegenApiKeyBySubaccountIdPut",
		Method:             "PUT",
		PathPattern:        "/subaccounts/{subaccount_id}/regen-api-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubaccountsRegenAPIKeyBySubaccountIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SubaccountsRegenAPIKeyBySubaccountIDPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
