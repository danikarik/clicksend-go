// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VoiceDeliveryReceiptAutomationPutReader is a Reader for the VoiceDeliveryReceiptAutomationPut structure.
type VoiceDeliveryReceiptAutomationPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceDeliveryReceiptAutomationPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceDeliveryReceiptAutomationPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceDeliveryReceiptAutomationPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceDeliveryReceiptAutomationPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceDeliveryReceiptAutomationPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceDeliveryReceiptAutomationPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceDeliveryReceiptAutomationPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceDeliveryReceiptAutomationPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceDeliveryReceiptAutomationPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceDeliveryReceiptAutomationPutOK creates a VoiceDeliveryReceiptAutomationPutOK with default headers values
func NewVoiceDeliveryReceiptAutomationPutOK() *VoiceDeliveryReceiptAutomationPutOK {
	return &VoiceDeliveryReceiptAutomationPutOK{}
}

/*VoiceDeliveryReceiptAutomationPutOK handles this case with default header values.

SUCCESS
*/
type VoiceDeliveryReceiptAutomationPutOK struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutOK) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutOK  %+v", 200, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutBadRequest creates a VoiceDeliveryReceiptAutomationPutBadRequest with default headers values
func NewVoiceDeliveryReceiptAutomationPutBadRequest() *VoiceDeliveryReceiptAutomationPutBadRequest {
	return &VoiceDeliveryReceiptAutomationPutBadRequest{}
}

/*VoiceDeliveryReceiptAutomationPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceDeliveryReceiptAutomationPutBadRequest struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutUnauthorized creates a VoiceDeliveryReceiptAutomationPutUnauthorized with default headers values
func NewVoiceDeliveryReceiptAutomationPutUnauthorized() *VoiceDeliveryReceiptAutomationPutUnauthorized {
	return &VoiceDeliveryReceiptAutomationPutUnauthorized{}
}

/*VoiceDeliveryReceiptAutomationPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceDeliveryReceiptAutomationPutUnauthorized struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutForbidden creates a VoiceDeliveryReceiptAutomationPutForbidden with default headers values
func NewVoiceDeliveryReceiptAutomationPutForbidden() *VoiceDeliveryReceiptAutomationPutForbidden {
	return &VoiceDeliveryReceiptAutomationPutForbidden{}
}

/*VoiceDeliveryReceiptAutomationPutForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceDeliveryReceiptAutomationPutForbidden struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutForbidden  %+v", 403, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutNotFound creates a VoiceDeliveryReceiptAutomationPutNotFound with default headers values
func NewVoiceDeliveryReceiptAutomationPutNotFound() *VoiceDeliveryReceiptAutomationPutNotFound {
	return &VoiceDeliveryReceiptAutomationPutNotFound{}
}

/*VoiceDeliveryReceiptAutomationPutNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationPutNotFound struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutNotFound  %+v", 404, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutMethodNotAllowed creates a VoiceDeliveryReceiptAutomationPutMethodNotAllowed with default headers values
func NewVoiceDeliveryReceiptAutomationPutMethodNotAllowed() *VoiceDeliveryReceiptAutomationPutMethodNotAllowed {
	return &VoiceDeliveryReceiptAutomationPutMethodNotAllowed{}
}

/*VoiceDeliveryReceiptAutomationPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationPutMethodNotAllowed struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutTooManyRequests creates a VoiceDeliveryReceiptAutomationPutTooManyRequests with default headers values
func NewVoiceDeliveryReceiptAutomationPutTooManyRequests() *VoiceDeliveryReceiptAutomationPutTooManyRequests {
	return &VoiceDeliveryReceiptAutomationPutTooManyRequests{}
}

/*VoiceDeliveryReceiptAutomationPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceDeliveryReceiptAutomationPutTooManyRequests struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPutDefault creates a VoiceDeliveryReceiptAutomationPutDefault with default headers values
func NewVoiceDeliveryReceiptAutomationPutDefault(code int) *VoiceDeliveryReceiptAutomationPutDefault {
	return &VoiceDeliveryReceiptAutomationPutDefault{
		_statusCode: code,
	}
}

/*VoiceDeliveryReceiptAutomationPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceDeliveryReceiptAutomationPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the voice delivery receipt automation put default response
func (o *VoiceDeliveryReceiptAutomationPutDefault) Code() int {
	return o._statusCode
}

func (o *VoiceDeliveryReceiptAutomationPutDefault) Error() string {
	return fmt.Sprintf("[PUT /automations/voice/receipts/{receipt_rule_id}][%d] VoiceDeliveryReceiptAutomationPut default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}