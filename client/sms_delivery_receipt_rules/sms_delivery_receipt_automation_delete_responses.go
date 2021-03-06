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

// SMSDeliveryReceiptAutomationDeleteReader is a Reader for the SMSDeliveryReceiptAutomationDelete structure.
type SMSDeliveryReceiptAutomationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSDeliveryReceiptAutomationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSDeliveryReceiptAutomationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSDeliveryReceiptAutomationDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSDeliveryReceiptAutomationDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSDeliveryReceiptAutomationDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSDeliveryReceiptAutomationDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSDeliveryReceiptAutomationDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSDeliveryReceiptAutomationDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSDeliveryReceiptAutomationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSDeliveryReceiptAutomationDeleteOK creates a SMSDeliveryReceiptAutomationDeleteOK with default headers values
func NewSMSDeliveryReceiptAutomationDeleteOK() *SMSDeliveryReceiptAutomationDeleteOK {
	return &SMSDeliveryReceiptAutomationDeleteOK{}
}

/*SMSDeliveryReceiptAutomationDeleteOK handles this case with default header values.

SUCCESS
*/
type SMSDeliveryReceiptAutomationDeleteOK struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteOK  %+v", 200, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteBadRequest creates a SMSDeliveryReceiptAutomationDeleteBadRequest with default headers values
func NewSMSDeliveryReceiptAutomationDeleteBadRequest() *SMSDeliveryReceiptAutomationDeleteBadRequest {
	return &SMSDeliveryReceiptAutomationDeleteBadRequest{}
}

/*SMSDeliveryReceiptAutomationDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSDeliveryReceiptAutomationDeleteBadRequest struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteUnauthorized creates a SMSDeliveryReceiptAutomationDeleteUnauthorized with default headers values
func NewSMSDeliveryReceiptAutomationDeleteUnauthorized() *SMSDeliveryReceiptAutomationDeleteUnauthorized {
	return &SMSDeliveryReceiptAutomationDeleteUnauthorized{}
}

/*SMSDeliveryReceiptAutomationDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSDeliveryReceiptAutomationDeleteUnauthorized struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteForbidden creates a SMSDeliveryReceiptAutomationDeleteForbidden with default headers values
func NewSMSDeliveryReceiptAutomationDeleteForbidden() *SMSDeliveryReceiptAutomationDeleteForbidden {
	return &SMSDeliveryReceiptAutomationDeleteForbidden{}
}

/*SMSDeliveryReceiptAutomationDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSDeliveryReceiptAutomationDeleteForbidden struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteNotFound creates a SMSDeliveryReceiptAutomationDeleteNotFound with default headers values
func NewSMSDeliveryReceiptAutomationDeleteNotFound() *SMSDeliveryReceiptAutomationDeleteNotFound {
	return &SMSDeliveryReceiptAutomationDeleteNotFound{}
}

/*SMSDeliveryReceiptAutomationDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSDeliveryReceiptAutomationDeleteNotFound struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteMethodNotAllowed creates a SMSDeliveryReceiptAutomationDeleteMethodNotAllowed with default headers values
func NewSMSDeliveryReceiptAutomationDeleteMethodNotAllowed() *SMSDeliveryReceiptAutomationDeleteMethodNotAllowed {
	return &SMSDeliveryReceiptAutomationDeleteMethodNotAllowed{}
}

/*SMSDeliveryReceiptAutomationDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSDeliveryReceiptAutomationDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteTooManyRequests creates a SMSDeliveryReceiptAutomationDeleteTooManyRequests with default headers values
func NewSMSDeliveryReceiptAutomationDeleteTooManyRequests() *SMSDeliveryReceiptAutomationDeleteTooManyRequests {
	return &SMSDeliveryReceiptAutomationDeleteTooManyRequests{}
}

/*SMSDeliveryReceiptAutomationDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSDeliveryReceiptAutomationDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationDeleteDefault creates a SMSDeliveryReceiptAutomationDeleteDefault with default headers values
func NewSMSDeliveryReceiptAutomationDeleteDefault(code int) *SMSDeliveryReceiptAutomationDeleteDefault {
	return &SMSDeliveryReceiptAutomationDeleteDefault{
		_statusCode: code,
	}
}

/*SMSDeliveryReceiptAutomationDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSDeliveryReceiptAutomationDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms delivery receipt automation delete default response
func (o *SMSDeliveryReceiptAutomationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SMSDeliveryReceiptAutomationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /automations/sms/receipts/{receipt_rule_id}][%d] SmsDeliveryReceiptAutomationDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
