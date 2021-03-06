// Code generated by go-swagger; DO NOT EDIT.

package inbound_sms_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSInboundAutomationsGetReader is a Reader for the SMSInboundAutomationsGet structure.
type SMSInboundAutomationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundAutomationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundAutomationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundAutomationsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundAutomationsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundAutomationsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundAutomationsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundAutomationsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundAutomationsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundAutomationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundAutomationsGetOK creates a SMSInboundAutomationsGetOK with default headers values
func NewSMSInboundAutomationsGetOK() *SMSInboundAutomationsGetOK {
	return &SMSInboundAutomationsGetOK{}
}

/*SMSInboundAutomationsGetOK handles this case with default header values.

SUCCESS
*/
type SMSInboundAutomationsGetOK struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetOK) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetOK  %+v", 200, o.Payload)
}

func (o *SMSInboundAutomationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetBadRequest creates a SMSInboundAutomationsGetBadRequest with default headers values
func NewSMSInboundAutomationsGetBadRequest() *SMSInboundAutomationsGetBadRequest {
	return &SMSInboundAutomationsGetBadRequest{}
}

/*SMSInboundAutomationsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundAutomationsGetBadRequest struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundAutomationsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetUnauthorized creates a SMSInboundAutomationsGetUnauthorized with default headers values
func NewSMSInboundAutomationsGetUnauthorized() *SMSInboundAutomationsGetUnauthorized {
	return &SMSInboundAutomationsGetUnauthorized{}
}

/*SMSInboundAutomationsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundAutomationsGetUnauthorized struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundAutomationsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetForbidden creates a SMSInboundAutomationsGetForbidden with default headers values
func NewSMSInboundAutomationsGetForbidden() *SMSInboundAutomationsGetForbidden {
	return &SMSInboundAutomationsGetForbidden{}
}

/*SMSInboundAutomationsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundAutomationsGetForbidden struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundAutomationsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetNotFound creates a SMSInboundAutomationsGetNotFound with default headers values
func NewSMSInboundAutomationsGetNotFound() *SMSInboundAutomationsGetNotFound {
	return &SMSInboundAutomationsGetNotFound{}
}

/*SMSInboundAutomationsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundAutomationsGetNotFound struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundAutomationsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetMethodNotAllowed creates a SMSInboundAutomationsGetMethodNotAllowed with default headers values
func NewSMSInboundAutomationsGetMethodNotAllowed() *SMSInboundAutomationsGetMethodNotAllowed {
	return &SMSInboundAutomationsGetMethodNotAllowed{}
}

/*SMSInboundAutomationsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundAutomationsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundAutomationsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetTooManyRequests creates a SMSInboundAutomationsGetTooManyRequests with default headers values
func NewSMSInboundAutomationsGetTooManyRequests() *SMSInboundAutomationsGetTooManyRequests {
	return &SMSInboundAutomationsGetTooManyRequests{}
}

/*SMSInboundAutomationsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundAutomationsGetTooManyRequests struct {
	Payload interface{}
}

func (o *SMSInboundAutomationsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] smsInboundAutomationsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundAutomationsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationsGetDefault creates a SMSInboundAutomationsGetDefault with default headers values
func NewSMSInboundAutomationsGetDefault(code int) *SMSInboundAutomationsGetDefault {
	return &SMSInboundAutomationsGetDefault{
		_statusCode: code,
	}
}

/*SMSInboundAutomationsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundAutomationsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms inbound automations get default response
func (o *SMSInboundAutomationsGetDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundAutomationsGetDefault) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound][%d] SmsInboundAutomationsGet default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundAutomationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
