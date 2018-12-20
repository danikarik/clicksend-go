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

// SMSInboundAutomationPostReader is a Reader for the SMSInboundAutomationPost structure.
type SMSInboundAutomationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSInboundAutomationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSInboundAutomationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSInboundAutomationPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSInboundAutomationPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSInboundAutomationPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSInboundAutomationPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSInboundAutomationPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSInboundAutomationPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSInboundAutomationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSInboundAutomationPostOK creates a SMSInboundAutomationPostOK with default headers values
func NewSMSInboundAutomationPostOK() *SMSInboundAutomationPostOK {
	return &SMSInboundAutomationPostOK{}
}

/*SMSInboundAutomationPostOK handles this case with default header values.

SUCCESS
*/
type SMSInboundAutomationPostOK struct {
	Payload string
}

func (o *SMSInboundAutomationPostOK) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostOK  %+v", 200, o.Payload)
}

func (o *SMSInboundAutomationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostBadRequest creates a SMSInboundAutomationPostBadRequest with default headers values
func NewSMSInboundAutomationPostBadRequest() *SMSInboundAutomationPostBadRequest {
	return &SMSInboundAutomationPostBadRequest{}
}

/*SMSInboundAutomationPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSInboundAutomationPostBadRequest struct {
	Payload string
}

func (o *SMSInboundAutomationPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostBadRequest  %+v", 400, o.Payload)
}

func (o *SMSInboundAutomationPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostUnauthorized creates a SMSInboundAutomationPostUnauthorized with default headers values
func NewSMSInboundAutomationPostUnauthorized() *SMSInboundAutomationPostUnauthorized {
	return &SMSInboundAutomationPostUnauthorized{}
}

/*SMSInboundAutomationPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSInboundAutomationPostUnauthorized struct {
	Payload string
}

func (o *SMSInboundAutomationPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSInboundAutomationPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostForbidden creates a SMSInboundAutomationPostForbidden with default headers values
func NewSMSInboundAutomationPostForbidden() *SMSInboundAutomationPostForbidden {
	return &SMSInboundAutomationPostForbidden{}
}

/*SMSInboundAutomationPostForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSInboundAutomationPostForbidden struct {
	Payload string
}

func (o *SMSInboundAutomationPostForbidden) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostForbidden  %+v", 403, o.Payload)
}

func (o *SMSInboundAutomationPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostNotFound creates a SMSInboundAutomationPostNotFound with default headers values
func NewSMSInboundAutomationPostNotFound() *SMSInboundAutomationPostNotFound {
	return &SMSInboundAutomationPostNotFound{}
}

/*SMSInboundAutomationPostNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSInboundAutomationPostNotFound struct {
	Payload string
}

func (o *SMSInboundAutomationPostNotFound) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostNotFound  %+v", 404, o.Payload)
}

func (o *SMSInboundAutomationPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostMethodNotAllowed creates a SMSInboundAutomationPostMethodNotAllowed with default headers values
func NewSMSInboundAutomationPostMethodNotAllowed() *SMSInboundAutomationPostMethodNotAllowed {
	return &SMSInboundAutomationPostMethodNotAllowed{}
}

/*SMSInboundAutomationPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSInboundAutomationPostMethodNotAllowed struct {
	Payload string
}

func (o *SMSInboundAutomationPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSInboundAutomationPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostTooManyRequests creates a SMSInboundAutomationPostTooManyRequests with default headers values
func NewSMSInboundAutomationPostTooManyRequests() *SMSInboundAutomationPostTooManyRequests {
	return &SMSInboundAutomationPostTooManyRequests{}
}

/*SMSInboundAutomationPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSInboundAutomationPostTooManyRequests struct {
	Payload string
}

func (o *SMSInboundAutomationPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] smsInboundAutomationPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSInboundAutomationPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSInboundAutomationPostDefault creates a SMSInboundAutomationPostDefault with default headers values
func NewSMSInboundAutomationPostDefault(code int) *SMSInboundAutomationPostDefault {
	return &SMSInboundAutomationPostDefault{
		_statusCode: code,
	}
}

/*SMSInboundAutomationPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSInboundAutomationPostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms inbound automation post default response
func (o *SMSInboundAutomationPostDefault) Code() int {
	return o._statusCode
}

func (o *SMSInboundAutomationPostDefault) Error() string {
	return fmt.Sprintf("[POST /automations/sms/inbound][%d] SmsInboundAutomationPost default  %+v", o._statusCode, o.Payload)
}

func (o *SMSInboundAutomationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
