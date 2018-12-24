// Code generated by go-swagger; DO NOT EDIT.

package fax_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// FAXDeliveryReceiptAutomationPutReader is a Reader for the FAXDeliveryReceiptAutomationPut structure.
type FAXDeliveryReceiptAutomationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FAXDeliveryReceiptAutomationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFAXDeliveryReceiptAutomationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFAXDeliveryReceiptAutomationPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewFAXDeliveryReceiptAutomationPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFAXDeliveryReceiptAutomationPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFAXDeliveryReceiptAutomationPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFAXDeliveryReceiptAutomationPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewFAXDeliveryReceiptAutomationPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFAXDeliveryReceiptAutomationPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFAXDeliveryReceiptAutomationPutOK creates a FAXDeliveryReceiptAutomationPutOK with default headers values
func NewFAXDeliveryReceiptAutomationPutOK() *FAXDeliveryReceiptAutomationPutOK {
	return &FAXDeliveryReceiptAutomationPutOK{}
}

/*FAXDeliveryReceiptAutomationPutOK handles this case with default header values.

SUCCESS
*/
type FAXDeliveryReceiptAutomationPutOK struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutOK) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutOK  %+v", 200, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutBadRequest creates a FAXDeliveryReceiptAutomationPutBadRequest with default headers values
func NewFAXDeliveryReceiptAutomationPutBadRequest() *FAXDeliveryReceiptAutomationPutBadRequest {
	return &FAXDeliveryReceiptAutomationPutBadRequest{}
}

/*FAXDeliveryReceiptAutomationPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type FAXDeliveryReceiptAutomationPutBadRequest struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutBadRequest  %+v", 400, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutUnauthorized creates a FAXDeliveryReceiptAutomationPutUnauthorized with default headers values
func NewFAXDeliveryReceiptAutomationPutUnauthorized() *FAXDeliveryReceiptAutomationPutUnauthorized {
	return &FAXDeliveryReceiptAutomationPutUnauthorized{}
}

/*FAXDeliveryReceiptAutomationPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type FAXDeliveryReceiptAutomationPutUnauthorized struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutUnauthorized  %+v", 401, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutForbidden creates a FAXDeliveryReceiptAutomationPutForbidden with default headers values
func NewFAXDeliveryReceiptAutomationPutForbidden() *FAXDeliveryReceiptAutomationPutForbidden {
	return &FAXDeliveryReceiptAutomationPutForbidden{}
}

/*FAXDeliveryReceiptAutomationPutForbidden handles this case with default header values.

FORBIDDEN
*/
type FAXDeliveryReceiptAutomationPutForbidden struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutForbidden  %+v", 403, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutNotFound creates a FAXDeliveryReceiptAutomationPutNotFound with default headers values
func NewFAXDeliveryReceiptAutomationPutNotFound() *FAXDeliveryReceiptAutomationPutNotFound {
	return &FAXDeliveryReceiptAutomationPutNotFound{}
}

/*FAXDeliveryReceiptAutomationPutNotFound handles this case with default header values.

NOT_FOUND
*/
type FAXDeliveryReceiptAutomationPutNotFound struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutNotFound  %+v", 404, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutMethodNotAllowed creates a FAXDeliveryReceiptAutomationPutMethodNotAllowed with default headers values
func NewFAXDeliveryReceiptAutomationPutMethodNotAllowed() *FAXDeliveryReceiptAutomationPutMethodNotAllowed {
	return &FAXDeliveryReceiptAutomationPutMethodNotAllowed{}
}

/*FAXDeliveryReceiptAutomationPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type FAXDeliveryReceiptAutomationPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutTooManyRequests creates a FAXDeliveryReceiptAutomationPutTooManyRequests with default headers values
func NewFAXDeliveryReceiptAutomationPutTooManyRequests() *FAXDeliveryReceiptAutomationPutTooManyRequests {
	return &FAXDeliveryReceiptAutomationPutTooManyRequests{}
}

/*FAXDeliveryReceiptAutomationPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type FAXDeliveryReceiptAutomationPutTooManyRequests struct {
	Payload interface{}
}

func (o *FAXDeliveryReceiptAutomationPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] faxDeliveryReceiptAutomationPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXDeliveryReceiptAutomationPutDefault creates a FAXDeliveryReceiptAutomationPutDefault with default headers values
func NewFAXDeliveryReceiptAutomationPutDefault(code int) *FAXDeliveryReceiptAutomationPutDefault {
	return &FAXDeliveryReceiptAutomationPutDefault{
		_statusCode: code,
	}
}

/*FAXDeliveryReceiptAutomationPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type FAXDeliveryReceiptAutomationPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Fax delivery receipt automation put default response
func (o *FAXDeliveryReceiptAutomationPutDefault) Code() int {
	return o._statusCode
}

func (o *FAXDeliveryReceiptAutomationPutDefault) Error() string {
	return fmt.Sprintf("[PUT /automations/fax/receipts/{receipt_rule_id}][%d] FaxDeliveryReceiptAutomationPut default  %+v", o._statusCode, o.Payload)
}

func (o *FAXDeliveryReceiptAutomationPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
