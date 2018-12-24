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

// ListsImportByListIDPostReader is a Reader for the ListsImportByListIDPost structure.
type ListsImportByListIDPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListsImportByListIDPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListsImportByListIDPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListsImportByListIDPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListsImportByListIDPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListsImportByListIDPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListsImportByListIDPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListsImportByListIDPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewListsImportByListIDPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListsImportByListIDPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListsImportByListIDPostOK creates a ListsImportByListIDPostOK with default headers values
func NewListsImportByListIDPostOK() *ListsImportByListIDPostOK {
	return &ListsImportByListIDPostOK{}
}

/*ListsImportByListIDPostOK handles this case with default header values.

SUCCESS
*/
type ListsImportByListIDPostOK struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostOK) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostOK  %+v", 200, o.Payload)
}

func (o *ListsImportByListIDPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostBadRequest creates a ListsImportByListIDPostBadRequest with default headers values
func NewListsImportByListIDPostBadRequest() *ListsImportByListIDPostBadRequest {
	return &ListsImportByListIDPostBadRequest{}
}

/*ListsImportByListIDPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ListsImportByListIDPostBadRequest struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostBadRequest  %+v", 400, o.Payload)
}

func (o *ListsImportByListIDPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostUnauthorized creates a ListsImportByListIDPostUnauthorized with default headers values
func NewListsImportByListIDPostUnauthorized() *ListsImportByListIDPostUnauthorized {
	return &ListsImportByListIDPostUnauthorized{}
}

/*ListsImportByListIDPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ListsImportByListIDPostUnauthorized struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ListsImportByListIDPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostForbidden creates a ListsImportByListIDPostForbidden with default headers values
func NewListsImportByListIDPostForbidden() *ListsImportByListIDPostForbidden {
	return &ListsImportByListIDPostForbidden{}
}

/*ListsImportByListIDPostForbidden handles this case with default header values.

FORBIDDEN
*/
type ListsImportByListIDPostForbidden struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostForbidden) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostForbidden  %+v", 403, o.Payload)
}

func (o *ListsImportByListIDPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostNotFound creates a ListsImportByListIDPostNotFound with default headers values
func NewListsImportByListIDPostNotFound() *ListsImportByListIDPostNotFound {
	return &ListsImportByListIDPostNotFound{}
}

/*ListsImportByListIDPostNotFound handles this case with default header values.

NOT_FOUND
*/
type ListsImportByListIDPostNotFound struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostNotFound) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostNotFound  %+v", 404, o.Payload)
}

func (o *ListsImportByListIDPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostMethodNotAllowed creates a ListsImportByListIDPostMethodNotAllowed with default headers values
func NewListsImportByListIDPostMethodNotAllowed() *ListsImportByListIDPostMethodNotAllowed {
	return &ListsImportByListIDPostMethodNotAllowed{}
}

/*ListsImportByListIDPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ListsImportByListIDPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListsImportByListIDPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostTooManyRequests creates a ListsImportByListIDPostTooManyRequests with default headers values
func NewListsImportByListIDPostTooManyRequests() *ListsImportByListIDPostTooManyRequests {
	return &ListsImportByListIDPostTooManyRequests{}
}

/*ListsImportByListIDPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ListsImportByListIDPostTooManyRequests struct {
	Payload interface{}
}

func (o *ListsImportByListIDPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] listsImportByListIdPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListsImportByListIDPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListsImportByListIDPostDefault creates a ListsImportByListIDPostDefault with default headers values
func NewListsImportByListIDPostDefault(code int) *ListsImportByListIDPostDefault {
	return &ListsImportByListIDPostDefault{
		_statusCode: code,
	}
}

/*ListsImportByListIDPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ListsImportByListIDPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the lists import by list Id post default response
func (o *ListsImportByListIDPostDefault) Code() int {
	return o._statusCode
}

func (o *ListsImportByListIDPostDefault) Error() string {
	return fmt.Sprintf("[POST /lists/{list_id}/import][%d] ListsImportByListIdPost default  %+v", o._statusCode, o.Payload)
}

func (o *ListsImportByListIDPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
