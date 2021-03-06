// Code generated by go-swagger; DO NOT EDIT.

package sms_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSDeliveryReceiptAutomationsGetReader is a Reader for the SMSDeliveryReceiptAutomationsGet structure.
type SMSDeliveryReceiptAutomationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSDeliveryReceiptAutomationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSDeliveryReceiptAutomationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSDeliveryReceiptAutomationsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSDeliveryReceiptAutomationsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSDeliveryReceiptAutomationsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSDeliveryReceiptAutomationsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSDeliveryReceiptAutomationsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSDeliveryReceiptAutomationsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSDeliveryReceiptAutomationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSDeliveryReceiptAutomationsGetOK creates a SMSDeliveryReceiptAutomationsGetOK with default headers values
func NewSMSDeliveryReceiptAutomationsGetOK() *SMSDeliveryReceiptAutomationsGetOK {
	return &SMSDeliveryReceiptAutomationsGetOK{}
}

/*SMSDeliveryReceiptAutomationsGetOK handles this case with default header values.

SUCCESS
*/
type SMSDeliveryReceiptAutomationsGetOK struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetOK) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetOK  %+v", 200, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetBadRequest creates a SMSDeliveryReceiptAutomationsGetBadRequest with default headers values
func NewSMSDeliveryReceiptAutomationsGetBadRequest() *SMSDeliveryReceiptAutomationsGetBadRequest {
	return &SMSDeliveryReceiptAutomationsGetBadRequest{}
}

/*SMSDeliveryReceiptAutomationsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSDeliveryReceiptAutomationsGetBadRequest struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetBadRequest  %+v", 400, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetUnauthorized creates a SMSDeliveryReceiptAutomationsGetUnauthorized with default headers values
func NewSMSDeliveryReceiptAutomationsGetUnauthorized() *SMSDeliveryReceiptAutomationsGetUnauthorized {
	return &SMSDeliveryReceiptAutomationsGetUnauthorized{}
}

/*SMSDeliveryReceiptAutomationsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSDeliveryReceiptAutomationsGetUnauthorized struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetForbidden creates a SMSDeliveryReceiptAutomationsGetForbidden with default headers values
func NewSMSDeliveryReceiptAutomationsGetForbidden() *SMSDeliveryReceiptAutomationsGetForbidden {
	return &SMSDeliveryReceiptAutomationsGetForbidden{}
}

/*SMSDeliveryReceiptAutomationsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSDeliveryReceiptAutomationsGetForbidden struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetForbidden  %+v", 403, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetNotFound creates a SMSDeliveryReceiptAutomationsGetNotFound with default headers values
func NewSMSDeliveryReceiptAutomationsGetNotFound() *SMSDeliveryReceiptAutomationsGetNotFound {
	return &SMSDeliveryReceiptAutomationsGetNotFound{}
}

/*SMSDeliveryReceiptAutomationsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSDeliveryReceiptAutomationsGetNotFound struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetNotFound  %+v", 404, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetMethodNotAllowed creates a SMSDeliveryReceiptAutomationsGetMethodNotAllowed with default headers values
func NewSMSDeliveryReceiptAutomationsGetMethodNotAllowed() *SMSDeliveryReceiptAutomationsGetMethodNotAllowed {
	return &SMSDeliveryReceiptAutomationsGetMethodNotAllowed{}
}

/*SMSDeliveryReceiptAutomationsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSDeliveryReceiptAutomationsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetTooManyRequests creates a SMSDeliveryReceiptAutomationsGetTooManyRequests with default headers values
func NewSMSDeliveryReceiptAutomationsGetTooManyRequests() *SMSDeliveryReceiptAutomationsGetTooManyRequests {
	return &SMSDeliveryReceiptAutomationsGetTooManyRequests{}
}

/*SMSDeliveryReceiptAutomationsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSDeliveryReceiptAutomationsGetTooManyRequests struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] smsDeliveryReceiptAutomationsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationsGetDefault creates a SMSDeliveryReceiptAutomationsGetDefault with default headers values
func NewSMSDeliveryReceiptAutomationsGetDefault(code int) *SMSDeliveryReceiptAutomationsGetDefault {
	return &SMSDeliveryReceiptAutomationsGetDefault{
		_statusCode: code,
	}
}

/*SMSDeliveryReceiptAutomationsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSDeliveryReceiptAutomationsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms delivery receipt automations get default response
func (o *SMSDeliveryReceiptAutomationsGetDefault) Code() int {
	return o._statusCode
}

func (o *SMSDeliveryReceiptAutomationsGetDefault) Error() string {
	return fmt.Sprintf("[GET /automations/sms/receipts][%d] SmsDeliveryReceiptAutomationsGet default  %+v", o._statusCode, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
