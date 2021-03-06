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

// PostReturnAddressesByReturnAddressIDPutReader is a Reader for the PostReturnAddressesByReturnAddressIDPut structure.
type PostReturnAddressesByReturnAddressIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReturnAddressesByReturnAddressIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostReturnAddressesByReturnAddressIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostReturnAddressesByReturnAddressIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostReturnAddressesByReturnAddressIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostReturnAddressesByReturnAddressIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostReturnAddressesByReturnAddressIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostReturnAddressesByReturnAddressIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostReturnAddressesByReturnAddressIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostReturnAddressesByReturnAddressIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostReturnAddressesByReturnAddressIDPutOK creates a PostReturnAddressesByReturnAddressIDPutOK with default headers values
func NewPostReturnAddressesByReturnAddressIDPutOK() *PostReturnAddressesByReturnAddressIDPutOK {
	return &PostReturnAddressesByReturnAddressIDPutOK{}
}

/*PostReturnAddressesByReturnAddressIDPutOK handles this case with default header values.

SUCCESS
*/
type PostReturnAddressesByReturnAddressIDPutOK struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutOK  %+v", 200, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutBadRequest creates a PostReturnAddressesByReturnAddressIDPutBadRequest with default headers values
func NewPostReturnAddressesByReturnAddressIDPutBadRequest() *PostReturnAddressesByReturnAddressIDPutBadRequest {
	return &PostReturnAddressesByReturnAddressIDPutBadRequest{}
}

/*PostReturnAddressesByReturnAddressIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostReturnAddressesByReturnAddressIDPutBadRequest struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutUnauthorized creates a PostReturnAddressesByReturnAddressIDPutUnauthorized with default headers values
func NewPostReturnAddressesByReturnAddressIDPutUnauthorized() *PostReturnAddressesByReturnAddressIDPutUnauthorized {
	return &PostReturnAddressesByReturnAddressIDPutUnauthorized{}
}

/*PostReturnAddressesByReturnAddressIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostReturnAddressesByReturnAddressIDPutUnauthorized struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutForbidden creates a PostReturnAddressesByReturnAddressIDPutForbidden with default headers values
func NewPostReturnAddressesByReturnAddressIDPutForbidden() *PostReturnAddressesByReturnAddressIDPutForbidden {
	return &PostReturnAddressesByReturnAddressIDPutForbidden{}
}

/*PostReturnAddressesByReturnAddressIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type PostReturnAddressesByReturnAddressIDPutForbidden struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutForbidden  %+v", 403, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutNotFound creates a PostReturnAddressesByReturnAddressIDPutNotFound with default headers values
func NewPostReturnAddressesByReturnAddressIDPutNotFound() *PostReturnAddressesByReturnAddressIDPutNotFound {
	return &PostReturnAddressesByReturnAddressIDPutNotFound{}
}

/*PostReturnAddressesByReturnAddressIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type PostReturnAddressesByReturnAddressIDPutNotFound struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutNotFound  %+v", 404, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutMethodNotAllowed creates a PostReturnAddressesByReturnAddressIDPutMethodNotAllowed with default headers values
func NewPostReturnAddressesByReturnAddressIDPutMethodNotAllowed() *PostReturnAddressesByReturnAddressIDPutMethodNotAllowed {
	return &PostReturnAddressesByReturnAddressIDPutMethodNotAllowed{}
}

/*PostReturnAddressesByReturnAddressIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostReturnAddressesByReturnAddressIDPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutTooManyRequests creates a PostReturnAddressesByReturnAddressIDPutTooManyRequests with default headers values
func NewPostReturnAddressesByReturnAddressIDPutTooManyRequests() *PostReturnAddressesByReturnAddressIDPutTooManyRequests {
	return &PostReturnAddressesByReturnAddressIDPutTooManyRequests{}
}

/*PostReturnAddressesByReturnAddressIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostReturnAddressesByReturnAddressIDPutTooManyRequests struct {
	Payload interface{}
}

func (o *PostReturnAddressesByReturnAddressIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] postReturnAddressesByReturnAddressIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostReturnAddressesByReturnAddressIDPutDefault creates a PostReturnAddressesByReturnAddressIDPutDefault with default headers values
func NewPostReturnAddressesByReturnAddressIDPutDefault(code int) *PostReturnAddressesByReturnAddressIDPutDefault {
	return &PostReturnAddressesByReturnAddressIDPutDefault{
		_statusCode: code,
	}
}

/*PostReturnAddressesByReturnAddressIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostReturnAddressesByReturnAddressIDPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the post return addresses by return address Id put default response
func (o *PostReturnAddressesByReturnAddressIDPutDefault) Code() int {
	return o._statusCode
}

func (o *PostReturnAddressesByReturnAddressIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /post/return-addresses/{return_address_id}][%d] PostReturnAddressesByReturnAddressIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *PostReturnAddressesByReturnAddressIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
