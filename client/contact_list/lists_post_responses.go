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

// ListsPostReader is a Reader for the ListsPost structure.
type ListsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListsPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListsPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListsPostOK creates a ListsPostOK with default headers values
func NewListsPostOK() *ListsPostOK {
	return &ListsPostOK{}
}

/*ListsPostOK handles this case with default header values.

SUCCESS
*/
type ListsPostOK struct {
	Payload interface{}
}

func (o *ListsPostOK) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostOK  %+v", 200, o.Payload)
}

func (o *ListsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostBadRequest creates a ListsPostBadRequest with default headers values
func NewListsPostBadRequest() *ListsPostBadRequest {
	return &ListsPostBadRequest{}
}

/*ListsPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ListsPostBadRequest struct {
	Payload interface{}
}

func (o *ListsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ListsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostUnauthorized creates a ListsPostUnauthorized with default headers values
func NewListsPostUnauthorized() *ListsPostUnauthorized {
	return &ListsPostUnauthorized{}
}

/*ListsPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ListsPostUnauthorized struct {
	Payload interface{}
}

func (o *ListsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ListsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostForbidden creates a ListsPostForbidden with default headers values
func NewListsPostForbidden() *ListsPostForbidden {
	return &ListsPostForbidden{}
}

/*ListsPostForbidden handles this case with default header values.

FORBIDDEN
*/
type ListsPostForbidden struct {
	Payload interface{}
}

func (o *ListsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostForbidden  %+v", 403, o.Payload)
}

func (o *ListsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostNotFound creates a ListsPostNotFound with default headers values
func NewListsPostNotFound() *ListsPostNotFound {
	return &ListsPostNotFound{}
}

/*ListsPostNotFound handles this case with default header values.

NOT_FOUND
*/
type ListsPostNotFound struct {
	Payload interface{}
}

func (o *ListsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostNotFound  %+v", 404, o.Payload)
}

func (o *ListsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostMethodNotAllowed creates a ListsPostMethodNotAllowed with default headers values
func NewListsPostMethodNotAllowed() *ListsPostMethodNotAllowed {
	return &ListsPostMethodNotAllowed{}
}

/*ListsPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ListsPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *ListsPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListsPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostTooManyRequests creates a ListsPostTooManyRequests with default headers values
func NewListsPostTooManyRequests() *ListsPostTooManyRequests {
	return &ListsPostTooManyRequests{}
}

/*ListsPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ListsPostTooManyRequests struct {
	Payload interface{}
}

func (o *ListsPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /lists][%d] listsPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListsPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsPostDefault creates a ListsPostDefault with default headers values
func NewListsPostDefault(code int) *ListsPostDefault {
	return &ListsPostDefault{
		_statusCode: code,
	}
}

/*ListsPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ListsPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the lists post default response
func (o *ListsPostDefault) Code() int {
	return o._statusCode
}

func (o *ListsPostDefault) Error() string {
	return fmt.Sprintf("[POST /lists][%d] ListsPost default  %+v", o._statusCode, o.Payload)
}

func (o *ListsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
