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

// EmailDeliveryReceiptAutomationDeleteReader is a Reader for the EmailDeliveryReceiptAutomationDelete structure.
type EmailDeliveryReceiptAutomationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailDeliveryReceiptAutomationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailDeliveryReceiptAutomationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailDeliveryReceiptAutomationDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailDeliveryReceiptAutomationDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailDeliveryReceiptAutomationDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailDeliveryReceiptAutomationDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailDeliveryReceiptAutomationDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailDeliveryReceiptAutomationDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailDeliveryReceiptAutomationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailDeliveryReceiptAutomationDeleteOK creates a EmailDeliveryReceiptAutomationDeleteOK with default headers values
func NewEmailDeliveryReceiptAutomationDeleteOK() *EmailDeliveryReceiptAutomationDeleteOK {
	return &EmailDeliveryReceiptAutomationDeleteOK{}
}

/*EmailDeliveryReceiptAutomationDeleteOK handles this case with default header values.

SUCCESS
*/
type EmailDeliveryReceiptAutomationDeleteOK struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteOK  %+v", 200, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteBadRequest creates a EmailDeliveryReceiptAutomationDeleteBadRequest with default headers values
func NewEmailDeliveryReceiptAutomationDeleteBadRequest() *EmailDeliveryReceiptAutomationDeleteBadRequest {
	return &EmailDeliveryReceiptAutomationDeleteBadRequest{}
}

/*EmailDeliveryReceiptAutomationDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailDeliveryReceiptAutomationDeleteBadRequest struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteUnauthorized creates a EmailDeliveryReceiptAutomationDeleteUnauthorized with default headers values
func NewEmailDeliveryReceiptAutomationDeleteUnauthorized() *EmailDeliveryReceiptAutomationDeleteUnauthorized {
	return &EmailDeliveryReceiptAutomationDeleteUnauthorized{}
}

/*EmailDeliveryReceiptAutomationDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailDeliveryReceiptAutomationDeleteUnauthorized struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteForbidden creates a EmailDeliveryReceiptAutomationDeleteForbidden with default headers values
func NewEmailDeliveryReceiptAutomationDeleteForbidden() *EmailDeliveryReceiptAutomationDeleteForbidden {
	return &EmailDeliveryReceiptAutomationDeleteForbidden{}
}

/*EmailDeliveryReceiptAutomationDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailDeliveryReceiptAutomationDeleteForbidden struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteForbidden  %+v", 403, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteNotFound creates a EmailDeliveryReceiptAutomationDeleteNotFound with default headers values
func NewEmailDeliveryReceiptAutomationDeleteNotFound() *EmailDeliveryReceiptAutomationDeleteNotFound {
	return &EmailDeliveryReceiptAutomationDeleteNotFound{}
}

/*EmailDeliveryReceiptAutomationDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailDeliveryReceiptAutomationDeleteNotFound struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteNotFound  %+v", 404, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteMethodNotAllowed creates a EmailDeliveryReceiptAutomationDeleteMethodNotAllowed with default headers values
func NewEmailDeliveryReceiptAutomationDeleteMethodNotAllowed() *EmailDeliveryReceiptAutomationDeleteMethodNotAllowed {
	return &EmailDeliveryReceiptAutomationDeleteMethodNotAllowed{}
}

/*EmailDeliveryReceiptAutomationDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailDeliveryReceiptAutomationDeleteMethodNotAllowed struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteTooManyRequests creates a EmailDeliveryReceiptAutomationDeleteTooManyRequests with default headers values
func NewEmailDeliveryReceiptAutomationDeleteTooManyRequests() *EmailDeliveryReceiptAutomationDeleteTooManyRequests {
	return &EmailDeliveryReceiptAutomationDeleteTooManyRequests{}
}

/*EmailDeliveryReceiptAutomationDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailDeliveryReceiptAutomationDeleteTooManyRequests struct {
	Payload string
}

func (o *EmailDeliveryReceiptAutomationDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] emailDeliveryReceiptAutomationDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationDeleteDefault creates a EmailDeliveryReceiptAutomationDeleteDefault with default headers values
func NewEmailDeliveryReceiptAutomationDeleteDefault(code int) *EmailDeliveryReceiptAutomationDeleteDefault {
	return &EmailDeliveryReceiptAutomationDeleteDefault{
		_statusCode: code,
	}
}

/*EmailDeliveryReceiptAutomationDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailDeliveryReceiptAutomationDeleteDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the email delivery receipt automation delete default response
func (o *EmailDeliveryReceiptAutomationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EmailDeliveryReceiptAutomationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /automations/email/receipts/{receipt_rule_id}][%d] EmailDeliveryReceiptAutomationDelete default  %+v", o._statusCode, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}