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

// EmailDeliveryReceiptAutomationsGetReader is a Reader for the EmailDeliveryReceiptAutomationsGet structure.
type EmailDeliveryReceiptAutomationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailDeliveryReceiptAutomationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailDeliveryReceiptAutomationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailDeliveryReceiptAutomationsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailDeliveryReceiptAutomationsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailDeliveryReceiptAutomationsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailDeliveryReceiptAutomationsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailDeliveryReceiptAutomationsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailDeliveryReceiptAutomationsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailDeliveryReceiptAutomationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailDeliveryReceiptAutomationsGetOK creates a EmailDeliveryReceiptAutomationsGetOK with default headers values
func NewEmailDeliveryReceiptAutomationsGetOK() *EmailDeliveryReceiptAutomationsGetOK {
	return &EmailDeliveryReceiptAutomationsGetOK{}
}

/*EmailDeliveryReceiptAutomationsGetOK handles this case with default header values.

SUCCESS
*/
type EmailDeliveryReceiptAutomationsGetOK struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetOK) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetOK  %+v", 200, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetBadRequest creates a EmailDeliveryReceiptAutomationsGetBadRequest with default headers values
func NewEmailDeliveryReceiptAutomationsGetBadRequest() *EmailDeliveryReceiptAutomationsGetBadRequest {
	return &EmailDeliveryReceiptAutomationsGetBadRequest{}
}

/*EmailDeliveryReceiptAutomationsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailDeliveryReceiptAutomationsGetBadRequest struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetBadRequest  %+v", 400, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetUnauthorized creates a EmailDeliveryReceiptAutomationsGetUnauthorized with default headers values
func NewEmailDeliveryReceiptAutomationsGetUnauthorized() *EmailDeliveryReceiptAutomationsGetUnauthorized {
	return &EmailDeliveryReceiptAutomationsGetUnauthorized{}
}

/*EmailDeliveryReceiptAutomationsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailDeliveryReceiptAutomationsGetUnauthorized struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetForbidden creates a EmailDeliveryReceiptAutomationsGetForbidden with default headers values
func NewEmailDeliveryReceiptAutomationsGetForbidden() *EmailDeliveryReceiptAutomationsGetForbidden {
	return &EmailDeliveryReceiptAutomationsGetForbidden{}
}

/*EmailDeliveryReceiptAutomationsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailDeliveryReceiptAutomationsGetForbidden struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetForbidden  %+v", 403, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetNotFound creates a EmailDeliveryReceiptAutomationsGetNotFound with default headers values
func NewEmailDeliveryReceiptAutomationsGetNotFound() *EmailDeliveryReceiptAutomationsGetNotFound {
	return &EmailDeliveryReceiptAutomationsGetNotFound{}
}

/*EmailDeliveryReceiptAutomationsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailDeliveryReceiptAutomationsGetNotFound struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetNotFound  %+v", 404, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetMethodNotAllowed creates a EmailDeliveryReceiptAutomationsGetMethodNotAllowed with default headers values
func NewEmailDeliveryReceiptAutomationsGetMethodNotAllowed() *EmailDeliveryReceiptAutomationsGetMethodNotAllowed {
	return &EmailDeliveryReceiptAutomationsGetMethodNotAllowed{}
}

/*EmailDeliveryReceiptAutomationsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailDeliveryReceiptAutomationsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetTooManyRequests creates a EmailDeliveryReceiptAutomationsGetTooManyRequests with default headers values
func NewEmailDeliveryReceiptAutomationsGetTooManyRequests() *EmailDeliveryReceiptAutomationsGetTooManyRequests {
	return &EmailDeliveryReceiptAutomationsGetTooManyRequests{}
}

/*EmailDeliveryReceiptAutomationsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailDeliveryReceiptAutomationsGetTooManyRequests struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] emailDeliveryReceiptAutomationsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationsGetDefault creates a EmailDeliveryReceiptAutomationsGetDefault with default headers values
func NewEmailDeliveryReceiptAutomationsGetDefault(code int) *EmailDeliveryReceiptAutomationsGetDefault {
	return &EmailDeliveryReceiptAutomationsGetDefault{
		_statusCode: code,
	}
}

/*EmailDeliveryReceiptAutomationsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailDeliveryReceiptAutomationsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the email delivery receipt automations get default response
func (o *EmailDeliveryReceiptAutomationsGetDefault) Code() int {
	return o._statusCode
}

func (o *EmailDeliveryReceiptAutomationsGetDefault) Error() string {
	return fmt.Sprintf("[GET /automations/email/receipts][%d] EmailDeliveryReceiptAutomationsGet default  %+v", o._statusCode, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
