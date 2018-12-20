// Code generated by go-swagger; DO NOT EDIT.

package detect_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new detect address API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for detect address API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DetectAddressPost detects address in uploaded file

Detects address in uploaded file.
*/
func (a *Client) DetectAddressPost(params *DetectAddressPostParams, authInfo runtime.ClientAuthInfoWriter) (*DetectAddressPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetectAddressPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DetectAddressPost",
		Method:             "POST",
		PathPattern:        "/post/letters/detect-address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetectAddressPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetectAddressPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
