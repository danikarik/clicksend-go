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

// SMSInboundReadByMessageIDPutReader is a Reader for the SMSInboundReadByMessageIDPut structure.
type SMSInboundReadByMessageIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundReadByMessageIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundReadByMessageIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundReadByMessageIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundReadByMessageIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundReadByMessageIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundReadByMessageIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundReadByMessageIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundReadByMessageIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundReadByMessageIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundReadByMessageIDPutOK creates a SMSInboundReadByMessageIDPutOK with default headers values
func NewSMSInboundReadByMessageIDPutOK() *SMSInboundReadByMessageIDPutOK {
	return &SMSInboundReadByMessageIDPutOK{}
}

/*SMSInboundReadByMessageIDPutOK handles this case with default header values.

SUCCESS
*/
type SMSInboundReadByMessageIDPutOK struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutOK  %+v", 200, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutBadRequest creates a SMSInboundReadByMessageIDPutBadRequest with default headers values
func NewSMSInboundReadByMessageIDPutBadRequest() *SMSInboundReadByMessageIDPutBadRequest {
	return &SMSInboundReadByMessageIDPutBadRequest{}
}

/*SMSInboundReadByMessageIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundReadByMessageIDPutBadRequest struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutUnauthorized creates a SMSInboundReadByMessageIDPutUnauthorized with default headers values
func NewSMSInboundReadByMessageIDPutUnauthorized() *SMSInboundReadByMessageIDPutUnauthorized {
	return &SMSInboundReadByMessageIDPutUnauthorized{}
}

/*SMSInboundReadByMessageIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundReadByMessageIDPutUnauthorized struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutForbidden creates a SMSInboundReadByMessageIDPutForbidden with default headers values
func NewSMSInboundReadByMessageIDPutForbidden() *SMSInboundReadByMessageIDPutForbidden {
	return &SMSInboundReadByMessageIDPutForbidden{}
}

/*SMSInboundReadByMessageIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundReadByMessageIDPutForbidden struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutNotFound creates a SMSInboundReadByMessageIDPutNotFound with default headers values
func NewSMSInboundReadByMessageIDPutNotFound() *SMSInboundReadByMessageIDPutNotFound {
	return &SMSInboundReadByMessageIDPutNotFound{}
}

/*SMSInboundReadByMessageIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundReadByMessageIDPutNotFound struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutMethodNotAllowed creates a SMSInboundReadByMessageIDPutMethodNotAllowed with default headers values
func NewSMSInboundReadByMessageIDPutMethodNotAllowed() *SMSInboundReadByMessageIDPutMethodNotAllowed {
	return &SMSInboundReadByMessageIDPutMethodNotAllowed{}
}

/*SMSInboundReadByMessageIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundReadByMessageIDPutMethodNotAllowed struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutTooManyRequests creates a SMSInboundReadByMessageIDPutTooManyRequests with default headers values
func NewSMSInboundReadByMessageIDPutTooManyRequests() *SMSInboundReadByMessageIDPutTooManyRequests {
	return &SMSInboundReadByMessageIDPutTooManyRequests{}
}

/*SMSInboundReadByMessageIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundReadByMessageIDPutTooManyRequests struct {
	Payload string
}

func (o *SMSInboundReadByMessageIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] smsInboundReadByMessageIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundReadByMessageIDPutDefault creates a SMSInboundReadByMessageIDPutDefault with default headers values
func NewSMSInboundReadByMessageIDPutDefault(code int) *SMSInboundReadByMessageIDPutDefault {
	return &SMSInboundReadByMessageIDPutDefault{
		_statusCode: code,
	}
}

/*SMSInboundReadByMessageIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundReadByMessageIDPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms inbound read by message Id put default response
func (o *SMSInboundReadByMessageIDPutDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundReadByMessageIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /sms/inbound-read/{message_id}][%d] SmsInboundReadByMessageIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundReadByMessageIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
