// Code generated by go-swagger; DO NOT EDIT.

package post_return_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new post return address API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for post return address API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostReturnAddressesByReturnAddressIDDelete deletes specific post return address

Delete specific post return address
*/
func (a *Client) PostReturnAddressesByReturnAddressIDDelete(params *PostReturnAddressesByReturnAddressIDDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PostReturnAddressesByReturnAddressIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReturnAddressesByReturnAddressIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReturnAddressesByReturnAddressIdDelete",
		Method:             "DELETE",
		PathPattern:        "/post/return-addresses/{return_address_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReturnAddressesByReturnAddressIDDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReturnAddressesByReturnAddressIDDeleteOK), nil

}

/*
PostReturnAddressesByReturnAddressIDGet gets specific post return address

Get specific post return address
*/
func (a *Client) PostReturnAddressesByReturnAddressIDGet(params *PostReturnAddressesByReturnAddressIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*PostReturnAddressesByReturnAddressIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReturnAddressesByReturnAddressIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReturnAddressesByReturnAddressIdGet",
		Method:             "GET",
		PathPattern:        "/post/return-addresses/{return_address_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReturnAddressesByReturnAddressIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReturnAddressesByReturnAddressIDGetOK), nil

}

/*
PostReturnAddressesByReturnAddressIDPut updates post return address

Update post return address
*/
func (a *Client) PostReturnAddressesByReturnAddressIDPut(params *PostReturnAddressesByReturnAddressIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*PostReturnAddressesByReturnAddressIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReturnAddressesByReturnAddressIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReturnAddressesByReturnAddressIdPut",
		Method:             "PUT",
		PathPattern:        "/post/return-addresses/{return_address_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReturnAddressesByReturnAddressIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReturnAddressesByReturnAddressIDPutOK), nil

}

/*
PostReturnAddressesGet gets list of post return addresses

Get list of post return addresses
*/
func (a *Client) PostReturnAddressesGet(params *PostReturnAddressesGetParams, authInfo runtime.ClientAuthInfoWriter) (*PostReturnAddressesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReturnAddressesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReturnAddressesGet",
		Method:             "GET",
		PathPattern:        "/post/return-addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReturnAddressesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReturnAddressesGetOK), nil

}

/*
PostReturnAddressesPost creates post return address

Create post return address
*/
func (a *Client) PostReturnAddressesPost(params *PostReturnAddressesPostParams, authInfo runtime.ClientAuthInfoWriter) (*PostReturnAddressesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReturnAddressesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReturnAddressesPost",
		Method:             "POST",
		PathPattern:        "/post/return-addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostReturnAddressesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReturnAddressesPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}