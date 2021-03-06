// Code generated by go-swagger; DO NOT EDIT.

package post_return_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostReturnAddressesPostReader is a Reader for the PostReturnAddressesPost structure.
type PostReturnAddressesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReturnAddressesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostReturnAddressesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostReturnAddressesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostReturnAddressesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostReturnAddressesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostReturnAddressesPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostReturnAddressesPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostReturnAddressesPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostReturnAddressesPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostReturnAddressesPostOK creates a PostReturnAddressesPostOK with default headers values
func NewPostReturnAddressesPostOK() *PostReturnAddressesPostOK {
	return &PostReturnAddressesPostOK{}
}

/*PostReturnAddressesPostOK handles this case with default header values.

SUCCESS
*/
type PostReturnAddressesPostOK struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostOK) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostOK  %+v", 200, o.Payload)
}

func (o *PostReturnAddressesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostBadRequest creates a PostReturnAddressesPostBadRequest with default headers values
func NewPostReturnAddressesPostBadRequest() *PostReturnAddressesPostBadRequest {
	return &PostReturnAddressesPostBadRequest{}
}

/*PostReturnAddressesPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostReturnAddressesPostBadRequest struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PostReturnAddressesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostUnauthorized creates a PostReturnAddressesPostUnauthorized with default headers values
func NewPostReturnAddressesPostUnauthorized() *PostReturnAddressesPostUnauthorized {
	return &PostReturnAddressesPostUnauthorized{}
}

/*PostReturnAddressesPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostReturnAddressesPostUnauthorized struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PostReturnAddressesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostForbidden creates a PostReturnAddressesPostForbidden with default headers values
func NewPostReturnAddressesPostForbidden() *PostReturnAddressesPostForbidden {
	return &PostReturnAddressesPostForbidden{}
}

/*PostReturnAddressesPostForbidden handles this case with default header values.

FORBIDDEN
*/
type PostReturnAddressesPostForbidden struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostForbidden) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostForbidden  %+v", 403, o.Payload)
}

func (o *PostReturnAddressesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostNotFound creates a PostReturnAddressesPostNotFound with default headers values
func NewPostReturnAddressesPostNotFound() *PostReturnAddressesPostNotFound {
	return &PostReturnAddressesPostNotFound{}
}

/*PostReturnAddressesPostNotFound handles this case with default header values.

NOT_FOUND
*/
type PostReturnAddressesPostNotFound struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostNotFound) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostNotFound  %+v", 404, o.Payload)
}

func (o *PostReturnAddressesPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostMethodNotAllowed creates a PostReturnAddressesPostMethodNotAllowed with default headers values
func NewPostReturnAddressesPostMethodNotAllowed() *PostReturnAddressesPostMethodNotAllowed {
	return &PostReturnAddressesPostMethodNotAllowed{}
}

/*PostReturnAddressesPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostReturnAddressesPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostReturnAddressesPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostTooManyRequests creates a PostReturnAddressesPostTooManyRequests with default headers values
func NewPostReturnAddressesPostTooManyRequests() *PostReturnAddressesPostTooManyRequests {
	return &PostReturnAddressesPostTooManyRequests{}
}

/*PostReturnAddressesPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostReturnAddressesPostTooManyRequests struct {
	Payload interface{}
}

func (o *PostReturnAddressesPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] postReturnAddressesPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostReturnAddressesPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesPostDefault creates a PostReturnAddressesPostDefault with default headers values
func NewPostReturnAddressesPostDefault(code int) *PostReturnAddressesPostDefault {
	return &PostReturnAddressesPostDefault{
		_statusCode: code,
	}
}

/*PostReturnAddressesPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostReturnAddressesPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the post return addresses post default response
func (o *PostReturnAddressesPostDefault) Code() int {
	return o._statusCode
}

func (o *PostReturnAddressesPostDefault) Error() string {
	return fmt.Sprintf("[POST /post/return-addresses][%d] PostReturnAddressesPost default  %+v", o._statusCode, o.Payload)
}

func (o *PostReturnAddressesPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
