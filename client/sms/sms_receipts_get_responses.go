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

// SMSReceiptsGetReader is a Reader for the SMSReceiptsGet structure.
type SMSReceiptsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSReceiptsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSReceiptsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSReceiptsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSReceiptsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSReceiptsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSReceiptsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSReceiptsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSReceiptsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSReceiptsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSReceiptsGetOK creates a SMSReceiptsGetOK with default headers values
func NewSMSReceiptsGetOK() *SMSReceiptsGetOK {
	return &SMSReceiptsGetOK{}
}

/*SMSReceiptsGetOK handles this case with default header values.

SUCCESS
*/
type SMSReceiptsGetOK struct {
	Payload string
}

func (o *SMSReceiptsGetOK) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetOK  %+v", 200, o.Payload)
}

func (o *SMSReceiptsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetBadRequest creates a SMSReceiptsGetBadRequest with default headers values
func NewSMSReceiptsGetBadRequest() *SMSReceiptsGetBadRequest {
	return &SMSReceiptsGetBadRequest{}
}

/*SMSReceiptsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSReceiptsGetBadRequest struct {
	Payload string
}

func (o *SMSReceiptsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetBadRequest  %+v", 400, o.Payload)
}

func (o *SMSReceiptsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetUnauthorized creates a SMSReceiptsGetUnauthorized with default headers values
func NewSMSReceiptsGetUnauthorized() *SMSReceiptsGetUnauthorized {
	return &SMSReceiptsGetUnauthorized{}
}

/*SMSReceiptsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSReceiptsGetUnauthorized struct {
	Payload string
}

func (o *SMSReceiptsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSReceiptsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetForbidden creates a SMSReceiptsGetForbidden with default headers values
func NewSMSReceiptsGetForbidden() *SMSReceiptsGetForbidden {
	return &SMSReceiptsGetForbidden{}
}

/*SMSReceiptsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSReceiptsGetForbidden struct {
	Payload string
}

func (o *SMSReceiptsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetForbidden  %+v", 403, o.Payload)
}

func (o *SMSReceiptsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetNotFound creates a SMSReceiptsGetNotFound with default headers values
func NewSMSReceiptsGetNotFound() *SMSReceiptsGetNotFound {
	return &SMSReceiptsGetNotFound{}
}

/*SMSReceiptsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSReceiptsGetNotFound struct {
	Payload string
}

func (o *SMSReceiptsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetNotFound  %+v", 404, o.Payload)
}

func (o *SMSReceiptsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetMethodNotAllowed creates a SMSReceiptsGetMethodNotAllowed with default headers values
func NewSMSReceiptsGetMethodNotAllowed() *SMSReceiptsGetMethodNotAllowed {
	return &SMSReceiptsGetMethodNotAllowed{}
}

/*SMSReceiptsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSReceiptsGetMethodNotAllowed struct {
	Payload string
}

func (o *SMSReceiptsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSReceiptsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetTooManyRequests creates a SMSReceiptsGetTooManyRequests with default headers values
func NewSMSReceiptsGetTooManyRequests() *SMSReceiptsGetTooManyRequests {
	return &SMSReceiptsGetTooManyRequests{}
}

/*SMSReceiptsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSReceiptsGetTooManyRequests struct {
	Payload string
}

func (o *SMSReceiptsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] smsReceiptsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSReceiptsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSReceiptsGetDefault creates a SMSReceiptsGetDefault with default headers values
func NewSMSReceiptsGetDefault(code int) *SMSReceiptsGetDefault {
	return &SMSReceiptsGetDefault{
		_statusCode: code,
	}
}

/*SMSReceiptsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSReceiptsGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms receipts get default response
func (o *SMSReceiptsGetDefault) Code() int {
	return o._statusCode
}

func (o *SMSReceiptsGetDefault) Error() string {
	return fmt.Sprintf("[GET /sms/receipts][%d] SmsReceiptsGet default  %+v", o._statusCode, o.Payload)
}

func (o *SMSReceiptsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}