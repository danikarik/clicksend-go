// Code generated by go-swagger; DO NOT EDIT.

package email_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmailDeliveryReceiptAutomationPutReader is a Reader for the EmailDeliveryReceiptAutomationPut structure.
type EmailDeliveryReceiptAutomationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailDeliveryReceiptAutomationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailDeliveryReceiptAutomationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailDeliveryReceiptAutomationPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailDeliveryReceiptAutomationPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailDeliveryReceiptAutomationPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailDeliveryReceiptAutomationPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailDeliveryReceiptAutomationPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailDeliveryReceiptAutomationPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailDeliveryReceiptAutomationPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailDeliveryReceiptAutomationPutOK creates a EmailDeliveryReceiptAutomationPutOK with default headers values
func NewEmailDeliveryReceiptAutomationPutOK() *EmailDeliveryReceiptAutomationPutOK {
	return &EmailDeliveryReceiptAutomationPutOK{}
}

/*EmailDeliveryReceiptAutomationPutOK handles this case with default header values.

SUCCESS
*/
type EmailDeliveryReceiptAutomationPutOK struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutOK) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutOK  %+v", 200, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutBadRequest creates a EmailDeliveryReceiptAutomationPutBadRequest with default headers values
func NewEmailDeliveryReceiptAutomationPutBadRequest() *EmailDeliveryReceiptAutomationPutBadRequest {
	return &EmailDeliveryReceiptAutomationPutBadRequest{}
}

/*EmailDeliveryReceiptAutomationPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailDeliveryReceiptAutomationPutBadRequest struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutBadRequest  %+v", 400, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutUnauthorized creates a EmailDeliveryReceiptAutomationPutUnauthorized with default headers values
func NewEmailDeliveryReceiptAutomationPutUnauthorized() *EmailDeliveryReceiptAutomationPutUnauthorized {
	return &EmailDeliveryReceiptAutomationPutUnauthorized{}
}

/*EmailDeliveryReceiptAutomationPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailDeliveryReceiptAutomationPutUnauthorized struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutForbidden creates a EmailDeliveryReceiptAutomationPutForbidden with default headers values
func NewEmailDeliveryReceiptAutomationPutForbidden() *EmailDeliveryReceiptAutomationPutForbidden {
	return &EmailDeliveryReceiptAutomationPutForbidden{}
}

/*EmailDeliveryReceiptAutomationPutForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailDeliveryReceiptAutomationPutForbidden struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutForbidden  %+v", 403, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutNotFound creates a EmailDeliveryReceiptAutomationPutNotFound with default headers values
func NewEmailDeliveryReceiptAutomationPutNotFound() *EmailDeliveryReceiptAutomationPutNotFound {
	return &EmailDeliveryReceiptAutomationPutNotFound{}
}

/*EmailDeliveryReceiptAutomationPutNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailDeliveryReceiptAutomationPutNotFound struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutNotFound  %+v", 404, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutMethodNotAllowed creates a EmailDeliveryReceiptAutomationPutMethodNotAllowed with default headers values
func NewEmailDeliveryReceiptAutomationPutMethodNotAllowed() *EmailDeliveryReceiptAutomationPutMethodNotAllowed {
	return &EmailDeliveryReceiptAutomationPutMethodNotAllowed{}
}

/*EmailDeliveryReceiptAutomationPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailDeliveryReceiptAutomationPutMethodNotAllowed struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutTooManyRequests creates a EmailDeliveryReceiptAutomationPutTooManyRequests with default headers values
func NewEmailDeliveryReceiptAutomationPutTooManyRequests() *EmailDeliveryReceiptAutomationPutTooManyRequests {
	return &EmailDeliveryReceiptAutomationPutTooManyRequests{}
}

/*EmailDeliveryReceiptAutomationPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailDeliveryReceiptAutomationPutTooManyRequests struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPutDefault creates a EmailDeliveryReceiptAutomationPutDefault with default headers values
func NewEmailDeliveryReceiptAutomationPutDefault(code int) *EmailDeliveryReceiptAutomationPutDefault {
	return &EmailDeliveryReceiptAutomationPutDefault{
		_statusCode: code,
	}
}

/*EmailDeliveryReceiptAutomationPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailDeliveryReceiptAutomationPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the email delivery receipt automation put default response
func (o *EmailDeliveryReceiptAutomationPutDefault) Code() int {
	return o._statusCode
}

func (o *EmailDeliveryReceiptAutomationPutDefault) Error() string {
	return fmt.Sprintf("[PUT /automations/email/receipts/{receipt_rule_id}][%d] EmailDeliveryReceiptAutomationPut default  %+v", o._statusCode, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
