// Code generated by go-swagger; DO NOT EDIT.

package transactional_email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new transactional email API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactional email API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EmailHistoryExportGet exports all transactional email history

Export all Transactional Email history
*/
func (a *Client) EmailHistoryExportGet(params *EmailHistoryExportGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmailHistoryExportGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailHistoryExportGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmailHistoryExportGet",
		Method:             "GET",
		PathPattern:        "/email/history/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailHistoryExportGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailHistoryExportGetOK), nil

}

/*
EmailHistoryGet gets all transactional email history

Get all transactional email history
*/
func (a *Client) EmailHistoryGet(params *EmailHistoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*EmailHistoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailHistoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmailHistoryGet",
		Method:             "GET",
		PathPattern:        "/email/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailHistoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailHistoryGetOK), nil

}

/*
EmailPricePost gets transactional email price

Get transactional email price
*/
func (a *Client) EmailPricePost(params *EmailPricePostParams, authInfo runtime.ClientAuthInfoWriter) (*EmailPricePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailPricePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmailPricePost",
		Method:             "POST",
		PathPattern:        "/email/price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailPricePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailPricePostOK), nil

}

/*
EmailSendPost sends transactional email

Send transactional email
*/
func (a *Client) EmailSendPost(params *EmailSendPostParams, authInfo runtime.ClientAuthInfoWriter) (*EmailSendPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailSendPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmailSendPost",
		Method:             "POST",
		PathPattern:        "/email/send",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailSendPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailSendPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
