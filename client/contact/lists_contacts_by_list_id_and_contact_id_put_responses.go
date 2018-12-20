// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ListsContactsByListIDAndContactIDPutReader is a Reader for the ListsContactsByListIDAndContactIDPut structure.
type ListsContactsByListIDAndContactIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsContactsByListIDAndContactIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListsContactsByListIDAndContactIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListsContactsByListIDAndContactIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListsContactsByListIDAndContactIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListsContactsByListIDAndContactIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListsContactsByListIDAndContactIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListsContactsByListIDAndContactIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListsContactsByListIDAndContactIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListsContactsByListIDAndContactIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListsContactsByListIDAndContactIDPutOK creates a ListsContactsByListIDAndContactIDPutOK with default headers values
func NewListsContactsByListIDAndContactIDPutOK() *ListsContactsByListIDAndContactIDPutOK {
	return &ListsContactsByListIDAndContactIDPutOK{}
}

/*ListsContactsByListIDAndContactIDPutOK handles this case with default header values.

SUCCESS
*/
type ListsContactsByListIDAndContactIDPutOK struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutOK  %+v", 200, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutBadRequest creates a ListsContactsByListIDAndContactIDPutBadRequest with default headers values
func NewListsContactsByListIDAndContactIDPutBadRequest() *ListsContactsByListIDAndContactIDPutBadRequest {
	return &ListsContactsByListIDAndContactIDPutBadRequest{}
}

/*ListsContactsByListIDAndContactIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ListsContactsByListIDAndContactIDPutBadRequest struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutUnauthorized creates a ListsContactsByListIDAndContactIDPutUnauthorized with default headers values
func NewListsContactsByListIDAndContactIDPutUnauthorized() *ListsContactsByListIDAndContactIDPutUnauthorized {
	return &ListsContactsByListIDAndContactIDPutUnauthorized{}
}

/*ListsContactsByListIDAndContactIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ListsContactsByListIDAndContactIDPutUnauthorized struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutForbidden creates a ListsContactsByListIDAndContactIDPutForbidden with default headers values
func NewListsContactsByListIDAndContactIDPutForbidden() *ListsContactsByListIDAndContactIDPutForbidden {
	return &ListsContactsByListIDAndContactIDPutForbidden{}
}

/*ListsContactsByListIDAndContactIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type ListsContactsByListIDAndContactIDPutForbidden struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutForbidden  %+v", 403, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutNotFound creates a ListsContactsByListIDAndContactIDPutNotFound with default headers values
func NewListsContactsByListIDAndContactIDPutNotFound() *ListsContactsByListIDAndContactIDPutNotFound {
	return &ListsContactsByListIDAndContactIDPutNotFound{}
}

/*ListsContactsByListIDAndContactIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type ListsContactsByListIDAndContactIDPutNotFound struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutNotFound  %+v", 404, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutMethodNotAllowed creates a ListsContactsByListIDAndContactIDPutMethodNotAllowed with default headers values
func NewListsContactsByListIDAndContactIDPutMethodNotAllowed() *ListsContactsByListIDAndContactIDPutMethodNotAllowed {
	return &ListsContactsByListIDAndContactIDPutMethodNotAllowed{}
}

/*ListsContactsByListIDAndContactIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ListsContactsByListIDAndContactIDPutMethodNotAllowed struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutTooManyRequests creates a ListsContactsByListIDAndContactIDPutTooManyRequests with default headers values
func NewListsContactsByListIDAndContactIDPutTooManyRequests() *ListsContactsByListIDAndContactIDPutTooManyRequests {
	return &ListsContactsByListIDAndContactIDPutTooManyRequests{}
}

/*ListsContactsByListIDAndContactIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ListsContactsByListIDAndContactIDPutTooManyRequests struct {
	Payload string
}

func (o *ListsContactsByListIDAndContactIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] listsContactsByListIdAndContactIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDAndContactIDPutDefault creates a ListsContactsByListIDAndContactIDPutDefault with default headers values
func NewListsContactsByListIDAndContactIDPutDefault(code int) *ListsContactsByListIDAndContactIDPutDefault {
	return &ListsContactsByListIDAndContactIDPutDefault{
		_statusCode: code,
	}
}

/*ListsContactsByListIDAndContactIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ListsContactsByListIDAndContactIDPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the lists contacts by list Id and contact Id put default response
func (o *ListsContactsByListIDAndContactIDPutDefault) Code() int {
	return o._statusCode
}

func (o *ListsContactsByListIDAndContactIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /lists/{list_id}/contacts/{contact_id}][%d] ListsContactsByListIdAndContactIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *ListsContactsByListIDAndContactIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}