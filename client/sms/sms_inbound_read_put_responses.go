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

// SMSInboundReadPutReader is a Reader for the SMSInboundReadPut structure.
type SMSInboundReadPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundReadPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundReadPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundReadPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundReadPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundReadPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundReadPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundReadPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundReadPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundReadPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundReadPutOK creates a SMSInboundReadPutOK with default headers values
func NewSMSInboundReadPutOK() *SMSInboundReadPutOK {
	return &SMSInboundReadPutOK{}
}

/*SMSInboundReadPutOK handles this case with default header values.

SUCCESS
*/
type SMSInboundReadPutOK struct {
	Payload string
}

func (o *SMSInboundReadPutOK) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutOK  %+v", 200, o.Payload)
}

func (o *SMSInboundReadPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutBadRequest creates a SMSInboundReadPutBadRequest with default headers values
func NewSMSInboundReadPutBadRequest() *SMSInboundReadPutBadRequest {
	return &SMSInboundReadPutBadRequest{}
}

/*SMSInboundReadPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundReadPutBadRequest struct {
	Payload string
}

func (o *SMSInboundReadPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundReadPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutUnauthorized creates a SMSInboundReadPutUnauthorized with default headers values
func NewSMSInboundReadPutUnauthorized() *SMSInboundReadPutUnauthorized {
	return &SMSInboundReadPutUnauthorized{}
}

/*SMSInboundReadPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundReadPutUnauthorized struct {
	Payload string
}

func (o *SMSInboundReadPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundReadPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutForbidden creates a SMSInboundReadPutForbidden with default headers values
func NewSMSInboundReadPutForbidden() *SMSInboundReadPutForbidden {
	return &SMSInboundReadPutForbidden{}
}

/*SMSInboundReadPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundReadPutForbidden struct {
	Payload string
}

func (o *SMSInboundReadPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundReadPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutNotFound creates a SMSInboundReadPutNotFound with default headers values
func NewSMSInboundReadPutNotFound() *SMSInboundReadPutNotFound {
	return &SMSInboundReadPutNotFound{}
}

/*SMSInboundReadPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundReadPutNotFound struct {
	Payload string
}

func (o *SMSInboundReadPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundReadPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutMethodNotAllowed creates a SMSInboundReadPutMethodNotAllowed with default headers values
func NewSMSInboundReadPutMethodNotAllowed() *SMSInboundReadPutMethodNotAllowed {
	return &SMSInboundReadPutMethodNotAllowed{}
}

/*SMSInboundReadPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundReadPutMethodNotAllowed struct {
	Payload string
}

func (o *SMSInboundReadPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundReadPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutTooManyRequests creates a SMSInboundReadPutTooManyRequests with default headers values
func NewSMSInboundReadPutTooManyRequests() *SMSInboundReadPutTooManyRequests {
	return &SMSInboundReadPutTooManyRequests{}
}

/*SMSInboundReadPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundReadPutTooManyRequests struct {
	Payload string
}

func (o *SMSInboundReadPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] smsInboundReadPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundReadPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadPutDefault creates a SMSInboundReadPutDefault with default headers values
func NewSMSInboundReadPutDefault(code int) *SMSInboundReadPutDefault {
	return &SMSInboundReadPutDefault{
		_statusCode: code,
	}
}

/*SMSInboundReadPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundReadPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms inbound read put default response
func (o *SMSInboundReadPutDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundReadPutDefault) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read][%d] SmsInboundReadPut default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundReadPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
