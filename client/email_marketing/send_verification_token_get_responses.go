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

// SendVerificationTokenGetReader is a Reader for the SendVerificationTokenGet structure.
type SendVerificationTokenGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVerificationTokenGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVerificationTokenGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVerificationTokenGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVerificationTokenGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVerificationTokenGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVerificationTokenGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSendVerificationTokenGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSendVerificationTokenGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSendVerificationTokenGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSendVerificationTokenGetOK creates a SendVerificationTokenGetOK with default headers values
func NewSendVerificationTokenGetOK() *SendVerificationTokenGetOK {
	return &SendVerificationTokenGetOK{}
}

/*SendVerificationTokenGetOK handles this case with default header values.

SUCCESS
*/
type SendVerificationTokenGetOK struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetOK) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetOK  %+v", 200, o.Payload)
}

func (o *SendVerificationTokenGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetBadRequest creates a SendVerificationTokenGetBadRequest with default headers values
func NewSendVerificationTokenGetBadRequest() *SendVerificationTokenGetBadRequest {
	return &SendVerificationTokenGetBadRequest{}
}

/*SendVerificationTokenGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SendVerificationTokenGetBadRequest struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetBadRequest  %+v", 400, o.Payload)
}

func (o *SendVerificationTokenGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetUnauthorized creates a SendVerificationTokenGetUnauthorized with default headers values
func NewSendVerificationTokenGetUnauthorized() *SendVerificationTokenGetUnauthorized {
	return &SendVerificationTokenGetUnauthorized{}
}

/*SendVerificationTokenGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SendVerificationTokenGetUnauthorized struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVerificationTokenGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetForbidden creates a SendVerificationTokenGetForbidden with default headers values
func NewSendVerificationTokenGetForbidden() *SendVerificationTokenGetForbidden {
	return &SendVerificationTokenGetForbidden{}
}

/*SendVerificationTokenGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SendVerificationTokenGetForbidden struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetForbidden) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetForbidden  %+v", 403, o.Payload)
}

func (o *SendVerificationTokenGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetNotFound creates a SendVerificationTokenGetNotFound with default headers values
func NewSendVerificationTokenGetNotFound() *SendVerificationTokenGetNotFound {
	return &SendVerificationTokenGetNotFound{}
}

/*SendVerificationTokenGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SendVerificationTokenGetNotFound struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetNotFound) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetNotFound  %+v", 404, o.Payload)
}

func (o *SendVerificationTokenGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetMethodNotAllowed creates a SendVerificationTokenGetMethodNotAllowed with default headers values
func NewSendVerificationTokenGetMethodNotAllowed() *SendVerificationTokenGetMethodNotAllowed {
	return &SendVerificationTokenGetMethodNotAllowed{}
}

/*SendVerificationTokenGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SendVerificationTokenGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SendVerificationTokenGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetTooManyRequests creates a SendVerificationTokenGetTooManyRequests with default headers values
func NewSendVerificationTokenGetTooManyRequests() *SendVerificationTokenGetTooManyRequests {
	return &SendVerificationTokenGetTooManyRequests{}
}

/*SendVerificationTokenGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SendVerificationTokenGetTooManyRequests struct {
	Payload interface{}
}

func (o *SendVerificationTokenGetTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] sendVerificationTokenGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SendVerificationTokenGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVerificationTokenGetDefault creates a SendVerificationTokenGetDefault with default headers values
func NewSendVerificationTokenGetDefault(code int) *SendVerificationTokenGetDefault {
	return &SendVerificationTokenGetDefault{
		_statusCode: code,
	}
}

/*SendVerificationTokenGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SendVerificationTokenGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the send verification token get default response
func (o *SendVerificationTokenGetDefault) Code() int {
	return o._statusCode
}

func (o *SendVerificationTokenGetDefault) Error() string {
	return fmt.Sprintf("[PUT /email/address-verify/{email_address_id}/send][%d] SendVerificationTokenGet default  %+v", o._statusCode, o.Payload)
}

func (o *SendVerificationTokenGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
