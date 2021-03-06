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

// SMSDeliveryReceiptAutomationPutReader is a Reader for the SMSDeliveryReceiptAutomationPut structure.
type SMSDeliveryReceiptAutomationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSDeliveryReceiptAutomationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSDeliveryReceiptAutomationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSDeliveryReceiptAutomationPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSDeliveryReceiptAutomationPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSDeliveryReceiptAutomationPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSDeliveryReceiptAutomationPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSDeliveryReceiptAutomationPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSDeliveryReceiptAutomationPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSDeliveryReceiptAutomationPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSDeliveryReceiptAutomationPutOK creates a SMSDeliveryReceiptAutomationPutOK with default headers values
func NewSMSDeliveryReceiptAutomationPutOK() *SMSDeliveryReceiptAutomationPutOK {
	return &SMSDeliveryReceiptAutomationPutOK{}
}

/*SMSDeliveryReceiptAutomationPutOK handles this case with default header values.

SUCCESS
*/
type SMSDeliveryReceiptAutomationPutOK struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutOK) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutOK  %+v", 200, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutBadRequest creates a SMSDeliveryReceiptAutomationPutBadRequest with default headers values
func NewSMSDeliveryReceiptAutomationPutBadRequest() *SMSDeliveryReceiptAutomationPutBadRequest {
	return &SMSDeliveryReceiptAutomationPutBadRequest{}
}

/*SMSDeliveryReceiptAutomationPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSDeliveryReceiptAutomationPutBadRequest struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutBadRequest  %+v", 400, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutUnauthorized creates a SMSDeliveryReceiptAutomationPutUnauthorized with default headers values
func NewSMSDeliveryReceiptAutomationPutUnauthorized() *SMSDeliveryReceiptAutomationPutUnauthorized {
	return &SMSDeliveryReceiptAutomationPutUnauthorized{}
}

/*SMSDeliveryReceiptAutomationPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSDeliveryReceiptAutomationPutUnauthorized struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutForbidden creates a SMSDeliveryReceiptAutomationPutForbidden with default headers values
func NewSMSDeliveryReceiptAutomationPutForbidden() *SMSDeliveryReceiptAutomationPutForbidden {
	return &SMSDeliveryReceiptAutomationPutForbidden{}
}

/*SMSDeliveryReceiptAutomationPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSDeliveryReceiptAutomationPutForbidden struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutForbidden  %+v", 403, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutNotFound creates a SMSDeliveryReceiptAutomationPutNotFound with default headers values
func NewSMSDeliveryReceiptAutomationPutNotFound() *SMSDeliveryReceiptAutomationPutNotFound {
	return &SMSDeliveryReceiptAutomationPutNotFound{}
}

/*SMSDeliveryReceiptAutomationPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSDeliveryReceiptAutomationPutNotFound struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutNotFound  %+v", 404, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutMethodNotAllowed creates a SMSDeliveryReceiptAutomationPutMethodNotAllowed with default headers values
func NewSMSDeliveryReceiptAutomationPutMethodNotAllowed() *SMSDeliveryReceiptAutomationPutMethodNotAllowed {
	return &SMSDeliveryReceiptAutomationPutMethodNotAllowed{}
}

/*SMSDeliveryReceiptAutomationPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSDeliveryReceiptAutomationPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutTooManyRequests creates a SMSDeliveryReceiptAutomationPutTooManyRequests with default headers values
func NewSMSDeliveryReceiptAutomationPutTooManyRequests() *SMSDeliveryReceiptAutomationPutTooManyRequests {
	return &SMSDeliveryReceiptAutomationPutTooManyRequests{}
}

/*SMSDeliveryReceiptAutomationPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSDeliveryReceiptAutomationPutTooManyRequests struct {
	Payload interface{}
}

func (o *SMSDeliveryReceiptAutomationPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] smsDeliveryReceiptAutomationPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSDeliveryReceiptAutomationPutDefault creates a SMSDeliveryReceiptAutomationPutDefault with default headers values
func NewSMSDeliveryReceiptAutomationPutDefault(code int) *SMSDeliveryReceiptAutomationPutDefault {
	return &SMSDeliveryReceiptAutomationPutDefault{
		_statusCode: code,
	}
}

/*SMSDeliveryReceiptAutomationPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSDeliveryReceiptAutomationPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms delivery receipt automation put default response
func (o *SMSDeliveryReceiptAutomationPutDefault) Code() int {
	return o._statusCode
}

func (o *SMSDeliveryReceiptAutomationPutDefault) Error() string {
	return fmt.Sprintf("[PUT /automations/sms/receipts/{receipt_rule_id}][%d] SmsDeliveryReceiptAutomationPut default  %+v", o._statusCode, o.Payload)
}

func (o *SMSDeliveryReceiptAutomationPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
