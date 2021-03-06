// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contact API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contact API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListsContactsByListIDAndContactIDDelete deletes a contact

Delete a contact
*/
func (a *Client) ListsContactsByListIDAndContactIDDelete(params *ListsContactsByListIDAndContactIDDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ListsContactsByListIDAndContactIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsContactsByListIDAndContactIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsContactsByListIdAndContactIdDelete",
		Method:             "DELETE",
		PathPattern:        "/lists/{list_id}/contacts/{contact_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsContactsByListIDAndContactIDDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsContactsByListIDAndContactIDDeleteOK), nil

}

/*
ListsContactsByListIDAndContactIDGet gets a specific contact

Get a specific contact
*/
func (a *Client) ListsContactsByListIDAndContactIDGet(params *ListsContactsByListIDAndContactIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*ListsContactsByListIDAndContactIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsContactsByListIDAndContactIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsContactsByListIdAndContactIdGet",
		Method:             "GET",
		PathPattern:        "/lists/{list_id}/contacts/{contact_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsContactsByListIDAndContactIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsContactsByListIDAndContactIDGetOK), nil

}

/*
ListsContactsByListIDAndContactIDPut updates specific contact

Update specific contact
*/
func (a *Client) ListsContactsByListIDAndContactIDPut(params *ListsContactsByListIDAndContactIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*ListsContactsByListIDAndContactIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsContactsByListIDAndContactIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsContactsByListIdAndContactIdPut",
		Method:             "PUT",
		PathPattern:        "/lists/{list_id}/contacts/{contact_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsContactsByListIDAndContactIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsContactsByListIDAndContactIDPutOK), nil

}

/*
ListsContactsByListIDGet gets all contacts in a list

Get all contacts in a list
*/
func (a *Client) ListsContactsByListIDGet(params *ListsContactsByListIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*ListsContactsByListIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsContactsByListIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsContactsByListIdGet",
		Method:             "GET",
		PathPattern:        "/lists/{list_id}/contacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsContactsByListIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsContactsByListIDGetOK), nil

}

/*
ListsContactsByListIDPost creates new contact

Create new contact
*/
func (a *Client) ListsContactsByListIDPost(params *ListsContactsByListIDPostParams, authInfo runtime.ClientAuthInfoWriter) (*ListsContactsByListIDPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsContactsByListIDPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsContactsByListIdPost",
		Method:             "POST",
		PathPattern:        "/lists/{list_id}/contacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsContactsByListIDPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsContactsByListIDPostOK), nil

}

/*
ListsRemoveOptedOutContactsByListIDAndOptOutListIDPut removes all opted out contacts

Remove all opted out contacts
*/
func (a *Client) ListsRemoveOptedOutContactsByListIDAndOptOutListIDPut(params *ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams, authInfo runtime.ClientAuthInfoWriter) (*ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsRemoveOptedOutContactsByListIDAndOptOutListIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsRemoveOptedOutContactsByListIdAndOptOutListIdPut",
		Method:             "PUT",
		PathPattern:        "/lists/{list_id}/remove-opted-out-contacts/{opt_out_list_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsRemoveOptedOutContactsByListIDAndOptOutListIDPutOK), nil

}

/*
ListsTransferContactPut transfers contact to another list

Transfer contact to another list
*/
func (a *Client) ListsTransferContactPut(params *ListsTransferContactPutParams, authInfo runtime.ClientAuthInfoWriter) (*ListsTransferContactPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListsTransferContactPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListsTransferContactPut",
		Method:             "PUT",
		PathPattern:        "/lists/{from_list_id}/contacts/{contact_id}/transfer/{to_list_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListsTransferContactPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListsTransferContactPutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
