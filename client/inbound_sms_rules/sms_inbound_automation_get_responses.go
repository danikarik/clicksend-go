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

// SMSInboundAutomationGetReader is a Reader for the SMSInboundAutomationGet structure.
type SMSInboundAutomationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundAutomationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundAutomationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundAutomationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundAutomationGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundAutomationGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundAutomationGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundAutomationGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundAutomationGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundAutomationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundAutomationGetOK creates a SMSInboundAutomationGetOK with default headers values
func NewSMSInboundAutomationGetOK() *SMSInboundAutomationGetOK {
	return &SMSInboundAutomationGetOK{}
}

/*SMSInboundAutomationGetOK handles this case with default header values.

SUCCESS
*/
type SMSInboundAutomationGetOK struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetOK) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetOK  %+v", 200, o.Payload)
}

func (o *SMSInboundAutomationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetBadRequest creates a SMSInboundAutomationGetBadRequest with default headers values
func NewSMSInboundAutomationGetBadRequest() *SMSInboundAutomationGetBadRequest {
	return &SMSInboundAutomationGetBadRequest{}
}

/*SMSInboundAutomationGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundAutomationGetBadRequest struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundAutomationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetUnauthorized creates a SMSInboundAutomationGetUnauthorized with default headers values
func NewSMSInboundAutomationGetUnauthorized() *SMSInboundAutomationGetUnauthorized {
	return &SMSInboundAutomationGetUnauthorized{}
}

/*SMSInboundAutomationGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundAutomationGetUnauthorized struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundAutomationGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetForbidden creates a SMSInboundAutomationGetForbidden with default headers values
func NewSMSInboundAutomationGetForbidden() *SMSInboundAutomationGetForbidden {
	return &SMSInboundAutomationGetForbidden{}
}

/*SMSInboundAutomationGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundAutomationGetForbidden struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetForbidden) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundAutomationGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetNotFound creates a SMSInboundAutomationGetNotFound with default headers values
func NewSMSInboundAutomationGetNotFound() *SMSInboundAutomationGetNotFound {
	return &SMSInboundAutomationGetNotFound{}
}

/*SMSInboundAutomationGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundAutomationGetNotFound struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetNotFound) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundAutomationGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetMethodNotAllowed creates a SMSInboundAutomationGetMethodNotAllowed with default headers values
func NewSMSInboundAutomationGetMethodNotAllowed() *SMSInboundAutomationGetMethodNotAllowed {
	return &SMSInboundAutomationGetMethodNotAllowed{}
}

/*SMSInboundAutomationGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundAutomationGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundAutomationGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetTooManyRequests creates a SMSInboundAutomationGetTooManyRequests with default headers values
func NewSMSInboundAutomationGetTooManyRequests() *SMSInboundAutomationGetTooManyRequests {
	return &SMSInboundAutomationGetTooManyRequests{}
}

/*SMSInboundAutomationGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundAutomationGetTooManyRequests struct {
	Payload interface{}
}

func (o *SMSInboundAutomationGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundAutomationGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationGetDefault creates a SMSInboundAutomationGetDefault with default headers values
func NewSMSInboundAutomationGetDefault(code int) *SMSInboundAutomationGetDefault {
	return &SMSInboundAutomationGetDefault{
		_statusCode: code,
	}
}

/*SMSInboundAutomationGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundAutomationGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms inbound automation get default response
func (o *SMSInboundAutomationGetDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundAutomationGetDefault) Error() string {
	return fmt.Sprintf("[GET /automations/sms/inbound/{inbound_rule_id}][%d] SmsInboundAutomationGet default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundAutomationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
