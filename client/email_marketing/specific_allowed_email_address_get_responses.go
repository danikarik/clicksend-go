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

// SpecificAllowedEmailAddressGetReader is a Reader for the SpecificAllowedEmailAddressGet structure.
type SpecificAllowedEmailAddressGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecificAllowedEmailAddressGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpecificAllowedEmailAddressGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpecificAllowedEmailAddressGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSpecificAllowedEmailAddressGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSpecificAllowedEmailAddressGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpecificAllowedEmailAddressGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSpecificAllowedEmailAddressGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSpecificAllowedEmailAddressGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSpecificAllowedEmailAddressGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSpecificAllowedEmailAddressGetOK creates a SpecificAllowedEmailAddressGetOK with default headers values
func NewSpecificAllowedEmailAddressGetOK() *SpecificAllowedEmailAddressGetOK {
	return &SpecificAllowedEmailAddressGetOK{}
}

/*SpecificAllowedEmailAddressGetOK handles this case with default header values.

SUCCESS
*/
type SpecificAllowedEmailAddressGetOK struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetOK) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetOK  %+v", 200, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetBadRequest creates a SpecificAllowedEmailAddressGetBadRequest with default headers values
func NewSpecificAllowedEmailAddressGetBadRequest() *SpecificAllowedEmailAddressGetBadRequest {
	return &SpecificAllowedEmailAddressGetBadRequest{}
}

/*SpecificAllowedEmailAddressGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SpecificAllowedEmailAddressGetBadRequest struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetBadRequest  %+v", 400, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetUnauthorized creates a SpecificAllowedEmailAddressGetUnauthorized with default headers values
func NewSpecificAllowedEmailAddressGetUnauthorized() *SpecificAllowedEmailAddressGetUnauthorized {
	return &SpecificAllowedEmailAddressGetUnauthorized{}
}

/*SpecificAllowedEmailAddressGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SpecificAllowedEmailAddressGetUnauthorized struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetForbidden creates a SpecificAllowedEmailAddressGetForbidden with default headers values
func NewSpecificAllowedEmailAddressGetForbidden() *SpecificAllowedEmailAddressGetForbidden {
	return &SpecificAllowedEmailAddressGetForbidden{}
}

/*SpecificAllowedEmailAddressGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SpecificAllowedEmailAddressGetForbidden struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetForbidden) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetForbidden  %+v", 403, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetNotFound creates a SpecificAllowedEmailAddressGetNotFound with default headers values
func NewSpecificAllowedEmailAddressGetNotFound() *SpecificAllowedEmailAddressGetNotFound {
	return &SpecificAllowedEmailAddressGetNotFound{}
}

/*SpecificAllowedEmailAddressGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SpecificAllowedEmailAddressGetNotFound struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetNotFound) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetNotFound  %+v", 404, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetMethodNotAllowed creates a SpecificAllowedEmailAddressGetMethodNotAllowed with default headers values
func NewSpecificAllowedEmailAddressGetMethodNotAllowed() *SpecificAllowedEmailAddressGetMethodNotAllowed {
	return &SpecificAllowedEmailAddressGetMethodNotAllowed{}
}

/*SpecificAllowedEmailAddressGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SpecificAllowedEmailAddressGetMethodNotAllowed struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetTooManyRequests creates a SpecificAllowedEmailAddressGetTooManyRequests with default headers values
func NewSpecificAllowedEmailAddressGetTooManyRequests() *SpecificAllowedEmailAddressGetTooManyRequests {
	return &SpecificAllowedEmailAddressGetTooManyRequests{}
}

/*SpecificAllowedEmailAddressGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SpecificAllowedEmailAddressGetTooManyRequests struct {
	Payload string
}

func (o *SpecificAllowedEmailAddressGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressGetDefault creates a SpecificAllowedEmailAddressGetDefault with default headers values
func NewSpecificAllowedEmailAddressGetDefault(code int) *SpecificAllowedEmailAddressGetDefault {
	return &SpecificAllowedEmailAddressGetDefault{
		_statusCode: code,
	}
}

/*SpecificAllowedEmailAddressGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SpecificAllowedEmailAddressGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the specific allowed email address get default response
func (o *SpecificAllowedEmailAddressGetDefault) Code() int {
	return o._statusCode
}

func (o *SpecificAllowedEmailAddressGetDefault) Error() string {
	return fmt.Sprintf("[GET /email/addresses/{email_address_id}][%d] SpecificAllowedEmailAddressGet default  %+v", o._statusCode, o.Payload)
}

func (o *SpecificAllowedEmailAddressGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
