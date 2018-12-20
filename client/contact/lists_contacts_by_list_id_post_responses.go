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

// ListsContactsByListIDPostReader is a Reader for the ListsContactsByListIDPost structure.
type ListsContactsByListIDPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsContactsByListIDPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListsContactsByListIDPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListsContactsByListIDPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListsContactsByListIDPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListsContactsByListIDPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListsContactsByListIDPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListsContactsByListIDPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListsContactsByListIDPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListsContactsByListIDPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListsContactsByListIDPostOK creates a ListsContactsByListIDPostOK with default headers values
func NewListsContactsByListIDPostOK() *ListsContactsByListIDPostOK {
	return &ListsContactsByListIDPostOK{}
}

/*ListsContactsByListIDPostOK handles this case with default header values.

SUCCESS
*/
type ListsContactsByListIDPostOK struct {
	Payload string
}

func (o *ListsContactsByListIDPostOK) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostOK  %+v", 200, o.Payload)
}

func (o *ListsContactsByListIDPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostBadRequest creates a ListsContactsByListIDPostBadRequest with default headers values
func NewListsContactsByListIDPostBadRequest() *ListsContactsByListIDPostBadRequest {
	return &ListsContactsByListIDPostBadRequest{}
}

/*ListsContactsByListIDPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ListsContactsByListIDPostBadRequest struct {
	Payload string
}

func (o *ListsContactsByListIDPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostBadRequest  %+v", 400, o.Payload)
}

func (o *ListsContactsByListIDPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostUnauthorized creates a ListsContactsByListIDPostUnauthorized with default headers values
func NewListsContactsByListIDPostUnauthorized() *ListsContactsByListIDPostUnauthorized {
	return &ListsContactsByListIDPostUnauthorized{}
}

/*ListsContactsByListIDPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ListsContactsByListIDPostUnauthorized struct {
	Payload string
}

func (o *ListsContactsByListIDPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ListsContactsByListIDPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostForbidden creates a ListsContactsByListIDPostForbidden with default headers values
func NewListsContactsByListIDPostForbidden() *ListsContactsByListIDPostForbidden {
	return &ListsContactsByListIDPostForbidden{}
}

/*ListsContactsByListIDPostForbidden handles this case with default header values.

FORBIDDEN
*/
type ListsContactsByListIDPostForbidden struct {
	Payload string
}

func (o *ListsContactsByListIDPostForbidden) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostForbidden  %+v", 403, o.Payload)
}

func (o *ListsContactsByListIDPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostNotFound creates a ListsContactsByListIDPostNotFound with default headers values
func NewListsContactsByListIDPostNotFound() *ListsContactsByListIDPostNotFound {
	return &ListsContactsByListIDPostNotFound{}
}

/*ListsContactsByListIDPostNotFound handles this case with default header values.

NOT_FOUND
*/
type ListsContactsByListIDPostNotFound struct {
	Payload string
}

func (o *ListsContactsByListIDPostNotFound) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostNotFound  %+v", 404, o.Payload)
}

func (o *ListsContactsByListIDPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostMethodNotAllowed creates a ListsContactsByListIDPostMethodNotAllowed with default headers values
func NewListsContactsByListIDPostMethodNotAllowed() *ListsContactsByListIDPostMethodNotAllowed {
	return &ListsContactsByListIDPostMethodNotAllowed{}
}

/*ListsContactsByListIDPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ListsContactsByListIDPostMethodNotAllowed struct {
	Payload string
}

func (o *ListsContactsByListIDPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListsContactsByListIDPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostTooManyRequests creates a ListsContactsByListIDPostTooManyRequests with default headers values
func NewListsContactsByListIDPostTooManyRequests() *ListsContactsByListIDPostTooManyRequests {
	return &ListsContactsByListIDPostTooManyRequests{}
}

/*ListsContactsByListIDPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ListsContactsByListIDPostTooManyRequests struct {
	Payload string
}

func (o *ListsContactsByListIDPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] listsContactsByListIdPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListsContactsByListIDPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsContactsByListIDPostDefault creates a ListsContactsByListIDPostDefault with default headers values
func NewListsContactsByListIDPostDefault(code int) *ListsContactsByListIDPostDefault {
	return &ListsContactsByListIDPostDefault{
		_statusCode: code,
	}
}

/*ListsContactsByListIDPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ListsContactsByListIDPostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the lists contacts by list Id post default response
func (o *ListsContactsByListIDPostDefault) Code() int {
	return o._statusCode
}

func (o *ListsContactsByListIDPostDefault) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/contacts][%d] ListsContactsByListIdPost default  %+v", o._statusCode, o.Payload)
}

func (o *ListsContactsByListIDPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}