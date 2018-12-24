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

// EmailDeliveryReceiptAutomationPostReader is a Reader for the EmailDeliveryReceiptAutomationPost structure.
type EmailDeliveryReceiptAutomationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailDeliveryReceiptAutomationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailDeliveryReceiptAutomationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailDeliveryReceiptAutomationPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailDeliveryReceiptAutomationPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailDeliveryReceiptAutomationPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailDeliveryReceiptAutomationPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailDeliveryReceiptAutomationPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailDeliveryReceiptAutomationPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailDeliveryReceiptAutomationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailDeliveryReceiptAutomationPostOK creates a EmailDeliveryReceiptAutomationPostOK with default headers values
func NewEmailDeliveryReceiptAutomationPostOK() *EmailDeliveryReceiptAutomationPostOK {
	return &EmailDeliveryReceiptAutomationPostOK{}
}

/*EmailDeliveryReceiptAutomationPostOK handles this case with default header values.

SUCCESS
*/
type EmailDeliveryReceiptAutomationPostOK struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostOK) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostOK  %+v", 200, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostBadRequest creates a EmailDeliveryReceiptAutomationPostBadRequest with default headers values
func NewEmailDeliveryReceiptAutomationPostBadRequest() *EmailDeliveryReceiptAutomationPostBadRequest {
	return &EmailDeliveryReceiptAutomationPostBadRequest{}
}

/*EmailDeliveryReceiptAutomationPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailDeliveryReceiptAutomationPostBadRequest struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostBadRequest  %+v", 400, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostUnauthorized creates a EmailDeliveryReceiptAutomationPostUnauthorized with default headers values
func NewEmailDeliveryReceiptAutomationPostUnauthorized() *EmailDeliveryReceiptAutomationPostUnauthorized {
	return &EmailDeliveryReceiptAutomationPostUnauthorized{}
}

/*EmailDeliveryReceiptAutomationPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailDeliveryReceiptAutomationPostUnauthorized struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostForbidden creates a EmailDeliveryReceiptAutomationPostForbidden with default headers values
func NewEmailDeliveryReceiptAutomationPostForbidden() *EmailDeliveryReceiptAutomationPostForbidden {
	return &EmailDeliveryReceiptAutomationPostForbidden{}
}

/*EmailDeliveryReceiptAutomationPostForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailDeliveryReceiptAutomationPostForbidden struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostForbidden) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostForbidden  %+v", 403, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostNotFound creates a EmailDeliveryReceiptAutomationPostNotFound with default headers values
func NewEmailDeliveryReceiptAutomationPostNotFound() *EmailDeliveryReceiptAutomationPostNotFound {
	return &EmailDeliveryReceiptAutomationPostNotFound{}
}

/*EmailDeliveryReceiptAutomationPostNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailDeliveryReceiptAutomationPostNotFound struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostNotFound) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostNotFound  %+v", 404, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostMethodNotAllowed creates a EmailDeliveryReceiptAutomationPostMethodNotAllowed with default headers values
func NewEmailDeliveryReceiptAutomationPostMethodNotAllowed() *EmailDeliveryReceiptAutomationPostMethodNotAllowed {
	return &EmailDeliveryReceiptAutomationPostMethodNotAllowed{}
}

/*EmailDeliveryReceiptAutomationPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailDeliveryReceiptAutomationPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostTooManyRequests creates a EmailDeliveryReceiptAutomationPostTooManyRequests with default headers values
func NewEmailDeliveryReceiptAutomationPostTooManyRequests() *EmailDeliveryReceiptAutomationPostTooManyRequests {
	return &EmailDeliveryReceiptAutomationPostTooManyRequests{}
}

/*EmailDeliveryReceiptAutomationPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailDeliveryReceiptAutomationPostTooManyRequests struct {
	Payload interface{}
}

func (o *EmailDeliveryReceiptAutomationPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] emailDeliveryReceiptAutomationPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailDeliveryReceiptAutomationPostDefault creates a EmailDeliveryReceiptAutomationPostDefault with default headers values
func NewEmailDeliveryReceiptAutomationPostDefault(code int) *EmailDeliveryReceiptAutomationPostDefault {
	return &EmailDeliveryReceiptAutomationPostDefault{
		_statusCode: code,
	}
}

/*EmailDeliveryReceiptAutomationPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailDeliveryReceiptAutomationPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the email delivery receipt automation post default response
func (o *EmailDeliveryReceiptAutomationPostDefault) Code() int {
	return o._statusCode
}

func (o *EmailDeliveryReceiptAutomationPostDefault) Error() string {
	return fmt.Sprintf("[POST /automations/email/receipts][%d] EmailDeliveryReceiptAutomationPost default  %+v", o._statusCode, o.Payload)
}

func (o *EmailDeliveryReceiptAutomationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
