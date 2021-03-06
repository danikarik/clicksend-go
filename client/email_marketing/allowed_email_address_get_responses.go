// Code generated by go-swagger; DO NOT EDIT.

package email_marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AllowedEmailAddressGetReader is a Reader for the AllowedEmailAddressGet structure.
type AllowedEmailAddressGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllowedEmailAddressGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllowedEmailAddressGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllowedEmailAddressGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAllowedEmailAddressGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAllowedEmailAddressGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllowedEmailAddressGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAllowedEmailAddressGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewAllowedEmailAddressGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAllowedEmailAddressGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllowedEmailAddressGetOK creates a AllowedEmailAddressGetOK with default headers values
func NewAllowedEmailAddressGetOK() *AllowedEmailAddressGetOK {
	return &AllowedEmailAddressGetOK{}
}

/*AllowedEmailAddressGetOK handles this case with default header values.

SUCCESS
*/
type AllowedEmailAddressGetOK struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetOK) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetOK  %+v", 200, o.Payload)
}

func (o *AllowedEmailAddressGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetBadRequest creates a AllowedEmailAddressGetBadRequest with default headers values
func NewAllowedEmailAddressGetBadRequest() *AllowedEmailAddressGetBadRequest {
	return &AllowedEmailAddressGetBadRequest{}
}

/*AllowedEmailAddressGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type AllowedEmailAddressGetBadRequest struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedEmailAddressGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetUnauthorized creates a AllowedEmailAddressGetUnauthorized with default headers values
func NewAllowedEmailAddressGetUnauthorized() *AllowedEmailAddressGetUnauthorized {
	return &AllowedEmailAddressGetUnauthorized{}
}

/*AllowedEmailAddressGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type AllowedEmailAddressGetUnauthorized struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedEmailAddressGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetForbidden creates a AllowedEmailAddressGetForbidden with default headers values
func NewAllowedEmailAddressGetForbidden() *AllowedEmailAddressGetForbidden {
	return &AllowedEmailAddressGetForbidden{}
}

/*AllowedEmailAddressGetForbidden handles this case with default header values.

FORBIDDEN
*/
type AllowedEmailAddressGetForbidden struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetForbidden) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetForbidden  %+v", 403, o.Payload)
}

func (o *AllowedEmailAddressGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetNotFound creates a AllowedEmailAddressGetNotFound with default headers values
func NewAllowedEmailAddressGetNotFound() *AllowedEmailAddressGetNotFound {
	return &AllowedEmailAddressGetNotFound{}
}

/*AllowedEmailAddressGetNotFound handles this case with default header values.

NOT_FOUND
*/
type AllowedEmailAddressGetNotFound struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetNotFound) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetNotFound  %+v", 404, o.Payload)
}

func (o *AllowedEmailAddressGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetMethodNotAllowed creates a AllowedEmailAddressGetMethodNotAllowed with default headers values
func NewAllowedEmailAddressGetMethodNotAllowed() *AllowedEmailAddressGetMethodNotAllowed {
	return &AllowedEmailAddressGetMethodNotAllowed{}
}

/*AllowedEmailAddressGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type AllowedEmailAddressGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AllowedEmailAddressGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetTooManyRequests creates a AllowedEmailAddressGetTooManyRequests with default headers values
func NewAllowedEmailAddressGetTooManyRequests() *AllowedEmailAddressGetTooManyRequests {
	return &AllowedEmailAddressGetTooManyRequests{}
}

/*AllowedEmailAddressGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type AllowedEmailAddressGetTooManyRequests struct {
	Payload interface{}
}

func (o *AllowedEmailAddressGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] allowedEmailAddressGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *AllowedEmailAddressGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressGetDefault creates a AllowedEmailAddressGetDefault with default headers values
func NewAllowedEmailAddressGetDefault(code int) *AllowedEmailAddressGetDefault {
	return &AllowedEmailAddressGetDefault{
		_statusCode: code,
	}
}

/*AllowedEmailAddressGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type AllowedEmailAddressGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the allowed email address get default response
func (o *AllowedEmailAddressGetDefault) Code() int {
	return o._statusCode
}

func (o *AllowedEmailAddressGetDefault) Error() string {
	return fmt.Sprintf("[GET /email/addresses][%d] AllowedEmailAddressGet default  %+v", o._statusCode, o.Payload)
}

func (o *AllowedEmailAddressGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
