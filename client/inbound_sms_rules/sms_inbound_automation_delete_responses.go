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

// SMSInboundAutomationDeleteReader is a Reader for the SMSInboundAutomationDelete structure.
type SMSInboundAutomationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundAutomationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundAutomationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundAutomationDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundAutomationDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundAutomationDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundAutomationDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundAutomationDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundAutomationDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundAutomationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundAutomationDeleteOK creates a SMSInboundAutomationDeleteOK with default headers values
func NewSMSInboundAutomationDeleteOK() *SMSInboundAutomationDeleteOK {
	return &SMSInboundAutomationDeleteOK{}
}

/*SMSInboundAutomationDeleteOK handles this case with default header values.

SUCCESS
*/
type SMSInboundAutomationDeleteOK struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteOK  %+v", 200, o.Payload)
}

func (o *SMSInboundAutomationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteBadRequest creates a SMSInboundAutomationDeleteBadRequest with default headers values
func NewSMSInboundAutomationDeleteBadRequest() *SMSInboundAutomationDeleteBadRequest {
	return &SMSInboundAutomationDeleteBadRequest{}
}

/*SMSInboundAutomationDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundAutomationDeleteBadRequest struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundAutomationDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteUnauthorized creates a SMSInboundAutomationDeleteUnauthorized with default headers values
func NewSMSInboundAutomationDeleteUnauthorized() *SMSInboundAutomationDeleteUnauthorized {
	return &SMSInboundAutomationDeleteUnauthorized{}
}

/*SMSInboundAutomationDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundAutomationDeleteUnauthorized struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundAutomationDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteForbidden creates a SMSInboundAutomationDeleteForbidden with default headers values
func NewSMSInboundAutomationDeleteForbidden() *SMSInboundAutomationDeleteForbidden {
	return &SMSInboundAutomationDeleteForbidden{}
}

/*SMSInboundAutomationDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundAutomationDeleteForbidden struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundAutomationDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteNotFound creates a SMSInboundAutomationDeleteNotFound with default headers values
func NewSMSInboundAutomationDeleteNotFound() *SMSInboundAutomationDeleteNotFound {
	return &SMSInboundAutomationDeleteNotFound{}
}

/*SMSInboundAutomationDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundAutomationDeleteNotFound struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundAutomationDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteMethodNotAllowed creates a SMSInboundAutomationDeleteMethodNotAllowed with default headers values
func NewSMSInboundAutomationDeleteMethodNotAllowed() *SMSInboundAutomationDeleteMethodNotAllowed {
	return &SMSInboundAutomationDeleteMethodNotAllowed{}
}

/*SMSInboundAutomationDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundAutomationDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundAutomationDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteTooManyRequests creates a SMSInboundAutomationDeleteTooManyRequests with default headers values
func NewSMSInboundAutomationDeleteTooManyRequests() *SMSInboundAutomationDeleteTooManyRequests {
	return &SMSInboundAutomationDeleteTooManyRequests{}
}

/*SMSInboundAutomationDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundAutomationDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *SMSInboundAutomationDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] smsInboundAutomationDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundAutomationDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationDeleteDefault creates a SMSInboundAutomationDeleteDefault with default headers values
func NewSMSInboundAutomationDeleteDefault(code int) *SMSInboundAutomationDeleteDefault {
	return &SMSInboundAutomationDeleteDefault{
		_statusCode: code,
	}
}

/*SMSInboundAutomationDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundAutomationDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms inbound automation delete default response
func (o *SMSInboundAutomationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundAutomationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/inbound/{inbound_rule_id}][%d] SmsInboundAutomationDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundAutomationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
