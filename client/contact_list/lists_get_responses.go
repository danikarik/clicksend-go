// Code generated by go-swagger; DO NOT EDIT.

package contact_list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ListsGetReader is a Reader for the ListsGet structure.
type ListsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListsGetOK creates a ListsGetOK with default headers values
func NewListsGetOK() *ListsGetOK {
	return &ListsGetOK{}
}

/*ListsGetOK handles this case with default header values.

SUCCESS
*/
type ListsGetOK struct {
	Payload interface{}
}

func (o *ListsGetOK) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetOK  %+v", 200, o.Payload)
}

func (o *ListsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetBadRequest creates a ListsGetBadRequest with default headers values
func NewListsGetBadRequest() *ListsGetBadRequest {
	return &ListsGetBadRequest{}
}

/*ListsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ListsGetBadRequest struct {
	Payload interface{}
}

func (o *ListsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ListsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetUnauthorized creates a ListsGetUnauthorized with default headers values
func NewListsGetUnauthorized() *ListsGetUnauthorized {
	return &ListsGetUnauthorized{}
}

/*ListsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ListsGetUnauthorized struct {
	Payload interface{}
}

func (o *ListsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ListsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetForbidden creates a ListsGetForbidden with default headers values
func NewListsGetForbidden() *ListsGetForbidden {
	return &ListsGetForbidden{}
}

/*ListsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type ListsGetForbidden struct {
	Payload interface{}
}

func (o *ListsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetForbidden  %+v", 403, o.Payload)
}

func (o *ListsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetNotFound creates a ListsGetNotFound with default headers values
func NewListsGetNotFound() *ListsGetNotFound {
	return &ListsGetNotFound{}
}

/*ListsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type ListsGetNotFound struct {
	Payload interface{}
}

func (o *ListsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetNotFound  %+v", 404, o.Payload)
}

func (o *ListsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetMethodNotAllowed creates a ListsGetMethodNotAllowed with default headers values
func NewListsGetMethodNotAllowed() *ListsGetMethodNotAllowed {
	return &ListsGetMethodNotAllowed{}
}

/*ListsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ListsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *ListsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetTooManyRequests creates a ListsGetTooManyRequests with default headers values
func NewListsGetTooManyRequests() *ListsGetTooManyRequests {
	return &ListsGetTooManyRequests{}
}

/*ListsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ListsGetTooManyRequests struct {
	Payload interface{}
}

func (o *ListsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /lists][%d] listsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsGetDefault creates a ListsGetDefault with default headers values
func NewListsGetDefault(code int) *ListsGetDefault {
	return &ListsGetDefault{
		_statusCode: code,
	}
}

/*ListsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ListsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the lists get default response
func (o *ListsGetDefault) Code() int {
	return o._statusCode
}

func (o *ListsGetDefault) Error() string {
	return fmt.Sprintf("[GET /lists][%d] ListsGet default  %+v", o._statusCode, o.Payload)
}

func (o *ListsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
