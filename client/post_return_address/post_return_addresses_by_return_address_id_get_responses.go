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

// PostReturnAddressesByReturnAddressIDGetReader is a Reader for the PostReturnAddressesByReturnAddressIDGet structure.
type PostReturnAddressesByReturnAddressIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReturnAddressesByReturnAddressIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostReturnAddressesByReturnAddressIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostReturnAddressesByReturnAddressIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostReturnAddressesByReturnAddressIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostReturnAddressesByReturnAddressIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostReturnAddressesByReturnAddressIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostReturnAddressesByReturnAddressIDGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostReturnAddressesByReturnAddressIDGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostReturnAddressesByReturnAddressIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostReturnAddressesByReturnAddressIDGetOK creates a PostReturnAddressesByReturnAddressIDGetOK with default headers values
func NewPostReturnAddressesByReturnAddressIDGetOK() *PostReturnAddressesByReturnAddressIDGetOK {
	return &PostReturnAddressesByReturnAddressIDGetOK{}
}

/*PostReturnAddressesByReturnAddressIDGetOK handles this case with default header values.

SUCCESS
*/
type PostReturnAddressesByReturnAddressIDGetOK struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetOK) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetOK  %+v", 200, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetBadRequest creates a PostReturnAddressesByReturnAddressIDGetBadRequest with default headers values
func NewPostReturnAddressesByReturnAddressIDGetBadRequest() *PostReturnAddressesByReturnAddressIDGetBadRequest {
	return &PostReturnAddressesByReturnAddressIDGetBadRequest{}
}

/*PostReturnAddressesByReturnAddressIDGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostReturnAddressesByReturnAddressIDGetBadRequest struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetUnauthorized creates a PostReturnAddressesByReturnAddressIDGetUnauthorized with default headers values
func NewPostReturnAddressesByReturnAddressIDGetUnauthorized() *PostReturnAddressesByReturnAddressIDGetUnauthorized {
	return &PostReturnAddressesByReturnAddressIDGetUnauthorized{}
}

/*PostReturnAddressesByReturnAddressIDGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostReturnAddressesByReturnAddressIDGetUnauthorized struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetForbidden creates a PostReturnAddressesByReturnAddressIDGetForbidden with default headers values
func NewPostReturnAddressesByReturnAddressIDGetForbidden() *PostReturnAddressesByReturnAddressIDGetForbidden {
	return &PostReturnAddressesByReturnAddressIDGetForbidden{}
}

/*PostReturnAddressesByReturnAddressIDGetForbidden handles this case with default header values.

FORBIDDEN
*/
type PostReturnAddressesByReturnAddressIDGetForbidden struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetForbidden  %+v", 403, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetNotFound creates a PostReturnAddressesByReturnAddressIDGetNotFound with default headers values
func NewPostReturnAddressesByReturnAddressIDGetNotFound() *PostReturnAddressesByReturnAddressIDGetNotFound {
	return &PostReturnAddressesByReturnAddressIDGetNotFound{}
}

/*PostReturnAddressesByReturnAddressIDGetNotFound handles this case with default header values.

NOT_FOUND
*/
type PostReturnAddressesByReturnAddressIDGetNotFound struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetNotFound  %+v", 404, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetMethodNotAllowed creates a PostReturnAddressesByReturnAddressIDGetMethodNotAllowed with default headers values
func NewPostReturnAddressesByReturnAddressIDGetMethodNotAllowed() *PostReturnAddressesByReturnAddressIDGetMethodNotAllowed {
	return &PostReturnAddressesByReturnAddressIDGetMethodNotAllowed{}
}

/*PostReturnAddressesByReturnAddressIDGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostReturnAddressesByReturnAddressIDGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetTooManyRequests creates a PostReturnAddressesByReturnAddressIDGetTooManyRequests with default headers values
func NewPostReturnAddressesByReturnAddressIDGetTooManyRequests() *PostReturnAddressesByReturnAddressIDGetTooManyRequests {
	return &PostReturnAddressesByReturnAddressIDGetTooManyRequests{}
}

/*PostReturnAddressesByReturnAddressIDGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostReturnAddressesByReturnAddressIDGetTooManyRequests struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDGetDefault creates a PostReturnAddressesByReturnAddressIDGetDefault with default headers values
func NewPostReturnAddressesByReturnAddressIDGetDefault(code int) *PostReturnAddressesByReturnAddressIDGetDefault {
	return &PostReturnAddressesByReturnAddressIDGetDefault{
		_statusCode: code,
	}
}

/*PostReturnAddressesByReturnAddressIDGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostReturnAddressesByReturnAddressIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the post return addresses by return address Id get default response
func (o *PostReturnAddressesByReturnAddressIDGetDefault) Code() int {
	return o._statusCode
}

func (o *PostReturnAddressesByReturnAddressIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /post/return-addresses/{return_address_id}][%d] PostReturnAddressesByReturnAddressIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
