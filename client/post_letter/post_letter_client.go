// Code generated by go-swagger; DO NOT EDIT.

package post_letter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new post letter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for post letter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostLettersExportGet exports post letter history

export post letter history
*/
func (a *Client) PostLettersExportGet(params *PostLettersExportGetParams, authInfo runtime.ClientAuthInfoWriter) (*PostLettersExportGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLettersExportGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLettersExportGet",
		Method:             "GET",
		PathPattern:        "/post/letters/history/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLettersExportGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLettersExportGetOK), nil

}

/*
PostLettersHistoryGet gets all post letter history

Get all post letter history
*/
func (a *Client) PostLettersHistoryGet(params *PostLettersHistoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*PostLettersHistoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLettersHistoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLettersHistoryGet",
		Method:             "GET",
		PathPattern:        "/post/letters/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLettersHistoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLettersHistoryGetOK), nil

}

/*
PostLettersPricePost calculates post letter price

Calculate post letter price
*/
func (a *Client) PostLettersPricePost(params *PostLettersPricePostParams, authInfo runtime.ClientAuthInfoWriter) (*PostLettersPricePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLettersPricePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLettersPricePost",
		Method:             "POST",
		PathPattern:        "/post/letters/price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLettersPricePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLettersPricePostOK), nil

}

/*
PostLettersSendPost sends post letter

Send post letter
*/
func (a *Client) PostLettersSendPost(params *PostLettersSendPostParams, authInfo runtime.ClientAuthInfoWriter) (*PostLettersSendPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLettersSendPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLettersSendPost",
		Method:             "POST",
		PathPattern:        "/post/letters/send",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLettersSendPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLettersSendPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
