// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSSendPostReader is a Reader for the SMSSendPost structure.
type SMSSendPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSSendPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSSendPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSSendPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSSendPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSSendPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSSendPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSSendPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSSendPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSSendPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSSendPostOK creates a SMSSendPostOK with default headers values
func NewSMSSendPostOK() *SMSSendPostOK {
	return &SMSSendPostOK{}
}

/*SMSSendPostOK handles this case with default header values.

SUCCESS
*/
type SMSSendPostOK struct {
	Payload interface{}
}

func (o *SMSSendPostOK) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostOK  %+v", 200, o.Payload)
}

func (o *SMSSendPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostBadRequest creates a SMSSendPostBadRequest with default headers values
func NewSMSSendPostBadRequest() *SMSSendPostBadRequest {
	return &SMSSendPostBadRequest{}
}

/*SMSSendPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSSendPostBadRequest struct {
	Payload interface{}
}

func (o *SMSSendPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostBadRequest  %+v", 400, o.Payload)
}

func (o *SMSSendPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostUnauthorized creates a SMSSendPostUnauthorized with default headers values
func NewSMSSendPostUnauthorized() *SMSSendPostUnauthorized {
	return &SMSSendPostUnauthorized{}
}

/*SMSSendPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSSendPostUnauthorized struct {
	Payload interface{}
}

func (o *SMSSendPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSSendPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostForbidden creates a SMSSendPostForbidden with default headers values
func NewSMSSendPostForbidden() *SMSSendPostForbidden {
	return &SMSSendPostForbidden{}
}

/*SMSSendPostForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSSendPostForbidden struct {
	Payload interface{}
}

func (o *SMSSendPostForbidden) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostForbidden  %+v", 403, o.Payload)
}

func (o *SMSSendPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostNotFound creates a SMSSendPostNotFound with default headers values
func NewSMSSendPostNotFound() *SMSSendPostNotFound {
	return &SMSSendPostNotFound{}
}

/*SMSSendPostNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSSendPostNotFound struct {
	Payload interface{}
}

func (o *SMSSendPostNotFound) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostNotFound  %+v", 404, o.Payload)
}

func (o *SMSSendPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostMethodNotAllowed creates a SMSSendPostMethodNotAllowed with default headers values
func NewSMSSendPostMethodNotAllowed() *SMSSendPostMethodNotAllowed {
	return &SMSSendPostMethodNotAllowed{}
}

/*SMSSendPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSSendPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSSendPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSSendPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostTooManyRequests creates a SMSSendPostTooManyRequests with default headers values
func NewSMSSendPostTooManyRequests() *SMSSendPostTooManyRequests {
	return &SMSSendPostTooManyRequests{}
}

/*SMSSendPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSSendPostTooManyRequests struct {
	Payload interface{}
}

func (o *SMSSendPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] smsSendPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSSendPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSSendPostDefault creates a SMSSendPostDefault with default headers values
func NewSMSSendPostDefault(code int) *SMSSendPostDefault {
	return &SMSSendPostDefault{
		_statusCode: code,
	}
}

/*SMSSendPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSSendPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms send post default response
func (o *SMSSendPostDefault) Code() int {
	return o._statusCode
}

func (o *SMSSendPostDefault) Error() string {
	return fmt.Sprintf("[POST /sms/send][%d] SmsSendPost default  %+v", o._statusCode, o.Payload)
}

func (o *SMSSendPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
